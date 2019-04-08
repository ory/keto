package configuration

import (
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"

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

func MustValidate(l logrus.FieldLogger, p Provider) {

}
