// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package step

import (
	"context"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/relationtuple"
)

// ComputedUsersetStep rewrites the tuple's relation to the computed relation
// and delegates to IsAllowedStep.
type ComputedUsersetStep struct {
	Relation string
}

func (ComputedUsersetStep) Kind() check.StepKind { return check.StepComputed }

func (s ComputedUsersetStep) Execute(ctx context.Context, req check.CheckRequest, ex check.Executor) check.Result {
	ex.Deps().Logger().
		WithField("request", req.Tuple.String()).
		WithField("computed subjectSet relation", s.Relation).
		Trace("check computed subjectSet")

	childReq := check.CheckRequest{
		Tuple: &relationtuple.RelationTuple{
			Namespace: req.Tuple.Namespace,
			Object:    req.Tuple.Object,
			Relation:  s.Relation,
			Subject:   req.Tuple.Subject,
		},
		RestDepth: req.RestDepth,
	}
	return ex.Run(ctx, IsAllowedStep{}, childReq)
}
