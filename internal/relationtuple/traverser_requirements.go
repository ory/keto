// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TraverserTest(t *testing.T, m Manager, tr Traverser) {
	nspace := strconv.Itoa(rand.Int()) // nolint

	t.Run("method=FindTupleWithRelations", func(t *testing.T) {
		t.Run("case=nil relations returns nil", func(t *testing.T) {
			result, err := tr.FindTupleWithRelations(t.Context(), &RelationTuple{}, nil)
			require.NoError(t, err)
			assert.Nil(t, result)
		})

		t.Run("case=matching relation returns tuple", func(t *testing.T) {
			tuple := &RelationTuple{
				Namespace: nspace,
				Object:    uuid.Must(uuid.NewV4()),
				Relation:  "viewer",
				Subject:   &SubjectSet{Namespace: "User", Object: uuid.Must(uuid.NewV4())},
			}
			require.NoError(t, m.WriteRelationTuples(t.Context(), tuple))

			cpy := *tuple
			cpy.Relation = "editor"
			result, err := tr.FindTupleWithRelations(t.Context(), &cpy, []string{"viewer"})
			require.NoError(t, err)
			assert.Equal(t, tuple, result)
		})

		t.Run("case=non-matching relation returns nil", func(t *testing.T) {
			tuple := &RelationTuple{
				Namespace: nspace,
				Object:    uuid.Must(uuid.NewV4()),
				Relation:  "viewer",
				Subject:   &SubjectID{ID: uuid.Must(uuid.NewV4())},
			}
			require.NoError(t, m.WriteRelationTuples(t.Context(), tuple))

			result, err := tr.FindTupleWithRelations(t.Context(), tuple, []string{"editor"})
			require.NoError(t, err)
			assert.Nil(t, result)
		})

		t.Run("case=multiple relations with one matching returns tuple", func(t *testing.T) {
			tuple := &RelationTuple{
				Namespace: nspace,
				Object:    uuid.Must(uuid.NewV4()),
				Relation:  "viewer",
				Subject:   &SubjectID{ID: uuid.Must(uuid.NewV4())},
			}
			require.NoError(t, m.WriteRelationTuples(t.Context(), tuple))

			result, err := tr.FindTupleWithRelations(t.Context(), tuple, []string{"editor", "viewer", "owner"})
			require.NoError(t, err)
			assert.Equal(t, tuple, result)
		})
	})

	t.Run("method=TraverseSubjectSetExpansion", func(t *testing.T) {
		t.Run("case=no subject set tuples returns empty", func(t *testing.T) {
			start := &RelationTuple{
				Namespace: nspace,
				Object:    uuid.Must(uuid.NewV4()),
				Relation:  "rel",
				Subject:   &SubjectID{ID: uuid.Must(uuid.NewV4())},
			}

			res, err := tr.TraverseSubjectSetExpansion(t.Context(), start)
			require.NoError(t, err)
			assert.Empty(t, res)
		})

		t.Run("case=subject found in subject set", func(t *testing.T) {
			groupNs := strconv.Itoa(rand.Int()) // nolint
			obj := uuid.Must(uuid.NewV4())
			groupObj := uuid.Must(uuid.NewV4())
			userID := uuid.Must(uuid.NewV4())

			// ns:obj#rel@(groupNs:groupObj#member)
			require.NoError(t, m.WriteRelationTuples(t.Context(),
				&RelationTuple{
					Namespace: nspace,
					Object:    obj,
					Relation:  "rel",
					Subject:   &SubjectSet{Namespace: groupNs, Object: groupObj, Relation: "member"},
				},
				// groupNs:groupObj#member@userID
				&RelationTuple{
					Namespace: groupNs,
					Object:    groupObj,
					Relation:  "member",
					Subject:   &SubjectID{ID: userID},
				},
			))

			start := &RelationTuple{
				Namespace: nspace,
				Object:    obj,
				Relation:  "rel",
				Subject:   &SubjectID{ID: userID},
			}
			res, err := tr.TraverseSubjectSetExpansion(t.Context(), start)
			require.NoError(t, err)
			require.Len(t, res, 1)
			assert.True(t, res[0].Found)
			assert.Equal(t, TraversalSubjectSetExpand, res[0].Via)
			assert.Equal(t, start, res[0].From)
			assert.Equal(t, groupNs, res[0].To.Namespace)
			assert.Equal(t, groupObj, res[0].To.Object)
			assert.Equal(t, "member", res[0].To.Relation)
		})

		t.Run("case=subject not found in subject set", func(t *testing.T) {
			groupNs := strconv.Itoa(rand.Int()) // nolint
			obj := uuid.Must(uuid.NewV4())
			groupObj := uuid.Must(uuid.NewV4())

			// ns:obj#rel@(groupNs:groupObj#member), but no membership tuple
			require.NoError(t, m.WriteRelationTuples(t.Context(),
				&RelationTuple{
					Namespace: nspace,
					Object:    obj,
					Relation:  "rel",
					Subject:   &SubjectSet{Namespace: groupNs, Object: groupObj, Relation: "member"},
				},
			))

			start := &RelationTuple{
				Namespace: nspace,
				Object:    obj,
				Relation:  "rel",
				Subject:   &SubjectID{ID: uuid.Must(uuid.NewV4())},
			}
			res, err := tr.TraverseSubjectSetExpansion(t.Context(), start)
			require.NoError(t, err)
			require.Len(t, res, 1)
			assert.False(t, res[0].Found)
		})
	})
}
