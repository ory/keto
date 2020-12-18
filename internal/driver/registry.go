package driver

import (
	"context"

	"github.com/ory/x/healthx"
	"github.com/ory/x/tracing"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

type (
	Registry interface {
		Init(context.Context) error
		BuildVersion() string
		BuildDate() string
		BuildHash() string
		Config() config.Provider

		x.LoggerProvider
		x.WriterProvider

		relationtuple.ManagerProvider
		namespace.MigratorProvider
		expand.EngineProvider
		check.EngineProvider
		persistence.MigratorProvider
		persistence.Provider

		HealthHandler() *healthx.Handler
		Tracer() *tracing.Tracer
	}

	contextKeys string
)

const LogrusHookContextKey contextKeys = "logrus hook"
