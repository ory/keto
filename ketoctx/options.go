// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package ketoctx

import (
	"io/fs"
	"net/http"

	"github.com/ory/pop/v6"
	"github.com/ory/x/contextx"
	"github.com/ory/x/healthx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/otelx"
	"github.com/ory/x/popx"
	"google.golang.org/grpc"
)

type (
	Opts struct {
		logger                 *logrusx.Logger
		TracerWrapper          TracerWrapper
		contextualizer         contextx.Contextualizer
		httpMiddlewares        []func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
		grpcUnaryInterceptors  []grpc.UnaryServerInterceptor
		grpcStreamInterceptors []grpc.StreamServerInterceptor
		grpcServerOptions      []grpc.ServerOption
		migrationOpts          []popx.MigrationBoxOption
		readyCheckers          healthx.ReadyCheckers
		extraMigrations        []fs.FS
		inspect                InspectFunc
		dbOpts                 []func(details *pop.ConnectionDetails)
	}
	Option        func(o *Opts)
	TracerWrapper func(*otelx.Tracer) *otelx.Tracer
	InspectFunc   func(*pop.Connection) error
)

// WithDBOptionsModifier adds database connection options that will be applied to the
// underlying connection.
func WithDBOptionsModifier(mods ...func(details *pop.ConnectionDetails)) Option {
	return func(o *Opts) {
		o.dbOpts = append(o.dbOpts, mods...)
	}
}

// WithLogger sets the logger.
func WithLogger(l *logrusx.Logger) Option {
	return func(o *Opts) { o.logger = l }
}

// WithTracerWrapper sets a function that wraps the tracer.
func WithTracerWrapper(wrapper TracerWrapper) Option {
	return func(o *Opts) { o.TracerWrapper = wrapper }
}

// WithContextualizer sets the contextualizer.
func WithContextualizer(ctxer contextx.Contextualizer) Option {
	return func(o *Opts) {
		o.contextualizer = ctxer
	}
}

// WithHTTPMiddlewares adds HTTP middlewares to the list of HTTP middlewares.
func WithHTTPMiddlewares(m ...func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)) Option {
	return func(o *Opts) {
		o.httpMiddlewares = m
	}
}

// WithGRPCUnaryInterceptors adds gRPC unary interceptors to the list of gRPC
// interceptors.
func WithGRPCUnaryInterceptors(i ...grpc.UnaryServerInterceptor) Option {
	return func(o *Opts) {
		o.grpcUnaryInterceptors = i
	}
}

// WithGRPCStreamInterceptors adds gRPC stream interceptors to the list of gRPC
// stream interceptors.
func WithGRPCStreamInterceptors(i ...grpc.StreamServerInterceptor) Option {
	return func(o *Opts) {
		o.grpcStreamInterceptors = i
	}
}

// WithGRPCServerOptions adds gRPC server options.
func WithGRPCServerOptions(serverOpts ...grpc.ServerOption) Option {
	return func(o *Opts) {
		o.grpcServerOptions = serverOpts
	}
}

// WithExtraMigrations adds additional database migrations.
func WithExtraMigrations(o ...fs.FS) Option {
	return func(opts *Opts) {
		opts.extraMigrations = append(opts.extraMigrations, o...)
	}
}

// WithMigrationOptions adds migration options to the list of migration options.
func WithMigrationOptions(o ...popx.MigrationBoxOption) Option {
	return func(opts *Opts) {
		opts.migrationOpts = o
	}
}

// WithReadinessCheck adds a new readness health checker to the list of
// checkers. Can be called multiple times. If the name is already taken, the
// checker will be overwritten.
func WithReadinessCheck(name string, rc healthx.ReadyChecker) Option {
	return func(o *Opts) {
		if o.readyCheckers == nil {
			o.readyCheckers = make(healthx.ReadyCheckers)
		}
		o.readyCheckers[name] = rc
	}
}

func Inspect(f InspectFunc) Option {
	return func(o *Opts) {
		o.inspect = f
	}
}

func (o *Opts) Logger() *logrusx.Logger {
	return o.logger
}

func (o *Opts) Contextualizer() contextx.Contextualizer {
	return o.contextualizer
}

func (o *Opts) HTTPMiddlewares() []func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	return o.httpMiddlewares
}

func (o *Opts) GRPCUnaryInterceptors() []grpc.UnaryServerInterceptor {
	return o.grpcUnaryInterceptors
}

func (o *Opts) GRPCStreamInterceptors() []grpc.StreamServerInterceptor {
	return o.grpcStreamInterceptors
}

func (o *Opts) GRPCServerOptions() []grpc.ServerOption {
	return o.grpcServerOptions
}

func (o *Opts) ExtraMigrations() []fs.FS {
	return o.extraMigrations
}

func (o *Opts) MigrationOptions() []popx.MigrationBoxOption {
	return o.migrationOpts
}

func (o *Opts) ReadyCheckers() healthx.ReadyCheckers {
	return o.readyCheckers
}

func (o *Opts) Inspect() InspectFunc {
	return o.inspect
}

func (o *Opts) DBOptionsModifiers() []func(details *pop.ConnectionDetails) {
	return o.dbOpts
}

func Options(options ...Option) *Opts {
	o := &Opts{
		contextualizer: &contextx.Default{},
	}
	for _, opt := range options {
		opt(o)
	}
	return o
}
