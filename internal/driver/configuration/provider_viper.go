package configuration

import (
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/ory/keto/internal/namespace"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/rs/cors"

	"github.com/ory/x/corsx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/tracing"
	"github.com/ory/x/viperx"
)

const (
	ViperKeyDSN = "dsn"

	ViperKeyHost = "serve.host"
	ViperKeyPort = "serve.port"

	ViperKeyNamespacePath = "namespaces.dir_path"
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
	return viperx.GetString(v.l, ViperKeyDSN, "memory", "DATABASE_URL")
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
	}
}

func (v *ViperProvider) Namespaces() []*namespace.Namespace {
	namespaceDir := viperx.GetString(v.l, ViperKeyNamespacePath, "./keto-namespaces")

	infos, err := ioutil.ReadDir(namespaceDir)
	if err != nil {
		v.l.WithError(err).Errorf("Could no read namespace directory %s.", namespaceDir)
		return nil
	}

	var namespaces []*namespace.Namespace
	for _, info := range infos {
		fn := info.Name()

		if info.IsDir() || !(strings.HasSuffix(fn, ".yaml") || strings.HasSuffix(fn, ".yml") || strings.HasSuffix(fn, ".json")) {
			v.l.Infof("Skipping file %s in namespace directory because it is not *.{yaml|yml|json}", fn)
			continue
		}

		fc, err := ioutil.ReadFile(filepath.Join(namespaceDir, fn))
		if err != nil {
			v.l.WithError(err).Errorf("Could not read namespace file %s.", fn)
			continue
		}

		namesp := namespace.Namespace{}
		if err := yaml.Unmarshal(fc, &namesp); err != nil {
			v.l.WithError(err).Errorf("Could not unmarshal namespace file %s.", fn)
			continue
		}

		namespaces = append(namespaces, &namesp)
	}

	return namespaces
}
