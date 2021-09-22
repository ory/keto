package relationtuple

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/ory/x/pointerx"

	"github.com/ory/herodot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/x"
)

func ManagerTest(t *testing.T, m Manager, addNamespace func(context.Context, *testing.T, string)) {
	t.Run("method=Write", func(t *testing.T) {
		t.Run("case=success", func(t *testing.T) {
			nspace := t.Name()
			addNamespace(context.Background(), t, nspace)

			tuples := []*InternalRelationTuple{
				{
					Namespace: nspace,
					Object:    "obj",
					Relation:  "rel",
					Subject:   &SubjectID{ID: "sub"},
				},
				{
					Namespace: nspace,
					Object:    "obj",
					Relation:  "rel",
					Subject: &SubjectSet{
						Namespace: nspace,
						Object:    "sub obj",
						Relation:  "sub rel",
					},
				},
			}

			require.NoError(t, m.WriteRelationTuples(context.Background(), tuples...))

			for _, tup := range tuples {
				tupC := *tup

				t.Run(fmt.Sprintf("subject_type=%T", tupC.Subject), func(t *testing.T) {
					resp, nextPage, err := m.GetRelationTuples(context.Background(), tupC.ToQuery())
					require.NoError(t, err)
					assert.Equal(t, "", nextPage)
					assert.Equal(t, []*InternalRelationTuple{&tupC}, resp)
				})
			}
		})

		t.Run("case=unknown namespace", func(t *testing.T) {
			err := m.WriteRelationTuples(context.Background(), &InternalRelationTuple{
				Namespace: "unknown namespace",
				Subject:   &SubjectID{},
			})
			assert.NotNil(t, err)
			assert.True(t, errors.Is(err, herodot.ErrNotFound), "actual error: %+v", err)
		})
	})

	t.Run("method=Get", func(t *testing.T) {
		t.Run("case=queries", func(t *testing.T) {
			nspace := t.Name()
			addNamespace(context.Background(), t, nspace)

			tuples := make([]*InternalRelationTuple, 10)
			for i := range tuples {
				tuples[i] = &InternalRelationTuple{
					Namespace: nspace,
					Object:    fmt.Sprintf("o %d", i%2),
					Relation:  fmt.Sprintf("r %d", i%4),
					Subject:   &SubjectID{ID: fmt.Sprintf("s %d", i)},
				}
			}

			require.NoError(t, m.WriteRelationTuples(context.Background(), tuples...))

			for i, tc := range []struct {
				query    *RelationQuery
				expected []*InternalRelationTuple
			}{
				{
					query: &RelationQuery{
						Namespace: nspace,
					},
					expected: tuples,
				},
				{
					query: &RelationQuery{
						Namespace: nspace,
						Object:    "o 0",
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
						Namespace: nspace,
						Relation:  "r 0",
					},
					expected: []*InternalRelationTuple{
						tuples[0],
						tuples[4],
						tuples[8],
					},
				},
				{
					query: &RelationQuery{
						Namespace: nspace,
						Object:    "o 0",
						Relation:  "r 0",
					},
					expected: []*InternalRelationTuple{
						tuples[0],
						tuples[4],
						tuples[8],
					},
				},
				{
					query: &RelationQuery{
						Namespace: nspace,
						SubjectID: pointerx.String("s 0"),
					},
					expected: []*InternalRelationTuple{
						tuples[0],
					},
				},
				{
					query: &RelationQuery{
						Namespace: nspace,
						Object:    "o 0",
						SubjectID: pointerx.String("s 0"),
					},
					expected: []*InternalRelationTuple{
						tuples[0],
					},
				},
				{
					query: &RelationQuery{
						Namespace: nspace,
						Relation:  "r 0",
						SubjectID: pointerx.String("s 0"),
					},
					expected: []*InternalRelationTuple{
						tuples[0],
					},
				},
				{
					query: &RelationQuery{
						Namespace: nspace,
						Object:    "o 0",
						Relation:  "r 0",
						SubjectID: pointerx.String("s 0"),
					},
					expected: []*InternalRelationTuple{
						tuples[0],
					},
				},
			} {
				t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
					res, nextPage, err := m.GetRelationTuples(context.Background(), tc.query)
					require.NoError(t, err)
					assert.Equal(t, "", nextPage)

					// assert equal elements but not equal order
					assert.Equal(t, len(tc.expected), len(res))

					for _, r := range tc.expected {
						assert.Contains(t, res, r)
					}

					for _, r := range res {
						assert.Contains(t, tc.expected, r)
					}
				})
			}
		})

		t.Run("case=pagination", func(t *testing.T) {
			nspace := t.Name()
			addNamespace(context.Background(), t, nspace)

			tuples := make([]*InternalRelationTuple, 20)
			for i := range tuples {
				tuples[i] = &InternalRelationTuple{
					Namespace: nspace,
					Object:    "o",
					Relation:  "r",
					Subject:   &SubjectID{ID: strconv.Itoa(i)},
				}
			}

			require.NoError(t, m.WriteRelationTuples(context.Background(), tuples...))

			notEncounteredTuples := make([]*InternalRelationTuple, len(tuples))
			copy(notEncounteredTuples, tuples)

			var nextPage string
			for range tuples[:len(tuples)-1] {
				var (
					res []*InternalRelationTuple
					err error
				)

				res, nextPage, err = m.GetRelationTuples(context.Background(), &RelationQuery{
					Namespace: nspace,
					Object:    "o",
					Relation:  "r",
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

			res, nextPage, err := m.GetRelationTuples(context.Background(), &RelationQuery{
				Namespace: nspace,
				Object:    "o",
				Relation:  "r",
			}, x.WithSize(1), x.WithToken(nextPage))
			require.NoError(t, err)
			assert.Equal(t, "", nextPage)
			assert.Len(t, res, 1)
			assert.Equal(t, notEncounteredTuples, res)
		})

		t.Run("case=empty list", func(t *testing.T) {
			nspace := t.Name()
			addNamespace(context.Background(), t, nspace)

			res, nextPage, err := m.GetRelationTuples(context.Background(), &RelationQuery{
				Namespace: nspace,
			})

			assert.NoError(t, err)
			assert.Equal(t, []*InternalRelationTuple{}, res)
			assert.Equal(t, "", nextPage)
		})
	})

	t.Run("method=Delete", func(t *testing.T) {
		t.Run("case=deletes tuple", func(t *testing.T) {
			nspace := t.Name()
			addNamespace(context.Background(), t, nspace)

			for _, rt := range []*InternalRelationTuple{
				{
					Namespace: nspace,
					Object:    "o to delete",
					Relation:  "r to delete",
					Subject:   &SubjectID{ID: "s to delete"},
				},
				{
					Namespace: nspace,
					Object:    "o to delete",
					Relation:  "r to delete",
					Subject: &SubjectSet{
						Namespace: nspace,
						Object:    "o2",
						Relation:  "r2",
					},
				},
			} {
				t.Run(fmt.Sprintf("subject_type=%T", rt.Subject), func(t *testing.T) {
					require.NoError(t, m.WriteRelationTuples(context.Background(), rt))

					res, _, err := m.GetRelationTuples(context.Background(), rt.ToQuery())
					require.NoError(t, err)
					assert.Equal(t, []*InternalRelationTuple{rt}, res)

					require.NoError(t, m.DeleteRelationTuples(context.Background(), rt))

					res, _, err = m.GetRelationTuples(context.Background(), rt.ToQuery())
					require.NoError(t, err)
					assert.Len(t, res, 0)
				})
			}
		})

		t.Run("case=deletes only one tuple", func(t *testing.T) {
			nspace := t.Name()
			addNamespace(context.Background(), t, nspace)

			rs := make([]*InternalRelationTuple, 4)
			for i := range rs {
				rs[i] = &InternalRelationTuple{
					Namespace: nspace,
					Object:    "o" + strconv.Itoa(i),
					Relation:  "r" + strconv.Itoa(i),
					Subject:   &SubjectID{ID: "s" + strconv.Itoa(i)},
				}
			}
			require.NoError(t, m.WriteRelationTuples(context.Background(), rs...))

			res, _, err := m.GetRelationTuples(context.Background(), &RelationQuery{
				Namespace: nspace,
			})
			require.NoError(t, err)
			for _, rt := range rs {
				assert.Contains(t, res, rt)
			}

			require.NoError(t, m.DeleteRelationTuples(context.Background(), rs[0], rs[2]))

			res, _, err = m.GetRelationTuples(context.Background(), &RelationQuery{
				Namespace: nspace,
			})
			require.NoError(t, err)
			assert.Equal(t, []*InternalRelationTuple{rs[1], rs[3]}, res)
		})

		t.Run("case=tuple and subject namespace differ", func(t *testing.T) {
			ctx := context.Background()

			n0, n1 := t.Name()+"0", t.Name()+"1"
			addNamespace(ctx, t, n0)
			addNamespace(ctx, t, n1)

			rt := &InternalRelationTuple{
				Namespace: n0,
				Object:    "o",
				Relation:  "r",
				Subject: &SubjectSet{
					Namespace: n1,
					Object:    "o",
					Relation:  "r",
				},
			}
			require.NoError(t, m.WriteRelationTuples(ctx, rt))

			actual, _, err := m.GetRelationTuples(ctx, &RelationQuery{Namespace: n0})
			require.NoError(t, err)
			assert.Equal(t, []*InternalRelationTuple{rt}, actual)

			require.NoError(t, m.DeleteRelationTuples(ctx, rt))

			actual, _, err = m.GetRelationTuples(ctx, &RelationQuery{Namespace: n0})
			require.NoError(t, err)
			assert.Len(t, actual, 0)
		})
	})

	t.Run("method=Transact", func(t *testing.T) {
		t.Run("case=success", func(t *testing.T) {
			nspace := t.Name()
			addNamespace(context.Background(), t, nspace)

			rs := make([]*InternalRelationTuple, 4)
			for i := range rs {
				rs[i] = &InternalRelationTuple{
					Namespace: nspace,
					Object:    "o" + strconv.Itoa(i),
					Relation:  "r" + strconv.Itoa(i),
					Subject:   &SubjectID{ID: "s" + strconv.Itoa(i)},
				}
			}
			require.NoError(t, m.WriteRelationTuples(context.Background(), rs[0], rs[1]))

			res, _, err := m.GetRelationTuples(context.Background(), &RelationQuery{
				Namespace: nspace,
			})
			require.NoError(t, err)
			assert.Equal(t, []*InternalRelationTuple{rs[0], rs[1]}, res)

			require.NoError(t, m.TransactRelationTuples(context.Background(), []*InternalRelationTuple{rs[2], rs[3]}, []*InternalRelationTuple{rs[0]}))

			res, _, err = m.GetRelationTuples(context.Background(), &RelationQuery{
				Namespace: nspace,
			})
			require.NoError(t, err)

			for _, rt := range []*InternalRelationTuple{rs[1], rs[2], rs[3]} {
				assert.Contains(t, res, rt)
			}
		})

		t.Run("case=err rolls back all", func(t *testing.T) {
			nspace := t.Name()
			addNamespace(context.Background(), t, nspace)

			rs := make([]*InternalRelationTuple, 2)
			for i := range rs {
				rs[i] = &InternalRelationTuple{
					Namespace: nspace,
					Object:    "o" + strconv.Itoa(i),
					Relation:  "r" + strconv.Itoa(i),
					Subject:   &SubjectID{ID: "s" + strconv.Itoa(i)},
				}
			}
			invalidRt := &InternalRelationTuple{
				Namespace: nspace,
				Object:    "o0",
				Relation:  "r0",
				Subject:   nil, // subject is not allowed to be nil
			}
			require.NoError(t, m.WriteRelationTuples(context.Background(), rs[0]))

			res, _, err := m.GetRelationTuples(context.Background(), &RelationQuery{
				Namespace: nspace,
			})
			require.NoError(t, err)
			assert.Equal(t, []*InternalRelationTuple{rs[0]}, res)

			t.Run("invalid=insert", func(t *testing.T) {
				assert.ErrorIs(t, m.TransactRelationTuples(context.Background(), []*InternalRelationTuple{invalidRt}, []*InternalRelationTuple{rs[0]}), ErrNilSubject)
			})

			res, _, err = m.GetRelationTuples(context.Background(), &RelationQuery{
				Namespace: nspace,
			})
			require.NoError(t, err)
			assert.Equal(t, []*InternalRelationTuple{rs[0]}, res)

			t.Run("invalid=delete", func(t *testing.T) {
				assert.ErrorIs(t, m.TransactRelationTuples(context.Background(), []*InternalRelationTuple{rs[1]}, []*InternalRelationTuple{invalidRt}), ErrNilSubject)
			})

			res, _, err = m.GetRelationTuples(context.Background(), &RelationQuery{
				Namespace: nspace,
			})
			require.NoError(t, err)
			assert.Equal(t, []*InternalRelationTuple{rs[0]}, res)
		})
	})
}
