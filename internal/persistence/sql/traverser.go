// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package sql

import (
	"context"
	"fmt"
	"strings"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/ory/x/otelx"
	"github.com/ory/x/sqlcon"
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/namespace"
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

	rewriteRelationTupleRow struct {
		RelationTuple
		Traversal relationtuple.Traversal `db:"traversal"`
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
                 %s  -- subject where clause
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
	if err != nil {
		return nil, sqlcon.HandleError(err)
	}

	res = make([]*relationtuple.TraversalResult, len(rows))

	for i, r := range rows {
		to, err := r.RelationTuple.ToInternal()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		to.Subject = start.Subject
		res[i] = &relationtuple.TraversalResult{
			From:  start,
			To:    to,
			Via:   relationtuple.TraversalSubjectSetExpand,
			Found: r.Found,
		}
	}

	return res, nil
}

func (t *Traverser) TraverseSubjectSetRewrite(ctx context.Context, start *relationtuple.RelationTuple, computedSubjectSets []string) (res []*relationtuple.TraversalResult, err error) {

	namespaceManager, err := t.d.Config(ctx).NamespaceManager()
	if err != nil {
		return nil, err
	}

	subjectSQL, sqlArgs, err := whereSubject(start.Subject)

	var targetRelationPlaceholders []string
	for _, relation := range computedSubjectSets {
		astRel, _ := namespace.ASTRelationFor(ctx, namespaceManager, start.Namespace, relation)
		if astRel != nil && astRel.SubjectSetRewrite != nil {
			continue
		}
		sqlArgs = append(sqlArgs, relation)
		targetRelationPlaceholders = append(targetRelationPlaceholders, "?")
	}
	sqlArgs = append([]any{t.p.NetworkID(ctx), start.Namespace, start.Object}, sqlArgs...)

	if len(targetRelationPlaceholders) > 0 {
		_, span := t.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.TraverseSubjectSetRewrite")
		defer otelx.End(span, &err)

		var rows []*rewriteRelationTupleRow
		err = t.conn.RawQuery(fmt.Sprintf(`
SELECT t.*,
       'computed userset' as traversal
FROM keto_relation_tuples AS t
WHERE t.nid = ? AND
      t.namespace = ? AND
      t.object = ? AND
      %s AND -- subject where clause
      t.relation IN (%s)
LIMIT 1;
`, subjectSQL, strings.Join(targetRelationPlaceholders, ", ")),
			sqlArgs...,
		).All(&rows)
		if err != nil {
			return nil, sqlcon.HandleError(err)
		}

		// If we got any rows back, success!
		if len(rows) > 0 {
			r := rows[0]
			to, err := r.RelationTuple.ToInternal()
			if err != nil {
				return nil, errors.WithStack(err)
			}
			return []*relationtuple.TraversalResult{{
				From:  start,
				To:    to,
				Via:   r.Traversal,
				Found: true,
			}}, nil
		}
	}

	// Otherwise, the next candidates are those tuples with relations from the rewrite
	for _, relation := range computedSubjectSets {
		res = append(res, &relationtuple.TraversalResult{
			From: start,
			To: &relationtuple.RelationTuple{
				Namespace: start.Namespace,
				Object:    start.Object,
				Relation:  relation,
				Subject:   start.Subject,
			},
			Via:   relationtuple.TraversalComputedUserset,
			Found: false,
		})
	}

	return res, nil
}

func NewTraverser(p *Persister) *Traverser {
	return &Traverser{
		conn: p.conn,
		d:    p.d,
		nid:  p.nid,
		p:    p,
	}
}
