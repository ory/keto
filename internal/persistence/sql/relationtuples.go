package sql

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ory/herodot"

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
)

func (r *relationTuple) TableName() string {
	return tableFromNamespace(r.Namespace)
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
	const (
		whereRelation = "relation = ?"
		whereObject   = "object = ?"
		whereSubject  = "subject = ?"
	)

	pagination, err := internalPaginationFromOptions(options...)
	if err != nil {
		return nil, x.PageTokenEnd, err
	}

	var (
		where []string
		args  []interface{}
	)

	if query.Relation != "" {
		where = append(where, whereRelation)
		args = append(args, query.Relation)
	}

	if query.Object != "" {
		where = append(where, whereObject)
		args = append(args, query.Object)
	}

	if query.Subject != nil {
		where = append(where, whereSubject)
		args = append(args, query.Subject.String())
	}

	var res []*relationTuple
	var rawQuery string

	n, err := p.namespaces.GetNamespace(ctx, query.Namespace)
	if err != nil {
		return nil, "-1", err
	}

	if len(where) == 0 {
		rawQuery = fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d",
			tableFromNamespace(n),
			pagination.Limit+1,
			pagination.Offset,
		)
	} else {
		rawQuery = fmt.Sprintf("SELECT * FROM %s WHERE %s ORDER BY object, relation, subject, commit_time LIMIT %d OFFSET %d",
			tableFromNamespace(n),
			strings.Join(where, " AND "),
			pagination.Limit+1,
			pagination.Offset,
		)
	}

	if err := p.conn.
		RawQuery(rawQuery, args...).
		All(&res); err != nil {
		return nil, x.PageTokenEnd, errors.WithStack(err)
	}

	if len(res) == 0 {
		return nil, x.PageTokenEnd, errors.WithStack(herodot.ErrNotFound)
	}

	cutOff := 1
	nextPageToken := pagination.encodeNextPageToken()
	if len(res) <= pagination.Limit {
		nextPageToken = x.PageTokenEnd
		cutOff = 0
	}

	internalRes := make([]*relationtuple.InternalRelationTuple, len(res)-cutOff)
	for i, r := range res[:len(res)-cutOff] {
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
