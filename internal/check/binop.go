// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ory/keto/internal/check/checkgroup"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/ketoapi"
)

type binaryOperator = func(ctx context.Context, checks []checkgroup.CheckFunc) checkgroup.Result

func or(ctx context.Context, checks []checkgroup.CheckFunc) checkgroup.Result {
	if len(checks) == 0 {
		return checkgroup.ResultNotMember
	}

	resultCh := make(chan checkgroup.Result, 1)

	for _, check := range checks {
		check(ctx, resultCh)
		select {
		case result := <-resultCh:
			// We return either the first error or the first success.
			if result.Err != nil || result.Membership == checkgroup.IsMember {
				return result
			}
		case <-ctx.Done():
			return checkgroup.Result{Err: errors.WithStack(ctx.Err())}
		}
	}

	return checkgroup.ResultNotMember
}

func and(ctx context.Context, checks []checkgroup.CheckFunc) checkgroup.Result {
	if len(checks) == 0 {
		return checkgroup.ResultNotMember
	}

	resultCh := make(chan checkgroup.Result, 1)

	tree := &ketoapi.Tree[*relationtuple.RelationTuple]{
		Type:     ketoapi.TreeNodeIntersection,
		Children: []*ketoapi.Tree[*relationtuple.RelationTuple]{},
	}

	for _, check := range checks {
		check(ctx, resultCh)
		select {
		case result := <-resultCh:
			// We return fast on either an error or if a subcheck returns "not a
			// member".
			if result.Err != nil || result.Membership != checkgroup.IsMember {
				return checkgroup.Result{Err: result.Err, Membership: checkgroup.NotMember}
			} else {
				tree.Children = append(tree.Children, result.Tree)
			}
		case <-ctx.Done():
			return checkgroup.Result{Err: errors.WithStack(ctx.Err())}
		}
	}

	return checkgroup.Result{
		Membership: checkgroup.IsMember,
		Tree:       tree,
	}
}
