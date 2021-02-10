package sql

import (
	"context"
	"fmt"
	"time"

	"github.com/gobuffalo/pop/v5"

	"github.com/pkg/errors"

	"github.com/ory/keto/internal/namespace"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

type (
	relationTuple struct {
		ShardID    string               `db:"shard_id"`
		Object     string               `db:"object"`
		Relation   string               `db:"relation"`
		Subject    string               `db:"subject"`
		CommitTime time.Time            `db:"commit_time"`
		Namespace  *namespace.Namespace `db:"-"`
	}
	relationTuples []*relationTuple
	whereStmts     struct {
		stmt string
		arg  interface{}
	}
)

const (
	namespaceContextKey contextKeys = "namespace"
)

func (relationTuples) TableName(ctx context.Context) string {
	n, ok := ctx.Value(namespaceContextKey).(*namespace.Namespace)
	if n == nil || !ok {
		panic("namespace context key not set")
	}
	return tableFromNamespace(n)
}

func (p *Persister) insertRelationTuple(ctx context.Context, rel *relationtuple.InternalRelationTuple) error {
	commitTime := time.Now()

	n, err := p.namespaces.GetNamespace(ctx, rel.Namespace)
	if err != nil {
		return err
	}

	// TODO sharding
	shardID := "default"

	p.l.WithFields(rel.ToLoggerFields()).Trace("creating in database")

	return p.connection(ctx).RawQuery(fmt.Sprintf(
		"INSERT INTO %s (shard_id, object, relation, subject, commit_time) VALUES (?, ?, ?, ?, ?)", tableFromNamespace(n)),
		shardID, rel.Object, rel.Relation, rel.Subject.String(), commitTime,
	).Exec()
}

func (r *relationTuple) toInternal() (*relationtuple.InternalRelationTuple, error) {
	if r == nil {
		return nil, nil
	}

	sub, err := relationtuple.SubjectFromString(r.Subject)
	return &relationtuple.InternalRelationTuple{
		Relation:  r.Relation,
		Object:    r.Object,
		Namespace: r.Namespace.Name,
		Subject:   sub,
	}, err
}

func (p *Persister) GetRelationTuples(ctx context.Context, query *relationtuple.RelationQuery, options ...x.PaginationOptionSetter) ([]*relationtuple.InternalRelationTuple, string, error) {
	pagination, err := internalPaginationFromOptions(options...)
	if err != nil {
		return nil, x.PageTokenEnd, err
	}

	var wheres []whereStmts

	if query.Relation != "" {
		wheres = append(wheres, whereStmts{stmt: "relation = ?", arg: query.Relation})
	}

	if query.Object != "" {
		wheres = append(wheres, whereStmts{stmt: "object = ?", arg: query.Object})
	}

	if query.Subject != nil {
		wheres = append(wheres, whereStmts{stmt: "subject = ?", arg: query.Subject.String()})
	}

	n, err := p.namespaces.GetNamespace(ctx, query.Namespace)
	if err != nil {
		return nil, x.PageTokenEnd, err
	}

	sqlQuery := p.connection(context.WithValue(ctx, namespaceContextKey, n)).
		Order("object, relation, subject, commit_time").
		Paginate(pagination.Page, pagination.PerPage)

	for _, w := range wheres {
		sqlQuery = sqlQuery.Where(w.stmt, w.arg)
	}

	var res relationTuples
	if err := sqlQuery.All(&res); err != nil {
		return nil, x.PageTokenEnd, errors.WithStack(err)
	}

	nextPageToken := pagination.encodeNextPageToken()
	if sqlQuery.Paginator.Page >= sqlQuery.Paginator.TotalPages {
		nextPageToken = x.PageTokenEnd
	}

	internalRes := make([]*relationtuple.InternalRelationTuple, len(res))
	for i, r := range res {
		r.Namespace = n

		var err error
		internalRes[i], err = r.toInternal()
		if err != nil {
			return nil, x.PageTokenEnd, err
		}
	}

	return internalRes, nextPageToken, nil
}

func (p *Persister) WriteRelationTuples(ctx context.Context, rs ...*relationtuple.InternalRelationTuple) error {
	return p.transaction(ctx, func(ctx context.Context, _ *pop.Connection) error {
		for _, r := range rs {
			if err := p.insertRelationTuple(ctx, r); err != nil {
				return err
			}
		}
		return nil
	})
}
