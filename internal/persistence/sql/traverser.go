package sql

import (
	"context"
	"fmt"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/ory/x/otelx"
	"github.com/ory/x/sqlcon"
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/ketoapi"
)

type (
	Traverser struct {
		conn *pop.Connection
		d    dependencies
		nid  uuid.UUID
		p    *Persister
	}

	subjectExpandedRelationTupleRow struct {
		RelationTuple

		Found bool `db:"found"`
	}
)

func whereSubject(sub relationtuple.Subject) (sqlFragment string, args []any, err error) {
	switch s := sub.(type) {
	case *relationtuple.SubjectID:
		sqlFragment = "subject_id = ? AND subject_set_namespace IS NULL AND subject_set_object IS NULL AND subject_set_relation IS NULL"
		args = []any{s.ID}

	case *relationtuple.SubjectSet:
		sqlFragment = "subject_id IS NULL AND subject_set_namespace = ? AND subject_set_object = ? AND subject_set_relation = ?"
		args = []any{s.Namespace, s.Object, s.Relation}

	case nil:
		return "", nil, errors.WithStack(ketoapi.ErrNilSubject)
	}
	return sqlFragment, args, nil
}

func (t *Traverser) TraverseSubjectSetExpansion(ctx context.Context, start *relationtuple.RelationTuple) (res []*relationtuple.TraversalResult, err error) {
	ctx, span := t.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.TraverseSubjectSetExpansion")
	defer otelx.End(span, &err)

	targetSubjectSQL, targetSubjectArgs, err := whereSubject(start.Subject)
	if err != nil {
		return nil, err
	}

	var rows []*subjectExpandedRelationTupleRow
	err = t.conn.RawQuery(fmt.Sprintf(`
SELECT current.subject_set_namespace AS namespace,
       current.subject_set_object AS object,
       current.subject_set_relation AS relation,
       EXISTS(
           SELECT 1 FROM keto_relation_tuples
           WHERE nid = current.nid AND
                 namespace = current.subject_set_namespace AND
                 object = current.subject_set_object AND
                 relation = current.subject_set_relation AND
                 %s
       ) AS found
FROM keto_relation_tuples AS current
WHERE current.nid = ? AND
      current.namespace = ? AND
      current.object = ? AND
      current.relation = ? AND
      current.subject_id IS NULL
`, targetSubjectSQL),
		append(targetSubjectArgs, t.p.NetworkID(ctx), start.Namespace, start.Object, start.Relation)...,
	).All(&rows)
	res = make([]*relationtuple.TraversalResult, len(rows))

	for i, r := range rows {
		to, err := r.RelationTuple.ToInternal()
		to.Subject = start.Subject
		if err != nil {
			return nil, errors.WithStack(err)
		}
		res[i] = &relationtuple.TraversalResult{
			From:  start,
			To:    to,
			Via:   relationtuple.TraversalSubjectSetExpand,
			Found: r.Found,
		}
	}

	return res, sqlcon.HandleError(err)
}

func NewTraverser(p *Persister) *Traverser {
	return &Traverser{
		conn: p.conn,
		d:    p.d,
		nid:  p.nid,
		p:    p,
	}
}
