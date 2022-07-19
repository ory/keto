package relationtuple_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ory/herodot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoapi"
)

func TestMapper(t *testing.T) {
	ctx := context.Background()
	reg := driver.NewSqliteTestRegistry(t, false)
	nspace := namespace.Namespace{
		ID:   1,
		Name: "test",
	}
	require.NoError(t, reg.Config(ctx).Set(config.KeyNamespaces, []*namespace.Namespace{&nspace}))

	t.Run("items=relation tuples", func(t *testing.T) {
		for _, tc := range []struct {
			name string
			rts  []*ketoapi.RelationTuple
			err  error
		}{
			{
				name: "no relation tuples",
				rts:  []*ketoapi.RelationTuple{},
			},
			{
				name: "one relation tuple",
				rts: []*ketoapi.RelationTuple{
					{
						Namespace: nspace.Name,
						Object:    "object",
						Relation:  "relation",
						SubjectID: x.Ptr("subject"),
					},
				},
			},
			{
				name: "many relation tuples",
				rts: func() []*ketoapi.RelationTuple {
					rts := make([]*ketoapi.RelationTuple, 10)
					for i := range rts {
						rts[i] = &ketoapi.RelationTuple{
							Namespace: nspace.Name,
							Object:    fmt.Sprintf("object %d", i),
							Relation:  "relation",
							SubjectID: x.Ptr("subject"),
						}
					}
					return rts
				}(),
			},
		} {
			t.Run(tc.name, func(t *testing.T) {
				m := reg.Mapper()
				mapped, err := m.FromTuple(ctx, tc.rts...)
				require.ErrorIs(t, err, tc.err)
				if tc.err != nil {
					// the rest only makes sense if we have no error
					return
				}

				assert.Len(t, mapped, len(tc.rts))
				actual, err := m.ToTuple(ctx, mapped...)
				require.NoError(t, err)
				assert.ElementsMatch(t, tc.rts, actual)
			})
		}
	})

	t.Run("item=relation query", func(t *testing.T) {
		for _, tc := range []struct {
			name  string
			query *ketoapi.RelationQuery
			err   error
		}{
			{
				name:  "all fields nil",
				query: &ketoapi.RelationQuery{},
			},
			{
				name: "all fields set",
				query: &ketoapi.RelationQuery{
					Namespace: x.Ptr(nspace.Name),
					Object:    x.Ptr("object"),
					Relation:  x.Ptr("relation"),
					SubjectID: x.Ptr("subject"),
				},
			},
			{
				name: "subject set",
				query: &ketoapi.RelationQuery{
					SubjectSet: &ketoapi.SubjectSet{
						Namespace: nspace.Name,
						Object:    "object",
						Relation:  "relation",
					},
				},
			},
			{
				name: "non-mapped fields",
				query: &ketoapi.RelationQuery{
					Relation: x.Ptr("relation"),
					SubjectSet: &ketoapi.SubjectSet{
						Namespace: nspace.Name,
						Relation:  "relation",
						Object:    "object",
					},
				},
			},
			{
				name: "unknown namespace",
				query: &ketoapi.RelationQuery{
					Namespace: x.Ptr("unknown"),
				},
				err: herodot.ErrNotFound,
			},
		} {
			t.Run(tc.name, func(t *testing.T) {
				m := reg.Mapper()
				mapped, err := m.FromQuery(ctx, tc.query)
				require.ErrorIs(t, err, tc.err)
				if tc.err != nil {
					// the rest only makes sense if we have no error
					return
				}

				actual, err := m.ToQuery(ctx, mapped)
				require.NoError(t, err)
				assert.Equal(t, tc.query, actual)
			})
		}
	})

	t.Run("item=subject set", func(t *testing.T) {
		for _, tc := range []struct {
			name string
			set  *ketoapi.SubjectSet
			err  error
		}{
			{
				name: "basic subject set",
				set: &ketoapi.SubjectSet{
					Namespace: nspace.Name,
					Object:    "object",
					Relation:  "relation",
				},
			},
			{
				name: "unknown namespace",
				set: &ketoapi.SubjectSet{
					Namespace: "unknown",
					Object:    "object",
					Relation:  "relation",
				},
				err: herodot.ErrNotFound,
			},
		} {
			t.Run(tc.name, func(t *testing.T) {
				mapped, err := reg.Mapper().FromSubjectSet(ctx, tc.set)
				require.ErrorIs(t, err, tc.err)
				if tc.err != nil {
					// the rest only makes sense if we have no error
					return
				}
				assert.NotZero(t, mapped.Object)
				assert.Equal(t, nspace.ID, mapped.Namespace)
				assert.Equal(t, tc.set.Relation, mapped.Relation)
			})
		}
	})

	t.Run("item=expand tree", func(t *testing.T) {
		strs := make([]string, 3)
		for i := range strs {
			strs[i] = fmt.Sprintf("foo %d", i)
		}
		uuids, err := reg.MappingManager().MapStringsToUUIDs(ctx, strs...)
		require.NoError(t, err)

		for _, tc := range []struct {
			name string
			tree *relationtuple.Tree
			err  error
		}{
			{
				name: "basic tree",
				tree: &relationtuple.Tree{
					Type:    ketoapi.Leaf,
					Subject: &relationtuple.SubjectID{ID: uuids[0]},
				},
			},
			{
				name: "basic tree with children",
				tree: &relationtuple.Tree{
					Type: ketoapi.Union,
					Subject: &relationtuple.SubjectSet{
						Namespace: nspace.ID,
						Object:    uuids[0],
						Relation:  "members",
					},
					Children: []*relationtuple.Tree{
						{
							Type:    ketoapi.Leaf,
							Subject: &relationtuple.SubjectID{ID: uuids[1]},
						},
						{
							Type:    ketoapi.Leaf,
							Subject: &relationtuple.SubjectID{ID: uuids[2]},
						},
					},
				},
			},
			{
				name: "deeply nested tree",
				tree: &relationtuple.Tree{
					Type: ketoapi.Union,
					Subject: &relationtuple.SubjectSet{
						Namespace: nspace.ID,
						Object:    uuids[0],
						Relation:  "members",
					},
					Children: []*relationtuple.Tree{
						{
							Type: ketoapi.Union,
							Subject: &relationtuple.SubjectSet{
								Namespace: nspace.ID,
								Object:    uuids[1],
								Relation:  "members",
							},
							Children: []*relationtuple.Tree{
								{
									Type:    ketoapi.Leaf,
									Subject: &relationtuple.SubjectID{ID: uuids[2]},
								},
							},
						},
					},
				},
			},
		} {
			t.Run(tc.name, func(t *testing.T) {
				mapped, err := reg.Mapper().ToTree(ctx, tc.tree)
				require.ErrorIs(t, err, tc.err)
				if tc.err != nil {
					// the rest only makes sense if we have no error
					return
				}

				var checkTree func(*ketoapi.ExpandTree, *relationtuple.Tree)
				checkTree = func(mapped *ketoapi.ExpandTree, original *relationtuple.Tree) {
					switch s := original.Subject.(type) {
					case *relationtuple.SubjectID:
						require.NotNil(t, mapped.SubjectID)
						assert.Nil(t, mapped.SubjectSet)
						assert.Equal(t, strs[slices.Index(uuids, s.ID)], *mapped.SubjectID)
					case *relationtuple.SubjectSet:
						require.NotNil(t, mapped.SubjectSet)
						assert.Nil(t, mapped.SubjectID)
						assert.Equal(t, nspace.Name, mapped.SubjectSet.Namespace)
						assert.Equal(t, strs[slices.Index(uuids, s.Object)], mapped.SubjectSet.Object)
						assert.Equal(t, s.Relation, mapped.SubjectSet.Relation)
					default:
						t.Fatalf("expected subject to be set: %+v", mapped)
					}

					assert.Equal(t, original.Type, mapped.Type)
					require.Len(t, mapped.Children, len(original.Children))
					for i := range original.Children {
						checkTree(mapped.Children[i], original.Children[i])
					}
				}
				checkTree(mapped, tc.tree)
			})
		}
	})

	t.Run("case=unknown namespace", func(t *testing.T) {
		t.Run("direction=from tuple", func(t *testing.T) {
			rts := make([]*ketoapi.RelationTuple, 20)
			var expected []*ketoapi.RelationTuple
			for i := 0; i < len(rts); i = i + 3 {
				rts[i] = &ketoapi.RelationTuple{
					Namespace: nspace.Name,
					Object:    "o",
					Relation:  "r",
					SubjectID: x.Ptr("s"),
				}
				expected = append(expected, rts[i])
			}
			for i := 1; i < len(rts); i = i + 3 {
				rts[i] = &ketoapi.RelationTuple{
					Namespace: "unknown",
					Object:    "o",
					Relation:  "r",
					SubjectID: x.Ptr("s"),
				}
			}
			for i := 2; i < len(rts); i = i + 3 {
				rts[i] = &ketoapi.RelationTuple{
					Namespace: nspace.Name,
					Object:    "o",
					Relation:  "r",
					SubjectSet: &ketoapi.SubjectSet{
						Namespace: "unknown",
						Object:    "o",
						Relation:  "r",
					},
				}
			}

			actual, err := reg.Mapper().FromTuple(ctx, rts...)
			require.NoError(t, err)

			expectedMapped, err := reg.Mapper().FromTuple(ctx, expected...)
			require.NoError(t, err)
			require.Len(t, expectedMapped, len(expected))
			assert.ElementsMatch(t, expectedMapped, actual)
		})
	})

	t.Run("direction=to tuple", func(t *testing.T) {
		ids, err := reg.MappingManager().MapStringsToUUIDs(ctx, "s")
		require.NoError(t, err)
		id := ids[0]

		rts := make([]*relationtuple.RelationTuple, 20)
		var expected []*relationtuple.RelationTuple
		for i := 0; i < len(rts); i = i + 3 {
			rts[i] = &relationtuple.RelationTuple{
				Namespace: nspace.ID,
				Object:    id,
				Relation:  "r",
				Subject:   &relationtuple.SubjectID{ID: id},
			}
			expected = append(expected, rts[i])
		}
		for i := 1; i < len(rts); i = i + 3 {
			rts[i] = &relationtuple.RelationTuple{
				Namespace: nspace.ID + 1,
				Object:    id,
				Relation:  "r",
				Subject:   &relationtuple.SubjectID{ID: id},
			}
		}
		for i := 2; i < len(rts); i = i + 3 {
			rts[i] = &relationtuple.RelationTuple{
				Namespace: nspace.ID,
				Object:    id,
				Relation:  "r",
				Subject: &relationtuple.SubjectSet{
					Namespace: nspace.ID + 1,
					Object:    id,
					Relation:  "r",
				},
			}
		}

		actual, err := reg.Mapper().ToTuple(ctx, rts...)
		require.NoError(t, err)

		expectedMapped, err := reg.Mapper().ToTuple(ctx, expected...)
		require.NoError(t, err)
		require.Len(t, expectedMapped, len(expected))
		assert.ElementsMatch(t, expectedMapped, actual)
	})
}
