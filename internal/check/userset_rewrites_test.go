package check_test

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/check/checkgroup"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/namespace/ast"
	"github.com/ory/keto/internal/relationtuple"
)

var namespaces = []*namespace.Namespace{
	{Name: "doc",
		ID: 1,
		Relations: []ast.Relation{
			{
				Name: "owner"},
			{
				Name: "editor",
				UsersetRewrite: &ast.UsersetRewrite{
					Children: ast.Children{ast.ComputedUserset{
						Relation: "owner"}}}},
			{
				Name: "viewer",
				UsersetRewrite: &ast.UsersetRewrite{
					Children: ast.Children{
						ast.ComputedUserset{
							Relation: "editor"},
						ast.TupleToUserset{
							Relation:                "parent",
							ComputedUsersetRelation: "viewer"}}}},
		}},
	{Name: "group",
		ID:        2,
		Relations: []ast.Relation{{Name: "member"}},
	},
	{Name: "level",
		ID:        3,
		Relations: []ast.Relation{{Name: "member"}},
	},
	{Name: "resource",
		ID: 4,
		Relations: []ast.Relation{
			{Name: "level"},
			{Name: "viewer",
				UsersetRewrite: &ast.UsersetRewrite{
					Children: ast.Children{
						ast.TupleToUserset{Relation: "owner", ComputedUsersetRelation: "member"}}}},
			{Name: "owner",
				UsersetRewrite: &ast.UsersetRewrite{
					Children: ast.Children{
						ast.TupleToUserset{Relation: "owner", ComputedUsersetRelation: "member"}}}},
			{Name: "read",
				UsersetRewrite: &ast.UsersetRewrite{
					Children: ast.Children{
						ast.ComputedUserset{Relation: "viewer"},
						ast.ComputedUserset{Relation: "owner"}}}},
			{Name: "update",
				UsersetRewrite: &ast.UsersetRewrite{
					Children: ast.Children{
						ast.ComputedUserset{Relation: "owner"}}}},
			{Name: "delete",
				UsersetRewrite: &ast.UsersetRewrite{
					Operation: ast.SetOperationIntersection,
					Children: ast.Children{
						ast.ComputedUserset{Relation: "owner"},
						ast.TupleToUserset{
							Relation:                "level",
							ComputedUsersetRelation: "member"}}}},
		}},
	{Name: "acl",
		ID: 5,
		Relations: []ast.Relation{
			{Name: "allow"},
			{Name: "deny"},
			{Name: "access",
				UsersetRewrite: &ast.UsersetRewrite{
					Operation: ast.SetOperationDifference,
					Children: ast.Children{
						ast.ComputedUserset{Relation: "allow"},
						ast.ComputedUserset{Relation: "deny"}}}}}},
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

type path []string

func TestUsersetRewrites(t *testing.T) {
	ctx := context.Background()

	reg := newDepsProvider(t, namespaces)
	reg.Logger().Logger.SetLevel(logrus.TraceLevel)

	insertFixtures(t, reg.RelationTupleManager(), []string{
		"doc:document#owner@user",                 // user owns doc
		"doc:doc_in_folder#parent@doc:folder#...", // doc_in_folder is in folder
		"doc:folder#owner@user",                   // user owns folder

		// Folder hierarchy folder_a -> folder_b -> folder_c -> file
		// and folder_a is owned by user. Then user should have access to file.
		"doc:file#parent@doc:folder_c#...",
		"doc:folder_c#parent@doc:folder_b#...",
		"doc:folder_b#parent@doc:folder_a#...",
		"doc:folder_a#owner@user",

		"group:editors#member@mark",
		"level:superadmin#member@mark",
		"level:superadmin#member@sandy",
		"resource:topsecret#owner@group:editors#...",
		"resource:topsecret#level@level:superadmin#...",
		"resource:topsecret#owner@mike",

		"acl:document#allow@alice",
		"acl:document#allow@bob",
		"acl:document#allow@mallory",
		"acl:document#deny@mallory",
	})

	e := check.NewEngine(reg)

	testCases := []struct {
		query         string
		expected      checkgroup.Result
		expectedPaths []path
	}{{
		// direct
		query: "doc:document#owner@user",
		expected: checkgroup.Result{
			Membership: checkgroup.IsMember,
		},
	}, {
		// userset rewrite
		query: "doc:document#editor@user",
		expected: checkgroup.Result{
			Membership: checkgroup.IsMember,
		},
	}, {
		// transitive userset rewrite
		query: "doc:document#viewer@user",
		expected: checkgroup.Result{
			Membership: checkgroup.IsMember,
		},
	}, {
		query:    "doc:document#editor@nobody",
		expected: checkgroup.ResultNotMember,
	}, {
		query: "doc:folder#viewer@user",
		expected: checkgroup.Result{
			Membership: checkgroup.IsMember,
		},
	}, {
		// tuple to userset
		query: "doc:doc_in_folder#viewer@user",
		expected: checkgroup.Result{
			Membership: checkgroup.IsMember,
		},
	}, {
		// tuple to userset
		query:    "doc:doc_in_folder#viewer@nobody",
		expected: checkgroup.ResultNotMember,
	}, {
		// tuple to userset
		query:    "doc:another_doc#viewer@user",
		expected: checkgroup.ResultNotMember,
	}, {
		query:    "doc:file#viewer@user",
		expected: checkgroup.ResultIsMember,
	}, {
		query:    "level:superadmin#member@mark",
		expected: checkgroup.ResultIsMember, // mark is both editor and has correct level
	}, {
		query:    "resource:topsecret#owner@mark",
		expected: checkgroup.ResultIsMember, // mark is both editor and has correct level
	}, {
		query:    "resource:topsecret#delete@mark",
		expected: checkgroup.ResultIsMember, // mark is both editor and has correct level
		expectedPaths: []path{
			{"*", "resource:topsecret#delete@mark", "level:superadmin#member@mark"},
			{"*", "resource:topsecret#delete@mark", "resource:topsecret#owner@mark", "group:editors#member@mark"},
		},
	}, {
		query:    "resource:topsecret#update@mike",
		expected: checkgroup.ResultIsMember, // mike owns the resource
	}, {
		query:    "level:superadmin#member@mike",
		expected: checkgroup.ResultNotMember, // mike does not have correct level
	}, {
		query:    "resource:topsecret#delete@mike",
		expected: checkgroup.ResultNotMember, // mike does not have correct level
	}, {
		query:    "resource:topsecret#delete@sandy",
		expected: checkgroup.ResultNotMember, // sandy is not in the editor group
	}, {
		query:         "acl:document#access@alice",
		expected:      checkgroup.ResultIsMember,
		expectedPaths: []path{{"*", "acl:document#access@alice", "acl:document#allow@alice"}},
	}, {
		query:    "acl:document#access@bob",
		expected: checkgroup.ResultIsMember,
	}, {
		query:    "acl:document#allow@mallory",
		expected: checkgroup.ResultIsMember,
	}, {
		query:    "acl:document#access@mallory",
		expected: checkgroup.ResultNotMember, // mallory is also on deny-list
	}}

	for _, tc := range testCases {
		t.Run(tc.query, func(t *testing.T) {
			rt, err := relationtuple.InternalFromString(tc.query)
			require.NoError(t, err)

			res := e.Check(ctx, rt, 100)
			assert.Equal(t, tc.expected.Err, res.Err)
			t.Logf("tree:\n%s", res.Tree)
			assert.Equal(t, tc.expected.Membership, res.Membership)

			if len(tc.expectedPaths) > 0 {
				for _, path := range tc.expectedPaths {
					assertPath(t, path, res.Tree)
				}
			}
		})
	}
}

// assertPath asserts that the given path can be found in the tree.
func assertPath(t *testing.T, path path, tree *expand.Tree) {
	require.NotNil(t, tree)
	assert.True(t, hasPath(path, tree), "could not find path %s in tree:\n%s", path, tree)
}

func hasPath(path path, tree *expand.Tree) bool {
	if len(path) == 0 {
		return true
	}
	treeLabel := tree.Label()
	if path[0] != "*" && path[0] != treeLabel {
		return false
	}

	if len(path) == 1 {
		return true
	}

	for _, child := range tree.Children {
		if hasPath(path[1:], child) {
			return true
		}
	}
	return false
}
