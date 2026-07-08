// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/gofrs/uuid"
	keysetpagination "github.com/ory/x/pagination/keysetpagination_v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/testhelpers"
	"github.com/ory/keto/ketoapi"
)

func ManagerTest(t *testing.T, m relationtuple.Manager) {
	t.Run("method=Write", func(t *testing.T) {
		t.Run("case=success", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()
			tuples := []*relationtuple.RelationTuple{
				tf.Tuple(t, "Obj:o1#rel@s1"),
				tf.Tuple(t, "Obj:o2#rel@Obj:o3#sub rel"),
			}

			require.NoError(t, m.WriteRelationTuples(t.Context(), tuples...))

			resp, nextPage, err := m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
				Namespace: new(tf.NS("Obj")),
			})
			require.NoError(t, err)
			assert.True(t, nextPage.IsLast())
			assert.ElementsMatch(t, tuples, resp)
		})
	})

	t.Run("method=Get", func(t *testing.T) {
		t.Run("case=queries", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()
			nspace := tf.NS("Obj")
			obj0 := tf.UUID("o0")

			// Tuple i has object o{i%2}, relation "r {i%4}", and subject s{i}.
			tuples := make([]*relationtuple.RelationTuple, 10)
			for i := range tuples {
				tuples[i] = tf.Tuple(t, fmt.Sprintf("Obj:o%d#r %d@s%d", i%2, i%4, i))
			}

			require.NoError(t, m.WriteRelationTuples(t.Context(), tuples...))

			for i, tc := range []struct {
				query    *relationtuple.RelationQuery
				expected []*relationtuple.RelationTuple
			}{
				{
					query: &relationtuple.RelationQuery{
						Namespace: &nspace,
					},
					expected: tuples,
				},
				{
					query: &relationtuple.RelationQuery{
						Namespace: &nspace,
						Object:    &obj0,
					},
					expected: []*relationtuple.RelationTuple{
						tuples[0],
						tuples[2],
						tuples[4],
						tuples[6],
						tuples[8],
					},
				},
				{
					query: &relationtuple.RelationQuery{
						Namespace: &nspace,
						Relation:  new("r 0"),
					},
					expected: []*relationtuple.RelationTuple{
						tuples[0],
						tuples[4],
						tuples[8],
					},
				},
				{
					query: &relationtuple.RelationQuery{
						Namespace: &nspace,
						Object:    &obj0,
						Relation:  new("r 0"),
					},
					expected: []*relationtuple.RelationTuple{
						tuples[0],
						tuples[4],
						tuples[8],
					},
				},
				{
					query: &relationtuple.RelationQuery{
						Namespace: &nspace,
						Subject:   tuples[0].Subject,
					},
					expected: []*relationtuple.RelationTuple{
						tuples[0],
					},
				},
				{
					query: &relationtuple.RelationQuery{
						Namespace: &nspace,
						Object:    &obj0,
						Subject:   tuples[0].Subject,
					},
					expected: []*relationtuple.RelationTuple{
						tuples[0],
					},
				},
				{
					query: &relationtuple.RelationQuery{
						Namespace: &nspace,
						Relation:  new("r 0"),
						Subject:   tuples[0].Subject,
					},
					expected: []*relationtuple.RelationTuple{
						tuples[0],
					},
				},
				{
					query: &relationtuple.RelationQuery{
						Namespace: &nspace,
						Object:    &obj0,
						Relation:  new("r 0"),
						Subject:   tuples[0].Subject,
					},
					expected: []*relationtuple.RelationTuple{
						tuples[0],
					},
				},
			} {
				t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
					res, nextPage, err := m.GetRelationTuples(t.Context(), tc.query)
					require.NoError(t, err)
					assert.True(t, nextPage.IsLast())
					assert.ElementsMatch(t, tc.expected, res)
				})
			}
		})

		t.Run("case=pagination", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()

			tuples := make([]*relationtuple.RelationTuple, 20)
			for i := range tuples {
				tuples[i] = tf.Tuple(t, fmt.Sprintf("Obj:o#r@s%d", i))
			}

			require.NoError(t, m.WriteRelationTuples(t.Context(), tuples...))

			tests := []struct {
				name           string
				searchCriteria *relationtuple.RelationQuery
			}{
				{
					name: "search=ns,obj,rel",
					searchCriteria: &relationtuple.RelationQuery{
						Namespace: new(tf.NS("Obj")),
						Object:    new(tf.UUID("o")),
						Relation:  new("r"),
					},
				},
				{
					name: "search=ns",
					searchCriteria: &relationtuple.RelationQuery{
						Namespace: new(tf.NS("Obj")),
						Object:    nil,
						Relation:  nil,
					},
				},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					var (
						res, thisPage []*relationtuple.RelationTuple
						err           error
					)
					nextPage, err := keysetpagination.NewPaginator(keysetpagination.WithSize(1))
					require.NoError(t, err)
					for range tuples[:len(tuples)-1] {
						thisPage, nextPage, err = m.GetRelationTuples(t.Context(), tt.searchCriteria, nextPage.ToOptions()...)
						require.NoError(t, err)
						assert.False(t, nextPage.IsLast())
						require.Len(t, thisPage, 1)

						res = append(res, thisPage[0])
					}

					thisPage, nextPage, err = m.GetRelationTuples(t.Context(), tt.searchCriteria, nextPage.ToOptions()...)
					require.NoError(t, err)
					assert.True(t, nextPage.IsLast())
					require.Len(t, thisPage, 1)

					res = append(res, thisPage[0])
					assert.ElementsMatch(t, tuples, res)
				})
			}
		})

		// Comparing paginated results against a non-paginated fetch,
		// asserting order and correctness with different page sizes
		t.Run("case=pagination returns all rows without skips in correct order", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()

			// 11 tuples, mixing SubjectID and SubjectSet subjects, including one
			// exact duplicate row.
			duplicate := tf.Tuple(t, "Obj:o#r@teams:t1#editor")
			tuples := []*relationtuple.RelationTuple{
				tf.Tuple(t, "Obj:o#r@groups:g1#member"),
				tf.Tuple(t, "Obj:o#r@s1"),
				tf.Tuple(t, "Obj:o#r@s2"),
				tf.Tuple(t, "Obj:o#r@teams:t2#viewer"),
				tf.Tuple(t, "Obj:o#r@groups:g2#owner"),
				tf.Tuple(t, "Obj:o#r@s3"),
				tf.Tuple(t, "Obj:o#r@orgs:org1#admin"),
				tf.Tuple(t, "Obj:o#r@s4"),
				duplicate,
				duplicate,
				tf.Tuple(t, "Obj:o#r@s5"),
			}

			require.NoError(t, m.WriteRelationTuples(t.Context(), tuples...))

			tests := []struct {
				name           string
				searchCriteria *relationtuple.RelationQuery
				pageSizes      []int
			}{
				{
					name: "search=ns",
					searchCriteria: &relationtuple.RelationQuery{
						Namespace: new(tf.NS("Obj")),
					},
					pageSizes: []int{1, 2, 50},
				},
				{
					name: "search=obj",
					searchCriteria: &relationtuple.RelationQuery{
						Object: new(tf.UUID("o")),
					},
					pageSizes: []int{1, 2, 50},
				},
				{
					name: "search=ns,obj,rel",
					searchCriteria: &relationtuple.RelationQuery{
						Namespace: new(tf.NS("Obj")),
						Object:    new(tf.UUID("o")),
						Relation:  new("r"),
					},
					pageSizes: []int{1, 2, 50},
				},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					expected, next, err := m.GetRelationTuples(t.Context(), tt.searchCriteria, keysetpagination.WithSize(9999))
					require.NoError(t, err)
					require.True(t, next.IsLast())
					require.Len(t, expected, len(tuples))

					for _, pageSize := range tt.pageSizes {
						nextPage, _ := keysetpagination.NewPaginator(keysetpagination.WithSize(pageSize))

						var got []*relationtuple.RelationTuple
						for {
							page, np, err := m.GetRelationTuples(t.Context(), tt.searchCriteria, nextPage.ToOptions()...)
							require.NoError(t, err)

							require.LessOrEqual(t, len(page), pageSize)
							if !np.IsLast() {
								require.Len(t, page, pageSize)
							}

							got = append(got, page...)
							nextPage = np

							if nextPage.IsLast() {
								break
							}
						}

						// Same order as non-paginated fetch.
						require.Equal(t, expected, got)
					}
				})
			}
		})

		t.Run("case=empty list", func(t *testing.T) {
			t.Parallel()
			nspace := t.Name()

			res, nextPage, err := m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
				Namespace: &nspace,
			})

			assert.NoError(t, err)
			assert.Equal(t, []*relationtuple.RelationTuple{}, res)
			assert.True(t, nextPage.IsLast())
		})
	})

	t.Run("method=Delete", func(t *testing.T) {
		t.Run("case=deletes tuple", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()

			for _, rt := range []*relationtuple.RelationTuple{
				tf.Tuple(t, "Obj:o#r to delete@s"),
				tf.Tuple(t, "Obj:o#r to delete@Obj:o2#r2"),
			} {
				t.Run(fmt.Sprintf("subject_type=%T", rt.Subject), func(t *testing.T) {
					require.NoError(t, m.WriteRelationTuples(t.Context(), rt))

					res, _, err := m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
						Namespace: new(tf.NS("Obj")),
					})
					require.NoError(t, err)
					assert.Equal(t, []*relationtuple.RelationTuple{rt}, res)

					require.NoError(t, m.DeleteRelationTuples(t.Context(), rt))

					res, _, err = m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
						Namespace: new(tf.NS("Obj")),
					})
					require.NoError(t, err)
					assert.Len(t, res, 0)
				})
			}
		})

		t.Run("case=deletes only one tuple", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()

			rs := make([]*relationtuple.RelationTuple, 4)
			for i := range rs {
				rs[i] = tf.Tuple(t, fmt.Sprintf("Obj:o%d#r%d@s%d", i, i, i))
			}
			require.NoError(t, m.WriteRelationTuples(t.Context(), rs...))

			res, _, err := m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
				Namespace: new(tf.NS("Obj")),
			})
			require.NoError(t, err)
			for _, rt := range rs {
				assert.Contains(t, res, rt)
			}

			require.NoError(t, m.DeleteRelationTuples(t.Context(), rs[0], rs[2]))

			res, _, err = m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
				Namespace: new(tf.NS("Obj")),
			})
			require.NoError(t, err)
			assert.ElementsMatch(t, []*relationtuple.RelationTuple{rs[1], rs[3]}, res)
		})

		t.Run("case=tuple and subject namespace differ", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()

			rt := tf.Tuple(t, "N0:o#r@N1:o#r")
			require.NoError(t, m.WriteRelationTuples(t.Context(), rt))

			actual, _, err := m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{Namespace: new(tf.NS("N0"))})
			require.NoError(t, err)
			assert.Equal(t, []*relationtuple.RelationTuple{rt}, actual)

			require.NoError(t, m.DeleteRelationTuples(t.Context(), rt))

			actual, _, err = m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{Namespace: new(tf.NS("N0"))})
			require.NoError(t, err)
			assert.Len(t, actual, 0)
		})
	})

	t.Run("method=Transact", func(t *testing.T) {
		t.Run("case=success", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()

			rs := make([]*relationtuple.RelationTuple, 4)
			for i := range rs {
				rs[i] = tf.Tuple(t, fmt.Sprintf("Obj:o%d#r%d@s%d", i, i, i))
			}
			require.NoError(t, m.WriteRelationTuples(t.Context(), rs[0], rs[1]))

			res, _, err := m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
				Namespace: new(tf.NS("Obj")),
			})
			require.NoError(t, err)
			assert.ElementsMatch(t, []*relationtuple.RelationTuple{rs[0], rs[1]}, res)

			require.NoError(t, m.TransactRelationTuples(t.Context(), []*relationtuple.RelationTuple{rs[2], rs[3]}, []*relationtuple.RelationTuple{rs[0]}))

			res, _, err = m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
				Namespace: new(tf.NS("Obj")),
			})
			require.NoError(t, err)

			for _, rt := range []*relationtuple.RelationTuple{rs[1], rs[2], rs[3]} {
				assert.Contains(t, res, rt)
			}
		})

		t.Run("case=err rolls back all", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()

			rs := []*relationtuple.RelationTuple{
				tf.Tuple(t, "Obj:o0#r0@s0"),
				tf.Tuple(t, "Obj:o1#r1@s1"),
			}
			invalidRt := &relationtuple.RelationTuple{
				Namespace: tf.NS("Obj"),
				Object:    tf.UUID("o0"),
				Relation:  "r0",
				Subject:   nil, // subject is not allowed to be nil
			}
			require.NoError(t, m.WriteRelationTuples(t.Context(), rs[0]))

			res, _, err := m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
				Namespace: new(tf.NS("Obj")),
			})
			require.NoError(t, err)
			assert.Equal(t, []*relationtuple.RelationTuple{rs[0]}, res)

			t.Run("invalid=insert", func(t *testing.T) {
				assert.ErrorIs(t, m.TransactRelationTuples(t.Context(), []*relationtuple.RelationTuple{invalidRt}, []*relationtuple.RelationTuple{rs[0]}), ketoapi.ErrNilSubject())
			})

			res, _, err = m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
				Namespace: new(tf.NS("Obj")),
			})
			require.NoError(t, err)
			assert.Equal(t, []*relationtuple.RelationTuple{rs[0]}, res)

			t.Run("invalid=delete", func(t *testing.T) {
				assert.ErrorIs(t, m.TransactRelationTuples(t.Context(), []*relationtuple.RelationTuple{rs[1]}, []*relationtuple.RelationTuple{invalidRt}), ketoapi.ErrNilSubject())
			})

			res, _, err = m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
				Namespace: new(tf.NS("Obj")),
			})
			require.NoError(t, err)
			assert.Equal(t, []*relationtuple.RelationTuple{rs[0]}, res)
		})
	})

	traversalTest(t, m)
}

func traversalTest(t *testing.T, m relationtuple.Manager) {
	t.Run("method=FindTupleWithRelations", func(t *testing.T) {
		t.Run("case=nil relations returns nil", func(t *testing.T) {
			t.Parallel()
			result, err := m.FindTupleWithRelations(t.Context(), &relationtuple.RelationTuple{}, nil)
			require.NoError(t, err)
			assert.Nil(t, result)
		})

		t.Run("case=matching relation returns tuple", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()
			tuple := tf.Tuple(t, "Obj:o#viewer@User:u")
			require.NoError(t, m.WriteRelationTuples(t.Context(), tuple))

			// The query tuple's own relation is ignored; only the relations list matters.
			result, err := m.FindTupleWithRelations(t.Context(), tf.Tuple(t, "Obj:o#editor@User:u"), []string{"viewer"})
			require.NoError(t, err)
			assert.Equal(t, tuple, result)
		})

		t.Run("case=non-matching relation returns nil", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()
			tuple := tf.Tuple(t, "Obj:o#viewer@u")
			require.NoError(t, m.WriteRelationTuples(t.Context(), tuple))

			result, err := m.FindTupleWithRelations(t.Context(), tuple, []string{"editor"})
			require.NoError(t, err)
			assert.Nil(t, result)
		})

		t.Run("case=multiple relations with one matching returns tuple", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()
			tuple := tf.Tuple(t, "Obj:o#viewer@u")
			require.NoError(t, m.WriteRelationTuples(t.Context(), tuple))

			result, err := m.FindTupleWithRelations(t.Context(), tuple, []string{"editor", "viewer", "owner"})
			require.NoError(t, err)
			assert.Equal(t, tuple, result)
		})
	})

	t.Run("method=TraverseSubjectSetExpansion", func(t *testing.T) {
		t.Run("case=no subject set tuples returns empty", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()
			res, err := m.TraverseSubjectSetExpansion(t.Context(), tf.Tuple(t, "Obj:o#rel@u"), nil)
			require.NoError(t, err)
			assert.Empty(t, res)
		})

		t.Run("case=subject found in subject set", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()
			membership := tf.Tuple(t, "Group:g#member@u")
			require.NoError(t, m.WriteRelationTuples(t.Context(),
				tf.Tuple(t, "Obj:o#rel@Group:g#member"),
				membership,
			))

			start := tf.Tuple(t, "Obj:o#rel@u")
			res, err := m.TraverseSubjectSetExpansion(t.Context(), start, nil)
			require.NoError(t, err)
			require.Len(t, res, 1)
			assert.True(t, res[0].Found)
			assert.Equal(t, relationtuple.TraversalSubjectSetExpand, res[0].Via)
			assert.Equal(t, start, res[0].From)
			assert.Equal(t, membership.Namespace, res[0].To.Namespace)
			assert.Equal(t, membership.Object, res[0].To.Object)
			assert.Equal(t, membership.Relation, res[0].To.Relation)
		})

		t.Run("case=subject not found in subject set", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()
			// Subject-set pointer without a membership tuple behind it.
			require.NoError(t, m.WriteRelationTuples(t.Context(),
				tf.Tuple(t, "Obj:o#rel@Group:g#member"),
			))

			res, err := m.TraverseSubjectSetExpansion(t.Context(), tf.Tuple(t, "Obj:o#rel@u"), nil)
			require.NoError(t, err)
			require.Len(t, res, 1)
			assert.False(t, res[0].Found)
		})

		t.Run("case=allowedSubjectSets filters out non-matching types", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()
			require.NoError(t, m.WriteRelationTuples(t.Context(),
				tf.Tuple(t, "Obj:o#rel@Doc:d#member"),
				tf.Tuple(t, "Obj:o#rel@Group:g#admin"),
			))

			res, err := m.TraverseSubjectSetExpansion(t.Context(), tf.Tuple(t, "Obj:o#rel@u"), []relationtuple.SubjectSetType{{Namespace: tf.NS("Doc"), Relation: "member"}})
			require.NoError(t, err)
			require.Len(t, res, 1)
			assert.Equal(t, tf.NS("Doc"), res[0].To.Namespace)
			assert.Equal(t, "member", res[0].To.Relation)
		})

		t.Run("case=nil allowedSubjectSets returns all types", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()
			require.NoError(t, m.WriteRelationTuples(t.Context(),
				tf.Tuple(t, "Obj:o#rel@Doc:d#member"),
				tf.Tuple(t, "Obj:o#rel@Group:g#admin"),
			))

			res, err := m.TraverseSubjectSetExpansion(t.Context(), tf.Tuple(t, "Obj:o#rel@u"), nil)
			require.NoError(t, err)
			assert.Len(t, res, 2)
		})

		t.Run("case=AllowsDirect=false direct relationship is not checked", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()
			require.NoError(t, m.WriteRelationTuples(t.Context(),
				tf.Tuple(t, "Obj:o#rel@Group:g#member"),
				tf.Tuple(t, "Group:g#member@User:u"),
			))

			start := tf.Tuple(t, "Obj:o#rel@User:u")
			res, err := m.TraverseSubjectSetExpansion(t.Context(), start, []relationtuple.SubjectSetType{{Namespace: tf.NS("Group"), Relation: "member", AllowsDirect: false}})
			require.NoError(t, err)
			require.Len(t, res, 1)
			assert.False(t, res[0].Found)
		})

		t.Run("case=AllowsDirect gates EXISTS per type", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()
			require.NoError(t, m.WriteRelationTuples(t.Context(),
				tf.Tuple(t, "Obj:o#rel@Allowed:a#member"),
				tf.Tuple(t, "Obj:o#rel@NotAllowed:n#member"),
				tf.Tuple(t, "Allowed:a#member@User:u"),
				tf.Tuple(t, "NotAllowed:n#member@User:u"),
			))

			start := tf.Tuple(t, "Obj:o#rel@User:u")

			t.Run("non-allowed type returns found=false", func(t *testing.T) {
				t.Parallel()
				res, err := m.TraverseSubjectSetExpansion(t.Context(), start,
					[]relationtuple.SubjectSetType{{Namespace: tf.NS("NotAllowed"), Relation: "member", AllowsDirect: false}},
				)
				require.NoError(t, err)
				require.Len(t, res, 1)
				assert.Equal(t, tf.NS("NotAllowed"), res[0].To.Namespace)
				assert.False(t, res[0].Found)
			})

			t.Run("mixed: allowed returns found=true, non-allowed returns found=false", func(t *testing.T) {
				t.Parallel()
				res, err := m.TraverseSubjectSetExpansion(t.Context(), start, []relationtuple.SubjectSetType{
					{Namespace: tf.NS("Allowed"), Relation: "member", AllowsDirect: true},
					{Namespace: tf.NS("NotAllowed"), Relation: "member", AllowsDirect: false},
				})
				require.NoError(t, err)
				// TraverseSubjectSetExpansion can short-circuit if it reads the allowed
				// namespace first, returning only 1 item.
				require.GreaterOrEqual(t, len(res), 1)

				byNs := make(map[string]*relationtuple.TraversalResult)
				for _, r := range res {
					byNs[r.To.Namespace] = r
				}
				assert.True(t, byNs[tf.NS("Allowed")].Found)
			})
		})

		t.Run("case=multiple AllowsDirect types are checked in one query", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()
			require.NoError(t, m.WriteRelationTuples(t.Context(),
				tf.Tuple(t, "Obj:o#rel@GroupA:a#member"),
				tf.Tuple(t, "Obj:o#rel@GroupB:b#member"),
				tf.Tuple(t, "Obj:o#rel@Stale:s#member"),
				// A stale membership under the type that does not allow direct matches.
				tf.Tuple(t, "Stale:s#member@User:u"),
			))

			start := tf.Tuple(t, "Obj:o#rel@User:u")
			types := []relationtuple.SubjectSetType{
				{Namespace: tf.NS("GroupA"), Relation: "member", AllowsDirect: true},
				{Namespace: tf.NS("GroupB"), Relation: "member", AllowsDirect: true},
				{Namespace: tf.NS("Stale"), Relation: "member", AllowsDirect: false},
			}

			// Nothing is found, so the traversal cannot short-circuit and must
			// return all three pointers with found=false.
			res, err := m.TraverseSubjectSetExpansion(t.Context(), start, types)
			require.NoError(t, err)
			require.Len(t, res, 3)
			for _, r := range res {
				assert.False(t, r.Found, "namespace %s must not be found", r.To.Namespace)
			}
		})

		t.Run("case=multiple allowedSubjectSets entries each match independently", func(t *testing.T) {
			t.Parallel()
			tf := testhelpers.NewTupleFactory()
			require.NoError(t, m.WriteRelationTuples(t.Context(),
				tf.Tuple(t, "Obj:o#rel@Doc:d#member"),
				tf.Tuple(t, "Obj:o#rel@Group:g#admin"),
				// Not in the allowed list below.
				tf.Tuple(t, "Obj:o#rel@Group:g2#viewer"),
			))

			res, err := m.TraverseSubjectSetExpansion(t.Context(), tf.Tuple(t, "Obj:o#rel@u"), []relationtuple.SubjectSetType{
				{Namespace: tf.NS("Doc"), Relation: "member"},
				{Namespace: tf.NS("Group"), Relation: "admin"},
			})
			require.NoError(t, err)
			assert.Len(t, res, 2)
		})
	})
}

func twice(t *testing.T, m0, m1 relationtuple.Manager) func(string, func(*testing.T, relationtuple.Manager, relationtuple.Manager)) {
	return func(desc string, run func(*testing.T, relationtuple.Manager, relationtuple.Manager)) {
		t.Run(desc+"-0", func(t *testing.T) {
			run(t, m0, m1)
		})
		t.Run(desc+"-1", func(t *testing.T) {
			run(t, m1, m0)
		})
	}
}

func reset(t *testing.T, ms ...relationtuple.Manager) {
	for _, m := range ms {
		for {
			rts, next, err := m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{})
			require.NoError(t, err)

			require.NoError(t, m.DeleteRelationTuples(t.Context(), rts...))

			if next.IsLast() {
				break
			}
		}
	}
}

func IsolationTest(t *testing.T, m0, m1 relationtuple.Manager) {
	run := twice(t, m0, m1)

	run("suite=lifecycle", func(t *testing.T, m0, m1 relationtuple.Manager) {
		nspace := t.Name()

		rts := []*relationtuple.RelationTuple{
			{
				Namespace: nspace,
				Object:    uuid.Must(uuid.NewV4()),
				Relation:  "r",
				Subject: &relationtuple.SubjectSet{
					Namespace: nspace,
					Object:    uuid.Must(uuid.NewV4()),
					Relation:  "r",
				},
			},
			{
				Namespace: nspace,
				Object:    uuid.Must(uuid.NewV4()),
				Relation:  "r",
				Subject:   &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
			},
		}

		t.Run("case=write and get", func(t *testing.T) {
			reset(t, m0, m1)

			require.NoError(t, m0.WriteRelationTuples(t.Context(), rts...))

			other, _, err := m1.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{Namespace: &nspace})
			require.NoError(t, err)
			assert.Len(t, other, 0)

			actual, _, err := m0.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{Namespace: &nspace})
			require.NoError(t, err)
			assert.ElementsMatch(t, rts, actual)
		})

		t.Run("case=delete", func(t *testing.T) {
			reset(t, m0, m1)

			require.NoError(t, m1.WriteRelationTuples(t.Context(), rts...))

			require.NoError(t, m0.DeleteRelationTuples(t.Context(), rts...))

			deleted, _, err := m0.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{Namespace: &nspace})
			require.NoError(t, err)
			assert.Len(t, deleted, 0)

			actual, _, err := m1.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{Namespace: &nspace})
			require.NoError(t, err)
			assert.ElementsMatch(t, rts, actual)
		})

		t.Run("case=transact", func(t *testing.T) {
			reset(t, m0, m1)

			// note that the reset is outside this subtest, so in the second run we actually delete what we had before
			twice(t, m0, m1)("insert and delete", func(t *testing.T, m0, m1 relationtuple.Manager) {
				require.NoError(t, m0.TransactRelationTuples(t.Context(), []*relationtuple.RelationTuple{rts[0]}, []*relationtuple.RelationTuple{rts[1]}))

				require.NoError(t, m1.TransactRelationTuples(t.Context(), []*relationtuple.RelationTuple{rts[1]}, []*relationtuple.RelationTuple{rts[0]}))

				r0, _, err := m0.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{Namespace: &nspace})
				require.NoError(t, err)

				r1, _, err := m1.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{Namespace: &nspace})
				require.NoError(t, err)

				assert.Equal(t, rts[:1], r0)
				assert.Equal(t, rts[1:], r1)
			})
		})

		t.Run("case=cancelled", func(t *testing.T) {
			reset(t, m0, m1)
			ctx, cancel := context.WithCancel(t.Context())

			require.NoError(t, m0.WriteRelationTuples(ctx, rts...))

			cancel()

			_, _, err := m0.GetRelationTuples(ctx, &relationtuple.RelationQuery{Namespace: &nspace})
			assert.ErrorIs(t, err, context.Canceled)
		})
	})
}
