package sql

import (
	"context"
	"database/sql"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/ory/x/sqlcon"
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

type (
	RelationTuple struct {
		// An ID field is required to make pop happy. The actual ID is a composite primary key.
		ID                    uuid.UUID      `db:"shard_id"`
		NetworkID             uuid.UUID      `db:"nid"`
		NamespaceID           int32          `db:"namespace_id"`
		Object                string         `db:"object"`
		Relation              string         `db:"relation"`
		SubjectID             sql.NullString `db:"subject_id"`
		SubjectSetNamespaceID sql.NullInt32  `db:"subject_set_namespace_id"`
		SubjectSetObject      sql.NullString `db:"subject_set_object"`
		SubjectSetRelation    sql.NullString `db:"subject_set_relation"`
		CommitTime            time.Time      `db:"commit_time"`
		HlcTimestamp          string         `db:"hlc_timestamp"`
	}
	relationTuples []*RelationTuple
)

func (relationTuples) TableName(_ context.Context) string {
	return "keto_relation_tuples"
}

func (RelationTuple) TableName(_ context.Context) string {
	return "keto_relation_tuples"
}

func (r *RelationTuple) toInternal(ctx context.Context, nm namespace.Manager, p *Persister) (*relationtuple.InternalRelationTuple, error) {
	if r == nil {
		return nil, nil
	}

	n, err := p.GetNamespaceByID(ctx, r.NamespaceID)
	if err != nil {
		return nil, err
	}

	rt := &relationtuple.InternalRelationTuple{
		Relation:  r.Relation,
		Object:    r.Object,
		Namespace: n.Name,
	}

	if r.SubjectID.Valid {
		rt.Subject = &relationtuple.SubjectID{
			ID: r.SubjectID.String,
		}
	} else {
		n, err := p.GetNamespaceByID(ctx, r.SubjectSetNamespaceID.Int32)
		if err != nil {
			return nil, err
		}
		sn, err := nm.GetNamespaceByConfigID(ctx, n.ID)
		if err != nil {
			return nil, err
		}
		rt.Subject = &relationtuple.SubjectSet{
			Namespace: sn.Name,
			Object:    r.SubjectSetObject.String,
			Relation:  r.SubjectSetRelation.String,
		}
	}

	return rt, nil
}

func (r *RelationTuple) insertSubject(ctx context.Context, p *Persister, s relationtuple.Subject) error {
	switch st := s.(type) {
	case *relationtuple.SubjectID:
		r.SubjectID = sql.NullString{
			String: st.ID,
			Valid:  true,
		}
		r.SubjectSetNamespaceID = sql.NullInt32{}
		r.SubjectSetObject = sql.NullString{}
		r.SubjectSetRelation = sql.NullString{}
	case *relationtuple.SubjectSet:
		n, err := p.GetNamespaceByName(ctx, st.Namespace)
		if err != nil {
			return err
		}

		r.SubjectID = sql.NullString{}
		r.SubjectSetNamespaceID = sql.NullInt32{
			Int32: n.ID,
			Valid: true,
		}
		r.SubjectSetObject = sql.NullString{
			String: st.Object,
			Valid:  true,
		}
		r.SubjectSetRelation = sql.NullString{
			String: st.Relation,
			Valid:  true,
		}
	}
	return nil
}

func (r *RelationTuple) FromInternal(ctx context.Context, p *Persister, rt *relationtuple.InternalRelationTuple) error {
	n, err := p.GetNamespaceByName(ctx, rt.Namespace)
	if err != nil {
		return err
	}

	r.NamespaceID = n.ID
	r.Object = rt.Object
	r.Relation = rt.Relation

	return r.insertSubject(ctx, p, rt.Subject)
}

func (p *Persister) InsertRelationTuple(ctx context.Context, rel *relationtuple.InternalRelationTuple) error {
	if rel.Subject == nil {
		return errors.WithStack(relationtuple.ErrNilSubject)
	}

	p.d.Logger().WithFields(rel.ToLoggerFields()).Trace("creating in database")

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

func (p *Persister) whereSubject(ctx context.Context, q *pop.Query, sub relationtuple.Subject) error {
	switch s := sub.(type) {
	case *relationtuple.SubjectID:
		q.
			Where("subject_id = ?", s.ID).
			// NULL checks to leverage partial indexes
			Where("subject_set_namespace_id IS NULL").
			Where("subject_set_object IS NULL").
			Where("subject_set_relation IS NULL")
	case *relationtuple.SubjectSet:
		n, err := p.GetNamespaceByName(ctx, s.Namespace)
		if err != nil {
			return err
		}

		q.
			Where("subject_set_namespace_id = ?", n.ID).
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
	if rq.Namespace != "" {
		n, err := p.GetNamespaceByName(ctx, rq.Namespace)
		if err != nil {
			return err
		}
		q.Where("namespace_id = ?", n.ID)
	}
	if rq.Object != "" {
		q.Where("object = ?", rq.Object)
	}
	if rq.Relation != "" {
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
	return p.Transaction(ctx, func(ctx context.Context, c *pop.Connection) error {
		for _, r := range rs {
			n, err := p.GetNamespaceByName(ctx, r.Namespace)
			if err != nil {
				return err
			}

			q := p.QueryWithNetwork(ctx).
				Where("namespace_id = ?", n.ID).
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
	return p.Transaction(ctx, func(ctx context.Context, c *pop.Connection) error {
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
	pagination, err := internalPaginationFromOptions(options...)
	if err != nil {
		return nil, "", err
	}

	nm, err := p.d.Config(ctx).NamespaceManager()
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
		internalRes[i], err = r.toInternal(ctx, nm, p)
		if err != nil {
			return nil, "", err
		}
	}

	return internalRes, nextPageToken, nil
}

func (p *Persister) WriteRelationTuples(ctx context.Context, rs ...*relationtuple.InternalRelationTuple) error {
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
	return p.Transaction(ctx, func(ctx context.Context, _ *pop.Connection) error {
		if err := p.WriteRelationTuples(ctx, ins...); err != nil {
			return err
		}
		return p.DeleteRelationTuples(ctx, del...)
	})
}
