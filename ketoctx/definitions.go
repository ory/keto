package ketoctx

import (
	"context"

	"github.com/ory/x/configx"

	"github.com/gofrs/uuid"
)

type (
	NetworkContextualizer interface {
		ContextualizeNetwork(ctx context.Context) uuid.UUID
	}
	ConfigContextualizer interface {
		ContextualizeConfig(ctx context.Context) *configx.Provider
	}
	contextKey string
)

const (
	NetworkContextualizerKey contextKey = "network contextualizer"
	ConfigContextualizerKey  contextKey = "config contextualizer"
)

func ContextualizeNetwork(ctx context.Context) uuid.UUID {
	if ctx == nil {
		panic("got unexpected nil context")
	}
	contextualizer, ok := ctx.Value(NetworkContextualizerKey).(NetworkContextualizer)
	if contextualizer == nil || !ok {
		panic("no network contextualizer found in context")
	}
	return contextualizer.ContextualizeNetwork(ctx)
}

func ContextualizeConfig(ctx context.Context) *configx.Provider {
	if ctx == nil {
		panic("got unexpected nil context")
	}
	contextualizer, ok := ctx.Value(ConfigContextualizerKey).(ConfigContextualizer)
	if contextualizer == nil || !ok {
		panic("no config contextualizer found in context")
	}
	return contextualizer.ContextualizeConfig(ctx)
}

func WithNetworkContextualizer(ctx context.Context, contextualizer NetworkContextualizer) context.Context {
	if ctx.Value(NetworkContextualizerKey) != nil {
		return ctx
	}
	return context.WithValue(ctx, NetworkContextualizerKey, contextualizer)
}

func WithConfigContextualizer(ctx context.Context, contextualizer ConfigContextualizer) context.Context {
	if ctx.Value(ConfigContextualizerKey) != nil {
		return ctx
	}
	return context.WithValue(ctx, ConfigContextualizerKey, contextualizer)
}
