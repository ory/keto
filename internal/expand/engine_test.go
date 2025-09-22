// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package expand_test

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	keysetpagination "github.com/ory/x/pagination/keysetpagination_v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoapi"
)

type (
	configProvider = config.Provider
	loggerProvider = x.LoggerProvider
)

// deps is defined to capture engine dependencies in a single struct
type deps struct {
	*relationtuple.ManagerWrapper // managerProvider
	configProvider
	loggerProvider
	x.TracingProvider
	x.NetworkIDProvider
}

func newTestEngine(t *testing.T, namespaces []*namespace.Namespace, paginationOpts ...keysetpagination.Option) (*relationtuple.ManagerWrapper, *expand.Engine) {
	innerReg := driver.NewSqliteTestRegistry(t, false)
	require.NoError(t, innerReg.Config(context.Background()).Set(config.KeyNamespaces, namespaces))
	reg := relationtuple.NewManagerWrapper(t, innerReg, paginationOpts...)
	e := expand.NewEngine(&deps{
		ManagerWrapper:    reg,
		configProvider:    innerReg,
		loggerProvider:    innerReg,
		TracingProvider:   innerReg,
		NetworkIDProvider: innerReg,
	})
	return reg, e
}

func TestEngine(t *testing.T) {
	t.Run("case=returns SubjectID on expand", func(t *testing.T) {
		user := &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())}
		_, e := newTestEngine(t, []*namespace.Namespace{})

		tree, err := e.BuildTree(context.Background(), user, 100)
		require.NoError(t, err)
		assert.Equal(t, &relationtuple.Tree{
			Type:    ketoapi.TreeNodeLeaf,
			Subject: user,
		}, tree)
	})

	t.Run("case=expands one level", func(t *testing.T) {
		tommy := &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())}
		paul := &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())}
		boulderGroup := uuid.Must(uuid.NewV4())
		bouldererUserSet := &relationtuple.SubjectSet{
			Relation: "member",
			Object:   boulderGroup,
		}
		boulderers := []*relationtuple.RelationTuple{
			{
				Relation: bouldererUserSet.Relation,
				Object:   boulderGroup,
				Subject:  tommy,
			},
			{
				Relation: bouldererUserSet.Relation,
				Object:   boulderGroup,
				Subject:  paul,
			},
		}
		reg, e := newTestEngine(t, []*namespace.Namespace{{}})

		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), boulderers...))

		tree, err := e.BuildTree(context.Background(), bouldererUserSet, 100)
		require.NoError(t, err)

		expand.AssertInternalTreesAreEqual(t, &relationtuple.Tree{
			Type:    ketoapi.TreeNodeUnion,
			Subject: bouldererUserSet,
			Children: []*relationtuple.Tree{
				{
					Type:    ketoapi.TreeNodeLeaf,
					Subject: paul,
				},
				{
					Type:    ketoapi.TreeNodeLeaf,
					Subject: tommy,
				},
			},
		}, tree)
	})

	t.Run("case=expands two levels", func(t *testing.T) {
		expectedTree := &relationtuple.Tree{
			Type: ketoapi.TreeNodeUnion,
			Subject: &relationtuple.SubjectSet{
				Object:   uuid.Must(uuid.NewV4()),
				Relation: "transitive member",
			},
			Children: []*relationtuple.Tree{
				{
					Type: ketoapi.TreeNodeUnion,
					Subject: &relationtuple.SubjectSet{
						Object:   uuid.Must(uuid.NewV4()),
						Relation: "member",
					},
					Children: []*relationtuple.Tree{
						{
							Type:    ketoapi.TreeNodeLeaf,
							Subject: &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
						},
						{
							Type:    ketoapi.TreeNodeLeaf,
							Subject: &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
						},
						{
							Type:    ketoapi.TreeNodeLeaf,
							Subject: &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
						},
					},
				},
				{
					Type: ketoapi.TreeNodeUnion,
					Subject: &relationtuple.SubjectSet{
						Object:   uuid.Must(uuid.NewV4()),
						Relation: "member",
					},
					Children: []*relationtuple.Tree{
						{
							Type:    ketoapi.TreeNodeLeaf,
							Subject: &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
						},
						{
							Type:    ketoapi.TreeNodeLeaf,
							Subject: &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
						},
						{
							Type:    ketoapi.TreeNodeLeaf,
							Subject: &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
						},
					},
				},
			},
		}

		reg, e := newTestEngine(t, []*namespace.Namespace{{}})

		for _, group := range expectedTree.Children {
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &relationtuple.RelationTuple{
				Object:   expectedTree.Subject.(*relationtuple.SubjectSet).Object,
				Relation: "transitive member",
				Subject: &relationtuple.SubjectSet{
					Object:   group.Subject.(*relationtuple.SubjectSet).Object,
					Relation: "member",
				},
			}))

			for _, user := range group.Children {
				require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &relationtuple.RelationTuple{
					Object:   group.Subject.(*relationtuple.SubjectSet).Object,
					Relation: "member",
					Subject:  user.Subject.(*relationtuple.SubjectID),
				}))
			}
		}

		actualTree, err := e.BuildTree(context.Background(), expectedTree.Subject, 100)
		require.NoError(t, err)
		expand.AssertInternalTreesAreEqual(t, expectedTree, actualTree)
	})

	t.Run("case=respects max depth", func(t *testing.T) {
		reg, e := newTestEngine(t, []*namespace.Namespace{{}})

		ids := x.UUIDs(5)
		for i := 1; i < len(ids); i++ {
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &relationtuple.RelationTuple{
				Object:   ids[i-1],
				Relation: "child",
				Subject: &relationtuple.SubjectSet{
					Object:   ids[i],
					Relation: "child",
				},
			}))
		}

		expectedTree := &relationtuple.Tree{
			Type: ketoapi.TreeNodeUnion,
			Subject: &relationtuple.SubjectSet{
				Object:   ids[0],
				Relation: "child",
			},
			Children: []*relationtuple.Tree{
				{
					Type: ketoapi.TreeNodeUnion,
					Subject: &relationtuple.SubjectSet{
						Object:   ids[1],
						Relation: "child",
					},
					Children: []*relationtuple.Tree{
						{
							Type: ketoapi.TreeNodeUnion,
							Subject: &relationtuple.SubjectSet{
								Object:   ids[2],
								Relation: "child",
							},
							Children: []*relationtuple.Tree{
								{
									Type: ketoapi.TreeNodeLeaf,
									Subject: &relationtuple.SubjectSet{
										Object:   ids[3],
										Relation: "child",
									},
								},
							},
						},
					},
				},
			},
		}

		actualTree, err := e.BuildTree(context.Background(), expectedTree.Subject, 4)
		require.NoError(t, err)

		assert.Equal(t, expectedTree, actualTree)
	})

	t.Run("case=paginates", func(t *testing.T) {
		reg, e := newTestEngine(t, []*namespace.Namespace{{}}, keysetpagination.WithSize(2))

		root := uuid.Must(uuid.NewV4())
		expectedTree := &relationtuple.Tree{
			Type:    ketoapi.TreeNodeUnion,
			Subject: &relationtuple.SubjectSet{Object: root, Relation: "access"},
		}

		for _, user := range x.UUIDs(4) {
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &relationtuple.RelationTuple{
				Object:   root,
				Relation: "access",
				Subject:  &relationtuple.SubjectID{ID: user},
			}))
			expectedTree.Children = append(expectedTree.Children, &relationtuple.Tree{
				Type:    ketoapi.TreeNodeLeaf,
				Subject: &relationtuple.SubjectID{ID: user},
			})
		}

		tree, err := e.BuildTree(context.Background(), &relationtuple.SubjectSet{
			Object:   root,
			Relation: "access",
		}, 10)
		require.NoError(t, err)

		assert.True(t, expand.AssertInternalTreesAreEqual(t, expectedTree, tree))
		assert.Len(t, reg.RequestedPages, 2)
	})

	t.Run("case=handles subject sets as leaf", func(t *testing.T) {
		reg, e := newTestEngine(t, []*namespace.Namespace{{}})

		expectedTree := &relationtuple.Tree{
			Type: ketoapi.TreeNodeUnion,
			Subject: &relationtuple.SubjectSet{
				Object:   uuid.Must(uuid.NewV4()),
				Relation: "rel",
			},
			Children: []*relationtuple.Tree{
				{
					Type: ketoapi.TreeNodeLeaf,
					Subject: &relationtuple.SubjectSet{
						Object:   uuid.Must(uuid.NewV4()),
						Relation: "sr",
					},
				},
			},
		}

		require.NoError(t, reg.WriteRelationTuples(context.Background(), &relationtuple.RelationTuple{
			Object:   expectedTree.Subject.(*relationtuple.SubjectSet).Object,
			Relation: expectedTree.Subject.(*relationtuple.SubjectSet).Relation,
			Subject:  expectedTree.Children[0].Subject,
		}))

		tree, err := e.BuildTree(context.Background(), expectedTree.Subject, 100)
		require.NoError(t, err)
		assert.Equal(t, expectedTree, tree)
	})

	t.Run("case=circular tuples", func(t *testing.T) {
		sendlingerTor, odeonsplatz, centralStation, connected, namesp := uuid.NewV5(uuid.Nil, "Sendlinger Tor"), uuid.NewV5(uuid.Nil, "Odeonsplatz"), uuid.NewV5(uuid.Nil, "Central Station"), "connected", "92384"

		sendlingerTorSS, odeonsplatzSS, centralStationSS := &relationtuple.SubjectSet{
			Namespace: namesp,
			Object:    sendlingerTor,
			Relation:  connected,
		}, &relationtuple.SubjectSet{
			Namespace: namesp,
			Object:    odeonsplatz,
			Relation:  connected,
		}, &relationtuple.SubjectSet{
			Namespace: namesp,
			Object:    centralStation,
			Relation:  connected,
		}

		reg, e := newTestEngine(t, []*namespace.Namespace{{Name: namesp}})

		expectedTree := &relationtuple.Tree{
			Type:    ketoapi.TreeNodeUnion,
			Subject: sendlingerTorSS,
			Children: []*relationtuple.Tree{
				{
					Type:    ketoapi.TreeNodeUnion,
					Subject: odeonsplatzSS,
					Children: []*relationtuple.Tree{
						{
							Type:    ketoapi.TreeNodeUnion,
							Subject: centralStationSS,
							Children: []*relationtuple.Tree{
								{
									Type:     ketoapi.TreeNodeLeaf,
									Subject:  sendlingerTorSS,
									Children: nil,
								},
							},
						},
					},
				},
			},
		}

		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), []*relationtuple.RelationTuple{
			{
				Namespace: namesp,
				Object:    sendlingerTor,
				Relation:  connected,
				Subject:   odeonsplatzSS,
			},
			{
				Namespace: namesp,
				Object:    odeonsplatz,
				Relation:  connected,
				Subject:   centralStationSS,
			},
			{
				Namespace: namesp,
				Object:    centralStation,
				Relation:  connected,
				Subject:   sendlingerTorSS,
			},
		}...))

		tree, err := e.BuildTree(context.Background(), &relationtuple.SubjectSet{
			Namespace: namesp,
			Object:    sendlingerTor,
			Relation:  connected,
		}, 100)
		require.NoError(t, err)
		assert.Equal(t, expectedTree, tree)
	})

	t.Run("case=returns result on unknown subject", func(t *testing.T) {
		_, e := newTestEngine(t, []*namespace.Namespace{})
		tree, err := e.BuildTree(context.Background(), &relationtuple.SubjectSet{
			Namespace: "unknown",
			Object:    uuid.Must(uuid.NewV4()),
			Relation:  "rel",
		}, 100)
		require.NoError(t, err)
		assert.Nil(t, tree)
	})
}
