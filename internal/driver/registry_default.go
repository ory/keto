package driver

import (
	"context"
	"net/http"
	"sync"

	"github.com/ory/x/configx"

	"github.com/ory/keto/ketoctx"

	prometheus "github.com/ory/x/prometheusx"

	"github.com/ory/x/networkx"
	"github.com/rs/cors"

	"github.com/gobuffalo/pop/v6"
	"github.com/ory/x/popx"
	"github.com/pkg/errors"

	"github.com/ory/x/dbal"

	"github.com/ory/x/metricsx"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/ory/x/reqlog"
	"github.com/urfave/negroni"
	"google.golang.org/grpc/reflection"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/ory/keto/internal/driver/config"

	"github.com/ory/herodot"
	"github.com/ory/x/healthx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/tracing"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/internal/persistence/sql"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

var (
	_ relationtuple.ManagerProvider = (*RegistryDefault)(nil)
	_ x.WriterProvider              = (*RegistryDefault)(nil)
	_ x.LoggerProvider              = (*RegistryDefault)(nil)
	_ Registry                      = (*RegistryDefault)(nil)
	_ acl.VersionServiceServer      = (*RegistryDefault)(nil)
)

type (
	RegistryDefault struct {
		p    persistence.Persister
		mb   *popx.MigrationBox
		l    *logrusx.Logger
		w    herodot.Writer
		ce   *check.Engine
		ee   *expand.Engine
		c    *config.Config
		conn *pop.Connection

		initialized    sync.Once
		healthH        *healthx.Handler
		healthServer   *health.Server
		handlers       []Handler
		sqaService     *metricsx.Service
		tracer         *tracing.Tracer
		pmm            *prometheus.MetricsManager
		metricsHandler *prometheus.Handler
	}
	Handler interface {
		RegisterReadRoutes(r *x.ReadRouter)
		RegisterWriteRoutes(r *x.WriteRouter)
		RegisterReadGRPC(s *grpc.Server)
		RegisterWriteGRPC(s *grpc.Server)
	}
)

func (r *RegistryDefault) ContextualizeConfig(_ context.Context) *configx.Provider {
	return nil
}

func (r *RegistryDefault) Config(ctx context.Context) *config.Config {
	provider := ketoctx.ContextualizeConfig(ketoctx.WithConfigContextualizer(ctx, r))
	if provider == nil {
		return r.c
	}

	return config.New(ctx, r.Logger(), provider)
}

func (r *RegistryDefault) HealthHandler() *healthx.Handler {
	if r.healthH == nil {
		r.healthH = healthx.NewHandler(r.Writer(), config.Version, healthx.ReadyCheckers{})
	}

	return r.healthH
}

func (r *RegistryDefault) HealthServer() *health.Server {
	if r.healthServer == nil {
		r.healthServer = health.NewServer()
	}

	return r.healthServer
}

func (r *RegistryDefault) GetVersion(_ context.Context, _ *acl.GetVersionRequest) (*acl.GetVersionResponse, error) {
	return &acl.GetVersionResponse{Version: config.Version}, nil
}

func (r *RegistryDefault) Tracer(ctx context.Context) *tracing.Tracer {
	if r.tracer == nil {
		// Tracing is initialized only once so it can not be hot reloaded or context-aware.
		t, err := tracing.New(r.Logger(), r.Config(ctx).TracingConfig())
		if err != nil {
			r.Logger().WithError(err).Fatalf("Unable to initialize Tracer.")
		}
		r.tracer = t
	}

	return r.tracer
}

func (r *RegistryDefault) MetricsHandler() *prometheus.Handler {
	if r.metricsHandler == nil {
		r.metricsHandler = prometheus.NewHandler(r.Writer(), config.Version)
	}
	return r.metricsHandler
}

func (r *RegistryDefault) PrometheusManager() *prometheus.MetricsManager {
	if r.pmm == nil {
		r.pmm = prometheus.NewMetricsManager("keto", config.Version, config.Commit, config.Date)
	}
	return r.pmm
}

func (r *RegistryDefault) Logger() *logrusx.Logger {
	if r.l == nil {
		r.l = logrusx.New("ORY Keto", config.Version)
	}
	return r.l
}

func (r *RegistryDefault) Writer() herodot.Writer {
	if r.w == nil {
		r.w = herodot.NewJSONWriter(r.Logger())
	}
	return r.w
}

func (r *RegistryDefault) RelationTupleManager() relationtuple.Manager {
	if r.p == nil {
		panic("no relation tuple manager, but expected to have one")
	}
	return r.p
}

func (r *RegistryDefault) Persister() persistence.Persister {
	if r.p == nil {
		panic("no persister, but expected to have one")
	}
	return r.p
}

func (r *RegistryDefault) PermissionEngine() *check.Engine {
	if r.ce == nil {
		r.ce = check.NewEngine(r)
	}
	return r.ce
}

func (r *RegistryDefault) ExpandEngine() *expand.Engine {
	if r.ee == nil {
		r.ee = expand.NewEngine(r)
	}
	return r.ee
}

func (r *RegistryDefault) MigrationBox(ctx context.Context) (*popx.MigrationBox, error) {
	if r.mb == nil {
		c, err := r.PopConnection(ctx)
		if err != nil {
			return nil, err
		}
		mb, err := sql.NewMigrationBox(c, r.Logger(), r.Tracer(ctx))
		if err != nil {
			return nil, err
		}
		r.mb = mb
	}
	return r.mb, nil
}

func (r *RegistryDefault) MigrateUp(ctx context.Context) error {
	mb, err := r.MigrationBox(ctx)
	if err != nil {
		return err
	}
	if err := mb.Up(ctx); err != nil {
		return err
	}
	return r.Init(ctx)
}

func (r *RegistryDefault) MigrateDown(ctx context.Context) error {
	mb, err := r.MigrationBox(ctx)
	if err != nil {
		return err
	}
	return mb.Up(ctx)
}

func (r *RegistryDefault) determineNetwork(ctx context.Context) (*networkx.Network, error) {
	c, err := r.PopConnection(ctx)
	if err != nil {
		return nil, err
	}
	mb, err := popx.NewMigrationBox(networkx.Migrations, popx.NewMigrator(c, r.Logger(), r.Tracer(ctx), 0))
	if err != nil {
		return nil, err
	}
	s, err := mb.Status(ctx)
	if err != nil {
		return nil, err
	}
	if s.HasPending() {
		return nil, errors.WithStack(persistence.ErrNetworkMigrationsMissing)
	}

	return networkx.NewManager(c, r.Logger(), r.Tracer(ctx)).Determine(ctx)
}

func (r *RegistryDefault) InitWithoutNetworkID(ctx context.Context) error {
	if dbal.IsMemorySQLite(r.Config(ctx).DSN()) {
		mb, err := r.MigrationBox(ctx)
		if err != nil {
			return err
		}

		if err := mb.Up(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (r *RegistryDefault) Init(ctx context.Context) (err error) {
	r.initialized.Do(func() {
		err = func() error {
			if err := r.InitWithoutNetworkID(ctx); err != nil {
				return err
			}

			network, err := r.determineNetwork(ctx)
			if err != nil {
				return err
			}

			r.p, err = sql.NewPersister(ctx, r, network.ID)
			if err != nil {
				return err
			}

			return nil
		}()
	})
	return
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
	n := negroni.New(reqlog.NewMiddlewareFromLogger(r.l, "read#Ory Keto").ExcludePaths(healthx.AliveCheckPath, healthx.ReadyCheckPath))

	br := &x.ReadRouter{Router: httprouter.New()}

	r.HealthHandler().SetHealthRoutes(br.Router, false)
	r.HealthHandler().SetVersionRoutes(br.Router)

	for _, h := range r.allHandlers() {
		h.RegisterReadRoutes(br)
	}

	n.UseHandler(br)

	if t := r.Tracer(ctx); t.IsLoaded() {
		n.Use(t)
	}

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
	n := negroni.New(reqlog.NewMiddlewareFromLogger(r.l, "write#Ory Keto").ExcludePaths(healthx.AliveCheckPath, healthx.ReadyCheckPath))

	pr := &x.WriteRouter{Router: httprouter.New()}

	r.HealthHandler().SetHealthRoutes(pr.Router, false)
	r.HealthHandler().SetVersionRoutes(pr.Router)

	for _, h := range r.allHandlers() {
		h.RegisterWriteRoutes(pr)
	}

	n.UseHandler(pr)

	if t := r.Tracer(ctx); t.IsLoaded() {
		n.Use(t)
	}

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

func (r *RegistryDefault) MetricsRouter() http.Handler {
	n := negroni.New(reqlog.NewMiddlewareFromLogger(r.Logger(), "keto").ExcludePaths(prometheus.MetricsPrometheusPath))
	router := httprouter.New()

	r.PrometheusManager().RegisterRouter(router)
	r.MetricsHandler().SetRoutes(router)
	n.UseHandler(router)
	n.Use(r.PrometheusManager())

	var handler http.Handler = n
	options, enabled := r.Config().CORS("metrics")
	if enabled {
		handler = cors.New(options).Handler(handler)
	}
	return handler
}

func (r *RegistryDefault) unaryInterceptors(ctx context.Context) []grpc.UnaryServerInterceptor {
	is := []grpc.UnaryServerInterceptor{
		herodot.UnaryErrorUnwrapInterceptor,
		grpcMiddleware.ChainUnaryServer(
			grpc_logrus.UnaryServerInterceptor(r.l.Entry),
		),
	}
	if r.Tracer(ctx).IsLoaded() {
		is = append(is, otgrpc.OpenTracingServerInterceptor(r.Tracer(ctx).Tracer()))
	}
	if r.sqaService != nil {
		is = append(is, r.sqaService.UnaryInterceptor)
	}
	return is
}

func (r *RegistryDefault) streamInterceptors(ctx context.Context) []grpc.StreamServerInterceptor {
	is := []grpc.StreamServerInterceptor{
		herodot.StreamErrorUnwrapInterceptor,
		grpcMiddleware.ChainStreamServer(
			grpc_logrus.StreamServerInterceptor(r.l.Entry),
		),
	}
	if r.Tracer(ctx).IsLoaded() {
		is = append(is, otgrpc.OpenTracingStreamServerInterceptor(r.Tracer(ctx).Tracer()))
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
	acl.RegisterVersionServiceServer(s, r)
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
	acl.RegisterVersionServiceServer(s, r)
	reflection.Register(s)

	for _, h := range r.allHandlers() {
		h.RegisterWriteGRPC(s)
	}

	return s
}
