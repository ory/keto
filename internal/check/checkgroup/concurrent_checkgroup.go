package checkgroup

import (
	"context"
	"sync"
)

// A concurrentCheckgroup is a collection of goroutines performing checks.
type concurrentCheckgroup struct {
	ctx context.Context

	cancel     context.CancelFunc
	resultCh   chan Result
	subcheckCh chan Result
	once       sync.Once
	counts     struct {
		sync.RWMutex
		totalChecks     int
		finishedChecks  int
		resultRequested bool
	}
}

func NewConcurrent(ctx context.Context) Checkgroup {
	return &concurrentCheckgroup{ctx: ctx}
}

func (g *concurrentCheckgroup) incrementRunningCheckCount() {
	g.counts.Lock()
	defer g.counts.Unlock()
	g.counts.totalChecks++
}
func (g *concurrentCheckgroup) incrementFinishedCheckCount() {
	g.counts.Lock()
	defer g.counts.Unlock()
	g.counts.finishedChecks++
}

func (g *concurrentCheckgroup) allCheckFinished() bool {
	g.counts.RLock()
	defer g.counts.RUnlock()
	return g.counts.resultRequested && g.counts.totalChecks == g.counts.finishedChecks
}

func (g *concurrentCheckgroup) noChecksAdded() bool {
	g.counts.RLock()
	defer g.counts.RUnlock()
	return g.counts.totalChecks == 0
}

// If freeze is called, no more checks can be added.
func (g *concurrentCheckgroup) freeze() {
	g.counts.Lock()
	defer g.counts.Unlock()
	g.counts.resultRequested = true
}

func (g *concurrentCheckgroup) frozen() bool {
	g.counts.RLock()
	defer g.counts.RUnlock()
	return g.counts.resultRequested
}

func (g *concurrentCheckgroup) startConsumer() {
	g.once.Do(func() {
		g.subcheckCh = make(chan Result)
		g.resultCh = make(chan Result)
		g.ctx, g.cancel = context.WithCancel(g.ctx)
		go func() {
			for {
				select {
				case result := <-g.subcheckCh:
					g.incrementFinishedCheckCount()
					if result.Err != nil || result.Membership == IsMember || g.allCheckFinished() {
						g.resultCh <- result
						g.cancel()
						return
					}

				case <-g.ctx.Done():
					g.resultCh <- Result{Err: context.Canceled}
					g.cancel()
					return
				}
			}
		}()
	})
}

func (g *concurrentCheckgroup) Done() bool {
	select {
	case <-g.ctx.Done():
		return true
	default:
		return false
	}
}

// Add adds the Func to the checkgroup and starts running it.
func (g *concurrentCheckgroup) Add(check Func) {
	if g.frozen() {
		panic("Trying to Add(), but result was already requested.")
	}
	if g.Done() {
		return
	}
	g.startConsumer()
	g.incrementRunningCheckCount()
	go check(g.ctx, g.subcheckCh)
}

// SetIsMember makes the checkgroup emit "IsMember" directly.
func (g *concurrentCheckgroup) SetIsMember() {
	g.Add(IsMemberFunc)
}

// Result returns the Result, possibly blocking.
func (g *concurrentCheckgroup) Result() Result {
	g.startConsumer()
	g.freeze()
	if g.noChecksAdded() {
		g.cancel()
		return Result{Membership: NotMember}
	}

	return <-g.resultCh
}

// CheckFunc returns a `Func` that writes the result to the result channel.
func (g *concurrentCheckgroup) CheckFunc() Func {
	g.startConsumer()
	g.freeze()
	if g.noChecksAdded() {
		g.cancel()
		return NotMemberFunc
	}

	return func(ctx context.Context, resultCh chan<- Result) {
		select {
		case result := <-g.resultCh:
			resultCh <- result
		case <-ctx.Done():
			resultCh <- Result{Err: context.Canceled}
			g.cancel()
		}
	}
}
