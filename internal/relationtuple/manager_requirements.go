package relationtuple

import (
	"context"
	"errors"
	"testing"

	"github.com/ory/herodot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/x"
)

func ManagerTest(t *testing.T, m Manager, nspace string) {
	t.Run("method=Write", func(t *testing.T) {
		t.Run("case=success", func(t *testing.T) {
			tuples := []*InternalRelationTuple{
				{
					Namespace: nspace,
					Object:    "obj for " + t.Name(),
					Relation:  "rel for " + t.Name(),
					Subject:   &SubjectID{ID: "sub for " + t.Name()},
				},
				{
					Namespace: nspace,
					Object:    "obj for " + t.Name(),
					Relation:  "rel for " + t.Name(),
					Subject: &SubjectSet{
						Namespace: nspace,
						Object:    "sub obj for " + t.Name(),
						Relation:  "sub rel for " + t.Name(),
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
				Namespace: "not " + nspace,
			})
			assert.NotNil(t, err)
			assert.True(t, errors.Is(err, herodot.ErrNotFound))
		})
	})

	//t.Run("method=Get", func(t *testing.T) {
	//	for i, tc :=
	//		m.GetRelationTuples(context.Background())
	//})
}
