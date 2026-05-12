// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package step

import (
	"context"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/namespace/ast"
)

// InvertStep evaluates its child and inverts the membership result.
// IsMember becomes NotMember and vice versa; Unknown and Error pass through.
type InvertStep struct {
	Child ast.Child
}

func (InvertStep) Kind() check.StepKind { return check.StepInvert }

func (s InvertStep) Execute(ctx context.Context, req check.CheckRequest, ex check.Executor) check.Result {
	if req.RestDepth <= 0 {
		ex.Deps().Logger().Debug("reached max-depth, therefore this query will not be further expanded")
		return check.Result{Membership: check.MembershipUnknown}
	}

	ex.Deps().Logger().
		WithField("request", req.Tuple.String()).
		Trace("invert check")

	childStep, err := rewriteChildToStep(s.Child)
	if err != nil {
		return check.Result{Err: err}
	}

	result := ex.Run(ctx, childStep, req)

	switch result.Membership {
	case check.IsMember:
		result.Membership = check.NotMember
	case check.NotMember:
		result.Membership = check.IsMember
	}
	return result
}
