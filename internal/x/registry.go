package x

import (
	"github.com/ory/herodot"
	"github.com/ory/x/logrusx"
)

type LoggerProvider interface {
	Logger() *logrusx.Logger
}

type WriterProvider interface {
	Writer() herodot.Writer
}
