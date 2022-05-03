package migratest

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"regexp"
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
	"github.com/ory/keto/internal/persistence/sql/migrations/uuidmapping"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x/dbx"
)

// TODO(hperl): move to ory/x
func withTestdata(t *testing.T, testdata fs.FS) func(*popx.MigrationBox) *popx.MigrationBox {
	return func(m *popx.MigrationBox) *popx.MigrationBox {
		err := fs.WalkDir(testdata, ".", func(path string, info fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			if m, _ := regexp.MatchString(`\d+_testdata.sql`, info.Name()); !m {
				return nil
			}
			version := strings.TrimSuffix(info.Name(), "_testdata.sql")
			m.Migrations["up"] = append(m.Migrations["up"], popx.Migration{
				Version:   version + "9", // run testdata after version
				Path:      path,
				Name:      "testdata",
				DBType:    "all",
				Direction: "up",
				Type:      "sql",
				Runner: func(m popx.Migration, _ *pop.Connection, tx *pop.Tx) error {
					b, err := fs.ReadFile(testdata, m.Path)
					if err != nil {
						return err
					}
					_, err = tx.Exec(string(b))
					return err
				},
			})
			m.Migrations["down"] = append(m.Migrations["down"], popx.Migration{
				Version:   version + "9", // run testdata after version
				Path:      path,
				Name:      "testdata",
				DBType:    "all",
				Direction: "down",
				Type:      "sql",
				Runner: func(m popx.Migration, _ *pop.Connection, tx *pop.Tx) error {
					return nil
				},
			})

			return nil
		})
		if err != nil {
			t.Fatalf("could not add all testdata migrations: %v", err)
		}

		return m
	}
}

func hasDownMigrationWithVersion(mb *popx.MigrationBox, version string) bool {
	for _, down := range mb.Migrations["down"] {
		if version == down.Version {
			return true
		}
	}
	return false
}

// check that every "up" migration has a corresponding "down" migration in
// reverse order.
// TODO(hperl): move to ory/x
func check(mb *popx.MigrationBox) error {
	for _, up := range mb.Migrations["up"] {
		if !hasDownMigrationWithVersion(mb, up.Version) {
			return fmt.Errorf("migration %s has no corresponding down migration", up.Version)
		}
	}
	return nil
}

func TestMigrations(t *testing.T) {
	const debugOnDisk = false

	for _, db := range dbx.GetDSNs(t, debugOnDisk) {
		t.Run("dsn="+db.Name, func(t *testing.T) {
			db.MigrateUp, db.MigrateDown = false, false

			ctx := context.Background()
			l := logrusx.New("", "", logrusx.ForceLevel(logrus.DebugLevel))

			var c *pop.Connection
			var err error
			c, err = pop.NewConnection(&pop.ConnectionDetails{URL: db.Conn})
			require.NoError(t, err)
			require.NoError(t, c.Open())
			t.Cleanup(func() { c.Close() })
			for i := 0; i < 120; i++ {
				if err := c.Store.(interface{ Ping() error }).Ping(); err == nil {
					break
				}
				time.Sleep(time.Second)
			}
			require.NoError(t, c.Store.(interface{ Ping() error }).Ping())

			tm, err := popx.NewMigrationBox(
				fsx.Merge(sql.Migrations, networkx.Migrations),
				popx.NewMigrator(c, l, nil, 1*time.Minute),
				popx.WithGoMigrations(uuidmapping.Migrations),
				withTestdata(t, os.DirFS("./testdata")),
			)
			require.NoError(t, err)
			if err := check(tm); err != nil {
				t.Log(err)
				t.Log("up migrations:")
				for _, m := range tm.Migrations["up"] {
					t.Logf("\t%s\t%s\t%s\n", m.Name, m.Version, m.DBType)
				}
				t.Log("down migrations:")
				for _, m := range tm.Migrations["down"] {
					t.Logf("\t%s\t%s\t%s\n", m.Name, m.Version, m.DBType)
				}
				t.FailNow()
			}

			// cleanup first
			require.NoError(t, tm.Down(ctx, -1))

			t.Log("before migration")
			logMigrationStatus(t, tm)

			t.Run("suite=up", func(t *testing.T) {
				if err := tm.Up(ctx); err != nil {
					t.Log("migrations failed:", err)
					t.Fail()
				}
				logMigrationStatus(t, tm)
			})

			t.Run("suite=fixtures", func(t *testing.T) {
				reg := driver.NewTestRegistry(t, db)
				require.NoError(t,
					reg.Config(ctx).Set(config.KeyNamespaces, []*namespace.Namespace{
						{ID: 1, Name: "foo"}}))
				p, err := sql.NewPersister(ctx, reg, uuid.Must(uuid.FromString("77fdc5e0-2260-49da-8aae-c36ba255d05b")))

				t.Run("table=legacy namespaces", func(t *testing.T) {
					// as they are legacy, we expect them to be actually dropped
					assert.ErrorIs(t, sqlcon.HandleError(c.RawQuery(
						"SELECT * FROM keto_namespace",
					).Exec()), sqlcon.ErrNoSuchTable)
				})

				t.Run("table=relation tuples", func(t *testing.T) {
					require.NoError(t, err)
					actualRts, next, err := p.GetRelationTuples(ctx, &relationtuple.RelationQuery{Namespace: "foo"})

					require.NoError(t, err)
					assert.Equal(t, "", next)
					t.Log("actual rts:", actualRts)

					expectedRts := []*relationtuple.InternalRelationTuple{
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
					}

					// The relationship tuples in the db have a UUID mapping, so
					// we need to convert our expectations to that.
					assert.NoError(t, p.MapFieldsToUUID(
						ctx, relationtuple.InternalRelationTuples(expectedRts)))
					assert.ElementsMatch(t, expectedRts, actualRts)
					logMigrationStatus(t, tm)
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
	status, err := m.Status(context.Background())
	require.NoError(t, err)
	s := strings.Builder{}
	_ = status.Write(&s)
	t.Log("Migration status:\n", s.String())
}
