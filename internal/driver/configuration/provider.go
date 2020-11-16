package configuration

import (
	"github.com/rs/cors"

	"github.com/ory/x/logrusx"
	"github.com/ory/x/tracing"
)

type Provider interface {
	CORSEnabled() bool
	CORSOptions() cors.Options
	ListenOn() string
	DSN() string
	TracingServiceName() string
	TracingProvider() string
	TracingJaegerConfig() *tracing.JaegerConfig
}

func MustValidate(l *logrusx.Logger, p Provider) {

}
