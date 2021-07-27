package acl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelationTupleToDeltas(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		rts      []*RelationTuple
		action   RelationTupleDelta_Action
		expected []*RelationTupleDelta
	}{
		{
			desc:     "empty list",
			rts:      nil,
			action:   RelationTupleDelta_INSERT,
			expected: []*RelationTupleDelta{},
		},
		{
			desc: "some tuples",
			rts: []*RelationTuple{
				{
					Namespace: "n",
					Object:    "o1",
					Relation:  "r",
					Subject:   NewSubjectID("s"),
				}, {
					Namespace: "n",
					Object:    "o2",
					Relation:  "r",
					Subject:   NewSubjectSet("sn", "so", "sr"),
				},
			},
			action: RelationTupleDelta_DELETE,
			expected: []*RelationTupleDelta{
				{
					RelationTuple: &RelationTuple{
						Namespace: "n",
						Object:    "o1",
						Relation:  "r",
						Subject:   NewSubjectID("s"),
					},
					Action: RelationTupleDelta_DELETE,
				},
				{
					RelationTuple: &RelationTuple{
						Namespace: "n",
						Object:    "o2",
						Relation:  "r",
						Subject:   NewSubjectSet("sn", "so", "sr"),
					},
					Action: RelationTupleDelta_DELETE,
				},
			},
		},
	} {
		t.Run("case="+tc.desc, func(t *testing.T) {
			assert.Equal(t, tc.expected, RelationTupleToDeltas(tc.rts, tc.action))
		})
	}
}
