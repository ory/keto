package config

import (
	"context"
	"io"
	"sync"

	"github.com/ory/x/logrusx"
	"github.com/ory/x/watcherx"

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

var _ namespace.Manager = (*oplConfigWatcher)(nil)

func newOPLConfigWatcher(ctx context.Context, l *logrusx.Logger, target string) (*oplConfigWatcher, error) {

	nw := &oplConfigWatcher{
		logger:                 l,
		target:                 target,
		files:                  configFiles{byPath: make(map[string]io.Reader)},
		memoryNamespaceManager: *NewMemoryNamespaceManager(),
	}

	return nw, watchTarget(ctx, target, nw, l)
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
		errs = append(errs, ee...)
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
