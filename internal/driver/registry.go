package driver

import (
	"github.com/pkg/errors"

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

	HealthHandler() *healthx.Handler
	Tracer() *tracing.Tracer
}

func NewRegistry(c configuration.Provider) (Registry, error) {
	driver, err := dbal.GetDriverFor(c.DSN())
	if err != nil {
		return nil, err
	}

	registry, ok := driver.(Registry)
	if !ok {
		return nil, errors.Errorf("driver of type %T does not implement interface Registry", driver)
	}

	registry = registry.WithConfig(c)

	if err := registry.Init(); err != nil {
		return nil, err
	}

	return registry, nil
}
