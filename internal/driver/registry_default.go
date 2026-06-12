// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"context"
	"io/fs"
	"maps"
	"net/http"
	"sync"

	"connectrpc.com/connect"
	"github.com/gofrs/uuid"
	"github.com/ory/herodot"
	"github.com/pkg/errors"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/ory/pop/v6"
	"github.com/ory/x/contextx"
	"github.com/ory/x/dbal"
	"github.com/ory/x/fsx"
	"github.com/ory/x/healthx"
	"github.com/ory/x/httprouterx"
	"github.com/ory/x/httpx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/metricsx"
	"github.com/ory/x/networkx"
	"github.com/ory/x/otelx"
	"github.com/ory/x/popx"

	rts "github.com/ory/keto/gen/go/ory/keto/relation_tuples/v1alpha2"
	"github.com/ory/keto/gen/go/ory/keto/relation_tuples/v1alpha2/relationtuplesconnect"
	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/check/step"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/internal/persistence/sql"
	"github.com/ory/keto/internal/persistence/sql/migrations/uuidmapping"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/ketoctx"
)

var (
	_ relationtuple.ManagerProvider               = (*RegistryDefault)(nil)
	_ relationtuple.MapperProvider                = (*RegistryDefault)(nil)
	_ relationtuple.MappingManagerProvider        = (*RegistryDefault)(nil)
	_ httpx.WriterProvider                        = (*RegistryDefault)(nil)
	_ logrusx.Provider                            = (*RegistryDefault)(nil)
	_ Registry                                    = (*RegistryDefault)(nil)
	_ relationtuplesconnect.VersionServiceHandler = (*RegistryDefault)(nil)
)

type (
	RegistryDefault struct {
		relationtuplesconnect.UnimplementedVersionServiceHandler
		p               persistence.Persister
		mb              *popx.MigrationBox
		extraMigrations []fs.FS
		l               *logrusx.Logger
		w               herodot.Writer
		ce              *check.Engine
		ck              *step.Executor
		ee              *expand.Engine
		c               *config.Config
		conn            *pop.Connection
		ctxer           contextx.Contextualizer
		mapper          *relationtuple.Mapper
		readOnlyMapper  *relationtuple.Mapper
		replaceMapperNs func(ctx context.Context) uuid.UUID
		replaceShardID  func() uuid.UUID

		init1, init2       sync.Once
		init1err, init2err error

		healthH       *healthx.Handler
		healthServer  *health.Server
		sqaService    *metricsx.Service
		tracer        *otelx.Tracer
		tracerWrapper ketoctx.TracerWrapper

		defaultHandlerOptions      []connect.HandlerOption
		defaultConnectInterceptors []connect.Interceptor
		defaultHttpMiddlewares     []func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
		grpcTransportCredentials   credentials.TransportCredentials
		defaultMigrationOptions    []popx.MigrationBoxOption
		healthReadyCheckers        healthx.ReadyCheckers
		dbOpts                     []func(details *pop.ConnectionDetails)
	}
	ReadHandler interface {
		Handler
		RegisterReadRoutes(r *httprouterx.RouterPublic)
	}
	WriteHandler interface {
		Handler
		RegisterWriteRoutes(r *httprouterx.RouterAdmin)
	}
	OPLSyntaxHandler interface {
		Handler
		RegisterSyntaxRoutes(r httprouterx.Router)
	}
	Handler interface {
		ProtoFiles() []protoreflect.FileDescriptor
	}
)

func (r *RegistryDefault) Mapper() *relationtuple.Mapper {
	if r.mapper == nil {
		r.mapper = &relationtuple.Mapper{D: r}
	}
	return r.mapper
}

func (r *RegistryDefault) ReadOnlyMapper() *relationtuple.Mapper {
	if r.readOnlyMapper == nil {
		r.readOnlyMapper = &relationtuple.Mapper{D: r, ReadOnly: true}
	}
	return r.readOnlyMapper
}

func (r *RegistryDefault) Contextualizer() contextx.Contextualizer {
	return r.ctxer
}

func (r *RegistryDefault) Config(ctx context.Context) *config.Config {
	if provider := r.ctxer.Config(ctx, r.c.Source()); provider != r.c.Source() {
		return config.New(ctx, r.Logger(), provider)
	}
	return r.c
}

func (r *RegistryDefault) HealthHandler() *healthx.Handler {
	if r.healthH != nil {
		return r.healthH
	}

	h := healthx.ReadyCheckers{
		"database": func(req *http.Request) error {
			return r.Persister().Connection(req.Context()).Store.PingContext(req.Context())
		},
		"migrations": func(req *http.Request) error {
			mb, err := r.MigrationBox(req.Context())
			if err != nil {
				return err
			}
			status, err := mb.Status(req.Context())
			if err != nil {
				return err
			}

			if status.HasPending() {
				return errors.Errorf("migrations have not yet been fully applied")
			}
			return nil
		},
	}
	maps.Copy(h, r.healthReadyCheckers) // possibly override default checkers

	r.healthH = healthx.NewHandler(r.Writer(), config.Version, h)

	return r.healthH
}

func (r *RegistryDefault) HealthServer() *health.Server {
	if r.healthServer == nil {
		r.healthServer = health.NewServer()
	}

	return r.healthServer
}

func (r *RegistryDefault) GetVersion(_ context.Context, _ *connect.Request[rts.GetVersionRequest]) (*connect.Response[rts.GetVersionResponse], error) {
	return connect.NewResponse(&rts.GetVersionResponse{Version: config.Version}), nil
}

func (r *RegistryDefault) Tracer(ctx context.Context) *otelx.Tracer {
	if r.tracer == nil {
		// Tracing is initialized only once, so it can not be hot reloaded or context-aware.
		t, err := otelx.New("Ory Keto", r.Logger(), r.Config(ctx).TracingConfig())
		if err != nil {
			r.Logger().WithError(err).Fatalf("Unable to initialize Tracer.")
		}

		// Wrap the tracer if required
		if r.tracerWrapper != nil {
			t = r.tracerWrapper(t)
		}

		r.tracer = t
	}

	return r.tracer
}

func (r *RegistryDefault) Logger() *logrusx.Logger {
	if r.l == nil {
		r.l = logrusx.New("Ory Keto", config.Version)
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

func (r *RegistryDefault) MappingManager() relationtuple.MappingManager {
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

func (r *RegistryDefault) NetworkID(ctx context.Context) uuid.UUID {
	if r.p == nil {
		panic("no persister, but expected to have one")
	}
	return r.p.NetworkID(ctx)
}

func (r *RegistryDefault) NewShardID() uuid.UUID {
	if r.replaceShardID != nil {
		return r.replaceShardID()
	}
	return uuid.Must(uuid.NewV4())
}

func (r *RegistryDefault) MapperNamespace(ctx context.Context) uuid.UUID {
	if r.replaceMapperNs != nil {
		return r.replaceMapperNs(ctx)
	}
	return r.p.NetworkID(ctx)
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

func (r *RegistryDefault) Checker() check.Checker {
	if r.ck == nil {
		r.ck = step.NewExecutor(r)
	}
	return r.ck
}

func (r *RegistryDefault) MigrationBox(ctx context.Context, opts ...popx.MigrationBoxOption) (*popx.MigrationBox, error) {
	if r.mb == nil {
		c, err := r.PopConnection(ctx)
		if err != nil {
			return nil, err
		}
		namespaces, err := r.Config(ctx).NamespaceManager()
		if err != nil {
			return nil, err
		}

		migrationBoxOptions := append(opts,
			popx.WithGoMigrations(uuidmapping.Migrations(namespaces)))

		migrationBoxOptions = append(migrationBoxOptions, r.defaultMigrationOptions...)

		mb, err := popx.NewMigrationBox(
			fsx.Merge(append([]fs.FS{sql.Migrations, networkx.Migrations}, r.extraMigrations...)...),
			c, r.Logger(),
			migrationBoxOptions...,
		)
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
	return mb.Down(ctx, -1)
}

func (r *RegistryDefault) DetermineNetwork(ctx context.Context) (*networkx.Network, error) {
	c, err := r.PopConnection(ctx)
	if err != nil {
		return nil, err
	}
	mb, err := popx.NewMigrationBox(networkx.Migrations, c, r.Logger())
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

	return networkx.Determine(c.WithContext(ctx))
}

func (r *RegistryDefault) InitWithoutNetworkID(ctx context.Context) error {
	r.init1.Do(func() {
		if dbal.IsMemorySQLite(r.Config(ctx).DSN()) {
			mb, err := r.MigrationBox(ctx)
			if err != nil {
				r.init1err = err
				return
			}

			if err := mb.Up(ctx); err != nil {
				r.init1err = err
				return
			}
		}

		p, err := sql.NewPersister(ctx, r, uuid.Nil)
		if err != nil {
			r.init1err = err
			return
		}
		r.p = p
		r.initEngines()
	})
	return r.init1err
}

func (r *RegistryDefault) initEngines() {
	_ = r.PermissionEngine()
	_ = r.ExpandEngine()
	_ = r.Checker()
}

func (r *RegistryDefault) Init(ctx context.Context) (err error) {
	r.init2.Do(func() {
		if err := r.InitWithoutNetworkID(ctx); err != nil {
			r.init2err = err
			return
		}

		network, err := r.DetermineNetwork(ctx)
		if err != nil {
			r.init2err = err
			return
		}

		r.p.SetNetwork(network.ID)
	})
	return r.init2err
}
