package memory

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/models"
)

func newTestSetup(t *testing.T) (p *Persister, rel1, rel2, rel3, rel4 *models.InternalRelationTuple) {
	p = NewPersister()
	rel1 = &models.InternalRelationTuple{
		Object:   &models.Object{ID: "obj", Namespace: "rel1"},
		Relation: "rel1 name",
		Subject:  &models.UserID{ID: "rel1 user"},
	}
	rel2 = &models.InternalRelationTuple{
		Object:   &models.Object{ID: "obj", Namespace: "rel2"},
		Relation: "rel2 name",
		Subject:  &models.UserID{ID: "rel2 user"},
	}
	rel3 = &models.InternalRelationTuple{
		Object:   &models.Object{ID: "obj", Namespace: "rel3"},
		Relation: "shared name",
		Subject: &models.UserSet{
			Object:   &models.Object{ID: "user set obj", Namespace: "rel3"},
			Relation: "rel3 user set",
		},
	}
	rel4 = &models.InternalRelationTuple{
		Object:   &models.Object{ID: "obj", Namespace: "rel4"},
		Relation: "shared name",
		Subject:  &models.UserID{ID: "rel4 user"},
	}

	require.NoError(t, p.WriteRelationTuples(context.Background(), rel1, rel2, rel3, rel4))
	return
}

func TestGetRelationTuples(t *testing.T) {
	p, rel1, rel2, rel3, rel4 := newTestSetup(t)

	for i, tc := range []struct {
		query    *models.RelationQuery
		expected []*models.InternalRelationTuple
	}{
		{
			query:    &models.RelationQuery{Object: rel1.Object},
			expected: []*models.InternalRelationTuple{rel1},
		},
		{
			query:    &models.RelationQuery{Subject: rel1.Subject},
			expected: []*models.InternalRelationTuple{rel1},
		},
		{
			query:    &models.RelationQuery{Relation: rel1.Relation},
			expected: []*models.InternalRelationTuple{rel1},
		},
		{
			query:    &models.RelationQuery{Object: rel1.Object, Subject: rel1.Subject},
			expected: []*models.InternalRelationTuple{rel1},
		},
		{
			query:    &models.RelationQuery{Object: rel1.Object, Relation: rel1.Relation},
			expected: []*models.InternalRelationTuple{rel1},
		},
		{
			query:    &models.RelationQuery{Subject: rel1.Subject, Relation: rel1.Relation},
			expected: []*models.InternalRelationTuple{rel1},
		},
		{
			query:    &models.RelationQuery{Object: rel1.Object, Subject: rel1.Subject, Relation: rel1.Relation},
			expected: []*models.InternalRelationTuple{rel1},
		},
		{
			query:    &models.RelationQuery{Object: rel2.Object},
			expected: []*models.InternalRelationTuple{rel2},
		},
		{
			query:    &models.RelationQuery{Subject: rel3.Subject},
			expected: []*models.InternalRelationTuple{rel3},
		},
		{
			query:    &models.RelationQuery{Object: rel3.Object, Relation: rel3.Relation},
			expected: []*models.InternalRelationTuple{rel3},
		},
		{
			query:    &models.RelationQuery{},
			expected: []*models.InternalRelationTuple{rel1, rel2, rel3, rel4},
		},
		{
			expected: []*models.InternalRelationTuple{},
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
