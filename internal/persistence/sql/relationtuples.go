package sql

import (
	"context"
	"database/sql"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/ory/x/sqlcon"
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

type (
	RelationTuple struct {
		// An ID field is required to make pop happy. The actual ID is a composite primary key.
		ID                    uuid.UUID      `db:"shard_id"`
		NetworkID             uuid.UUID      `db:"nid"`
		NamespaceID           int32          `db:"namespace_id"`
		Object                uuid.UUID      `db:"object"`
		Relation              string         `db:"relation"`
		SubjectID             uuid.NullUUID  `db:"subject_id"`
		SubjectSetNamespaceID sql.NullInt32  `db:"subject_set_namespace_id"`
		SubjectSetObject      uuid.NullUUID  `db:"subject_set_object"`
		SubjectSetRelation    sql.NullString `db:"subject_set_relation"`
		CommitTime            time.Time      `db:"commit_time"`
	}
	relationTuples []*RelationTuple
)

func (relationTuples) TableName() string {
	return "keto_relation_tuples"
}

func (RelationTuple) TableName() string {
	return "keto_relation_tuples"
}

func (r *RelationTuple) toInternal() (*relationtuple.InternalRelationTuple, error) {
	if r == nil {
		return nil, nil
	}

	rt := &relationtuple.InternalRelationTuple{
		Relation:  r.Relation,
		Object:    r.Object,
		Namespace: r.NamespaceID,
	}

	if r.SubjectID.Valid {
		rt.Subject = &relationtuple.SubjectID{
			ID: r.SubjectID.UUID,
		}
	} else {
		rt.Subject = &relationtuple.SubjectSet{
			Namespace: r.SubjectSetNamespaceID.Int32,
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
		r.SubjectSetNamespaceID = sql.NullInt32{}
		r.SubjectSetObject = uuid.NullUUID{}
		r.SubjectSetRelation = sql.NullString{}
	case *relationtuple.SubjectSet:
		r.SubjectID = uuid.NullUUID{}
		r.SubjectSetNamespaceID = sql.NullInt32{
			Int32: st.Namespace,
			Valid: true,
		}
		r.SubjectSetObject = uuid.NullUUID{
			UUID:  st.Object,
			Valid: true,
		}
		r.SubjectSetRelation = sql.NullString{
			String: st.Relation,
			Valid:  true,
		}
	}
	return nil
}

func (r *RelationTuple) FromInternal(ctx context.Context, p *Persister, rt *relationtuple.InternalRelationTuple) error {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.FromInternal")
	defer span.End()

	r.NamespaceID = rt.Namespace
	r.Object = rt.Object
	r.Relation = rt.Relation

	return r.insertSubject(ctx, rt.Subject)
}

func (p *Persister) InsertRelationTuple(ctx context.Context, rel *relationtuple.InternalRelationTuple) error {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.InsertRelationTuple")
	defer span.End()

	if rel.Subject == nil {
		return errors.WithStack(relationtuple.ErrNilSubject)
	}

	p.d.Logger().Trace("creating tuples in database")

	rt := &RelationTuple{
		ID:         uuid.Must(uuid.NewV4()),
		CommitTime: time.Now(),
	}
	if err := rt.FromInternal(ctx, p, rel); err != nil {
		return err
	}

	if err := sqlcon.HandleError(
		p.CreateWithNetwork(ctx, rt),
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
			Where("subject_set_namespace_id IS NULL").
			Where("subject_set_object IS NULL").
			Where("subject_set_relation IS NULL")
	case *relationtuple.SubjectSet:
		q.
			Where("subject_set_namespace_id = ?", s.Namespace).
			Where("subject_set_object = ?", s.Object).
			Where("subject_set_relation = ?", s.Relation).
			// NULL checks to leverage partial indexes
			Where("subject_id IS NULL")
	case nil:
		return errors.WithStack(relationtuple.ErrNilSubject)
	}
	return nil
}

func (p *Persister) whereQuery(ctx context.Context, q *pop.Query, rq *relationtuple.RelationQuery) error {
	if rq.Namespace != nil {
		q.Where("namespace_id = ?", rq.Namespace)
	}
	if rq.Object != nil {
		q.Where("object = ?", rq.Object)
	}
	if rq.Relation != nil {
		q.Where("relation = ?", rq.Relation)
	}
	if s := rq.Subject(); s != nil {
		if err := p.whereSubject(ctx, q, s); err != nil {
			return err
		}
	}
	return nil
}

func (p *Persister) DeleteRelationTuples(ctx context.Context, rs ...*relationtuple.InternalRelationTuple) error {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.DeleteRelationTuples")
	defer span.End()

	return p.Transaction(ctx, func(ctx context.Context, _ *pop.Connection) error {
		for _, r := range rs {
			q := p.QueryWithNetwork(ctx).
				Where("namespace_id = ?", r.Namespace).
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

func (p *Persister) DeleteAllRelationTuples(ctx context.Context, query *relationtuple.RelationQuery) error {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.DeleteAllRelationTuples")
	defer span.End()

	return p.Transaction(ctx, func(ctx context.Context, _ *pop.Connection) error {
		sqlQuery := p.QueryWithNetwork(ctx)
		err := p.whereQuery(ctx, sqlQuery, query)
		if err != nil {
			return err
		}

		var res relationTuples
		return sqlQuery.Delete(&res)
	})
}

func (p *Persister) GetRelationTuples(ctx context.Context, query *relationtuple.RelationQuery, options ...x.PaginationOptionSetter) ([]*relationtuple.InternalRelationTuple, string, error) {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.GetRelationTuples")
	defer span.End()

	pagination, err := internalPaginationFromOptions(options...)
	if err != nil {
		return nil, "", err
	}

	sqlQuery := p.QueryWithNetwork(ctx).
		Order("nid, namespace_id, object, relation, subject_id, subject_set_namespace_id, subject_set_object, subject_set_relation, commit_time").
		Paginate(pagination.Page, pagination.PerPage)

	err = p.whereQuery(ctx, sqlQuery, query)
	if err != nil {
		return nil, "", err
	}
	var res relationTuples
	if err := sqlQuery.All(&res); err != nil {
		return nil, "", sqlcon.HandleError(err)
	}

	nextPageToken := pagination.encodeNextPageToken()
	if sqlQuery.Paginator.Page >= sqlQuery.Paginator.TotalPages {
		nextPageToken = ""
	}

	internalRes := make([]*relationtuple.InternalRelationTuple, len(res))
	for i, r := range res {
		var err error
		internalRes[i], err = r.toInternal()
		if err != nil {
			return nil, "", err
		}
	}

	return internalRes, nextPageToken, nil
}

func (p *Persister) WriteRelationTuples(ctx context.Context, rs ...*relationtuple.InternalRelationTuple) error {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.WriteRelationTuples")
	defer span.End()

	return p.Transaction(ctx, func(ctx context.Context, _ *pop.Connection) error {
		for _, r := range rs {
			if err := p.InsertRelationTuple(ctx, r); err != nil {
				return err
			}
		}
		return nil
	})
}

func (p *Persister) TransactRelationTuples(ctx context.Context, ins []*relationtuple.InternalRelationTuple, del []*relationtuple.InternalRelationTuple) error {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.TransactRelationTuples")
	defer span.End()

	return p.Transaction(ctx, func(ctx context.Context, _ *pop.Connection) error {
		if err := p.WriteRelationTuples(ctx, ins...); err != nil {
			return err
		}
		return p.DeleteRelationTuples(ctx, del...)
	})
}
