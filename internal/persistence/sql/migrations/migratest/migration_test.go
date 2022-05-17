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
	t.Parallel()
	const debugOnDisk = false

	for _, db := range dbx.GetDSNs(t, debugOnDisk) {
		db := db
		t.Run("dsn="+db.Name, func(t *testing.T) {
			t.Parallel()

			db.MigrateUp, db.MigrateDown = false, false

			ctx := context.Background()
			l := logrusx.New("", "", logrusx.ForceLevel(logrus.DebugLevel))

			var conn *pop.Connection
			var err error
			conn, err = pop.NewConnection(&pop.ConnectionDetails{URL: db.Conn})
			require.NoError(t, err)
			for i := 0; i < 120; i++ {
				require.NoError(t, conn.Open())
				if err := dbx.Ping(conn); err == nil {
					t.Cleanup(func() { conn.Close() })
					break
				}
				time.Sleep(time.Second)
			}
			require.NoError(t, dbx.Ping(conn))

			tm, err := popx.NewMigrationBox(
				fsx.Merge(sql.Migrations, networkx.Migrations),
				popx.NewMigrator(conn, l, nil, 1*time.Minute),
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

			t.Run("suite=up", func(t *testing.T) {
				if err := tm.Up(ctx); err != nil {
					t.Log("migrations failed:", err)
					t.Fail()
				}
				logMigrationStatus(t, tm)
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

			t.Run("suite=uuid_migrations", func(t *testing.T) {
				require.NoError(t, tm.Down(ctx, -1))

				// Migrate up to (including) "migrate-strings-to-uuids"
				migrateTo(t, tm, "migrate-strings-to-uuids")
				t.Log("status after up migration")
				logMigrationStatus(t, tm)

				// Assert that relationtuples have UUIDs
				tuples, _, err := p.GetRelationTuples(ctx, &relationtuple.RelationQuery{Namespace: "uuid_test"})
				require.NoError(t, err)
				assertIsUUID(t, *tuples[0].Subject.SubjectID())
				assertIsUUID(t, tuples[0].Object)

				// Migrate down to before "migrate-strings-to-uuids"
				require.NoError(t, tm.Down(ctx, 1))
				t.Log("status after down migration")
				logMigrationStatus(t, tm)

				// Assert that relationtuples have strings
				tuples, _, err = p.GetRelationTuples(ctx, &relationtuple.RelationQuery{Namespace: "uuid_test"})
				require.NoError(t, err)
				assert.Equal(t, "user", *tuples[0].Subject.SubjectID())
				assert.Equal(t, "object", tuples[0].Object)
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

func migrateTo(t *testing.T, tm *popx.MigrationBox, name string) {
	statuses, err := tm.Status(context.Background())
	require.NoError(t, err)

	for i, status := range statuses {
		if status.Name == name {
			_, err = tm.UpTo(context.Background(), i+1)
			require.NoError(t, err)
			return
		}
	}
	t.Fatal("could not find ", name)
}

func assertIsUUID(t *testing.T, id string) {
	t.Helper()
	if _, err := uuid.FromString(id); err != nil {
		t.Errorf("expected %q to be a UUID", id)
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
