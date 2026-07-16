// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package step

import (
	"context"
	"net/http"

	"github.com/ory/herodot"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/ory/keto/internal/check"
)

type callBudgetKeyType struct{}

var callBudgetKey callBudgetKeyType

// maxCallSteps bounds the number of steps along a single resolution path. Most
// recursion is already bounded by RestDepth, but same-object computed hops
// (ComputedUsersetStep, InvertStep) don't decrement it, so a cyclic config
// (e.g. two permits that reference each other) would otherwise recurse until
// the stack overflows and crashes the process. This is a blanket per-path
// ceiling, well above any real config and well below a stack overflow.
const maxCallSteps = 1000

func errCallBudgetExceeded() error {
	return errors.WithStack(&herodot.DefaultError{
		StatusField:   http.StatusText(http.StatusUnprocessableEntity),
		ErrorField:    "The check could not be completed",
		ReasonField:   "The check was aborted after too many resolution steps. This usually indicates a cyclic permit reference in the namespace configuration.",
		CodeField:     http.StatusUnprocessableEntity,
		GRPCCodeField: codes.FailedPrecondition,
	})
}

// CallLimitMiddleware aborts a check once any single resolution path exceeds
// maxCallSteps. It returns a decisive error so the whole check stops
// immediately, rather than a limitation (which is non-decisive and would let a
// cyclic config keep exploring exponentially many branches). The budget is
// carried per-path in the context and seeded lazily.
func CallLimitMiddleware() check.Middleware {
	return func(ctx context.Context, s check.Step, req check.CheckRequest, next func(context.Context, check.Step, check.CheckRequest) check.Result) check.Result {
		remaining, ok := ctx.Value(callBudgetKey).(int)
		if !ok {
			remaining = maxCallSteps
		}
		if remaining <= 0 {
			return check.Result{Err: errCallBudgetExceeded()}
		}
		return next(context.WithValue(ctx, callBudgetKey, remaining-1), s, req)
	}
}
