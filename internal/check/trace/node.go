// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package trace

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/relationtuple"
)

type (
	NodeKind   string
	NodeResult string
)

const (
	NodeIsAllowed     NodeKind = "is_allowed"
	NodeDirect        NodeKind = "direct"
	NodeMultiDirect   NodeKind = "direct_multi"
	NodeExpandSubject NodeKind = "expand_subject"
	NodeFoundSubject  NodeKind = "expand_found"
	NodeComputed      NodeKind = "computed"
	NodeTraverse      NodeKind = "traverse"
	NodeUnion         NodeKind = "union"
	NodeIntersection  NodeKind = "intersection"
	NodeInvert        NodeKind = "invert"
)

func (k NodeKind) String() string { return string(k) }

const (
	// ResultMember indicates the subject is a member.
	ResultMember NodeResult = "member"
	// ResultNotMember indicates the subject is not a member.
	ResultNotMember NodeResult = "not_member"
	// ResultError indicates an error occurred during the check.
	ResultError NodeResult = "error"
	// ResultSkipped indicates the check was registered but never executed
	// because a sibling check resolved the parent before this one ran.
	ResultSkipped NodeResult = "skipped"
	// ResultCancelled indicates the check was cancelled mid-execution
	// because a sibling check resolved.
	ResultCancelled NodeResult = "cancelled"
	// ResultUnknown indicates the check was cut short because restDepth
	// reached zero; membership could not be determined.
	ResultUnknown NodeResult = "unknown"
)

func (r NodeResult) prunable() bool {
	return r == ResultNotMember || r == ResultSkipped
}

// MaxExpandChildren is the maximum number of non-significant children kept on
// expand_subject and traverse nodes. Children whose result is member, error,
// or cancelled are always kept regardless of this limit.
const MaxExpandChildren = 5

// Node is one step in the debug tree produced by a permission check. It
// records what kind of operation was performed, what the input tuple was, what
// the result was, how many database records were fetched, and any sub-steps
// that were spawned.
type Node struct {
	Kind   NodeKind
	Tuple  *relationtuple.RelationTuple
	Result NodeResult

	// Relations holds the batched relation names for NodeMultiDirect nodes.
	Relations []string

	// TuplesLoaded is the number of database records fetched by this node.
	TuplesLoaded int

	// TruncatedChildren is the number of children that were dropped to keep
	// the tree size manageable.
	TruncatedChildren int

	Children []*Node

	// Duration is the wall-clock time spent in this node (including children).
	Duration time.Duration

	// parent is set when this node is appended to a parent, so that
	// dropIfOverLimit can remove it after the result is known.
	parent *Node

	// mu protects Children, TruncatedChildren, TuplesLoaded, and prunableCount.
	mu sync.Mutex

	// prunableCount counts how many children have a prunable result.
	prunableCount int
}

func (n *Node) setResult(result check.Result) {
	if result.Err != nil {
		n.Result = ResultError
		if errors.Is(result.Err, context.Canceled) {
			n.Result = ResultCancelled
		}
		return
	}
	switch result.Membership {
	case check.IsMember:
		n.Result = ResultMember
	case check.NotMember:
		n.Result = ResultNotMember
	default:
		n.Result = ResultUnknown
	}
}

// dropIfOverLimit removes this node from its parent when the parent is an
// expand or TTU node that already has MaxExpandChildren prunable children.
func (n *Node) dropIfOverLimit() {
	parent := n.parent
	if parent == nil {
		return
	}
	if parent.Kind != NodeExpandSubject && parent.Kind != NodeTraverse {
		return
	}
	if !n.Result.prunable() {
		return
	}
	parent.mu.Lock()
	defer parent.mu.Unlock()
	parent.prunableCount++
	if parent.prunableCount <= MaxExpandChildren {
		return
	}
	for i, c := range parent.Children {
		if c == n {
			parent.Children = append(parent.Children[:i], parent.Children[i+1:]...)
			break
		}
	}
	parent.TruncatedChildren++
}

func (n *Node) recordTuplesLoaded(count int) {
	n.mu.Lock()
	n.TuplesLoaded += count
	n.mu.Unlock()
}

func (n *Node) recordFoundSubject(tuple *relationtuple.RelationTuple) {
	child := &Node{Kind: NodeFoundSubject, Tuple: tuple, Result: ResultMember}
	n.mu.Lock()
	n.Children = append(n.Children, child)
	n.mu.Unlock()
}

// String is a human-friendly representation of the debug tree, useful during testing.
func (n *Node) String() string {
	if n == nil {
		return "<nil>"
	}

	var b strings.Builder

	var walk func(cur *Node, depth int)
	walk = func(cur *Node, depth int) {
		for range depth {
			b.WriteString("  ")
		}

		input := ""
		if cur.Tuple != nil {
			input = cur.Tuple.String()
			if len(cur.Relations) > 0 {
				input = strings.Replace(input, "#"+cur.Tuple.Relation+"@", "#{"+strings.Join(cur.Relations, ",")+"}"+"@", 1)
			}
		}
		b.WriteString(string(cur.Kind))
		b.WriteByte(' ')
		b.WriteString(strconv.Quote(input))
		b.WriteString(" -> ")
		b.WriteString(string(cur.Result))

		if cur.TuplesLoaded > 0 {
			b.WriteString(" tuples_loaded=")
			b.WriteString(strconv.Itoa(cur.TuplesLoaded))
		}
		b.WriteByte('\n')

		for _, child := range cur.Children {
			walk(child, depth+1)
		}
	}

	walk(n, 0)
	return strings.TrimSuffix(b.String(), "\n")
}
