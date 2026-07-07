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
		return maxDepthReached(ex, req)
	}

	ex.Deps().Logger().
		WithField("request", req.Tuple.String()).
		Trace("check is allowed")

	nm, err := ex.Deps().Config(ctx).NamespaceManager()
	if err != nil {
		return check.Result{Err: err}
	}

	relation, err := astRelationFor(ctx, nm, req.Tuple)
	if err != nil {
		return check.Result{Err: err}
	}

	hasRewrite := relation != nil && relation.SubjectSetRewrite != nil
	strictMode := ex.Deps().Config(ctx).StrictMode()

	// In strict mode, only run steps that OPL explicitly permits.
	// In non-strict mode, run all applicable steps for backwards compatibility.
	canHaveDirect := !strictMode || AllowsDirectMember(relation, req.Tuple.Subject)

	steps := make([]check.PlannedStep, 0, 3)

	if hasRewrite {
		steps = append(steps, check.PlannedStep{
			Step: RewriteStep{Rewrite: relation.SubjectSetRewrite},
			Req:  req,
		})
	}

	if canHaveDirect && !s.skipDirect {
		steps = append(steps, check.PlannedStep{
			Step: DirectStep{},
			Req:  req,
		})
	}

	// In strict mode, only expand when OPL explicitly declares subject-set types, and
	// filter expansion to only those declared types.
	// In non-strict mode, always expand to preserve backwards compatibility.
	if strictMode {
		subjectSetTypes, err := subjectSetTypesFor(ctx, nm, req.Tuple.Subject, relation)
		if err != nil {
			return check.Result{Err: err}
		}
		if len(subjectSetTypes) > 0 {
			steps = append(steps, check.PlannedStep{
				Step: ExpandSubjectStep{SubjectSetTypes: subjectSetTypes},
				Req:  req,
			})
		}
	} else {
		steps = append(steps, check.PlannedStep{
			Step: ExpandSubjectStep{},
			Req:  req,
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
