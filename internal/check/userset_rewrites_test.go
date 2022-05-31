package check_test

import (
	"context"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/check/checkgroup"
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

func pathFromString(t *testing.T, s string) (path checkgroup.Path) {
	for _, el := range strings.Split(s, "->") {
		tuple, trans, found := strings.Cut(el, " as ")
		if !found {
			t.Fatalf("could not parse path from %q", s)
			return
		}
		tuple = strings.TrimSpace(tuple)
		trans = strings.TrimSpace(trans)
		rt, err := relationtuple.InternalFromString(tuple)
		if err != nil {
			t.Fatalf("could not parse tuple from string %q: %v", s, err)
			return
		}
		path.Edges = append(path.Edges, checkgroup.Edge{
			Tuple:          *rt,
			Transformation: checkgroup.TransformationFromString(trans),
		})
	}

	return
}

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
		query    string
		expected checkgroup.Result
	}{{
		// direct
		query: "doc:document#owner@user",
		expected: checkgroup.Result{
			Membership: checkgroup.IsMember,
			Path:       pathFromString(t, "doc:document#owner@user as direct"),
		},
	}, {
		// userset rewrite
		query: "doc:document#editor@user",
		expected: checkgroup.Result{
			Membership: checkgroup.IsMember,
			Path:       pathFromString(t, "doc:document#editor@user as computed-userset -> doc:document#owner@user as direct"),
		},
	}, {
		// transitive userset rewrite
		query: "doc:document#viewer@user",
		expected: checkgroup.Result{
			Membership: checkgroup.IsMember,
			Path:       pathFromString(t, "doc:document#viewer@user as computed-userset -> doc:document#editor@user as computed-userset -> doc:document#owner@user as direct"),
		},
	}, {
		query:    "doc:document#editor@nobody",
		expected: checkgroup.ResultNotMember,
	}, {
		query: "doc:folder#viewer@user",
		expected: checkgroup.Result{
			Membership: checkgroup.IsMember,
			Path:       pathFromString(t, "doc:folder#viewer@user as computed-userset -> doc:folder#editor@user as computed-userset -> doc:folder#owner@user as direct"),
		},
	}, {
		// tuple to userset
		query: "doc:doc_in_folder#viewer@user",
		expected: checkgroup.Result{
			Membership: checkgroup.IsMember,
			Path: pathFromString(t, `
			doc:doc_in_folder#viewer@user as tuple-to-userset ->
			doc:folder#viewer@user as computed-userset -> 
			doc:folder#editor@user as computed-userset -> 
			doc:folder#owner@user as direct`),
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
		query: "doc:file#viewer@user",
		expected: checkgroup.Result{
			Membership: checkgroup.IsMember,
			Path: pathFromString(t, `
			doc:file#viewer@user as tuple-to-userset ->
			doc:folder_c#viewer@user as tuple-to-userset -> 
			doc:folder_b#viewer@user as tuple-to-userset -> 
			doc:folder_a#viewer@user as computed-userset -> 
			doc:folder_a#editor@user as computed-userset -> 
			doc:folder_a#owner@user as direct`),
		},
	}, {
		query:    "level:superadmin#member@mark",
		expected: checkgroup.ResultIsMember, // mark is both editor and has correct level
	}, {
		query:    "resource:topsecret#owner@mark",
		expected: checkgroup.ResultIsMember, // mark is both editor and has correct level
	}, {
		query:    "resource:topsecret#delete@mark",
		expected: checkgroup.ResultIsMember, // mark is both editor and has correct level
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
		query:    "acl:document#access@alice",
		expected: checkgroup.ResultIsMember,
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
			assert.Equal(t, tc.expected.Membership, res.Membership)
			if len(tc.expected.Path.Edges) > 0 {
				assert.Equalf(t, tc.expected.Path, res.Path, "\nwant: %s\ngot:  %s\n", &tc.expected.Path, &res.Path)
			}
		})
	}
}
