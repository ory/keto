// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package step

import (
	"context"
	"sync"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/relationtuple"
)

type (
	relationTuple = relationtuple.RelationTuple
	query         = relationtuple.RelationQuery
)

// runMode defines the decisiveness and default result semantics for an execution mode.
//
// IsDecisive determines whether a step result ends execution early (short-circuit).
//   - Union: IsMember or any Error is decisive.
//   - Intersection: NotMember or any Error is decisive.
//
// NoDecisionResult is returned when all steps ran without triggering a short-circuit.
//   - Union: NotMember (no member found among alternatives).
//   - Intersection: IsMember (all conditions met).
//
// NoStepsResult is returned when no steps were added (empty input slice).
//   - Both: NotMember (no tuples to check).
type runMode interface {
	IsDecisive(check.Result) bool
	NoDecisionResult() check.Result
	NoStepsResult() check.Result
}

// unionRunMode implements OR semantics: first IsMember or Error wins.
type unionRunMode struct{}

func (unionRunMode) IsDecisive(r check.Result) bool {
	return r.Membership == check.IsMember || r.Err != nil
}
func (unionRunMode) NoDecisionResult() check.Result { return check.ResultNotMember }
func (unionRunMode) NoStepsResult() check.Result    { return check.ResultNotMember }

// intersectionRunMode implements AND semantics: first NotMember or Error wins.
type intersectionRunMode struct{}

func (intersectionRunMode) IsDecisive(r check.Result) bool {
	return r.Membership != check.IsMember || r.Err != nil
}
func (intersectionRunMode) NoDecisionResult() check.Result { return check.ResultIsMember }
func (intersectionRunMode) NoStepsResult() check.Result    { return check.ResultNotMember }

var (
	unionMode        runMode = unionRunMode{}
	intersectionMode runMode = intersectionRunMode{}
)

// Executor implements check.Executor and runs Steps with short-circuit
// union/intersection semantics and a middleware chain.
type Executor struct {
	dep         check.EngineDependencies
	middlewares []check.Middleware
}

func NewExecutor(dep check.EngineDependencies, middlewares ...check.Middleware) *Executor {
	return &Executor{dep: dep, middlewares: middlewares}
}

func (e *Executor) Deps() check.EngineDependencies { return e.dep }

func (e *Executor) Run(ctx context.Context, s check.Step, req check.CheckRequest) check.Result {
	return e.runAt(ctx, s, req, 0)
}

// runAt executes the middleware chain at position i, then the step's Execute method.
// It implements a functional middleware chain: each middleware receives a next() function
// that continues the chain.
func (e *Executor) runAt(ctx context.Context, s check.Step, req check.CheckRequest, i int) check.Result {
	if i >= len(e.middlewares) {
		return s.Execute(ctx, req, e)
	}
	return e.middlewares[i](ctx, s, req, func(ctx context.Context, s check.Step, req check.CheckRequest) check.Result {
		return e.runAt(ctx, s, req, i+1)
	})
}

func (e *Executor) RunUnion(ctx context.Context, steps ...check.PlannedStep) check.Result {
	return runAll(ctx, e, unionMode, 1, steps...)
}

func (e *Executor) RunIntersection(ctx context.Context, steps ...check.PlannedStep) check.Result {
	return runAll(ctx, e, intersectionMode, 1, steps...)
}

func (e *Executor) NewUnionRunner(ctx context.Context) check.StepRunner {
	return newStepGroup(ctx, 1, unionMode, e)
}

// CheckRelationTuple checks if the tuple's subject has the relation on the object.
func (e *Executor) CheckRelationTuple(ctx context.Context, r *relationTuple, restDepth int) check.Result {
	restDepth = clampDepth(ctx, e.dep, restDepth)
	return e.Run(ctx, IsAllowedStep{}, check.CheckRequest{
		Tuple:     r,
		RestDepth: restDepth,
	})
}

func clampDepth(ctx context.Context, c config.Provider, restDepth int) int {
	if globalMaxDepth := c.Config(ctx).MaxReadDepth(); restDepth <= 0 || globalMaxDepth < restDepth {
		return globalMaxDepth
	}
	return restDepth
}

func runAll(ctx context.Context, ex check.Executor, mode runMode, limit int, steps ...check.PlannedStep) check.Result {
	if len(steps) == 0 {
		return mode.NoStepsResult()
	}
	g := newStepGroup(ctx, limit, mode, ex)

	for _, ps := range steps {
		if g.Done() {
			break
		}
		g.Add(ps)
	}
	return g.Result()
}

// stepGroup runs sub-checks with short-circuit semantics. When sequential (limit=1)
// steps run inline with no goroutine overhead. When limit>1 steps run
// concurrently, cancelling siblings on the first decisive result.
// stepGroup implements check.StepRunner.
type stepGroup struct {
	ctx        context.Context
	cancel     context.CancelFunc
	ex         check.Executor
	sequential bool

	sem chan struct{}
	wg  sync.WaitGroup

	mu     sync.Mutex
	result check.Result
	set    bool

	mode runMode
}

// newStepGroup creates a StepGroup that runs steps with the given concurrencyLimit and mode.
// If concurrencyLimit < 1, it's clamped to 1 (sequential). Sequential execution (concurrencyLimit=1)
// runs steps inline without goroutines; concurrent execution spawns workers up to
// concurrencyLimit. The returned context is derived from ctx via WithCancel and is cancelled
// when a decisive result is reached (triggering sibling cancellation).
func newStepGroup(ctx context.Context, concurrencyLimit int, mode runMode, ex check.Executor) *stepGroup {
	if concurrencyLimit < 1 {
		concurrencyLimit = 1
	}
	g := &stepGroup{
		ex:         ex,
		sequential: concurrencyLimit == 1,
		mode:       mode,
	}
	g.ctx, g.cancel = context.WithCancel(ctx)
	if !g.sequential {
		g.sem = make(chan struct{}, concurrencyLimit)
	}
	return g
}

// Add submits ps for execution. No-op if Done returns true.
func (g *stepGroup) Add(ps check.PlannedStep) {
	g.add(func(ctx context.Context) check.Result {
		return g.ex.Run(ctx, ps.Step, ps.Req)
	})
}

func (g *stepGroup) add(f func(context.Context) check.Result) {
	if g.sequential {
		if g.Done() {
			return
		}
		g.handle(f(g.ctx))
		return
	}

	// currently dead path as we always use concurrentLimit = 1
	select {
	case g.sem <- struct{}{}:
	case <-g.ctx.Done():
		return
	}
	if g.Done() {
		<-g.sem
		return
	}
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		defer func() { <-g.sem }()
		g.handle(f(g.ctx))
	}()
}

func (g *stepGroup) handle(r check.Result) {
	if g.mode.IsDecisive(r) && g.trySetResult(r) {
		g.cancel()
	}
}

// Done reports whether a decisive result has been reached.
func (g *stepGroup) Done() bool {
	select {
	case <-g.ctx.Done():
		return true
	default:
		return false
	}
}

// Ctx returns the group's context.
func (g *stepGroup) Ctx() context.Context { return g.ctx }

// Result waits for all in-flight steps to complete and returns the aggregated result.
// Safe to call multiple times; subsequent calls return the same value.
func (g *stepGroup) Result() check.Result {
	g.wg.Wait()
	// Capture ctx.Err() before g.cancel() to distinguish an external cancellation
	// (deadline exceeded, parent cancelled) from the internal cancel used to signal siblings.
	ctxErr := g.ctx.Err()
	g.cancel()
	g.mu.Lock()
	defer g.mu.Unlock()
	if !g.set {
		if ctxErr != nil {
			g.result = check.Result{Membership: check.MembershipUnknown, Err: ctxErr}
		} else {
			g.result = g.mode.NoDecisionResult()
		}
		g.set = true
	}
	return g.result
}

func (g *stepGroup) trySetResult(r check.Result) bool {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.set {
		return false
	}
	g.set = true
	g.result = r
	return true
}
