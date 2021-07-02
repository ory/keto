package e2e

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/ory/keto/internal/x/dbx"

	"github.com/tidwall/sjson"

	"github.com/ory/keto/internal/relationtuple"

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

func newInitializedReg(t testing.TB, dsn *dbx.DsnT) (context.Context, driver.Registry, func(*testing.T, ...*namespace.Namespace)) {
	hook, ctx := setup(t)

	ports, err := freeport.GetFreePorts(2)
	require.NoError(t, err)

	flags := pflag.NewFlagSet("", pflag.ContinueOnError)
	configx.RegisterConfigFlag(flags, nil)

	nspaces := make([]*namespace.Namespace, 0)
	cf := dbx.ConfigFile(t, map[string]interface{}{
		config.KeyDSN:               dsn.Conn,
		config.KeyNamespaces:        nspaces,
		"log.level":                 "debug",
		"log.leak_sensitive_values": true,
		config.KeyReadAPIHost:       "127.0.0.1",
		config.KeyReadAPIPort:       ports[0],
		config.KeyWriteAPIHost:      "127.0.0.1",
		config.KeyWriteAPIPort:      ports[1],
	})
	require.NoError(t, flags.Parse([]string{"--" + configx.FlagConfig, cf}))

	reg, err := driver.NewDefaultRegistry(ctx, flags)
	require.NoError(t, err)

	if dsn.Name != "memory" {
		migrateEverythingUp(ctx, t, reg)
	}
	assertMigrated(ctx, t, reg)

	return ctx, reg, func(t *testing.T, nn ...*namespace.Namespace) {
		for _, n := range nn {
			n.ID = int64(len(nspaces))
			nspaces = append(nspaces, n)
		}

		cc, err := os.ReadFile(cf)
		require.NoError(t, err)
		cc, err = sjson.SetBytes(cc, config.KeyNamespaces, nspaces)
		require.NoError(t, err)
		require.NoError(t, os.WriteFile(cf, cc, 0644))

		select {
		case <-time.After(time.Second):
			t.Errorf("did not get namespace update %+v", nspaces)
		case <-func() chan struct{} {
			c := make(chan struct{})
			go func() {
				defer close(c)

			pollLogEntries:
				for {
					ee := hook.AllEntries()
					for _, e := range ee {
						if f, ok := e.Data["file"]; ok && f == cf && e.Message == "Configuration change processed successfully." {
							hook.Reset()
							break pollLogEntries
						}
					}
					t.Logf("waiting for last entry to notify about config %s change, got %+v", cf, ee)
					time.Sleep(10 * time.Millisecond)
				}

				for {
					nm, err := reg.Config().NamespaceManager()
					require.NoError(t, err)

					if func() (allNamespacesThere bool) {
						for _, n := range nn {
							a, err := nm.GetNamespaceByID(ctx, n.ID)
							if err != nil {
								return false
							}
							assert.Equal(t, n, a)
						}
						return true
					}() {
						break
					}
					nn, err := nm.Namespaces(ctx)
					require.NoError(t, err)
					t.Logf("not all namespaces there yet %+v", nn)
					time.Sleep(10 * time.Millisecond)
				}
			}()
			return c
		}():
		}

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
}

func migrateEverythingUp(ctx context.Context, t testing.TB, r driver.Registry) {
	mb, err := r.Migrator().MigrationBox(ctx)
	require.NoError(t, err)
	s, err := mb.Status(ctx)
	require.NoError(t, err)

	if s.HasPending() {
		require.NoError(t, mb.Up(ctx))
	}

	t.Cleanup(func() {
		require.NoError(t, err)
		require.NoError(t, mb.Down(context.Background(), 0))
	})
}

func assertMigrated(ctx context.Context, t testing.TB, r driver.Registry) {
	mb, err := r.Migrator().MigrationBox(ctx)
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
