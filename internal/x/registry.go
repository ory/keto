package x

import (
	"context"

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
