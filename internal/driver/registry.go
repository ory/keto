package driver

import (
	"github.com/ory/x/dbal"
	"github.com/ory/x/healthx"
	"github.com/ory/x/tracing"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"

	"github.com/ory/keto/internal/x"
)

type Registry interface {
	dbal.Driver
	Init() error
	BuildVersion() string
	BuildDate() string
	BuildHash() string

	x.LoggerProvider
	x.WriterProvider

	relationtuple.ManagerProvider
	namespace.MigratorProvider
	expand.EngineProvider
	check.EngineProvider

	HealthHandler() *healthx.Handler
	Tracer() *tracing.Tracer
}
