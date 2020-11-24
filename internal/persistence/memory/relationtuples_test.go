package memory

import (
	"context"
	"fmt"
	"testing"

	"github.com/ory/keto/internal/namespace"

	"github.com/ory/keto/internal/relationtuple"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestSetup(t *testing.T) (p *Persister, rel1, rel2, rel3, rel4 *relationtuple.InternalRelationTuple) {
	p = NewPersister()
	rel1 = &relationtuple.InternalRelationTuple{
		Object:    "obj",
		Namespace: "rel1",
		Relation:  "rel1 name",
		Subject:   &relationtuple.SubjectID{ID: "rel1 user"},
	}
	rel2 = &relationtuple.InternalRelationTuple{
		Object:    "obj",
		Namespace: "rel2",
		Relation:  "rel2 name",
		Subject:   &relationtuple.SubjectID{ID: "rel2 user"},
	}
	rel3 = &relationtuple.InternalRelationTuple{
		Object:    "obj",
		Namespace: "rel3",
		Relation:  "shared name",
		Subject: &relationtuple.SubjectSet{
			Object:    "user set obj",
			Namespace: "rel3",
			Relation:  "rel3 user set",
		},
	}
	rel4 = &relationtuple.InternalRelationTuple{
		Object:    "obj",
		Namespace: "rel4",
		Relation:  "shared name",
		Subject:   &relationtuple.SubjectID{ID: "rel4 user"},
	}

	require.NoError(t, p.MigrateNamespaceUp(&namespace.Namespace{Name: rel1.Namespace, ID: 0}))
	require.NoError(t, p.MigrateNamespaceUp(&namespace.Namespace{Name: rel2.Namespace, ID: 1}))
	require.NoError(t, p.MigrateNamespaceUp(&namespace.Namespace{Name: rel3.Namespace, ID: 2}))
	require.NoError(t, p.MigrateNamespaceUp(&namespace.Namespace{Name: rel4.Namespace, ID: 3}))
	require.NoError(t, p.WriteRelationTuples(context.Background(), rel1, rel2, rel3, rel4))
	return
}

func TestGetRelationTuples(t *testing.T) {
	p, rel1, rel2, rel3, _ := newTestSetup(t)

	for i, tc := range []struct {
		query    *relationtuple.RelationQuery
		expected []*relationtuple.InternalRelationTuple
	}{
		{
			query:    &relationtuple.RelationQuery{Object: rel1.Object, Namespace: rel1.Namespace},
			expected: []*relationtuple.InternalRelationTuple{rel1},
		},
		{
			query:    &relationtuple.RelationQuery{Subject: rel1.Subject, Namespace: rel1.Namespace},
			expected: []*relationtuple.InternalRelationTuple{rel1},
		},
		{
			query:    &relationtuple.RelationQuery{Relation: rel1.Relation, Namespace: rel1.Namespace},
			expected: []*relationtuple.InternalRelationTuple{rel1},
		},
		{
			query:    &relationtuple.RelationQuery{Object: rel1.Object, Subject: rel1.Subject, Namespace: rel1.Namespace},
			expected: []*relationtuple.InternalRelationTuple{rel1},
		},
		{
			query:    &relationtuple.RelationQuery{Object: rel1.Object, Relation: rel1.Relation, Namespace: rel1.Namespace},
			expected: []*relationtuple.InternalRelationTuple{rel1},
		},
		{
			query:    &relationtuple.RelationQuery{Subject: rel1.Subject, Relation: rel1.Relation, Namespace: rel1.Namespace},
			expected: []*relationtuple.InternalRelationTuple{rel1},
		},
		{
			query:    &relationtuple.RelationQuery{Object: rel1.Object, Subject: rel1.Subject, Relation: rel1.Relation, Namespace: rel1.Namespace},
			expected: []*relationtuple.InternalRelationTuple{rel1},
		},
		{
			query:    &relationtuple.RelationQuery{Object: rel2.Object, Namespace: rel2.Namespace},
			expected: []*relationtuple.InternalRelationTuple{rel2},
		},
		{
			query:    &relationtuple.RelationQuery{Subject: rel3.Subject, Namespace: rel3.Namespace},
			expected: []*relationtuple.InternalRelationTuple{rel3},
		},
		{
			query:    &relationtuple.RelationQuery{Object: rel3.Object, Relation: rel3.Relation, Namespace: rel3.Namespace},
			expected: []*relationtuple.InternalRelationTuple{rel3},
		},
		{
			expected: []*relationtuple.InternalRelationTuple{},
		},
	} {
		t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
			res, _, err := p.GetRelationTuples(context.Background(), tc.query)
			require.NoError(t, err)

			assert.Equal(t, len(tc.expected), len(res), "query: %s\nexpected: %+v\n got: %+v", tc.query, tc.expected, res)
			for _, r := range tc.expected {
				assert.Contains(t, res, r)
			}
		})
	}
}
