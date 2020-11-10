package expand_test

import (
	"context"
	"github.com/ory/keto/expand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/driver"
	"github.com/ory/keto/models"
)

func newTestEngine(_ *testing.T) (*driver.RegistryDefault, *expand.Engine) {
	reg := &driver.RegistryDefault{}
	e := expand.NewEngine(reg)
	return reg, e
}

func TestEngine(t *testing.T) {
	t.Run("case=returns UserID on expand", func(t *testing.T) {
		user := &models.UserID{ID: "user"}
		_, e := newTestEngine(t)

		tree, err := e.BuildTree(context.Background(), user, 100)
		require.NoError(t, err)
		assert.Equal(t, &expand.Tree{
			Type:    expand.Leaf,
			Subject: user,
		}, tree)
	})

	t.Run("case=expands one level", func(t *testing.T) {
		tommy := &models.UserID{ID: "Tommy"}
		paul := &models.UserID{ID: "Paul"}
		boulderGroup := &models.Object{
			ID:        "boulder group",
			Namespace: "default",
		}
		bouldererUserSet := &models.UserSet{
			Relation: "member",
			Object:   boulderGroup,
		}
		boulderers := []*models.InternalRelationTuple{
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
			Subject: &models.UserSet{
				Object:   &models.Object{ID: "z"},
				Relation: "transitive member",
			},
			Children: []*expand.Tree{
				{
					Type: expand.Union,
					Subject: &models.UserSet{
						Object:   &models.Object{ID: "x"},
						Relation: "member",
					},
					Children: []*expand.Tree{
						{
							Type:    expand.Leaf,
							Subject: &models.UserID{ID: "a"},
						},
						{
							Type:    expand.Leaf,
							Subject: &models.UserID{ID: "b"},
						},
						{
							Type:    expand.Leaf,
							Subject: &models.UserID{ID: "c"},
						},
					},
				},
				{
					Type: expand.Union,
					Subject: &models.UserSet{
						Object:   &models.Object{ID: "y"},
						Relation: "member",
					},
					Children: []*expand.Tree{
						{
							Type:    expand.Leaf,
							Subject: &models.UserID{ID: "d"},
						},
						{
							Type:    expand.Leaf,
							Subject: &models.UserID{ID: "e"},
						},
						{
							Type:    expand.Leaf,
							Subject: &models.UserID{ID: "f"},
						},
					},
				},
			},
		}

		for _, group := range expectedTree.Children {
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &models.InternalRelationTuple{
				Object:   expectedTree.Subject.(*models.UserSet).Object,
				Relation: "transitive member",
				Subject: &models.UserSet{
					Object:   group.Subject.(*models.UserSet).Object,
					Relation: "member",
				},
			}))

			for _, user := range group.Children {
				require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &models.InternalRelationTuple{
					Object:   group.Subject.(*models.UserSet).Object,
					Relation: "member",
					Subject:  user.Subject.(*models.UserID),
				}))
			}
		}

		actualTree, err := e.BuildTree(context.Background(), expectedTree.Subject, 100)
		require.NoError(t, err)
		assert.Equal(t, expectedTree, actualTree, "%+v", actualTree.Children[0].Children)
	})

	t.Run("case=respects max depth", func(t *testing.T) {
		reg, e := newTestEngine(t)
		root := &models.Object{ID: "root"}
		prev := root
		for _, sub := range []string{"0", "1", "2", "3"} {
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &models.InternalRelationTuple{
				Object:   prev,
				Relation: "child",
				Subject: &models.UserSet{
					Object:   &models.Object{ID: sub},
					Relation: "child",
				},
			}))
			prev = &models.Object{ID: sub}
		}

		expectedTree := &expand.Tree{
			Type: expand.Union,
			Subject: &models.UserSet{
				Object:   root,
				Relation: "child",
			},
			Children: []*expand.Tree{
				{
					Type: expand.Union,
					Subject: &models.UserSet{
						Object:   &models.Object{ID: "0"},
						Relation: "child",
					},
					Children: []*expand.Tree{
						{
							Type: expand.Union,
							Subject: &models.UserSet{
								Object:   &models.Object{ID: "1"},
								Relation: "child",
							},
							Children: []*expand.Tree{
								{
									Type: expand.Leaf,
									Subject: &models.UserSet{
										Object:   &models.Object{ID: "2"},
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
