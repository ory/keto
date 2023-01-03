// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ory/herodot"
	"github.com/ory/x/pointerx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/ketoapi"
)

func TestMapper(t *testing.T) {
	ctx := context.Background()
	reg := driver.NewSqliteTestRegistry(t, false)
	nspace := namespace.Namespace{
		Name: "test",
	}
	require.NoError(t, reg.Config(ctx).Set(config.KeyNamespaces, []*namespace.Namespace{&nspace}))

	t.Run("items=relationships", func(t *testing.T) {
		for _, tc := range []struct {
			name string
			rts  []*ketoapi.RelationTuple
			err  error
		}{
			{
				name: "no relationships",
				rts:  []*ketoapi.RelationTuple{},
			},
			{
				name: "one relation tuple",
				rts: []*ketoapi.RelationTuple{
					{
						Namespace: nspace.Name,
						Object:    "object",
						Relation:  "relation",
						SubjectID: pointerx.Ptr("subject"),
					},
				},
			},
			{
				name: "relation tuple without subject",
				rts: []*ketoapi.RelationTuple{
					{
						Namespace: nspace.Name,
						Object:    "object",
						Relation:  "relation",
					},
				},
				err: ketoapi.ErrNilSubject,
			},
			{
				name: "many relationships",
				rts: func() []*ketoapi.RelationTuple {
					rts := make([]*ketoapi.RelationTuple, 10)
					for i := range rts {
						rts[i] = &ketoapi.RelationTuple{
							Namespace: nspace.Name,
							Object:    fmt.Sprintf("object %d", i),
							Relation:  "relation",
							SubjectID: pointerx.Ptr("subject"),
						}
					}
					return rts
				}(),
			},
			{
				name: "unknown namespace",
				rts: []*ketoapi.RelationTuple{
					{
						Namespace: "unknown",
						Object:    "object",
						Relation:  "relation",
						SubjectID: pointerx.Ptr("subject"),
					},
				},
				err: herodot.ErrNotFound,
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
					Namespace: pointerx.Ptr(nspace.Name),
					Object:    pointerx.Ptr("object"),
					Relation:  pointerx.Ptr("relation"),
					SubjectID: pointerx.Ptr("subject"),
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
					Relation: pointerx.Ptr("relation"),
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
					Namespace: pointerx.Ptr("unknown"),
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
				assert.Equal(t, nspace.Name, mapped.Namespace)
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
					Type:    ketoapi.TreeNodeLeaf,
					Subject: &relationtuple.SubjectID{ID: uuids[0]},
				},
			},
			{
				name: "basic tree with children",
				tree: &relationtuple.Tree{
					Type: ketoapi.TreeNodeUnion,
					Subject: &relationtuple.SubjectSet{
						Namespace: nspace.Name,
						Object:    uuids[0],
						Relation:  "members",
					},
					Children: []*relationtuple.Tree{
						{
							Type:    ketoapi.TreeNodeLeaf,
							Subject: &relationtuple.SubjectID{ID: uuids[1]},
						},
						{
							Type:    ketoapi.TreeNodeLeaf,
							Subject: &relationtuple.SubjectID{ID: uuids[2]},
						},
					},
				},
			},
			{
				name: "deeply nested tree",
				tree: &relationtuple.Tree{
					Type: ketoapi.TreeNodeUnion,
					Subject: &relationtuple.SubjectSet{
						Namespace: nspace.Name,
						Object:    uuids[0],
						Relation:  "members",
					},
					Children: []*relationtuple.Tree{
						{
							Type: ketoapi.TreeNodeUnion,
							Subject: &relationtuple.SubjectSet{
								Namespace: nspace.Name,
								Object:    uuids[1],
								Relation:  "members",
							},
							Children: []*relationtuple.Tree{
								{
									Type:    ketoapi.TreeNodeLeaf,
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

				var checkTree func(*ketoapi.Tree[*ketoapi.RelationTuple], *relationtuple.Tree)
				checkTree = func(mapped *ketoapi.Tree[*ketoapi.RelationTuple], original *relationtuple.Tree) {
					switch s := original.Subject.(type) {
					case *relationtuple.SubjectID:
						require.NotNil(t, mapped.Tuple.SubjectID)
						assert.Nil(t, mapped.Tuple.SubjectSet)
						assert.Equal(t, strs[slices.Index(uuids, s.ID)], *mapped.Tuple.SubjectID)
					case *relationtuple.SubjectSet:
						require.NotNil(t, mapped.Tuple.SubjectSet)
						assert.Nil(t, mapped.Tuple.SubjectID)
						assert.Equal(t, nspace.Name, mapped.Tuple.SubjectSet.Namespace)
						assert.Equal(t, strs[slices.Index(uuids, s.Object)], mapped.Tuple.SubjectSet.Object)
						assert.Equal(t, s.Relation, mapped.Tuple.SubjectSet.Relation)
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
}
