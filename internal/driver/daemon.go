// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"context"
	stderrors "errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"syscall"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"github.com/ory/analytics-go/v5"
	"github.com/ory/graceful"
	"github.com/ory/herodot"
	"github.com/ory/x/healthx"
	"github.com/ory/x/httprouterx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/metricsx"
	"github.com/ory/x/otelx"
	"github.com/ory/x/otelx/semconv"
	"github.com/ory/x/prometheusx"
	"github.com/ory/x/reqlog"
	"github.com/ory/x/urlx"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"github.com/urfave/negroni"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"golang.org/x/sync/errgroup"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/reflect/protoregistry"

	rts "github.com/ory/keto/gen/go/ory/keto/relation_tuples/v1alpha2"
	"github.com/ory/keto/gen/go/ory/keto/relation_tuples/v1alpha2/relationtuplesconnect"
	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/namespace/namespacehandler"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/schema"
)

func (r *RegistryDefault) enableSqa(cmd *cobra.Command) {
	ctx := cmd.Context()

	var urls []string
	addr, _ := r.Config(ctx).ReadAPIListenOn()
	urls = append(urls, addr)
	addr, _ = r.Config(ctx).WriteAPIListenOn()
	urls = append(urls, addr)
	addr, _ = r.Config(ctx).MetricsListenOn()
	urls = append(urls, addr)
	addr, _ = r.Config(ctx).OPLSyntaxAPIListenOn()
	urls = append(urls, addr)

	if c, y := r.Config(ctx).CORS("read"); y {
		urls = append(urls, c.AllowedOrigins...)
	}
	if c, y := r.Config(ctx).CORS("write"); y {
		urls = append(urls, c.AllowedOrigins...)
	}
	if c, y := r.Config(ctx).CORS("metrics"); y {
		urls = append(urls, c.AllowedOrigins...)
	}

	host := urlx.ExtractPublicAddress(urls...)

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
				relationtuplesconnect.VersionServiceGetVersionProcedure,
				"/" + grpchealth.HealthV1ServiceName + "/Check",
				"/" + grpchealth.HealthV1ServiceName + "/Watch",

				relationtuple.ReadRouteBase,
				relationtuplesconnect.ReadServiceListRelationTuplesProcedure,
				check.RouteBase,
				relationtuplesconnect.CheckServiceCheckProcedure,
				relationtuplesconnect.CheckServiceBatchCheckProcedure,
				expand.RouteBase,
				relationtuplesconnect.ExpandServiceExpandProcedure,
				relationtuplesconnect.WriteServiceTransactRelationTuplesProcedure,
				relationtuplesconnect.WriteServiceDeleteRelationTuplesProcedure,
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
			Hostname: host,
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
	rt := r.ReadRouter(ctx)

	if tracer := r.Tracer(ctx); tracer.IsLoaded() {
		rt = otelx.NewMiddleware(rt, "serveRead",
			otelhttp.WithTracerProvider(tracer.Provider()),
		)
	}

	return func() error {
		addr, listenFile := r.Config(ctx).ReadAPIListenOn()
		return servePort(ctx, r.Logger().WithField("endpoint", "read"), addr, listenFile, rt, done)
	}
}

func (r *RegistryDefault) serveWrite(ctx context.Context, done chan<- struct{}) func() error {
	rt := r.WriteRouter(ctx)

	if tracer := r.Tracer(ctx); tracer.IsLoaded() {
		rt = otelx.NewMiddleware(rt, "serveWrite",
			otelhttp.WithTracerProvider(tracer.Provider()),
		)
	}

	return func() error {
		addr, listenFile := r.Config(ctx).WriteAPIListenOn()
		return servePort(ctx, r.Logger().WithField("endpoint", "write"), addr, listenFile, rt, done)
	}
}

func (r *RegistryDefault) serveOPLSyntax(ctx context.Context, done chan<- struct{}) func() error {
	rt := r.OPLSyntaxRouter(ctx)

	if tracer := r.Tracer(ctx); tracer.IsLoaded() {
		rt = otelx.NewMiddleware(rt, "serveOPLSyntax",
			otelhttp.WithTracerProvider(tracer.Provider()),
		)
	}

	return func() error {
		addr, listenFile := r.Config(ctx).OPLSyntaxAPIListenOn()
		return servePort(ctx, r.Logger().WithField("endpoint", "opl"), addr, listenFile, rt, done)
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

func servePort(ctx context.Context, log *logrusx.Logger, addr, listenFile string, router http.Handler, done chan<- struct{}) error {
	l, err := listenAndWriteFile(ctx, addr, listenFile)
	if err != nil {
		return err
	}

	p := new(http.Protocols)
	p.SetHTTP1(true)
	// For gRPC clients, it's convenient to support HTTP/2 without TLS.
	p.SetUnencryptedHTTP2(true)
	p.SetHTTP2(true)
	//nolint:gosec // graceful.WithDefaults already sets a timeout
	restS := graceful.WithDefaults(&http.Server{
		Handler:   router,
		Protocols: p,
	})

	eg := &errgroup.Group{}

	eg.Go(func() error {
		if err := restS.Serve(l); !errors.Is(err, http.ErrServerClosed) {
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

		shutdownEg := errgroup.Group{}
		shutdownEg.Go(func() error {
			// we ignore net.ErrClosed, because a cmux listener's close func is actually the one of the root listener (which is closed in a racy fashion)
			if err := restS.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) && !errors.Is(err, net.ErrClosed) {
				// unexpected error
				return errors.WithStack(err)
			}
			return nil
		})
		shutdownEg.Go(func() error {
			gracefulDone := make(chan struct{})
			go func() {
				close(gracefulDone)
			}()
			select {
			case <-gracefulDone:
				return nil
			case <-ctx.Done():
				return errors.New("graceful stop of gRPC server timed out, had to force it")
			}
		})

		return shutdownEg.Wait()
	})

	return eg.Wait()
}

func (r *RegistryDefault) allHandlers() []Handler {
	return []Handler{
		relationtuple.NewReadHandler(r),
		relationtuple.NewWriteHandler(r),
		check.NewHandler(r),
		expand.NewHandler(r),
		namespacehandler.New(r),
		schema.NewHandler(r),
	}
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

var httpMetrics = prometheusx.NewHTTPMetrics("keto", prometheusx.HTTPPrefix, config.Version, config.Commit, config.Date)

func (r *RegistryDefault) ReadRouter(ctx context.Context) http.Handler {
	recovery := negroni.NewRecovery()
	recovery.Logger = r.Logger()

	router := httprouterx.NewRouterPublic()
	n := negroni.New(
		recovery,
		httpMetrics,
	)
	for _, f := range r.defaultHttpMiddlewares {
		n.UseFunc(f)
	}
	n.UseFunc(semconv.Middleware)
	n.Use(reqlog.NewMiddlewareFromLogger(r.l, "read#Ory Keto").ExcludePaths(healthx.AliveCheckPath, healthx.ReadyCheckPath))

	prometheusx.SetMuxRoutes(router)
	r.HealthHandler().SetHealthRoutes(router, false)
	r.HealthHandler().SetVersionRoutes(router)

	for _, h := range registerReflectionAndHealth[ReadHandler](r, router) {
		h.RegisterReadRoutes(router)
	}

	if r.sqaService != nil {
		n.Use(r.sqaService)
	}

	n.UseHandler(router)

	var handler http.Handler = n
	options, enabled := r.Config(ctx).CORS("read")
	if enabled {
		handler = cors.New(options).Handler(handler)
	}

	return handler
}

func (r *RegistryDefault) WriteRouter(ctx context.Context) http.Handler {
	recovery := negroni.NewRecovery()
	recovery.Logger = r.Logger()

	router := httprouterx.NewRouterAdmin()
	n := negroni.New(
		recovery,
		httpMetrics,
	)
	for _, f := range r.defaultHttpMiddlewares {
		n.UseFunc(f)
	}
	n.UseFunc(semconv.Middleware)
	n.Use(reqlog.NewMiddlewareFromLogger(r.l, "write#Ory Keto").ExcludePaths(healthx.AliveCheckPath, healthx.ReadyCheckPath))

	prometheusx.SetMuxRoutes(router)

	r.HealthHandler().SetHealthRoutes(router, false)
	r.HealthHandler().SetVersionRoutes(router)

	for _, h := range registerReflectionAndHealth[WriteHandler](r, router) {
		h.RegisterWriteRoutes(router)
	}

	if r.sqaService != nil {
		n.Use(r.sqaService)
	}

	n.UseHandler(router)

	var handler http.Handler = n
	options, enabled := r.Config(ctx).CORS("write")
	if enabled {
		handler = cors.New(options).Handler(handler)
	}

	return handler
}

func (r *RegistryDefault) OPLSyntaxRouter(ctx context.Context) http.Handler {
	recovery := negroni.NewRecovery()
	recovery.Logger = r.Logger()

	router := httprouterx.NewRouter()
	n := negroni.New(
		recovery,
		httpMetrics,
	)
	for _, f := range r.defaultHttpMiddlewares {
		n.UseFunc(f)
	}
	n.UseFunc(semconv.Middleware)
	n.Use(reqlog.NewMiddlewareFromLogger(r.l, "syntax#Ory Keto").ExcludePaths(healthx.AliveCheckPath, healthx.ReadyCheckPath))

	prometheusx.SetMuxRoutes(router)

	r.HealthHandler().SetHealthRoutes(router, false)
	r.HealthHandler().SetVersionRoutes(router)

	for _, h := range registerReflectionAndHealth[OPLSyntaxHandler](r, router) {
		h.RegisterSyntaxRoutes(router)
	}

	if r.sqaService != nil {
		n.Use(r.sqaService)
	}

	n.UseHandler(router)

	var handler http.Handler = n
	options, enabled := r.Config(ctx).CORS("write")
	if enabled {
		handler = cors.New(options).Handler(handler)
	}

	return handler
}

func (r *RegistryDefault) recoveryHandler(_ context.Context, spec connect.Spec, _ http.Header, val any) error {
	r.Logger().
		WithField("reason", val).
		WithField("procedure", spec.Procedure).
		WithField("stack_trace", string(debug.Stack())).
		Error("panic recovered")
	return connect.NewError(connect.CodeInternal, fmt.Errorf("%v", val))
}

func (r *RegistryDefault) HandlerOptions() []connect.HandlerOption {
	return append([]connect.HandlerOption{
		connect.WithRecover(r.recoveryHandler),
		connect.WithInterceptors(connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
			return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
				resp, err := next(ctx, req)
				if hErr, ok := stderrors.AsType[*herodot.DefaultError](err); ok {
					cErr := connect.NewError(connect.Code(hErr.GRPCCodeField), hErr)
					if reason := hErr.Reason(); reason != "" {
						d, err := connect.NewErrorDetail(&errdetails.ErrorInfo{Reason: reason})
						if err == nil {
							cErr.AddDetail(d)
						}
					}
					return resp, cErr
				}
				return resp, err
			}
		})),
		connect.WithInterceptors(r.defaultConnectInterceptors...),
	},
		r.defaultHandlerOptions...)
}

func (r *RegistryDefault) metricsRouter(ctx context.Context) http.Handler {
	recovery := negroni.NewRecovery()
	recovery.Logger = r.Logger()

	n := negroni.New(
		recovery,
		reqlog.NewMiddlewareFromLogger(r.Logger(), "keto").ExcludePaths(prometheusx.MetricsPrometheusPath),
	)

	router := http.NewServeMux()
	prometheusx.SetMuxRoutes(router)
	n.UseHandler(router)

	var handler http.Handler = n
	options, enabled := r.Config(ctx).CORS("metrics")
	if enabled {
		handler = cors.New(options).Handler(handler)
	}
	return handler
}

func registerReflectionAndHealth[H Handler](r *RegistryDefault, router httprouterx.Router) (handlers []H) {
	protoFiles := &protoregistry.Files{}
	var protoServices []string
	for _, h := range r.allHandlers() {
		if h, ok := h.(H); ok {
			handlers = append(handlers, h)
			for _, f := range h.ProtoFiles() {
				if err := protoFiles.RegisterFile(f); err != nil && !strings.Contains(err.Error(), "is already registered") {
					r.Logger().WithError(err).
						WithField("file_path", f.Path()).
						Error("unable to register proto file for reflection")
				}
				s := f.Services()
				for i := range s.Len() {
					protoServices = append(protoServices, string(s.Get(i).FullName()))
				}
			}
		}
	}

	router.Handle(relationtuplesconnect.NewVersionServiceHandler(r, r.HandlerOptions()...))
	protoServices = append(protoServices, relationtuplesconnect.VersionServiceName)
	if err := protoFiles.RegisterFile(rts.File_ory_keto_relation_tuples_v1alpha2_version_proto); err != nil {
		r.Logger().WithError(err).Error("unable to register version proto file for reflection")
	}

	reflector := grpcreflect.NewReflector(
		grpcreflect.NamerFunc(func() []string { return protoServices }),
		grpcreflect.WithDescriptorResolver(protoFiles),
	)
	router.Handle(grpcreflect.NewHandlerV1(reflector, r.HandlerOptions()...))
	router.Handle(grpcreflect.NewHandlerV1Alpha(reflector, r.HandlerOptions()...))

	router.Handle(grpchealth.NewHandler(grpchealth.NewStaticChecker(protoServices...), r.HandlerOptions()...))

	return
}
