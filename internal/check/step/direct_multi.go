// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package step

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ory/keto/internal/check"
)

// DirectMultiStep checks whether any of the given relations exist directly in
// the database for the tuple.
type DirectMultiStep struct {
	Relations []string
}

func (DirectMultiStep) Kind() check.StepKind { return check.StepDirectMulti }

func (s DirectMultiStep) Execute(ctx context.Context, req check.CheckRequest, ex check.Executor) check.Result {
	if req.RestDepth <= 0 {
		ex.Deps().Logger().
			WithField("method", "DirectMultiStep").
			Debug("reached max-depth, therefore this query will not be further expanded")
		return check.Result{Membership: check.MembershipUnknown}
	}

	ex.Deps().Logger().
		WithField("request", req.Tuple.String()).
		WithField("relations", s.Relations).
		Trace("check direct multi")

	res, err := ex.Deps().Traverser().FindTupleWithRelations(ctx, req.Tuple, s.Relations)
	if err != nil {
		return check.Result{Err: errors.WithStack(err)}
	}
	if res != nil {
		return check.ResultIsMember
	}
	return check.ResultNotMember
}
