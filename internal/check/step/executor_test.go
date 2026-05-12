// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package step

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/check"
)

func TestBaseExecutorRunSequential(t *testing.T) {
	t.Parallel()

	t.Run("RunUnion", func(t *testing.T) {
		t.Parallel()

		t.Run("NotMember when there is no step", func(t *testing.T) {
			t.Parallel()
			result := NewExecutor(nil).RunUnion(t.Context())
			assert.Equal(t, check.ResultNotMember, result)
		})

		t.Run("return error when called after context is canceled", func(t *testing.T) {
			t.Parallel()
			ctx, cancel := context.WithCancel(t.Context())
			cancel()

			calls := make([]int, 0, 1)
			steps := plannedSteps([]check.Result{check.ResultIsMember}, &calls)
			result := NewExecutor(nil).RunUnion(ctx, steps...)

			// context is cancelled, therefore steps did not run, and the result Unknown
			assert.Equal(t, check.MembershipUnknown, result.Membership)
			assert.ErrorIs(t, result.Err, context.Canceled)
		})

		t.Run("table", func(t *testing.T) {
			t.Parallel()

			sentinel := errors.New("planted error")

			type testCase struct {
				name           string
				results        []check.Result
				expectedResult check.Result
				expectedErr    error
				expectedCalls  []int
			}

			for _, tc := range []testCase{
				{
					name:    "runs all steps when none is member",
					results: []check.Result{check.ResultUnknown, check.ResultNotMember, check.ResultUnknown, check.ResultNotMember, check.ResultNotMember},
					// todo: ideally, this should be Unknown
					expectedResult: check.ResultNotMember,
					expectedCalls:  []int{0, 1, 2, 3, 4},
				},
				{
					name:           "stops on first Member",
					results:        []check.Result{check.ResultNotMember, check.ResultIsMember, check.ResultNotMember, check.ResultNotMember},
					expectedResult: check.ResultIsMember,
					expectedCalls:  []int{0, 1},
				},
				{
					name:           "stops on first Member; last step is Member",
					results:        []check.Result{check.ResultNotMember, check.ResultNotMember, check.ResultNotMember, check.ResultIsMember},
					expectedResult: check.ResultIsMember,
					expectedCalls:  []int{0, 1, 2, 3},
				},
				{
					name:           "stops on first Member; therefore doesn't run error after Member",
					results:        []check.Result{check.ResultNotMember, check.ResultIsMember, {Err: sentinel}, check.ResultNotMember},
					expectedResult: check.ResultIsMember,
					expectedCalls:  []int{0, 1},
				},
				{
					name:           "error short-circuits",
					results:        []check.Result{check.ResultNotMember, {Err: sentinel}, check.ResultNotMember, check.ResultNotMember},
					expectedResult: check.ResultUnknown,
					expectedErr:    sentinel,
					expectedCalls:  []int{0, 1},
				},
			} {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					calls := make([]int, 0, len(tc.results))
					steps := plannedSteps(tc.results, &calls)

					result := NewExecutor(nil).RunUnion(t.Context(), steps...)
					assert.Equal(t, tc.expectedResult.Membership, result.Membership)
					assert.ErrorIs(t, result.Err, tc.expectedErr)
					assert.Equal(t, tc.expectedCalls, calls)
				})
			}
		})
	})

	t.Run("RunIntersection", func(t *testing.T) {
		t.Parallel()

		t.Run("NotMember when there is no step", func(t *testing.T) {
			t.Parallel()
			result := NewExecutor(nil).RunIntersection(t.Context())
			assert.Equal(t, check.ResultNotMember, result)
		})

		t.Run("return error when called after context is canceled", func(t *testing.T) {
			t.Parallel()
			ctx, cancel := context.WithCancel(t.Context())
			cancel()

			calls := make([]int, 0, 1)
			steps := plannedSteps([]check.Result{check.ResultIsMember}, &calls)
			result := NewExecutor(nil).RunIntersection(ctx, steps...)

			// context is cancelled, therefore steps did not run, and the result Unknown
			assert.Equal(t, check.MembershipUnknown, result.Membership)
			assert.ErrorIs(t, result.Err, context.Canceled)
		})

		t.Run("table", func(t *testing.T) {
			t.Parallel()

			sentinel := errors.New("planted error")

			type testCase struct {
				name           string
				results        []check.Result
				expectedResult check.Result
				expectedErr    error
				expectedCalls  []int
			}

			for _, tc := range []testCase{
				{
					name:           "runs all steps when no NotMember",
					results:        []check.Result{check.ResultIsMember, check.ResultIsMember, check.ResultIsMember},
					expectedResult: check.ResultIsMember,
					expectedCalls:  []int{0, 1, 2},
				},
				{
					name:           "stops on first NotMember",
					results:        []check.Result{check.ResultIsMember, check.ResultNotMember, check.ResultIsMember, check.ResultIsMember},
					expectedResult: check.ResultNotMember,
					expectedCalls:  []int{0, 1},
				},
				{
					name:           "stops on first NotMember; last step is NotMember",
					results:        []check.Result{check.ResultIsMember, check.ResultIsMember, check.ResultIsMember, check.ResultNotMember},
					expectedResult: check.ResultNotMember,
					expectedCalls:  []int{0, 1, 2, 3},
				},
				{
					name:           "error short-circuits",
					results:        []check.Result{check.ResultIsMember, {Err: sentinel}, check.ResultIsMember, check.ResultIsMember},
					expectedResult: check.ResultUnknown,
					expectedErr:    sentinel,
					expectedCalls:  []int{0, 1},
				},
				{
					name:           "should stop on first unknown: unknown is last",
					results:        []check.Result{check.ResultIsMember, check.ResultUnknown},
					expectedResult: check.ResultUnknown,
					expectedCalls:  []int{0, 1},
				},
				{
					name:           "should stop on first unknown: unknown is first",
					results:        []check.Result{check.ResultUnknown, check.ResultIsMember},
					expectedResult: check.ResultUnknown,
					expectedCalls:  []int{0},
				},
			} {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					calls := make([]int, 0, len(tc.results))
					steps := plannedSteps(tc.results, &calls)

					result := NewExecutor(nil).RunIntersection(t.Context(), steps...)
					assert.Equal(t, tc.expectedResult.Membership, result.Membership)
					assert.ErrorIs(t, result.Err, tc.expectedErr)
					assert.Equal(t, tc.expectedCalls, calls)
				})
			}
		})
	})
}

func TestGroupSequential(t *testing.T) {
	t.Parallel()

	t.Run("Add after decisive result is no-op", func(t *testing.T) {
		t.Parallel()

		calls := make([]int, 0, 3)
		runner := newStepGroup(t.Context(), 1, unionMode, NewExecutor(nil))

		runner.Add(check.PlannedStep{Step: stepFunc(func(_ context.Context) check.Result {
			calls = append(calls, 0)
			return check.ResultNotMember
		})})
		require.False(t, runner.Done())

		runner.Add(check.PlannedStep{Step: stepFunc(func(_ context.Context) check.Result {
			calls = append(calls, 1)
			return check.ResultIsMember
		})})
		require.True(t, runner.Done())

		// this step is no-op because the runner is already done with a decisive result
		runner.Add(check.PlannedStep{Step: stepFunc(func(_ context.Context) check.Result {
			calls = append(calls, 2)
			return check.ResultIsMember
		})})

		result := runner.Result()
		require.Equal(t, check.IsMember, result.Membership)
		assert.Equal(t, []int{0, 1}, calls)
	})

	t.Run("Add after cancelled context does not execute the step", func(t *testing.T) {
		t.Parallel()

		calls := 0
		ctx, cancel := context.WithCancel(t.Context())
		cancel()

		runner := newStepGroup(ctx, 1, unionMode, NewExecutor(nil))

		require.True(t, runner.Done())
		runner.Add(check.PlannedStep{Step: stepFunc(func(_ context.Context) check.Result {
			calls++
			return check.ResultIsMember
		})})

		assert.Equal(t, 0, calls)
	})

	t.Run("Result after cancelled context returns MembershipUnknown and context.Canceled", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(t.Context())
		cancel()

		runner := newStepGroup(ctx, 1, unionMode, NewExecutor(nil))
		result := runner.Result()
		assert.Equal(t, check.MembershipUnknown, result.Membership)
		assert.ErrorIs(t, result.Err, context.Canceled)
	})

	t.Run("limit lower than one is clamped to sequential", func(t *testing.T) {
		t.Parallel()

		for _, limit := range []int{0, -1} {
			t.Run("limit", func(t *testing.T) {
				t.Parallel()

				calls := make([]int, 0, 3)
				runner := newStepGroup(t.Context(), limit, unionMode, NewExecutor(nil))

				runner.Add(check.PlannedStep{Step: stepFunc(func(_ context.Context) check.Result {
					calls = append(calls, 0)
					return check.ResultNotMember
				})})
				runner.Add(check.PlannedStep{Step: stepFunc(func(_ context.Context) check.Result {
					calls = append(calls, 1)
					return check.ResultIsMember
				})})
				runner.Add(check.PlannedStep{Step: stepFunc(func(_ context.Context) check.Result {
					calls = append(calls, 2)
					return check.ResultNotMember
				})})

				result := runner.Result()
				require.Equal(t, check.IsMember, result.Membership)
				assert.Equal(t, []int{0, 1}, calls)
			})
		}
	})

	t.Run("Result on empty group uses mode default", func(t *testing.T) {
		t.Parallel()

		unionRunner := newStepGroup(t.Context(), 1, unionMode, NewExecutor(nil))
		assert.Equal(t, check.ResultNotMember, unionRunner.Result())

		intersectionRunner := newStepGroup(t.Context(), 1, intersectionMode, NewExecutor(nil))
		assert.Equal(t, check.ResultIsMember, intersectionRunner.Result())
	})

	t.Run("Result called twice returns the same value", func(t *testing.T) {
		t.Parallel()

		t.Run("decisive result", func(t *testing.T) {
			t.Parallel()

			runner := newStepGroup(t.Context(), 1, unionMode, NewExecutor(nil))
			runner.Add(check.PlannedStep{Step: stepFunc(func(_ context.Context) check.Result {
				return check.ResultIsMember
			})})
			assert.Equal(t, check.IsMember, runner.Result().Membership)
			assert.Equal(t, check.IsMember, runner.Result().Membership)
		})

		t.Run("no decisive result", func(t *testing.T) {
			t.Parallel()

			// When g.set is false after all steps complete, Result() commits the final value.
			// A second call must not return context.Canceled due to the internal g.cancel() from the first call.
			runner := newStepGroup(t.Context(), 1, intersectionMode, NewExecutor(nil))
			runner.Add(check.PlannedStep{Step: stepFunc(func(_ context.Context) check.Result {
				return check.ResultIsMember
			})})
			assert.Equal(t, check.IsMember, runner.Result().Membership)
			assert.Equal(t, check.IsMember, runner.Result().Membership)
		})

		t.Run("cancelled context", func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithCancel(t.Context())
			cancel()

			runner := newStepGroup(ctx, 1, unionMode, NewExecutor(nil))
			assert.ErrorIs(t, runner.Result().Err, context.Canceled)
			assert.ErrorIs(t, runner.Result().Err, context.Canceled)
		})
	})
}

func TestGroupConcurrent(t *testing.T) {
	t.Parallel()

	t.Run("union", func(t *testing.T) {
		t.Parallel()

		t.Run("siblings gets context cancellation when one sibling returns a decisive result", func(t *testing.T) {
			t.Parallel()

			var ran [3]atomic.Bool
			var blockedStepCancelled atomic.Bool

			runner := newStepGroup(t.Context(), 2, unionMode, NewExecutor(nil))

			runner.Add(check.PlannedStep{Step: stepFunc(func(ctx context.Context) check.Result {
				ran[0].Store(true)
				<-ctx.Done()
				blockedStepCancelled.Store(true)
				return check.ResultNotMember
			})})
			runner.Add(check.PlannedStep{Step: stepFunc(func(_ context.Context) check.Result {
				ran[1].Store(true)
				return check.ResultIsMember
			})})
			runner.Add(check.PlannedStep{Step: stepFunc(func(_ context.Context) check.Result {
				ran[2].Store(true)
				return check.ResultNotMember
			})})

			result := runner.Result()

			require.Equal(t, check.IsMember, result.Membership)
			assert.True(t, ran[0].Load())
			assert.True(t, ran[1].Load())
			assert.True(t, blockedStepCancelled.Load())
			assert.False(t, ran[2].Load())
		})

		t.Run("internal sibling cancellation does not leak as error when decisive result exists", func(t *testing.T) {
			t.Parallel()

			runner := newStepGroup(t.Context(), 2, unionMode, NewExecutor(nil))

			// Step 1 propagates its context error on cancellation.
			runner.Add(check.PlannedStep{Step: stepFunc(func(ctx context.Context) check.Result {
				<-ctx.Done()
				return check.Result{Err: ctx.Err()}
			})})
			// Step 2 returns a decisive result, which internally cancels step 1.
			runner.Add(check.PlannedStep{Step: stepFunc(func(_ context.Context) check.Result {
				return check.ResultIsMember
			})})

			result := runner.Result()
			assert.Equal(t, check.IsMember, result.Membership)
			assert.NoError(t, result.Err)
		})

		t.Run("Result blocks until all in-flight goroutines finish", func(t *testing.T) {
			t.Parallel()

			var stepCompleted atomic.Bool

			runner := newStepGroup(t.Context(), 2, unionMode, NewExecutor(nil))

			runner.Add(check.PlannedStep{Step: stepFunc(func(ctx context.Context) check.Result {
				<-ctx.Done()
				stepCompleted.Store(true)
				return check.ResultNotMember
			})})
			runner.Add(check.PlannedStep{Step: stepFunc(func(_ context.Context) check.Result {
				return check.ResultIsMember
			})})

			result := runner.Result()
			require.Equal(t, check.IsMember, result.Membership)
			assert.True(t, stepCompleted.Load(), "Result() returned before the in-flight goroutine finished")
		})

		t.Run("Add after cancelled context does not execute the step", func(t *testing.T) {
			t.Parallel()

			var calls atomic.Int64
			ctx, cancel := context.WithCancel(t.Context())
			cancel()

			runner := newStepGroup(ctx, 2, unionMode, NewExecutor(nil))

			require.True(t, runner.Done())
			runner.Add(check.PlannedStep{Step: stepFunc(func(_ context.Context) check.Result {
				calls.Add(1)
				return check.ResultIsMember
			})})

			assert.Equal(t, int64(0), calls.Load())
		})

		t.Run("Result after cancelled context returns MembershipUnknown and context.Canceled", func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithCancel(t.Context())
			cancel()

			runner := newStepGroup(ctx, 2, unionMode, NewExecutor(nil))
			result := runner.Result()
			assert.Equal(t, check.MembershipUnknown, result.Membership)
			assert.ErrorIs(t, result.Err, context.Canceled)
		})

		t.Run("context cancelled mid-step: error propagates to Result", func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithCancel(t.Context())

			runner := newStepGroup(ctx, 2, unionMode, NewExecutor(nil))

			// Step 1 blocks until the context is cancelled, then returns the cancellation error.
			runner.Add(check.PlannedStep{Step: stepFunc(func(ctx context.Context) check.Result {
				<-ctx.Done()
				return check.Result{Err: ctx.Err()}
			})})
			cancel()

			// Step 2 is a no-op: g.ctx is already cancelled so Add returns immediately without running it.
			runner.Add(check.PlannedStep{Step: stepFunc(func(ctx context.Context) check.Result {
				return check.ResultNotMember
			})})

			result := runner.Result()
			assert.ErrorIs(t, result.Err, context.Canceled)
			assert.Equal(t, check.MembershipUnknown, result.Membership)
		})

		t.Run("returns NotMember when no step is decisive", func(t *testing.T) {
			t.Parallel()

			runner := newStepGroup(t.Context(), 3, unionMode, NewExecutor(nil))
			for range 3 {
				runner.Add(check.PlannedStep{Step: stepFunc(func(_ context.Context) check.Result {
					return check.ResultNotMember
				})})
			}

			result := runner.Result()
			assert.Equal(t, check.NotMember, result.Membership)
			assert.NoError(t, result.Err)
		})
	})

	t.Run("intersection", func(t *testing.T) {
		t.Parallel()

		t.Run("all steps IsMember returns IsMember without error", func(t *testing.T) {
			t.Parallel()

			// For intersection, IsMember is not decisive so g.set stays false throughout.
			// Result() relies on ctxErr being nil (no internal g.cancel() fires) to reach NoDecisionResult().
			// This test documents that the success path produces no spurious error.
			runner := newStepGroup(t.Context(), 3, intersectionMode, NewExecutor(nil))
			for range 3 {
				runner.Add(check.PlannedStep{Step: stepFunc(func(_ context.Context) check.Result {
					return check.ResultIsMember
				})})
			}

			result := runner.Result()
			assert.Equal(t, check.IsMember, result.Membership)
			assert.NoError(t, result.Err)
		})
	})
}

// plannedSteps is a test helper that creates a slice of PlannedStep from a slice of Results.
// Each step records its index in calls when executed, allowing tests to verify execution order and count.
func plannedSteps(results []check.Result, calls *[]int) []check.PlannedStep {
	var mu sync.Mutex
	steps := make([]check.PlannedStep, 0, len(results))
	for i, result := range results {
		steps = append(steps, check.PlannedStep{Step: stepFunc(func(_ context.Context) check.Result {
			mu.Lock()
			*calls = append(*calls, i)
			mu.Unlock()
			return result
		})})
	}
	return steps
}

// stepFunc is a Step whose Execute calls f directly, ignoring req and executor.
type stepFunc func(ctx context.Context) check.Result

func (stepFunc) Kind() check.StepKind { return check.StepIsAllowed }
func (f stepFunc) Execute(ctx context.Context, _ check.CheckRequest, _ check.Executor) check.Result {
	return f(ctx)
}
