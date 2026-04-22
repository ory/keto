// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package expand_test

import (
	"context"
	"slices"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	keysetpagination "github.com/ory/x/pagination/keysetpagination_v2"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/testhelpers"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoapi"
	"github.com/ory/keto/schema"
)

func TestEngine(t *testing.T) {
	emptyNs := driver.WithNamespaces([]*namespace.Namespace{{}})

	t.Run("case=returns SubjectID on expand", func(t *testing.T) {
		groupObj := uuid.Must(uuid.NewV4())
		tuple := &relationtuple.RelationTuple{
			Namespace: "Group",
			Object:    groupObj,
			Relation:  "rel",
			Subject:   &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
		}

		reg := driver.NewSqliteTestRegistry(t, false)
		e := expand.NewEngine(reg)
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), tuple))

		t.Run(`tuple does not exist, but ns:obj exists`, func(t *testing.T) {
			tree, err := e.BuildTree(context.Background(), &relationtuple.SubjectSet{
				Namespace: "Group",
				Object:    groupObj,
				Relation:  "rel2",
			}, 100)
			require.NoError(t, err)

			expectedTree := &relationtuple.Tree{
				Type: ketoapi.TreeNodeUnion,
				Subject: &relationtuple.SubjectSet{
					Namespace: "Group",
					Object:    groupObj,
					Relation:  "rel2",
				},
				Children: nil,
			}
			assert.Equal(t, expectedTree, tree)
		})

		t.Run(`no such ns:obj exists at all`, func(t *testing.T) {
			tree, err := e.BuildTree(context.Background(), &relationtuple.SubjectSet{
				Namespace: "Folder",
				Object:    groupObj,
				Relation:  "rel",
			}, 100)
			require.NoError(t, err)

			expectedTree := &relationtuple.Tree{
				Type: ketoapi.TreeNodeUnion,
				Subject: &relationtuple.SubjectSet{
					Namespace: "Folder",
					Object:    groupObj,
					Relation:  "rel",
				},
				Children: nil,
			}
			assert.Equal(t, expectedTree, tree)
		})
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
		reg := driver.NewSqliteTestRegistry(t, false, emptyNs)
		e := expand.NewEngine(reg)

		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), boulderers...))

		tree, err := e.BuildTree(context.Background(), bouldererUserSet, 100)
		require.NoError(t, err)

		require.True(t, expand.AssertInternalTreesAreEqual(t, &relationtuple.Tree{
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
		}, tree))
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

		reg := driver.NewSqliteTestRegistry(t, false, emptyNs)
		e := expand.NewEngine(reg)

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

		actualTree, err := e.BuildTree(context.Background(), expectedTree.Subject.(*relationtuple.SubjectSet), 100)
		require.NoError(t, err)
		require.True(t, expand.AssertInternalTreesAreEqual(t, expectedTree, actualTree))
	})

	t.Run("case=respects max depth", func(t *testing.T) {
		reg := driver.NewSqliteTestRegistry(t, false, emptyNs)
		e := expand.NewEngine(reg)

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
									Type: ketoapi.TreeNodeUnion,
									Subject: &relationtuple.SubjectSet{
										Object:   ids[3],
										Relation: "child",
									},
									Truncation: &relationtuple.Truncation{
										Reason: relationtuple.TruncationReasonDepthLimit,
										Cursor: &relationtuple.ExpandCursor{
											Kind: relationtuple.ExpandCursorKindDirect,
											SubjectSet: &relationtuple.SubjectSet{
												Object:   ids[3],
												Relation: "child",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		actualTree, err := e.BuildTree(context.Background(), expectedTree.Subject.(*relationtuple.SubjectSet), 4)
		require.NoError(t, err)

		assert.Equal(t, expectedTree, actualTree)
	})

	t.Run("case=paginates", func(t *testing.T) {
		innerReg := driver.NewSqliteTestRegistry(t, false, emptyNs)
		reg, mw := testhelpers.RegistryWithManagerWrapper(t, innerReg, keysetpagination.WithSize(2))

		e := expand.NewEngine(reg)
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
		assert.Len(t, mw.RequestedPages, 2)
	})

	t.Run("case=handles subject sets as leaf", func(t *testing.T) {
		reg := driver.NewSqliteTestRegistry(t, false, emptyNs)
		e := expand.NewEngine(reg)

		expectedTree := &relationtuple.Tree{
			Type: ketoapi.TreeNodeUnion,
			Subject: &relationtuple.SubjectSet{
				Object:   uuid.Must(uuid.NewV4()),
				Relation: "rel",
			},
			Children: []*relationtuple.Tree{
				{
					Type: ketoapi.TreeNodeUnion,
					Subject: &relationtuple.SubjectSet{
						Object:   uuid.Must(uuid.NewV4()),
						Relation: "sr",
					},
				},
			},
		}

		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &relationtuple.RelationTuple{
			Object:   expectedTree.Subject.(*relationtuple.SubjectSet).Object,
			Relation: expectedTree.Subject.(*relationtuple.SubjectSet).Relation,
			Subject:  expectedTree.Children[0].Subject,
		}))

		tree, err := e.BuildTree(context.Background(), expectedTree.Subject.(*relationtuple.SubjectSet), 100)
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

		reg := driver.NewSqliteTestRegistry(t, false, driver.WithNamespaces([]*namespace.Namespace{{Name: namesp}}))
		e := expand.NewEngine(reg)

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
									Type:    ketoapi.TreeNodeUnion,
									Subject: sendlingerTorSS,
									Truncation: &relationtuple.Truncation{
										Reason: relationtuple.TruncationReasonCycle,
									},
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
		reg := driver.NewSqliteTestRegistry(t, false, emptyNs)
		e := expand.NewEngine(reg)
		ss := &relationtuple.SubjectSet{
			Namespace: "unknown",
			Object:    uuid.Must(uuid.NewV4()),
			Relation:  "rel",
		}
		tree, err := e.BuildTree(context.Background(), ss, 100)
		require.NoError(t, err)
		assert.Equal(t, &relationtuple.Tree{
			Type:    ketoapi.TreeNodeUnion,
			Subject: ss,
		}, tree)
	})
}

func TestEngineOpl(t *testing.T) {
	type expandTest struct {
		name        string
		opl         string
		inputTuples []string
		expandInput string
		expected    *relationtuple.Tree
		strict      bool
		maxNodes    int
		only        bool
	}
	tests := []expandTest{
		{
			name:   "case=Multile levels of union and exclusion",
			strict: true,
			opl: `
			class User implements Namespace {}
			class Corporate implements Namespace {
				related: {
					releasers: User[]
				}
			}					
			class UserGroup implements Namespace {
				related: {
					members: User[]
				}
			}
			class AccountGroup implements Namespace {
				related: {
					managers: SubjectSet<UserGroup, "members">[]
				}
			}
			class Account implements Namespace {
				related: {
					managers: SubjectSet<AccountGroup, "managers">[]
				}
			}
			class Task implements Namespace {
				related: {
					managers: (SubjectSet<Account, "managers"> | SubjectSet<Corporate, "releasers">)[]
					
					blocklist: SubjectSet<UserGroup, "members">[]
				}
				permits = {
					viewer: (ctx: Context) => this.related.managers.includes(ctx.subject) && !this.related.blocklist.includes(ctx.subject)
				}
			}
			`,

			inputTuples: []string{
				"UserGroup:UG1#members@User:1",
				"UserGroup:UG1#members@User:2",
				"UserGroup:UG2#members@User:2",
				"UserGroup:UG2#members@User:3",
				"UserGroup:UG2#members@User:4",
				"UserGroup:UG3#members@User:4",
				"UserGroup:UG3#members@User:5",

				"AccountGroup:AG1#managers@UserGroup:UG1#members",
				"AccountGroup:AG1#managers@UserGroup:UG2#members",
				"AccountGroup:AG2#managers@UserGroup:UG3#members",

				"Account:A1#managers@AccountGroup:AG1#managers",
				"Account:A1#managers@AccountGroup:AG2#managers",

				"Task:T1#managers@Account:A1#managers",

				// thru corporate
				"Task:T1#managers@Corporate:C1#releasers",

				// blocklist
				"Task:T1#blocklist@UserGroup:UG-1#members",
				"UserGroup:UG-1#members@User:5",
			},
			expandInput: "Task:T1#viewer",
			maxNodes:    20,
			expected: treeNode(t, intersection, "Task:T1#viewer", nil,
				treeNode(t, union, "Task:T1#managers", nil,
					treeNode(t, union, "Account:A1#managers", nil,
						treeNode(t, union, "AccountGroup:AG1#managers", nil,
							treeNode(t, union, "UserGroup:UG1#members", nil,
								treeNode(t, leaf, "User:1", nil),
								treeNode(t, leaf, "User:2", nil),
							),
							treeNode(t, union, "UserGroup:UG2#members", nil,
								treeNode(t, leaf, "User:2", nil),
								treeNode(t, leaf, "User:3", nil),
								treeNode(t, leaf, "User:4", nil),
							),
						),
						treeNode(t, union, "AccountGroup:AG2#managers", nil,
							treeNode(t, union, "UserGroup:UG3#members", nil,
								treeNode(t, leaf, "User:4", nil),
								treeNode(t, leaf, "User:5", nil),
							),
						),
					),
					treeNode(t, union, "Corporate:C1#releasers", nil),
				),
				treeNode(t, exclusion, "", nil,
					treeNode(t, union, "Task:T1#blocklist", nil,
						treeNode(t, union, "UserGroup:UG-1#members", nil,
							treeNode(t, leaf, "User:5", nil),
						),
					),
				),
			),
		},
		{
			name: "case=expands 1 level without OPL",
			opl:  "",
			inputTuples: []string{
				"Group:parent#users@User:1",
				"Group:parent#users@User:2",
				"Group:parent#users@User:3",
				"Group:parent#users@User4",
				"Group:parent#users@ApiKey:x",
			},
			expandInput: "Group:parent#users",
			expected: treeNode(t, union, "Group:parent#users", nil,
				treeNode(t, leaf, "User:1", nil),
				treeNode(t, leaf, "User:2", nil),
				treeNode(t, leaf, "User:3", nil),
				treeNode(t, leaf, "User4", nil),
				treeNode(t, leaf, "ApiKey:x", nil),
			),
		},
		{
			name: "case=expands simple ComputedSubjectSet level with OPL",
			opl: `
				class User implements Namespace{}
				class File implements Namespace{
					related: {
						viewers: User[]
					}
					permits = {
						view: (ctx: Context) => this.related.viewers.includes(ctx.subject)
					}
				}
			`,
			inputTuples: []string{
				"File:f1#viewers@User:1",
				"File:f1#viewers@User:2",
				"File:f1#viewers@User:3",
			},
			expandInput: "File:f1#view",
			expected: treeNode(t, union, "File:f1#view", nil,
				treeNode(t, union, "File:f1#viewers", nil,
					treeNode(t, leaf, "User:1", nil),
					treeNode(t, leaf, "User:2", nil),
					treeNode(t, leaf, "User:3", nil),
				),
			),
		},
		{
			name:   "case=expands simple Inverted ComputedSubjectSet level with OPL",
			strict: true,
			opl: `
				class User implements Namespace{}
				class File implements Namespace{
					related: {
						viewers: User[]
					}
					permits = {
						view: (ctx: Context) => !this.related.viewers.includes(ctx.subject)
					}
				}
			`,
			inputTuples: []string{
				"File:f1#viewers@User:1",
				"File:f1#viewers@User:2",
				"File:f1#viewers@User:3",
			},
			expandInput: "File:f1#view",
			expected: treeNode(t, union, "File:f1#view", nil,
				treeNode(t, exclusion, "", nil,
					treeNode(t, union, "File:f1#viewers", nil,
						treeNode(t, leaf, "User:1", nil),
						treeNode(t, leaf, "User:2", nil),
						treeNode(t, leaf, "User:3", nil),
					),
				),
			),
		},
		{
			name: "case=expands blocklistexample",
			opl: `
				class User implements Namespace{}
				class File implements Namespace{
					related: {
						viewers: User[]
						blocklist: User[]
					}
					permits = {
						view: (ctx: Context) => this.related.viewers.includes(ctx.subject) && !this.related.blocklist.includes(ctx.subject)
					}
				}
			`,
			inputTuples: []string{
				"File:f1#viewers@User:1",
				"File:f1#viewers@User:2",
				"File:f1#blocklist@User:1",
			},
			expandInput: "File:f1#view",
			expected: treeNode(t, intersection, "File:f1#view", nil,
				treeNode(t, union, "File:f1#viewers", nil,
					treeNode(t, leaf, "User:1", nil),
					treeNode(t, leaf, "User:2", nil),
				),
				treeNode(t, exclusion, "", nil,
					treeNode(t, union, "File:f1#blocklist", nil,
						treeNode(t, leaf, "User:1", nil),
					),
				),
			),
		},
		{
			name: "case=expands 2 blocklists example",
			opl: `
				class User implements Namespace{}
				class File implements Namespace{
					related: {
						viewers: User[]
						blocklist1: User[]
						blocklist2: User[]
					}
					permits = {
						view: (ctx: Context) => 
							this.related.viewers.includes(ctx.subject) &&
							!(
								this.related.blocklist1.includes(ctx.subject) &&
								this.related.blocklist2.includes(ctx.subject)
							)
					}
				}
			`,
			inputTuples: []string{
				"File:f1#viewers@User:1",
				"File:f1#blocklist1@User:1",
				"File:f1#blocklist2@User:2",
			},
			expandInput: "File:f1#view",
			expected: treeNode(t, intersection, "File:f1#view", nil,
				treeNode(t, union, "File:f1#viewers", nil,
					treeNode(t, leaf, "User:1", nil),
				),
				treeNode(t, exclusion, "", nil,
					treeNode(t, intersection, "", nil,
						treeNode(t, union, "File:f1#blocklist2", nil,
							treeNode(t, leaf, "User:2", nil),
						),
						treeNode(t, union, "File:f1#blocklist1", nil,
							treeNode(t, leaf, "User:1", nil),
						),
					),
				),
			),
		},
		{
			name: "case=expands (includes or traverse)",
			opl: `
				class User implements Namespace{}
				class Group implements Namespace{
					related: {
						members: User[]
					}
				}
				class File implements Namespace{
					related: {
						viewers: User[]
						viewerGroups: Group[]
					}
					permits = {
						view: (ctx: Context) => this.related.viewers.includes(ctx.subject) ||
							this.related.viewerGroups.traverse(g => g.related.members.includes(ctx.subject))
					}
				}
			`,
			inputTuples: []string{
				"File:f1#viewers@User:1",
				"File:f1#viewers@User:2",
				"File:f1#viewers@User:3",

				"Group:g1#members@User:4",
				"File:f1#viewerGroups@Group:g1",
			},
			expandInput: "File:f1#view",
			expected: treeNode(t, union, "File:f1#view", nil,
				treeNode(t, union, "File:f1#viewers", nil,
					treeNode(t, leaf, "User:1", nil),
					treeNode(t, leaf, "User:2", nil),
					treeNode(t, leaf, "User:3", nil),
				),
				treeNode(t, union, "File:f1#viewerGroups", nil,
					treeNode(t, union, "Group:g1#members", nil,
						treeNode(t, leaf, "User:4", nil),
					),
				),
			),
		},
		{
			name: "case=expands (includes or traverse) - traverse group has no member tuples",
			opl: `
				class User implements Namespace{}
				class Group implements Namespace{
					related: {
						members: User[]
					}
				}
				class File implements Namespace{
					related: {
						viewers: User[]
						viewerGroups: Group[]
					}
					permits = {
						view: (ctx: Context) => this.related.viewers.includes(ctx.subject) ||
							this.related.viewerGroups.traverse(g => g.related.members.includes(ctx.subject))
					}
				}
			`,
			inputTuples: []string{
				"File:f1#viewers@User:1",
				"File:f1#viewers@User:2",
				"File:f1#viewers@User:3",

				"File:f1#viewerGroups@Group:g1",
			},
			expandInput: "File:f1#view",
			expected: treeNode(t, union, "File:f1#view", nil,
				treeNode(t, union, "File:f1#viewers", nil,
					treeNode(t, leaf, "User:1", nil),
					treeNode(t, leaf, "User:2", nil),
					treeNode(t, leaf, "User:3", nil),
				),
				treeNode(t, union, "File:f1#viewerGroups", nil,
					treeNode(t, union, "Group:g1#members", nil),
				),
			),
		},
		{
			name: "case=expands (includes and traverse) - traverse has no tuples",
			opl: `
				class User implements Namespace{}
				class Group implements Namespace{
					related: {
						members: User[]
					}
				}
				class File implements Namespace{
					related: {
						viewers: User[]
						viewerGroups: Group[]
					}
					permits = {
						view: (ctx: Context) => this.related.viewers.includes(ctx.subject) &&
							this.related.viewerGroups.traverse(g => g.related.members.includes(ctx.subject))
					}
				}
			`,
			inputTuples: []string{
				"File:f1#viewers@User:1",
				"File:f1#viewers@User:2",
				"File:f1#viewers@User:3",
			},
			expandInput: "File:f1#view",
			expected: treeNode(t, intersection, "File:f1#view", nil,
				treeNode(t, union, "File:f1#viewers", nil,
					treeNode(t, leaf, "User:1", nil),
					treeNode(t, leaf, "User:2", nil),
					treeNode(t, leaf, "User:3", nil),
				),
				treeNode(t, union, "File:f1#viewerGroups", nil),
			),
		},
		{
			name: "case=expands multiple child permits",
			opl: `
				class User implements Namespace{}
				class Group implements Namespace{
					related: {
						members: User[]
					}
				}
				class File implements Namespace{
					related: {
						viewers: User[]
						viewerGroups: Group[]
					}
					permits = {
						viewUser: (ctx: Context) => this.related.viewers.includes(ctx.subject),
						viewGroup: (ctx: Context) => this.related.viewerGroups.traverse(g => g.related.members.includes(ctx.subject)),

						view: (ctx: Context) => this.permits.viewUser(ctx) && this.permits.viewGroup(ctx)
					}
				}
			`,
			inputTuples: []string{
				"File:f1#viewers@User:1",
				"File:f1#viewers@User:2",
				"File:f1#viewers@User:3",

				"Group:g1#members@User:4",
				"File:f1#viewerGroups@Group:g1",
			},
			expandInput: "File:f1#view",
			expected: treeNode(t, intersection, "File:f1#view", nil,
				treeNode(t, union, "File:f1#viewUser", nil,
					treeNode(t, union, "File:f1#viewers", nil,
						treeNode(t, leaf, "User:1", nil),
						treeNode(t, leaf, "User:2", nil),
						treeNode(t, leaf, "User:3", nil),
					),
				),
				treeNode(t, union, "File:f1#viewGroup", nil,
					treeNode(t, union, "File:f1#viewerGroups", nil,
						treeNode(t, union, "Group:g1#members", nil,
							treeNode(t, leaf, "User:4", nil),
						),
					),
				),
			),
		},
		{
			name: "case=expands compound relations",
			opl: `
				class User implements Namespace{}
				class File implements Namespace{
					related: {
						viewers: User[]
						editors: User[]
						owners: User[]
					}
					permits = {
						view: (ctx: Context) => 
							(
								this.related.viewers.includes(ctx.subject) ||
								this.related.editors.includes(ctx.subject)
							) &&
							this.related.owners.includes(ctx.subject)
					}
				}
			`,
			inputTuples: []string{
				"File:f1#viewers@User:1",
				"File:f1#editors@User:2",
				"File:f1#editors@User:3",
				"File:f1#owners@User:3",
			},
			expandInput: "File:f1#view",
			expected: treeNode(t, intersection, "File:f1#view", nil,
				treeNode(t, union, "File:f1#owners", nil,
					treeNode(t, leaf, "User:3", nil),
				),
				treeNode(t, union, "", nil,
					treeNode(t, union, "File:f1#viewers", nil,
						treeNode(t, leaf, "User:1", nil),
					),
					treeNode(t, union, "File:f1#editors", nil,
						treeNode(t, leaf, "User:2", nil),
						treeNode(t, leaf, "User:3", nil),
					),
				),
			),
		},
		{
			name: "case=expands repeated child permit calls",
			opl: `
				class User implements Namespace{}
				class Group implements Namespace{
					related: {
						members: User[]
					}
				}
				class File implements Namespace{
					related: {
						viewers: User[]
						viewerGroups: Group[]
					}
					permits = {
						viewUser: (ctx: Context) => this.related.viewers.includes(ctx.subject),
						viewGroup: (ctx: Context) => this.related.viewerGroups.traverse(g => g.related.members.includes(ctx.subject)),

						view: (ctx: Context) => (
							( this.permits.viewUser(ctx) && this.permits.viewGroup(ctx) )
							 ||
							( this.permits.viewUser(ctx) || this.permits.viewGroup(ctx) )
						)
					}
				}
			`,
			inputTuples: []string{
				"File:f1#viewers@User:1",
				"File:f1#viewers@User:2",
				"File:f1#viewers@User:3",

				"Group:g1#members@User:4",
				"File:f1#viewerGroups@Group:g1",
			},
			expandInput: "File:f1#view",
			expected: treeNode(t, union, "File:f1#view", nil,
				treeNode(t, intersection, "", nil,
					treeNode(t, union, "File:f1#viewUser", nil,
						treeNode(t, union, "File:f1#viewers", nil,
							treeNode(t, leaf, "User:1", nil),
							treeNode(t, leaf, "User:2", nil),
							treeNode(t, leaf, "User:3", nil),
						),
					),
					treeNode(t, union, "File:f1#viewGroup", nil,
						treeNode(t, union, "File:f1#viewerGroups", nil,
							treeNode(t, union, "Group:g1#members", nil,
								treeNode(t, leaf, "User:4", nil),
							),
						),
					),
				),
				treeNode(t, union, "File:f1#viewUser", nil,
					treeNode(t, union, "File:f1#viewers", nil,
						treeNode(t, leaf, "User:1", nil),
						treeNode(t, leaf, "User:2", nil),
						treeNode(t, leaf, "User:3", nil),
					),
				),
				treeNode(t, union, "File:f1#viewGroup", nil,
					treeNode(t, union, "File:f1#viewerGroups", nil,
						treeNode(t, union, "Group:g1#members", nil,
							treeNode(t, leaf, "User:4", nil),
						),
					),
				),
			),
		},
		{
			name: "case=expands recursive groups should detect circular references",
			opl: `
				class User implements Namespace{}
				class Group implements Namespace{
					related: {
						members: User[]
						parent: Group[]
					}
					permits = {
						isMember: (ctx: Context) => this.related.members.includes(ctx.subject) ||
							this.related.parent.traverse(g => g.permits.isMember(ctx))
					}
				}
			`,
			inputTuples: []string{
				"Group:g1#members@User:4",
				"Group:g1#parent@Group:g2",
				"Group:g2#members@User:5",
				"Group:g2#parent@Group:g1",
			},
			expandInput: "Group:g1#isMember",
			expected: treeNode(t, union, "Group:g1#isMember", nil,
				treeNode(t, union, "Group:g1#members", nil,
					treeNode(t, leaf, "User:4", nil),
				),
				treeNode(t, union, "Group:g1#parent", nil,
					treeNode(t, union, "Group:g2#isMember", nil,
						treeNode(t, union, "Group:g2#members", nil,
							treeNode(t, leaf, "User:5", nil),
						),
						treeNode(t, union, "Group:g2#parent", nil,
							treeNode(t, union, "Group:g1#isMember", &relationtuple.Truncation{
								Reason: relationtuple.TruncationReasonCycle,
							}),
						),
					),
				),
			),
		},
		{
			// node 1.= Group:g2#members
			// node 2 = Group:g3#members
			// rest is truncated.
			name:     "case=node-limit is counted correctly for SubjectSet relationships",
			maxNodes: 2,
			inputTuples: []string{
				"Group:g1#members@Group:g2#members",
				"Group:g2#members@Group:g3#members",
				"Group:g3#members@User:1",
			},
			expandInput: "Group:g1#members",
			expected: treeNode(t, union, "Group:g1#members", nil,
				treeNode(t, union, "Group:g2#members", nil,
					treeNode(t, union, "Group:g3#members", &relationtuple.Truncation{
						Reason: relationtuple.TruncationReasonTupleLimit,
						Cursor: &relationtuple.ExpandCursor{
							Kind:       relationtuple.ExpandCursorKindDirect,
							SubjectSet: testhelpers.SubjectSetFromString(t, "Group:g3#members"),
						},
					}),
				),
			),
		},
		{
			name: "case=stops at empty SubjectSet",
			inputTuples: []string{
				"Group:g1#members@Group:g2#members",
			},
			expandInput: "Group:g1#members",
			expected: treeNode(t, union, "Group:g1#members", nil,
				treeNode(t, union, "Group:g2#members", nil),
			),
		},
		{
			name:     "case=node limit does not block OPL structural traversal when there's data",
			maxNodes: 3,
			opl: `
				class User implements Namespace{}
				class Group implements Namespace{
					related: {
						members: User[]
					}
				}
				class File implements Namespace{
					related: {
						viewers: User[]
						viewerGroups: Group[]
					}
					permits = {
						viewUser: (ctx: Context) => this.related.viewers.includes(ctx.subject),
						viewGroup: (ctx: Context) => this.related.viewerGroups.traverse(g => g.related.members.includes(ctx.subject)),
						view: (ctx: Context) => this.permits.viewUser(ctx) || this.permits.viewGroup(ctx)
					}
				}
			`,
			inputTuples: []string{
				"File:f1#viewers@User:1",
				"File:f1#viewers@User:2",
				"File:f1#viewers@User:3",

				"File:f1#viewerGroups@Group:g1",
			},
			expandInput: "File:f1#view",
			expected: treeNode(t, union, "File:f1#view", nil,
				treeNode(t, union, "File:f1#viewUser", nil,
					treeNode(t, union, "File:f1#viewers", nil,
						treeNode(t, leaf, "User:1", nil),
						treeNode(t, leaf, "User:2", nil),
						treeNode(t, leaf, "User:3", nil),
					),
				),
				treeNode(t, union, "File:f1#viewGroup", nil,
					treeNode(t, union, "File:f1#viewerGroups", &relationtuple.Truncation{
						Reason: relationtuple.TruncationReasonTupleLimit,
						Cursor: &relationtuple.ExpandCursor{
							Kind:             relationtuple.ExpandCursorKindTTU,
							SubjectSet:       testhelpers.SubjectSetFromString(t, "File:f1#viewerGroups"),
							TraverseRelation: new("members"),
						},
					}),
				),
			),
		},
		{
			name:   "case=unknown 2nd relationship on strict mode",
			strict: true,
			opl: `
				class User implements Namespace{}
				class Group implements Namespace{
					related: {
						viewers: User[]
					}
				}
				class File implements Namespace{
					related: {
						groupViewers: SubjectSet<Group, "viewers">[]
						groups: Group[]
					}
					permits = {
						view: (ctx: Context) => this.related.groupViewers.includes(ctx.subject) ||
							|| this.related.groups.traverse(g => g.related.viewers.includes(ctx.subject))
					}
				}
			`,
			inputTuples: []string{
				"Group:g1#editors@User:1", // "editor" relationship does not exist on Group
				"Group:g2#editors@User:2",

				"File:f1#groupViewers@Group:g1#readers", // "reader" relationship does not exist on Group
				"File:f1#groups@Group:g2",
			},
			expandInput: "File:f1#view",
			expected: treeNode(t, union, "File:f1#view", nil,
				treeNode(t, union, "File:f1#groupViewers", nil),
				treeNode(t, union, "File:f1#groups", nil,
					treeNode(t, union, "Group:g2#viewers", nil),
				),
			),
		},
		{
			name:     "case=unknown relationship on strict mode",
			maxNodes: 3,
			strict:   true,
			opl: `
				class User implements Namespace{}
				class File implements Namespace{
					related: {
						viewers: User[]
					}
				}
			`,
			inputTuples: []string{
				"File:f1#undefined@User:1",
			},
			expandInput: "File:f1#undefined",
			expected:    nil,
		},
		{
			name:     "case=unknown namespace on strict mode",
			maxNodes: 3,
			strict:   true,
			opl: `
				class User implements Namespace{}
				class File implements Namespace{
					related: {
						viewers: User[]
					}
				}
			`,
			inputTuples: []string{
				"File:f1#undefined@User:1",
			},
			expandInput: "File:f1#undefined",
			expected:    nil,
		},
		{
			name:     "case=unknown relationship on non-strict mode",
			maxNodes: 3,
			opl: `
				class User implements Namespace{}
				class File implements Namespace{
					related: {
						viewers: User[]
					}
				}
			`,
			inputTuples: []string{
				"File:f1#undefined@User:1",
			},
			expandInput: "File:f1#undefined",
			expected: treeNode(t, union, "File:f1#undefined", nil,
				treeNode(t, leaf, "User:1", nil),
			),
		},
		{
			name:     "case=unknown namespace on non-strict mode",
			maxNodes: 3,
			opl: `
				class User implements Namespace{}
				class File implements Namespace{
					related: {
						viewers: User[]
					}
				}
			`,
			inputTuples: []string{},
			expandInput: "X:f1#undefined",
			expected:    treeNode(t, union, "X:f1#undefined", nil),
		},
		{
			name:   "case=limit should be applied and prevent nesting",
			strict: true,
			opl: `
				class File implements Namespace{}
				class Folder implements Namespace{
					related: {
						members: (File | SubjectSet<Folder, "members">)[]
					}
				}
			`,
			// the 1st digit refers to the depth, 2nd digit refers to the position in that depth.
			inputTuples: []string{
				"Folder:11#members@Folder:21#members",
				"Folder:11#members@Folder:22#members",

				"Folder:21#members@Folder:31#members",
				"Folder:21#members@Folder:32#members",
				"Folder:21#members@Folder:33#members",

				// Folder:22 has no members

				"Folder:31#members@File:41", "Folder:31#members@File:42",
				"Folder:32#members@File:43", "Folder:32#members@File:44",
			},
			maxNodes:    4,
			expandInput: "Folder:11#members",
			expected: treeNode(t, union, "Folder:11#members", nil,
				treeNode(t, union, "Folder:21#members", nil,
					treeNode(t, union, "Folder:31#members", &relationtuple.Truncation{
						Reason: relationtuple.TruncationReasonTupleLimit,
						Cursor: &relationtuple.ExpandCursor{
							Kind:       relationtuple.ExpandCursorKindDirect,
							SubjectSet: testhelpers.SubjectSetFromString(t, "Folder:31#members"),
						},
					}),
					treeNode(t, union, "Folder:32#members", &relationtuple.Truncation{
						Reason: relationtuple.TruncationReasonTupleLimit,
						Cursor: &relationtuple.ExpandCursor{
							Kind:       relationtuple.ExpandCursorKindDirect,
							SubjectSet: testhelpers.SubjectSetFromString(t, "Folder:32#members"),
						},
					}),
					treeNode(t, union, "Folder:33#members", &relationtuple.Truncation{
						Reason: relationtuple.TruncationReasonTupleLimit,
						Cursor: &relationtuple.ExpandCursor{
							Kind:       relationtuple.ExpandCursorKindDirect,
							SubjectSet: testhelpers.SubjectSetFromString(t, "Folder:33#members"),
						},
					}),
				),
				treeNode(t, union, "Folder:22#members", &relationtuple.Truncation{
					Reason: relationtuple.TruncationReasonTupleLimit,
					Cursor: &relationtuple.ExpandCursor{
						Kind:       relationtuple.ExpandCursorKindDirect,
						SubjectSet: testhelpers.SubjectSetFromString(t, "Folder:22#members"),
					},
				}),
			),
		},
		{
			name: "case=direct relations should be prioritized over rewrites",
			opl: `
				class User implements Namespace{}
				class Group implements Namespace{
					related: {
						members: User[]
						parent: Group[]
						hiddenMembers: User[]
					}
					permits = {
						isMember: (ctx: Context) => this.related.members.includes(ctx.subject) 
							||	this.related.parent.traverse(p => p.permits.isMember(ctx))
							||	this.related.hiddenMembers.includes(ctx.subject) 
					}
				}
			`,
			inputTuples: []string{
				"Group:Sales#parent@Group:All",

				"Group:Sales#members@User:1",
				"Group:Sales#members@User:2",

				"Group:Sales#hiddenMembers@User:3",
				"Group:Sales#hiddenMembers@User:4",

				"Group:All#members@User:5",
				"Group:All#hiddenMembers@User:6",
			},
			maxNodes:    4,
			expandInput: "Group:Sales#isMember",
			expected: treeNode(t, union, "Group:Sales#isMember", nil,
				treeNode(t, union, "Group:Sales#members", nil,
					treeNode(t, leaf, "User:1", nil),
					treeNode(t, leaf, "User:2", nil),
				),
				treeNode(t, union, "Group:Sales#hiddenMembers", nil,
					treeNode(t, leaf, "User:3", nil),
					treeNode(t, leaf, "User:4", nil),
				),
				treeNode(t, union, "Group:Sales#parent", &relationtuple.Truncation{
					Reason: relationtuple.TruncationReasonTupleLimit,
					Cursor: &relationtuple.ExpandCursor{
						Kind:             relationtuple.ExpandCursorKindTTU,
						SubjectSet:       testhelpers.SubjectSetFromString(t, "Group:Sales#parent"),
						TraverseRelation: new("isMember"),
					},
				}),
			),
		},
	}

	t.Run("loop tests", func(t *testing.T) {
		focused := slices.ContainsFunc(tests, func(tt expandTest) bool { return tt.only })
		for _, tt := range tests {
			if focused && !tt.only {
				continue
			}
			t.Run(tt.name, func(t *testing.T) {
				if tt.opl != "" {
					_, errs := schema.Parse(tt.opl)
					require.Empty(t, errs)
				}
				opts := []driver.TestRegistryOption{
					driver.WithConfig(config.KeyFeatureFlagExpandRewrites, true),
					driver.WithConfig(config.KeyLimitMaxReadDepth, 50),
				}

				if tt.opl != "" {
					opts = append(opts, driver.WithOPL(tt.opl))
				} else {
					opts = append(opts, driver.WithConfig(config.KeyNamespaces, []*namespace.Namespace{}))
				}
				if tt.maxNodes > 0 {
					opts = append(opts, driver.WithConfig(config.KeyLimitExpandMaxSize, tt.maxNodes))
				}
				if tt.strict {
					opts = append(opts, driver.WithConfig(config.KeyNamespacesExperimentalStrictMode, tt.strict))
				}

				reg := driver.NewSqliteTestRegistry(t, false, opts...)

				e := expand.NewEngine(reg)

				ctx := context.Background()
				testhelpers.MapAndInsertTuplesFromString(t, reg, tt.inputTuples)

				tree, err := e.BuildTree(ctx, testhelpers.SubjectSetFromString(t, tt.expandInput), 100)
				require.NoError(t, err)

				expand.RequireInternalTreesAreEqual(t, tt.expected, tree)
			})
		}
	})
}

var (
	union        = ketoapi.TreeNodeUnion
	leaf         = ketoapi.TreeNodeLeaf
	intersection = ketoapi.TreeNodeIntersection
	exclusion    = ketoapi.TreeNodeExclusion
)

// treeNode builds a *relationtuple.Tree for use in test expectations.
// Pass subject="" to leave Subject nil (needed for Exclusion/Intersection
// nodes that have no subject). Omit children to leave Children nil.
func treeNode(t testing.TB, nodeType ketoapi.TreeNodeType, subject string, truncation *relationtuple.Truncation, children ...*relationtuple.Tree) *relationtuple.Tree {
	t.Helper()
	node := &relationtuple.Tree{
		Type:       nodeType,
		Truncation: truncation,
		Children:   children,
	}
	if subject != "" {
		node.Subject = testhelpers.SubjectFromString(t, subject)
	}
	if len(children) == 0 {
		node.Children = nil
	}
	return node
}
