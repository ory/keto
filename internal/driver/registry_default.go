package driver

import (
	"context"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/ory/x/sqlcon"

	"github.com/ory/keto/internal/driver/config"

	"github.com/pkg/errors"

	"github.com/gobuffalo/pop/v5"
	"github.com/ory/herodot"
	"github.com/ory/x/healthx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/tracing"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/internal/persistence/sql"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

var _ relationtuple.ManagerProvider = &RegistryDefault{}
var _ x.WriterProvider = &RegistryDefault{}
var _ x.LoggerProvider = &RegistryDefault{}
var _ Registry = &RegistryDefault{}

type RegistryDefault struct {
	p      persistence.Persister
	l      *logrusx.Logger
	w      herodot.Writer
	ce     *check.Engine
	ee     *expand.Engine
	conn   *pop.Connection
	c      config.Provider
	health *healthx.Handler
}

func (r *RegistryDefault) BuildVersion() string {
	return config.Version
}

func (r *RegistryDefault) BuildDate() string {
	return config.Date
}

func (r *RegistryDefault) BuildHash() string {
	return config.Commit
}

func (r *RegistryDefault) Config() config.Provider {
	return r.c
}

func (r *RegistryDefault) HealthHandler() *healthx.Handler {
	if r.health == nil {
		r.health = healthx.NewHandler(r.Writer(), config.Version, healthx.ReadyCheckers{})
	}

	return r.health
}

func (r *RegistryDefault) Tracer() *tracing.Tracer {
	panic("implement me")
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
	return r.p
}

func (r *RegistryDefault) NamespaceMigrator() namespace.Migrator {
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

func (r *RegistryDefault) Persister() persistence.Persister {
	return r.p
}

func (r *RegistryDefault) Migrator() persistence.Migrator {
	return r.p.(persistence.Migrator)
}

func (r *RegistryDefault) Init(ctx context.Context) error {
	bc := backoff.NewExponentialBackOff()
	bc.MaxElapsedTime = time.Minute * 5
	bc.Reset()

	if err := backoff.Retry(func() error {
		pool, idlePool, connMaxLifetime, cleanedDSN := sqlcon.ParseConnectionOptions(r.l, r.c.DSN())
		c, err := pop.NewConnection(&pop.ConnectionDetails{
			URL:             sqlcon.FinalizeDSN(r.l, cleanedDSN),
			IdlePool:        idlePool,
			ConnMaxLifetime: connMaxLifetime,
			Pool:            pool,
		})
		if err != nil {
			r.Logger().WithError(err).Warnf("Unable to connect to database, retrying.")
			return errors.WithStack(err)
		}

		r.conn = c
		if err := c.Open(); err != nil {
			r.Logger().WithError(err).Warnf("Unable to open the database connection, retrying.")
			return errors.WithStack(err)
		}

		return nil
	}, bc); err != nil {
		return err
	}

	nm, err := r.c.NamespaceManager()
	if err != nil {
		return err
	}

	r.p, err = sql.NewPersister(r.conn, r.Logger(), nm)
	if err != nil {
		return err
	}

	m := r.Migrator()
	if r.c.DSN() == config.DSNMemory {
		if err := m.MigrateUp(context.Background()); err != nil {
			return err
		}
	}

	namespaceConfigs, err := nm.Namespaces(ctx)
	if err != nil {
		return err
	}
	for _, n := range namespaceConfigs {
		s, err := r.NamespaceMigrator().NamespaceStatus(ctx, n.ID)
		if err != nil {
			if r.c.DSN() == config.DSNMemory {
				// auto migrate when DSN is memory
				if err := r.NamespaceMigrator().MigrateNamespaceUp(ctx, n); err != nil {
					r.l.WithError(err).Errorf("Could not auto-migrate namespace %s.", n.Name)
				}
				continue
			}

			r.l.Warnf("Namespace %s is defined in the config but not yet migrated. It is ignored until you explicitly migrate it.", n.Name)
			continue
		}

		if s.CurrentVersion != s.NextVersion {
			r.l.Warnf("Namespace %s is not migrated to the latest version, it will be ignored until you explicitly migrate it.", n.Name)
		}
	}

	return nil
}
