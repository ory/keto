package relationtuple

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"testing"

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

				resp, nextPage, err := m.GetRelationTuples(context.Background(), (*RelationQuery)(&tupC))
				require.NoError(t, err)
				assert.Equal(t, x.PageTokenEnd, nextPage)
				assert.Equal(t, []*InternalRelationTuple{&tupC}, resp)
			}
		})

		t.Run("case=unknown namespace", func(t *testing.T) {
			err := m.WriteRelationTuples(context.Background(), &InternalRelationTuple{
				Namespace: "unknown namespace",
			})
			assert.NotNil(t, err)
			assert.True(t, errors.Is(err, herodot.ErrNotFound))
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
					Subject:   &SubjectID{ID: fmt.Sprintf("s %d", i%8)},
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
						Subject:   &SubjectID{ID: "s 0"},
					},
					expected: []*InternalRelationTuple{
						tuples[0],
						tuples[8],
					},
				},
				{
					query: &RelationQuery{
						Namespace: nspace,
						Object:    "o 0",
						Subject:   &SubjectID{ID: "s 0"},
					},
					expected: []*InternalRelationTuple{
						tuples[0],
						tuples[8],
					},
				},
				{
					query: &RelationQuery{
						Namespace: nspace,
						Relation:  "r 0",
						Subject:   &SubjectID{ID: "s 0"},
					},
					expected: []*InternalRelationTuple{
						tuples[0],
						tuples[8],
					},
				},
				{
					query: &RelationQuery{
						Namespace: nspace,
						Object:    "o 0",
						Relation:  "r 0",
						Subject:   &SubjectID{ID: "s 0"},
					},
					expected: []*InternalRelationTuple{
						tuples[0],
						tuples[8],
					},
				},
			} {
				t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
					res, nextPage, err := m.GetRelationTuples(context.Background(), tc.query)
					require.NoError(t, err)
					assert.Equal(t, x.PageTokenEnd, nextPage)

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
				assert.NotEqual(t, x.PageTokenEnd, nextPage)
				assert.Len(t, res, 1)

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
			assert.Equal(t, x.PageTokenEnd, nextPage)
			assert.Len(t, res, 1)
			assert.Equal(t, notEncounteredTuples, res)
		})
	})
}
