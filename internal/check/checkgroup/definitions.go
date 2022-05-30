package checkgroup

import "context"

type Func func(ctx context.Context, resultCh chan<- Result)

type Result struct {
	Membership Membership
	Err        error
}

type Membership int

const (
	MembershipUnknown Membership = iota
	IsMember
	NotMember
)

var (
	ResultIsMember  = Result{Membership: IsMember}
	ResultNotMember = Result{Membership: NotMember}
)

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
