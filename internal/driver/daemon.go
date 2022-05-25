package driver

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	prometheus "github.com/ory/x/prometheusx"
	grpcOtel "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"

	"github.com/ory/x/logrusx"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"github.com/ory/x/reqlog"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"

	"github.com/ory/analytics-go/v4"
	"github.com/ory/x/healthx"
	"github.com/ory/x/metricsx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver/config"

	"github.com/ory/graceful"
	"github.com/pkg/errors"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func (r *RegistryDefault) enableSqa(cmd *cobra.Command) {
	ctx := cmd.Context()

	r.sqaService = metricsx.New(
		cmd,
		r.Logger(),
		r.Config(ctx).Source(),
		&metricsx.Options{
			Service:       "ory-keto",
			ClusterID:     metricsx.Hash(r.Config(ctx).DSN()),
			IsDevelopment: strings.HasPrefix(r.Config(ctx).DSN(), "sqlite"),
			WriteKey:      "qQlI6q8Q4WvkzTjKQSor4sHYOikHIvvi",
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
				Endpoint: "https://sqa.ory.sh",
			},
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

	doneShutdown := make(chan struct{}, 3)

	go func() {
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

		select {
		case <-osSignals:
			cancel()
		case <-innerCtx.Done():
		}

		ctx, cancel := context.WithTimeout(context.Background(), graceful.DefaultShutdownTimeout)
		defer cancel()

		nWaitingForShutdown := cap(doneShutdown)
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
	}()

	eg := &errgroup.Group{}

	eg.Go(r.serveRead(innerCtx, doneShutdown))
	eg.Go(r.serveWrite(innerCtx, doneShutdown))
	eg.Go(r.serveMetrics(innerCtx, doneShutdown))

	return eg.Wait()
}

func (r *RegistryDefault) serveRead(ctx context.Context, done chan<- struct{}) func() error {
	rt, s := r.ReadRouter(ctx), r.ReadGRPCServer(ctx)

	if tracer := r.Tracer(ctx); tracer.IsLoaded() {
		rt = x.TraceHandler(rt)
	}

	return func() error {
		return multiplexPort(ctx, r.Logger().WithField("endpoint", "read"), r.Config(ctx).ReadAPIListenOn(), rt, s, done)
	}
}

func (r *RegistryDefault) serveWrite(ctx context.Context, done chan<- struct{}) func() error {
	rt, s := r.WriteRouter(ctx), r.WriteGRPCServer(ctx)

	if tracer := r.Tracer(ctx); tracer.IsLoaded() {
		rt = x.TraceHandler(rt)
	}

	return func() error {
		return multiplexPort(ctx, r.Logger().WithField("endpoint", "write"), r.Config(ctx).WriteAPIListenOn(), rt, s, done)
	}
}

func (r *RegistryDefault) serveMetrics(ctx context.Context, done chan<- struct{}) func() error {
	return func() error {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		eg := &errgroup.Group{}
		s := graceful.WithDefaults(&http.Server{
			Handler: r.metricsRouter(ctx),
			Addr:    r.Config(ctx).MetricsListenOn(),
		})

		eg.Go(func() error {
			if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
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

func multiplexPort(ctx context.Context, log *logrusx.Logger, addr string, router http.Handler, grpcS *grpc.Server, done chan<- struct{}) error {
	l, err := (&net.ListenConfig{}).Listen(ctx, "tcp", addr)
	if err != nil {
		return err
	}

	m := cmux.New(l)
	m.SetReadTimeout(graceful.DefaultReadTimeout)

	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1())

	restS := graceful.WithDefaults(&http.Server{
		Handler: router,
	})

	eg := &errgroup.Group{}

	eg.Go(func() error {
		if err := grpcS.Serve(grpcL); !errors.Is(err, cmux.ErrServerClosed) {
			return errors.WithStack(err)
		}
		return nil
	})

	eg.Go(func() error {
		if err := restS.Serve(httpL); !errors.Is(err, http.ErrServerClosed) && !errors.Is(err, cmux.ErrServerClosed) {
			return errors.WithStack(err)
		}
		return nil
	})

	eg.Go(func() error {
		err := m.Serve()
		if err != nil && !errors.Is(err, net.ErrClosed) {
			// unexpected error
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

		ctx, cancel := context.WithTimeout(context.Background(), graceful.DefaultShutdownTimeout)
		defer cancel()

		shutdownEg := errgroup.Group{}
		shutdownEg.Go(func() error {
			// we ignore net.ErrClosed, because a cmux listener's close func is actually the one of the root listener (which is closed in a racy fashion)
			if err := restS.Shutdown(ctx); !(err == nil || errors.Is(err, http.ErrServerClosed) || errors.Is(err, net.ErrClosed)) {
				// unexpected error
				return errors.WithStack(err)
			}
			return nil
		})
		shutdownEg.Go(func() error {
			gracefulDone := make(chan struct{})
			go func() {
				grpcS.GracefulStop()
				close(gracefulDone)
			}()
			select {
			case <-gracefulDone:
				return nil
			case <-ctx.Done():
				grpcS.Stop()
				return errors.New("graceful stop of gRPC server canceled, had to force it")
			}
		})

		return shutdownEg.Wait()
	})

	return eg.Wait()
}

func (r *RegistryDefault) allHandlers() []Handler {
	if len(r.handlers) == 0 {
		r.handlers = []Handler{
			relationtuple.NewHandler(r),
			check.NewHandler(r),
			expand.NewHandler(r),
		}
	}
	return r.handlers
}

func (r *RegistryDefault) ReadRouter(ctx context.Context) http.Handler {
	n := negroni.New()
	for _, f := range r.defaultHttpMiddlewares {
		n.UseFunc(f)
	}
	n.Use(reqlog.NewMiddlewareFromLogger(r.l, "read#Ory Keto").ExcludePaths(healthx.AliveCheckPath, healthx.ReadyCheckPath))

	br := &x.ReadRouter{Router: httprouter.New()}

	r.HealthHandler().SetHealthRoutes(br.Router, false)
	r.HealthHandler().SetVersionRoutes(br.Router)

	for _, h := range r.allHandlers() {
		h.RegisterReadRoutes(br)
	}

	n.UseHandler(br)

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

func (r *RegistryDefault) WriteRouter(ctx context.Context) http.Handler {
	n := negroni.New()
	for _, f := range r.defaultHttpMiddlewares {
		n.UseFunc(f)
	}
	n.Use(reqlog.NewMiddlewareFromLogger(r.l, "write#Ory Keto").ExcludePaths(healthx.AliveCheckPath, healthx.ReadyCheckPath))

	pr := &x.WriteRouter{Router: httprouter.New()}

	r.HealthHandler().SetHealthRoutes(pr.Router, false)
	r.HealthHandler().SetVersionRoutes(pr.Router)

	for _, h := range r.allHandlers() {
		h.RegisterWriteRoutes(pr)
	}

	n.UseHandler(pr)

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

func (r *RegistryDefault) unaryInterceptors(ctx context.Context) []grpc.UnaryServerInterceptor {
	is := make([]grpc.UnaryServerInterceptor, len(r.defaultUnaryInterceptors), len(r.defaultUnaryInterceptors)+2)
	copy(is, r.defaultUnaryInterceptors)
	is = append(is,
		herodot.UnaryErrorUnwrapInterceptor,
		grpcMiddleware.ChainUnaryServer(
			grpc_logrus.UnaryServerInterceptor(r.l.Entry),
		),
	)
	if r.Tracer(ctx).IsLoaded() {
		is = append(is, grpcOtel.UnaryServerInterceptor(grpcOtel.WithTracerProvider(otel.GetTracerProvider())))
	}
	if r.sqaService != nil {
		is = append(is, r.sqaService.UnaryInterceptor)
	}
	return is
}

func (r *RegistryDefault) streamInterceptors(ctx context.Context) []grpc.StreamServerInterceptor {
	is := make([]grpc.StreamServerInterceptor, len(r.defaultStreamInterceptors), len(r.defaultStreamInterceptors)+2)
	copy(is, r.defaultStreamInterceptors)
	is = append(is,
		herodot.StreamErrorUnwrapInterceptor,
		grpcMiddleware.ChainStreamServer(
			grpc_logrus.StreamServerInterceptor(r.l.Entry),
		),
	)
	if r.Tracer(ctx).IsLoaded() {
		is = append(is, grpcOtel.StreamServerInterceptor(grpcOtel.WithTracerProvider(otel.GetTracerProvider())))
	}
	if r.sqaService != nil {
		is = append(is, r.sqaService.StreamInterceptor)
	}
	return is
}

func (r *RegistryDefault) ReadGRPCServer(ctx context.Context) *grpc.Server {
	s := grpc.NewServer(
		grpc.ChainStreamInterceptor(r.streamInterceptors(ctx)...),
		grpc.ChainUnaryInterceptor(r.unaryInterceptors(ctx)...),
	)

	grpcHealthV1.RegisterHealthServer(s, r.HealthServer())
	rts.RegisterVersionServiceServer(s, r)
	reflection.Register(s)

	for _, h := range r.allHandlers() {
		h.RegisterReadGRPC(s)
	}

	return s
}

func (r *RegistryDefault) WriteGRPCServer(ctx context.Context) *grpc.Server {
	s := grpc.NewServer(
		grpc.ChainStreamInterceptor(r.streamInterceptors(ctx)...),
		grpc.ChainUnaryInterceptor(r.unaryInterceptors(ctx)...),
	)

	grpcHealthV1.RegisterHealthServer(s, r.HealthServer())
	rts.RegisterVersionServiceServer(s, r)
	reflection.Register(s)

	for _, h := range r.allHandlers() {
		h.RegisterWriteGRPC(s)
	}

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
