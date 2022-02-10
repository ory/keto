package ketoctx

import (
	"net/http"

	"github.com/ory/x/logrusx"
	"google.golang.org/grpc"
)

type (
	opts struct {
		logger                 *logrusx.Logger
		contextualizer         Contextualizer
		httpMiddlewares        []func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
		grpcUnaryInterceptors  []grpc.UnaryServerInterceptor
		grpcStreamInterceptors []grpc.StreamServerInterceptor
	}
	Option func(o *opts)
)

func WithLogger(l *logrusx.Logger) Option {
	return func(o *opts) {
		o.logger = l
	}
}

func WithContextualizer(ctxer Contextualizer) Option {
	return func(o *opts) {
		o.contextualizer = ctxer
	}
}

func WithHTTPMiddlewares(m ...func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)) Option {
	return func(o *opts) {
		o.httpMiddlewares = m
	}
}

func WithGRPCUnaryInterceptors(i ...grpc.UnaryServerInterceptor) Option {
	return func(o *opts) {
		o.grpcUnaryInterceptors = i
	}
}

func WithGRPCStreamInterceptors(i ...grpc.StreamServerInterceptor) Option {
	return func(o *opts) {
		o.grpcStreamInterceptors = i
	}
}

func (o *opts) Logger() *logrusx.Logger {
	return o.logger
}

func (o *opts) Contextualizer() Contextualizer {
	return o.contextualizer
}

func (o *opts) HTTPMiddlewares() []func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	return o.httpMiddlewares
}

func (o *opts) GRPCUnaryInterceptors() []grpc.UnaryServerInterceptor {
	return o.grpcUnaryInterceptors
}

func (o *opts) GRPCStreamInterceptors() []grpc.StreamServerInterceptor {
	return o.grpcStreamInterceptors
}

func Options(options ...Option) *opts {
	o := &opts{
		contextualizer: &DefaultContextualizer{},
	}
	for _, opt := range options {
		opt(o)
	}
	return o
}
