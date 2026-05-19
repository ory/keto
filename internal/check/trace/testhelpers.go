// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package trace

import (
	"slices"
	"strings"
	"testing"

	"github.com/ory/keto/internal/testhelpers"
)

// SortNode sorts children of NodeExpandSubject and NodeTraverse nodes by
// Tuple.String() recursively. Both node kinds have non-deterministic child
// order because the underlying queries order by shard_id (a random UUIDv4
// assigned at write time). All other sibling orderings are engine-controlled
// and must be asserted as-is.
func SortNode(n *Node) *Node {
	if n == nil {
		return nil
	}
	if n.Kind == NodeExpandSubject || n.Kind == NodeTraverse {
		slices.SortFunc(n.Children, func(a, b *Node) int {
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
		SortNode(child)
	}
	return n
}

// StripTiming zeroes out Duration on every node in the tree so that test
// assertions are not sensitive to wall-clock time.
func StripTiming(n *Node) *Node {
	if n == nil {
		return nil
	}
	n.Duration = 0
	for _, child := range n.Children {
		StripTiming(child)
	}
	return n
}

// CheckNode builds a NodeIsAllowed node.
func CheckNode(t testing.TB, result NodeResult, tuple string, children ...*Node) *Node {
	t.Helper()
	return &Node{Kind: NodeIsAllowed, Tuple: testhelpers.TupleFromString(t, tuple), Result: result, Children: children}
}

// DirectNode builds a NodeDirect leaf (no children).
func DirectNode(t testing.TB, result NodeResult, tuple string) *Node {
	t.Helper()
	return &Node{Kind: NodeDirect, Tuple: testhelpers.TupleFromString(t, tuple), Result: result}
}

// ExpandNode builds a NodeExpandSubject node.
func ExpandNode(t testing.TB, result NodeResult, tuple string, tuplesLoaded int, children ...*Node) *Node {
	t.Helper()
	return &Node{Kind: NodeExpandSubject, Tuple: testhelpers.TupleFromString(t, tuple), Result: result, TuplesLoaded: tuplesLoaded, Children: children}
}

// FoundNode builds a NodeFoundSubject leaf (no children).
func FoundNode(t testing.TB, tuple string) *Node {
	t.Helper()
	return &Node{Kind: NodeFoundSubject, Tuple: testhelpers.TupleFromString(t, tuple), Result: ResultMember}
}

// UnionNode builds a NodeUnion node.
func UnionNode(t testing.TB, result NodeResult, tuple string, children ...*Node) *Node {
	t.Helper()
	return &Node{Kind: NodeUnion, Tuple: testhelpers.TupleFromString(t, tuple), Result: result, Children: children}
}

// IntersectionNode builds a NodeIntersection node.
func IntersectionNode(t testing.TB, result NodeResult, tuple string, children ...*Node) *Node {
	t.Helper()
	return &Node{Kind: NodeIntersection, Tuple: testhelpers.TupleFromString(t, tuple), Result: result, Children: children}
}

// InvertNode builds a NodeInvert node.
func InvertNode(t testing.TB, result NodeResult, tuple string, children ...*Node) *Node {
	t.Helper()
	return &Node{Kind: NodeInvert, Tuple: testhelpers.TupleFromString(t, tuple), Result: result, Children: children}
}

// ComputedNode builds a NodeComputed node.
func ComputedNode(t testing.TB, result NodeResult, tuple string, children ...*Node) *Node {
	t.Helper()
	return &Node{Kind: NodeComputed, Tuple: testhelpers.TupleFromString(t, tuple), Result: result, Children: children}
}

// TraverseNode builds a NodeTraverse node.
func TraverseNode(t testing.TB, result NodeResult, tuple string, tuplesLoaded int, children ...*Node) *Node {
	t.Helper()
	return &Node{Kind: NodeTraverse, Tuple: testhelpers.TupleFromString(t, tuple), Result: result, TuplesLoaded: tuplesLoaded, Children: children}
}

// MultiDirectNode builds a NodeMultiDirect leaf (no children).
func MultiDirectNode(t testing.TB, result NodeResult, tuple string, relations []string) *Node {
	t.Helper()
	return &Node{Kind: NodeMultiDirect, Tuple: testhelpers.TupleFromString(t, tuple), Relations: relations, Result: result}
}
