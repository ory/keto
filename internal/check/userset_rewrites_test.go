package check_test

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/namespace/ast"
	"github.com/ory/keto/internal/relationtuple"
)

var docsNS = namespace.Namespace{Name: "docs",
	ID: 1,
	Relations: []ast.Relation{
		{
			Name: "owner",
		},
		{
			Name: "editor",
			UsersetRewrite: &ast.UsersetRewrite{
				Children: ast.Children{
					ComputedUsersets: []ast.ComputedUserset{
						{
							Relation: "owner",
						},
					},
				},
			},
		},
		{
			Name: "viewer",
			UsersetRewrite: &ast.UsersetRewrite{
				Children: ast.Children{
					ComputedUsersets: []ast.ComputedUserset{{
						Relation: "editor",
					}},
					TupleToUsersets: []ast.TupleToUserset{{
						Relation:                "parent",
						ComputedUsersetRelation: "viewer",
					}},
				},
			},
		},
	},
}

func insertFixtures(t *testing.T, m relationtuple.Manager, tuples []string) {
	t.Helper()
	relationTuples := make([]*relationtuple.InternalRelationTuple, len(tuples))
	var err error
	for i, tuple := range tuples {
		relationTuples[i], err = relationtuple.InternalFromString(tuple)
		require.NoError(t, err)
	}
	require.NoError(t, m.WriteRelationTuples(context.Background(), relationTuples...))
}

func TestUsersetRewrites_DocsExample(t *testing.T) {
	ctx := context.Background()

	reg := newDepsProvider(t, []*namespace.Namespace{&docsNS})
	reg.Logger().Logger.SetLevel(logrus.TraceLevel)
	nsMgr, err := reg.Config(ctx).NamespaceManager()
	require.NoError(t, err)
	ns, err := nsMgr.GetNamespaceByName(ctx, docsNS.Name)
	require.NoError(t, err)
	require.Equal(t, &docsNS, ns)

	insertFixtures(t, reg.RelationTupleManager(), []string{
		"docs:document#owner@user",                  // user owns doc
		"docs:doc_in_folder#parent@docs:folder#...", // doc_in_folder is in folder
		"docs:folder#owner@user",                    // user owns folder

		// Folder hierarchy folder_a -> folder_b -> folder_c -> file
		// and folder_a is owned by user. Then user should have access to file.
		"docs:file#parent@docs:folder_c#...",
		"docs:folder_c#parent@docs:folder_b#...",
		"docs:folder_b#parent@docs:folder_a#...",
		"docs:folder_a#owner@user",
	})

	e := check.NewEngine(reg)

	testCases := []struct {
		query   string
		allowed bool
	}{{
		// direct
		query:   "docs:document#owner@user",
		allowed: true,
	}, {
		// userset rewrite
		query:   "docs:document#editor@user",
		allowed: true,
	}, {
		// transitive userset rewrite
		query:   "docs:document#viewer@user",
		allowed: true,
	}, {
		query:   "docs:document#editor@nobody",
		allowed: false,
	}, {
		query:   "docs:folder#viewer@user",
		allowed: true,
	}, {
		// tuple to userset
		query:   "docs:doc_in_folder#viewer@user",
		allowed: true,
	}, {
		// tuple to userset
		query:   "docs:doc_in_folder#viewer@nobody",
		allowed: false,
	}, {
		// tuple to userset
		query:   "docs:another_doc#viewer@user",
		allowed: false,
	}, {
		query:   "docs:file#viewer@user",
		allowed: true,
	}}

	for _, tc := range testCases {
		t.Run(tc.query, func(t *testing.T) {
			rt, err := relationtuple.InternalFromString(tc.query)
			require.NoError(t, err)

			res, err := e.SubjectIsAllowed(ctx, rt, 100)
			require.NoError(t, err)
			assert.Equal(t, tc.allowed, res)
		})
	}
}
