// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package trace

import (
	"context"
	"sync"
	"time"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/check/step"
	"github.com/ory/keto/internal/namespace/ast"
	"github.com/ory/keto/internal/relationtuple"
)

// traceNodeKey is the context key used to store and retrieve the current
// trace node as execution descends through the step tree.
type traceNodeKey struct{}

func nodeFromCtx(ctx context.Context) *Node {
	n, _ := ctx.Value(traceNodeKey{}).(*Node)
	return n
}

func withTraceNode(ctx context.Context, n *Node) context.Context {
	return context.WithValue(ctx, traceNodeKey{}, n)
}

// Session collects the debug tree for a single runtime execution.
type Session struct {
	mu   sync.Mutex
	root *Node
	wg   sync.WaitGroup
}

func NewSession() *Session {
	return &Session{}
}

func (t *Session) Root() *Node {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.root
}

func (t *Session) Wait() {
	t.wg.Wait()
}

// nodeKindForStep maps a Step's execution-model kind to the trace-model NodeKind.
func nodeKindForStep(s check.Step) NodeKind {
	if s, ok := s.(step.RewriteStep); ok {
		if s.Rewrite.Operation == ast.OperatorOr {
			return NodeUnion
		}
		return NodeIntersection
	}
	switch s.Kind() {
	case check.StepIsAllowed:
		return NodeIsAllowed
	case check.StepDirect:
		return NodeDirect
	case check.StepDirectMulti:
		return NodeMultiDirect
	case check.StepExpand:
		return NodeExpandSubject
	case check.StepComputed:
		return NodeComputed
	case check.StepTraverse:
		return NodeTraverse
	case check.StepInvert:
		return NodeInvert
	default:
		return NodeKind(s.Kind())
	}
}

// Middleware returns a Middleware that records a Node for every step execution,
// building the debug trace tree as execution unfolds.
func (t *Session) Middleware() check.Middleware {
	return func(ctx context.Context, s check.Step, req check.CheckRequest, next func(context.Context, check.Step, check.CheckRequest) check.Result) check.Result {
		t.wg.Add(1)
		defer t.wg.Done()

		tuple := req.Tuple
		var relations []string

		switch s := s.(type) {
		case step.ComputedUsersetStep:
			clone := *req.Tuple
			clone.Relation = s.Relation
			tuple = &clone
		case step.TraverseStep:
			clone := *req.Tuple
			clone.Relation = s.TTU.Relation
			tuple = &clone
		case step.DirectMultiStep:
			relations = s.Relations
		}

		node, ctx := t.newTraceNode(ctx, nodeKindForStep(s), tuple, relations...)
		switch s.Kind() {
		case check.StepExpand:
			ctx = check.WithTraceHooks(ctx, check.TraceHooks{
				TuplesLoaded: node.recordTuplesLoaded,
				SubjectFound: node.recordFoundSubject,
			})
		case check.StepTraverse:
			ctx = check.WithTraceHooks(ctx, check.TraceHooks{TuplesLoaded: node.recordTuplesLoaded})
		}

		start := time.Now()
		result := next(ctx, s, req)
		node.Duration = time.Since(start)
		node.setResult(result)
		node.dropIfOverLimit()

		return result
	}
}

// newTraceNode creates a Node, links it as a child of the current parent node
// in ctx (or sets it as the root if there is none), and returns both the node
// and a derived context with the new node set as the current parent.
func (t *Session) newTraceNode(ctx context.Context, kind NodeKind, tuple *relationtuple.RelationTuple, relations ...string) (*Node, context.Context) {
	node := &Node{
		Kind:      kind,
		Tuple:     tuple,
		Relations: relations,

		// Initial Result is Skipped, as it's possible that the step won't run at all
		// due to other sibling branch finishing with success.
		Result: ResultSkipped,
	}
	if parent := nodeFromCtx(ctx); parent != nil {
		node.parent = parent
		parent.mu.Lock()
		parent.Children = append(parent.Children, node)
		parent.mu.Unlock()
	} else {
		t.mu.Lock()
		t.root = node
		t.mu.Unlock()
	}
	return node, withTraceNode(ctx, node)
}
