package sql_test

import (
	"context"
	"testing"

	"github.com/ory/x/configx"
	"github.com/spf13/pflag"

	"github.com/ory/keto/internal/x/dbx"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/persistence/sql"
)

func TestNetwork(t *testing.T) {
	var setup = func(t *testing.T, dsn *dbx.DsnT) (context.Context, driver.Registry) {
		ctx := context.Background()

		flags := pflag.NewFlagSet("", pflag.ContinueOnError)
		configx.RegisterConfigFlag(flags, nil)
		cf := dbx.ConfigFile(t, map[string]interface{}{
			config.KeyDSN: dsn.Conn,
			config.KeyNamespaces: []*namespace.Namespace{{
				ID:   0,
				Name: "default",
			}},
			"log.level":                 "debug",
			"log.leak_sensitive_values": true,
			config.KeyReadAPIPort:       0,
			config.KeyWriteAPIPort:      0,
		})
		require.NoError(t, flags.Parse([]string{"--" + configx.FlagConfig, cf}))

		reg, err := driver.NewDefaultRegistry(ctx, flags)
		require.NoError(t, err)

		mb, err := reg.Migrator().MigrationBox(ctx)
		require.NoError(t, err)
		require.NoError(t, mb.Up(ctx))

		return ctx, reg
	}

	for _, dsn := range dbx.GetDSNs(t, false) {
		t.Run("dsn="+dsn.Name, func(t *testing.T) {
			ctx, reg := setup(t, dsn)

			t.Run("case=creates network ID", func(t *testing.T) {
				nid, err := reg.Persister().NetworkID(ctx)
				require.NoError(t, err)
				assert.NotEqual(t, uuid.Nil, nid)

				t.Run("case=gets same ID again", func(t *testing.T) {
					nid2, err := reg.Persister().NetworkID(ctx)
					require.NoError(t, err)
					assert.Equal(t, nid, nid2)
				})

				t.Run("case=gets same ID with new persister", func(t *testing.T) {
					ctx, reg := setup(t, dsn)

					nid2, err := reg.Persister().NetworkID(ctx)
					require.NoError(t, err)
					assert.Equal(t, nid, nid2)
				})

				t.Run("case=db enforces single network ID", func(t *testing.T) {
					err := reg.Persister().Connection(ctx).Create(&sql.NetworkID{
						ID: uuid.Must(uuid.NewV4()),
					})
					require.NotNil(t, err)
					uniqueMsgs := map[string]string{
						"mysql":     "Duplicate entry",
						"memory":    "UNIQUE constraint",
						"postgres":  "unique constraint",
						"cockroach": "unique constraint",
					}
					assert.Contains(t, err.Error(), uniqueMsgs[dsn.Name], "%v", err)

					err = reg.Persister().Connection(ctx).Create(&sql.NetworkID{
						ID:      uuid.Must(uuid.NewV4()),
						Limiter: 1,
					})
					require.NotNil(t, err)
					checkMsgs := map[string]string{
						"mysql":     "Check constraint",
						"memory":    "CHECK constraint",
						"postgres":  "check constraint",
						"cockroach": "CHECK constraint",
					}
					assert.Contains(t, err.Error(), checkMsgs[dsn.Name], "%v", err)
				})
			})
		})
	}
}
