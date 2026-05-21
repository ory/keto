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

// ExpandSubjectStep expands subject-set pointers under the tuple's relation and recursively checks each one.
type ExpandSubjectStep struct {
	StrictMode bool
	// SubjectSetTypes are the subject-set types declared in OPL for this relation.
	SubjectSetTypes []relationtuple.SubjectSetType
}

func (ExpandSubjectStep) Kind() check.StepKind { return check.StepExpand }

func (s ExpandSubjectStep) Execute(ctx context.Context, req check.CheckRequest, ex check.Executor) check.Result {
	if req.RestDepth <= 0 {
		return maxDepthReached(ex, req)
	}

	ex.Deps().Logger().
		WithField("request", req.Tuple.String()).
		Trace("check expand subject")

	subjectSetTypes := s.SubjectSetTypes
	if !s.StrictMode {
		subjectSetTypes = nil
	}

	results, err := ex.Deps().Traverser().TraverseSubjectSetExpansion(ctx, req.Tuple, subjectSetTypes)
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

	if len(results) == 0 {
		return check.ResultNotMember
	}

	// If we have results, but no direct hits, we need to check the next level of expansion.
	// We check the depth here to prevent planning a large number of child checks when they will already be at max depth.
	if req.RestDepth <= 1 {
		return maxDepthReached(ex, req)
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

		runner.AddLimitation(check.LimitationMaxWidthExceeded)
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
			Req:  check.CheckRequest{Tuple: result.To, RestDepth: req.RestDepth - 1},
		})
	}
	return runner.Result()
}
