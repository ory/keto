package relationtuple

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/x"
)

func ManagerTest(t *testing.T, m Manager) {
	ctx := context.Background()

	t.Run("method=Write", func(t *testing.T) {
		t.Run("case=success", func(t *testing.T) {
			nspace := rand.Int31()

			tuples := []*InternalRelationTuple{
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
				Namespace: x.Ptr(nspace),
			})
			require.NoError(t, err)
			assert.Equal(t, "", nextPage)
			assert.ElementsMatch(t, tuples, resp)
		})
	})

	t.Run("method=Get", func(t *testing.T) {
		t.Run("case=queries", func(t *testing.T) {
			nspace := rand.Int31()

			tuples := make([]*InternalRelationTuple, 10)
			ids := make([]uuid.UUID, len(tuples))
			for i := range ids {
				ids[i] = uuid.Must(uuid.NewV4())
			}

			for i := range tuples {
				tuples[i] = &InternalRelationTuple{
					Namespace: nspace,
					Object:    ids[i%2],
					Relation:  fmt.Sprintf("r %d", i%4),
					Subject:   &SubjectID{ID: ids[i]},
				}
			}

			require.NoError(t, m.WriteRelationTuples(ctx, tuples...))

			for i, tc := range []struct {
				query    *RelationQuery
				expected []*InternalRelationTuple
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
					expected: []*InternalRelationTuple{
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
						Relation:  x.Ptr("r 0"),
					},
					expected: []*InternalRelationTuple{
						tuples[0],
						tuples[4],
						tuples[8],
					},
				},
				{
					query: &RelationQuery{
						Namespace: &nspace,
						Object:    &ids[0],
						Relation:  x.Ptr("r 0"),
					},
					expected: []*InternalRelationTuple{
						tuples[0],
						tuples[4],
						tuples[8],
					},
				},
				{
					query: &RelationQuery{
						Namespace: &nspace,
						SubjectID: &ids[0],
					},
					expected: []*InternalRelationTuple{
						tuples[0],
					},
				},
				{
					query: &RelationQuery{
						Namespace: &nspace,
						Object:    &ids[0],
						SubjectID: &ids[0],
					},
					expected: []*InternalRelationTuple{
						tuples[0],
					},
				},
				{
					query: &RelationQuery{
						Namespace: &nspace,
						Relation:  x.Ptr("r 0"),
						SubjectID: &ids[0],
					},
					expected: []*InternalRelationTuple{
						tuples[0],
					},
				},
				{
					query: &RelationQuery{
						Namespace: &nspace,
						Object:    &ids[0],
						Relation:  x.Ptr("r 0"),
						SubjectID: &ids[0],
					},
					expected: []*InternalRelationTuple{
						tuples[0],
					},
				},
			} {
				t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
					res, nextPage, err := m.GetRelationTuples(ctx, tc.query)
					require.NoError(t, err)
					assert.Equal(t, "", nextPage)
					assert.ElementsMatch(t, tc.expected, res)
				})
			}
		})

		t.Run("case=pagination", func(t *testing.T) {
			nspace := rand.Int31()

			tuples := make([]*InternalRelationTuple, 20)
			oID := uuid.Must(uuid.NewV4())
			for i := range tuples {
				tuples[i] = &InternalRelationTuple{
					Namespace: nspace,
					Object:    oID,
					Relation:  "r",
					Subject:   &SubjectID{uuid.Must(uuid.NewV4())},
				}
			}

			require.NoError(t, m.WriteRelationTuples(ctx, tuples...))

			notEncounteredTuples := make([]*InternalRelationTuple, len(tuples))
			copy(notEncounteredTuples, tuples)

			var nextPage string
			for range tuples[:len(tuples)-1] {
				var (
					res []*InternalRelationTuple
					err error
				)

				res, nextPage, err = m.GetRelationTuples(ctx, &RelationQuery{
					Namespace: x.Ptr(nspace),
					Object:    &oID,
					Relation:  x.Ptr("r"),
				}, x.WithSize(1), x.WithToken(nextPage))
				require.NoError(t, err)
				assert.NotEqual(t, "", nextPage)
				require.Len(t, res, 1)

				var found bool
				for i, r := range notEncounteredTuples {
					if assert.ObjectsAreEqual(r, res[0]) {
						found = true
						notEncounteredTuples[i] = notEncounteredTuples[len(notEncounteredTuples)-1]
						notEncounteredTuples = notEncounteredTuples[:len(notEncounteredTuples)-1]
						break
					}
				}
				assert.True(t, found, "not encountered: %+v, res: %+v", notEncounteredTuples, res[0])
			}

			res, nextPage, err := m.GetRelationTuples(ctx, &RelationQuery{
				Namespace: x.Ptr(nspace),
				Object:    &oID,
				Relation:  x.Ptr("r"),
			}, x.WithSize(1), x.WithToken(nextPage))
			require.NoError(t, err)
			assert.Equal(t, "", nextPage)
			assert.Len(t, res, 1)
			assert.Equal(t, notEncounteredTuples, res)
		})

		t.Run("case=empty list", func(t *testing.T) {
			nspace := rand.Int31()

			res, nextPage, err := m.GetRelationTuples(ctx, &RelationQuery{
				Namespace: &nspace,
			})

			assert.NoError(t, err)
			assert.Equal(t, []*InternalRelationTuple{}, res)
			assert.Equal(t, "", nextPage)
		})
	})

	t.Run("method=Delete", func(t *testing.T) {
		t.Run("case=deletes tuple", func(t *testing.T) {
			nspace := rand.Int31()
			oID := uuid.Must(uuid.NewV4())
			sID := uuid.Must(uuid.NewV4())

			for _, rt := range []*InternalRelationTuple{
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
						Namespace: x.Ptr(nspace),
					})
					require.NoError(t, err)
					assert.Equal(t, []*InternalRelationTuple{rt}, res)

					require.NoError(t, m.DeleteRelationTuples(ctx, rt))

					res, _, err = m.GetRelationTuples(ctx, &RelationQuery{
						Namespace: x.Ptr(nspace),
					})
					require.NoError(t, err)
					assert.Len(t, res, 0)
				})
			}
		})

		t.Run("case=deletes only one tuple", func(t *testing.T) {
			nspace := rand.Int31()

			rs := make([]*InternalRelationTuple, 4)
			oIDs, sIDs := make([]uuid.UUID, len(rs)), make([]uuid.UUID, len(rs))
			for i := range oIDs {
				oIDs[i] = uuid.Must(uuid.NewV4())
				sIDs[i] = uuid.Must(uuid.NewV4())
			}

			for i := range rs {
				rs[i] = &InternalRelationTuple{
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
			assert.Equal(t, []*InternalRelationTuple{rs[1], rs[3]}, res)
		})

		t.Run("case=tuple and subject namespace differ", func(t *testing.T) {
			ctx := ctx

			n0, n1 := rand.Int31(), rand.Int31()
			oID := uuid.Must(uuid.NewV4())

			rt := &InternalRelationTuple{
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
			assert.Equal(t, []*InternalRelationTuple{rt}, actual)

			require.NoError(t, m.DeleteRelationTuples(ctx, rt))

			actual, _, err = m.GetRelationTuples(ctx, &RelationQuery{Namespace: &n0})
			require.NoError(t, err)
			assert.Len(t, actual, 0)
		})
	})

	t.Run("method=Transact", func(t *testing.T) {
		t.Run("case=success", func(t *testing.T) {
			nspace := rand.Int31()

			rs := make([]*InternalRelationTuple, 4)
			oIDs, sIDs := make([]uuid.UUID, len(rs)), make([]uuid.UUID, len(rs))
			for i := range oIDs {
				oIDs[i] = uuid.Must(uuid.NewV4())
				sIDs[i] = uuid.Must(uuid.NewV4())
			}

			for i := range rs {
				rs[i] = &InternalRelationTuple{
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
			assert.Equal(t, []*InternalRelationTuple{rs[0], rs[1]}, res)

			require.NoError(t, m.TransactRelationTuples(ctx, []*InternalRelationTuple{rs[2], rs[3]}, []*InternalRelationTuple{rs[0]}))

			res, _, err = m.GetRelationTuples(ctx, &RelationQuery{
				Namespace: &nspace,
			})
			require.NoError(t, err)

			for _, rt := range []*InternalRelationTuple{rs[1], rs[2], rs[3]} {
				assert.Contains(t, res, rt)
			}
		})

		t.Run("case=err rolls back all", func(t *testing.T) {
			nspace := rand.Int31()

			rs := make([]*InternalRelationTuple, 2)
			oIDs, sIDs := make([]uuid.UUID, len(rs)), make([]uuid.UUID, len(rs))
			for i := range oIDs {
				oIDs[i] = uuid.Must(uuid.NewV4())
				sIDs[i] = uuid.Must(uuid.NewV4())
			}
			for i := range rs {
				rs[i] = &InternalRelationTuple{
					Namespace: nspace,
					Object:    oIDs[i],
					Relation:  "r" + strconv.Itoa(i),
					Subject:   &SubjectID{ID: sIDs[i]},
				}
			}
			invalidRt := &InternalRelationTuple{
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
			assert.Equal(t, []*InternalRelationTuple{rs[0]}, res)

			t.Run("invalid=insert", func(t *testing.T) {
				assert.ErrorIs(t, m.TransactRelationTuples(ctx, []*InternalRelationTuple{invalidRt}, []*InternalRelationTuple{rs[0]}), ErrNilSubject)
			})

			res, _, err = m.GetRelationTuples(ctx, &RelationQuery{
				Namespace: &nspace,
			})
			require.NoError(t, err)
			assert.Equal(t, []*InternalRelationTuple{rs[0]}, res)

			t.Run("invalid=delete", func(t *testing.T) {
				assert.ErrorIs(t, m.TransactRelationTuples(ctx, []*InternalRelationTuple{rs[1]}, []*InternalRelationTuple{invalidRt}), ErrNilSubject)
			})

			res, _, err = m.GetRelationTuples(ctx, &RelationQuery{
				Namespace: &nspace,
			})
			require.NoError(t, err)
			assert.Equal(t, []*InternalRelationTuple{rs[0]}, res)
		})
	})
}
