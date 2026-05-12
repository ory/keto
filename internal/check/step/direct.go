// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package step

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ory/keto/internal/check"
)

// DirectStep checks whether the exact relation tuple exists in the database.
type DirectStep struct{}

func (DirectStep) Kind() check.StepKind { return check.StepDirect }

func (DirectStep) Execute(ctx context.Context, req check.CheckRequest, ex check.Executor) check.Result {
	if req.RestDepth <= 0 {
		ex.Deps().Logger().
			WithField("request", req.Tuple.String()).
			WithField("method", "DirectStep").
			Debug("reached max-depth, therefore this query will not be further expanded")
		return check.Result{Membership: check.MembershipUnknown}
	}

	ex.Deps().Logger().
		WithField("request", req.Tuple.String()).
		Trace("check direct")

	found, err := ex.Deps().RelationTupleManager().ExistsRelationTuples(ctx, req.Tuple.ToQuery())
	switch {
	case err != nil:
		ex.Deps().Logger().
			WithField("method", "DirectStep").
			WithError(err).
			Error("failed to look up direct access in db")
		return check.Result{Err: errors.WithStack(err)}
	case found:
		return check.ResultIsMember
	default:
		return check.ResultNotMember
	}
}
