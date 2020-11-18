package driver

import (
	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/x/dbal"
	"github.com/ory/x/healthx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/tracing"

	"github.com/ory/keto/internal/driver/configuration"
	"github.com/ory/keto/internal/x"
)

type Registry interface {
	dbal.Driver
	Init() error
	WithConfig(c configuration.Provider) Registry
	WithLogger(l *logrusx.Logger) Registry
	WithBuildInfo(version, hash, date string) Registry
	BuildVersion() string
	BuildDate() string
	BuildHash() string

	x.LoggerProvider
	x.WriterProvider

	relationtuple.ManagerProvider
	namespace.ManagerProvider
	expand.EngineProvider
	check.EngineProvider

	HealthHandler() *healthx.Handler
	Tracer() *tracing.Tracer
}

func NewRegistry(c configuration.Provider) (Registry, error) {
	//driver, err := dbal.GetDriverFor(c.DSN())
	//if err != nil {
	//	return nil, err
	//}
	//
	//registry, ok := driver.(Registry)
	//if !ok {
	//	return nil, errors.Errorf("driver of type %T does not implement interface Registry", driver)
	//}

	registry := (&RegistryDefault{}).WithConfig(c)

	return registry, nil
}
