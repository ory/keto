// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/herodot"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/testhelpers"
	"github.com/ory/keto/ketoapi"
)

// ExpansionRecursive is a three-level group hierarchy where engineering branches
// into two sibling sub-groups. Safe to run with WithStrict(true).
//
// company ⊃ engineering ⊃ (frontend, backend)
// company ⊃ User:cto
// engineering ⊃ User:manager
// backend ⊃ User:Alice, Alice
// frontend ⊃ User:Bob, Bob
var ExpansionRecursive = testhelpers.Scenario{
	Name: "Expansion Recursive",
	Opl: `
		class User implements Namespace{}
		class Group implements Namespace{
			related: {
				members: (User | SubjectSet<Group, "members">)[]
			}
		}
	`,
	InputTuples: []string{
		"Group:company#members@Group:engineering#members",

		"Group:company#members@User:cto",
		"Group:engineering#members@User:manager",

		"Group:engineering#members@Group:frontend#members",
		"Group:engineering#members@Group:backend#members",

		"Group:backend#members@User:Alice",
		"Group:frontend#members@User:Bob",

		"Group:backend#members@Alice", // SubjectID
		"Group:frontend#members@Bob",  // SubjectID
	},
}

var ExpansionRecursiveNamespaceOnly = testhelpers.Scenario{
	Name:        ExpansionRecursive.Name + " (namespaces only)",
	Opl:         `class User implements Namespace{} class Group implements Namespace{}`,
	InputTuples: ExpansionRecursive.InputTuples,
}

// SubjectSetExpansion3Hop combines cross-namespace and within-namespace SubjectSet expansion.
//
// File:team-report#viewers ⊃ File:team-report#editors ⊃ Group:finance#member → (User:Alice, User:Bob).
// File:personal-note#viewers ⊃ User:Alice, User:Bob
var SubjectSetExpansion3Hop = testhelpers.Scenario{
	Name: "SubjectSet Expansion 3 Hops",
	Opl: `
		class User implements Namespace{}
		class Group implements Namespace{
			related: {
				member: User[]
			}
		}
		class File implements Namespace{
			related: {
				editors: (User | SubjectSet<Group, "member">)[]
				viewers: (User | SubjectSet<File, "editors">)[]
			}
		}
	`,
	InputTuples: []string{
		"Group:finance#member@User:Alice",
		"Group:finance#member@User:Bob",

		"File:team-report#editors@Group:finance#member",     // group members are editors
		"File:team-report#viewers@File:team-report#editors", // editors are also viewers

		"File:personal-note#viewers@User:Alice",
		"File:personal-note#viewers@User:Bob",
	},
}

var SubjectSetExpansion3HopNamespacesOnly = testhelpers.Scenario{
	Name:        SubjectSetExpansion3Hop.Name + " (namespaces only)",
	Opl:         `class User implements Namespace{} class Group implements Namespace{} class File implements Namespace{}`,
	InputTuples: SubjectSetExpansion3Hop.InputTuples,
}

// CircularGraphNamespacesOnly covers a fully circular SubjectSet graph with no leaf subjects.
//
// A → B → C → A.
var CircularGraphNamespacesOnly = testhelpers.Scenario{
	Name: "Circular Graph (namespaces only)",
	Opl: `
		class Folder implements Namespace{}
	`,
	InputTuples: []string{
		"Folder:a#parent@Folder:b#parent",
		"Folder:b#parent@Folder:c#parent",
		"Folder:c#parent@Folder:a#parent",
	},
}

// FileWithNonConformingTuples covers a schema where a client incorrectly stores
// tuples directly under permit names (canEdit, canView) instead of the backing
// related fields (editors, viewers).
//
// In strict mode the engine ignores those tuples and resolves access only via
// OPL rewrites. In non-strict mode the direct tuples are visible alongside the
// rewrites.
var FileWithNonConformingTuples = testhelpers.Scenario{
	Name: "File With Non-Conforming Tuples",
	Opl: `
		class User implements Namespace {}
		class File implements Namespace {
			related: {
				editors: User[]
				viewers: User[]
			}
			permits = {
				canEdit: (ctx: Context) => this.related.editors.includes(ctx.subject),
				canView: (ctx: Context) => this.permits.canEdit(ctx) || this.related.viewers.includes(ctx.subject),
			}
		}
	`,
	InputTuples: []string{
		"File:secret-doc#editors@User:Alice",
		"File:secret-doc#viewers@User:Bob",

		// Bad client data: permit names used as direct relation targets.
		"File:secret-doc#canEdit@User:Eve",
		"File:secret-doc#canEdit@Eve", // subjectID
	},
}

// NotionPage models a simplified Notion-like workspace with nested pages.
//
// canView requires direct or parent access.
// canEdit requires direct editor access AND workspace membership.
var NotionPage = testhelpers.Scenario{
	Name: "Notion Page",
	Opl: `
		class User implements Namespace {}
		class Workspace implements Namespace {
			related: {
				member: User[]
			}
		}
		class Page implements Namespace {
			related: {
				workspace: Workspace[]
				parent: Page[]
				viewer: User[]
				editor: User[]
			}
			permits = {
				canView: (ctx: Context) =>
					this.related.viewer.includes(ctx.subject) ||
					this.related.editor.includes(ctx.subject) ||
					this.related.parent.traverse(p => p.permits.canView(ctx)),
				canEdit: (ctx: Context) =>
					this.related.editor.includes(ctx.subject) &&
					this.related.workspace.traverse(w => w.related.member.includes(ctx.subject)),
			}
		}
	`,
	InputTuples: []string{
		"Workspace:acme#member@User:Alice",
		"Workspace:acme#member@User:Bob",

		"Page:home#workspace@Workspace:acme",
		"Page:home#viewer@User:Alice",
		"Page:home#editor@User:Bob",
		"Page:home#editor@User:Eve", // Eve is not workspace member

		"Page:intro#workspace@Workspace:acme",
		"Page:intro#parent@Page:home", // intro inherits canView from home
		"Page:intro#editor@User:Alice",
	},
}

var testcases = []struct {
	scenarios          []testhelpers.Scenario
	expectedMembers    []string
	expectedNonMembers []string
}{
	{
		scenarios: []testhelpers.Scenario{ExpansionRecursive, ExpansionRecursiveNamespaceOnly},
		expectedMembers: []string{
			// direct: SubjectSet and SubjectID members
			"Group:backend#members@User:Alice", "Group:backend#members@Alice",
			"Group:frontend#members@User:Bob", "Group:frontend#members@Bob",

			// engineering: own direct member plus frontend and backend
			"Group:engineering#members@User:manager",
			"Group:engineering#members@User:Alice", "Group:engineering#members@Alice",
			"Group:engineering#members@User:Bob", "Group:engineering#members@Bob",

			// company: own direct member plus everything from engineering
			"Group:company#members@User:cto",
			"Group:company#members@User:manager",
			"Group:company#members@User:Alice", "Group:company#members@Alice",
			"Group:company#members@User:Bob", "Group:company#members@Bob",
		},
		expectedNonMembers: []string{
			// frontend and backend members do not cross.
			"Group:frontend#members@User:Alice",
			"Group:backend#members@User:Bob",

			// Membership flows upward only.
			"Group:backend#members@User:manager", "Group:backend#members@User:cto",
			"Group:frontend#members@User:manager", "Group:frontend#members@User:cto",
			"Group:engineering#members@User:cto",
		},
	},
	{
		scenarios: []testhelpers.Scenario{SubjectSetExpansion3Hop, SubjectSetExpansion3HopNamespacesOnly},
		expectedMembers: []string{
			// direct
			"Group:finance#member@User:Alice", "Group:finance#member@User:Bob",

			// via group expansion
			"File:team-report#editors@User:Alice", "File:team-report#editors@User:Bob",

			// via editors are viewers
			"File:team-report#viewers@User:Alice", "File:team-report#viewers@User:Bob",

			// direct viewers on personal note
			"File:personal-note#viewers@User:Alice", "File:personal-note#viewers@User:Bob",
		},
		expectedNonMembers: []string{
			// File:personal-note is missing subject-set relationship;
			// Therefore editors are not viewers;
			"File:personal-note#editors@User:Alice",
			"File:personal-note#editors@User:Bob",
		},
	},
	{
		scenarios: []testhelpers.Scenario{CircularGraphNamespacesOnly},
		expectedMembers: []string{
			"Folder:a#parent@Folder:a#parent",
			"Folder:a#parent@Folder:b#parent",
			"Folder:a#parent@Folder:c#parent",

			"Folder:b#parent@Folder:a#parent",
			"Folder:b#parent@Folder:b#parent",
			"Folder:b#parent@Folder:c#parent",

			"Folder:c#parent@Folder:a#parent",
			"Folder:c#parent@Folder:b#parent",
			"Folder:c#parent@Folder:c#parent",
		},
	},
	{
		scenarios: []testhelpers.Scenario{{
			Name: "rejects transitive relation",
			Opl: `
						class User implements Namespace{}
						class Directory implements Namespace{}
						class File implements Namespace{}
					`,
			InputTuples: []string{
				"File:report#parent@Directory:docs",
				"Directory:docs#access@User:Alice",
			},
		}},
		expectedNonMembers: []string{
			"File:report#access@User:Alice",
		},
	},
	{
		scenarios: []testhelpers.Scenario{FileWithNonConformingTuples.WithStrict(true)},
		expectedMembers: []string{
			"File:secret-doc#viewers@User:Bob",
			"File:secret-doc#editors@User:Alice",

			// via OPL
			"File:secret-doc#canView@User:Bob",
			"File:secret-doc#canEdit@User:Alice",
			"File:secret-doc#canView@User:Alice",
		},
		expectedNonMembers: []string{
			"File:secret-doc#canEdit@User:Eve",
			"File:secret-doc#canView@User:Eve",
		},
	},
	{
		scenarios: []testhelpers.Scenario{NotionPage},
		expectedMembers: []string{
			// direct viewer/editor on home
			"Page:home#canView@User:Alice",
			"Page:home#canView@User:Bob", // editor also satisfies canView
			"Page:home#canView@User:Eve", // stale editor tuple still satisfies canView (no workspace check)
			// canEdit requires editor AND workspace member
			"Page:home#canEdit@User:Bob",
			// intro: direct editor + workspace member
			"Page:intro#canEdit@User:Alice",
			// intro canView via TTU: intro → parent → home → Alice/Bob/Eve
			"Page:intro#canView@User:Alice",
			"Page:intro#canView@User:Bob",
			"Page:intro#canView@User:Eve",
		},
		expectedNonMembers: []string{
			// Eve has editor tuple but is not a workspace member — AND fails
			"Page:home#canEdit@User:Eve",
			// Bob is editor of home, not intro — canEdit requires explicit editor tuple
			"Page:intro#canEdit@User:Bob",
			// Alice is only a viewer on home, not an editor
			"Page:home#canEdit@User:Alice",
		},
	},
	{
		scenarios: []testhelpers.Scenario{FileWithNonConformingTuples},
		expectedMembers: []string{
			"File:secret-doc#viewers@User:Bob",
			"File:secret-doc#editors@User:Alice",

			// via OPL
			"File:secret-doc#canView@User:Bob",
			"File:secret-doc#canEdit@User:Alice",
			"File:secret-doc#canView@User:Alice",

			// canEdit is a direct tuple, visible in non-strict mode
			"File:secret-doc#canEdit@User:Eve",
			// canView is visible via OPL rewrite in non-strict mode
			"File:secret-doc#canView@User:Eve",
		},
		expectedNonMembers: []string{},
	},
}

func TestEngine(t *testing.T) {
	t.Run(`membership tests`, func(t *testing.T) {
		for _, tc := range testcases {
			for _, scenario := range tc.scenarios {
				scenario.Run(t, func(t *testing.T, reg driver.Registry) {
					e := check.NewEngine(reg)
					ctx := t.Context()

					for _, tuple := range tc.expectedMembers {
						res, err := e.CheckIsMember(ctx, testhelpers.TupleFromString(t, tuple), 5)
						require.NoError(t, err)
						assert.Truef(t, res, "expected %q to be a member", tuple)
					}

					for _, tuple := range tc.expectedNonMembers {
						res, err := e.CheckIsMember(ctx, testhelpers.TupleFromString(t, tuple), 5)
						require.NoError(t, err)
						assert.Falsef(t, res, "expected %q to not be a member", tuple)
					}
				})
			}
		}
	})

	t.Run("respects max-depth", func(t *testing.T) {
		ExpansionRecursive.Run(t, func(t *testing.T, reg driver.Registry) {
			e := check.NewEngine(reg)
			ctx := t.Context()

			// company → engineering → backend → User:Alice requires exactly 3 hops.
			tuple := testhelpers.TupleFromString(t, "Group:company#members@User:Alice")

			assert.Equal(t, 5, reg.Config(ctx).MaxReadDepth())

			res, err := e.CheckIsMember(ctx, tuple, 2)
			require.NoError(t, err)
			assert.False(t, res, "depth=2 is not enough for a 3-hop chain")

			res, err = e.CheckIsMember(ctx, tuple, 3)
			require.NoError(t, err)
			assert.True(t, res, "depth=3 is enough for a 3-hop chain")

			require.NoError(t, reg.Config(ctx).Set(config.KeyLimitMaxReadDepth, 2))
			res, err = e.CheckIsMember(ctx, tuple, 5)
			require.NoError(t, err)
			assert.False(t, res, "global max-depth=2 overrides request depth=5")

			require.NoError(t, reg.Config(ctx).Set(config.KeyLimitMaxReadDepth, 3))
			res, err = e.CheckIsMember(ctx, tuple, 0)
			require.NoError(t, err)
			assert.True(t, res, "global max-depth=3 is enough")
		})
	})

	t.Run("batch check", func(t *testing.T) {
		SubjectSetExpansion3Hop.Run(t, func(t *testing.T, reg driver.Registry) {
			e := check.NewEngine(reg)
			ctx := t.Context()

			targetTuples := []*ketoapi.RelationTuple{
				testhelpers.APITupleFromString(t, "Group:finance#member@User:Alice"),                   // direct
				testhelpers.APITupleFromString(t, "File:team-report#editors@User:Alice"),               // 2-hop
				testhelpers.APITupleFromString(t, "File:team-report#viewers@User:Alice"),               // 3-hop
				testhelpers.APITupleFromString(t, "NonExistent:x#rel@User:Alice"),                      // non-existent namespace
				testhelpers.APITupleFromString(t, "Group:finance#member@User:Eve"),                     // unknown subject
				testhelpers.APITupleFromString(t, "File:team-report#viewers@File:team-report#editors"), // via subject set
			}

			// At depth=2, the 3-hop viewers chain is not resolved.
			results, err := e.BatchCheck(ctx, targetTuples, 2)
			require.NoError(t, err)

			require.Equal(t, check.IsMember, results[0].Membership)
			require.Equal(t, check.IsMember, results[1].Membership)
			require.Equal(t, check.NotMember, results[2].Membership)
			require.Equal(t, check.MembershipUnknown, results[3].Membership)
			require.EqualError(t, results[3].Err, herodot.ErrNotFound().Error())
			require.Equal(t, check.NotMember, results[4].Membership)
			require.Equal(t, check.IsMember, results[5].Membership)

			// At depth=3, the 3-hop viewers chain resolves
			results, err = e.BatchCheck(ctx, targetTuples, 3)
			require.NoError(t, err)
			require.Equal(t, check.IsMember, results[2].Membership)
		})
	})
}

// Alice is in both deep and short chains.
// Depth of 4 is needed to reach Alice via shortRel.
// Depth of 6 is needed to reach Alice via deepRel.
var depthLimitScenario = testhelpers.Scenario{
	Name: "Depth Limit",
	Opl: `
		class User implements Namespace {}
		class Group implements Namespace {
			related: {
				members: (User | SubjectSet<Group, "members">)[]
			}
		}
		class Resource implements Namespace {
			related: {
				deepRel: (User | SubjectSet<Group, "members">)[]
				shortRel: (User | SubjectSet<Group, "members">)[]
			}
			permits = {
				union: (ctx:Context) => 
						this.related.deepRel.includes(ctx.subject) 
						|| this.related.shortRel.includes(ctx.subject),
				intersection: (ctx:Context) => 
						this.related.deepRel.includes(ctx.subject) 
						&& this.related.shortRel.includes(ctx.subject),
				shortButNotDeep: (ctx:Context) => this.related.shortRel.includes(ctx.subject) 
						&& !this.related.deepRel.includes(ctx.subject),

				// notEither simulates a client that authorizes users when Keto returns NotMember
				notEither: (ctx:Context) => !this.permits.union(ctx),
			}
		}
	`,
	InputTuples: []string{
		"Resource:doc#deepRel@Group:deep#members",
		"Resource:doc#shortRel@Group:short#members",
		"Group:deep#members@Group:deep1#members",
		"Group:deep1#members@Group:deep2#members",
		"Group:deep2#members@Group:deep3#members",
		"Group:deep3#members@Group:deep4#members",
		"Group:deep4#members@User:Alice",
		"Group:short#members@Group:short1#members",
		"Group:short1#members@User:Alice",
	},
	Strict: false,
}

func TestDepthLimit(t *testing.T) {
	type tc struct {
		name               string
		tuple              string
		depth              int
		expectedMembership check.Membership
		expectedIsMember   bool
		expectedErr        error
	}

	tests := []tc{
		{
			name:               "union: too shallow depth for both branches yields NotMember",
			tuple:              "Resource:doc#union@User:Alice",
			depth:              2,
			expectedMembership: check.NotMember,
			expectedIsMember:   false,
		},
		{
			name:               "union: enough depth for 1 branch yields Member",
			tuple:              "Resource:doc#union@User:Alice",
			depth:              4,
			expectedMembership: check.IsMember,
			expectedIsMember:   true,
		},
		{
			name:               "intersection: too shallow depth for both branches yields NotMember",
			tuple:              "Resource:doc#intersection@User:Alice",
			depth:              2,
			expectedMembership: check.NotMember,
			expectedIsMember:   false,
		},
		{
			name:               "intersection: enough depth for only 1 branch yields NotMember",
			tuple:              "Resource:doc#intersection@User:Alice",
			depth:              4,
			expectedMembership: check.NotMember,
			expectedIsMember:   false,
		},
		{
			name:               "intersection: enough depth for both branches yields Member",
			tuple:              "Resource:doc#intersection@User:Alice",
			depth:              6,
			expectedMembership: check.IsMember,
			expectedIsMember:   true,
		},

		{
			name:               "shortButNotDeep: too shallow depth for both branches yields NotMember",
			tuple:              "Resource:doc#shortButNotDeep@User:Alice",
			depth:              2,
			expectedMembership: check.NotMember,
			expectedIsMember:   false,
		},
		{
			name:               "shortButNotDeep: enough depth for only 1 branch yields NotMember",
			tuple:              "Resource:doc#shortButNotDeep@User:Alice",
			depth:              4,
			expectedMembership: check.IsMember,
			expectedIsMember:   true,
		},
		{
			name:               "notEither: not enough depth for any branch yields Member",
			tuple:              "Resource:doc#notEither@User:Alice",
			depth:              2,
			expectedMembership: check.IsMember,
			expectedIsMember:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			depthLimitScenario.Run(t, func(t *testing.T, reg driver.Registry) {
				ctx := t.Context()
				e := check.NewEngine(reg)

				if tc.depth > 0 {
					require.NoError(t, reg.Config(ctx).Set(config.KeyLimitMaxReadDepth, tc.depth))
				}

				res := e.CheckRelationTuple(ctx, testhelpers.TupleFromString(t, tc.tuple), 10)

				require.NoError(t, res.Err)
				assert.Equal(t, tc.expectedMembership, res.Membership)
				if tc.expectedErr != nil {
					assert.EqualError(t, res.Err, tc.expectedErr.Error())
				} else {
					assert.NoError(t, res.Err)
				}

				assert.Equal(t, tc.expectedIsMember, res.Membership == check.IsMember)
			})
		})
	}
}
