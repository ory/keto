package checkgroup

import (
	"context"
	"errors"
	"sync"
	"time"
)

type sequentialCheckgroup struct {
	ctx    context.Context
	lock   sync.Mutex
	checks []Func
	done   bool
}

func NewSequential(ctx context.Context) Checkgroup {
	return &sequentialCheckgroup{ctx: ctx}
}

func (g *sequentialCheckgroup) Done() bool {
	g.lock.Lock()
	defer g.lock.Unlock()
	return g.done
}

func (g *sequentialCheckgroup) Add(check Func) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.done {
		panic("already done")
	}
	g.checks = append(g.checks, check)
}

func (g *sequentialCheckgroup) SetIsMember() {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.checks = append(g.checks, IsMemberFunc)
}

func (g *sequentialCheckgroup) Result() Result {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.done = true

	if len(g.checks) == 0 {
		return ResultNotMember
	}
	resultCh := make(chan Result)

	for _, check := range g.checks {
		childCtx, cancel := context.WithTimeout(g.ctx, 1*time.Second)
		defer cancel()
		go check(childCtx, resultCh)
		select {
		case result := <-resultCh:
			if errors.Is(result.Err, context.DeadlineExceeded) {
				continue
			}
			if result.Err != nil || result.Membership == IsMember {
				return result
			}
		case <-g.ctx.Done():
			return Result{Err: context.Canceled}
		}
	}

	return ResultNotMember
}

func (g *sequentialCheckgroup) CheckFunc() Func {
	g.lock.Lock()
	defer g.lock.Unlock()

	return func(_ context.Context, resultCh chan<- Result) {
		resultCh <- g.Result()
	}
}
