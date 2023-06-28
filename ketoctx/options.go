// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package ketoctx

import (
	"io/fs"
	"net/http"

	"github.com/ory/x/healthx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/otelx"
	"github.com/ory/x/popx"
	"google.golang.org/grpc"
)

type (
	opts struct {
		logger                 *logrusx.Logger
		TracerWrapper          TracerWrapper
		contextualizer         Contextualizer
		httpMiddlewares        []func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
		grpcUnaryInterceptors  []grpc.UnaryServerInterceptor
		grpcStreamInterceptors []grpc.StreamServerInterceptor
		migrationOpts          []popx.MigrationBoxOption
		readyCheckers          healthx.ReadyCheckers
		extraMigrations        []fs.FS
	}
	Option        func(o *opts)
	TracerWrapper func(*otelx.Tracer) *otelx.Tracer
)

// WithLogger sets the logger.
func WithLogger(l *logrusx.Logger) Option {
	return func(o *opts) { o.logger = l }
}

// WithTracerWrapper sets a function that wraps the tracer.
func WithTracerWrapper(wrapper TracerWrapper) Option {
	return func(o *opts) { o.TracerWrapper = wrapper }
}

// WithContextualizer sets the contextualizer.
func WithContextualizer(ctxer Contextualizer) Option {
	return func(o *opts) {
		o.contextualizer = ctxer
	}
}

// WithHTTPMiddlewares adds HTTP middlewares to the list of HTTP middlewares.
func WithHTTPMiddlewares(m ...func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)) Option {
	return func(o *opts) {
		o.httpMiddlewares = m
	}
}

// WithGRPCUnaryInterceptors adds gRPC unary interceptors to the list of gRPC
// interceptors.
func WithGRPCUnaryInterceptors(i ...grpc.UnaryServerInterceptor) Option {
	return func(o *opts) {
		o.grpcUnaryInterceptors = i
	}
}

// WithGRPCStreamInterceptors adds gRPC stream interceptors to the list of gRPC
// stream interceptors.
func WithGRPCStreamInterceptors(i ...grpc.StreamServerInterceptor) Option {
	return func(o *opts) {
		o.grpcStreamInterceptors = i
	}
}

// WithExtraMigrations adds additional database migrations.
func WithExtraMigrations(o ...fs.FS) Option {
	return func(opts *opts) {
		opts.extraMigrations = append(opts.extraMigrations, o...)
	}
}

// WithMigrationOptions adds migration options to the list of migration options.
func WithMigrationOptions(o ...popx.MigrationBoxOption) Option {
	return func(opts *opts) {
		opts.migrationOpts = o
	}
}

// WithReadinessCheck adds a new readness health checker to the list of
// checkers. Can be called multiple times. If the name is already taken, the
// checker will be overwritten.
func WithReadinessCheck(name string, rc healthx.ReadyChecker) Option {
	return func(o *opts) {
		if o.readyCheckers == nil {
			o.readyCheckers = make(healthx.ReadyCheckers)
		}
		o.readyCheckers[name] = rc
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

func (o *opts) ExtraMigrations() []fs.FS {
	return o.extraMigrations
}

func (o *opts) MigrationOptions() []popx.MigrationBoxOption {
	return o.migrationOpts
}

func (o *opts) ReadyCheckers() healthx.ReadyCheckers {
	return o.readyCheckers
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
