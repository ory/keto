// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package step

import (
	"context"

	"github.com/ory/herodot"
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x/graph"
)

// ExpandSubjectStep expands all subject-set pointers stored under the tuple's
// relation and recursively checks each one.
type ExpandSubjectStep struct{}

func (ExpandSubjectStep) Kind() check.StepKind { return check.StepExpand }

func (ExpandSubjectStep) Execute(ctx context.Context, req check.CheckRequest, ex check.Executor) check.Result {
	if req.RestDepth <= 0 {
		ex.Deps().Logger().
			WithField("request", req.Tuple.String()).
			Debug("reached max-depth, therefore this query will not be further expanded")
		return check.Result{Membership: check.MembershipUnknown}
	}

	ex.Deps().Logger().
		WithField("request", req.Tuple.String()).
		Trace("check expand subject")

	results, err := ex.Deps().Traverser().TraverseSubjectSetExpansion(ctx, req.Tuple)
	if errors.Is(err, herodot.ErrNotFound()) {
		return check.ResultNotMember
	} else if err != nil {
		return check.Result{Err: errors.WithStack(err)}
	}

	check.ReportTuplesLoaded(ctx, len(results))

	// Check whether the current hop already resolved the query.
	for _, result := range results {
		if result.Found {
			check.ReportSubjectFound(ctx, result.To)
			return check.ResultIsMember
		}
	}

	// Build the visited-tracking context from the current ctx so that child
	// checks share the same visited set and prevent infinite loops.
	var visited bool
	visitCtx := graph.InitVisited(ctx)

	runner := ex.NewUnionRunner(visitCtx)
	if maxWidth := ex.Deps().Config(ctx).MaxReadWidth(); len(results) > maxWidth {
		ex.Deps().Logger().
			WithField("method", "ExpandSubjectStep").
			WithField("request", req.Tuple.String()).
			WithField("max_width", maxWidth).
			WithField("results", len(results)).
			Debug("too many results, truncating to width limit")
		results = results[:maxWidth]
	}

	for _, result := range results {
		if runner.Done() {
			break
		}
		sub := &relationtuple.SubjectSet{
			Namespace: result.To.Namespace,
			Object:    result.To.Object,
			Relation:  result.To.Relation,
		}
		visitCtx, visited = graph.CheckAndAddVisited(visitCtx, sub)
		if visited {
			continue
		}
		runner.Add(check.PlannedStep{
			Step: IsAllowedStep{skipDirect: true},
			Req:  check.CheckRequest{Tuple: result.To, RestDepth: req.RestDepth},
		})
	}
	return runner.Result()
}
