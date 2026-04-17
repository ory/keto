// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package expand

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/ketoapi"
)

func AssertExternalTreesAreEqual(t *testing.T, expected, actual *ketoapi.Tree[*ketoapi.RelationTuple]) {
	t.Helper()
	assert.Truef(t, treesAreEqual(t, expected, actual),
		"expected:\n%+v\n\nactual:\n%+v", expected, actual)
}

// TODO(hperl): Refactor to generic tree equality helper.
func treesAreEqual(t *testing.T, expected, actual *ketoapi.Tree[*ketoapi.RelationTuple]) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	if expected.Type != actual.Type {
		t.Logf("expected type %q, actual type %q", expected.Type, actual.Type)
		return false
	}
	if !assert.ObjectsAreEqual(expected.Tuple.SubjectID, actual.Tuple.SubjectID) || !assert.ObjectsAreEqual(expected.Tuple.SubjectSet, actual.Tuple.SubjectSet) {
		t.Logf("expected subject: %+v %+v, actual subject: %+v %+v", expected.Tuple.SubjectID, expected.Tuple.SubjectSet, actual.Tuple.SubjectID, actual.Tuple.SubjectSet)
		return false
	}
	if len(expected.Children) != len(actual.Children) {
		t.Logf("expected len(children)=%d, actual len(children)=%d", len(expected.Children), len(actual.Children))
		return false
	}

	// For children, we check for equality disregarding the order
outer:
	for _, expectedChild := range expected.Children {
		for _, actualChild := range actual.Children {
			if treesAreEqual(t, expectedChild, actualChild) {
				continue outer
			}
		}
		t.Logf("expected child:\n%+v\n\nactual child:\n%+v", expectedChild, actual)
		return false
	}
	return true
}

// formatTree returns a human-readable, indented representation of a tree.
func formatTree(tree *relationtuple.Tree, indent string) string {
	if tree == nil {
		return indent + "<nil>\n"
	}
	var b strings.Builder
	fmt.Fprintf(&b, "%s(%s) %+v", indent, tree.Type, tree.Subject)
	if e := tree.Truncation; e != nil {
		fmt.Fprintf(&b, " [truncation: reason=%s cursor=%s]", e.Reason, formatExpandCursor(e.Cursor))
	}
	b.WriteByte('\n')
	for _, child := range tree.Children {
		b.WriteString(formatTree(child, indent+"  "))
	}
	return b.String()
}

func formatExpandCursor(c *relationtuple.ExpandCursor) string {
	if c == nil {
		return "<nil>"
	}
	var parts []string
	parts = append(parts, "kind="+string(c.Kind))
	parts = append(parts, "subject_set="+c.SubjectSet.String())
	if c.TraverseRelation != nil {
		parts = append(parts, "traverseRel="+*c.TraverseRelation)
	}
	return "{" + strings.Join(parts, " ") + "}"
}

// sortTree recursively sorts the children of a tree by subject string, in-place.
func sortTree(tree *relationtuple.Tree) {
	if tree == nil {
		return
	}
	for _, child := range tree.Children {
		sortTree(child)
	}
	if tree.Children == nil {
		return
	}
	slices.SortFunc(tree.Children, func(a, b *relationtuple.Tree) int {
		if a.Subject == nil && b.Subject == nil {
			return 0
		}
		if a.Subject == nil {
			return 1
		}
		if b.Subject == nil {
			return -1
		}
		return strings.Compare(a.Subject.String(), b.Subject.String())
	})
}

// AssertInternalTreesAreEqual compares two trees for equality, ignoring child order.
func AssertInternalTreesAreEqual(t *testing.T, expected, actual *relationtuple.Tree) bool {
	t.Helper()
	sortTree(expected)
	sortTree(actual)
	return assert.Equal(t, formatTree(expected, ""), formatTree(actual, ""))
}

// RequireInternalTreesAreEqual compares two trees for equality, ignoring child order.
func RequireInternalTreesAreEqual(t *testing.T, expected, actual *relationtuple.Tree) {
	t.Helper()
	sortTree(expected)
	sortTree(actual)
	require.Equal(t, formatTree(expected, ""), formatTree(actual, ""))
}
