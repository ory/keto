package checkgroup

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/ketoapi"
)

type (
	Checkgroup interface {
		Done() bool
		Add(check CheckFunc)
		SetIsMember()
		Result() Result
		CheckFunc() CheckFunc
	}

	Factory = func(ctx context.Context) Checkgroup

	CheckFunc = func(ctx context.Context, resultCh chan<- Result)

	Result struct {
		Membership Membership
		Tree       *ketoapi.Tree[*relationtuple.RelationTuple]
		Err        error
	}

	Edge struct {
		Tuple relationtuple.RelationTuple
		Type  ketoapi.TreeNodeType
	}

	Transformation int

	Membership int

	tree = ketoapi.Tree[*relationtuple.RelationTuple]
)

//go:generate stringer -type Membership
const (
	MembershipUnknown Membership = iota
	IsMember
	NotMember
)

var (
	ResultIsMember  = Result{Membership: IsMember}
	ResultNotMember = Result{Membership: NotMember}
)

var DefaultFactory = NewConcurrent

func New(ctx context.Context) Checkgroup {
	return DefaultFactory(ctx)
}

func ErrorFunc(err error) CheckFunc {
	return func(_ context.Context, resultCh chan<- Result) {
		resultCh <- Result{Err: errors.WithStack(err)}
	}
}

func IsMemberFunc(_ context.Context, resultCh chan<- Result) {
	resultCh <- Result{Membership: IsMember}
}

func NotMemberFunc(_ context.Context, resultCh chan<- Result) {
	resultCh <- Result{Membership: NotMember}
}

func UnknownMemberFunc(_ context.Context, resultCh chan<- Result) {
	resultCh <- Result{Membership: MembershipUnknown}
}

// WithEdge adds the edge e to the result of the function.
func WithEdge(e Edge, f CheckFunc) CheckFunc {
	return func(ctx context.Context, resultCh chan<- Result) {
		childCh := make(chan Result, 1)
		go f(ctx, childCh)
		select {
		case result := <-childCh:
			if result.Tree == nil {
				result.Tree = &tree{
					Type:  ketoapi.TreeNodeLeaf,
					Tuple: &e.Tuple,
				}
			} else {
				result.Tree = &tree{
					Type:     e.Type,
					Tuple:    &e.Tuple,
					Children: []*tree{result.Tree},
				}
			}
			resultCh <- result
		case <-ctx.Done():
			resultCh <- Result{Err: ctx.Err()}
		}
	}
}
