// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/gofrs/uuid"
	keysetpagination "github.com/ory/x/pagination/keysetpagination_v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoapi"
)

func ManagerTest(t *testing.T, m relationtuple.Manager) {
	t.Run("method=Write", func(t *testing.T) {
		t.Run("case=success", func(t *testing.T) {
			nspace := strconv.Itoa(rand.Int()) // nolint

			tuples := []*relationtuple.RelationTuple{
				{
					Namespace: nspace,
					Object:    uuid.Must(uuid.NewV4()),
					Relation:  "rel",
					Subject:   &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
				},
				{
					Namespace: nspace,
					Object:    uuid.Must(uuid.NewV4()),
					Relation:  "rel",
					Subject: &relationtuple.SubjectSet{
						Namespace: nspace,
						Object:    uuid.Must(uuid.NewV4()),
						Relation:  "sub rel",
					},
				},
			}

			require.NoError(t, m.WriteRelationTuples(t.Context(), tuples...))

			resp, nextPage, err := m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
				Namespace: new(nspace),
			})
			require.NoError(t, err)
			assert.True(t, nextPage.IsLast())
			assert.ElementsMatch(t, tuples, resp)
		})
	})

	t.Run("method=Get", func(t *testing.T) {
		t.Run("case=queries", func(t *testing.T) {
			nspace := strconv.Itoa(rand.Int()) // nolint

			tuples := make([]*relationtuple.RelationTuple, 10)
			ids := x.UUIDs(len(tuples))

			for i := range tuples {
				tuples[i] = &relationtuple.RelationTuple{
					Namespace: nspace,
					Object:    ids[i%2],
					Relation:  fmt.Sprintf("r %d", i%4),
					Subject:   &relationtuple.SubjectID{ID: ids[i]},
				}
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
						Object:    &ids[0],
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
						Object:    &ids[0],
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
						Subject:   &relationtuple.SubjectID{ID: ids[0]},
					},
					expected: []*relationtuple.RelationTuple{
						tuples[0],
					},
				},
				{
					query: &relationtuple.RelationQuery{
						Namespace: &nspace,
						Object:    &ids[0],
						Subject:   &relationtuple.SubjectID{ID: ids[0]},
					},
					expected: []*relationtuple.RelationTuple{
						tuples[0],
					},
				},
				{
					query: &relationtuple.RelationQuery{
						Namespace: &nspace,
						Relation:  new("r 0"),
						Subject:   &relationtuple.SubjectID{ID: ids[0]},
					},
					expected: []*relationtuple.RelationTuple{
						tuples[0],
					},
				},
				{
					query: &relationtuple.RelationQuery{
						Namespace: &nspace,
						Object:    &ids[0],
						Relation:  new("r 0"),
						Subject:   &relationtuple.SubjectID{ID: ids[0]},
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
			nspace := strconv.Itoa(rand.Int()) // nolint

			tuples := make([]*relationtuple.RelationTuple, 20)
			oID := uuid.Must(uuid.NewV4())
			for i := range tuples {
				tuples[i] = &relationtuple.RelationTuple{
					Namespace: nspace,
					Object:    oID,
					Relation:  "r",
					Subject:   &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
				}
			}

			require.NoError(t, m.WriteRelationTuples(t.Context(), tuples...))

			tests := []struct {
				name           string
				searchCriteria *relationtuple.RelationQuery
			}{
				{
					name: "search=ns,obj,rel",
					searchCriteria: &relationtuple.RelationQuery{
						Namespace: new(nspace),
						Object:    &oID,
						Relation:  new("r"),
					},
				},
				{
					name: "search=ns",
					searchCriteria: &relationtuple.RelationQuery{
						Namespace: new(nspace),
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
			ns := "ns_" + strconv.Itoa(rand.Int()) // nolint:gosec
			oID := uuid.Must(uuid.NewV4())

			// 11 tuples manually, mixing relationtuple.SubjectID and relationtuple.SubjectSet.
			duplicate := &relationtuple.RelationTuple{Namespace: ns, Object: oID, Relation: "r", Subject: &relationtuple.SubjectSet{Namespace: "teams", Object: uuid.Must(uuid.NewV4()), Relation: "editor"}}
			tuples := []*relationtuple.RelationTuple{
				{Namespace: ns, Object: oID, Relation: "r", Subject: &relationtuple.SubjectSet{Namespace: "groups", Object: uuid.Must(uuid.NewV4()), Relation: "member"}},
				{Namespace: ns, Object: oID, Relation: "r", Subject: &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())}},
				{Namespace: ns, Object: oID, Relation: "r", Subject: &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())}},
				{Namespace: ns, Object: oID, Relation: "r", Subject: &relationtuple.SubjectSet{Namespace: "teams", Object: uuid.Must(uuid.NewV4()), Relation: "viewer"}},
				{Namespace: ns, Object: oID, Relation: "r", Subject: &relationtuple.SubjectSet{Namespace: "groups", Object: uuid.Must(uuid.NewV4()), Relation: "owner"}},
				{Namespace: ns, Object: oID, Relation: "r", Subject: &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())}},
				{Namespace: ns, Object: oID, Relation: "r", Subject: &relationtuple.SubjectSet{Namespace: "orgs", Object: uuid.Must(uuid.NewV4()), Relation: "admin"}},
				{Namespace: ns, Object: oID, Relation: "r", Subject: &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())}},
				duplicate,
				duplicate,
				{Namespace: ns, Object: oID, Relation: "r", Subject: &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())}},
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
						Namespace: new(ns),
					},
					pageSizes: []int{1, 2, 50},
				},
				{
					name: "search=obj",
					searchCriteria: &relationtuple.RelationQuery{
						Object: new(oID),
					},
					pageSizes: []int{1, 2, 50},
				},
				{
					name: "search=ns,obj,rel",
					searchCriteria: &relationtuple.RelationQuery{
						Namespace: new(ns),
						Object:    new(oID),
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
			nspace := strconv.Itoa(rand.Int()) // nolint
			oID := uuid.Must(uuid.NewV4())
			sID := uuid.Must(uuid.NewV4())

			for _, rt := range []*relationtuple.RelationTuple{
				{
					Namespace: nspace,
					Object:    oID,
					Relation:  "r to delete",
					Subject:   &relationtuple.SubjectID{ID: sID},
				},
				{
					Namespace: nspace,
					Object:    oID,
					Relation:  "r to delete",
					Subject: &relationtuple.SubjectSet{
						Namespace: nspace,
						Object:    uuid.Must(uuid.NewV4()),
						Relation:  "r2",
					},
				},
			} {
				t.Run(fmt.Sprintf("subject_type=%T", rt.Subject), func(t *testing.T) {
					require.NoError(t, m.WriteRelationTuples(t.Context(), rt))

					res, _, err := m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
						Namespace: new(nspace),
					})
					require.NoError(t, err)
					assert.Equal(t, []*relationtuple.RelationTuple{rt}, res)

					require.NoError(t, m.DeleteRelationTuples(t.Context(), rt))

					res, _, err = m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
						Namespace: new(nspace),
					})
					require.NoError(t, err)
					assert.Len(t, res, 0)
				})
			}
		})

		t.Run("case=deletes only one tuple", func(t *testing.T) {
			nspace := strconv.Itoa(rand.Int()) // nolint

			rs := make([]*relationtuple.RelationTuple, 4)
			oIDs, sIDs := make([]uuid.UUID, len(rs)), make([]uuid.UUID, len(rs))
			for i := range oIDs {
				oIDs[i] = uuid.Must(uuid.NewV4())
				sIDs[i] = uuid.Must(uuid.NewV4())
			}

			for i := range rs {
				rs[i] = &relationtuple.RelationTuple{
					Namespace: nspace,
					Object:    oIDs[i],
					Relation:  "r" + strconv.Itoa(i),
					Subject:   &relationtuple.SubjectID{ID: sIDs[i]},
				}
			}
			require.NoError(t, m.WriteRelationTuples(t.Context(), rs...))

			res, _, err := m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
				Namespace: &nspace,
			})
			require.NoError(t, err)
			for _, rt := range rs {
				assert.Contains(t, res, rt)
			}

			require.NoError(t, m.DeleteRelationTuples(t.Context(), rs[0], rs[2]))

			res, _, err = m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
				Namespace: &nspace,
			})
			require.NoError(t, err)
			assert.ElementsMatch(t, []*relationtuple.RelationTuple{rs[1], rs[3]}, res)
		})

		t.Run("case=tuple and subject namespace differ", func(t *testing.T) {
			n0, n1 := t.Name()[:60], t.Name()[:60]+"1"
			oID := uuid.Must(uuid.NewV4())

			rt := &relationtuple.RelationTuple{
				Namespace: n0,
				Object:    oID,
				Relation:  "r",
				Subject: &relationtuple.SubjectSet{
					Namespace: n1,
					Object:    oID,
					Relation:  "r",
				},
			}
			require.NoError(t, m.WriteRelationTuples(t.Context(), rt))

			actual, _, err := m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{Namespace: &n0})
			require.NoError(t, err)
			assert.Equal(t, []*relationtuple.RelationTuple{rt}, actual)

			require.NoError(t, m.DeleteRelationTuples(t.Context(), rt))

			actual, _, err = m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{Namespace: &n0})
			require.NoError(t, err)
			assert.Len(t, actual, 0)
		})
	})

	t.Run("method=Transact", func(t *testing.T) {
		t.Run("case=success", func(t *testing.T) {
			nspace := strconv.Itoa(rand.Int()) // nolint

			rs := make([]*relationtuple.RelationTuple, 4)
			oIDs, sIDs := make([]uuid.UUID, len(rs)), make([]uuid.UUID, len(rs))
			for i := range oIDs {
				oIDs[i] = uuid.Must(uuid.NewV4())
				sIDs[i] = uuid.Must(uuid.NewV4())
			}

			for i := range rs {
				rs[i] = &relationtuple.RelationTuple{
					Namespace: nspace,
					Object:    oIDs[i],
					Relation:  "r" + strconv.Itoa(i),
					Subject:   &relationtuple.SubjectID{ID: sIDs[i]},
				}
			}
			require.NoError(t, m.WriteRelationTuples(t.Context(), rs[0], rs[1]))

			res, _, err := m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
				Namespace: &nspace,
			})
			require.NoError(t, err)
			assert.ElementsMatch(t, []*relationtuple.RelationTuple{rs[0], rs[1]}, res)

			require.NoError(t, m.TransactRelationTuples(t.Context(), []*relationtuple.RelationTuple{rs[2], rs[3]}, []*relationtuple.RelationTuple{rs[0]}))

			res, _, err = m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
				Namespace: &nspace,
			})
			require.NoError(t, err)

			for _, rt := range []*relationtuple.RelationTuple{rs[1], rs[2], rs[3]} {
				assert.Contains(t, res, rt)
			}
		})

		t.Run("case=err rolls back all", func(t *testing.T) {
			nspace := strconv.Itoa(rand.Int()) // nolint

			rs := make([]*relationtuple.RelationTuple, 2)
			oIDs, sIDs := make([]uuid.UUID, len(rs)), make([]uuid.UUID, len(rs))
			for i := range oIDs {
				oIDs[i] = uuid.Must(uuid.NewV4())
				sIDs[i] = uuid.Must(uuid.NewV4())
			}
			for i := range rs {
				rs[i] = &relationtuple.RelationTuple{
					Namespace: nspace,
					Object:    oIDs[i],
					Relation:  "r" + strconv.Itoa(i),
					Subject:   &relationtuple.SubjectID{ID: sIDs[i]},
				}
			}
			invalidRt := &relationtuple.RelationTuple{
				Namespace: nspace,
				Object:    oIDs[0],
				Relation:  "r0",
				Subject:   nil, // subject is not allowed to be nil
			}
			require.NoError(t, m.WriteRelationTuples(t.Context(), rs[0]))

			res, _, err := m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
				Namespace: &nspace,
			})
			require.NoError(t, err)
			assert.Equal(t, []*relationtuple.RelationTuple{rs[0]}, res)

			t.Run("invalid=insert", func(t *testing.T) {
				assert.ErrorIs(t, m.TransactRelationTuples(t.Context(), []*relationtuple.RelationTuple{invalidRt}, []*relationtuple.RelationTuple{rs[0]}), ketoapi.ErrNilSubject())
			})

			res, _, err = m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
				Namespace: &nspace,
			})
			require.NoError(t, err)
			assert.Equal(t, []*relationtuple.RelationTuple{rs[0]}, res)

			t.Run("invalid=delete", func(t *testing.T) {
				assert.ErrorIs(t, m.TransactRelationTuples(t.Context(), []*relationtuple.RelationTuple{rs[1]}, []*relationtuple.RelationTuple{invalidRt}), ketoapi.ErrNilSubject())
			})

			res, _, err = m.GetRelationTuples(t.Context(), &relationtuple.RelationQuery{
				Namespace: &nspace,
			})
			require.NoError(t, err)
			assert.Equal(t, []*relationtuple.RelationTuple{rs[0]}, res)
		})
	})

	traversalTest(t, m)
}

func traversalTest(t *testing.T, m relationtuple.Manager) {
	nspace := strconv.Itoa(rand.Int()) // nolint

	t.Run("method=FindTupleWithRelations", func(t *testing.T) {
		t.Run("case=nil relations returns nil", func(t *testing.T) {
			result, err := m.FindTupleWithRelations(t.Context(), &relationtuple.RelationTuple{}, nil)
			require.NoError(t, err)
			assert.Nil(t, result)
		})

		t.Run("case=matching relation returns tuple", func(t *testing.T) {
			tuple := &relationtuple.RelationTuple{
				Namespace: nspace,
				Object:    uuid.Must(uuid.NewV4()),
				Relation:  "viewer",
				Subject:   &relationtuple.SubjectSet{Namespace: "User", Object: uuid.Must(uuid.NewV4())},
			}
			require.NoError(t, m.WriteRelationTuples(t.Context(), tuple))

			cpy := *tuple
			cpy.Relation = "editor"
			result, err := m.FindTupleWithRelations(t.Context(), &cpy, []string{"viewer"})
			require.NoError(t, err)
			assert.Equal(t, tuple, result)
		})

		t.Run("case=non-matching relation returns nil", func(t *testing.T) {
			tuple := &relationtuple.RelationTuple{
				Namespace: nspace,
				Object:    uuid.Must(uuid.NewV4()),
				Relation:  "viewer",
				Subject:   &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
			}
			require.NoError(t, m.WriteRelationTuples(t.Context(), tuple))

			result, err := m.FindTupleWithRelations(t.Context(), tuple, []string{"editor"})
			require.NoError(t, err)
			assert.Nil(t, result)
		})

		t.Run("case=multiple relations with one matching returns tuple", func(t *testing.T) {
			tuple := &relationtuple.RelationTuple{
				Namespace: nspace,
				Object:    uuid.Must(uuid.NewV4()),
				Relation:  "viewer",
				Subject:   &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
			}
			require.NoError(t, m.WriteRelationTuples(t.Context(), tuple))

			result, err := m.FindTupleWithRelations(t.Context(), tuple, []string{"editor", "viewer", "owner"})
			require.NoError(t, err)
			assert.Equal(t, tuple, result)
		})
	})

	t.Run("method=TraverseSubjectSetExpansion", func(t *testing.T) {
		t.Run("case=no subject set tuples returns empty", func(t *testing.T) {
			start := &relationtuple.RelationTuple{
				Namespace: nspace,
				Object:    uuid.Must(uuid.NewV4()),
				Relation:  "rel",
				Subject:   &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
			}

			res, err := m.TraverseSubjectSetExpansion(t.Context(), start, nil)
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
				&relationtuple.RelationTuple{
					Namespace: nspace,
					Object:    obj,
					Relation:  "rel",
					Subject:   &relationtuple.SubjectSet{Namespace: groupNs, Object: groupObj, Relation: "member"},
				},
				// groupNs:groupObj#member@userID
				&relationtuple.RelationTuple{
					Namespace: groupNs,
					Object:    groupObj,
					Relation:  "member",
					Subject:   &relationtuple.SubjectID{ID: userID},
				},
			))

			start := &relationtuple.RelationTuple{
				Namespace: nspace,
				Object:    obj,
				Relation:  "rel",
				Subject:   &relationtuple.SubjectID{ID: userID},
			}
			res, err := m.TraverseSubjectSetExpansion(t.Context(), start, nil)
			require.NoError(t, err)
			require.Len(t, res, 1)
			assert.True(t, res[0].Found)
			assert.Equal(t, relationtuple.TraversalSubjectSetExpand, res[0].Via)
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
				&relationtuple.RelationTuple{
					Namespace: nspace,
					Object:    obj,
					Relation:  "rel",
					Subject:   &relationtuple.SubjectSet{Namespace: groupNs, Object: groupObj, Relation: "member"},
				},
			))

			start := &relationtuple.RelationTuple{
				Namespace: nspace,
				Object:    obj,
				Relation:  "rel",
				Subject:   &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
			}
			res, err := m.TraverseSubjectSetExpansion(t.Context(), start, nil)
			require.NoError(t, err)
			require.Len(t, res, 1)
			assert.False(t, res[0].Found)
		})

		t.Run("case=allowedSubjectSets filters out non-matching types", func(t *testing.T) {
			docNs := strconv.Itoa(rand.Int())   // nolint
			groupNs := strconv.Itoa(rand.Int()) // nolint
			obj := uuid.Must(uuid.NewV4())

			// ns:obj#rel@(docNs:docObj#member) and ns:obj#rel@(groupNs:groupObj#admin)
			require.NoError(t, m.WriteRelationTuples(t.Context(),
				&relationtuple.RelationTuple{
					Namespace: nspace,
					Object:    obj,
					Relation:  "rel",
					Subject:   &relationtuple.SubjectSet{Namespace: docNs, Object: uuid.Must(uuid.NewV4()), Relation: "member"},
				},
				&relationtuple.RelationTuple{
					Namespace: nspace,
					Object:    obj,
					Relation:  "rel",
					Subject:   &relationtuple.SubjectSet{Namespace: groupNs, Object: uuid.Must(uuid.NewV4()), Relation: "admin"},
				},
			))

			start := &relationtuple.RelationTuple{
				Namespace: nspace,
				Object:    obj,
				Relation:  "rel",
				Subject:   &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
			}
			res, err := m.TraverseSubjectSetExpansion(t.Context(), start, []relationtuple.SubjectSetType{{Namespace: docNs, Relation: "member"}})
			require.NoError(t, err)
			require.Len(t, res, 1)
			assert.Equal(t, docNs, res[0].To.Namespace)
			assert.Equal(t, "member", res[0].To.Relation)
		})

		t.Run("case=nil allowedSubjectSets returns all types", func(t *testing.T) {
			docNs := strconv.Itoa(rand.Int())   // nolint
			groupNs := strconv.Itoa(rand.Int()) // nolint
			obj := uuid.Must(uuid.NewV4())

			// ns:obj#rel@(docNs:docObj#member) and ns:obj#rel@(groupNs:groupObj#admin)
			require.NoError(t, m.WriteRelationTuples(t.Context(),
				&relationtuple.RelationTuple{
					Namespace: nspace,
					Object:    obj,
					Relation:  "rel",
					Subject:   &relationtuple.SubjectSet{Namespace: docNs, Object: uuid.Must(uuid.NewV4()), Relation: "member"},
				},
				&relationtuple.RelationTuple{
					Namespace: nspace,
					Object:    obj,
					Relation:  "rel",
					Subject:   &relationtuple.SubjectSet{Namespace: groupNs, Object: uuid.Must(uuid.NewV4()), Relation: "admin"},
				},
			))

			start := &relationtuple.RelationTuple{
				Namespace: nspace,
				Object:    obj,
				Relation:  "rel",
				Subject:   &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
			}
			res, err := m.TraverseSubjectSetExpansion(t.Context(), start, nil)
			require.NoError(t, err)
			assert.Len(t, res, 2)
		})

		t.Run("case=multiple allowedSubjectSets entries each match independently", func(t *testing.T) {
			docNs := strconv.Itoa(rand.Int())   // nolint
			groupNs := strconv.Itoa(rand.Int()) // nolint
			obj := uuid.Must(uuid.NewV4())

			// ns:obj#rel@(docNs:docObj#member), ns:obj#rel@(groupNs:groupObj#admin), and
			// ns:obj#rel@(groupNs:otherObj#viewer) which is not in the allowed list
			require.NoError(t, m.WriteRelationTuples(t.Context(),
				&relationtuple.RelationTuple{
					Namespace: nspace,
					Object:    obj,
					Relation:  "rel",
					Subject:   &relationtuple.SubjectSet{Namespace: docNs, Object: uuid.Must(uuid.NewV4()), Relation: "member"},
				},
				&relationtuple.RelationTuple{
					Namespace: nspace,
					Object:    obj,
					Relation:  "rel",
					Subject:   &relationtuple.SubjectSet{Namespace: groupNs, Object: uuid.Must(uuid.NewV4()), Relation: "admin"},
				},
				&relationtuple.RelationTuple{
					Namespace: nspace,
					Object:    obj,
					Relation:  "rel",
					Subject:   &relationtuple.SubjectSet{Namespace: groupNs, Object: uuid.Must(uuid.NewV4()), Relation: "viewer"},
				},
			))

			start := &relationtuple.RelationTuple{
				Namespace: nspace,
				Object:    obj,
				Relation:  "rel",
				Subject:   &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
			}
			res, err := m.TraverseSubjectSetExpansion(t.Context(), start, []relationtuple.SubjectSetType{
				{Namespace: docNs, Relation: "member"},
				{Namespace: groupNs, Relation: "admin"},
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
