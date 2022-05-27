package migratest

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/ory/x/fsx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/networkx"
	"github.com/ory/x/popx"
	"github.com/ory/x/sqlcon"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/persistence/sql"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x/dbx"
)

func TestMigrations(t *testing.T) {
	t.Parallel()
	const debugOnDisk = false

	for _, db := range dbx.GetDSNs(t, debugOnDisk) {
		db := db
		t.Run("dsn="+db.Name, func(t *testing.T) {
			t.Parallel()

			db.MigrateUp, db.MigrateDown = false, false

			ctx := context.Background()

			var conn *pop.Connection
			var err error
			conn, err = pop.NewConnection(&pop.ConnectionDetails{URL: db.Conn})
			require.NoError(t, err)
			for i := 0; i < 120; i++ {
				require.NoError(t, conn.Open())
				if err := dbx.Ping(conn); err == nil {
					break
				}
				time.Sleep(time.Second)
			}
			require.NoError(t, dbx.Ping(conn))
			t.Cleanup(func() { conn.Close() })

			tm, err := popx.NewMigrationBox(
				fsx.Merge(sql.Migrations, networkx.Migrations),
				popx.NewMigrator(conn, logrusx.New("", "", logrusx.ForceLevel(logrus.DebugLevel)), nil, 1*time.Minute),
				popx.WithTestdata(t, os.DirFS("./testdata")),
			)
			require.NoError(t, err)

			// cleanup first
			require.NoError(t, tm.Down(ctx, -1))

			t.Run("suite=up", func(t *testing.T) {
				if err := tm.Up(ctx); err != nil {
					t.Log("migrations failed:", err)
					logMigrationStatus(t, tm)
					t.FailNow()
				}
			})

			reg := driver.NewTestRegistry(t, db)
			require.NoError(t,
				reg.Config(ctx).Set(config.KeyNamespaces, []*namespace.Namespace{
					{ID: 1, Name: "foo"},
					{ID: 2, Name: "uuid_test"},
				}))
			p, err := sql.NewPersister(ctx, reg, uuid.Must(uuid.FromString("77fdc5e0-2260-49da-8aae-c36ba255d05b")))

			t.Run("suite=fixtures", func(t *testing.T) {
				t.Run("table=legacy namespaces", func(t *testing.T) {
					// as they are legacy, we expect them to be actually dropped
					assert.ErrorIs(t, sqlcon.HandleError(conn.RawQuery(
						"SELECT * FROM keto_namespace",
					).Exec()), sqlcon.ErrNoSuchTable)
				})

				t.Run("table=relation tuples", func(t *testing.T) {
					require.NoError(t, err)
					actualRts, next, err := p.GetRelationTuples(ctx, &relationtuple.RelationQuery{Namespace: "foo"})
					require.NoError(t, err)
					assert.Equal(t, "", next)

					for _, rt := range []*relationtuple.InternalRelationTuple{
						{
							Namespace: "foo",
							Object:    "object",
							Relation:  "relation",
							Subject:   &relationtuple.SubjectID{ID: "user"},
						},
						{
							Namespace: "foo",
							Object:    "object",
							Relation:  "relation",
							Subject: &relationtuple.SubjectSet{
								Namespace: "foo",
								Object:    "s_object",
								Relation:  "s_relation",
							},
						},
					} {
						assert.Contains(t, actualRts, rt)
					}
				})
			})

			t.Run("suite=down", func(t *testing.T) {
				if debugOnDisk && db.Name == "sqlite" {
					t.SkipNow()
				}
				require.NoError(t, tm.Down(ctx, -1))
			})
		})
	}
}

func logMigrationStatus(t *testing.T, m *popx.MigrationBox) {
	t.Helper()

	status, err := m.Status(context.Background())
	require.NoError(t, err)
	s := strings.Builder{}
	_ = status.Write(&s)
	t.Log("Migration status:\n", s.String())
}
