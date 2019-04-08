package driver

import (
	"github.com/open-policy-agent/opa/ast"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/ory/keto/driver/configuration"
	"github.com/ory/keto/engine"
	"github.com/ory/keto/engine/ladon"
	"github.com/ory/keto/storage"
	"github.com/ory/keto/x"
	"github.com/ory/x/dbal"
	"github.com/ory/x/healthx"
	"github.com/ory/x/tracing"
)

type Registry interface {
	dbal.Driver
	Init() error
	WithConfig(c configuration.Provider) Registry
	WithLogger(l logrus.FieldLogger) Registry
	WithBuildInfo(version, hash, date string) Registry
	BuildVersion() string
	BuildDate() string
	BuildHash() string

	x.RegistryLogger
	x.RegistryWriter
	engine.Registry
	storage.Registry

	EngineCompiler() *ast.Compiler
	StorageHandler() *storage.Handler
	HealthHandler() *healthx.Handler
	LadonEngine() *ladon.Engine
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
