package e2e

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/internal/x"

	"github.com/ory/x/configx"
	"github.com/phayes/freeport"
	"github.com/spf13/pflag"

	"github.com/ory/keto/internal/driver/config"

	"github.com/ory/keto/internal/namespace"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
)

func setup(t testing.TB) (*test.Hook, context.Context) {
	hook := &test.Hook{}
	ctx, cancel := context.WithCancel(context.WithValue(context.Background(), driver.LogrusHookContextKey, hook))
	t.Cleanup(func() {
		cancel()
	})

	return hook, ctx
}

func newInitializedReg(t testing.TB, dsn *x.DsnT, nspaces []*namespace.Namespace) (context.Context, driver.Registry) {
	_, ctx := setup(t)

	ports, err := freeport.GetFreePorts(2)
	require.NoError(t, err)

	t.Logf("starting server with dsn %+v; ports %+v", dsn.Name, ports)

	flags := pflag.NewFlagSet("", pflag.ContinueOnError)
	configx.RegisterConfigFlag(flags, nil)

	require.NoError(t, flags.Parse(
		[]string{"--" + configx.FlagConfig, x.ConfigFile(t, map[string]interface{}{
			config.KeyDSN:          dsn.Conn,
			config.KeyNamespaces:   nspaces,
			"log.level":            "debug",
			config.KeyReadAPIHost:  "127.0.0.1",
			config.KeyReadAPIPort:  ports[0],
			config.KeyWriteAPIHost: "127.0.0.1",
			config.KeyWriteAPIPort: ports[1],
		})},
	))

	reg, err := driver.NewDefaultRegistry(ctx, flags)
	require.NoError(t, err)

	if dsn.Name != "memory" {
		migrateEverythingUp(ctx, t, reg, nspaces)
	}
	assertMigrated(ctx, t, reg, nspaces)

	return ctx, reg
}

func migrateEverythingUp(ctx context.Context, t testing.TB, r driver.Registry, nn []*namespace.Namespace) {
	s, err := r.Migrator().MigrationStatus(ctx)
	require.NoError(t, err)

	if s.HasPending() {
		require.NoError(t, r.Migrator().MigrateUp(ctx))
	}

	for _, n := range nn {
		require.NoError(t, r.NamespaceMigrator().MigrateNamespaceUp(ctx, n))
	}

	t.Cleanup(func() {
		for _, n := range nn {
			require.NoError(t, r.NamespaceMigrator().MigrateNamespaceDown(context.Background(), n, 0))
		}

		require.NoError(t, r.Migrator().MigrateDown(context.Background(), 0))
	})
}

func assertMigrated(ctx context.Context, t testing.TB, r driver.Registry, nn []*namespace.Namespace) {
	s, err := r.Migrator().MigrationStatus(ctx)
	require.NoError(t, err)
	assert.False(t, s.HasPending())

	for _, n := range nn {
		s, err := r.NamespaceMigrator().NamespaceStatus(ctx, n)
		require.NoError(t, err)
		assert.False(t, s.HasPending())
	}
}

func startServer(ctx context.Context, t testing.TB, reg driver.Registry) func() {
	// Start the server
	serverCtx, serverCancel := context.WithCancel(ctx)
	serverErr := make(chan error)
	go func() {
		serverErr <- reg.ServeAll(serverCtx)
	}()

	// defer this close function to make sure it is shutdown on test failure as well
	return func() {
		// stop the server
		serverCancel()
		// wait for it to stop
		require.NoError(t, <-serverErr)
	}
}
