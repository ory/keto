// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/urlx"
	"github.com/ory/x/watcherx"
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/schema"
)

type (
	configFiles struct {
		byPath map[string]io.Reader
		sync.Mutex
	}

	oplConfigWatcher struct {
		logger *logrusx.Logger
		target string
		files  configFiles

		memoryNamespaceManager
	}
)

var (
	_        namespace.Manager = (*oplConfigWatcher)(nil)
	cache, _                   = ristretto.NewCache(&ristretto.Config{
		MaxCost:     20_000_000, // 20 MB max size, each item ca. 10 KB => max 2000 items
		NumCounters: 20_000,     // max 2000 items => 20000 counters
		BufferItems: 64,
	})
)

func newOPLConfigWatcher(ctx context.Context, c *Config, target string) (*oplConfigWatcher, error) {
	nw := &oplConfigWatcher{
		logger:                 c.l,
		target:                 target,
		files:                  configFiles{byPath: make(map[string]io.Reader)},
		memoryNamespaceManager: *NewMemoryNamespaceManager(),
	}

	targetUrl, err := urlx.Parse(target)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	switch targetUrl.Scheme {
	case "file", "":
		return nw, watchTarget(ctx, target, nw, c.l)
	case "base64":
		file, err := c.Fetcher().FetchContext(ctx, target)
		if err != nil {
			return nil, err
		}
		nw.files.byPath[targetUrl.String()] = file
		nw.parseFiles()
		return nw, err
	case "http", "https":
		var file io.Reader
		if item, ok := cache.Get(target); ok {
			file = bytes.NewReader(item.([]byte))
		} else {
			buf, err := c.Fetcher().FetchContext(ctx, target)
			if err != nil {
				return nil, err
			}
			b := buf.Bytes()
			cache.SetWithTTL(target, b, int64(cap(b)), 30*time.Minute)
			file = bytes.NewReader(b)
		}
		nw.files.byPath[targetUrl.String()] = file
		nw.parseFiles()
		return nw, err
	default:
		return nil, fmt.Errorf("unexpected url scheme: %q", targetUrl.Scheme)
	}
}

func (nw *oplConfigWatcher) handleChange(e *watcherx.ChangeEvent) {
	// the lock is acquired before parsing to ensure that the getters are
	// waiting for the updated values
	nw.files.Lock()
	defer nw.files.Unlock()
	nw.files.byPath[e.Source()] = e.Reader()
	nw.parseFiles()
}

func (nw *oplConfigWatcher) handleRemove(e *watcherx.RemoveEvent) {
	nw.files.Lock()
	defer nw.files.Unlock()
	delete(nw.files.byPath, e.Source())
	nw.parseFiles()
}

func (nw *oplConfigWatcher) handleError(e *watcherx.ErrorEvent) {
	nw.logger.
		WithError(e).
		Errorf("Received error while watching OPL config files at target %s.",
			nw.target)
}

// parseFiles loops through all files, parsing each and getting the namespaces.
// It then sets the namespaces only if there were no errors.
//
// The caller must  hold the lock to nw.files.
func (nw *oplConfigWatcher) parseFiles() {
	var (
		namespaces = make([]*namespace.Namespace, 0)
		errs       []error
	)
	for _, reader := range nw.files.byPath {
		content, err := io.ReadAll(reader)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		nn, ee := schema.Parse(string(content))
		for _, e := range ee {
			errs = append(errs, e)
		}
		for _, n := range nn {
			n := n // alias because we want a reference
			namespaces = append(namespaces, &n)
		}
	}
	if len(errs) > 0 {
		for _, err := range errs {
			nw.logger.
				WithError(err).
				Errorf("Failed to parse OPL config files at target %s.",
					nw.target)
		}
		return
	}
	nw.set(namespaces)
}
