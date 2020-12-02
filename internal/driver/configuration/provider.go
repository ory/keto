package configuration

import (
	"github.com/rs/cors"

	"github.com/ory/keto/internal/namespace"

	"github.com/ory/x/logrusx"
	"github.com/ory/x/tracing"
)

type Provider interface {
	namespace.Manager

	CORSEnabled() bool
	CORSOptions() cors.Options
	ListenOn() string
	DSN() string
	TracingServiceName() string
	TracingProvider() string
	TracingJaegerConfig() *tracing.JaegerConfig
	Namespaces() []*namespace.Namespace
}

func MustValidate(l *logrusx.Logger, p Provider) {

}

const DSNMemory = "memory"
