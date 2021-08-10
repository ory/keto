package driver

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/ory/x/networkx"

	"github.com/cenkalti/backoff/v3"
	"github.com/gobuffalo/pop/v5"
	"github.com/luna-duclos/instrumentedsql"
	"github.com/luna-duclos/instrumentedsql/opentracing"
	"github.com/ory/x/popx"
	"github.com/ory/x/sqlcon"
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

		initialized  sync.Once
		healthH      *healthx.Handler
		healthServer *health.Server
		handlers     []Handler
		sqaService   *metricsx.Service
		tracer       *tracing.Tracer
	}
	Handler interface {
		RegisterReadRoutes(r *x.ReadRouter)
		RegisterWriteRoutes(r *x.WriteRouter)
		RegisterReadGRPC(s *grpc.Server)
		RegisterWriteGRPC(s *grpc.Server)
	}
)

func (r *RegistryDefault) BuildVersion() string {
	return config.Version
}

func (r *RegistryDefault) BuildDate() string {
	return config.Date
}

func (r *RegistryDefault) BuildHash() string {
	return config.Commit
}

func (r *RegistryDefault) Config() *config.Config {
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

func (r *RegistryDefault) Tracer() *tracing.Tracer {
	if r.tracer == nil {
		// Tracing is initialized only once so it can not be hot reloaded or context-aware.
		t, err := tracing.New(r.Logger(), r.Config().TracingConfig())
		if err != nil {
			r.Logger().WithError(err).Fatalf("Unable to initialize Tracer.")
		}
		r.tracer = t
	}

	return r.tracer
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

func (r *RegistryDefault) MigrationBox() (*popx.MigrationBox, error) {
	if r.mb == nil {
		c, err := r.PopConnection()
		if err != nil {
			return nil, err
		}
		mb, err := sql.NewMigrationBox(c, r.Logger(), r.Tracer())
		if err != nil {
			return nil, err
		}
		r.mb = mb
	}
	return r.mb, nil
}

func (r *RegistryDefault) MigrateUp(ctx context.Context) error {
	mb, err := r.MigrationBox()
	if err != nil {
		return err
	}
	if err := mb.Up(ctx); err != nil {
		return err
	}
	return r.Init(ctx)
}

func (r *RegistryDefault) MigrateDown(ctx context.Context) error {
	mb, err := r.MigrationBox()
	if err != nil {
		return err
	}
	return mb.Up(ctx)
}

func (r *RegistryDefault) PopConnection() (*pop.Connection, error) {
	if r.conn == nil {
		tracer := r.Tracer()

		var opts []instrumentedsql.Opt
		if tracer.IsLoaded() {
			opts = []instrumentedsql.Opt{
				instrumentedsql.WithTracer(opentracing.NewTracer(true)),
				instrumentedsql.WithOmitArgs(),
			}
		}
		pool, idlePool, connMaxLifetime, connMaxIdleTime, cleanedDSN := sqlcon.ParseConnectionOptions(r.Logger(), r.Config().DSN())
		connDetails := &pop.ConnectionDetails{
			URL:                       sqlcon.FinalizeDSN(r.Logger(), cleanedDSN),
			IdlePool:                  idlePool,
			ConnMaxLifetime:           connMaxLifetime,
			ConnMaxIdleTime:           connMaxIdleTime,
			Pool:                      pool,
			UseInstrumentedDriver:     tracer != nil && tracer.IsLoaded(),
			InstrumentedDriverOptions: opts,
		}

		bc := backoff.NewExponentialBackOff()
		bc.MaxElapsedTime = time.Minute * 5
		bc.Reset()

		if err := backoff.Retry(func() (err error) {
			conn, err := pop.NewConnection(connDetails)
			if err != nil {
				r.Logger().WithError(err).Error("Unable to connect to database, retrying.")
				return errors.WithStack(err)
			}

			if err := conn.Open(); err != nil {
				r.Logger().WithError(err).Error("Unable to open the database connection, retrying.")
				return errors.WithStack(err)
			}

			if err := conn.Store.(interface{ Ping() error }).Ping(); err != nil {
				r.Logger().WithError(err).Error("Unable to ping the database connection, retrying.")
				return errors.WithStack(err)
			}

			r.conn = conn
			return nil
		}, bc); err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return r.conn, nil
}

func (r *RegistryDefault) determineNetwork(ctx context.Context) (*networkx.Network, error) {
	c, err := r.PopConnection()
	if err != nil {
		return nil, err
	}
	mb, err := popx.NewMigrationBox(networkx.Migrations, popx.NewMigrator(c, r.Logger(), r.Tracer(), 0))
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

	return networkx.NewManager(c, r.Logger(), r.Tracer()).Determine(ctx)
}

func (r *RegistryDefault) InitWithoutNetworkID(ctx context.Context) error {
	if dbal.IsMemorySQLite(r.c.DSN()) {
		mb, err := r.MigrationBox()
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

			r.p, err = sql.NewPersister(r, network.ID)
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

func (r *RegistryDefault) ReadRouter() http.Handler {
	n := negroni.New(reqlog.NewMiddlewareFromLogger(r.l, "write#Ory Keto"))

	br := &x.ReadRouter{Router: httprouter.New()}

	r.HealthHandler().SetHealthRoutes(br.Router, false)

	for _, h := range r.allHandlers() {
		h.RegisterReadRoutes(br)
	}

	n.UseHandler(br)

	if t := r.Tracer(); t.IsLoaded() {
		n.Use(t)
	}

	if r.sqaService != nil {
		n.Use(r.sqaService)
	}

	return n
}

func (r *RegistryDefault) WriteRouter() http.Handler {
	n := negroni.New(reqlog.NewMiddlewareFromLogger(r.l, "write#Ory Keto"))

	pr := &x.WriteRouter{Router: httprouter.New()}

	r.HealthHandler().SetHealthRoutes(pr.Router, false)

	for _, h := range r.allHandlers() {
		h.RegisterWriteRoutes(pr)
	}

	n.UseHandler(pr)

	if t := r.Tracer(); t.IsLoaded() {
		n.Use(t)
	}

	if r.sqaService != nil {
		n.Use(r.sqaService)
	}

	return n
}

func (r *RegistryDefault) unaryInterceptors() []grpc.UnaryServerInterceptor {
	is := []grpc.UnaryServerInterceptor{
		herodot.UnaryErrorUnwrapInterceptor,
		grpcMiddleware.ChainUnaryServer(
			grpc_logrus.UnaryServerInterceptor(r.l.Entry),
		),
	}
	if r.Tracer().IsLoaded() {
		is = append(is, otgrpc.OpenTracingServerInterceptor(r.Tracer().Tracer()))
	}
	return is
}

func (r *RegistryDefault) streamInterceptors() []grpc.StreamServerInterceptor {
	is := []grpc.StreamServerInterceptor{
		herodot.StreamErrorUnwrapInterceptor,
		grpcMiddleware.ChainStreamServer(
			grpc_logrus.StreamServerInterceptor(r.l.Entry),
		),
	}
	if r.Tracer().IsLoaded() {
		is = append(is, otgrpc.OpenTracingStreamServerInterceptor(r.Tracer().Tracer()))
	}
	return is
}

func (r *RegistryDefault) ReadGRPCServer() *grpc.Server {
	s := grpc.NewServer(
		grpc.ChainStreamInterceptor(r.streamInterceptors()...),
		grpc.ChainUnaryInterceptor(r.unaryInterceptors()...),
	)

	grpcHealthV1.RegisterHealthServer(s, r.HealthServer())
	acl.RegisterVersionServiceServer(s, r)
	reflection.Register(s)

	for _, h := range r.allHandlers() {
		h.RegisterReadGRPC(s)
	}

	return s
}

func (r *RegistryDefault) WriteGRPCServer() *grpc.Server {
	s := grpc.NewServer(
		grpc.ChainStreamInterceptor(r.streamInterceptors()...),
		grpc.ChainUnaryInterceptor(r.unaryInterceptors()...),
	)

	grpcHealthV1.RegisterHealthServer(s, r.HealthServer())
	acl.RegisterVersionServiceServer(s, r)
	reflection.Register(s)

	for _, h := range r.allHandlers() {
		h.RegisterWriteGRPC(s)
	}

	return s
}
