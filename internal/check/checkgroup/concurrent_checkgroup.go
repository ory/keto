// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package checkgroup

import (
	"context"
	"sync"
)

// A concurrentCheckgroup is a collection of goroutines performing checks.
type concurrentCheckgroup struct {
	// ctx is the main context of the checkgroup. If ctx is cancelled, all
	// subchecks are also cancelled and the result is set to Result{Err:
	// ctx.Err()}.
	ctx context.Context

	// subcheckCtx is the context used for the subchecks.
	subcheckCtx context.Context

	// cancel cancels the subcheckCtx, either because a result was found, or
	// because the context of the CheckFunc() was cancelled.
	cancel context.CancelFunc

	// sync.Once to ensure that we only ever start one consumer.
	startConsumerOnce sync.Once

	// addCheckCh is used to add a check to the consumer.
	addCheckCh chan CheckFunc

	// finalizeCh is used to signal that a result was requested.
	finalizeCh chan struct{}

	// doneCh is closed by the consumer if a result is ready. Methods that want
	// to retrieve the result need to wait for the doneCh to be closed first.
	doneCh chan struct{}

	// result is only written once by the consumer, and  can only be read after
	// the doneCh channel is closed.
	result Result

	// reading from reserveCheckCh reserves the right to create a concurrent
	// check.
	reserveCheckCh chan struct{}
}

func NewConcurrent(ctx context.Context) Checkgroup {
	g := &concurrentCheckgroup{
		ctx:            ctx,
		finalizeCh:     make(chan struct{}),
		doneCh:         make(chan struct{}),
		addCheckCh:     make(chan CheckFunc),
		reserveCheckCh: make(chan struct{}, 1),
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
				resultCh       = make(chan Result, 1)
				totalChecks    = 0
				finishedChecks = 0
				finalizing     = false
			)

			defer g.cancel()

			// Closing the doneCh will signal that the result is ready.
			defer close(g.doneCh)

			// We don't care about the subcheck results (most will be
			// `context.Canceled`), but we still want to receive these results
			// so that there are no dangling goroutines.
			defer func() {
				go receiveRemaining(resultCh, totalChecks-finishedChecks)
			}()

			// Start with one reservation available.
			g.reserveCheckCh <- struct{}{}

			for {
				select {
				case check := <-g.addCheckCh:
					if finalizing {
						continue
					}
					totalChecks++
					go check(g.subcheckCtx, resultCh)

				case <-g.finalizeCh:
					if finalizing {
						// we're already finalizing, so we don't want to
						// accidentally set the result to ResultNotMember on a
						// second finalize request
						continue
					}
					finalizing = true
					if finishedChecks == totalChecks {
						g.result = ResultNotMember
						return
					}

				case result := <-resultCh:
					finishedChecks++
					if result.Err != nil || result.Membership == IsMember {
						g.result = result
						return
					}

					if finalizing && finishedChecks == totalChecks {
						g.result = ResultNotMember
						return
					}

					// ready for a new check
					select {
					case g.reserveCheckCh <- struct{}{}:
					default:
					}

				case <-g.subcheckCtx.Done():
					g.result = Result{Err: g.ctx.Err()}
					return
				}
			}
		}()
	})
}

func (g *concurrentCheckgroup) Done() bool {
	select {
	case <-g.doneCh:
		return true
	default:
		return false
	}
}

// Add adds the CheckFunc to the checkgroup and starts running it.
func (g *concurrentCheckgroup) Add(check CheckFunc) {
	select {
	case <-g.reserveCheckCh:
		select {
		case g.addCheckCh <- check:
		case <-g.subcheckCtx.Done():
		}
	case <-g.subcheckCtx.Done():
	}
}

// SetIsMember makes the checkgroup emit "IsMember" directly.
func (g *concurrentCheckgroup) SetIsMember() {
	g.Add(IsMemberFunc)
}

// tryFinalize tries to set the group state to finalize, i.e, signal the
// consumer that the result was requested and that no more checks will be added.
// If the consumer is already done, finalizing is not necessary anymore. This
// should never block.
func (g *concurrentCheckgroup) tryFinalize() {
	select {
	case g.finalizeCh <- struct{}{}:
	case <-g.doneCh:
	}
}

// Result returns the Result, possibly blocking.
func (g *concurrentCheckgroup) Result() Result {
	g.tryFinalize()
	<-g.doneCh
	return g.result
}

// CheckFunc returns a CheckFunc that writes the result to the result channel.
func (g *concurrentCheckgroup) CheckFunc() CheckFunc {
	return func(ctx context.Context, resultCh chan<- Result) {
		g.tryFinalize()

		select {
		case <-g.doneCh:
			resultCh <- g.result
		case <-ctx.Done():
			g.cancel()
			<-g.doneCh
			resultCh <- g.result
		}
	}
}
