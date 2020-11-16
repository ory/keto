package sql

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/gobuffalo/pop/v5"

	"github.com/ory/keto/relationtuple"
	"github.com/ory/keto/x"
)

type (
	relationTuple struct {
		ShardID    string    `db:"shard_id"`
		ObjectID   string    `db:"object_id"`
		Relation   string    `db:"relation"`
		Subject    string    `db:"subject"`
		CommitTime time.Time `db:"commit_time"`
		Namespace  string    `db:"-"`
	}
)

func (r *relationTuple) TableName() string {
	return sqlSafeTableFromNamespace(r.Namespace)
}

func (r *relationTuple) fromInternal(rel *relationtuple.InternalRelationTuple) *relationTuple {
	if rel == nil {
		return r
	}
	if r == nil {
		*r = relationTuple{}
	}

	r.Relation = rel.Relation
	r.ShardID = "default"

	if rel.Object != nil {
		r.Namespace = rel.Object.Namespace
		r.ObjectID = rel.Object.ID
	}

	if rel.Subject != nil {
		r.Subject = rel.Subject.String()
	}

	return r
}

func (r *relationTuple) toInternal() *relationtuple.InternalRelationTuple {
	if r == nil {
		return nil
	}

	return &relationtuple.InternalRelationTuple{
		Relation: r.Relation,
		Object: &relationtuple.Object{
			ID:        r.ObjectID,
			Namespace: r.Namespace,
		},
		Subject: relationtuple.SubjectFromString(r.Subject),
	}
}

func (p *Persister) GetRelationTuples(_ context.Context, query *relationtuple.RelationQuery, options ...x.PaginationOptionSetter) ([]*relationtuple.InternalRelationTuple, error) {
	const (
		whereRelation = "relation = ?"
		whereObjectID = "object_id = ?"
		whereSubject  = "subject = ?"
	)

	var (
		where []string
		args  []interface{}
	)

	if query.Relation != "" {
		where = append(where, whereRelation)
		args = append(args, query.Relation)
	}

	if query.ObjectID != "" {
		where = append(where, whereObjectID)
		args = append(args, query.ObjectID)
	}

	if query.Subject != nil {
		where = append(where, whereSubject)
		args = append(args, query.Subject.String())
	}

	var res []*relationTuple
	var rawQuery string
	pagination := x.GetPaginationOptions(options...)

	if len(where) == 0 {
		rawQuery = fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d",
			sqlSafeTableFromNamespace(query.Namespace),
			pagination.PerPage,
			pagination.Page*pagination.PerPage,
		)
	} else {
		rawQuery = fmt.Sprintf("SELECT * FROM %s WHERE %s LIMIT %d OFFSET %d",
			sqlSafeTableFromNamespace(query.Namespace),
			strings.Join(where, " AND "),
			pagination.PerPage,
			pagination.Page*pagination.PerPage,
		)
	}

	if err := p.conn.
		RawQuery(rawQuery, args...).
		All(&res); err != nil {
		return nil, err
	}

	internalRes := make([]*relationtuple.InternalRelationTuple, len(res))
	for i, r := range res {
		internalRes[i] = r.toInternal()
	}
	return internalRes, nil
}

func (r *relationTuple) insert(c *pop.Connection) error {
	// TODO fix setting the commit time?
	r.CommitTime = time.Now()

	t := reflect.TypeOf(*r)
	v := reflect.ValueOf(*r)

	var rows []string
	var vals []interface{}
	for i := 0; i < t.NumField(); i++ {
		row, ok := t.Field(i).Tag.Lookup("db")
		if !ok || row == "-" {
			break
		}

		rows = append(rows, row)
		vals = append(vals, v.Field(i).Interface())
	}

	placeholders := strings.Repeat("?, ", len(rows))
	placeholders = placeholders[:len(placeholders)-2] // remove the last ", "

	return c.RawQuery(
		fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
			sqlSafeTableFromNamespace(r.Namespace),
			strings.Join(rows, ", "),
			placeholders),
		vals...).Exec()
}

func (p *Persister) WriteRelationTuples(_ context.Context, rs ...*relationtuple.InternalRelationTuple) error {
	return p.conn.Transaction(func(tx *pop.Connection) error {
		for _, r := range rs {
			if err := (&relationTuple{}).fromInternal(r).insert(tx); err != nil {
				return err
			}
		}
		return nil
	})
}
