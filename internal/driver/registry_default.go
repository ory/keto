package driver

import (
	"context"
	"net/http"
	"sync"

	"github.com/ory/keto/ketoctx"

	prometheus "github.com/ory/x/prometheusx"

	"github.com/ory/x/networkx"
	"github.com/rs/cors"

	"github.com/gobuffalo/pop/v6"
	"github.com/ory/x/networkx"
	"github.com/ory/x/popx"
	"github.com/pkg/errors"

	"github.com/ory/x/dbal"

	"github.com/ory/x/metricsx"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"

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
	_ relationtuple.ManagerProvider  = (*RegistryDefault)(nil)
	_ x.WriterProvider               = (*RegistryDefault)(nil)
	_ x.LoggerProvider               = (*RegistryDefault)(nil)
	_ Registry                       = (*RegistryDefault)(nil)
	_ acl.VersionServiceServer       = (*RegistryDefault)(nil)
	_ ketoctx.ContextualizerProvider = (*RegistryDefault)(nil)
)

type (
	RegistryDefault struct {
		p     persistence.Persister
		mb    *popx.MigrationBox
		l     *logrusx.Logger
		w     herodot.Writer
		ce    *check.Engine
		ee    *expand.Engine
		c     *config.Config
		conn  *pop.Connection
		ctxer ketoctx.Contextualizer

		initialized    sync.Once
		healthH        *healthx.Handler
		healthServer   *health.Server
		handlers       []Handler
		sqaService     *metricsx.Service
		tracer         *tracing.Tracer
		pmm            *prometheus.MetricsManager
		metricsHandler *prometheus.Handler

		defaultUnaryInterceptors  []grpc.UnaryServerInterceptor
		defaultStreamInterceptors []grpc.StreamServerInterceptor
		defaultHttpMiddlewares    []func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
	}
	Handler interface {
		RegisterReadRoutes(r *x.ReadRouter)
		RegisterWriteRoutes(r *x.WriteRouter)
		RegisterReadGRPC(s *grpc.Server)
		RegisterWriteGRPC(s *grpc.Server)
	}
)

func (r *RegistryDefault) Contextualizer() ketoctx.Contextualizer {
	return r.ctxer
}

func (r *RegistryDefault) Config(ctx context.Context) *config.Config {
	if provider := r.ctxer.Config(ctx, nil); provider != nil {
		return config.New(ctx, r.Logger(), provider)
	}
	return r.c
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
