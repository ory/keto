package driver

import (
	"context"
	"net/http"

	"github.com/ory/keto/internal/driver/config"

	"github.com/spf13/cobra"

	"google.golang.org/grpc"

	"github.com/ory/x/healthx"
	"github.com/ory/x/tracing"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/expand"
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

		config.Provider
		x.LoggerProvider
		x.WriterProvider

		relationtuple.ManagerProvider
		expand.EngineProvider
		check.EngineProvider
		persistence.Migrator

		HealthHandler() *healthx.Handler
		Tracer() *tracing.Tracer

		ReadRouter() http.Handler
		WriteRouter() http.Handler

		ReadGRPCServer() *grpc.Server
		WriteGRPCServer() *grpc.Server

		ServeAll(ctx context.Context) error
		ServeAllSQA(cmd *cobra.Command) error
		ServeRead(ctx context.Context) func() error
		ServeWrite(ctx context.Context) func() error
	}

	contextKeys string
)

const LogrusHookContextKey contextKeys = "logrus hook"
