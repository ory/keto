package driver

import (
	"github.com/ory/herodot"
	"github.com/ory/keto/internal/driver/configuration"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/x/healthx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/tracing"

	"github.com/ory/keto/internal/persistence"

	"github.com/ory/keto/internal/expand"

	"github.com/ory/keto/internal/check"

	"github.com/ory/keto/internal/persistence/memory"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

var _ relationtuple.ManagerProvider = &RegistryDefault{}
var _ x.WriterProvider = &RegistryDefault{}
var _ x.LoggerProvider = &RegistryDefault{}
var _ Registry = &RegistryDefault{}

type RegistryDefault struct {
	p  persistence.Persister
	l  *logrusx.Logger
	w  herodot.Writer
	ce *check.Engine
	ee *expand.Engine
	c  configuration.Provider
}

func (r *RegistryDefault) CanHandle(dsn string) bool {
	panic("implement me")
}

func (r *RegistryDefault) Ping() error {
	panic("implement me")
}

func (r *RegistryDefault) Init() error {
	namespaceConfigs := r.c.Namespaces()
	for _, n := range namespaceConfigs {
		s, err := r.NamespaceManager().NamespaceStatus(n)

		if err != nil {
			if r.c.DSN() == configuration.DSNMemory {
				// auto migrate on memory
				if err := r.NamespaceManager().MigrateNamespaceUp(n); err != nil {
					r.l.WithError(err).Errorf("Could not auto-migrate namespace %s.", n.Name)
				}
				continue
			}

			r.l.Warnf("Namespace %s is defined in the config but not yet migrated. It is ignored until you explicitly migrate it.", n.Name)

			continue
		}

		r.l.Infof("Namespace %s is migrated to version %d.", n.Name, s.Version)
	}

	return nil
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
	return r
}

func (r *RegistryDefault) BuildVersion() string {
	panic("implement me")
}

func (r *RegistryDefault) BuildDate() string {
	panic("implement me")
}

func (r *RegistryDefault) BuildHash() string {
	panic("implement me")
}

func (r *RegistryDefault) HealthHandler() *healthx.Handler {
	panic("implement me")
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
	if r.p == nil {
		r.p = memory.NewPersister()
	}
	return r.p
}

func (r *RegistryDefault) NamespaceManager() namespace.Manager {
	if r.p == nil {
		r.p = memory.NewPersister()
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
