// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package checkgroup

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/ketoapi"
)

type (
	Checkgroup interface {
		// Done returns true if a result is available.
		Done() bool

		// Add adds the CheckFunc to the checkgroup and starts running it.
		Add(check CheckFunc)

		// SetIsMember makes the checkgroup emit "IsMember" directly.
		SetIsMember()

		// Result returns the result, possibly blocking.
		Result() Result

		// CheckFunc returns a CheckFunc that writes the result to the result
		// channel.
		CheckFunc() CheckFunc
	}

	Pool interface {
		// Add adds the function to the pool and schedules it. The function will
		// only be run if there is a free worker available in the pool, thus
		// limiting the concurrent workloads in flight.
		Add(check func())

		// TryAdd tries to add the check function if the pool has capacity.
		// Otherwise, it returns false and does not add the check.
		TryAdd(check func()) bool
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
