// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package x

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/ory/herodot"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/otelx"
)

type LoggerProvider interface {
	Logger() *logrusx.Logger
}

type WriterProvider interface {
	Writer() herodot.Writer
}

type TracingProvider interface {
	Tracer(ctx context.Context) *otelx.Tracer
}

type NetworkIDProvider interface {
	NetworkID(context.Context) uuid.UUID
}

type TransactorProvider interface {
	Transactor() interface {
		Transaction(ctx context.Context, f func(ctx context.Context) error) error
	}
}
