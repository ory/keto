package config

import (
	"github.com/rs/cors"

	"github.com/ory/keto/internal/namespace"

	"github.com/ory/x/tracing"
)

type Provider interface {
	namespace.ManagerProvider

	CORS() (cors.Options, bool)
	RESTListenOn() string
	GRPCListenOn() string
	DSN() string
	TracingServiceName() string
	TracingProvider() string
	TracingConfig() *tracing.Config
	Set(key string, v interface{})
}

const DSNMemory = "sqlite://:memory:?_fk=true"
