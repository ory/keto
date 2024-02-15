// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package sql

import (
	"context"
	"embed"
	"reflect"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/ory/x/otelx"
	"github.com/ory/x/popx"
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoctx"
)

type (
	Persister struct {
		conn *pop.Connection
		d    dependencies
		nid  uuid.UUID
	}
	internalPagination struct {
		PerPage int
		LastID  uuid.UUID
	}
	dependencies interface {
		x.LoggerProvider
		x.TracingProvider
		ketoctx.ContextualizerProvider
		config.Provider

		PopConnection(ctx context.Context) (*pop.Connection, error)
	}
)

const (
	defaultPageSize int = 100
)

var (
	//go:embed migrations/sql/*.sql
	Migrations embed.FS

	_ persistence.Persister = &Persister{}
)

func NewPersister(ctx context.Context, reg dependencies, nid uuid.UUID) (*Persister, error) {
	conn, err := reg.PopConnection(ctx)
	if err != nil {
		return nil, err
	}

	p := &Persister{
		d:    reg,
		nid:  nid,
		conn: conn,
	}

	return p, nil
}

func (p *Persister) Connection(ctx context.Context) *pop.Connection {
	return popx.GetConnection(ctx, p.conn.WithContext(ctx))
}

func (p *Persister) createWithNetwork(ctx context.Context, v interface{}) (err error) {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.createWithNetwork")
	defer otelx.End(span, &err)

	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Ptr && rv.Elem().Kind() != reflect.Struct {
		panic("expected to get *struct in create")
	}
	nID := rv.Elem().FieldByName("NetworkID")
	if !nID.IsValid() || !nID.CanSet() {
		panic("expected struct to have a 'NetworkID uuid.UUID' field")
	}
	nID.Set(reflect.ValueOf(p.NetworkID(ctx)))

	return p.Connection(ctx).Create(v)
}

func (p *Persister) queryWithNetwork(ctx context.Context) *pop.Query {
	return p.Connection(ctx).Where("nid = ?", p.NetworkID(ctx))
}

func (p *Persister) Transaction(ctx context.Context, f func(ctx context.Context) error) error {
	return popx.Transaction(ctx, p.conn.WithContext(ctx), func(ctx context.Context, _ *pop.Connection) error { return f(ctx) })
}

func (p *Persister) NetworkID(ctx context.Context) uuid.UUID {
	return p.d.Contextualizer().Network(ctx, p.nid)
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
		p.LastID = uuid.Nil
		return nil
	}

	i, err := uuid.FromString(t)
	if err != nil {
		return errors.WithStack(persistence.ErrMalformedPageToken)
	}

	p.LastID = i
	return nil
}

func (p *internalPagination) encodeNextPageToken(lastID uuid.UUID) string {
	return lastID.String()
}
