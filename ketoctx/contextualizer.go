// Copyright © 2022 Ory Corp

package ketoctx

import (
	"context"

	"github.com/ory/x/configx"

	"github.com/gofrs/uuid"
)

type (
	Contextualizer interface {
		Network(ctx context.Context, network uuid.UUID) uuid.UUID
		Config(ctx context.Context, config *configx.Provider) *configx.Provider
	}
	ContextualizerProvider interface {
		Contextualizer() Contextualizer
	}
	DefaultContextualizer struct{}
)

var _ Contextualizer = (*DefaultContextualizer)(nil)

func (d *DefaultContextualizer) Network(_ context.Context, network uuid.UUID) uuid.UUID {
	return network
}

func (d *DefaultContextualizer) Config(_ context.Context, config *configx.Provider) *configx.Provider {
	return config
}
