package sql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gofrs/uuid"

	"github.com/ory/x/sqlcon"

	"github.com/gobuffalo/pop/v5"

	"github.com/pkg/errors"

	"github.com/ory/keto/internal/namespace"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

type (
	relationTuple struct {
		// An ID field is required to make pop happy. The actual ID is a composite primary key.
		ID                    uuid.UUID      `db:"shard_id"`
		NetworkID             uuid.UUID      `db:"network_id"`
		NamespaceID           int64          `db:"namespace_id"`
		Object                string         `db:"object"`
		Relation              string         `db:"relation"`
		SubjectID             sql.NullString `db:"subject_id"`
		SubjectSetNamespaceID sql.NullInt64  `db:"subject_set_namespace_id"`
		SubjectSetObject      sql.NullString `db:"subject_set_object"`
		SubjectSetRelation    sql.NullString `db:"subject_set_relation"`
		CommitTime            time.Time      `db:"commit_time"`
	}
	relationTuples []*relationTuple
	whereStmts     struct {
		stmt string
		arg  interface{}
	}
)

func (relationTuples) TableName(_ context.Context) string {
	return "keto_relation_tuples"
}

func (relationTuple) TableName(_ context.Context) string {
	return "keto_relation_tuples"
}

func (r *relationTuple) toInternal(ctx context.Context, nm namespace.Manager) (*relationtuple.InternalRelationTuple, error) {
	if r == nil {
		return nil, nil
	}

	n, err := nm.GetNamespaceByID(ctx, r.NamespaceID)
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
		sn, err := nm.GetNamespaceByID(ctx, r.SubjectSetNamespaceID.Int64)
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

func (r *relationTuple) insertSubject(ctx context.Context, nm namespace.Manager, s relationtuple.Subject) error {
	switch st := s.(type) {
	case *relationtuple.SubjectID:
		r.SubjectID = sql.NullString{
			String: st.ID,
			Valid:  true,
		}
		r.SubjectSetNamespaceID = sql.NullInt64{}
		r.SubjectSetObject = sql.NullString{}
		r.SubjectSetRelation = sql.NullString{}
	case *relationtuple.SubjectSet:
		n, err := nm.GetNamespaceByName(ctx, st.Namespace)
		if err != nil {
			return err
		}

		r.SubjectID = sql.NullString{}
		r.SubjectSetNamespaceID = sql.NullInt64{
			Int64: n.ID,
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

func (r *relationTuple) fromInternal(ctx context.Context, p *Persister, rt *relationtuple.InternalRelationTuple) error {
	nID, err := p.NetworkID(ctx)
	if err != nil {
		return err
	}

	nm, err := p.d.Config().NamespaceManager()
	if err != nil {
		return err
	}

	n, err := nm.GetNamespaceByName(ctx, rt.Namespace)
	if err != nil {
		return err
	}

	r.NetworkID = nID
	r.NamespaceID = n.ID
	r.Object = rt.Object
	r.Relation = rt.Relation

	return r.insertSubject(ctx, nm, rt.Subject)
}

func (p *Persister) insertRelationTuple(ctx context.Context, rel *relationtuple.InternalRelationTuple) error {
	if rel.Subject == nil {
		return errors.New("subject is not allowed to be nil")
	}

	p.d.Logger().WithFields(rel.ToLoggerFields()).Trace("creating in database")

	nID, err := p.NetworkID(ctx)
	if err != nil {
		return err
	}

	rt := &relationTuple{
		ID:         uuid.Must(uuid.NewV4()),
		NetworkID:  nID,
		CommitTime: time.Now(),
	}
	if err := rt.fromInternal(ctx, p, rel); err != nil {
		return err
	}

	if err := sqlcon.HandleError(
		p.Connection(ctx).Create(rt),
	); err != nil {
		fmt.Print("\n\n")
		return err
	}
	return nil
}

func (p *Persister) deleteRelationTupleSubjectID(ctx context.Context, r *relationtuple.InternalRelationTuple, s *relationtuple.SubjectID, n *namespace.Namespace, network uuid.UUID) error {
	if err := p.Connection(ctx).RawQuery(
		"DELETE FROM keto_relation_tuples WHERE network_id = ? AND namespace_id = ? AND object = ? AND relation = ? AND subject_id = ?",
		network,
		n.ID,
		r.Object,
		r.Relation,
		s.ID,
	).Exec(); err != nil {
		return sqlcon.HandleError(err)
	}

	return nil
}

func (p *Persister) deleteRelationTupleSubjectSet(ctx context.Context, r *relationtuple.InternalRelationTuple, s *relationtuple.SubjectSet, n *namespace.Namespace, network uuid.UUID) error {
	nm, err := p.d.Config().NamespaceManager()
	if err != nil {
		return err
	}

	sn, err := nm.GetNamespaceByName(ctx, s.Namespace)
	if err != nil {
		return err
	}

	if err := p.Connection(ctx).RawQuery(
		"DELETE FROM keto_relation_tuples WHERE network_id = ? AND namespace_id = ? AND object = ? AND relation = ? AND subject_set_namespace_id = ? AND subject_set_object = ? AND subject_set_relation = ?",
		network,
		n.ID,
		r.Object,
		r.Relation,
		sn.ID,
		s.Object,
		s.Relation,
	).Exec(); err != nil {
		return sqlcon.HandleError(err)
	}
	return nil
}

func (p *Persister) DeleteRelationTuples(ctx context.Context, rs ...*relationtuple.InternalRelationTuple) error {
	network, err := p.NetworkID(ctx)
	if err != nil {
		return err
	}

	nm, err := p.d.Config().NamespaceManager()
	if err != nil {
		return err
	}

	return p.transaction(ctx, func(ctx context.Context, c *pop.Connection) error {
		for _, r := range rs {
			n, err := nm.GetNamespaceByName(ctx, r.Namespace)
			if err != nil {
				return err
			}

			switch s := r.Subject.(type) {
			case *relationtuple.SubjectID:
				if err := p.deleteRelationTupleSubjectID(ctx, r, s, n, network); err != nil {
					return err
				}
			case *relationtuple.SubjectSet:
				if err := p.deleteRelationTupleSubjectSet(ctx, r, s, n, network); err != nil {
					return err
				}
			default:
				return errors.New("subject is not allowed to be nil")
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

	var wheres []whereStmts
	if query.Relation != "" {
		wheres = append(wheres, whereStmts{stmt: "relation = ?", arg: query.Relation})
	}
	if query.Object != "" {
		wheres = append(wheres, whereStmts{stmt: "object = ?", arg: query.Object})
	}
	if query.Subject != nil {
		switch s := query.Subject.(type) {
		case *relationtuple.SubjectID:
			wheres = append(wheres, whereStmts{stmt: "subject_id = ?", arg: s.ID})
		case *relationtuple.SubjectSet:
			n, err := nm.GetNamespaceByName(ctx, s.Namespace)
			if err != nil {
				return nil, "", err
			}

			wheres = append(
				wheres,
				whereStmts{stmt: "subject_set_namespace_id = ?", arg: n.ID},
				whereStmts{stmt: "subject_set_object = ?", arg: s.Object},
				whereStmts{stmt: "subject_set_relation = ?", arg: s.Relation},
			)
		}
	}
	if query.Namespace != "" {
		n, err := nm.GetNamespaceByName(ctx, query.Namespace)
		if err != nil {
			return nil, "", err
		}
		wheres = append(wheres, whereStmts{stmt: "namespace_id = ?", arg: n.ID})
	}

	nID, err := p.NetworkID(ctx)
	if err != nil {
		return nil, "", err
	}
	sqlQuery := p.Connection(ctx).
		Where("network_id = ?", nID).
		Order("network_id, namespace_id, object, relation, subject_id, subject_set_namespace_id, subject_set_object, subject_set_relation").
		Paginate(pagination.Page, pagination.PerPage)

	for _, w := range wheres {
		sqlQuery = sqlQuery.Where(w.stmt, w.arg)
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
		internalRes[i], err = r.toInternal(ctx, nm)
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
