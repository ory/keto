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
	NamespaceWatcher struct {
		sync.RWMutex
		namespaces map[string]*namespace.Namespace
		ec         watcherx.EventChannel
		l          *logrusx.Logger
		target     string
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
		namespaces: make(map[string]*namespace.Namespace),
	}

	var w watcherx.Watcher

	info, err := os.Stat(u.Path)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if info.IsDir() {
		w, err = watcherx.WatchDirectory(ctx, u.Path, nw.ec)
	} else {
		w, err = watcherx.Watch(ctx, u, nw.ec)
	}
	// this handles the watcher init error
	if err != nil {
		return nil, err
	}

	// trigger initial load
	done, err := w.DispatchNow()
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

					n := parseNamespace(nw.l, e.Reader(), e.Source())
					if n == nil {
						return
					}

					nw.namespaces[e.Source()] = n
				}()
			case *watcherx.ErrorEvent:
				nw.l.WithError(etyped).Errorf("Received error while watching namespace files at target %s.", nw.target)
			}
		}
	}
}

func parseNamespace(l *logrusx.Logger, r io.Reader, source string) *namespace.Namespace {
	var parser func([]byte, interface{}) error

	knownFormats := stringsx.RegisteredCases{}
	switch ext := filepath.Ext(source); ext {
	case knownFormats.AddCase(".yaml"), knownFormats.AddCase(".yml"):
		parser = yaml.Unmarshal
	case knownFormats.AddCase(".json"):
		parser = json.Unmarshal
	case knownFormats.AddCase(".toml"):
		parser = toml.Unmarshal
	default:
		l.WithError(knownFormats.ToUnknownCaseErr(ext)).WithField("file_name", source).Warn("could not infer format from file extension")
		return nil
	}

	raw, err := ioutil.ReadAll(r)
	if err != nil {
		l.WithError(errors.WithStack(err)).WithField("file_name", source).Error("could not read namespace file")
		return nil
	}

	n := namespace.Namespace{}
	if err := parser(raw, &n); err != nil {
		l.WithError(errors.WithStack(err)).WithField("file_name", source).Error("could not parse namespace file")
		return nil
	}

	return &n
}

func (n *NamespaceWatcher) GetNamespace(_ context.Context, name string) (*namespace.Namespace, error) {
	n.RLock()
	defer n.RUnlock()

	for _, nspace := range n.namespaces {
		if nspace.Name == name {
			return nspace, nil
		}
	}

	return nil, errors.WithStack(herodot.ErrNotFound.WithError("unknown namespace " + name))
}

func (n *NamespaceWatcher) Namespaces(_ context.Context) ([]*namespace.Namespace, error) {
	n.RLock()
	defer n.RUnlock()

	nspaces := make([]*namespace.Namespace, 0, len(n.namespaces))
	for _, nspace := range n.namespaces {
		nspaces = append(nspaces, nspace)
	}

	return nspaces, nil
}
