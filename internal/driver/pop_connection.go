package driver

import (
	"context"
	"time"

	"github.com/cenkalti/backoff/v3"
	"github.com/gobuffalo/pop/v6"
	"github.com/luna-duclos/instrumentedsql"
	otelsql "github.com/ory/x/otelx/sql"
	"github.com/ory/x/sqlcon"
	"github.com/pkg/errors"
)

func (r *RegistryDefault) PopConnectionWithOpts(ctx context.Context, popOpts ...func(*pop.ConnectionDetails)) (*pop.Connection, error) {
	tracer := r.Tracer(ctx)

	var opts []instrumentedsql.Opt
	if tracer.IsLoaded() {
		opts = []instrumentedsql.Opt{
			instrumentedsql.WithTracer(otelsql.NewTracer()),
			instrumentedsql.WithOmitArgs(),
		}
	}
	pool, idlePool, connMaxLifetime, connMaxIdleTime, cleanedDSN := sqlcon.ParseConnectionOptions(r.Logger(), r.Config(ctx).DSN())
	connDetails := &pop.ConnectionDetails{
		URL:                       sqlcon.FinalizeDSN(r.Logger(), cleanedDSN),
		IdlePool:                  idlePool,
		ConnMaxLifetime:           connMaxLifetime,
		ConnMaxIdleTime:           connMaxIdleTime,
		Pool:                      pool,
		UseInstrumentedDriver:     tracer != nil && tracer.IsLoaded(),
		InstrumentedDriverOptions: opts,
	}
	for _, o := range popOpts {
		o(connDetails)
	}

	bc := backoff.NewExponentialBackOff()
	bc.MaxElapsedTime = time.Minute * 5
	bc.Reset()

	var conn *pop.Connection
	if err := backoff.Retry(func() (err error) {
		conn, err = pop.NewConnection(connDetails)
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

		return nil
	}, bc); err != nil {
		return nil, errors.WithStack(err)
	}

	// Close this connection when the context is closed.
	go func() {
		<-ctx.Done()
		conn.Close()
	}()

	return conn.WithContext(ctx), nil
}

// PopConnection returns the standard connection that is kept for the whole time.
func (r *RegistryDefault) PopConnection(ctx context.Context) (*pop.Connection, error) {
	if r.conn == nil {
		var err error
		r.conn, err = r.PopConnectionWithOpts(ctx)
		return r.conn, err
	}
	return r.conn, nil
}
