package driver

import (
	"context"

	"github.com/pkg/errors"

	"github.com/gobuffalo/pop/v5"
	"github.com/ory/herodot"
	"github.com/ory/x/healthx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/tracing"

	"github.com/ory/keto/internal/driver/configuration"
	"github.com/ory/keto/internal/namespace"

	"github.com/ory/keto/internal/persistence/sql"

	"github.com/ory/keto/internal/persistence"

	"github.com/ory/keto/internal/expand"

	"github.com/ory/keto/internal/check"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

var _ relationtuple.ManagerProvider = &RegistryDefault{}
var _ x.WriterProvider = &RegistryDefault{}
var _ x.LoggerProvider = &RegistryDefault{}
var _ Registry = &RegistryDefault{}

type RegistryDefault struct {
	p    persistence.Persister
	l    *logrusx.Logger
	w    herodot.Writer
	ce   *check.Engine
	ee   *expand.Engine
	conn *pop.Connection
	c    configuration.Provider
	hh   *healthx.Handler

	version, hash, date string
}

func (r *RegistryDefault) CanHandle(dsn string) bool {
	return true
}

func (r *RegistryDefault) Ping() error {
	return r.conn.Open()
}

func (r *RegistryDefault) WithConfig(c configuration.Provider) Registry {
	r.c = c
	return r
}

func (r *RegistryDefault) WithLogger(l *logrusx.Logger) Registry {
	r.l = l
	return r
}

func (r *RegistryDefault) WithBuildInfo(version, hash, date string) Registry {
	r.version, r.hash, r.date = version, hash, date
	return r
}

func (r *RegistryDefault) BuildVersion() string {
	return r.version
}

func (r *RegistryDefault) BuildDate() string {
	return r.date
}

func (r *RegistryDefault) BuildHash() string {
	return r.hash
}

func (r *RegistryDefault) HealthHandler() *healthx.Handler {
	if r.hh == nil {
		r.hh = healthx.NewHandler(r.Writer(), r.version, healthx.ReadyCheckers{})
	}

	return r.hh
}

func (r *RegistryDefault) Tracer() *tracing.Tracer {
	panic("implement me")
}

func (r *RegistryDefault) Logger() *logrusx.Logger {
	if r.l == nil {
		r.l = logrusx.New("keto", "dev")
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

func (r *RegistryDefault) NamespaceManager() namespace.Manager {
	return r.c
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

func (r *RegistryDefault) Migrator() (persistence.Migrator, error) {
	return r.p.(persistence.Migrator), nil
}

func (r *RegistryDefault) Init() error {
	c, err := pop.NewConnection(&pop.ConnectionDetails{
		URL: "sqlite://:memory:?_fk=true",
	})
	if err != nil {
		return errors.WithStack(err)
	}

	r.conn = c
	if err := c.Open(); err != nil {
		return errors.WithStack(err)
	}

	r.p, err = sql.NewPersister(r.conn, r.Logger(), r.c)
	if err != nil {
		return err
	}

	m, err := r.Migrator()
	if err != nil {
		return err
	}
	if err := m.MigrateUp(context.Background()); err != nil {
		return err
	}

	namespaceConfigs := r.c.Namespaces()
	for _, n := range namespaceConfigs {
		s, err := r.NamespaceMigrator().NamespaceStatus(context.Background(), n.ID)
		if err != nil {
			if r.c.DSN() == configuration.DSNMemory {
				// auto migrate when DSN is memory
				if err := r.NamespaceMigrator().MigrateNamespaceUp(context.Background(), n); err != nil {
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
