// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check

import (
	"context"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/ketoapi"
)

type (
	Result struct {
		Membership Membership
		Tree       *ketoapi.Tree[*relationtuple.RelationTuple]
		Err        error
	}

	Membership int
)

//go:generate stringer -type Membership
const (
	MembershipUnknown Membership = iota
	IsMember
	NotMember
)

var (
	ResultUnknown   = Result{Membership: MembershipUnknown}
	ResultIsMember  = Result{Membership: IsMember}
	ResultNotMember = Result{Membership: NotMember}
)

// CheckRequest carries all inputs for a single check evaluation step.
type CheckRequest struct {
	Tuple     *relationtuple.RelationTuple
	RestDepth int
}

func (r CheckRequest) WithDepth(d int) CheckRequest {
	r.RestDepth = d
	return r
}

// PlannedStep pairs a Step with the CheckRequest it should run against.
type PlannedStep struct {
	Step Step
	Req  CheckRequest
}

// StepRunner orchestrates execution of multiple steps with short-circuit semantics.
// Typical usage:
//
//	runner := ex.NewUnionRunner(ctx)
//	for _, ps := range steps {
//		if runner.Done() { break }
//		runner.Add(ps)
//	}
//	return runner.Result()
type StepRunner interface {
	// Add enqueues a step for execution. Returns early if Done() is true.
	// No-op if the runner has already reached a decisive result or been cancelled.
	Add(PlannedStep)

	// Done reports whether a decisive result was reached or context cancelled.
	// Use to short-circuit the enqueue loop.
	Done() bool

	// Ctx returns the runner's context, usable for shared state like visited tracking.
	// The context is cancelled when a decisive result is reached (to signal siblings).
	Ctx() context.Context

	// Result waits for all in-flight steps to complete and returns the aggregated result.
	// Call it after the Add loop finishes (or breaks early on Done()).
	Result() Result
}

// Executor coordinates permission check execution with middleware support.
// It runs individual steps (Run), sequences steps with union/intersection logic
// (RunUnion, RunIntersection), and supports custom execution strategies via
// middleware. All Executor implementations are stateless; dependencies come
// from EngineDependencies.
//
// Typical usage:
//
//	result := executor.Run(ctx, DirectStep{}, req)
//	unionResult := executor.RunUnion(ctx, step1, step2)
//	intersectionResult := executor.RunIntersection(ctx, step1, step2)
type Executor interface {
	// Run executes a single step through the middleware chain.
	Run(ctx context.Context, step Step, req CheckRequest) Result

	// RunUnion runs steps with OR semantics: returns IsMember on first match,
	// short-circuits siblings, falls through to NotMember if none match.
	RunUnion(ctx context.Context, steps ...PlannedStep) Result

	// RunIntersection runs steps with AND semantics: returns NotMember or Error
	// on first failure, short-circuits siblings.
	RunIntersection(ctx context.Context, steps ...PlannedStep) Result

	// NewUnionRunner creates a StepRunner for streaming union execution.
	// It is useful when Steps are created dynamically and the full list is not known upfront.
	NewUnionRunner(ctx context.Context) StepRunner

	Deps() EngineDependencies
}

// Middleware intercepts a step execution within the pipeline. next continues
// down the remaining middlewares and ultimately calls step.Execute.
type Middleware func(
	ctx context.Context,
	step Step,
	req CheckRequest,
	next func(context.Context, Step, CheckRequest) Result,
) Result

// Checker performs a full permission check for a single relation tuple.
// Implemented by Executor in the step package.
type (
	CheckerProvider interface {
		Checker() Checker
	}
	Checker interface {
		CheckRelationTuple(ctx context.Context, r *relationtuple.RelationTuple, restDepth int) Result
	}
)
