package configuration

import (
	"fmt"

	"github.com/rs/cors"

	"github.com/ory/x/corsx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/tracing"
	"github.com/ory/x/viperx"
)

const (
	ViperKeyDSN  = "dsn"
	ViperKeyHost = "serve.host"
	ViperKeyPort = "serve.port"
)

type ViperProvider struct {
	l *logrusx.Logger
}

func NewViperProvider(l *logrusx.Logger) Provider {
	return &ViperProvider{l: l}
}

func (v *ViperProvider) ListenOn() string {
	return fmt.Sprintf(
		"%s:%d",
		viperx.GetString(v.l, ViperKeyHost, "", "HOST"),
		viperx.GetInt(v.l, ViperKeyPort, 4466, "PORT"),
	)
}

func (v *ViperProvider) CORSEnabled() bool {
	return corsx.IsEnabled(v.l, "serve")
}

func (v *ViperProvider) CORSOptions() cors.Options {
	return corsx.ParseOptions(v.l, "serve")
}

func (v *ViperProvider) DSN() string {
	return viperx.GetString(v.l, ViperKeyDSN, "", "DATABASE_URL")
}

func (v *ViperProvider) TracingServiceName() string {
	return viperx.GetString(v.l, "tracing.service_name", "ORY Keto")
}

func (v *ViperProvider) TracingProvider() string {
	return viperx.GetString(v.l, "tracing.provider", "", "TRACING_PROVIDER")
}

func (v *ViperProvider) TracingJaegerConfig() *tracing.JaegerConfig {
	return &tracing.JaegerConfig{
		LocalAgentHostPort: viperx.GetString(v.l, "tracing.providers.jaeger.local_agent_address", "", "TRACING_PROVIDER_JAEGER_LOCAL_AGENT_ADDRESS"),
		SamplerType:        viperx.GetString(v.l, "tracing.providers.jaeger.sampling.type", "const", "TRACING_PROVIDER_JAEGER_SAMPLING_TYPE"),
		SamplerValue:       viperx.GetFloat64(v.l, "tracing.providers.jaeger.sampling.value", float64(1), "TRACING_PROVIDER_JAEGER_SAMPLING_VALUE"),
		SamplerServerURL:   viperx.GetString(v.l, "tracing.providers.jaeger.sampling.server_url", "", "TRACING_PROVIDER_JAEGER_SAMPLING_SERVER_URL"),
		Propagation:        viperx.GetString(v.l, "tracing.providers.jaeger.propagation", "", "TRACING_PROVIDER_JAEGER_PROPAGATION"),
	}
}
