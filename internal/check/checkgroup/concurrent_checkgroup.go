package checkgroup

import (
	"context"
	"sync"
)

// A concurrentCheckgroup is a collection of goroutines performing checks.
type concurrentCheckgroup struct {
	ctx               context.Context
	subcheckCtx       context.Context
	cancel            context.CancelFunc
	startConsumerOnce sync.Once
	resultCh          chan Result
	addCheckCh        chan Func
	freezeCh          chan struct{}
	setResultOnce     sync.Once
	result            Result
}

func NewConcurrent(ctx context.Context) Checkgroup {
	g := &concurrentCheckgroup{
		ctx:        ctx,
		freezeCh:   make(chan struct{}),
		resultCh:   make(chan Result, 1),
		addCheckCh: make(chan Func),
	}
	g.subcheckCtx, g.cancel = context.WithCancel(g.ctx)
	g.startConsumer()
	return g
}

func receiveRemaining(ch <-chan Result, remaining int) {
	for i := 0; i < remaining; i++ {
		<-ch
	}
}

func (g *concurrentCheckgroup) startConsumer() {
	g.startConsumerOnce.Do(func() {
		go func() {
			var (
				subcheckCh     = make(chan Result, 1)
				totalChecks    = 0
				finishedChecks = 0
				frozen         = false
				sendResultOnce sync.Once
			)

			defer g.cancel()
			defer close(g.resultCh)

			// We don't care about the subcheck results (most will be
			// `context.Canceled`), but we still want to receive these results
			// so that there are no dangling goroutines.
			defer receiveRemaining(subcheckCh, totalChecks-finishedChecks)

			for {
				select {
				case f := <-g.addCheckCh:
					if frozen {
						continue
					}
					totalChecks++
					go f(g.subcheckCtx, subcheckCh)

				case <-g.freezeCh:
					frozen = true
					if finishedChecks == totalChecks {
						sendResultOnce.Do(func() { g.resultCh <- ResultNotMember })
						return
					}

				case result := <-subcheckCh:
					finishedChecks++
					if result.Err != nil || result.Membership == IsMember {
						sendResultOnce.Do(func() { g.resultCh <- result })
						return
					}

					if frozen && finishedChecks == totalChecks {
						sendResultOnce.Do(func() { g.resultCh <- ResultNotMember })
						return
					}

				case <-g.subcheckCtx.Done():
					sendResultOnce.Do(func() { g.resultCh <- Result{Err: g.ctx.Err()} })
					return
				}
			}
		}()
	})
}

func (g *concurrentCheckgroup) Done() bool {
	select {
	case <-g.subcheckCtx.Done():
		return true
	default:
		return false
	}
}

// Add adds the Func to the checkgroup and starts running it.
func (g *concurrentCheckgroup) Add(check Func) {
	select {
	case g.addCheckCh <- check:
	case <-g.subcheckCtx.Done():
	}
}

// SetIsMember makes the checkgroup emit "IsMember" directly.
func (g *concurrentCheckgroup) SetIsMember() {
	g.Add(IsMemberFunc)
}

// Result returns the Result, possibly blocking.
func (g *concurrentCheckgroup) Result() Result {
	g.setResultOnce.Do(func() {
		select {
		case g.freezeCh <- struct{}{}:
		case <-g.subcheckCtx.Done():
		}
		g.result = <-g.resultCh
	})

	return g.result
}

// CheckFunc returns a `Func` that writes the result to the result channel.
func (g *concurrentCheckgroup) CheckFunc() Func {
	return func(ctx context.Context, resultCh chan<- Result) {
		select {
		case g.freezeCh <- struct{}{}:
		case <-g.subcheckCtx.Done():
		}

		select {
		case result := <-g.resultCh:
			resultCh <- result
		case <-ctx.Done():
			g.cancel()
			resultCh <- <-g.resultCh
		}
	}
}
