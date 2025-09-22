// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package sql

import (
	"context"
	"database/sql"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/ory/pop/v6"
	"github.com/ory/x/otelx"
	keysetpagination "github.com/ory/x/pagination/keysetpagination_v2"
	"github.com/ory/x/sqlcon"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/ketoapi"
)

// Typical database limits for placeholders/bind vars are 1<<15 (32k, MySQL, SQLite) and 1<<16 (64k, PostgreSQL, CockroachDB).
const (
	chunkSizeInsertUUIDMappings = 15000 // two placeholders per mapping
	chunkSizeInsertTuple        = 3000  // ten placeholders per tuple
	chunkSizeDeleteTuple        = 100   // the database must build an expression tree for each chunk, so we must limit more aggressively
)

type (
	RelationTuple struct {
		// An ID field is required to make pop happy. The actual ID is a composite primary key.
		ID                  uuid.UUID      `db:"shard_id"`
		NetworkID           uuid.UUID      `db:"nid"`
		Namespace           string         `db:"namespace"`
		Object              uuid.UUID      `db:"object"`
		Relation            string         `db:"relation"`
		SubjectID           uuid.NullUUID  `db:"subject_id"`
		SubjectSetNamespace sql.NullString `db:"subject_set_namespace"`
		SubjectSetObject    uuid.NullUUID  `db:"subject_set_object"`
		SubjectSetRelation  sql.NullString `db:"subject_set_relation"`
		CommitTime          time.Time      `db:"commit_time"`
	}
	relationTuples []RelationTuple
)

func (relationTuples) TableName() string {
	return "keto_relation_tuples"
}

func (*RelationTuple) TableName() string {
	return "keto_relation_tuples"
}

func (r *RelationTuple) ToInternal() *relationtuple.RelationTuple {
	if r == nil {
		return nil
	}

	rt := &relationtuple.RelationTuple{
		Relation:  r.Relation,
		Object:    r.Object,
		Namespace: r.Namespace,
	}

	if r.SubjectID.Valid {
		rt.Subject = &relationtuple.SubjectID{
			ID: r.SubjectID.UUID,
		}
	} else {
		rt.Subject = &relationtuple.SubjectSet{
			Namespace: r.SubjectSetNamespace.String,
			Object:    r.SubjectSetObject.UUID,
			Relation:  r.SubjectSetRelation.String,
		}
	}

	return rt
}

func (r *RelationTuple) insertSubject(s relationtuple.Subject) error {
	switch st := s.(type) {
	case *relationtuple.SubjectID:
		r.SubjectID = uuid.NullUUID{
			UUID:  st.ID,
			Valid: true,
		}
		r.SubjectSetNamespace = sql.NullString{}
		r.SubjectSetObject = uuid.NullUUID{}
		r.SubjectSetRelation = sql.NullString{}
	case *relationtuple.SubjectSet:
		r.SubjectID = uuid.NullUUID{}
		_ = r.SubjectSetNamespace.Scan(st.Namespace)
		_ = r.SubjectSetObject.Scan(st.Object)
		_ = r.SubjectSetRelation.Scan(st.Relation)
	}
	return nil
}

func (r *RelationTuple) FromInternal(rt *relationtuple.RelationTuple) (err error) {
	r.Namespace = rt.Namespace
	r.Object = rt.Object
	r.Relation = rt.Relation

	return r.insertSubject(rt.Subject)
}

func (p *Persister) whereSubject(_ context.Context, q *pop.Query, sub relationtuple.Subject) error {
	switch s := sub.(type) {
	case *relationtuple.SubjectID:
		q.
			Where("subject_id = ?", s.ID).
			// NULL checks to leverage partial indexes
			Where("subject_set_namespace IS NULL").
			Where("subject_set_object IS NULL").
			Where("subject_set_relation IS NULL")
	case *relationtuple.SubjectSet:
		q.
			Where("subject_set_namespace = ?", s.Namespace).
			Where("subject_set_object = ?", s.Object).
			Where("subject_set_relation = ?", s.Relation).
			// NULL checks to leverage partial indexes
			Where("subject_id IS NULL")
	case nil:
		return errors.WithStack(ketoapi.ErrNilSubject)
	}
	return nil
}

func (p *Persister) whereQuery(ctx context.Context, q *pop.Query, rq *relationtuple.RelationQuery) error {
	if rq.Namespace != nil {
		q.Where("namespace = ?", *rq.Namespace)
	}
	if rq.Object != nil {
		q.Where("object = ?", *rq.Object)
	}
	if rq.Relation != nil {
		q.Where("relation = ?", *rq.Relation)
	}
	if s := rq.Subject; s != nil {
		if err := p.whereSubject(ctx, q, s); err != nil {
			return err
		}
	}
	return nil
}

func buildDelete(nid uuid.UUID, rs []*relationtuple.RelationTuple) (query string, args []any, err error) {
	if len(rs) == 0 {
		return "", nil, errors.WithStack(ketoapi.ErrMalformedInput)
	}

	args = make([]any, 0, 6*len(rs)+1)
	ors := make([]string, 0, len(rs))
	for _, rt := range rs {
		switch s := rt.Subject.(type) {
		case *relationtuple.SubjectID:
			ors = append(ors, "(namespace = ? AND object = ? AND relation = ? AND subject_id = ? AND subject_set_namespace IS NULL AND subject_set_object IS NULL AND subject_set_relation IS NULL)")
			args = append(args, rt.Namespace, rt.Object, rt.Relation, s.ID)
		case *relationtuple.SubjectSet:
			ors = append(ors, "(namespace = ? AND object = ? AND relation = ? AND subject_id IS NULL AND subject_set_namespace = ? AND subject_set_object = ? AND subject_set_relation = ?)")
			args = append(args, rt.Namespace, rt.Object, rt.Relation, s.Namespace, s.Object, s.Relation)
		case nil:
			return "", nil, errors.WithStack(ketoapi.ErrNilSubject)
		}
	}

	query = fmt.Sprintf("DELETE FROM %s WHERE (%s) AND nid = ?", (&RelationTuple{}).TableName(), strings.Join(ors, " OR "))
	args = append(args, nid)
	return query, args, nil
}

func (p *Persister) DeleteRelationTuples(ctx context.Context, rs ...*relationtuple.RelationTuple) (err error) {
	if len(rs) == 0 {
		return nil
	}

	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.DeleteRelationTuples",
		trace.WithAttributes(attribute.Int("count", len(rs))))
	defer otelx.End(span, &err)

	return p.Transaction(ctx, func(ctx context.Context) error {
		for chunk := range slices.Chunk(rs, chunkSizeDeleteTuple) {
			q, args, err := buildDelete(p.NetworkID(ctx), chunk)
			if err != nil {
				return err
			}
			if q == "" {
				continue
			}
			if err := p.Connection(ctx).RawQuery(q, args...).Exec(); err != nil {
				return sqlcon.HandleError(err)
			}
		}
		return nil
	})
}

func (p *Persister) DeleteAllRelationTuples(ctx context.Context, query *relationtuple.RelationQuery) (err error) {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.DeleteAllRelationTuples")
	defer otelx.End(span, &err)

	return p.Transaction(ctx, func(ctx context.Context) error {
		sqlQuery := p.queryWithNetwork(ctx)
		err := p.whereQuery(ctx, sqlQuery, query)
		if err != nil {
			return err
		}

		var res relationTuples
		return sqlQuery.Delete(&res)
	})
}

func (p *Persister) GetRelationTuples(ctx context.Context, query *relationtuple.RelationQuery, pageOpts ...keysetpagination.Option) (_ []*relationtuple.RelationTuple, _ *keysetpagination.Paginator, err error) {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.GetRelationTuples")
	defer otelx.End(span, &err)

	paginator := keysetpagination.NewPaginator(append(pageOpts,
		keysetpagination.WithDefaultToken(keysetpagination.NewPageToken(keysetpagination.Column{Name: "shard_id", Value: uuid.Nil})),
	)...)

	sqlQuery := p.queryWithNetwork(ctx).Scope(keysetpagination.Paginate[*RelationTuple](paginator))

	err = p.whereQuery(ctx, sqlQuery, query)
	if err != nil {
		return nil, nil, err
	}
	var res relationTuples
	if err := sqlQuery.All(&res); err != nil {
		return nil, nil, sqlcon.HandleError(err)
	}

	res, nextPage := keysetpagination.Result(res, paginator)
	internalRes := make([]*relationtuple.RelationTuple, len(res))
	for i := range res {
		internalRes[i] = res[i].ToInternal()
	}

	return internalRes, nextPage, nil
}

func (p *Persister) ExistsRelationTuples(ctx context.Context, query *relationtuple.RelationQuery) (_ bool, err error) {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.ExistsRelationTuples")
	defer otelx.End(span, &err)

	sqlQuery := p.queryWithNetwork(ctx)

	err = p.whereQuery(ctx, sqlQuery, query)
	if err != nil {
		return false, err
	}
	exists, err := sqlQuery.Exists(&RelationTuple{})
	return exists, sqlcon.HandleError(err)
}

func buildInsert(commitTime time.Time, nid uuid.UUID, rs []*relationtuple.RelationTuple) (query string, args []any, err error) {
	if len(rs) == 0 {
		return "", nil, errors.WithStack(ketoapi.ErrMalformedInput)
	}

	var q strings.Builder
	fmt.Fprintf(&q, "INSERT INTO %s (shard_id, nid, namespace, object, relation, subject_id, subject_set_namespace, subject_set_object, subject_set_relation, commit_time) VALUES ", (&RelationTuple{}).TableName())
	const placeholders = "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	const separator = ", "
	q.Grow(len(rs) * (len(placeholders) + len(separator)))
	args = make([]any, 0, 10*len(rs))

	for i, r := range rs {
		if r.Subject == nil {
			return "", nil, errors.WithStack(ketoapi.ErrNilSubject)
		}

		rt := &RelationTuple{
			ID:         uuid.Must(uuid.NewV4()),
			NetworkID:  nid,
			CommitTime: commitTime,
		}
		if err := rt.FromInternal(r); err != nil {
			return "", nil, err
		}

		if i > 0 {
			q.WriteString(separator)
		}
		q.WriteString(placeholders)
		args = append(args, rt.ID, rt.NetworkID, rt.Namespace, rt.Object, rt.Relation, rt.SubjectID, rt.SubjectSetNamespace, rt.SubjectSetObject, rt.SubjectSetRelation, rt.CommitTime)
	}

	query = q.String()
	return query, args, nil
}

func (p *Persister) WriteRelationTuples(ctx context.Context, rs ...*relationtuple.RelationTuple) (err error) {
	if len(rs) == 0 {
		return nil
	}

	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.WriteRelationTuples",
		trace.WithAttributes(attribute.Int("count", len(rs))))
	defer otelx.End(span, &err)

	commitTime := time.Now()

	return p.Transaction(ctx, func(ctx context.Context) error {
		for chunk := range slices.Chunk(rs, chunkSizeInsertTuple) {
			q, args, err := buildInsert(commitTime, p.NetworkID(ctx), chunk)
			if err != nil {
				return err
			}
			if err := p.Connection(ctx).RawQuery(q, args...).Exec(); err != nil {
				return sqlcon.HandleError(err)
			}
		}
		return nil
	})
}

func (p *Persister) TransactRelationTuples(ctx context.Context, ins []*relationtuple.RelationTuple, del []*relationtuple.RelationTuple) (err error) {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.TransactRelationTuples")
	defer otelx.End(span, &err)

	if len(ins)+len(del) == 0 {
		return nil
	}

	return p.Transaction(ctx, func(ctx context.Context) error {
		if err := p.WriteRelationTuples(ctx, ins...); err != nil {
			return err
		}
		return p.DeleteRelationTuples(ctx, del...)
	})
}
