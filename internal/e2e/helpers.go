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
			config.KeyDSN:               dsn.Conn,
			config.KeyNamespaces:        nspaces,
			"log.level":                 "debug",
			"log.leak_sensitive_values": true,
			config.KeyReadAPIHost:       "127.0.0.1",
			config.KeyReadAPIPort:       ports[0],
			config.KeyWriteAPIHost:      "127.0.0.1",
			config.KeyWriteAPIPort:      ports[1],
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
	mb, err := r.Migrator().MigrationBox(ctx)
	require.NoError(t, err)
	s, err := mb.Status(ctx)
	require.NoError(t, err)

	if s.HasPending() {
		require.NoError(t, mb.Up(ctx))
	}

	for _, n := range nn {
		nmb, err := r.NamespaceMigrator().NamespaceMigrationBox(ctx, n)
		require.NoError(t, err)
		require.NoError(t, nmb.Up(ctx))
	}

	t.Cleanup(func() {
		for _, n := range nn {
			nmb, err := r.NamespaceMigrator().NamespaceMigrationBox(context.Background(), n)
			require.NoError(t, err)
			require.NoError(t, nmb.Down(context.Background(), 0))
		}

		require.NoError(t, err)
		require.NoError(t, mb.Down(context.Background(), 0))
	})
}

func assertMigrated(ctx context.Context, t testing.TB, r driver.Registry, nn []*namespace.Namespace) {
	mb, err := r.Migrator().MigrationBox(ctx)
	require.NoError(t, err)
	s, err := mb.Status(ctx)
	require.NoError(t, err)
	assert.False(t, s.HasPending())

	for _, n := range nn {
		nmb, err := r.NamespaceMigrator().NamespaceMigrationBox(ctx, n)
		require.NoError(t, err)
		s, err := nmb.Status(ctx)
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
