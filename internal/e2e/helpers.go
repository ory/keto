package e2e

import (
	"context"
	"errors"
	"testing"

	"github.com/ory/keto/internal/x/dbx"

	"github.com/ory/keto/internal/relationtuple"

	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/internal/x"

	"github.com/ory/x/configx"
	"github.com/phayes/freeport"
	"github.com/spf13/pflag"

	"github.com/ory/keto/internal/driver/config"

	"github.com/ory/keto/internal/namespace"

	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
)

func newInitializedReg(t testing.TB, dsn *dbx.DsnT, cfgOverwrites map[string]interface{}) (context.Context, driver.Registry, func(*testing.T, ...*namespace.Namespace)) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(func() {
		cancel()
	})

	ports, err := freeport.GetFreePorts(3)
	require.NoError(t, err)

	flags := pflag.NewFlagSet("", pflag.ContinueOnError)
	configx.RegisterConfigFlag(flags, nil)

	cfgValues := map[string]interface{}{
		config.KeyDSN:               dsn.Conn,
		"log.level":                 "debug",
		"log.leak_sensitive_values": true,
		config.KeyReadAPIHost:       "127.0.0.1",
		config.KeyReadAPIPort:       ports[0],
		config.KeyWriteAPIHost:      "127.0.0.1",
		config.KeyWriteAPIPort:      ports[1],
		config.KeyMetricsHost:       "127.0.0.1",
		config.KeyMetricsPort:       ports[2],
	}
	for k, v := range cfgOverwrites {
		cfgValues[k] = v
	}

	cf := dbx.ConfigFile(t, cfgValues)
	require.NoError(t, flags.Parse([]string{"--" + configx.FlagConfig, cf}))

	reg, err := driver.NewDefaultRegistry(ctx, flags, true)
	require.NoError(t, err)

	require.NoError(t, reg.MigrateUp(ctx))
	assertMigrated(ctx, t, reg)

	nspaces := make([]*namespace.Namespace, 0)
	addNamespaces := func(t *testing.T, nn ...*namespace.Namespace) {
		for _, n := range nn {
			n.ID = int32(len(nspaces))
			nspaces = append(nspaces, n)
		}

		require.NoError(t, reg.Config().Set(config.KeyNamespaces, nspaces))

		t.Cleanup(func() {
			for _, n := range nn {
				err := errors.New("not nil")
				var pt string
				var ts []*relationtuple.InternalRelationTuple
				for pt != "" || err != nil {
					ts, pt, err = reg.RelationTupleManager().GetRelationTuples(ctx, &relationtuple.RelationQuery{
						Namespace: n.Name,
					}, x.WithToken(pt))
					require.NoError(t, err)
					require.NoError(t, reg.RelationTupleManager().DeleteRelationTuples(ctx, ts...))
				}
			}
		})
	}

	return ctx, reg, addNamespaces
}

func assertMigrated(ctx context.Context, t testing.TB, r driver.Registry) {
	mb, err := r.MigrationBox()
	require.NoError(t, err)
	s, err := mb.Status(ctx)
	require.NoError(t, err)
	assert.False(t, s.HasPending())
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
