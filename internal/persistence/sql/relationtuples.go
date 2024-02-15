// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package sql

import (
	"context"
	"database/sql"
	"time"

	"github.com/ory/keto/ketoapi"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/ory/x/otelx"
	"github.com/ory/x/sqlcon"
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
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
	relationTuples []*RelationTuple
)

func (relationTuples) TableName() string {
	return "keto_relation_tuples"
}

func (*RelationTuple) TableName() string {
	return "keto_relation_tuples"
}

func (r *RelationTuple) ToInternal() (*relationtuple.RelationTuple, error) {
	if r == nil {
		return nil, nil
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

	return rt, nil
}

func (r *RelationTuple) insertSubject(_ context.Context, s relationtuple.Subject) error {
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

func (r *RelationTuple) FromInternal(ctx context.Context, p *Persister, rt *relationtuple.RelationTuple) (err error) {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.FromInternal")
	defer otelx.End(span, &err)

	r.Namespace = rt.Namespace
	r.Object = rt.Object
	r.Relation = rt.Relation

	return r.insertSubject(ctx, rt.Subject)
}

func (p *Persister) InsertRelationTuple(ctx context.Context, rel *relationtuple.RelationTuple) (err error) {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.InsertRelationTuple")
	defer otelx.End(span, &err)

	if rel.Subject == nil {
		return errors.WithStack(ketoapi.ErrNilSubject)
	}

	rt := &RelationTuple{
		ID:         uuid.Must(uuid.NewV4()),
		CommitTime: time.Now(),
	}
	if err := rt.FromInternal(ctx, p, rel); err != nil {
		return err
	}

	if err := sqlcon.HandleError(
		p.createWithNetwork(ctx, rt),
	); err != nil {
		return err
	}
	return nil
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
		q.Where("namespace = ?", rq.Namespace)
	}
	if rq.Object != nil {
		q.Where("object = ?", rq.Object)
	}
	if rq.Relation != nil {
		q.Where("relation = ?", rq.Relation)
	}
	if s := rq.Subject; s != nil {
		if err := p.whereSubject(ctx, q, s); err != nil {
			return err
		}
	}
	return nil
}

func (p *Persister) DeleteRelationTuples(ctx context.Context, rs ...*relationtuple.RelationTuple) (err error) {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.DeleteRelationTuples")
	defer otelx.End(span, &err)

	return p.Transaction(ctx, func(ctx context.Context) error {
		for _, r := range rs {
			q := p.queryWithNetwork(ctx).
				Where("namespace = ?", r.Namespace).
				Where("object = ?", r.Object).
				Where("relation = ?", r.Relation)
			if err := p.whereSubject(ctx, q, r.Subject); err != nil {
				return err
			}

			if err := q.Delete(&RelationTuple{}); err != nil {
				return err
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

func (p *Persister) GetRelationTuples(ctx context.Context, query *relationtuple.RelationQuery, options ...x.PaginationOptionSetter) (_ []*relationtuple.RelationTuple, nextPageToken string, err error) {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.GetRelationTuples")
	defer otelx.End(span, &err)

	pagination, err := internalPaginationFromOptions(options...)
	if err != nil {
		return nil, "", err
	}

	sqlQuery := p.queryWithNetwork(ctx).
		Order("shard_id, nid").
		Where("shard_id > ?", pagination.LastID).
		Limit(pagination.PerPage + 1)

	err = p.whereQuery(ctx, sqlQuery, query)
	if err != nil {
		return nil, "", err
	}
	var res relationTuples
	if err := sqlQuery.All(&res); err != nil {
		return nil, "", sqlcon.HandleError(err)
	}
	if len(res) == 0 {
		return make([]*relationtuple.RelationTuple, 0), "", nil
	}

	if len(res) > pagination.PerPage {
		res = res[:len(res)-1]
		nextPageToken = pagination.encodeNextPageToken(res[len(res)-1].ID)
	}

	internalRes := make([]*relationtuple.RelationTuple, 0, len(res))
	for _, r := range res {
		if rt, err := r.ToInternal(); err == nil {
			// Ignore error here, which stems from a deleted namespace.
			internalRes = append(internalRes, rt)
		}
	}

	return internalRes, nextPageToken, nil
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

func (p *Persister) WriteRelationTuples(ctx context.Context, rs ...*relationtuple.RelationTuple) (err error) {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.WriteRelationTuples")
	defer otelx.End(span, &err)

	return p.Transaction(ctx, func(ctx context.Context) error {
		for _, r := range rs {
			if err := p.InsertRelationTuple(ctx, r); err != nil {
				return err
			}
		}
		return nil
	})
}

func (p *Persister) TransactRelationTuples(ctx context.Context, ins []*relationtuple.RelationTuple, del []*relationtuple.RelationTuple) (err error) {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.TransactRelationTuples")
	defer otelx.End(span, &err)

	return p.Transaction(ctx, func(ctx context.Context) error {
		if err := p.WriteRelationTuples(ctx, ins...); err != nil {
			return err
		}
		return p.DeleteRelationTuples(ctx, del...)
	})
}
