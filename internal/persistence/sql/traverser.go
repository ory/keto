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
	args := make([]any, 0, len(allowedSubjectSets)*2)
	for _, f := range allowedSubjectSets {
		args = append(args, f.Namespace, f.Relation)
	}
	return "\nAND (current.subject_set_namespace, current.subject_set_relation) IN (" + tuplePlaceholders(len(allowedSubjectSets)) + ")", args
}

// tuplePlaceholders returns n comma-separated "(?, ?)" pairs for use in a
// row-value IN list.
func tuplePlaceholders(n int) string {
	return strings.TrimSuffix(strings.Repeat("(?, ?), ", n), ", ")
}

// buildFoundExpr builds the SQL expression for the found column.
//
// When allowedSubjectSets is empty, found is the result of the EXISTS subquery for every row.
// When non-empty, AllowsDirect on each entry controls whether found can be true for that type:
//   - no entries have AllowsDirect: found is always false (literal, no subquery).
//   - all entries have AllowsDirect: found is the EXISTS result for every row.
//   - mixed: found can be true only for rows whose type is in the AllowsDirect set
func buildFoundExpr(allowedSubjectSets []relationtuple.SubjectSetType, subject relationtuple.Subject) (string, []any, error) {
	var directTypes []relationtuple.SubjectSetType
	if len(allowedSubjectSets) > 0 {
		for _, t := range allowedSubjectSets {
			if t.AllowsDirect {
				directTypes = append(directTypes, t)
			}
		}
		if len(directTypes) == 0 {
			return "false", nil, nil
		}
	}

	subjectSQL, subjectArgs, err := whereSubject(subject)
	if err != nil {
		return "", nil, err
	}

	existsQuery := fmt.Sprintf(`EXISTS(
           SELECT 1 FROM keto_relation_tuples
           WHERE nid = current.nid AND
                 namespace = current.subject_set_namespace AND
                 object = current.subject_set_object AND
                 relation = current.subject_set_relation AND
                 %s
       )`, subjectSQL)

	// No filter, or all types allow direct membership: EXISTS runs for every matched row.
	if len(allowedSubjectSets) == 0 || len(directTypes) == len(allowedSubjectSets) {
		return existsQuery, subjectArgs, nil
	}

	// Only some types allow direct membership: gate the EXISTS result to those
	// types. The EXISTS subquery may still be evaluated for every returned row.
	args := make([]any, 0, len(subjectArgs)+len(directTypes)*2)
	args = append(args, subjectArgs...)
	for _, t := range directTypes {
		args = append(args, t.Namespace, t.Relation)
	}
	return fmt.Sprintf("(%s AND (current.subject_set_namespace, current.subject_set_relation) IN (%s))", existsQuery, tuplePlaceholders(len(directTypes))), args, nil
}

// TraverseSubjectSetExpansion gets all subject-set pointers for the object#relation and
// checks whether the requested subject is a direct member of each. When allowedSubjectSets
// is non-empty, only pointers matching the declared (namespace, relation) pairs are returned,
// and AllowsDirect on each type controls whether a direct membership check(subquery) is performed.
func (t *Persister) TraverseSubjectSetExpansion(ctx context.Context, start *relationtuple.RelationTuple, allowedSubjectSets []relationtuple.SubjectSetType) (res []*relationtuple.TraversalResult, err error) {
	ctx, span := t.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.TraverseSubjectSetExpansion")
	defer otelx.End(span, &err)

	var rowCount int
	defer func() { span.SetAttributes(attribute.Int("tuples_loaded", rowCount)) }()

	foundExpr, foundArgs, err := buildFoundExpr(allowedSubjectSets, start.Subject)
	if err != nil {
		return nil, err
	}

	subjectSetFilterSQL, subjectSetFilterArgs := buildSubjectSetTypeSQL(allowedSubjectSets)
	shardID := uuid.Nil
	for {
		var (
			rows  []*subjectExpandedRelationTupleRow
			limit = 1000
		)
		queryArgs := make([]any, 0)
		queryArgs = append(queryArgs, foundArgs...)
		queryArgs = append(queryArgs, t.NetworkID(ctx), shardID, start.Namespace, start.Object, start.Relation)
		queryArgs = append(queryArgs, subjectSetFilterArgs...)
		queryArgs = append(queryArgs, limit)

		err = t.Connection(ctx).RawQuery(fmt.Sprintf(`
SELECT current.shard_id AS shard_id,
       current.subject_set_namespace AS namespace,
       current.subject_set_object AS object,
       current.subject_set_relation AS relation,
       %s AS found
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
`, foundExpr, subjectSetFilterSQL),
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
