package e2e

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/internal/x"

	"github.com/ory/x/configx"
	"github.com/ory/x/healthx"
	"github.com/phayes/freeport"
	"github.com/spf13/pflag"

	"github.com/ory/keto/internal/driver/config"

	"github.com/ory/keto/internal/namespace"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/sjson"

	"github.com/ory/keto/internal/driver"
)

func configFile(t testing.TB, values map[string]interface{}) string {
	dir := t.TempDir()
	fn := filepath.Join(dir, "keto.yml")

	c := []byte("{}")
	for key, val := range values {
		var err error
		c, err = sjson.SetBytes(c, key, val)
		require.NoError(t, err)
	}

	require.NoError(t, ioutil.WriteFile(fn, c, 0600))

	return fn
}

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
		[]string{"--" + configx.FlagConfig, configFile(t, map[string]interface{}{
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
	status := &bytes.Buffer{}

	require.NoError(t, r.Migrator().MigrationStatus(ctx, status))

	if strings.Contains(status.String(), "Pending") {
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
	status := &bytes.Buffer{}
	require.NoError(t, r.Migrator().MigrationStatus(ctx, status))
	assert.Contains(t, status.String(), "Applied")
	assert.NotContains(t, status.String(), "Pending")

	for _, n := range nn {
		status := &bytes.Buffer{}
		require.NoError(t, r.NamespaceMigrator().NamespaceStatus(ctx, status, n))
		assert.Contains(t, status.String(), "Applied")
		assert.NotContains(t, status.String(), "Pending")
	}
}

func startServer(ctx context.Context, t testing.TB, reg driver.Registry) func() {
	// Start the server
	serverCtx, serverCancel := context.WithCancel(ctx)
	serverErr := make(chan error)
	go func() {
		serverErr <- reg.ServeAll(serverCtx)
	}()

	var healthReady = func() error {
		ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
		defer cancel()

		r, err := http.NewRequestWithContext(ctx, "GET", "http://"+reg.Config().ReadAPIListenOn()+healthx.ReadyCheckPath, nil)
		if err != nil {
			return err
		}
		_, err = http.DefaultClient.Do(r)
		return err
	}
	// wait for /health/ready
	for err := healthReady(); err != nil; err = healthReady() {
		time.Sleep(10 * time.Millisecond)
	}

	// defer this close function to make sure it is shutdown on test failure as well
	return func() {
		// stop the server
		serverCancel()
		// wait for it to stop
		require.NoError(t, <-serverErr)
	}
}
