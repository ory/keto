package config

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/ghodss/yaml"
	"github.com/ory/herodot"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/stringsx"
	"github.com/ory/x/urlx"
	"github.com/ory/x/watcherx"
	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/namespace"
)

type (
	Parser func([]byte, interface{}) error

	NamespaceFile struct {
		Name     string
		Contents []byte
		Parser   Parser

		namespace *namespace.Namespace // last successfully parsed namespace. May be nil if there was a parse error
	}

	NamespaceWatcher struct {
		sync.RWMutex
		namespaces map[string]*NamespaceFile
		ec         watcherx.EventChannel
		l          *logrusx.Logger
		target     string
		w          watcherx.Watcher
	}
)

var _ namespace.Manager = &NamespaceWatcher{}

func NewNamespaceWatcher(ctx context.Context, l *logrusx.Logger, target string) (*NamespaceWatcher, error) {
	u, err := urlx.Parse(target)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	nw := NamespaceWatcher{
		ec:         make(watcherx.EventChannel),
		l:          l,
		target:     target,
		namespaces: make(map[string]*NamespaceFile),
	}

	info, err := os.Stat(u.Path)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if info.IsDir() {
		nw.w, err = watcherx.WatchDirectory(ctx, u.Path, nw.ec)
	} else {
		nw.w, err = watcherx.Watch(ctx, u, nw.ec)
	}
	// this handles the watcher init error
	if err != nil {
		return nil, err
	}

	// trigger initial load
	done, err := nw.w.DispatchNow()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	initialEventsProcessed := make(chan struct{})
	go eventHandler(ctx, &nw, done, initialEventsProcessed)

	// wait for initial load to be done
	<-initialEventsProcessed

	return &nw, nil
}

func eventHandler(ctx context.Context, nw *NamespaceWatcher, done <-chan int, initialEventsProcessed chan<- struct{}) {
	for {
		select {
		// because we use an unbuffered chan we can be sure that at least all initial events are handled
		case <-done:
			close(initialEventsProcessed)
		case <-ctx.Done():
			return
		case e, open := <-nw.ec:
			if !open {
				return
			}

			switch etyped := e.(type) {
			case *watcherx.RemoveEvent:
				func() {
					nw.Lock()
					defer nw.Unlock()

					delete(nw.namespaces, e.Source())
				}()
			case *watcherx.ChangeEvent:
				// the lock is acquired before parsing to ensure that the getters are waiting for the updated values
				func() {
					nw.Lock()
					defer nw.Unlock()

					n := readNamespaceFile(nw.l, e.Reader(), e.Source())
					if n == nil {
						return
					} else if n.namespace == nil {
						// parse failed, rolling back to previous working version
						if existing, ok := nw.namespaces[e.Source()]; ok {
							existing.Contents = n.Contents
						} else {
							nw.namespaces[e.Source()] = n
						}
					} else {
						nw.namespaces[e.Source()] = n
					}
				}()
			case *watcherx.ErrorEvent:
				nw.l.WithError(etyped).Errorf("Received error while watching namespace files at target %s.", nw.target)
			}
		}
	}
}

func readNamespaceFile(l *logrusx.Logger, r io.Reader, source string) *NamespaceFile {
	var parse Parser
	parse, err := GetParser(source)
	if err != nil {
		l.WithError(err).WithField("file_name", source).Warn("could not infer format from file extension")
		return nil
	}

	raw, err := ioutil.ReadAll(r)
	if err != nil {
		l.WithError(errors.WithStack(err)).WithField("file_name", source).Error("could not read namespace file")
		return nil
	}

	n := namespace.Namespace{}
	if err := parse(raw, &n); err != nil {
		l.WithError(errors.WithStack(err)).WithField("file_name", source).Error("could not parse namespace file")
		return &NamespaceFile{Name: source, Contents: raw, Parser: parse}
	}

	return &NamespaceFile{Name: source, Contents: raw, Parser: parse, namespace: &n}
}

func (n *NamespaceWatcher) GetNamespaceByName(_ context.Context, name string) (*namespace.Namespace, error) {
	n.RLock()
	defer n.RUnlock()

	for _, nsf := range n.namespaces {
		if nsf.namespace != nil && nsf.namespace.Name == name {
			return nsf.namespace, nil
		}
	}

	return nil, errors.WithStack(herodot.ErrNotFound.WithErrorf("Unknown namespace with name %s", name))
}

func (n *NamespaceWatcher) GetNamespaceByConfigID(_ context.Context, id int64) (*namespace.Namespace, error) {
	n.RLock()
	defer n.RUnlock()

	for _, nspace := range n.namespaces {
		if nspace.namespace.ID == id {
			return nspace.namespace, nil
		}
	}

	return nil, errors.WithStack(herodot.ErrNotFound.WithErrorf("Unknown namespace with ID %d", id))
}

func (n *NamespaceWatcher) Namespaces(_ context.Context) ([]*namespace.Namespace, error) {
	n.RLock()
	defer n.RUnlock()

	nspaces := make([]*namespace.Namespace, 0, len(n.namespaces))
	for _, nsf := range n.namespaces {
		if nsf.namespace != nil {
			nspaces = append(nspaces, nsf.namespace)
		}
	}
	return nspaces, nil
}

func (n *NamespaceWatcher) NamespaceFiles() []*NamespaceFile {
	n.RLock()
	defer n.RUnlock()

	nsfs := make([]*NamespaceFile, 0, len(n.namespaces))
	for _, nsf := range n.namespaces {
		nsfs = append(nsfs, nsf)
	}
	return nsfs
}

func GetParser(fn string) (Parser, error) {
	knownFormats := stringsx.RegisteredCases{}
	switch ext := filepath.Ext(fn); ext {
	case knownFormats.AddCase(".yaml"), knownFormats.AddCase(".yml"):
		return yaml.Unmarshal, nil
	case knownFormats.AddCase(".json"):
		return json.Unmarshal, nil
	case knownFormats.AddCase(".toml"):
		return toml.Unmarshal, nil
	default:
		return nil, knownFormats.ToUnknownCaseErr(ext)
	}
}
