package driver

import (
	"github.com/jmoiron/sqlx"
	"github.com/ory/keto/storage"
	"github.com/ory/x/dbal"
	"github.com/ory/x/sqlcon"
	"github.com/ory/x/urlx"
	"time"
)

type RegistrySQL struct {
	*RegistryBase
	db          *sqlx.DB
	sm          storage.Manager
	dbalOptions []sqlcon.OptionModifier
}

var _ Registry = new(RegistrySQL)

func init() {
	dbal.RegisterDriver(NewRegistrySQL())
}

func NewRegistrySQL() *RegistrySQL {
	r := &RegistrySQL{
		RegistryBase: new(RegistryBase),
	}
	r.RegistryBase.with(r)
	return r
}

func (m *RegistrySQL) DB() *sqlx.DB {
	if m.db == nil {
		if err := m.Init(); err != nil {
			m.Logger().WithError(err).Fatalf("Unable to initialize database.")
		}
	}

	return m.db
}

func (m *RegistrySQL) Init() error {
	if m.db != nil {
		return nil
	}

	options := append([]sqlcon.OptionModifier{}, m.dbalOptions...)
	if m.Tracer().IsLoaded() {
		options = append(options, sqlcon.WithDistributedTracing(), sqlcon.WithOmitArgsFromTraceSpans())
	}

	connection, err := sqlcon.NewSQLConnection(m.c.DSN(), m.Logger(), options...)
	if err != nil {
		return err
	}

	m.db, err = connection.GetDatabaseRetry(time.Second*5, time.Minute*5)
	if err != nil {
		return err
	}

	return err
}

func (m *RegistrySQL) StorageManager() storage.Manager {
	if m.sm == nil {
		m.sm = storage.NewSQLManager(m.DB())
	}
	return m.sm
}

func (m *RegistrySQL) CanHandle(dsn string) bool {
	s := dbal.Canonicalize(urlx.ParseOrFatal(m.l, dsn).Scheme)
	return s == dbal.DriverMySQL || s == dbal.DriverPostgreSQL
}

func (m *RegistrySQL) Ping() error {
	return m.DB().Ping()
}
