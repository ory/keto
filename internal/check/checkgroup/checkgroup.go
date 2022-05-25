package checkgroup

import (
	"context"
	"sync"
)

// A Checkgroup is a collection of goroutines performing checks.
type Checkgroup struct {
	Ctx context.Context

	cancel     context.CancelFunc
	resultCh   chan Result
	subcheckCh chan Result
	once       sync.Once
	counts     struct {
		sync.RWMutex
		totalChecks    int
		finishedChecks int
	}
}

func New(ctx context.Context) *Checkgroup {
	return &Checkgroup{Ctx: ctx}
}

func (g *Checkgroup) incrementRunningCheckCount() {
	g.counts.Lock()
	defer g.counts.Unlock()
	g.counts.totalChecks++
}
func (g *Checkgroup) incrementFinishedCheckCount() {
	g.counts.Lock()
	defer g.counts.Unlock()
	g.counts.finishedChecks++
}

func (g *Checkgroup) allCheckFinished() bool {
	g.counts.RLock()
	defer g.counts.RUnlock()
	return g.counts.totalChecks == g.counts.finishedChecks
}

func (g *Checkgroup) startConsumer() {
	g.once.Do(func() {
		g.subcheckCh = make(chan Result)
		g.resultCh = make(chan Result)
		g.Ctx, g.cancel = context.WithCancel(g.Ctx)
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

				case <-g.Ctx.Done():
					g.resultCh <- Result{Err: context.Canceled}
					g.cancel()
					return
				}
			}
		}()
	})
}

func (g *Checkgroup) Done() bool {
	select {
	case <-g.Ctx.Done():
		return true
	default:
		return false
	}
}

// Add adds the Func to the checkgroup and starts running it.
func (g *Checkgroup) Add(check Func) {
	g.startConsumer()
	g.incrementRunningCheckCount()
	go check(g.Ctx, g.subcheckCh)
}

func (g *Checkgroup) noChecksAdded() bool {
	return g.counts.totalChecks == 0
}

// Result returns the Result, possibly blocking.
func (g *Checkgroup) Result() Result {
	g.startConsumer()
	if g.noChecksAdded() {
		g.cancel()
		return Result{Membership: MembershipUnknown}
	}
	if g.allCheckFinished() {
		// Cancel the consumer to catch the case where all checks were already
		// done by the time Result() is called.
		g.cancel()
	}

	return <-g.resultCh
}
