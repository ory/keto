// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package sql

import (
	"context"
	"embed"

	"github.com/gofrs/uuid"
	"github.com/ory/pop/v6"
	"github.com/ory/x/contextx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/otelx"
	"github.com/ory/x/popx"

	keysetpagination "github.com/ory/x/pagination/keysetpagination_v2"
	"github.com/ory/x/pagination/paginationplanner"

	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/persistence"
)

type (
	Persister struct {
		conn *pop.Connection
		d    dependencies
		nid  uuid.UUID

		planner *paginationplanner.PaginationPlanner
	}
	dependencies interface {
		logrusx.Provider
		otelx.Provider
		contextx.Provider
		config.Provider

		PopConnection(ctx context.Context) (*pop.Connection, error)
	}
)

var (
	//go:embed migrations/sql/*.sql
	Migrations embed.FS

	_ persistence.Persister = &Persister{}
)

// Pagination Planner for relation_tuple table
var (
	t                           = paginationplanner.NewTable()
	TupleColNamespace           = t.NewColumn("namespace")
	TupleColObject              = t.NewColumn("object")
	TupleColRel                 = t.NewColumn("relation")
	TupleColSubjectId           = t.NewColumn("subject_id")
	TupleColSubjectSetNamespace = t.NewColumn("subject_set_namespace")
	TupleColSubjectSetObject    = t.NewColumn("subject_set_object")
	TupleColSubjectSetRel       = t.NewColumn("subject_set_relation")
)

var (
	fullIdxPlan = paginationplanner.PaginationPlan{
		DefaultPageToken: keysetpagination.NewPageToken([]keysetpagination.Column{
			{Name: TupleColNamespace.Name(), Order: keysetpagination.OrderAscending, Value: ""},
			{Name: TupleColObject.Name(), Order: keysetpagination.OrderAscending, Value: uuid.Nil},
			{Name: TupleColRel.Name(), Order: keysetpagination.OrderAscending, Value: ""},
			{Name: TupleColSubjectId.Name(), Order: keysetpagination.OrderAscending, Nullable: true},
			{Name: TupleColSubjectSetNamespace.Name(), Order: keysetpagination.OrderAscending, Nullable: true},
			{Name: TupleColSubjectSetObject.Name(), Order: keysetpagination.OrderAscending, Nullable: true},
			{Name: TupleColSubjectSetRel.Name(), Order: keysetpagination.OrderAscending, Nullable: true},
			{Name: "commit_time", Order: keysetpagination.OrderAscending, Value: "1970-01-01"},
			{Name: "shard_id", Order: keysetpagination.OrderAscending, Value: uuid.Nil},
		}...),
		ApplicableQueries: [][]paginationplanner.Column{
			{},
			{TupleColNamespace},
			{TupleColNamespace, TupleColObject},
			{TupleColNamespace, TupleColObject, TupleColRel},
		},
	}

	fallbackPlan = paginationplanner.PaginationPlan{
		DefaultPageToken: keysetpagination.NewPageToken([]keysetpagination.Column{
			{Name: "shard_id", Order: keysetpagination.OrderAscending, Value: uuid.Nil},
		}...),
	}
)

func NewPersister(ctx context.Context, reg dependencies, nid uuid.UUID) (*Persister, error) {
	conn, err := reg.PopConnection(ctx)
	if err != nil {
		return nil, err
	}

	planner, err := paginationplanner.NewPaginationPlanner(fallbackPlan, []paginationplanner.PaginationPlan{fullIdxPlan})
	if err != nil {
		return nil, err
	}
	p := &Persister{
		d:       reg,
		nid:     nid,
		conn:    conn,
		planner: planner,
	}
	return p, nil
}

func (p *Persister) Connection(ctx context.Context) *pop.Connection {
	return popx.GetConnection(ctx, p.conn.WithContext(ctx))
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

func (p *Persister) SetNetwork(nid uuid.UUID) {
	p.nid = nid
}
