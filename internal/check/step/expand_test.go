// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package step_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/check/trace"
	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/testhelpers"
)

// TestExpandSubjectStep_MemberFound verifies that ExpandSubjectStep produces
// the correct short-circuit trace when the subject is found via subject-set expansion.
func TestExpandSubjectStep_MemberFound(t *testing.T) {
	t.Parallel()

	namespaceOnlyOPL := `class User implements Namespace{} class Group implements Namespace{} class File implements Namespace{}`

	tests := []struct {
		scenario           testhelpers.Scenario
		checkInput         string
		expectedMembership check.Membership
		expectedTrace      func(t testing.TB) *trace.Node
	}{
		{
			scenario: testhelpers.Scenario{
				Name:        "1-hop subject-set hit resolves without recursive sub-check",
				Opl:         namespaceOnlyOPL,
				InputTuples: []string{"File:f1#viewers@Group:g1#members", "Group:g1#members@User:Alice"},
			},
			checkInput:         "File:f1#viewers@User:Alice",
			expectedMembership: check.IsMember,
			expectedTrace: func(t testing.TB) *trace.Node {
				return trace.CheckNode(t, trace.ResultMember, "File:f1#viewers@User:Alice",
					trace.DirectNode(t, trace.ResultNotMember, "File:f1#viewers@User:Alice"),
					trace.ExpandNode(t, trace.ResultMember, "File:f1#viewers@User:Alice", 1,
						trace.FoundNode(t, "Group:g1#members@User:Alice"),
					),
				)
			},
		},
		{
			scenario: testhelpers.Scenario{
				Name: "2-hop subject-set hit resolves without recursive sub-check",
				Opl:  namespaceOnlyOPL,
				InputTuples: []string{
					"File:f1#viewers@Group:g1#members",
					"Group:g1#members@Group:g2#members",
					"Group:g2#members@User:Alice",
				},
			},
			checkInput:         "File:f1#viewers@User:Alice",
			expectedMembership: check.IsMember,
			expectedTrace: func(t testing.TB) *trace.Node {
				return trace.CheckNode(t, trace.ResultMember, "File:f1#viewers@User:Alice",
					trace.DirectNode(t, trace.ResultNotMember, "File:f1#viewers@User:Alice"),
					trace.ExpandNode(t, trace.ResultMember, "File:f1#viewers@User:Alice", 1,
						trace.CheckNode(t, trace.ResultMember, "Group:g1#members@User:Alice",
							trace.ExpandNode(t, trace.ResultMember, "Group:g1#members@User:Alice", 1,
								trace.FoundNode(t, "Group:g2#members@User:Alice"),
							),
						),
					),
				)
			},
		},
	}

	for _, tt := range tests {
		tt.scenario.Run(t, func(t *testing.T, reg driver.Registry) {
			t.Parallel()

			e := trace.NewEngine(reg)
			res, tree := e.CheckRelationTupleWithTrace(t.Context(), testhelpers.TupleFromString(t, tt.checkInput), 100)
			require.Equal(t, tt.expectedMembership, res.Membership)
			require.NoError(t, res.Err)
			require.Equal(t, trace.SortNode(tt.expectedTrace(t)).String(), trace.SortNode(trace.StripTiming(tree)).String())
		})
	}
}

// TestExpandSubjectStep_NotMember documents the not-member traces across the three
// OPL modes (namespace-only; typed OPL; typed OPL with strict mode).
//
// Without typed OPL, the engine follows every subject-set pointer in the DB, including
// leaf subjects stored as subject-sets with an empty relation. With typed OPL the engine
// filters expansion to only the subject-set types declared in the schema, skipping leaf
// subjects that are not valid subject-set types for the relation.
func TestExpandSubjectStep_NotMember(t *testing.T) {
	t.Parallel()

	typedOPL := `
		class User implements Namespace{}
		class Group implements Namespace{
			related: {
				members: (User|SubjectSet<Group, "members">)[]
			}
		}
		class File implements Namespace{
			related: {
				viewers: SubjectSet<Group, "members">[]
			}
		}
	`
	inputTuples := []string{
		"File:f1#viewers@Group:g1#members",
		"Group:g1#members@User:Alice",
	}

	tests := []struct {
		scenario           testhelpers.Scenario
		checkInput         string
		expectedMembership check.Membership
		expectedTrace      func(t testing.TB) *trace.Node
	}{
		{
			// The engine follows every subject-set pointer including leaf subjects stored as SubjectSets.
			// In only-namespaces setup, everything	in DB is treated as a valid.
			scenario: testhelpers.Scenario{
				Name:        "namespace-only-OPL: no filter available, recurses into leaf subject-sets",
				Opl:         `class User implements Namespace{} class Group implements Namespace{} class File implements Namespace{}`,
				InputTuples: inputTuples,
			},
			checkInput:         "File:f1#viewers@User:Bob",
			expectedMembership: check.NotMember,
			expectedTrace: func(t testing.TB) *trace.Node {
				return trace.CheckNode(t, trace.ResultNotMember, "File:f1#viewers@User:Bob",
					trace.DirectNode(t, trace.ResultNotMember, "File:f1#viewers@User:Bob"),
					trace.ExpandNode(t, trace.ResultNotMember, "File:f1#viewers@User:Bob", 1,
						trace.CheckNode(t, trace.ResultNotMember, "Group:g1#members@User:Bob",
							trace.ExpandNode(t, trace.ResultNotMember, "Group:g1#members@User:Bob", 1,
								// Redundant: the engine doesn't know Alice is a leaf subject,
								// so it creates a sub-check that always returns not_member.
								trace.CheckNode(t, trace.ResultNotMember, "User:Alice#@User:Bob",
									trace.ExpandNode(t, trace.ResultNotMember, "User:Alice#@User:Bob", 0),
								),
							),
						),
					),
				)
			},
		},
		{
			// In non-strict mode the OPL filter is not applied, so the engine recurses
			// into every subject-set pointer including leaf Users stored as empty-relation
			// SubjectSets. This matches the namespace-only behavior.
			scenario: testhelpers.Scenario{
				Name:        "valid-OPL: without filter, recurses into leaf subject-sets",
				Opl:         typedOPL,
				InputTuples: inputTuples,
			},
			checkInput:         "File:f1#viewers@User:Bob",
			expectedMembership: check.NotMember,
			expectedTrace: func(t testing.TB) *trace.Node {
				return trace.CheckNode(t, trace.ResultNotMember, "File:f1#viewers@User:Bob",
					trace.DirectNode(t, trace.ResultNotMember, "File:f1#viewers@User:Bob"),
					trace.ExpandNode(t, trace.ResultNotMember, "File:f1#viewers@User:Bob", 1,
						trace.CheckNode(t, trace.ResultNotMember, "Group:g1#members@User:Bob",
							trace.ExpandNode(t, trace.ResultNotMember, "Group:g1#members@User:Bob", 1,
								trace.CheckNode(t, trace.ResultNotMember, "User:Alice#@User:Bob",
									trace.ExpandNode(t, trace.ResultNotMember, "User:Alice#@User:Bob", 0),
								),
							),
						),
					),
				)
			},
		},
		{
			scenario: testhelpers.Scenario{
				Name:        "valid-OPL-strict: with filter, skips leaf subject-sets",
				Strict:      true,
				Opl:         typedOPL,
				InputTuples: inputTuples,
			},
			checkInput:         "File:f1#viewers@User:Bob",
			expectedMembership: check.NotMember,
			expectedTrace: func(t testing.TB) *trace.Node {
				return trace.CheckNode(t, trace.ResultNotMember, "File:f1#viewers@User:Bob",
					trace.DirectNode(t, trace.ResultNotMember, "File:f1#viewers@User:Bob"),
					trace.ExpandNode(t, trace.ResultNotMember, "File:f1#viewers@User:Bob", 1,
						trace.CheckNode(t, trace.ResultNotMember, "Group:g1#members@User:Bob",
							trace.ExpandNode(t, trace.ResultNotMember, "Group:g1#members@User:Bob", 0),
						),
					),
				)
			},
		},
	}

	for _, tt := range tests {
		tt.scenario.Run(t, func(t *testing.T, reg driver.Registry) {
			t.Parallel()

			e := trace.NewEngine(reg)
			res, tree := e.CheckRelationTupleWithTrace(t.Context(), testhelpers.TupleFromString(t, tt.checkInput), 100)
			require.Equal(t, tt.expectedMembership, res.Membership)
			require.NoError(t, res.Err)
			require.Equal(t, trace.SortNode(tt.expectedTrace(t)).String(), trace.SortNode(trace.StripTiming(tree)).String())
		})
	}
}

// TestExpandSubjectStep_StaleSubjectSet shows how stale subject-set tuples behave
// differently in strict vs non-strict mode.
//
// Scenario: OPL declares viewers: SubjectSet<Group, "admins">[], but the DB still
// contains a tuple for the old "members" relation. In strict mode the OPL filter
// blocks the stale pointer and access is denied. In non-strict mode no filter is
// applied, so the stale pointer is followed and access is incorrectly granted.
// Clients should migrate to strict mode to prevent stale tuples from granting access.
func TestExpandSubjectStep_StaleSubjectSet(t *testing.T) {
	t.Parallel()

	opl := `
		class User implements Namespace{}
		class Group implements Namespace{
			related: {
				admins: User[]
			}
		}
		class File implements Namespace{
			related: {
				viewers: SubjectSet<Group, "admins">[]
			}
		}
	`
	inputTuples := []string{
		// "members" was removed from OPL but the tuple was not cleaned up in the DB.
		"File:f1#viewers@Group:g1#members",
		"Group:g1#members@User:Alice",
	}

	tests := []struct {
		scenario           testhelpers.Scenario
		expectedMembership check.Membership
		expectedTrace      func(t testing.TB) *trace.Node
	}{
		{
			// Without the OPL filter, the stale "members" pointer is followed and
			// Alice is found as a direct member of Group:g1#members.
			scenario: testhelpers.Scenario{
				Name:        "valid-OPL: stale subject-set tuples are followed without OPL filter",
				Opl:         opl,
				InputTuples: inputTuples,
			},
			expectedMembership: check.IsMember,
			expectedTrace: func(t testing.TB) *trace.Node {
				return trace.CheckNode(t, trace.ResultMember, "File:f1#viewers@User:Alice",
					trace.DirectNode(t, trace.ResultNotMember, "File:f1#viewers@User:Alice"),
					trace.ExpandNode(t, trace.ResultMember, "File:f1#viewers@User:Alice", 1,
						trace.FoundNode(t, "Group:g1#members@User:Alice"),
					),
				)
			},
		},
		{
			// In strict mode the OPL filter is applied: only SubjectSet<Group,"admins">
			// pointers are expanded. The stale "members" tuple is filtered out.
			scenario: testhelpers.Scenario{
				Name:        "valid-OPL-strict: stale subject-set tuples are filtered by OPL",
				Strict:      true,
				Opl:         opl,
				InputTuples: inputTuples,
			},
			expectedMembership: check.NotMember,
			expectedTrace: func(t testing.TB) *trace.Node {
				return trace.CheckNode(t, trace.ResultNotMember, "File:f1#viewers@User:Alice",
					trace.DirectNode(t, trace.ResultNotMember, "File:f1#viewers@User:Alice"),
					// tuples_loaded=0: expand found no valid subject-set tuples with "admins" relation.
					trace.ExpandNode(t, trace.ResultNotMember, "File:f1#viewers@User:Alice", 0),
				)
			},
		},
	}

	for _, tt := range tests {
		tt.scenario.Run(t, func(t *testing.T, reg driver.Registry) {
			t.Parallel()

			e := trace.NewEngine(reg)
			res, tree := e.CheckRelationTupleWithTrace(t.Context(), testhelpers.TupleFromString(t, "File:f1#viewers@User:Alice"), 100)
			require.Equal(t, tt.expectedMembership, res.Membership)
			require.NoError(t, res.Err)
			require.Equal(t, trace.SortNode(tt.expectedTrace(t)).String(), trace.SortNode(trace.StripTiming(tree)).String())
		})
	}
}
