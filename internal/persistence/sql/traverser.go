// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package sql

import (
	"context"
	"fmt"

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

// TraverseSubjectSetExpansion gets all subject sets for the object#relation.
// It also checks whether the requested subject is a member of each of the returned subject sets.
func (t *Traverser) TraverseSubjectSetExpansion(ctx context.Context, start *relationtuple.RelationTuple) (res []*relationtuple.TraversalResult, err error) {
	ctx, span := t.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.TraverseSubjectSetExpansion")
	defer otelx.End(span, &err)

	targetSubjectSQL, targetSubjectArgs, err := whereSubject(start.Subject)
	if err != nil {
		return nil, err
	}

	shardID := uuid.Nil
	for {
		var (
			rows  []*subjectExpandedRelationTupleRow
			limit = 1000
		)
		err = t.conn.WithContext(ctx).RawQuery(fmt.Sprintf(`
SELECT current.shard_id AS shard_id,
       current.subject_set_namespace AS namespace,
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
      current.shard_id > ? AND
      current.namespace = ? AND
      current.object = ? AND
      current.relation = ? AND
      current.subject_id IS NULL
ORDER BY current.nid, current.shard_id
LIMIT ?
`, targetSubjectSQL),
			append(targetSubjectArgs, t.p.NetworkID(ctx), shardID, start.Namespace, start.Object, start.Relation, limit)...,
		).All(&rows)
		if err != nil {
			return nil, sqlcon.HandleError(err)
		}

		for _, r := range rows {
			to, err := r.RelationTuple.ToInternal()
			if err != nil {
				return nil, errors.WithStack(err)
			}
			to.Subject = start.Subject
			res = append(res, &relationtuple.TraversalResult{
				From:  start,
				To:    to,
				Via:   relationtuple.TraversalSubjectSetExpand,
				Found: r.Found,
			})
			if r.Found {
				return res, nil
			}
		}
		if len(rows) == limit {
			shardID = rows[limit-1].ID
		} else {
			break
		}
	}

	return res, nil
}

func (t *Traverser) TraverseSubjectSetRewrite(ctx context.Context, start *relationtuple.RelationTuple, computedSubjectSets []string) (res []*relationtuple.TraversalResult, err error) {
	ctx, span := t.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.TraverseSubjectSetRewrite")
	defer otelx.End(span, &err)

	namespaceManager, err := t.d.Config(ctx).NamespaceManager()
	if err != nil {
		return nil, err
	}

	var relations []string
	for _, relation := range computedSubjectSets {
		astRel, _ := namespace.ASTRelationFor(ctx, namespaceManager, start.Namespace, relation)
		// In strict mode, we can skip querying for those relations that have userset rewrites defined,
		// because we can already apply those rewrites in memory.
		if t.d.Config(ctx).StrictMode() && astRel != nil && astRel.SubjectSetRewrite != nil {
			continue
		}
		relations = append(relations, relation)
	}

	if len(relations) > 0 {
		var rows relationTuples

		query := t.p.queryWithNetwork(ctx)
		if err := t.p.whereQuery(ctx, query, &relationtuple.RelationQuery{
			Namespace: &start.Namespace,
			Object:    &start.Object,
			Subject:   start.Subject,
		}); err != nil {
			return nil, err
		}
		err = query.Where("relation IN (?)", relations).Limit(1).All(&rows)
		if err != nil {
			return nil, sqlcon.HandleError(err)
		}

		// If we got any rows back, success!
		if len(rows) > 0 {
			r := rows[0]
			to, err := r.ToInternal()
			if err != nil {
				return nil, errors.WithStack(err)
			}
			return []*relationtuple.TraversalResult{{
				From:  start,
				To:    to,
				Via:   relationtuple.TraversalComputedUserset,
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
