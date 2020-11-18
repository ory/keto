package memory

import (
	"context"
	"fmt"
	"testing"

	"github.com/ory/keto/internal/relationtuple"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestSetup(t *testing.T) (p *Persister, rel1, rel2, rel3, rel4 *relationtuple.InternalRelationTuple) {
	p = NewPersister()
	rel1 = &relationtuple.InternalRelationTuple{
		Object:   &relationtuple.Object{ID: "obj", Namespace: "rel1"},
		Relation: "rel1 name",
		Subject:  &relationtuple.UserID{ID: "rel1 user"},
	}
	rel2 = &relationtuple.InternalRelationTuple{
		Object:   &relationtuple.Object{ID: "obj", Namespace: "rel2"},
		Relation: "rel2 name",
		Subject:  &relationtuple.UserID{ID: "rel2 user"},
	}
	rel3 = &relationtuple.InternalRelationTuple{
		Object:   &relationtuple.Object{ID: "obj", Namespace: "rel3"},
		Relation: "shared name",
		Subject: &relationtuple.UserSet{
			Object:   &relationtuple.Object{ID: "user set obj", Namespace: "rel3"},
			Relation: "rel3 user set",
		},
	}
	rel4 = &relationtuple.InternalRelationTuple{
		Object:   &relationtuple.Object{ID: "obj", Namespace: "rel4"},
		Relation: "shared name",
		Subject:  &relationtuple.UserID{ID: "rel4 user"},
	}

	require.NoError(t, p.WriteRelationTuples(context.Background(), rel1, rel2, rel3, rel4))
	return
}

func TestGetRelationTuples(t *testing.T) {
	p, rel1, rel2, rel3, rel4 := newTestSetup(t)

	for i, tc := range []struct {
		query    *relationtuple.RelationQuery
		expected []*relationtuple.InternalRelationTuple
	}{
		{
			query:    &relationtuple.RelationQuery{ObjectID: rel1.Object},
			expected: []*relationtuple.InternalRelationTuple{rel1},
		},
		{
			query:    &relationtuple.RelationQuery{Subject: rel1.Subject},
			expected: []*relationtuple.InternalRelationTuple{rel1},
		},
		{
			query:    &relationtuple.RelationQuery{Relation: rel1.Relation},
			expected: []*relationtuple.InternalRelationTuple{rel1},
		},
		{
			query:    &relationtuple.RelationQuery{ObjectID: rel1.Object, Subject: rel1.Subject},
			expected: []*relationtuple.InternalRelationTuple{rel1},
		},
		{
			query:    &relationtuple.RelationQuery{ObjectID: rel1.Object, Relation: rel1.Relation},
			expected: []*relationtuple.InternalRelationTuple{rel1},
		},
		{
			query:    &relationtuple.RelationQuery{Subject: rel1.Subject, Relation: rel1.Relation},
			expected: []*relationtuple.InternalRelationTuple{rel1},
		},
		{
			query:    &relationtuple.RelationQuery{ObjectID: rel1.Object, Subject: rel1.Subject, Relation: rel1.Relation},
			expected: []*relationtuple.InternalRelationTuple{rel1},
		},
		{
			query:    &relationtuple.RelationQuery{ObjectID: rel2.Object},
			expected: []*relationtuple.InternalRelationTuple{rel2},
		},
		{
			query:    &relationtuple.RelationQuery{Subject: rel3.Subject},
			expected: []*relationtuple.InternalRelationTuple{rel3},
		},
		{
			query:    &relationtuple.RelationQuery{ObjectID: rel3.Object, Relation: rel3.Relation},
			expected: []*relationtuple.InternalRelationTuple{rel3},
		},
		{
			query:    &relationtuple.RelationQuery{},
			expected: []*relationtuple.InternalRelationTuple{rel1, rel2, rel3, rel4},
		},
		{
			expected: []*relationtuple.InternalRelationTuple{},
		},
	} {
		t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
			res, err := p.GetRelationTuples(context.Background(), tc.query)
			require.NoError(t, err)

			assert.Equal(t, len(tc.expected), len(res), "expected: %+v\n got: %+v", tc.expected, res)
			for _, r := range tc.expected {
				assert.Contains(t, res, r)
			}
		})
	}
}
