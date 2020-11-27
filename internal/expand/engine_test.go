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

func newTestEngine(t *testing.T) (driver.Registry, *expand.Engine) {
	reg := driver.NewMemoryTestDriver(t).Registry()
	e := expand.NewEngine(reg)
	return reg, e
}

func TestEngine(t *testing.T) {
	t.Run("case=returns SubjectID on expand", func(t *testing.T) {
		user := &relationtuple.SubjectID{ID: "user"}
		_, e := newTestEngine(t)

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
		reg, e := newTestEngine(t)

		require.NoError(t, reg.NamespaceManager().MigrateNamespaceUp(context.Background(), &namespace.Namespace{Name: ""}))
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

		require.NoError(t, reg.NamespaceManager().MigrateNamespaceUp(context.Background(), &namespace.Namespace{Name: ""}))

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
		reg, e := newTestEngine(t)
		require.NoError(t, reg.NamespaceManager().MigrateNamespaceUp(context.Background(), &namespace.Namespace{Name: ""}))

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
}
