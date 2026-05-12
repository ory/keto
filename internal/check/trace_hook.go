// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check

import (
	"context"

	"github.com/ory/keto/internal/relationtuple"
)

type traceHooksKey struct{}

// TraceHooks carries optional instrumentation callbacks injected into context
// by tracing middleware. Step implementations can emit events through helpers
// in this file without importing the trace package.
type TraceHooks struct {
	TuplesLoaded func(n int)
	SubjectFound func(tuple *relationtuple.RelationTuple)
}

// WithTraceHooks returns a context that carries the given trace hooks.
func WithTraceHooks(ctx context.Context, hooks TraceHooks) context.Context {
	return context.WithValue(ctx, traceHooksKey{}, hooks)
}

func traceHooksFromContext(ctx context.Context) (TraceHooks, bool) {
	hooks, ok := ctx.Value(traceHooksKey{}).(TraceHooks)
	return hooks, ok
}

// ReportTuplesLoaded calls the tuples-loaded hook stored in ctx, if any.
func ReportTuplesLoaded(ctx context.Context, n int) {
	if hooks, ok := traceHooksFromContext(ctx); ok && hooks.TuplesLoaded != nil {
		hooks.TuplesLoaded(n)
	}
}

// ReportSubjectFound calls the subject-found hook stored in ctx, if any.
func ReportSubjectFound(ctx context.Context, tuple *relationtuple.RelationTuple) {
	if hooks, ok := traceHooksFromContext(ctx); ok && hooks.SubjectFound != nil {
		hooks.SubjectFound(tuple)
	}
}
