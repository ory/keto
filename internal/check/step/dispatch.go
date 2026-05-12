// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package step

import (
	"context"

	"github.com/ory/keto/internal/check"
)

// IsAllowedStep is the central semantic dispatcher. It reads the AST rule for
// the tuple and determines what checks are applicable via RunUnion.
type IsAllowedStep struct {
	skipDirect bool
}

func (IsAllowedStep) Kind() check.StepKind { return check.StepIsAllowed }

func (s IsAllowedStep) Execute(ctx context.Context, req check.CheckRequest, ex check.Executor) check.Result {
	if req.RestDepth <= 0 {
		ex.Deps().Logger().
			WithField("method", "IsAllowedStep").
			Debug("reached max-depth, therefore this query will not be further expanded")
		return check.Result{Membership: check.MembershipUnknown}
	}

	ex.Deps().Logger().
		WithField("request", req.Tuple.String()).
		Trace("check is allowed")

	relation, err := astRelationFor(ctx, ex.Deps(), req.Tuple)
	if err != nil {
		return check.Result{Err: err}
	}

	hasRewrite := relation != nil && relation.SubjectSetRewrite != nil
	strictMode := ex.Deps().Config(ctx).StrictMode()
	canHaveSubjectSets := !strictMode || relation == nil || containsSubjectSetExpand(relation)

	steps := make([]check.PlannedStep, 0, 3)

	if hasRewrite {
		steps = append(steps, check.PlannedStep{
			Step: RewriteStep{Rewrite: relation.SubjectSetRewrite},
			Req:  req,
		})
	}
	if (!strictMode || !hasRewrite) && !s.skipDirect {
		steps = append(steps, check.PlannedStep{
			Step: DirectStep{},
			Req:  req.WithDepth(req.RestDepth - 1),
		})
	}
	if canHaveSubjectSets {
		steps = append(steps, check.PlannedStep{
			Step: ExpandSubjectStep{},
			Req:  req.WithDepth(req.RestDepth - 1),
		})
	}

	switch len(steps) {
	case 0:
		return check.ResultNotMember
	case 1:
		return ex.Run(ctx, steps[0].Step, steps[0].Req)
	default:
		return ex.RunUnion(ctx, steps...)
	}
}
