// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package events

import (
	"context"

	"github.com/ory/x/otelx/semconv"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/ory/keto/internal/x"
)

type Event string

const (
	RelationtuplesCreated Event = "RelationtuplesCreated"
	RelationtuplesDeleted Event = "RelationtuplesDeleted"
	RelationtuplesChanged Event = "RelationtuplesChanged"

	PermissionsExpanded Event = "PermissionsExpanded"
	PermissionsChecked  Event = "PermissionsChecked"
)

// Add adds an event to the current span in the context.
func Add(ctx context.Context, p x.NetworkIDProvider, event Event, opt ...attribute.KeyValue) {
	trace.SpanFromContext(semconv.ContextWithAttributes(ctx)).AddEvent(
		string(event),
		trace.WithAttributes(
			opt...,
		),
	)
}
