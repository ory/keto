// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"syscall"
	"time"

	grpcOtel "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"

	"github.com/ory/x/otelx/semconv"

	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ory/keto/internal/namespace/namespacehandler"
	"github.com/ory/keto/internal/schema"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/ory/x/logrusx"
	prometheus "github.com/ory/x/prometheusx"

	grpcLogrus "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/soheilhy/cmux"
	"github.com/spf13/cobra"
	"github.com/urfave/negroni"
	grpcOtel "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"

	"github.com/ory/analytics-go/v5"
	"github.com/ory/graceful"
	"github.com/ory/herodot"
	"github.com/ory/x/healthx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/metricsx"
	"github.com/ory/x/otelx"
	"github.com/ory/x/otelx/semconv"
	prometheus "github.com/ory/x/prometheusx"
	"github.com/ory/x/reqlog"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/namespace/namespacehandler"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/schema"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/internal/x/api"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

func (r *RegistryDefault) enableSqa(cmd *cobra.Command) {
	ctx := cmd.Context()

	r.sqaService = metricsx.New(
		cmd,
		r.Logger(),
		r.Config(ctx).Source(),
		&metricsx.Options{
			Service:       "keto",
			DeploymentId:  metricsx.Hash(r.Persister().NetworkID(ctx).String()),
			IsDevelopment: strings.HasPrefix(r.Config(ctx).DSN(), "sqlite"),
			WriteKey:      "jk32cFATnj9GKbQdFL7fBB9qtKZdX9j7",
			WhitelistedPaths: []string{
				"/",
				healthx.AliveCheckPath,
				healthx.ReadyCheckPath,
				healthx.VersionPath,

				relationtuple.ReadRouteBase,
				check.RouteBase,
				expand.RouteBase,
			},
			BuildVersion: config.Version,
			BuildHash:    config.Commit,
			BuildTime:    config.Date,
			Config: &analytics.Config{
				Endpoint:             "https://sqa.ory.sh",
				GzipCompressionLevel: 6,
				BatchMaxSize:         500 * 1000,
				BatchSize:            1000,
				Interval:             time.Hour * 6,
			},
			Hostname: "", // TODO: figure out config to use
		},
	)
}

func (r *RegistryDefault) ServeAllSQA(cmd *cobra.Command) error {
	r.enableSqa(cmd)
	return r.ServeAll(cmd.Context())
}

func (r *RegistryDefault) ServeAll(ctx context.Context) error {
	innerCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	serveFuncs := []func(context.Context, chan<- struct{}) func() error{
		r.serveRead,
		r.serveWrite,
		r.serveOPLSyntax,
		r.serveMetrics,
	}

	doneShutdown := make(chan struct{}, len(serveFuncs))

	go func() {
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

		select {
		case <-osSignals:
			cancel()
		case <-innerCtx.Done():
		}

		ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), graceful.DefaultShutdownTimeout)
		defer cancel()

		nWaitingForShutdown := len(serveFuncs)
		for {
			select {
			case <-ctx.Done():
				return
			case <-doneShutdown:
				nWaitingForShutdown--
				if nWaitingForShutdown == 0 {
					// graceful shutdown done
					return
				}
			}
		}
	}()

	eg := &errgroup.Group{}

	// We need to separate the setup (invoking the functions that return the serve functions) from running the serve
	// functions to mitigate race conditions in the HTTP router.
	for _, serve := range serveFuncs {
		eg.Go(serve(innerCtx, doneShutdown))
	}

	return eg.Wait()
}

func (r *RegistryDefault) serveRead(ctx context.Context, done chan<- struct{}) func() error {
	server := r.newAPIServer(ctx)
	r.registerCommonGRPCServices(server.GRPCServer)
	r.registerReadGRPCServices(server.GRPCServer)

	apiHandler, err := server.Handler()
	if err != nil {
		return func() error { return err }
	}

	handler := r.ReadRouter(ctx, apiHandler)

	if tracer := r.Tracer(ctx); tracer.IsLoaded() {
		handler = otelx.TraceHandler(handler, otelhttp.WithTracerProvider(tracer.Provider()))
	}

	return func() error {
		addr, listenFile := r.Config(ctx).ReadAPIListenOn()
		return serve(ctx, r.Logger().WithField("endpoint", "read"), addr, listenFile, handler, done)
	}
}

func (r *RegistryDefault) serveWrite(ctx context.Context, done chan<- struct{}) func() error {
	server := r.newAPIServer(ctx)
	r.registerCommonGRPCServices(server.GRPCServer)
	r.registerWriteGRPCServices(server.GRPCServer)

	apiHandler, err := server.Handler()
	if err != nil {
		return func() error { return err }
	}

	handler := r.WriteRouter(ctx, apiHandler)

	if tracer := r.Tracer(ctx); tracer.IsLoaded() {
		handler = otelx.TraceHandler(handler, otelhttp.WithTracerProvider(tracer.Provider()))
	}

	return func() error {
		addr, listenFile := r.Config(ctx).WriteAPIListenOn()
		return serve(ctx, r.Logger().WithField("endpoint", "write"), addr, listenFile, handler, done)
	}
}

func (r *RegistryDefault) serveOPLSyntax(ctx context.Context, done chan<- struct{}) func() error {
	server := r.newAPIServer(ctx)
	r.registerCommonGRPCServices(server.GRPCServer)
	r.registerOPLGRPCServices(server.GRPCServer)

	apiHandler, err := server.Handler()
	if err != nil {
		return func() error { return err }
	}

	handler := r.OPLSyntaxRouter(ctx, apiHandler)

	if tracer := r.Tracer(ctx); tracer.IsLoaded() {
		handler = otelx.TraceHandler(handler, otelhttp.WithTracerProvider(tracer.Provider()))
	}

	return func() error {
		addr, listenFile := r.Config(ctx).OPLSyntaxAPIListenOn()
		return serve(ctx, r.Logger().WithField("endpoint", "opl"), addr, listenFile, handler, done)
	}
}

func (r *RegistryDefault) serveMetrics(ctx context.Context, done chan<- struct{}) func() error {
	ctx, cancel := context.WithCancel(ctx)

	//nolint:gosec // graceful.WithDefaults already sets a timeout
	s := graceful.WithDefaults(&http.Server{
		Handler: r.metricsRouter(ctx),
	})

	return func() error {
		defer cancel()
		eg := &errgroup.Group{}

		addr, listenFile := r.Config(ctx).MetricsListenOn()
		l, err := listenAndWriteFile(ctx, addr, listenFile)
		if err != nil {
			return err
		}

		eg.Go(func() error {
			if err := s.Serve(l); !errors.Is(err, http.ErrServerClosed) {
				return errors.WithStack(err)
			}
			return nil
		})
		eg.Go(func() (err error) {
			defer func() {
				l := r.Logger().WithField("endpoint", "metrics")
				if err != nil {
					l.WithError(err).Error("graceful shutdown failed")
				} else {
					l.Info("gracefully shutdown server")
				}
				done <- struct{}{}
			}()

			<-ctx.Done()
			ctx, cancel := context.WithTimeout(context.Background(), graceful.DefaultShutdownTimeout)
			defer cancel()
			return s.Shutdown(ctx)
		})

		return eg.Wait()
	}
}

func serve(ctx context.Context, log *logrusx.Logger, addr string, listenFile string, handler http.Handler, done chan<- struct{}) error {
	l, err := listenAndWriteFile(ctx, addr, listenFile)
	if err != nil {
		return err
	}

	//nolint:gosec // graceful.WithDefaults already sets a timeout
	server := graceful.WithDefaults(&http.Server{
		Handler: http.MaxBytesHandler(h2c.NewHandler(handler, &http2.Server{}), 1024*1024*4),
	})

	eg := &errgroup.Group{}

	eg.Go(func() error {
		if err := server.Serve(l); !errors.Is(err, http.ErrServerClosed) && !errors.Is(err, cmux.ErrServerClosed) {
			return errors.WithStack(err)
		}
		return nil
	})

	eg.Go(func() (err error) {
		defer func() {
			if err != nil {
				log.WithError(err).Error("graceful shutdown failed")
			} else {
				log.Info("gracefully shutdown server")
			}
			done <- struct{}{}
		}()

		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), graceful.DefaultShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			return errors.WithStack(err)
		}
		return nil
	})

	return eg.Wait()
}

func (r *RegistryDefault) allHandlers() []Handler {
	if len(r.handlers) == 0 {
		r.handlers = []Handler{
			relationtuple.NewHandler(r),
			check.NewHandler(r),
			expand.NewHandler(r),
			namespacehandler.New(r),
			schema.NewHandler(r),
		}
	}
	return r.handlers
}

func listenAndWriteFile(ctx context.Context, addr, listenFile string) (net.Listener, error) {
	l, err := (&net.ListenConfig{}).Listen(ctx, "tcp", addr)
	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("unable to listen on %q: %w", addr, err))
	}
	const filePrefix = "file://"
	if strings.HasPrefix(listenFile, filePrefix) {
		if err := os.WriteFile(listenFile[len(filePrefix):], []byte(l.Addr().String()), 0o600); err != nil {
			return nil, errors.WithStack(fmt.Errorf("unable to write listen file %q: %w", listenFile, err))
		}
	}
	return l, nil
}

type RouterOrHandler struct {
	Router  *httprouter.Router
	Handler http.Handler
}

func (h *RouterOrHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if handle, params, _ := h.Router.Lookup(r.Method, r.URL.Path); handle != nil {
		handle(rw, r, params)
		return
	}
	h.Handler.ServeHTTP(rw, r)
}

func (r *RegistryDefault) ReadRouter(ctx context.Context, apiHandler http.Handler) http.Handler {
	n := negroni.New()
	for _, f := range r.defaultHttpMiddlewares {
		n.UseFunc(f)
	}
	n.UseFunc(semconv.Middleware)
	n.Use(reqlog.NewMiddlewareFromLogger(r.l, "read#Ory Keto").ExcludePaths(healthx.AliveCheckPath, healthx.ReadyCheckPath))

	br := &x.ReadRouter{Router: httprouter.New()}
	r.PrometheusManager().RegisterRouter(br.Router)
	r.MetricsHandler().SetRoutes(br.Router)

	r.HealthHandler().SetHealthRoutes(br.Router, false)
	r.HealthHandler().SetVersionRoutes(br.Router)

	n.UseHandler(&RouterOrHandler{Router: br.Router, Handler: apiHandler})
	n.Use(r.PrometheusManager())

	if r.sqaService != nil {
		n.Use(r.sqaService)
	}

	var handler http.Handler = n
	options, enabled := r.Config(ctx).CORS("read")
	if enabled {
		handler = cors.New(options).Handler(handler)
	}

	return handler
}

func (r *RegistryDefault) WriteRouter(ctx context.Context, apiHandler http.Handler) http.Handler {
	n := negroni.New()
	for _, f := range r.defaultHttpMiddlewares {
		n.UseFunc(f)
	}
	n.UseFunc(semconv.Middleware)
	n.Use(reqlog.NewMiddlewareFromLogger(r.l, "write#Ory Keto").ExcludePaths(healthx.AliveCheckPath, healthx.ReadyCheckPath))

	pr := &x.WriteRouter{Router: httprouter.New()}
	r.PrometheusManager().RegisterRouter(pr.Router)
	r.MetricsHandler().SetRoutes(pr.Router)

	r.HealthHandler().SetHealthRoutes(pr.Router, false)
	r.HealthHandler().SetVersionRoutes(pr.Router)

	n.UseHandler(&RouterOrHandler{Router: pr.Router, Handler: apiHandler})
	n.Use(r.PrometheusManager())

	if r.sqaService != nil {
		n.Use(r.sqaService)
	}

	var handler http.Handler = n
	options, enabled := r.Config(ctx).CORS("write")
	if enabled {
		handler = cors.New(options).Handler(handler)
	}

	return handler
}

func (r *RegistryDefault) OPLSyntaxRouter(ctx context.Context, apiHandler http.Handler) http.Handler {
	n := negroni.New()
	for _, f := range r.defaultHttpMiddlewares {
		n.UseFunc(f)
	}
	n.UseFunc(semconv.Middleware)
	n.Use(reqlog.NewMiddlewareFromLogger(r.l, "syntax#Ory Keto").ExcludePaths(healthx.AliveCheckPath, healthx.ReadyCheckPath))
	n.UseFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if r.URL.Path == "/opl/syntax/check" && r.Header.Get("Content-Type") != "application/json" {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			newBody, err := json.Marshal(body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			r.Body = io.NopCloser(bytes.NewReader(newBody))
			r.ContentLength = int64(len(newBody))
			r.Header.Set("Content-Type", "application/json")
		}
		next(w, r)
	})

	pr := &x.OPLSyntaxRouter{Router: httprouter.New()}
	r.PrometheusManager().RegisterRouter(pr.Router)
	r.MetricsHandler().SetRoutes(pr.Router)

	r.HealthHandler().SetHealthRoutes(pr.Router, false)
	r.HealthHandler().SetVersionRoutes(pr.Router)

	n.UseHandler(&RouterOrHandler{Router: pr.Router, Handler: apiHandler})
	n.Use(r.PrometheusManager())

	if r.sqaService != nil {
		n.Use(r.sqaService)
	}

	var handler http.Handler = n
	options, enabled := r.Config(ctx).CORS("write")
	if enabled {
		handler = cors.New(options).Handler(handler)
	}

	return handler
}

func (r *RegistryDefault) grpcRecoveryHandler(p interface{}) error {
	r.Logger().
		WithField("reason", p).
		WithField("stack_trace", string(debug.Stack())).
		WithField("handler", "rate_limit").
		Error("panic recovered")
	return status.Errorf(codes.Internal, "%v", p)
}

func (r *RegistryDefault) unaryInterceptors(ctx context.Context) []grpc.UnaryServerInterceptor {
	is := []grpc.UnaryServerInterceptor{
		grpcRecovery.UnaryServerInterceptor(grpcRecovery.WithRecoveryHandler(r.grpcRecoveryHandler)),
	}
	is = append(is, r.defaultUnaryInterceptors...)
	is = append(is,
		herodot.UnaryErrorUnwrapInterceptor,
		x.ValidationInterceptor,
		grpcLogrus.UnaryServerInterceptor(InterceptorLogger(r.l.Logrus())),
		r.pmm.UnaryServerInterceptor,
	)
	if r.sqaService != nil {
		is = append(is, r.sqaService.UnaryInterceptor)
	}
	return is
}

func (r *RegistryDefault) streamInterceptors(ctx context.Context) []grpc.StreamServerInterceptor {
	is := []grpc.StreamServerInterceptor{
		grpcRecovery.StreamServerInterceptor(grpcRecovery.WithRecoveryHandler(r.grpcRecoveryHandler)),
	}
	is = append(is, r.defaultStreamInterceptors...)
	is = append(is,
		herodot.StreamErrorUnwrapInterceptor,
		grpcLogrus.StreamServerInterceptor(InterceptorLogger(r.l.Logrus())),
		r.pmm.StreamServerInterceptor,
	)
	if r.sqaService != nil {
		is = append(is, r.sqaService.StreamInterceptor)
	}
	return is
}

func (r *RegistryDefault) newGrpcServer(ctx context.Context) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.ChainStreamInterceptor(r.streamInterceptors(ctx)...),
		grpc.ChainUnaryInterceptor(r.unaryInterceptors(ctx)...),
	}
	if r.Tracer(ctx).IsLoaded() {
		opts = append(opts,
			grpc.StatsHandler(grpcOtel.NewServerHandler(grpcOtel.WithTracerProvider(otel.GetTracerProvider()))),
		)
	}

	opts = append(opts, r.defaultGRPCServerOptions...)
	if r.grpcTransportCredentials != nil {
		opts = append(opts, grpc.Creds(r.grpcTransportCredentials))
	}

	return grpc.NewServer(opts...)
}

func (r *RegistryDefault) newAPIServer(ctx context.Context) *api.Server {
	opts := []api.ServerOption{
		api.WithGRPCOption(grpc.ChainStreamInterceptor(r.streamInterceptors(ctx)...)),
		api.WithGRPCOption(grpc.ChainUnaryInterceptor(r.unaryInterceptors(ctx)...)),
		api.WithGRPCOption(grpc.MaxRecvMsgSize(1024 * 1024 * 8)),
	}
	for _, o := range r.defaultGRPCServerOptions {
		opts = append(opts, api.WithGRPCOption(o))
	}
	if r.grpcTransportCredentials != nil {
		opts = append(opts, api.WithGRPCOption(grpc.Creds(r.grpcTransportCredentials)))
	}

	return api.NewServer(opts...)
}

func (r *RegistryDefault) registerCommonGRPCServices(s *grpc.Server) {
	grpcHealthV1.RegisterHealthServer(s, r.HealthServer())
	rts.RegisterVersionServiceServer(s, r)
	r.pmm.Register(s)
}

func (r *RegistryDefault) registerReadGRPCServices(s *grpc.Server) {
	for _, h := range r.allHandlers() {
		if h, ok := h.(ReadHandler); ok {
			h.RegisterReadGRPC(s)
		}
	}
}

func (r *RegistryDefault) registerWriteGRPCServices(s *grpc.Server) {
	for _, h := range r.allHandlers() {
		if h, ok := h.(WriteHandler); ok {
			h.RegisterWriteGRPC(s)
		}
	}
}

func (r *RegistryDefault) registerOPLGRPCServices(s *grpc.Server) {
	for _, h := range r.allHandlers() {
		if h, ok := h.(OPLSyntaxHandler); ok {
			h.RegisterSyntaxGRPC(s)
		}
	}
}

func (r *RegistryDefault) ReadGRPCServer(ctx context.Context) *grpc.Server {
	s := r.newGrpcServer(ctx)
	r.registerCommonGRPCServices(s)
	r.registerReadGRPCServices(s)

	return s
}

func (r *RegistryDefault) WriteGRPCServer(ctx context.Context) *grpc.Server {
	s := r.newGrpcServer(ctx)
	r.registerCommonGRPCServices(s)
	r.registerWriteGRPCServices(s)

	return s
}

func (r *RegistryDefault) OplGRPCServer(ctx context.Context) *grpc.Server {
	s := r.newGrpcServer(ctx)
	r.registerCommonGRPCServices(s)
	r.registerOPLGRPCServices(s)

	return s
}

func (r *RegistryDefault) metricsRouter(ctx context.Context) http.Handler {
	n := negroni.New(reqlog.NewMiddlewareFromLogger(r.Logger(), "keto").ExcludePaths(prometheus.MetricsPrometheusPath))
	router := httprouter.New()

	r.PrometheusManager().RegisterRouter(router)
	r.MetricsHandler().SetRoutes(router)
	n.UseHandler(router)
	n.Use(r.PrometheusManager())

	var handler http.Handler = n
	options, enabled := r.Config(ctx).CORS("metrics")
	if enabled {
		handler = cors.New(options).Handler(handler)
	}
	return handler
}
