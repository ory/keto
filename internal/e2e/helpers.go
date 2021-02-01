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

	"github.com/ory/x/configx"
	"github.com/ory/x/healthx"
	"github.com/ory/x/sqlcon/dockertest"
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

func migrateEverythingUp(ctx context.Context, t testing.TB, r driver.Registry, nn []*namespace.Namespace) {
	status := &bytes.Buffer{}

	require.NoError(t, r.Migrator().MigrationStatus(ctx, status))

	if strings.Contains(status.String(), "Pending") {
		require.NoError(t, r.Migrator().MigrateUp(ctx))
	}

	for _, n := range nn {
		require.NoError(t, r.NamespaceMigrator().MigrateNamespaceUp(ctx, n))
	}

	// TODO
	//t.Cleanup(func() {
	//	for _, n := range nn {
	//		c.ExecNoErr(t, "namespace", "migrate", "down", n.Name, "1")
	//	}
	//
	//	c.ExecNoErr(t, "migrate", "down", "1")
	//})
}

type DsnT struct {
	Name    string
	Conn    string
	Prepare func(context.Context, testing.TB, driver.Registry, []*namespace.Namespace)
}

func GetDSNs(t testing.TB) []*DsnT {
	// we use a slice of structs here to always have the same execution order
	dsns := []*DsnT{
		{
			Name: "memory",
			Conn: "memory",
		},
	}
	if !testing.Short() {
		dsns = append(dsns,
			&DsnT{
				Name:    "mysql",
				Conn:    dockertest.RunTestMySQL(t),
				Prepare: migrateEverythingUp,
			},
			&DsnT{
				Name:    "postgres",
				Conn:    dockertest.RunTestPostgreSQL(t),
				Prepare: migrateEverythingUp,
			},
			&DsnT{
				Name:    "cockroach",
				Conn:    dockertest.RunTestCockroachDB(t),
				Prepare: migrateEverythingUp,
			},
		)
	}
	t.Cleanup(dockertest.KillAllTestDatabases)

	return dsns
}

func NewInitializedReg(t testing.TB, dsn *DsnT, nspaces []*namespace.Namespace) (context.Context, driver.Registry) {
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

	if dsn.Prepare != nil {
		dsn.Prepare(ctx, t, reg, nspaces)
	}

	return ctx, reg
}

func startServer(t testing.TB, dsn *DsnT, nspaces []*namespace.Namespace) (context.Context, driver.Registry, func()) {
	ctx, reg := NewInitializedReg(t, dsn, nspaces)
	// Initialization done

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

	return ctx,
		reg,
		// defer this close function to make sure it is shutdown on test failure as well
		func() {
			// stop the server
			serverCancel()
			// wait for it to stop
			require.NoError(t, <-serverErr)
		}
}
