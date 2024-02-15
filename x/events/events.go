// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package events

import (
	"context"

	"github.com/ory/x/otelx/semconv"
	"go.opentelemetry.io/otel/trace"
)

const (
	RelationtuplesCreated semconv.Event = "RelationtuplesCreated"
	RelationtuplesDeleted semconv.Event = "RelationtuplesDeleted"
	RelationtuplesChanged semconv.Event = "RelationtuplesChanged"

	PermissionsExpanded semconv.Event = "PermissionsExpanded"
	PermissionsChecked  semconv.Event = "PermissionsChecked"
)

func NewRelationtuplesCreated(ctx context.Context) (string, trace.EventOption) {
	return RelationtuplesCreated.String(),
		trace.WithAttributes(
			semconv.AttributesFromContext(ctx)...,
		)
}

func NewRelationtuplesDeleted(ctx context.Context) (string, trace.EventOption) {
	return RelationtuplesDeleted.String(),
		trace.WithAttributes(
			semconv.AttributesFromContext(ctx)...,
		)
}

func NewRelationtuplesChanged(ctx context.Context) (string, trace.EventOption) {
	return RelationtuplesChanged.String(),
		trace.WithAttributes(
			semconv.AttributesFromContext(ctx)...,
		)
}

func NewPermissionsExpanded(ctx context.Context) (string, trace.EventOption) {
	return PermissionsExpanded.String(),
		trace.WithAttributes(
			semconv.AttributesFromContext(ctx)...,
		)
}

func NewPermissionsChecked(ctx context.Context) (string, trace.EventOption) {
	return PermissionsChecked.String(),
		trace.WithAttributes(
			semconv.AttributesFromContext(ctx)...,
		)
}
