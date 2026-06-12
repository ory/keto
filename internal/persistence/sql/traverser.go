// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package sql

import (
	"context"
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/ory/x/otelx"
	"github.com/ory/x/sqlcon"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/attribute"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/ketoapi"
)

type (
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
		return "", nil, errors.WithStack(ketoapi.ErrNilSubject())
	}
	return sqlFragment, args, nil
}

// buildSubjectSetTypeSQL builds an optional SQL fragment and args to filter
// subject-set tuples by (namespace, relation) pairs. Returns empty string and
// nil args when allowedSubjectSets is nil (no filter applied).
func buildSubjectSetTypeSQL(allowedSubjectSets []relationtuple.SubjectSetType) (string, []any) {
	if len(allowedSubjectSets) == 0 {
		return "", nil
	}
	parts := make([]string, len(allowedSubjectSets))
	args := make([]any, 0, len(allowedSubjectSets)*2)
	for i, f := range allowedSubjectSets {
		parts[i] = "(current.subject_set_namespace = ? AND current.subject_set_relation = ?)"
		args = append(args, f.Namespace, f.Relation)
	}
	return "\nAND (" + strings.Join(parts, " OR ") + ")", args
}

// TraverseSubjectSetExpansion gets all subject sets for the object#relation and checks
// whether the requested subject is a member of each. When allowedSubjectSets is non-empty,
// only subject-set pointers matching the declared (namespace, relation) pairs are returned.
func (t *Persister) TraverseSubjectSetExpansion(ctx context.Context, start *relationtuple.RelationTuple, allowedSubjectSets []relationtuple.SubjectSetType) (res []*relationtuple.TraversalResult, err error) {
	ctx, span := t.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.TraverseSubjectSetExpansion")
	defer otelx.End(span, &err)

	targetSubjectSQL, targetSubjectArgs, err := whereSubject(start.Subject)
	if err != nil {
		return nil, err
	}

	var rowCount int
	defer func() { span.SetAttributes(attribute.Int("tuples_loaded", rowCount)) }()

	subjectSetFilterSQL, subjectSetFilterArgs := buildSubjectSetTypeSQL(allowedSubjectSets)

	shardID := uuid.Nil
	for {
		var (
			rows  []*subjectExpandedRelationTupleRow
			limit = 1000
		)
		queryArgs := make([]any, 0)
		queryArgs = append(queryArgs, targetSubjectArgs...)
		queryArgs = append(queryArgs, t.NetworkID(ctx), shardID, start.Namespace, start.Object, start.Relation)
		queryArgs = append(queryArgs, subjectSetFilterArgs...)
		queryArgs = append(queryArgs, limit)

		err = t.Connection(ctx).RawQuery(fmt.Sprintf(`
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
	  %s
ORDER BY current.shard_id
LIMIT ?
`, targetSubjectSQL, subjectSetFilterSQL),
			queryArgs...,
		).All(&rows)
		if err != nil {
			return nil, sqlcon.HandleError(err)
		}
		rowCount += len(rows)

		for _, r := range rows {
			to := r.ToInternal()
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

func (t *Persister) FindTupleWithRelations(ctx context.Context, tuple *relationtuple.RelationTuple, relations []string) (_ *relationtuple.RelationTuple, err error) {
	ctx, span := t.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.FindTupleWithRelations")
	defer otelx.End(span, &err)

	if len(relations) == 0 {
		return nil, nil
	}
	var rows relationTuples

	query := t.queryWithNetwork(ctx)
	if err := t.whereQuery(ctx, query, &relationtuple.RelationQuery{
		Namespace: &tuple.Namespace,
		Object:    &tuple.Object,
		Subject:   tuple.Subject,
	}); err != nil {
		return nil, err
	}
	err = query.Where("relation IN (?)", relations).Limit(1).All(&rows)
	if err != nil {
		return nil, sqlcon.HandleError(err)
	}

	// If we got any rows back, success!
	if len(rows) > 0 {
		return rows[0].ToInternal(), nil
	}

	return nil, nil
}
