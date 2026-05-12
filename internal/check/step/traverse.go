// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package step

import (
	"context"

	"github.com/ory/x/otelx"
	keysetpagination "github.com/ory/x/pagination/keysetpagination_v2"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/attribute"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/namespace/ast"
	"github.com/ory/keto/internal/relationtuple"
)

// TraverseStep implements the tuple-to-subject-set (TTU) rewrite.
// It loads all tuples matching the TTU relation on the request object, then
// for each subject-set subject checks membership under the computed relation.
type TraverseStep struct {
	TTU *ast.TupleToSubjectSet
}

func (TraverseStep) Kind() check.StepKind { return check.StepTraverse }

func (s TraverseStep) Execute(ctx context.Context, req check.CheckRequest, ex check.Executor) (res check.Result) {
	ctx, span := ex.Deps().Tracer(ctx).Tracer().Start(ctx, "check.step.TraverseStep")
	var rowCount int
	defer func() {
		span.SetAttributes(attribute.Int("tuples_loaded", rowCount))
		otelx.End(span, &res.Err)
	}()

	if req.RestDepth <= 0 {
		ex.Deps().Logger().Debug("reached max-depth, therefore this query will not be further expanded")
		return check.Result{Membership: check.MembershipUnknown}
	}

	ex.Deps().Logger().
		WithField("request", req.Tuple.String()).
		WithField("tuple to subject-set relation", s.TTU.Relation).
		WithField("tuple to subject-set computed", s.TTU.ComputedSubjectSetRelation).
		Trace("check tuple to subjectSet")

	runner := ex.NewUnionRunner(ctx)
	for nextPage, _ := keysetpagination.NewPaginator(); !nextPage.IsLast() && !runner.Done(); {
		var tuples []*relationtuple.RelationTuple
		var err error
		tuples, nextPage, err = ex.Deps().RelationTupleManager().GetRelationTuples(runner.Ctx(), &query{
			Namespace: &req.Tuple.Namespace,
			Object:    &req.Tuple.Object,
			Relation:  &s.TTU.Relation,
		}, nextPage.ToOptions()...)
		if err != nil {
			return check.Result{Err: errors.WithStack(err)}
		}
		rowCount += len(tuples)
		check.ReportTuplesLoaded(runner.Ctx(), len(tuples))

		for _, t := range tuples {
			subSet, ok := t.Subject.(*relationtuple.SubjectSet)
			if !ok {
				continue
			}
			runner.Add(check.PlannedStep{
				Step: IsAllowedStep{},
				Req: check.CheckRequest{
					Tuple: &relationtuple.RelationTuple{
						Namespace: subSet.Namespace,
						Object:    subSet.Object,
						Relation:  s.TTU.ComputedSubjectSetRelation,
						Subject:   req.Tuple.Subject,
					},
					RestDepth: req.RestDepth - 1,
				},
			})
			if runner.Done() {
				break
			}
		}
	}
	return runner.Result()
}
