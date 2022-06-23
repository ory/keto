package checkgroup_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/internal/check/checkgroup"
)

var neverFinishesCheckFn checkgroup.Func = func(context.Context, chan<- checkgroup.Result) {}

func TestCheckgroup_cancels(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	g := checkgroup.New(ctx)
	g.Add(neverFinishesCheckFn)
	cancel()
	assert.Equal(t, checkgroup.Result{Err: context.Canceled}, g.Result())
}

func TestCheckgroup_reports_first_result(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g := checkgroup.New(ctx)
	g.Add(neverFinishesCheckFn)
	g.Add(checkgroup.IsMemberFunc)
	assert.Equal(t, checkgroup.Result{Membership: checkgroup.IsMember}, g.Result())
}

func TestCheckgroup_cancels_all_other_subchecks(t *testing.T) {
	t.Parallel()

	wasCancelled := make(chan bool)
	var mockCheckFn checkgroup.Func = func(ctx context.Context, _ chan<- checkgroup.Result) {
		<-ctx.Done()
		wasCancelled <- true
	}

	ctx := context.Background()

	g := checkgroup.New(ctx)
	g.Add(neverFinishesCheckFn)
	g.Add(checkgroup.IsMemberFunc)
	g.Add(mockCheckFn)
	g.Result()
	assert.True(t, <-wasCancelled)
	assert.NotNil(t, <-g.Ctx.Done())
	assert.True(t, g.Done())
}

func TestCheckgroup_returns_first_successful_is_member(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	g := checkgroup.New(ctx)
	g.Add(neverFinishesCheckFn)
	g.Add(checkgroup.NotMemberFunc)
	g.Add(checkgroup.NotMemberFunc)
	time.Sleep(1 * time.Millisecond)
	assert.False(t, g.Done())
	g.Add(checkgroup.IsMemberFunc)

	assert.Equal(t, checkgroup.Result{Membership: checkgroup.IsMember}, g.Result())
	assert.NotNil(t, <-g.Ctx.Done())
	assert.True(t, g.Done())
}

func TestCheckgroup_returns_immediately_if_nothing_to_check(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g := checkgroup.New(ctx)
	assert.Equal(t,
		checkgroup.Result{Membership: checkgroup.MembershipUnknown},
		g.Result())
}
