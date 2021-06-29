package expand_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ory/keto/internal/x"

	"github.com/ory/keto/internal/namespace"

	"github.com/ory/keto/internal/relationtuple"

	"github.com/ory/keto/internal/expand"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
)

func newTestEngine(t *testing.T, namespaces []*namespace.Namespace, paginationOpts ...x.PaginationOptionSetter) (*relationtuple.ManagerWrapper, *expand.Engine) {
	innerReg := driver.NewMemoryTestRegistry(t, namespaces)
	reg := relationtuple.NewManagerWrapper(t, innerReg, paginationOpts...)
	e := expand.NewEngine(reg)
	return reg, e
}

func TestEngine(t *testing.T) {
	t.Run("case=returns SubjectID on expand", func(t *testing.T) {
		user := &relationtuple.SubjectID{ID: "user"}
		_, e := newTestEngine(t, []*namespace.Namespace{})

		tree, err := e.BuildTree(context.Background(), user, 100)
		require.NoError(t, err)
		assert.Equal(t, &expand.Tree{
			Type:    expand.Leaf,
			Subject: user,
		}, tree)
	})

	t.Run("case=expands one level", func(t *testing.T) {
		tommy := &relationtuple.SubjectID{ID: "Tommy"}
		paul := &relationtuple.SubjectID{ID: "Paul"}
		boulderGroup := "boulder group"
		bouldererUserSet := &relationtuple.SubjectSet{
			Relation: "member",
			Object:   boulderGroup,
		}
		boulderers := []*relationtuple.InternalRelationTuple{
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
		assert.Equal(t, &expand.Tree{
			Type:    expand.Union,
			Subject: bouldererUserSet,
			Children: []*expand.Tree{
				{
					Type:    expand.Leaf,
					Subject: paul,
				},
				{
					Type:    expand.Leaf,
					Subject: tommy,
				},
			},
		}, tree)
	})

	t.Run("case=expands two levels", func(t *testing.T) {
		expectedTree := &expand.Tree{
			Type: expand.Union,
			Subject: &relationtuple.SubjectSet{
				Object:   "z",
				Relation: "transitive member",
			},
			Children: []*expand.Tree{
				{
					Type: expand.Union,
					Subject: &relationtuple.SubjectSet{
						Object:   "x",
						Relation: "member",
					},
					Children: []*expand.Tree{
						{
							Type:    expand.Leaf,
							Subject: &relationtuple.SubjectID{ID: "a"},
						},
						{
							Type:    expand.Leaf,
							Subject: &relationtuple.SubjectID{ID: "b"},
						},
						{
							Type:    expand.Leaf,
							Subject: &relationtuple.SubjectID{ID: "c"},
						},
					},
				},
				{
					Type: expand.Union,
					Subject: &relationtuple.SubjectSet{
						Object:   "y",
						Relation: "member",
					},
					Children: []*expand.Tree{
						{
							Type:    expand.Leaf,
							Subject: &relationtuple.SubjectID{ID: "d"},
						},
						{
							Type:    expand.Leaf,
							Subject: &relationtuple.SubjectID{ID: "e"},
						},
						{
							Type:    expand.Leaf,
							Subject: &relationtuple.SubjectID{ID: "f"},
						},
					},
				},
			},
		}
		fmt.Println(expectedTree.String())

		reg, e := newTestEngine(t, []*namespace.Namespace{{}})

		for _, group := range expectedTree.Children {
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &relationtuple.InternalRelationTuple{
				Object:   expectedTree.Subject.(*relationtuple.SubjectSet).Object,
				Relation: "transitive member",
				Subject: &relationtuple.SubjectSet{
					Object:   group.Subject.(*relationtuple.SubjectSet).Object,
					Relation: "member",
				},
			}))

			for _, user := range group.Children {
				require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &relationtuple.InternalRelationTuple{
					Object:   group.Subject.(*relationtuple.SubjectSet).Object,
					Relation: "member",
					Subject:  user.Subject.(*relationtuple.SubjectID),
				}))
			}
		}

		actualTree, err := e.BuildTree(context.Background(), expectedTree.Subject, 100)
		require.NoError(t, err)
		assert.Equal(t, expectedTree, actualTree, "%+v", actualTree.Children[0].Children)
	})

	t.Run("case=respects max depth", func(t *testing.T) {
		reg, e := newTestEngine(t, []*namespace.Namespace{{}})

		root := "root"
		prev := root
		for _, sub := range []string{"0", "1", "2", "3"} {
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &relationtuple.InternalRelationTuple{
				Object:   prev,
				Relation: "child",
				Subject: &relationtuple.SubjectSet{
					Object:   sub,
					Relation: "child",
				},
			}))
			prev = sub
		}

		expectedTree := &expand.Tree{
			Type: expand.Union,
			Subject: &relationtuple.SubjectSet{
				Object:   root,
				Relation: "child",
			},
			Children: []*expand.Tree{
				{
					Type: expand.Union,
					Subject: &relationtuple.SubjectSet{
						Object:   "0",
						Relation: "child",
					},
					Children: []*expand.Tree{
						{
							Type: expand.Union,
							Subject: &relationtuple.SubjectSet{
								Object:   "1",
								Relation: "child",
							},
							Children: []*expand.Tree{
								{
									Type: expand.Leaf,
									Subject: &relationtuple.SubjectSet{
										Object:   "2",
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
		reg, e := newTestEngine(t, []*namespace.Namespace{{}}, x.WithSize(2))

		users := []string{"u1", "u2", "u3", "u4"}
		expectedTree := &expand.Tree{
			Type:    expand.Union,
			Subject: &relationtuple.SubjectSet{Object: "root", Relation: "access"},
		}

		for _, user := range users {
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &relationtuple.InternalRelationTuple{
				Object:   "root",
				Relation: "access",
				Subject:  &relationtuple.SubjectID{ID: user},
			}))
			expectedTree.Children = append(expectedTree.Children, &expand.Tree{
				Type:    expand.Leaf,
				Subject: &relationtuple.SubjectID{ID: user},
			})
		}

		tree, err := e.BuildTree(context.Background(), &relationtuple.SubjectSet{
			Object:   "root",
			Relation: "access",
		}, 10)
		require.NoError(t, err)

		assert.Equal(t, expectedTree, tree)
		assert.Len(t, reg.RequestedPages, 2)
	})

	t.Run("case=handles subject sets as leaf", func(t *testing.T) {
		reg, e := newTestEngine(t, []*namespace.Namespace{{}})

		expectedTree := &expand.Tree{
			Type: expand.Union,
			Subject: &relationtuple.SubjectSet{
				Object:   "root",
				Relation: "rel",
			},
			Children: []*expand.Tree{
				{
					Type: expand.Leaf,
					Subject: &relationtuple.SubjectSet{
						Object:   "so",
						Relation: "sr",
					},
				},
			},
		}

		require.NoError(t, reg.WriteRelationTuples(context.Background(), &relationtuple.InternalRelationTuple{
			Object:   expectedTree.Subject.(*relationtuple.SubjectSet).Object,
			Relation: expectedTree.Subject.(*relationtuple.SubjectSet).Relation,
			Subject:  expectedTree.Children[0].Subject,
		}))

		tree, err := e.BuildTree(context.Background(), expectedTree.Subject, 100)
		require.NoError(t, err)
		assert.Equal(t, expectedTree, tree)
	})

	t.Run("case=circular tuples", func(t *testing.T) {
		sendlingerTor, odeonsplatz, centralStation, connected, namesp := "Sendlinger Tor", "Odeonsplatz", "Central Station", "connected", "munich transport"

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

		expectedTree := &expand.Tree{
			Type:    expand.Union,
			Subject: sendlingerTorSS,
			Children: []*expand.Tree{
				{
					Type:    expand.Union,
					Subject: odeonsplatzSS,
					Children: []*expand.Tree{
						{
							Type:    expand.Union,
							Subject: centralStationSS,
							Children: []*expand.Tree{
								{
									Type:     expand.Leaf,
									Subject:  sendlingerTorSS,
									Children: nil,
								},
							},
						},
					},
				},
			},
		}

		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), []*relationtuple.InternalRelationTuple{
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
}
