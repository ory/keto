// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/gofrs/uuid"
	keysetpagination "github.com/ory/x/pagination/keysetpagination_v2"
	"github.com/ory/x/pointerx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoapi"
)

func ManagerTest(t *testing.T, m Manager) {
	ctx := context.Background()

	t.Run("method=Write", func(t *testing.T) {
		t.Run("case=success", func(t *testing.T) {
			nspace := strconv.Itoa(rand.Int()) // nolint

			tuples := []*RelationTuple{
				{
					Namespace: nspace,
					Object:    uuid.Must(uuid.NewV4()),
					Relation:  "rel",
					Subject:   &SubjectID{ID: uuid.Must(uuid.NewV4())},
				},
				{
					Namespace: nspace,
					Object:    uuid.Must(uuid.NewV4()),
					Relation:  "rel",
					Subject: &SubjectSet{
						Namespace: nspace,
						Object:    uuid.Must(uuid.NewV4()),
						Relation:  "sub rel",
					},
				},
			}

			require.NoError(t, m.WriteRelationTuples(ctx, tuples...))

			resp, nextPage, err := m.GetRelationTuples(ctx, &RelationQuery{
				Namespace: pointerx.Ptr(nspace),
			})
			require.NoError(t, err)
			assert.True(t, nextPage.IsLast())
			assert.ElementsMatch(t, tuples, resp)
		})
	})

	t.Run("method=Get", func(t *testing.T) {
		t.Run("case=queries", func(t *testing.T) {
			nspace := strconv.Itoa(rand.Int()) // nolint

			tuples := make([]*RelationTuple, 10)
			ids := x.UUIDs(len(tuples))

			for i := range tuples {
				tuples[i] = &RelationTuple{
					Namespace: nspace,
					Object:    ids[i%2],
					Relation:  fmt.Sprintf("r %d", i%4),
					Subject:   &SubjectID{ID: ids[i]},
				}
			}

			require.NoError(t, m.WriteRelationTuples(ctx, tuples...))

			for i, tc := range []struct {
				query    *RelationQuery
				expected []*RelationTuple
			}{
				{
					query: &RelationQuery{
						Namespace: &nspace,
					},
					expected: tuples,
				},
				{
					query: &RelationQuery{
						Namespace: &nspace,
						Object:    &ids[0],
					},
					expected: []*RelationTuple{
						tuples[0],
						tuples[2],
						tuples[4],
						tuples[6],
						tuples[8],
					},
				},
				{
					query: &RelationQuery{
						Namespace: &nspace,
						Relation:  pointerx.Ptr("r 0"),
					},
					expected: []*RelationTuple{
						tuples[0],
						tuples[4],
						tuples[8],
					},
				},
				{
					query: &RelationQuery{
						Namespace: &nspace,
						Object:    &ids[0],
						Relation:  pointerx.Ptr("r 0"),
					},
					expected: []*RelationTuple{
						tuples[0],
						tuples[4],
						tuples[8],
					},
				},
				{
					query: &RelationQuery{
						Namespace: &nspace,
						Subject:   &SubjectID{ids[0]},
					},
					expected: []*RelationTuple{
						tuples[0],
					},
				},
				{
					query: &RelationQuery{
						Namespace: &nspace,
						Object:    &ids[0],
						Subject:   &SubjectID{ids[0]},
					},
					expected: []*RelationTuple{
						tuples[0],
					},
				},
				{
					query: &RelationQuery{
						Namespace: &nspace,
						Relation:  pointerx.Ptr("r 0"),
						Subject:   &SubjectID{ids[0]},
					},
					expected: []*RelationTuple{
						tuples[0],
					},
				},
				{
					query: &RelationQuery{
						Namespace: &nspace,
						Object:    &ids[0],
						Relation:  pointerx.Ptr("r 0"),
						Subject:   &SubjectID{ids[0]},
					},
					expected: []*RelationTuple{
						tuples[0],
					},
				},
			} {
				t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
					res, nextPage, err := m.GetRelationTuples(ctx, tc.query)
					require.NoError(t, err)
					assert.True(t, nextPage.IsLast())
					assert.ElementsMatch(t, tc.expected, res)
				})
			}
		})

		t.Run("case=pagination", func(t *testing.T) {
			nspace := strconv.Itoa(rand.Int()) // nolint

			tuples := make([]*RelationTuple, 20)
			oID := uuid.Must(uuid.NewV4())
			for i := range tuples {
				tuples[i] = &RelationTuple{
					Namespace: nspace,
					Object:    oID,
					Relation:  "r",
					Subject:   &SubjectID{uuid.Must(uuid.NewV4())},
				}
			}

			require.NoError(t, m.WriteRelationTuples(ctx, tuples...))

			notEncounteredTuples := make([]*RelationTuple, len(tuples))
			copy(notEncounteredTuples, tuples)

			var (
				res, thisPage []*RelationTuple
				err           error
			)
			nextPage := keysetpagination.NewPaginator(keysetpagination.WithSize(1))
			for range tuples[:len(tuples)-1] {
				thisPage, nextPage, err = m.GetRelationTuples(ctx, &RelationQuery{
					Namespace: pointerx.Ptr(nspace),
					Object:    &oID,
					Relation:  pointerx.Ptr("r"),
				}, nextPage.ToOptions()...)
				require.NoError(t, err)
				assert.False(t, nextPage.IsLast())
				require.Len(t, thisPage, 1)

				res = append(res, thisPage[0])
			}

			thisPage, nextPage, err = m.GetRelationTuples(ctx, &RelationQuery{
				Namespace: pointerx.Ptr(nspace),
				Object:    &oID,
				Relation:  pointerx.Ptr("r"),
			}, nextPage.ToOptions()...)
			require.NoError(t, err)
			assert.True(t, nextPage.IsLast())
			require.Len(t, thisPage, 1)

			res = append(res, thisPage[0])
			assert.ElementsMatch(t, tuples, res)
		})

		t.Run("case=empty list", func(t *testing.T) {
			nspace := t.Name()

			res, nextPage, err := m.GetRelationTuples(ctx, &RelationQuery{
				Namespace: &nspace,
			})

			assert.NoError(t, err)
			assert.Equal(t, []*RelationTuple{}, res)
			assert.True(t, nextPage.IsLast())
		})
	})

	t.Run("method=Delete", func(t *testing.T) {
		t.Run("case=deletes tuple", func(t *testing.T) {
			nspace := strconv.Itoa(rand.Int()) // nolint
			oID := uuid.Must(uuid.NewV4())
			sID := uuid.Must(uuid.NewV4())

			for _, rt := range []*RelationTuple{
				{
					Namespace: nspace,
					Object:    oID,
					Relation:  "r to delete",
					Subject:   &SubjectID{ID: sID},
				},
				{
					Namespace: nspace,
					Object:    oID,
					Relation:  "r to delete",
					Subject: &SubjectSet{
						Namespace: nspace,
						Object:    uuid.Must(uuid.NewV4()),
						Relation:  "r2",
					},
				},
			} {
				t.Run(fmt.Sprintf("subject_type=%T", rt.Subject), func(t *testing.T) {
					require.NoError(t, m.WriteRelationTuples(ctx, rt))

					res, _, err := m.GetRelationTuples(ctx, &RelationQuery{
						Namespace: pointerx.Ptr(nspace),
					})
					require.NoError(t, err)
					assert.Equal(t, []*RelationTuple{rt}, res)

					require.NoError(t, m.DeleteRelationTuples(ctx, rt))

					res, _, err = m.GetRelationTuples(ctx, &RelationQuery{
						Namespace: pointerx.Ptr(nspace),
					})
					require.NoError(t, err)
					assert.Len(t, res, 0)
				})
			}
		})

		t.Run("case=deletes only one tuple", func(t *testing.T) {
			nspace := strconv.Itoa(rand.Int()) // nolint

			rs := make([]*RelationTuple, 4)
			oIDs, sIDs := make([]uuid.UUID, len(rs)), make([]uuid.UUID, len(rs))
			for i := range oIDs {
				oIDs[i] = uuid.Must(uuid.NewV4())
				sIDs[i] = uuid.Must(uuid.NewV4())
			}

			for i := range rs {
				rs[i] = &RelationTuple{
					Namespace: nspace,
					Object:    oIDs[i],
					Relation:  "r" + strconv.Itoa(i),
					Subject:   &SubjectID{ID: sIDs[i]},
				}
			}
			require.NoError(t, m.WriteRelationTuples(ctx, rs...))

			res, _, err := m.GetRelationTuples(ctx, &RelationQuery{
				Namespace: &nspace,
			})
			require.NoError(t, err)
			for _, rt := range rs {
				assert.Contains(t, res, rt)
			}

			require.NoError(t, m.DeleteRelationTuples(ctx, rs[0], rs[2]))

			res, _, err = m.GetRelationTuples(ctx, &RelationQuery{
				Namespace: &nspace,
			})
			require.NoError(t, err)
			assert.ElementsMatch(t, []*RelationTuple{rs[1], rs[3]}, res)
		})

		t.Run("case=tuple and subject namespace differ", func(t *testing.T) {
			ctx := ctx

			n0, n1 := t.Name()[:60], t.Name()[:60]+"1"
			oID := uuid.Must(uuid.NewV4())

			rt := &RelationTuple{
				Namespace: n0,
				Object:    oID,
				Relation:  "r",
				Subject: &SubjectSet{
					Namespace: n1,
					Object:    oID,
					Relation:  "r",
				},
			}
			require.NoError(t, m.WriteRelationTuples(ctx, rt))

			actual, _, err := m.GetRelationTuples(ctx, &RelationQuery{Namespace: &n0})
			require.NoError(t, err)
			assert.Equal(t, []*RelationTuple{rt}, actual)

			require.NoError(t, m.DeleteRelationTuples(ctx, rt))

			actual, _, err = m.GetRelationTuples(ctx, &RelationQuery{Namespace: &n0})
			require.NoError(t, err)
			assert.Len(t, actual, 0)
		})
	})

	t.Run("method=Transact", func(t *testing.T) {
		t.Run("case=success", func(t *testing.T) {
			nspace := strconv.Itoa(rand.Int()) // nolint

			rs := make([]*RelationTuple, 4)
			oIDs, sIDs := make([]uuid.UUID, len(rs)), make([]uuid.UUID, len(rs))
			for i := range oIDs {
				oIDs[i] = uuid.Must(uuid.NewV4())
				sIDs[i] = uuid.Must(uuid.NewV4())
			}

			for i := range rs {
				rs[i] = &RelationTuple{
					Namespace: nspace,
					Object:    oIDs[i],
					Relation:  "r" + strconv.Itoa(i),
					Subject:   &SubjectID{ID: sIDs[i]},
				}
			}
			require.NoError(t, m.WriteRelationTuples(ctx, rs[0], rs[1]))

			res, _, err := m.GetRelationTuples(ctx, &RelationQuery{
				Namespace: &nspace,
			})
			require.NoError(t, err)
			assert.ElementsMatch(t, []*RelationTuple{rs[0], rs[1]}, res)

			require.NoError(t, m.TransactRelationTuples(ctx, []*RelationTuple{rs[2], rs[3]}, []*RelationTuple{rs[0]}))

			res, _, err = m.GetRelationTuples(ctx, &RelationQuery{
				Namespace: &nspace,
			})
			require.NoError(t, err)

			for _, rt := range []*RelationTuple{rs[1], rs[2], rs[3]} {
				assert.Contains(t, res, rt)
			}
		})

		t.Run("case=err rolls back all", func(t *testing.T) {
			nspace := strconv.Itoa(rand.Int()) // nolint

			rs := make([]*RelationTuple, 2)
			oIDs, sIDs := make([]uuid.UUID, len(rs)), make([]uuid.UUID, len(rs))
			for i := range oIDs {
				oIDs[i] = uuid.Must(uuid.NewV4())
				sIDs[i] = uuid.Must(uuid.NewV4())
			}
			for i := range rs {
				rs[i] = &RelationTuple{
					Namespace: nspace,
					Object:    oIDs[i],
					Relation:  "r" + strconv.Itoa(i),
					Subject:   &SubjectID{ID: sIDs[i]},
				}
			}
			invalidRt := &RelationTuple{
				Namespace: nspace,
				Object:    oIDs[0],
				Relation:  "r0",
				Subject:   nil, // subject is not allowed to be nil
			}
			require.NoError(t, m.WriteRelationTuples(ctx, rs[0]))

			res, _, err := m.GetRelationTuples(ctx, &RelationQuery{
				Namespace: &nspace,
			})
			require.NoError(t, err)
			assert.Equal(t, []*RelationTuple{rs[0]}, res)

			t.Run("invalid=insert", func(t *testing.T) {
				assert.ErrorIs(t, m.TransactRelationTuples(ctx, []*RelationTuple{invalidRt}, []*RelationTuple{rs[0]}), ketoapi.ErrNilSubject)
			})

			res, _, err = m.GetRelationTuples(ctx, &RelationQuery{
				Namespace: &nspace,
			})
			require.NoError(t, err)
			assert.Equal(t, []*RelationTuple{rs[0]}, res)

			t.Run("invalid=delete", func(t *testing.T) {
				assert.ErrorIs(t, m.TransactRelationTuples(ctx, []*RelationTuple{rs[1]}, []*RelationTuple{invalidRt}), ketoapi.ErrNilSubject)
			})

			res, _, err = m.GetRelationTuples(ctx, &RelationQuery{
				Namespace: &nspace,
			})
			require.NoError(t, err)
			assert.Equal(t, []*RelationTuple{rs[0]}, res)
		})
	})
}
