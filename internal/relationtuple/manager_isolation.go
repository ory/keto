package relationtuple

import (
	"context"
	"github.com/gofrs/uuid"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func twice(t *testing.T, m0, m1 Manager) func(string, func(*testing.T, Manager, Manager)) {
	return func(desc string, run func(*testing.T, Manager, Manager)) {
		t.Run(desc+"-0", func(t *testing.T) {
			run(t, m0, m1)
		})
		t.Run(desc+"-1", func(t *testing.T) {
			run(t, m1, m0)
		})
	}
}

func reset(t *testing.T, ms ...Manager) {
	ctx := context.Background()

	for _, m := range ms {
		for {
			rts, next, err := m.GetRelationTuples(ctx, &RelationQuery{})
			require.NoError(t, err)

			require.NoError(t, m.DeleteRelationTuples(ctx, rts...))

			if next == "" {
				break
			}
		}
	}
}

func IsolationTest(t *testing.T, m0, m1 Manager) {
	ctx := context.Background()
	run := twice(t, m0, m1)

	run("suite=lifecycle", func(t *testing.T, m0, m1 Manager) {
		nspace := rand.Int31()

		rts := []*InternalRelationTuple{
			{
				Namespace: nspace,
				Object:    uuid.Must(uuid.NewV4()),
				Relation:  "r",
				Subject: &SubjectSet{
					Namespace: nspace,
					Object:    uuid.Must(uuid.NewV4()),
					Relation:  "r",
				},
			},
			{
				Namespace: nspace,
				Object:    uuid.Must(uuid.NewV4()),
				Relation:  "r",
				Subject:   &SubjectID{ID: uuid.Must(uuid.NewV4())},
			},
		}

		t.Run("case=write and get", func(t *testing.T) {
			reset(t, m0, m1)

			require.NoError(t, m0.WriteRelationTuples(ctx, rts...))

			other, _, err := m1.GetRelationTuples(ctx, &RelationQuery{Namespace: &nspace})
			require.NoError(t, err)
			assert.Len(t, other, 0)

			actual, _, err := m0.GetRelationTuples(ctx, &RelationQuery{Namespace: &nspace})
			require.NoError(t, err)
			assert.Equal(t, rts, actual)
		})

		t.Run("case=delete", func(t *testing.T) {
			reset(t, m0, m1)

			require.NoError(t, m1.WriteRelationTuples(ctx, rts...))

			require.NoError(t, m0.DeleteRelationTuples(ctx, rts...))

			deleted, _, err := m0.GetRelationTuples(ctx, &RelationQuery{Namespace: &nspace})
			require.NoError(t, err)
			assert.Len(t, deleted, 0)

			actual, _, err := m1.GetRelationTuples(ctx, &RelationQuery{Namespace: &nspace})
			require.NoError(t, err)
			assert.Equal(t, rts, actual)
		})

		t.Run("case=transact", func(t *testing.T) {
			reset(t, m0, m1)

			// note that the reset is outside this subtest, so in the second run we actually delete what we had before
			twice(t, m0, m1)("insert and delete", func(t *testing.T, m0, m1 Manager) {
				require.NoError(t, m0.TransactRelationTuples(ctx, []*InternalRelationTuple{rts[0]}, []*InternalRelationTuple{rts[1]}))

				require.NoError(t, m1.TransactRelationTuples(ctx, []*InternalRelationTuple{rts[1]}, []*InternalRelationTuple{rts[0]}))

				r0, _, err := m0.GetRelationTuples(ctx, &RelationQuery{Namespace: &nspace})
				require.NoError(t, err)

				r1, _, err := m1.GetRelationTuples(ctx, &RelationQuery{Namespace: &nspace})
				require.NoError(t, err)

				assert.Equal(t, rts[:1], r0)
				assert.Equal(t, rts[1:], r1)
			})
		})
	})
}
