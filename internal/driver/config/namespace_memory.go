// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"context"
	"reflect"
	"sync"

	"github.com/ory/herodot"
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/namespace"
)

type (
	memoryNamespaceManager struct {
		byName map[string]*namespace.Namespace
		sync.RWMutex
	}
)

var _ namespace.Manager = &memoryNamespaceManager{}

func NewMemoryNamespaceManager(nn ...*namespace.Namespace) *memoryNamespaceManager {
	s := &memoryNamespaceManager{}
	s.set(nn)
	return s
}

func (s *memoryNamespaceManager) GetNamespaceByName(_ context.Context, name string) (*namespace.Namespace, error) {
	s.RLock()
	defer s.RUnlock()

	if n, ok := s.byName[name]; ok {
		return n, nil
	}

	return nil, errors.WithStack(herodot.ErrNotFound.WithReasonf("Unknown namespace with name %q.", name))
}

func (s *memoryNamespaceManager) GetNamespaceByConfigID(_ context.Context, id int32) (*namespace.Namespace, error) {
	s.RLock()
	defer s.RUnlock()

	for _, n := range s.byName {
		//lint:ignore SA1019 backwards compatibility
		//nolint:staticcheck
		if n.ID == id {
			return n, nil
		}
	}

	return nil, errors.WithStack(herodot.ErrNotFound.WithReasonf("Unknown namespace with id %d.", id))
}

func (s *memoryNamespaceManager) Namespaces(_ context.Context) ([]*namespace.Namespace, error) {
	s.RLock()
	defer s.RUnlock()

	nn := make([]*namespace.Namespace, 0, len(s.byName))
	for _, n := range s.byName {
		nn = append(nn, n)
	}

	return nn, nil
}

func (s *memoryNamespaceManager) ShouldReload(newValue interface{}) bool {
	s.RLock()
	defer s.RUnlock()

	nn, _ := s.Namespaces(context.Background())

	return !reflect.DeepEqual(newValue, nn)
}

func (s *memoryNamespaceManager) set(nn []*namespace.Namespace) {
	s.Lock()
	defer s.Unlock()

	s.byName = make(map[string]*namespace.Namespace, len(nn))
	for _, n := range nn {
		s.byName[n.Name] = n
	}
}
