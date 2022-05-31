package checkgroup

import "context"

type (
	Checkgroup interface {
		Done() bool
		Add(check Func)
		SetIsMember()
		Result() Result
		CheckFunc() Func
	}

	Factory func(ctx context.Context) Checkgroup

	Func   func(ctx context.Context, resultCh chan<- Result)
	Result struct {
		Membership Membership
		Err        error
	}

	Membership int
)

const (
	MembershipUnknown Membership = iota
	IsMember
	NotMember
)

var (
	ResultIsMember  = Result{Membership: IsMember}
	ResultNotMember = Result{Membership: NotMember}
)

var DefaultFactory Factory = NewSequential

func New(ctx context.Context) Checkgroup {
	return DefaultFactory(ctx)
}

func ErrorFunc(err error) Func {
	return func(_ context.Context, resultCh chan<- Result) {
		resultCh <- Result{Err: err}
	}
}

var IsMemberFunc Func = func(_ context.Context, resultCh chan<- Result) {
	resultCh <- Result{Membership: IsMember}
}

var NotMemberFunc Func = func(_ context.Context, resultCh chan<- Result) {
	resultCh <- Result{Membership: NotMember}
}

var UnknownMemberFunc Func = func(_ context.Context, resultCh chan<- Result) {
	resultCh <- Result{Membership: MembershipUnknown}
}
