package driver

import (
	"context"
	"net/http"

	"github.com/ory/x/otelx"
	prometheus "github.com/ory/x/prometheusx"

	"github.com/gobuffalo/pop/v6"

	"github.com/ory/keto/internal/driver/config"

	"github.com/spf13/cobra"

	"google.golang.org/grpc"

	"github.com/ory/x/healthx"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

type (
	Registry interface {
		Init(context.Context) error

		config.Provider
		x.LoggerProvider
		x.WriterProvider

		relationtuple.ManagerProvider
		expand.EngineProvider
		check.EngineProvider
		persistence.Migrator
		persistence.Provider

		PopConnection(ctx context.Context) (*pop.Connection, error)
		PopConnectionWithOpts(ctx context.Context, f ...func(*pop.ConnectionDetails)) (*pop.Connection, error)

		HealthHandler() *healthx.Handler
		Tracer(ctx context.Context) *otelx.Tracer
		MetricsHandler() *prometheus.Handler
		PrometheusManager() *prometheus.MetricsManager

		ReadRouter(ctx context.Context) http.Handler
		WriteRouter(ctx context.Context) http.Handler

		ReadGRPCServer(ctx context.Context) *grpc.Server
		WriteGRPCServer(ctx context.Context) *grpc.Server

		ServeAll(ctx context.Context) error
		ServeAllSQA(cmd *cobra.Command) error
	}

	contextKeys string
)

const (
	LogrusHookContextKey contextKeys = "logrus hook"
	RegistryContextKey   contextKeys = "registry"
)
