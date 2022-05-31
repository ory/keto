package checkgroup

import (
	"context"
	"fmt"
	"strings"

	"github.com/ory/keto/internal/relationtuple"
)

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
		Path       Path
		Err        error
	}
	Path struct {
		Edges []Edge
	}

	Edge struct {
		Tuple          relationtuple.InternalRelationTuple
		Transformation Transformation
	}

	Transformation int

	Membership int
)

const (
	MembershipUnknown Membership = iota
	IsMember
	NotMember
)

const (
	TransformationUnknown Transformation = iota
	TransformationDirect
	TransformationTupleToUserset
	TransformationComputedUserset
)

func TransformationFromString(s string) Transformation {
	switch s {
	case "direct":
		return TransformationDirect
	case "tuple-to-userset":
		return TransformationTupleToUserset
	case "computed-userset":
		return TransformationComputedUserset
	default:
		return TransformationUnknown
	}
}

func (t Transformation) String() string {
	switch t {
	case TransformationDirect:
		return "direct"
	case TransformationTupleToUserset:
		return "tuple-to-userset"
	case TransformationComputedUserset:
		return "computed-userset"
	default:
		return "unknown"
	}
}

func (p *Path) String() string {
	parts := []string{}
	for _, edge := range p.Edges {
		parts = append(parts, fmt.Sprintf("%s as %s", edge.Tuple.String(), edge.Transformation))
	}
	return strings.Join(parts, " -> ")
}

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

// WithEdge adds the edge e to the result of the function.
func WithEdge(e Edge, f Func) Func {
	return func(ctx context.Context, resultCh chan<- Result) {
		childCh := make(chan Result)
		go f(ctx, childCh)
		select {
		case result := <-childCh:
			result.Path.Edges = append([]Edge{e}, result.Path.Edges...)
			resultCh <- result
		case <-ctx.Done():
			resultCh <- Result{Err: ctx.Err()}
		}
	}
}
