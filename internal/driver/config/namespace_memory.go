package config

import (
	"context"

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

func (s *memoryNamespaceManager) GetNamespace(_ context.Context, name string) (*namespace.Namespace, error) {
	for _, n := range *s {
		if n.Name == name {
			return n, nil
		}
	}

	return nil, errors.WithStack(herodot.ErrNotFound)
}

func (s *memoryNamespaceManager) Namespaces(_ context.Context) ([]*namespace.Namespace, error) {
	nn := make([]*namespace.Namespace, 0, len(*s))

	for _, n := range *s {
		nc := *n
		nn = append(nn, &nc)
	}

	return nn, nil
}
