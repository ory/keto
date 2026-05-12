// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package trace

import (
	"context"

	"github.com/ory/x/otelx"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/check/step"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/x/events"
)

// Engine builds a runtime with tracing enabled and returns the debug Node
// tree alongside the check result.
type Engine struct {
	deps check.EngineDependencies
}

func NewEngine(d check.EngineDependencies) *Engine {
	return &Engine{deps: d}
}

func (e *Engine) d() check.EngineDependencies { return e.deps }

// CheckIsMemberWithTrace is like Engine.CheckIsMember but also returns a debug
// tree that traces every sub-check performed.
func (e *Engine) CheckIsMemberWithTrace(ctx context.Context, tuple *relationtuple.RelationTuple, restDepth int) (bool, *Node, error) {
	res, tree := e.CheckRelationTupleWithTrace(ctx, tuple, restDepth)
	return res.Membership == check.IsMember, tree, res.Err
}

// CheckRelationTupleWithTrace returns the check result and the debug tree.
func (e *Engine) CheckRelationTupleWithTrace(ctx context.Context, tuple *relationtuple.RelationTuple, restDepth int) (res check.Result, tree *Node) {
	ctx, span := e.d().Tracer(ctx).Tracer().Start(ctx, "Engine.CheckRelationTupleWithTrace")
	defer otelx.End(span, &res.Err)

	session := NewSession()
	executor := step.NewExecutor(e.deps, session.Middleware())

	res = executor.CheckRelationTuple(ctx, tuple, restDepth)
	session.Wait()
	if res.Err == nil {
		span.AddEvent(events.NewPermissionsChecked(ctx))
	}
	return res, session.Root()
}
