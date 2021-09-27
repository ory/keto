package config

import (
	"context"
	"reflect"

	"github.com/ory/herodot"
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/namespace"
)

type (
	memoryNamespaceManager []*namespace.Namespace
)

var _ namespace.Manager = &memoryNamespaceManager{}

func NewMemoryNamespaceManager(nn ...*namespace.Namespace) *memoryNamespaceManager {
	nm := make(memoryNamespaceManager, len(nn))

	for i, np := range nn {
		n := *np
		nm[i] = &n
	}

	return &nm
}

func (s *memoryNamespaceManager) GetNamespaceByName(_ context.Context, name string) (*namespace.Namespace, error) {
	for _, n := range *s {
		if n.Name == name {
			return n, nil
		}
	}

	return nil, errors.WithStack(herodot.ErrNotFound.WithReasonf("Unknown namespace with name %s.", name))
}

func (s *memoryNamespaceManager) GetNamespaceByConfigID(_ context.Context, id int32) (*namespace.Namespace, error) {
	for _, n := range *s {
		if n.ID == id {
			return n, nil
		}
	}

	return nil, errors.WithStack(herodot.ErrNotFound.WithReasonf("Unknown namespace with id %d.", id))
}

func (s *memoryNamespaceManager) Namespaces(_ context.Context) ([]*namespace.Namespace, error) {
	nn := make([]*namespace.Namespace, 0, len(*s))

	for _, n := range *s {
		nc := *n
		nn = append(nn, &nc)
	}

	return nn, nil
}

func (s *memoryNamespaceManager) ShouldReload(newValue interface{}) bool {
	return !reflect.DeepEqual(newValue, []*namespace.Namespace(*s))
}
