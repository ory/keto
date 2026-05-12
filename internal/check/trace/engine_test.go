// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package trace_test

import (
	"context"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/check/trace"
	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/testhelpers"
	"github.com/ory/keto/schema"
)

// # Non-determinism constraint
//
// Expand and traverse (tuple-to-subject-set) children are ordered by
// shard_id (random UUIDv4) and require sortNode normalization.
//
// As a result, we cannot write tests where multiple Expand or Traverse siblings
// compete to produce member; which one wins first is non-deterministic, as the
// child tuples are not executed in the order we create them.
// Therefore:
//   - sortNode is used to sort the children of an Expand or Traverse nodes
//   - Tests where all branches return not-member are always safe: every branch
//     runs to completion before not-member is returned.
//   - Do not write tests where Expand or Traverse has multiple children, and one child
//     returns member, the other one notMember; which one runs first is non-deterministic;

type testcase struct {
	name        string
	opl         string
	inputTuples []string
	maxDepth    int

	checkInput string
	expected   func(t testing.TB) *trace.Node
	strict     bool
	only       bool
}

func TestTraceEngine(t *testing.T) {
	t.Parallel()

	tests := []testcase{
		// ── no-OPL mode ──────────────────────────────────────────────────────────

		{
			// Engine finds the tuple immediately via direct lookup.
			// Expand is registered eagerly, but the Executor already has a decisive
			// result from the direct match, so the expansion work is skipped.
			name: "no-OPL: direct hit",
			inputTuples: []string{
				"Group:eng#members@User:1",
			},
			checkInput: "Group:eng#members@User:1",
			expected: func(t testing.TB) *trace.Node {
				return checkNode(t, member, "Group:eng#members@User:1",
					directNode(t, member, "Group:eng#members@User:1"),
				)
			},
		},
		{
			// Engine follows a three-hop SubjectSet chain to find a member.
			name: "no-OPL: multi-hop SubjectSet expansion",
			inputTuples: []string{
				"Group:a#members@User:1",
				"Group:b#members@Group:a#members",
				"Group:c#members@Group:b#members",
			},
			checkInput: "Group:c#members@User:1",
			expected: func(t testing.TB) *trace.Node {
				return checkNode(t, member, "Group:c#members@User:1",
					directNode(t, notMember, "Group:c#members@User:1"),
					expandNode(t, member, "Group:c#members@User:1", 1,
						checkNode(t, member, "Group:b#members@User:1",
							expandNode(t, member, "Group:b#members@User:1", 1,
								foundNode(t, "Group:a#members@User:1"),
							),
						),
					),
				)
			},
		},
		{
			// Engine hits the depth cap at restDepth=0 and returns unknown.
			name:     "no-OPL: depth limit stops expansion",
			maxDepth: 2,
			inputTuples: []string{
				"Group:a#members@User:1",
				"Group:b1#members@Group:a#members",
				"Group:c#members@Group:b1#members",
			},
			checkInput: "Group:c#members@User:99",
			expected: func(t testing.TB) *trace.Node {
				return checkNode(t, notMember, "Group:c#members@User:99",
					directNode(t, notMember, "Group:c#members@User:99"),
					expandNode(t, notMember, "Group:c#members@User:99", 1,
						checkNode(t, unknown, "Group:b1#members@User:99",
							expandNode(t, unknown, "Group:b1#members@User:99", 0),
						),
					),
				)
			},
		},

		// ── strict OPL mode ───────────────────────────────────────────────────────

		{
			// In strict mode the only type-safe checks are fired;
			// no useless direct/expand on relations that are "permit";
			// Subject-set fires only when OPL field declares a SubjectSet<..>;
			name:   "strict OPL: SubjectSet expand fires only for declared types",
			strict: true,
			opl: `
				class User implements Namespace{}
				class Group implements Namespace{
					related: {
						members: User[]
					}
				}
				class File implements Namespace {
					related: {
						viewers: SubjectSet<Group, "members">[]
					}
					permits = {
						view: (ctx: Context) => this.related.viewers.includes(ctx.subject)
					}
				}
			`,
			inputTuples: []string{
				"File:f1#viewers@Group:g1#members",
			},
			checkInput: "File:f1#view@User:u1",
			expected: func(t testing.TB) *trace.Node {
				return checkNode(t, notMember, "File:f1#view@User:u1",
					unionNode(t, notMember, "File:f1#view@User:u1",
						multiDirectNode(t, notMember, "File:f1#view@User:u1", []string{"viewers"}),
						checkNode(t, notMember, "File:f1#viewers@User:u1",
							// expand fires because viewers declares SubjectSet<Group,"members">.
							// expand found 1 node, which is Group:g1#members, which is not a direct hit → not-member.
							expandNode(t, notMember, "File:f1#viewers@User:u1", 1,
								// skipDirect=true + no SubjectSet types on members → empty check
								checkNode(t, notMember, "Group:g1#members@User:u1"),
							),
						),
					),
				)
			},
		},
		{
			// Strict OR with multiple computed-usersets and a TTU.
			// The OR shortcut batches all plain computed-usersets into a single multiDirect check.
			// Plain related fields (viewers, viewers2) with no SubjectSet types produce
			// empty checks because multiDirect already covers them, and sets skipDirect=true.
			name:   "strict OPL: OR batches computed-usersets, suppresses spurious permit checks",
			strict: true,
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
						viewers2: User[]
						viewerGroups: Group[]
					}
					permits = {
						view: (ctx: Context) => this.related.viewers.includes(ctx.subject) ||
							this.related.viewerGroups.traverse(g => g.related.members.includes(ctx.subject)) ||
							this.related.viewers2.includes(ctx.subject)
					}
				}
			`,
			inputTuples: []string{
				"File:f1#viewers@User:1",
				"File:f1#viewers@User:2",

				"File:f1#viewers2@User:10",
				"File:f1#viewers2@User:11",

				"File:f1#viewerGroups@Group:g1",
				"Group:g1#members@User:4",
				"Group:g1#members@User:5",
			},
			checkInput: "File:f1#view@User:100",
			expected: func(t testing.TB) *trace.Node {
				return checkNode(t, notMember, "File:f1#view@User:100",
					unionNode(t, notMember, "File:f1#view@User:100",
						// OR shortcut batches viewers + viewers2 into one multi-relation direct check.
						multiDirectNode(t, notMember, "File:f1#view@User:100", []string{"viewers", "viewers2"}),
						// viewers and viewers2 are plain User[] fields: empty checks in strict mode.
						checkNode(t, notMember, "File:f1#viewers@User:100"),
						checkNode(t, notMember, "File:f1#viewers2@User:100"),
						// TTU traverses viewerGroups; Group:g1 has no User:100 → not-member.
						traverseNode(t, notMember, "File:f1#viewerGroups@User:100", 1,
							checkNode(t, notMember, "Group:g1#members@User:100",
								directNode(t, notMember, "Group:g1#members@User:100"),
							),
						),
					),
				)
			},
		},
		{
			// Strict OR where the TTU branch finds a member.
			name:   "strict OPL: OR with TTU finds member",
			strict: true,
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
				"Group:g1#members@User:4",
			},
			checkInput: "File:f1#view@User:4",
			expected: func(t testing.TB) *trace.Node {
				return checkNode(t, member, "File:f1#view@User:4",
					unionNode(t, member, "File:f1#view@User:4",
						// multiDirect for viewers: User:4 is not a direct viewer.
						multiDirectNode(t, notMember, "File:f1#view@User:4", []string{"viewers"}),
						// viewers is User[]: empty check in strict+skipDirect.
						checkNode(t, notMember, "File:f1#viewers@User:4"),
						// TTU loads Group:g1; User:4 is a direct member → member.
						// strict: members is User[] (no SubjectSet types) → no expand registered.
						traverseNode(t, member, "File:f1#viewerGroups@User:4", 1,
							checkNode(t, member, "Group:g1#members@User:4",
								directNode(t, member, "Group:g1#members@User:4"),
							),
						),
					),
				)
			},
		},
		{
			// NOT inverts an inner member result to not-member.
			name:   "strict OPL: NOT inverts member to not-member",
			strict: true,
			opl: `
				class User implements Namespace{}
				class File implements Namespace{
					related: {
						blocklist: User[]
					}
					permits = {
						view: (ctx: Context) => !this.related.blocklist.includes(ctx.subject)
					}
				}
			`,
			inputTuples: []string{
				"File:f1#blocklist@User:blocked",
			},
			checkInput: "File:f1#view@User:blocked",
			expected: func(t testing.TB) *trace.Node {
				return checkNode(t, notMember, "File:f1#view@User:blocked",
					unionNode(t, notMember, "File:f1#view@User:blocked",
						invertNode(t, notMember, "File:f1#view@User:blocked",
							computedNode(t, member, "File:f1#blocklist@User:blocked",
								checkNode(t, member, "File:f1#blocklist@User:blocked",
									directNode(t, member, "File:f1#blocklist@User:blocked"),
								),
							),
						),
					),
				)
			},
		},
		{
			// NOT inverts an inner not-member result to member.
			name:   "strict OPL: NOT inverts not-member to member",
			strict: true,
			opl: `
				class User implements Namespace{}
				class File implements Namespace{
					related: {
						blocklist: User[]
					}
					permits = {
						view: (ctx: Context) => !this.related.blocklist.includes(ctx.subject)
					}
				}
			`,
			inputTuples: []string{
				"File:f1#blocklist@User:blocked",
			},
			checkInput: "File:f1#view@User:free",
			expected: func(t testing.TB) *trace.Node {
				return checkNode(t, member, "File:f1#view@User:free",
					unionNode(t, member, "File:f1#view@User:free",
						invertNode(t, member, "File:f1#view@User:free",
							computedNode(t, notMember, "File:f1#blocklist@User:free",
								checkNode(t, notMember, "File:f1#blocklist@User:free",
									directNode(t, notMember, "File:f1#blocklist@User:free"),
								),
							),
						),
					),
				)
			},
		},
		{
			name:   "strict OPL: TTU with permit call inside traverse",
			strict: true,
			opl: `
				class User implements Namespace{}
				class Group implements Namespace{
					related: {
						members: User[]
					}
					permits = {
						isMember: (ctx: Context) => this.related.members.includes(ctx.subject)
					}
				}
				class File implements Namespace{
					related: {
						viewerGroups: Group[]
					}
					permits = {
						view: (ctx: Context) => this.related.viewerGroups.traverse(g => g.permits.isMember(ctx))
					}
				}
			`,
			inputTuples: []string{
				"Group:g1#members@User:1",
				"Group:g1#members@User:2",
				"File:f1#viewerGroups@Group:g1",
			},
			checkInput: "File:f1#view@User:99",
			expected: func(t testing.TB) *trace.Node {
				return checkNode(t, notMember, "File:f1#view@User:99",
					unionNode(t, notMember, "File:f1#view@User:99",
						traverseNode(t, notMember, "File:f1#viewerGroups@User:99", 1,
							checkNode(t, notMember, "Group:g1#isMember@User:99",
								unionNode(t, notMember, "Group:g1#isMember@User:99",
									multiDirectNode(t, notMember, "Group:g1#members@User:99", []string{"members"}),
									checkNode(t, notMember, "Group:g1#members@User:99"),
								),
							),
						),
					),
				)
			},
		},

		// ── non-strict OPL mode ───────────────────────────────────────────────────

		{
			// Non-strict mode fires the rewrite, a direct check, AND a subject-set
			// expand concurrently even for a permit relation (view).
			name: "non-strict OPL: OR on permit reference emits spurious direct and expand",
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
						viewGroup: (ctx: Context) => this.related.viewerGroups.traverse(g => g.related.members.includes(ctx.subject)),
						view: (ctx: Context) => this.related.viewers.includes(ctx.subject) || this.permits.viewGroup(ctx)
					}
				}
			`,
			inputTuples: []string{
				"File:f1#viewers@User:1",
				"File:f1#viewers@User:2",
				"File:f1#viewers@User:3",

				"Group:g1#members@User:4",
				"Group:g1#members@User:5",
				"Group:g2#members@User:6",
				"Group:g2#members@User:7",
				"Group:g3#members@User:8",
				"Group:g3#members@User:9",

				"File:f1#viewerGroups@Group:g1",
				"File:f1#viewerGroups@Group:g2",
				"File:f1#viewerGroups@Group:g3",
			},
			checkInput: "File:f1#view@User:100",
			expected: func(t testing.TB) *trace.Node {
				leafCheck := func(s string) *trace.Node {
					return checkNode(t, notMember, s,
						expandNode(t, notMember, s, 0),
					)
				}
				groupCheck := func(g string, users ...string) *trace.Node {
					members := make([]*trace.Node, len(users))
					for i, u := range users {
						members[i] = leafCheck("User:" + u + "#@User:100")
					}
					return checkNode(t, notMember, "Group:"+g+"#members@User:100",
						directNode(t, notMember, "Group:"+g+"#members@User:100"),
						expandNode(t, notMember, "Group:"+g+"#members@User:100", len(users), members...),
					)
				}
				return checkNode(t, notMember, "File:f1#view@User:100",
					unionNode(t, notMember, "File:f1#view@User:100",
						multiDirectNode(t, notMember, "File:f1#view@User:100", []string{"viewers", "viewGroup"}),
						checkNode(t, notMember, "File:f1#viewers@User:100",
							expandNode(t, notMember, "File:f1#viewers@User:100", 3,
								leafCheck("User:1#@User:100"),
								leafCheck("User:2#@User:100"),
								leafCheck("User:3#@User:100"),
							),
						),
						checkNode(t, notMember, "File:f1#viewGroup@User:100",
							unionNode(t, notMember, "File:f1#viewGroup@User:100",
								traverseNode(t, notMember, "File:f1#viewerGroups@User:100", 3,
									groupCheck("g1", "4", "5"),
									groupCheck("g2", "6", "7"),
									groupCheck("g3", "8", "9"),
								),
							),
							expandNode(t, notMember, "File:f1#viewGroup@User:100", 0),
						),
					),
					directNode(t, notMember, "File:f1#view@User:100"),
					expandNode(t, notMember, "File:f1#view@User:100", 0),
				)
			},
		},
		{
			// AND: both arms must complete before the result is known, so
			// the full tree is always explored regardless of intermediate results.
			name: "non-strict OPL: AND requires both arms; permit-calls-permit",
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
						viewGroup: (ctx: Context) => this.related.viewerGroups.traverse(g => g.related.members.includes(ctx.subject)),

						view: (ctx: Context) => this.related.viewers.includes(ctx.subject) &&
							this.permits.viewGroup(ctx)
					}
				}
			`,
			inputTuples: []string{
				"File:f1#viewers@User:1",
				"File:f1#viewers@User:2",
				"File:f1#viewers@User:3",

				"Group:g1#members@User:4",
				"Group:g1#members@User:5",
				"Group:g2#members@User:6",
				"Group:g2#members@User:7",
				"Group:g3#members@User:8",
				"Group:g3#members@User:9",

				"File:f1#viewerGroups@Group:g1",
				"File:f1#viewerGroups@Group:g2",
				"File:f1#viewerGroups@Group:g3",
			},
			checkInput: "File:f1#view@User:1",
			expected: func(t testing.TB) *trace.Node {
				leafCheck := func(s string) *trace.Node {
					return checkNode(t, notMember, s,
						expandNode(t, notMember, s, 0),
					)
				}
				groupCheck := func(g string, users ...string) *trace.Node {
					members := make([]*trace.Node, len(users))
					for i, u := range users {
						members[i] = leafCheck("User:" + u + "#@User:1")
					}
					return checkNode(t, notMember, "Group:"+g+"#members@User:1",
						directNode(t, notMember, "Group:"+g+"#members@User:1"),
						expandNode(t, notMember, "Group:"+g+"#members@User:1", len(users), members...),
					)
				}
				return checkNode(t, notMember, "File:f1#view@User:1",
					intersectionNode(t, notMember, "File:f1#view@User:1",
						computedNode(t, member, "File:f1#viewers@User:1",
							checkNode(t, member, "File:f1#viewers@User:1",
								directNode(t, member, "File:f1#viewers@User:1"),
							),
						),
						computedNode(t, notMember, "File:f1#viewGroup@User:1",
							checkNode(t, notMember, "File:f1#viewGroup@User:1",
								unionNode(t, notMember, "File:f1#viewGroup@User:1",
									traverseNode(t, notMember, "File:f1#viewerGroups@User:1", 3,
										groupCheck("g1", "4", "5"),
										groupCheck("g2", "6", "7"),
										groupCheck("g3", "8", "9"),
									),
								),
								directNode(t, notMember, "File:f1#viewGroup@User:1"),
								expandNode(t, notMember, "File:f1#viewGroup@User:1", 0),
							),
						),
					),
					directNode(t, notMember, "File:f1#view@User:1"),
					expandNode(t, notMember, "File:f1#view@User:1", 0),
				)
			},
		},
	}
	focused := slices.ContainsFunc(tests, func(tt testcase) bool { return tt.only })
	for _, tt := range tests {
		if focused && !tt.only {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			opts := []driver.TestRegistryOption{driver.WithLogLevel("info")}

			if tt.opl != "" {
				_, errs := schema.Parse(tt.opl)
				require.Empty(t, errs, "invalid OPL in test case: %v", errs)
			}

			if tt.opl != "" {
				opts = append(opts, driver.WithOPL(tt.opl))
			} else {
				opts = append(opts, driver.WithConfig(config.KeyNamespaces, []*namespace.Namespace{}))
			}

			if tt.maxDepth != 0 {
				opts = append(opts, driver.WithConfig(config.KeyLimitMaxReadDepth, tt.maxDepth))
			}

			if tt.strict {
				opts = append(opts, driver.WithConfig(config.KeyNamespacesExperimentalStrictMode, tt.strict))
			}

			reg := driver.NewSqliteTestRegistry(t, opts...)
			e := trace.NewEngine(reg)

			ctx := context.Background()
			for _, row := range tt.inputTuples {
				require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(ctx, testhelpers.TupleFromString(t, row)))
			}

			_, tree, err := e.CheckIsMemberWithTrace(ctx, testhelpers.TupleFromString(t, tt.checkInput), 100)
			require.NoError(t, err)

			expected := sortNode(tt.expected(t))
			actual := sortNode(stripTiming(tree))
			require.Equal(t, expected.String(), actual.String())
		})
	}
}

// sortNode sorts children of NodeExpand and NodeTupleToSet nodes by Tuple.String(),
// recursively. Both node kinds have non-deterministic child order because the
// underlying queries order by shard_id (a random UUIDv4 assigned at write
// time). All other sibling orderings are engine-controlled and must be
// asserted as-is.
func sortNode(n *trace.Node) *trace.Node {
	if n == nil {
		return nil
	}
	if n.Kind == trace.NodeExpandSubject || n.Kind == trace.NodeTraverse {
		slices.SortFunc(n.Children, func(a, b *trace.Node) int {
			as, bs := "", ""
			if a.Tuple != nil {
				as = a.Tuple.String()
			}
			if b.Tuple != nil {
				bs = b.Tuple.String()
			}
			if c := strings.Compare(as, bs); c != 0 {
				return c
			}
			return strings.Compare(string(a.Kind), string(b.Kind))
		})
	}
	for _, child := range n.Children {
		sortNode(child)
	}
	return n
}

func stripTiming(n *trace.Node) *trace.Node {
	if n == nil {
		return nil
	}
	n.Duration = 0
	for _, child := range n.Children {
		stripTiming(child)
	}
	return n
}

// result helpers — used as the first argument to every node builder.
const (
	member    = trace.ResultMember
	notMember = trace.ResultNotMember
	skipped   = trace.ResultSkipped
	unknown   = trace.ResultUnknown
)

// checkNode builds a NodeCheck node.
func checkNode(t testing.TB, result trace.NodeResult, tuple string, children ...*trace.Node) *trace.Node {
	t.Helper()
	return &trace.Node{Kind: trace.NodeIsAllowed, Tuple: testhelpers.TupleFromString(t, tuple), Result: result, Children: children}
}

// directNode builds a NodeDirect leaf (no children).
func directNode(t testing.TB, result trace.NodeResult, tuple string) *trace.Node {
	t.Helper()
	return &trace.Node{Kind: trace.NodeDirect, Tuple: testhelpers.TupleFromString(t, tuple), Result: result}
}

// expandNode builds a NodeExpand node.
func expandNode(t testing.TB, result trace.NodeResult, tuple string, tuplesLoaded int, children ...*trace.Node) *trace.Node {
	t.Helper()
	return &trace.Node{Kind: trace.NodeExpandSubject, Tuple: testhelpers.TupleFromString(t, tuple), Result: result, TuplesLoaded: tuplesLoaded, Children: children}
}

// foundNode builds a NodeFoundSubject leaf (no children).
func foundNode(t testing.TB, tuple string) *trace.Node {
	t.Helper()
	return &trace.Node{Kind: trace.NodeFoundSubject, Tuple: testhelpers.TupleFromString(t, tuple), Result: trace.ResultMember}
}

// unionNode builds a NodeUnion node.
func unionNode(t testing.TB, result trace.NodeResult, tuple string, children ...*trace.Node) *trace.Node {
	t.Helper()
	return &trace.Node{Kind: trace.NodeUnion, Tuple: testhelpers.TupleFromString(t, tuple), Result: result, Children: children}
}

// intersectionNode builds a NodeIntersection node.
func intersectionNode(t testing.TB, result trace.NodeResult, tuple string, children ...*trace.Node) *trace.Node {
	t.Helper()
	return &trace.Node{Kind: trace.NodeIntersection, Tuple: testhelpers.TupleFromString(t, tuple), Result: result, Children: children}
}

// invertNode builds a NodeInvert node.
func invertNode(t testing.TB, result trace.NodeResult, tuple string, children ...*trace.Node) *trace.Node {
	t.Helper()
	return &trace.Node{Kind: trace.NodeInvert, Tuple: testhelpers.TupleFromString(t, tuple), Result: result, Children: children}
}

// computedNode builds a NodeComputed node.
func computedNode(t testing.TB, result trace.NodeResult, tuple string, children ...*trace.Node) *trace.Node {
	t.Helper()
	return &trace.Node{Kind: trace.NodeComputed, Tuple: testhelpers.TupleFromString(t, tuple), Result: result, Children: children}
}

// traverseNode builds a NodeTupleToSet node.
func traverseNode(t testing.TB, result trace.NodeResult, tuple string, tuplesLoaded int, children ...*trace.Node) *trace.Node {
	t.Helper()
	return &trace.Node{Kind: trace.NodeTraverse, Tuple: testhelpers.TupleFromString(t, tuple), Result: result, TuplesLoaded: tuplesLoaded, Children: children}
}

// multiDirectNode builds a NodeMultiDirect leaf (no children).
func multiDirectNode(t testing.TB, result trace.NodeResult, tuple string, relations []string) *trace.Node {
	t.Helper()
	return &trace.Node{Kind: trace.NodeMultiDirect, Tuple: testhelpers.TupleFromString(t, tuple), Relations: relations, Result: result}
}
