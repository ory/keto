package checkgroup_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/internal/check/checkgroup"
)

var neverFinishesCheckFn checkgroup.Func = func(ctx context.Context, resultCh chan<- checkgroup.Result) {
	<-ctx.Done()
	resultCh <- checkgroup.Result{Err: ctx.Err()}
}

var checkgroups = []struct {
	name string
	new  checkgroup.Factory
}{
	{name: "sequential", new: checkgroup.NewSequential},
	{name: "concurrent", new: checkgroup.NewConcurrent},
}

func runWithCheckgroup(t *testing.T, test func(t *testing.T, new checkgroup.Factory)) {
	for _, group := range checkgroups {
		group := group
		t.Run(group.name, func(t *testing.T) {
			t.Parallel()
			test(t, group.new)
		})
	}
}

func TestCheckgroup_cancels(t *testing.T) {
	t.Parallel()

	runWithCheckgroup(t, func(t *testing.T, new checkgroup.Factory) {
		ctx, cancel := context.WithCancel(context.Background())
		g := new(ctx)
		g.Add(neverFinishesCheckFn)
		g.Add(neverFinishesCheckFn)
		g.Add(neverFinishesCheckFn)
		g.Add(neverFinishesCheckFn)
		g.Add(neverFinishesCheckFn)
		cancel()
		assert.Equal(t, checkgroup.Result{Err: context.Canceled}, g.Result())
	})
}

func TestCheckgroup_reports_first_result(t *testing.T) {
	t.Parallel()

	runWithCheckgroup(t, func(t *testing.T, new checkgroup.Factory) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		g := new(ctx)
		g.Add(neverFinishesCheckFn)
		g.Add(checkgroup.IsMemberFunc)
		assert.Equal(t, checkgroup.Result{Membership: checkgroup.IsMember}, g.Result())
	})
}

func TestCheckgroup_cancels_all_other_subchecks(t *testing.T) {
	t.Parallel()

	wasCancelled := make(chan bool)
	var mockCheckFn checkgroup.Func = func(ctx context.Context, _ chan<- checkgroup.Result) {
		<-ctx.Done()
		wasCancelled <- true
	}

	ctx := context.Background()

	g := checkgroup.NewConcurrent(ctx)
	g.Add(mockCheckFn)
	g.Add(neverFinishesCheckFn)
	g.Add(checkgroup.IsMemberFunc)
	result := g.Result()

	assert.Equal(t, checkgroup.ResultIsMember, result)
	assert.True(t, <-wasCancelled)
	assert.True(t, g.Done())
}

func TestCheckgroup_returns_first_successful_is_member(t *testing.T) {
	t.Parallel()

	runWithCheckgroup(t, func(t *testing.T, new checkgroup.Factory) {
		ctx := context.Background()

		g := new(ctx)
		g.Add(neverFinishesCheckFn)
		g.Add(checkgroup.NotMemberFunc)
		g.Add(checkgroup.NotMemberFunc)
		time.Sleep(1 * time.Millisecond)
		assert.False(t, g.Done())
		g.Add(func(_ context.Context, resultCh chan<- checkgroup.Result) {
			time.Sleep(10 * time.Millisecond)
			resultCh <- checkgroup.ResultIsMember
		})

		assert.Equal(t, checkgroup.Result{Membership: checkgroup.IsMember}, g.Result())
		// assert.Equal(t, checkgroup.Result{Membership: checkgroup.IsMember}, g.Result())
		assert.True(t, g.Done())
	})
}

func TestCheckgroup_returns_immediately_if_nothing_to_check(t *testing.T) {
	t.Parallel()
	runWithCheckgroup(t, func(t *testing.T, new checkgroup.Factory) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		g := new(ctx)
		assert.Equal(t, checkgroup.ResultNotMember, g.Result())
	})
}

func TestCheckgroup_propagates_not_member_results(t *testing.T) {
	t.Parallel()
	runWithCheckgroup(t, func(t *testing.T, new checkgroup.Factory) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		g := new(ctx)
		for i := 0; i < 10; i++ {
			i := i
			g.Add(func(ctx context.Context, resultCh chan<- checkgroup.Result) {
				select {
				case <-time.After(time.Duration(i) * time.Millisecond):
					resultCh <- checkgroup.ResultNotMember
				case <-ctx.Done():
					resultCh <- checkgroup.Result{Err: context.Canceled}
				}
			})
		}

		resultCh := make(chan checkgroup.Result)
		go g.CheckFunc()(ctx, resultCh)
		result := <-resultCh

		assert.Equal(t, checkgroup.ResultNotMember, result)
	})
}
