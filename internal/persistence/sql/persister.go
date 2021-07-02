package sql

import (
	"context"
	"embed"
	"fmt"
	"strconv"
	"time"

	"github.com/ory/keto/internal/driver/config"

	"github.com/gofrs/uuid"

	"github.com/luna-duclos/instrumentedsql"
	"github.com/luna-duclos/instrumentedsql/opentracing"
	"github.com/ory/x/tracing"

	"github.com/ory/x/popx"

	"github.com/cenkalti/backoff/v3"
	"github.com/ory/x/sqlcon"

	"github.com/gobuffalo/pop/v5"
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/internal/x"
)

type (
	Persister struct {
		conn            *pop.Connection
		mb              *popx.MigrationBox
		d               dependencies
		dsn             string
		tracer          *tracing.Tracer
		networkIDCached uuid.UUID
	}
	internalPagination struct {
		Page, PerPage int
	}
	dependencies interface {
		config.Provider
		x.LoggerProvider
	}
)

const (
	defaultPageSize int = 100
)

var (
	//go:embed migrations/sql/*.sql
	migrations embed.FS

	////go:embed namespace_migrations/*.sql
	//namespaceMigrations embed.FS

	_ persistence.Persister = &Persister{}
)

func NewPersister(dsn string, reg dependencies, tracer *tracing.Tracer) (*Persister, error) {
	pop.SetLogger(reg.Logger().PopLogger)

	p := &Persister{
		d:      reg,
		dsn:    dsn,
		tracer: tracer,
	}

	var err error
	p.conn, err = p.newConnection(nil)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Persister) newConnection(options map[string]string) (c *pop.Connection, err error) {
	var opts []instrumentedsql.Opt
	if p.tracer.IsLoaded() {
		opts = []instrumentedsql.Opt{
			instrumentedsql.WithTracer(opentracing.NewTracer(true)),
			instrumentedsql.WithOmitArgs(),
		}
	}
	pool, idlePool, connMaxLifetime, connMaxIdleTime, cleanedDSN := sqlcon.ParseConnectionOptions(p.d.Logger(), p.dsn)
	connDetails := &pop.ConnectionDetails{
		URL:                       sqlcon.FinalizeDSN(p.d.Logger(), cleanedDSN),
		IdlePool:                  idlePool,
		ConnMaxLifetime:           connMaxLifetime,
		ConnMaxIdleTime:           connMaxIdleTime,
		Pool:                      pool,
		Options:                   options,
		UseInstrumentedDriver:     p.tracer != nil && p.tracer.IsLoaded(),
		InstrumentedDriverOptions: opts,
	}

	bc := backoff.NewExponentialBackOff()
	bc.MaxElapsedTime = time.Minute * 5
	bc.Reset()

	if err := backoff.Retry(func() (err error) {
		c, err = pop.NewConnection(connDetails)
		if err != nil {
			p.d.Logger().WithError(err).Error("Unable to connect to database, retrying.")
			return errors.WithStack(err)
		}

		if err := c.Open(); err != nil {
			p.d.Logger().WithError(err).Error("Unable to open the database connection, retrying.")
			return errors.WithStack(err)
		}

		if err := c.Store.(interface{ Ping() error }).Ping(); err != nil {
			p.d.Logger().WithError(err).Error("Unable to ping the database connection, retrying.")
			return errors.WithStack(err)
		}

		return nil
	}, bc); err != nil {
		return nil, errors.WithStack(err)
	}

	return c, nil
}

func (p *Persister) MigrationBox(ctx context.Context) (*popx.MigrationBox, error) {
	if p.mb == nil {
		var err error
		p.mb, err = popx.NewMigrationBox(migrations, popx.NewMigrator(p.Connection(ctx), p.d.Logger(), nil, 0))
		if err != nil {
			return nil, err
		}
	}

	return p.mb, nil
}

func (p *Persister) Connection(ctx context.Context) *pop.Connection {
	return popx.GetConnection(ctx, p.conn.WithContext(ctx))
}

func (p *Persister) transaction(ctx context.Context, f func(ctx context.Context, c *pop.Connection) error) error {
	return popx.Transaction(ctx, p.conn.WithContext(ctx), f)
}

func internalPaginationFromOptions(opts ...x.PaginationOptionSetter) (*internalPagination, error) {
	xp := x.GetPaginationOptions(opts...)
	ip := &internalPagination{
		PerPage: xp.Size,
	}
	if ip.PerPage == 0 {
		ip.PerPage = defaultPageSize
	}
	return ip, ip.parsePageToken(xp.Token)
}

func (p *internalPagination) parsePageToken(t string) error {
	if t == "" {
		p.Page = 1
		return nil
	}

	i, err := strconv.ParseUint(t, 10, 32)
	if err != nil {
		return errors.WithStack(persistence.ErrMalformedPageToken)
	}

	p.Page = int(i)
	return nil
}

func (p *internalPagination) encodeNextPageToken() string {
	return fmt.Sprintf("%d", p.Page+1)
}
