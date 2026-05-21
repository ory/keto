// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/check/trace"
	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/testhelpers"
)

var rewrites = testhelpers.Scenario{
	Name: "userset rewrites",
	Opl: `
	class User implements Namespace {}

	class Doc implements Namespace {
		related: {
			owner: User[]
			parent: Doc[]
		}
		permits = {
			editor: (ctx: Context) => this.related.owner.includes(ctx.subject),
			viewer: (ctx: Context) => this.permits.editor(ctx) ||
				this.related.parent.traverse(p => p.permits.viewer(ctx))
		}
	}

	class Group implements Namespace {
		related: {
			member: User[]
		}
	}

	class Level implements Namespace {
		related: {
			member: User[]
		}
	}

	class Resource implements Namespace {
		related: {
			level: Level[]
			owner: Group[]
		}
		permits = {
			viewer: (ctx: Context) => this.related.owner.traverse(g => g.related.member.includes(ctx.subject)),
			isOwner: (ctx: Context) => this.related.owner.traverse(g => g.related.member.includes(ctx.subject)),

			read: (ctx: Context) => this.permits.viewer(ctx) || this.permits.isOwner(ctx),
			update: (ctx: Context) => this.permits.isOwner(ctx),
			delete: (ctx: Context) => this.permits.isOwner(ctx) && this.related.level.traverse(l => l.related.member.includes(ctx.subject)),
		}
	}

	class Acl implements Namespace {
		related: {
			allow: User[]
			deny: User[]
		}
		permits = {
			access: (ctx: Context) => this.related.allow.includes(ctx.subject) && !this.related.deny.includes(ctx.subject)
		}
	}
	`,
	InputTuples: []string{
		// Direct ownership — plain and namespace-qualified subjects both work.
		"Doc:document#owner@plain_user",
		"Doc:document#owner@User:user",

		// Single-level parent: doc_in_folder inherits viewer from its parent folder.
		"Doc:folder#owner@plain_user",
		"Doc:folder#owner@User:user",
		"Doc:doc_in_folder#parent@Doc:folder",

		// Deep hierarchy: user owns folder_a; viewer propagates down 4 levels to file.
		"Doc:folder_a#owner@user",
		"Doc:folder_b#parent@Doc:folder_a",
		"Doc:folder_c#parent@Doc:folder_b",
		"Doc:file#parent@Doc:folder_c",

		// Group + level gating: delete requires owner-group membership AND correct level.
		"Group:editors#member@mark",
		"Group:editors#member@mike",
		"Level:superadmin#member@mark",
		"Level:superadmin#member@sandy",
		"Resource:topsecret#owner@Group:editors#",
		"Resource:topsecret#level@Level:superadmin#",

		// Allow/deny ACL: deny overrides allow (mallory is on both lists).
		"Acl:document#allow@alice",
		"Acl:document#allow@bob",
		"Acl:document#allow@mallory",
		"Acl:document#deny@mallory",
	},
}

type path []string

func TestUsersetRewrites(t *testing.T) {
	testCases := []struct {
		query         string
		expected      check.Result
		expectedPaths []path
	}{
		{
			// direct
			query:    "Doc:document#owner@User:user",
			expected: check.ResultIsMember,
		},
		{
			// userset rewrite
			query:    "Doc:document#editor@User:user",
			expected: check.ResultIsMember,
		},
		{
			// userset rewrite
			query:    "Doc:document#editor@plain_user",
			expected: check.ResultIsMember,
		},
		{
			// transitive userset rewrite
			query:    "Doc:document#viewer@User:user",
			expected: check.ResultIsMember,
		},
		{
			query:    "Doc:document#editor@nobody",
			expected: check.ResultNotMember,
		},
		{
			query:    "Doc:folder#viewer@User:user",
			expected: check.ResultIsMember,
		},
		{
			// tuple to userset
			query:    "Doc:doc_in_folder#viewer@User:user",
			expected: check.ResultIsMember,
		},
		{
			// tuple to userset
			query:    "Doc:doc_in_folder#viewer@plain_user",
			expected: check.ResultIsMember,
		},
		{
			// tuple to userset
			query:    "Doc:doc_in_folder#viewer@nobody",
			expected: check.ResultNotMember,
		},
		{
			// tuple to userset
			query:    "Doc:another_doc#viewer@user",
			expected: check.ResultNotMember,
		},
		{
			query:    "Doc:file#viewer@user",
			expected: check.ResultIsMember,
		},
		{
			query:    "Level:superadmin#member@mark",
			expected: check.ResultIsMember, // mark is both editor and has correct level
		},
		{
			query:    "Resource:topsecret#isOwner@mark",
			expected: check.ResultIsMember, // mark is both editor and has correct level
		},
		{
			query:    "Resource:topsecret#delete@mark",
			expected: check.ResultIsMember, // mark is both editor and has correct level
			expectedPaths: []path{
				{"*", "Resource:topsecret#delete@mark", "*", "Level:superadmin#member@mark"},
				{"*", "Resource:topsecret#delete@mark", "Resource:topsecret#isOwner@mark", "*", "*", "*", "Group:editors#member@mark"},
			},
		},
		{
			query:    "Resource:topsecret#update@mike",
			expected: check.ResultIsMember, // mike owns the resource
		},
		{
			query:    "Level:superadmin#member@mike",
			expected: check.ResultNotMember, // mike does not have correct level
		},
		{
			query:    "Resource:topsecret#delete@mike",
			expected: check.ResultNotMember, // mike does not have correct level
		},
		{
			query:    "Resource:topsecret#delete@sandy",
			expected: check.ResultNotMember, // sandy is not in the editor Group
		},
		{
			query:         "Acl:document#access@alice",
			expected:      check.ResultIsMember,
			expectedPaths: []path{{"*", "Acl:document#access@alice", "Acl:document#allow@alice"}},
		},
		{
			query:    "Acl:document#access@bob",
			expected: check.ResultIsMember,
		},
		{
			query:    "Acl:document#allow@mallory",
			expected: check.ResultIsMember,
		},
		{
			query:    "Acl:document#access@mallory",
			expected: check.ResultNotMember, // mallory is also on deny-list
		},
	}

	rewrites.Run(t, func(t *testing.T, reg driver.Registry) {
		e := trace.NewEngine(reg)
		defer goleak.VerifyNone(t, goleak.IgnoreCurrent())

		t.Run("suite=testcases", func(t *testing.T) {
			defer goleak.VerifyNone(t, goleak.IgnoreCurrent())

			for _, tc := range testCases {
				t.Run("case="+tc.query, func(t *testing.T) {
					rt := testhelpers.TupleFromString(t, tc.query)

					res, tree := e.CheckRelationTupleWithTrace(t.Context(), rt, 100)
					require.NoError(t, res.Err)
					assert.Equal(t, tc.expected.Membership.String(), res.Membership.String())

					if len(tc.expectedPaths) > 0 {
						for _, path := range tc.expectedPaths {
							assertPath(t, path, tree)
						}
					}
				})
			}
		})
		t.Run("suite=one worker", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			e := trace.NewEngine(reg)

			rt := testhelpers.TupleFromString(t, "Doc:file#viewer@user")
			res, _ := e.CheckRelationTupleWithTrace(ctx, rt, 100)
			require.NoError(t, res.Err)
			assert.Equal(t, check.ResultIsMember.Membership, res.Membership)
		})
	})
}

// assertPath asserts that the given path can be found in the tree.
func assertPath(t *testing.T, path path, tree *trace.Node) {
	require.NotNil(t, tree)
	assert.True(t, hasPath(t, path, tree), "could not find path %s in tree:\n%s", path, tree)
}

func hasPath(t *testing.T, path path, tree *trace.Node) bool {
	if len(path) == 0 {
		return true
	}
	treeLabel := tree.Tuple.String()
	if path[0] != "*" {
		// use testhelpers.TupleFromString to compare against paths with UUIDs.
		tuple := testhelpers.TupleFromString(t, path[0])
		tupleStr := tuple.String()
		if tupleStr != treeLabel {
			return false
		}
	}

	if len(path) == 1 {
		return true
	}

	for _, child := range tree.Children {
		if hasPath(t, path[1:], child) {
			return true
		}
	}
	return false
}
