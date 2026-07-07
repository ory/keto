// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package sql

import (
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/relationtuple"
)

// existsFor returns the EXISTS subquery text that buildFoundExpr produces for the given subject.
func existsFor(t *testing.T, subject relationtuple.Subject) string {
	t.Helper()
	subjectSQL, _, err := whereSubject(subject)
	require.NoError(t, err)
	return `EXISTS(
           SELECT 1 FROM keto_relation_tuples
           WHERE nid = current.nid AND
                 namespace = current.subject_set_namespace AND
                 object = current.subject_set_object AND
                 relation = current.subject_set_relation AND
                 ` + subjectSQL + `
       )`
}

func TestBuildFoundExpr(t *testing.T) {
	t.Parallel()
	t.Run("case=subject as SubjectID", func(t *testing.T) {
		subject := &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())}
		existsExpr := existsFor(t, subject)

		// when SubjectID is used, no `allowedSubjectSets` is expected
		t.Run("no allowed types: EXISTS runs for all rows", func(t *testing.T) {
			expr, args, err := buildFoundExpr(nil, subject)
			require.NoError(t, err)

			assert.Equal(t, existsExpr, expr)
			assert.Equal(t, []any{subject.ID}, args)
		})
	})

	t.Run("case=subject as SubjectSet", func(t *testing.T) {
		subject := &relationtuple.SubjectSet{
			Namespace: "User",
			Object:    uuid.Must(uuid.NewV4()),
		}
		existsExpr := existsFor(t, subject)

		t.Run("no allowed types: EXISTS runs for all rows", func(t *testing.T) {
			expr, args, err := buildFoundExpr(nil, subject)
			require.NoError(t, err)

			assert.Equal(t, existsExpr, expr)
			assert.Equal(t, []any{subject.Namespace, subject.Object, subject.Relation}, args)
		})

		// All types have AllowsDirect=false: OPL forbids direct membership for every type,
		// so the found column is always false without touching the inner table.
		t.Run("all AllowsDirect=false: found is always false", func(t *testing.T) {
			types := []relationtuple.SubjectSetType{
				{Namespace: "Group", Relation: "member", AllowsDirect: false},
				{Namespace: "Org", Relation: "admin", AllowsDirect: false},
			}
			expr, args, err := buildFoundExpr(types, subject)
			require.NoError(t, err)

			assert.Equal(t, "false", expr)
			assert.Nil(t, args)
		})

		// All types allow direct: the WHERE clause already restricts rows to those types,
		// so EXISTS runs unconditionally — no extra condition needed.
		t.Run("all AllowsDirect=true: plain EXISTS, no extra conditions ", func(t *testing.T) {
			types := []relationtuple.SubjectSetType{
				{Namespace: "Group", Relation: "member", AllowsDirect: true},
				{Namespace: "Org", Relation: "admin", AllowsDirect: true},
			}
			expr, args, err := buildFoundExpr(types, subject)
			require.NoError(t, err)

			assert.Equal(t, existsExpr, expr)
			assert.Equal(t, []any{subject.Namespace, subject.Object, subject.Relation}, args)
		})

		// Mixed: EXISTS is gated to fire only for the AllowsDirect=true type;
		// the false type is excluded from the gate even though it passes the WHERE filter.
		t.Run("mixed AllowsDirect: EXISTS gated only on allowed type", func(t *testing.T) {
			types := []relationtuple.SubjectSetType{
				{Namespace: "Group", Relation: "member", AllowsDirect: true},
				{Namespace: "Org", Relation: "admin", AllowsDirect: false},
			}
			expr, args, err := buildFoundExpr(types, subject)
			require.NoError(t, err)

			assert.Equal(t,
				`(`+existsExpr+` AND (current.subject_set_namespace, current.subject_set_relation) IN ((?, ?)))`,
				expr,
			)
			assert.Equal(t, []any{subject.Namespace, subject.Object, subject.Relation, "Group", "member"}, args)
		})

		t.Run("mixed AllowsDirect: multiple allowed types with not-allowed type", func(t *testing.T) {
			types := []relationtuple.SubjectSetType{
				{Namespace: "Group", Relation: "member", AllowsDirect: true},
				{Namespace: "Org", Relation: "admin", AllowsDirect: true},
				{Namespace: "Team", Relation: "lead", AllowsDirect: false},
			}
			expr, args, err := buildFoundExpr(types, subject)
			require.NoError(t, err)

			assert.Equal(t,
				`(`+existsExpr+` AND (current.subject_set_namespace, current.subject_set_relation) IN ((?, ?), (?, ?)))`,
				expr,
			)
			assert.Equal(t, []any{subject.Namespace, subject.Object, subject.Relation, "Group", "member", "Org", "admin"}, args)
		})
	})
}

func TestBuildSubjectSetTypeSQL(t *testing.T) {
	t.Parallel()

	t.Run("no types: no filter", func(t *testing.T) {
		sqlFragment, args := buildSubjectSetTypeSQL(nil)
		assert.Empty(t, sqlFragment)
		assert.Nil(t, args)
	})

	t.Run("multiple types share one IN list", func(t *testing.T) {
		sqlFragment, args := buildSubjectSetTypeSQL([]relationtuple.SubjectSetType{
			{Namespace: "Group", Relation: "member"},
			{Namespace: "Org", Relation: "admin"},
		})
		assert.Equal(t, "\nAND (current.subject_set_namespace, current.subject_set_relation) IN ((?, ?), (?, ?))", sqlFragment)
		assert.Equal(t, []any{"Group", "member", "Org", "admin"}, args)
	})
}
