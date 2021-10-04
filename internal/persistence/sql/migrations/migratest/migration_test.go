package migratest

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/gobuffalo/pop/v5"

	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"

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
	"github.com/ory/keto/internal/persistence/sql"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x/dbx"
)

func TestMigrations(t *testing.T) {
	const debugOnDisk = false

	for _, db := range dbx.GetDSNs(t, debugOnDisk) {
		t.Run("dsn="+db.Name, func(t *testing.T) {
			db.MigrateUp, db.MigrateDown = false, false

			ctx := context.Background()
			l := logrusx.New("", "", logrusx.ForceLevel(logrus.DebugLevel))

			var c *pop.Connection
			var err error
			for i := 0; i < 120; i++ {
				c, err = pop.NewConnection(&pop.ConnectionDetails{URL: db.Conn})
				require.NoError(t, err)
				require.NoError(t, c.Open())
				if err := c.Store.(interface{ Ping() error }).Ping(); err == nil {
					break
				}
				time.Sleep(time.Second)
			}
			require.NoError(t, c.Store.(interface{ Ping() error }).Ping())

			tm := popx.NewTestMigrator(t, c, fsx.Merge(networkx.Migrations, os.DirFS("../sql")), os.DirFS("./testdata"), l)
			// cleanup first
			require.NoError(t, tm.Down(ctx, -1))

			t.Run("suite=up", func(t *testing.T) {
				require.NoError(t, tm.Up(ctx))
			})

			t.Run("suite=fixtures", func(t *testing.T) {
				t.Run("table=legacy namespaces", func(t *testing.T) {
					// as they are legacy, we expect them to be actually dropped

					assert.ErrorIs(t, sqlcon.HandleError(c.RawQuery("SELECT * FROM keto_namespace").Exec()), sqlcon.ErrNoSuchTable)
				})

				t.Run("table=relation tuples", func(t *testing.T) {
					reg := driver.NewTestRegistry(t, db)
					require.NoError(t,
						reg.Config().Set(config.KeyNamespaces, []*namespace.Namespace{{ID: 1, Name: "foo"}}))

					p, err := sql.NewPersister(reg, uuid.Must(uuid.FromString("77fdc5e0-2260-49da-8aae-c36ba255d05b")))
					require.NoError(t, err)
					rts, next, err := p.GetRelationTuples(context.Background(), &relationtuple.RelationQuery{Namespace: "foo"})
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
						assert.Contains(t, rts, rt)
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
