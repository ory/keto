package ketoctx

import (
	"context"
	"net/http"

	"google.golang.org/grpc"

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
	MiddlewareContextualizer interface {
		ContextualizeHTTPMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
		ContextualizeGRPCUnaryMiddleware(ctx context.Context, req interface{}, serverInfo *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
		ContextualizeGRPCStreamMiddleware(srv interface{}, stream grpc.ServerStream, serverInfo *grpc.StreamServerInfo, handler grpc.StreamHandler) error
	}
	contextKey string
)

const (
	NetworkContextualizerKey    contextKey = "network contextualizer"
	ConfigContextualizerKey     contextKey = "config contextualizer"
	MiddlewareContextualizerKey contextKey = "middleware"
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

func ContextualizeMiddleware(ctx context.Context) MiddlewareContextualizer {
	if ctx == nil {
		panic("got unexpected nil context")
	}
	contextualizer, ok := ctx.Value(MiddlewareContextualizerKey).(MiddlewareContextualizer)
	if contextualizer == nil || !ok {
		panic("no middleware contextualizer found in context")
	}
	return contextualizer
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

func WithMiddlewareContextualizer(ctx context.Context, contextualizer MiddlewareContextualizer) context.Context {
	if ctx.Value(MiddlewareContextualizerKey) != nil {
		return ctx
	}
	return context.WithValue(ctx, MiddlewareContextualizerKey, contextualizer)
}
