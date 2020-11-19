package expand_test

import (
	"context"
	"testing"

	"github.com/ory/keto/internal/namespace"

	"github.com/ory/keto/internal/relationtuple"

	"github.com/ory/keto/internal/expand"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
)

func newTestEngine(_ *testing.T) (*driver.RegistryDefault, *expand.Engine) {
	reg := &driver.RegistryDefault{}
	e := expand.NewEngine(reg)
	return reg, e
}

func TestEngine(t *testing.T) {
	t.Run("case=returns UserID on expand", func(t *testing.T) {
		user := &relationtuple.UserID{ID: "user"}
		_, e := newTestEngine(t)

		tree, err := e.BuildTree(context.Background(), user, 100)
		require.NoError(t, err)
		assert.Equal(t, &expand.Tree{
			Type:    expand.Leaf,
			Subject: user,
		}, tree)
	})

	t.Run("case=expands one level", func(t *testing.T) {
		tommy := &relationtuple.UserID{ID: "Tommy"}
		paul := &relationtuple.UserID{ID: "Paul"}
		boulderGroup := "boulder group"
		defaultNamespace := "default"

		bouldererUserSet := &relationtuple.UserSet{
			Relation:  "member",
			ObjectID:  boulderGroup,
			Namespace: defaultNamespace,
		}
		boulderers := []*relationtuple.InternalRelationTuple{
			{
				Relation:  bouldererUserSet.Relation,
				Namespace: defaultNamespace,
				ObjectID:  boulderGroup,
				Subject:   tommy,
			},
			{
				Relation:  bouldererUserSet.Relation,
				Namespace: defaultNamespace,
				ObjectID:  boulderGroup,
				Subject:   paul,
			},
		}
		reg, e := newTestEngine(t)

		require.NoError(t, reg.NamespaceManager().MigrateNamespaceUp(&namespace.Namespace{Name: defaultNamespace}))
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), boulderers...))

		tree, err := e.BuildTree(context.Background(), bouldererUserSet, 100)
		require.NoError(t, err)
		assert.Equal(t, &expand.Tree{
			Type:    expand.Union,
			Subject: bouldererUserSet,
			Children: []*expand.Tree{
				{
					Type:    expand.Leaf,
					Subject: tommy,
				},
				{
					Type:    expand.Leaf,
					Subject: paul,
				},
			},
		}, tree)
	})

	t.Run("case=expands two levels", func(t *testing.T) {
		reg, e := newTestEngine(t)
		expectedTree := &expand.Tree{
			Type: expand.Union,
			Subject: &relationtuple.UserSet{
				ObjectID: "z",
				Relation: "transitive member",
			},
			Children: []*expand.Tree{
				{
					Type: expand.Union,
					Subject: &relationtuple.UserSet{
						ObjectID: "x",
						Relation: "member",
					},
					Children: []*expand.Tree{
						{
							Type:    expand.Leaf,
							Subject: &relationtuple.UserID{ID: "a"},
						},
						{
							Type:    expand.Leaf,
							Subject: &relationtuple.UserID{ID: "b"},
						},
						{
							Type:    expand.Leaf,
							Subject: &relationtuple.UserID{ID: "c"},
						},
					},
				},
				{
					Type: expand.Union,
					Subject: &relationtuple.UserSet{
						ObjectID: "y",
						Relation: "member",
					},
					Children: []*expand.Tree{
						{
							Type:    expand.Leaf,
							Subject: &relationtuple.UserID{ID: "d"},
						},
						{
							Type:    expand.Leaf,
							Subject: &relationtuple.UserID{ID: "e"},
						},
						{
							Type:    expand.Leaf,
							Subject: &relationtuple.UserID{ID: "f"},
						},
					},
				},
			},
		}

		require.NoError(t, reg.NamespaceManager().MigrateNamespaceUp(&namespace.Namespace{Name: ""}))

		for _, group := range expectedTree.Children {
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &relationtuple.InternalRelationTuple{
				ObjectID: expectedTree.Subject.(*relationtuple.UserSet).ObjectID,
				Relation: "transitive member",
				Subject: &relationtuple.UserSet{
					ObjectID: group.Subject.(*relationtuple.UserSet).ObjectID,
					Relation: "member",
				},
			}))

			for _, user := range group.Children {
				require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &relationtuple.InternalRelationTuple{
					ObjectID: group.Subject.(*relationtuple.UserSet).ObjectID,
					Relation: "member",
					Subject:  user.Subject.(*relationtuple.UserID),
				}))
			}
		}

		actualTree, err := e.BuildTree(context.Background(), expectedTree.Subject, 100)
		require.NoError(t, err)
		assert.Equal(t, expectedTree, actualTree, "%+v", actualTree.Children[0].Children)
	})

	t.Run("case=respects max depth", func(t *testing.T) {
		reg, e := newTestEngine(t)
		require.NoError(t, reg.NamespaceManager().MigrateNamespaceUp(&namespace.Namespace{Name: ""}))

		root := "root"
		prev := root
		for _, sub := range []string{"0", "1", "2", "3"} {
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &relationtuple.InternalRelationTuple{
				ObjectID: prev,
				Relation: "child",
				Subject: &relationtuple.UserSet{
					ObjectID: sub,
					Relation: "child",
				},
			}))
			prev = sub
		}

		expectedTree := &expand.Tree{
			Type: expand.Union,
			Subject: &relationtuple.UserSet{
				ObjectID: root,
				Relation: "child",
			},
			Children: []*expand.Tree{
				{
					Type: expand.Union,
					Subject: &relationtuple.UserSet{
						ObjectID: "0",
						Relation: "child",
					},
					Children: []*expand.Tree{
						{
							Type: expand.Union,
							Subject: &relationtuple.UserSet{
								ObjectID: "1",
								Relation: "child",
							},
							Children: []*expand.Tree{
								{
									Type: expand.Leaf,
									Subject: &relationtuple.UserSet{
										ObjectID: "2",
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
}
