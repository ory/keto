package config

import (
	"context"
	"github.com/ory/herodot"
	"github.com/ory/keto/internal/namespace"
	"github.com/pkg/errors"
)

type (
	staticNamespaces []*namespace.Namespace
)

var _ namespace.Manager = &staticNamespaces{}

func (s *staticNamespaces) GetNamespace(_ context.Context, name string) (*namespace.Namespace, error) {
	for _, n := range *s {
		if n.Name == name {
			return n, nil
		}
	}

	return nil, errors.WithStack(herodot.ErrNotFound)
}

func (s *staticNamespaces) Namespaces(_ context.Context) ([]*namespace.Namespace, error) {
	nn := make([]*namespace.Namespace, 0, len(*s))

	for _, n := range *s {
		nc := *n
		nn = append(nn, &nc)
	}

	return nn, nil
}
