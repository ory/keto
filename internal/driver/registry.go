// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/ory/pop/v6"
	"github.com/ory/x/healthx"
	"github.com/ory/x/httpx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/otelx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/internal/relationtuple"
)

type (
	Registry interface {
		Init(context.Context) error

		config.Provider
		logrusx.Provider
		httpx.WriterProvider

		relationtuple.ManagerProvider
		relationtuple.MapperProvider
		expand.EngineProvider
		check.CheckerProvider
		check.EngineProvider
		persistence.Migrator
		persistence.Provider

		PopConnection(ctx context.Context) (*pop.Connection, error)
		PopConnectionWithOpts(ctx context.Context, f ...func(*pop.ConnectionDetails)) (*pop.Connection, error)

		HealthHandler() *healthx.Handler
		Tracer(ctx context.Context) *otelx.Tracer

		ReadRouter(ctx context.Context) http.Handler
		WriteRouter(ctx context.Context) http.Handler
		OPLSyntaxRouter(ctx context.Context) http.Handler

		ServeAll(ctx context.Context) error
		ServeAllSQA(cmd *cobra.Command) error

		HandlerOptions() []connect.HandlerOption
	}

	contextKeys string
)

const (
	LogrusHookContextKey contextKeys = "logrus hook"
	RegistryContextKey   contextKeys = "registry"
)
