// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package events

import (
	"context"

	"github.com/ory/x/otelx/semconv"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const (
	RelationtuplesCreated semconv.Event = "RelationtuplesCreated"
	RelationtuplesDeleted semconv.Event = "RelationtuplesDeleted"
	RelationtuplesChanged semconv.Event = "RelationtuplesChanged"

	PermissionsExpanded semconv.Event = "PermissionsExpanded"
	PermissionsChecked  semconv.Event = "PermissionsChecked"
)

// Emit adds an event to the current span in the context.
func Emit(ctx context.Context, event semconv.Event, opt ...attribute.KeyValue) {
	trace.SpanFromContext(ctx).AddEvent(
		event.String(),
		trace.WithAttributes(
			append(
				semconv.AttributesFromContext(ctx),
				opt...,
			)...,
		),
	)
}
