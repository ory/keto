// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"context"
	"time"

	"github.com/cenkalti/backoff/v3"
	"github.com/pkg/errors"

	"github.com/ory/pop/v6"
	"github.com/ory/x/sqlcon"
)

func (r *RegistryDefault) PopConnectionWithOpts(ctx context.Context, popOpts ...func(*pop.ConnectionDetails)) (*pop.Connection, error) {
	pool, idlePool, connMaxLifetime, connMaxIdleTime, cleanedDSN := sqlcon.ParseConnectionOptions(r.Logger(), r.Config(ctx).DSN())
	connDetails := &pop.ConnectionDetails{
		URL:             sqlcon.FinalizeDSN(r.Logger(), cleanedDSN),
		IdlePool:        idlePool,
		ConnMaxLifetime: connMaxLifetime,
		ConnMaxIdleTime: connMaxIdleTime,
		Pool:            pool,
		TracerProvider:  r.Tracer(ctx).Provider(),
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
		if err := conn.Close(); err != nil {
			r.Logger().WithError(err).Error("Unable to close the database connection.")
		}
	}()

	return conn.WithContext(ctx), nil
}

// PopConnection returns the standard connection that is kept for the whole time.
func (r *RegistryDefault) PopConnection(ctx context.Context) (*pop.Connection, error) {
	if r.conn == nil {
		var err error
		r.conn, err = r.PopConnectionWithOpts(ctx, r.dbOpts...)
		return r.conn, err
	}
	return r.conn, nil
}
