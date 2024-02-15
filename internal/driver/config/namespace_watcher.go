// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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
		logger     *logrusx.Logger
		target     string
	}

	eventHandler interface {
		handleRemove(*watcherx.RemoveEvent)
		handleChange(*watcherx.ChangeEvent)
		handleError(*watcherx.ErrorEvent)
	}
)

var _ namespace.Manager = (*NamespaceWatcher)(nil)

func NewNamespaceWatcher(ctx context.Context, l *logrusx.Logger, target string) (*NamespaceWatcher, error) {
	nw := NamespaceWatcher{
		logger:     l,
		target:     target,
		namespaces: make(map[string]*NamespaceFile),
	}

	return &nw, watchTarget(ctx, target, &nw, l)
}

func (nw *NamespaceWatcher) handleRemove(e *watcherx.RemoveEvent) {
	nw.Lock()
	defer nw.Unlock()

	delete(nw.namespaces, e.Source())
}

func (nw *NamespaceWatcher) handleChange(e *watcherx.ChangeEvent) {
	// the lock is acquired before parsing to ensure that the getters are
	// waiting for the updated values
	nw.Lock()
	defer nw.Unlock()

	n := nw.readNamespaceFile(e.Reader(), e.Source())
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
}

func (nw *NamespaceWatcher) handleError(e *watcherx.ErrorEvent) {
	nw.logger.
		WithError(e).
		Errorf("Received error while watching namespace files at target %s.", nw.target)
}

func watchTarget(ctx context.Context, target string, handler eventHandler, log *logrusx.Logger) error {
	var (
		eventCh = make(watcherx.EventChannel)
		watcher watcherx.Watcher
	)

	targetUrl, err := urlx.Parse(target)
	if err != nil {
		return errors.WithStack(err)
	}
	info, err := os.Stat(targetUrl.Path)
	if err != nil {
		return errors.WithStack(err)
	}
	if info.IsDir() {
		watcher, err = watcherx.WatchDirectory(ctx, urlx.GetURLFilePath(targetUrl), eventCh)
	} else {
		watcher, err = watcherx.Watch(ctx, targetUrl, eventCh)
	}
	// this handles the watcher init error
	if err != nil {
		return err
	}

	// trigger initial load
	done, err := watcher.DispatchNow()
	if err != nil {
		return errors.WithStack(err)
	}

	initialEventsProcessed := make(chan struct{})
	go startEventHandler(ctx, eventCh, handler, done, initialEventsProcessed, log)

	// wait for initial load to be done
	<-initialEventsProcessed

	return nil
}

func startEventHandler(ctx context.Context,
	eventCh watcherx.EventChannel,
	handler eventHandler,
	done <-chan int,
	initialEventsProcessed chan<- struct{},
	log *logrusx.Logger) {

	initalDone := false
	for {
		select {
		// because we use an unbuffered chan we can be sure that at least all
		// initial events are handled
		case <-done:
			initalDone = true
			close(initialEventsProcessed)

		case <-ctx.Done():
			return

		case e, open := <-eventCh:
			if !open {
				return
			}

			if initalDone {
				log.
					WithField("file", e.Source()).
					WithField("event_type", fmt.Sprintf("%T", e)).
					Info("A change to a namespace file was detected.")
			}

			switch e := e.(type) {
			case *watcherx.RemoveEvent:
				handler.handleRemove(e)
			case *watcherx.ChangeEvent:
				handler.handleChange(e)
			case *watcherx.ErrorEvent:
				handler.handleError(e)
			default:
				log.Warnf("Ignored unknown event %T", e)
			}
		}
	}
}

func (nw *NamespaceWatcher) readNamespaceFile(r io.Reader, source string) *NamespaceFile {
	parse, err := GetParser(source)
	if err != nil {
		nw.logger.
			WithError(err).
			WithField("file_name", source).
			Warn("could not infer format from file extension")
		return nil
	}

	raw, err := io.ReadAll(r)
	if err != nil {
		nw.logger.
			WithError(errors.WithStack(err)).
			WithField("file_name", source).
			Error("could not read namespace file")
		return nil
	}

	n := namespace.Namespace{}
	if err := parse(raw, &n); err != nil {
		nw.logger.
			WithError(errors.WithStack(err)).
			WithField("file_name", source).
			Error("could not parse namespace file")
		return &NamespaceFile{Name: source, Contents: raw, Parser: parse}
	}

	return &NamespaceFile{Name: source, Contents: raw, Parser: parse, namespace: &n}
}

func (nw *NamespaceWatcher) GetNamespaceByName(_ context.Context, name string) (*namespace.Namespace, error) {
	nw.RLock()
	defer nw.RUnlock()

	for _, nsf := range nw.namespaces {
		if nsf.namespace != nil && nsf.namespace.Name == name {
			return nsf.namespace, nil
		}
	}

	return nil, errors.WithStack(herodot.ErrNotFound.WithErrorf(
		"Unknown namespace with name %s", name))
}

func (nw *NamespaceWatcher) GetNamespaceByConfigID(_ context.Context, id int32) (*namespace.Namespace, error) {
	nw.RLock()
	defer nw.RUnlock()

	for _, nspace := range nw.namespaces {
		//lint:ignore SA1019 backwards compatibility
		//nolint:staticcheck
		if nspace.namespace.ID == id {
			return nspace.namespace, nil
		}
	}

	return nil, errors.WithStack(herodot.ErrNotFound.WithErrorf(
		"Unknown namespace with ID %d", id))
}

func (nw *NamespaceWatcher) Namespaces(_ context.Context) ([]*namespace.Namespace, error) {
	nw.RLock()
	defer nw.RUnlock()

	nspaces := make([]*namespace.Namespace, 0, len(nw.namespaces))
	for _, nsf := range nw.namespaces {
		if nsf.namespace != nil {
			nspaces = append(nspaces, nsf.namespace)
		}
	}
	return nspaces, nil
}

func (nw *NamespaceWatcher) NamespaceFiles() []*NamespaceFile {
	nw.RLock()
	defer nw.RUnlock()

	nsfs := make([]*NamespaceFile, 0, len(nw.namespaces))
	for _, nsf := range nw.namespaces {
		nsfs = append(nsfs, nsf)
	}
	return nsfs
}

func (nw *NamespaceWatcher) ShouldReload(newValue interface{}) bool {
	v, ok := newValue.(string)
	if !ok {
		// the manager type changed
		return true
	}
	// reload if target changed
	return v != nw.target
}

func GetParser(fn string) (Parser, error) {
	switch ext := stringsx.SwitchExact(filepath.Ext(fn)); {
	case ext.AddCase(".yaml"), ext.AddCase(".yml"):
		return func(b []byte, i interface{}) error {
			return yaml.Unmarshal(b, i)
		}, nil
	case ext.AddCase(".json"):
		return json.Unmarshal, nil
	case ext.AddCase(".toml"):
		return toml.Unmarshal, nil
	default:
		return nil, ext.ToUnknownCaseErr()
	}
}
