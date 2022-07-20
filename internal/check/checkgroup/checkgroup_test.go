package checkgroup_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"

	"github.com/ory/keto/internal/check/checkgroup"
)

var neverFinishesCheckFunc = func(ctx context.Context, resultCh chan<- checkgroup.Result) {
	<-ctx.Done()
	resultCh <- checkgroup.Result{Err: ctx.Err()}
}

func isMemberAfterDelayFunc(delay time.Duration) checkgroup.CheckFunc {
	return func(ctx context.Context, resultCh chan<- checkgroup.Result) {
		select {
		case <-time.After(delay):
			resultCh <- checkgroup.ResultIsMember
		case <-ctx.Done():
			resultCh <- checkgroup.Result{Err: ctx.Err()}
		}
	}
}

func notMemberAfterDelayFunc(delay time.Duration) checkgroup.CheckFunc {
	return func(ctx context.Context, resultCh chan<- checkgroup.Result) {
		select {
		case <-time.After(delay):
			resultCh <- checkgroup.ResultNotMember
		case <-ctx.Done():
			resultCh <- checkgroup.Result{Err: ctx.Err()}
		}
	}
}

func TestCheckgroup_cancels(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	g := checkgroup.New(ctx)
	g.Add(neverFinishesCheckFunc)
	g.Add(neverFinishesCheckFunc)
	g.Add(neverFinishesCheckFunc)
	g.Add(neverFinishesCheckFunc)
	g.Add(neverFinishesCheckFunc)
	cancel()
	assert.Equal(t, checkgroup.Result{Err: context.Canceled}, g.Result())
}

func TestCheckgroup_reports_first_result(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g := checkgroup.New(ctx)
	g.Add(neverFinishesCheckFunc)
	g.Add(checkgroup.IsMemberFunc)
	assert.Equal(t, checkgroup.Result{Membership: checkgroup.IsMember}, g.Result())
}

func TestCheckgroup_cancels_all_other_subchecks(t *testing.T) {
	t.Parallel()

	wasCancelled := make(chan bool)
	var mockCheckFn = func(ctx context.Context, resultCh chan<- checkgroup.Result) {
		<-ctx.Done()
		wasCancelled <- true
		resultCh <- checkgroup.Result{Err: ctx.Err()}
	}

	ctx := context.Background()

	g := checkgroup.New(ctx)
	g.Add(mockCheckFn)
	g.Add(neverFinishesCheckFunc)
	g.Add(checkgroup.IsMemberFunc)
	result := g.Result()

	assert.Equal(t, checkgroup.ResultIsMember, result)
	assert.True(t, <-wasCancelled)
	assert.True(t, g.Done())
}

func TestCheckgroup_returns_first_successful_is_member(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	g := checkgroup.New(ctx)
	g.Add(checkgroup.NotMemberFunc)
	g.Add(checkgroup.NotMemberFunc)
	time.Sleep(1 * time.Millisecond)

	assert.False(t, g.Done())

	g.Add(func(_ context.Context, resultCh chan<- checkgroup.Result) {
		resultCh <- checkgroup.ResultIsMember
	})

	resultCh := make(chan checkgroup.Result)
	go g.CheckFunc()(ctx, resultCh)

	assert.Equal(t, checkgroup.ResultIsMember, g.Result())
	assert.Equal(t, checkgroup.ResultIsMember, g.Result())
	assert.Equal(t, checkgroup.ResultIsMember, g.Result())
	assert.Equal(t, checkgroup.ResultIsMember, <-resultCh)
	assert.True(t, g.Done())
}

func TestCheckgroup_returns_immediately_if_nothing_to_check(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g := checkgroup.New(ctx)
	assert.Equal(t, checkgroup.ResultNotMember, g.Result())
}

func TestCheckgroup_has_no_leaks(t *testing.T) {
	testCases := []struct {
		name     string
		checks   []checkgroup.CheckFunc
		expected checkgroup.Result
	}{
		{
			name: "is member after delay",
			checks: []checkgroup.CheckFunc{
				checkgroup.UnknownMemberFunc,
				isMemberAfterDelayFunc(5 * time.Millisecond),
				notMemberAfterDelayFunc(1 * time.Millisecond),
				neverFinishesCheckFunc,
				neverFinishesCheckFunc,
				neverFinishesCheckFunc,
			},
			expected: checkgroup.ResultIsMember,
		},
		{
			name: "is member immediately",
			checks: []checkgroup.CheckFunc{
				checkgroup.IsMemberFunc,
				checkgroup.IsMemberFunc,
				checkgroup.IsMemberFunc,
				checkgroup.UnknownMemberFunc,
				isMemberAfterDelayFunc(5 * time.Millisecond),
				notMemberAfterDelayFunc(1 * time.Millisecond),
				neverFinishesCheckFunc,
				neverFinishesCheckFunc,
				neverFinishesCheckFunc,
			},
			expected: checkgroup.ResultIsMember,
		},
		{
			name: "is not member immediately",
			checks: []checkgroup.CheckFunc{
				checkgroup.NotMemberFunc,
				checkgroup.NotMemberFunc,
				checkgroup.NotMemberFunc,
				checkgroup.UnknownMemberFunc,
			},
			expected: checkgroup.ResultNotMember,
		},
		{
			name: "is not member after delay",
			checks: []checkgroup.CheckFunc{
				checkgroup.NotMemberFunc,
				checkgroup.NotMemberFunc,
				checkgroup.NotMemberFunc,
				checkgroup.UnknownMemberFunc,
				notMemberAfterDelayFunc(5 * time.Millisecond),
				notMemberAfterDelayFunc(1 * time.Millisecond),
			},
			expected: checkgroup.ResultNotMember,
		},
		{
			name: "never finishes",
			checks: []checkgroup.CheckFunc{
				neverFinishesCheckFunc,
				neverFinishesCheckFunc,
				checkgroup.UnknownMemberFunc,
				checkgroup.UnknownMemberFunc,
			},
			expected: checkgroup.Result{Err: context.DeadlineExceeded},
		},
	}

	for _, tc := range testCases {
		t.Run("tc="+tc.name, func(t *testing.T) {
			defer goleak.VerifyNone(t, goleak.IgnoreCurrent())

			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
			defer cancel()

			g := checkgroup.New(ctx)
			for _, check := range tc.checks {
				g.Add(check)
			}

			resultCh := make(chan checkgroup.Result)
			go g.CheckFunc()(ctx, resultCh)
			result := <-resultCh

			assert.Equal(t, tc.expected, result)
		})
	}
}
