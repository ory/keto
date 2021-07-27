package sql

import (
	"context"
	"database/sql"
	"time"

	"github.com/gobuffalo/pop/v5"
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
		NamespaceID           uuid.UUID      `db:"namespace_id"`
		Object                string         `db:"object"`
		Relation              string         `db:"relation"`
		SubjectID             sql.NullString `db:"subject_id"`
		SubjectSetNamespaceID uuid.NullUUID  `db:"subject_set_namespace_id"`
		SubjectSetObject      sql.NullString `db:"subject_set_object"`
		SubjectSetRelation    sql.NullString `db:"subject_set_relation"`
		CommitTime            time.Time      `db:"commit_time"`
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

	n, err := p.GetNamespaceConfigID(ctx, r.NamespaceID)
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
		n, err := p.GetNamespaceConfigID(ctx, r.SubjectSetNamespaceID.UUID)
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
		r.SubjectSetNamespaceID = uuid.NullUUID{}
		r.SubjectSetObject = sql.NullString{}
		r.SubjectSetRelation = sql.NullString{}
	case *relationtuple.SubjectSet:
		nID, err := p.GetNamespaceID(ctx, st.Namespace)
		if err != nil {
			return err
		}

		r.SubjectID = sql.NullString{}
		r.SubjectSetNamespaceID = uuid.NullUUID{
			UUID:  nID,
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

func (r *RelationTuple) fromInternal(ctx context.Context, p *Persister, rt *relationtuple.InternalRelationTuple) error {
	nID, err := p.GetNamespaceID(ctx, rt.Namespace)
	if err != nil {
		return err
	}

	r.NamespaceID = nID
	r.Object = rt.Object
	r.Relation = rt.Relation

	return r.insertSubject(ctx, p, rt.Subject)
}

func (p *Persister) insertRelationTuple(ctx context.Context, rel *relationtuple.InternalRelationTuple) error {
	if rel.Subject == nil {
		return errors.New("subject is not allowed to be nil")
	}

	p.d.Logger().WithFields(rel.ToLoggerFields()).Trace("creating in database")

	rt := &RelationTuple{
		ID:         uuid.Must(uuid.NewV4()),
		CommitTime: time.Now(),
	}
	if err := rt.fromInternal(ctx, p, rel); err != nil {
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
		nID, err := p.GetNamespaceID(ctx, s.Namespace)
		if err != nil {
			return err
		}

		q.
			Where("subject_set_namespace_id = ?", nID).
			Where("subject_set_object = ?", s.Object).
			Where("subject_set_relation = ?", s.Relation).
			// NULL checks to leverage partial indexes
			Where("subject_id IS NULL")
	case nil:
		return errors.New("subject is not allowed to be nil")
	}
	return nil
}

func (p *Persister) DeleteRelationTuples(ctx context.Context, rs ...*relationtuple.InternalRelationTuple) error {
	return p.transaction(ctx, func(ctx context.Context, c *pop.Connection) error {
		for _, r := range rs {
			nID, err := p.GetNamespaceID(ctx, r.Namespace)
			if err != nil {
				return err
			}

			q := p.QueryWithNetwork(ctx).
				Where("namespace_id = ?", nID).
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

func (p *Persister) GetRelationTuples(ctx context.Context, query *relationtuple.RelationQuery, options ...x.PaginationOptionSetter) ([]*relationtuple.InternalRelationTuple, string, error) {
	pagination, err := internalPaginationFromOptions(options...)
	if err != nil {
		return nil, "", err
	}

	nm, err := p.d.Config().NamespaceManager()
	if err != nil {
		return nil, "", err
	}

	sqlQuery := p.QueryWithNetwork(ctx).
		Order("nid, namespace_id, object, relation, subject_id, subject_set_namespace_id, subject_set_object, subject_set_relation, commit_time").
		Paginate(pagination.Page, pagination.PerPage)

	if query.Relation != "" {
		sqlQuery.Where("relation = ?", query.Relation)
	}
	if query.Object != "" {
		sqlQuery.Where("object = ?", query.Object)
	}
	if query.Subject != nil {
		if err := p.whereSubject(ctx, sqlQuery, query.Subject); err != nil {
			return nil, "", err
		}
	}

	if query.Namespace != "" {
		nID, err := p.GetNamespaceID(ctx, query.Namespace)
		if err != nil {
			return nil, "", err
		}
		sqlQuery.Where("namespace_id = ?", nID)
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
	return p.transaction(ctx, func(ctx context.Context, _ *pop.Connection) error {
		for _, r := range rs {
			if err := p.insertRelationTuple(ctx, r); err != nil {
				return err
			}
		}
		return nil
	})
}

func (p *Persister) TransactRelationTuples(ctx context.Context, ins []*relationtuple.InternalRelationTuple, del []*relationtuple.InternalRelationTuple) error {
	return p.transaction(ctx, func(ctx context.Context, _ *pop.Connection) error {
		if err := p.WriteRelationTuples(ctx, ins...); err != nil {
			return err
		}
		return p.DeleteRelationTuples(ctx, del...)
	})
}
