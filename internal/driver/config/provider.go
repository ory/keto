package config

import (
	"github.com/rs/cors"

	"github.com/ory/keto/internal/namespace"

	"github.com/ory/x/tracing"
)

type Provider interface {
	namespace.Manager

	CORS() (cors.Options, bool)
	ListenOn() string
	DSN() string
	TracingServiceName() string
	TracingProvider() string
	TracingConfig() *tracing.Config
	Set(key string, v interface{})
}

const DSNMemory = "memory"
