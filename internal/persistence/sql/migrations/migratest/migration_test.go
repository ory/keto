// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package migratest

import (
	"context"
	stdSql "database/sql"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/ory/pop/v6"
	"github.com/ory/x/fsx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/networkx"
	keysetpagination "github.com/ory/x/pagination/keysetpagination_v2"
	"github.com/ory/x/pointerx"
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
	"github.com/ory/keto/ketoapi"
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

			namespaces := []*namespace.Namespace{
				{ID: 1, Name: "foo"},
				{ID: 2, Name: "uuid_test"},
			}
			nm := config.NewMemoryNamespaceManager(namespaces...)
			tm, err := popx.NewMigrationBox(
				fsx.Merge(sql.Migrations, networkx.Migrations),
				popx.NewMigrator(conn, logrusx.New("", "", logrusx.ForceLevel(logrus.DebugLevel)), nil, 1*time.Minute),
				popx.WithGoMigrations(uuidmapping.Migrations(nm)),
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

			reg := driver.NewTestRegistry(t, db, driver.WithNamespaces(namespaces))
			p, err := sql.NewPersister(ctx, reg, uuid.Must(uuid.FromString("77fdc5e0-2260-49da-8aae-c36ba255d05b")))
			require.NoError(t, err)

			t.Run("suite=fixtures", func(t *testing.T) {
				t.Run("table=legacy namespaces", func(t *testing.T) {
					// as they are legacy, we expect them to be actually dropped
					assert.ErrorIs(t, sqlcon.HandleError(conn.RawQuery(
						"SELECT * FROM keto_namespace",
					).Exec()), sqlcon.ErrNoSuchTable)
				})

				t.Run("table=relationships", func(t *testing.T) {
					actualRts, next, err := p.GetRelationTuples(ctx, &relationtuple.RelationQuery{Namespace: &namespaces[0].Name})
					require.NoError(t, err)
					assert.True(t, next.IsLast())
					t.Log("actual rts:", actualRts)

					expectedRts := []*ketoapi.RelationTuple{
						{
							Namespace: "foo",
							Object:    "object",
							Relation:  "relation",
							SubjectID: pointerx.Ptr("user"),
						},
						{
							Namespace: "foo",
							Object:    "object",
							Relation:  "relation",
							SubjectSet: &ketoapi.SubjectSet{
								Namespace: "foo",
								Object:    "s_object",
								Relation:  "s_relation",
							},
						},
					}

					// The relationship tuples in the db have a UUID mapping, so
					// we need to convert our expectations to that.
					expectedUUID, err := reg.Mapper().FromTuple(ctx, expectedRts...)
					require.NoError(t, err)
					assert.ElementsMatch(t, expectedUUID, actualRts)
					logMigrationStatus(t, tm)
				})
			})

			t.Run("suite=uuid_migrations", func(t *testing.T) {
				t.Run("correct types", func(t *testing.T) {
					ctx, cancel := context.WithTimeout(ctx, time.Minute)
					defer cancel()
					require.NoError(t, tm.Down(ctx, -1))

					// Migrate up to (including) "drop old non-uuid table"
					migrateUpTo(t, tm, "20220513200600000001")
					t.Log("status after up migration")
					logMigrationStatus(t, tm)

					// Assert that relationtuples have UUIDs
					tuples, _, err := p.GetRelationTuples(ctx, &relationtuple.RelationQuery{Namespace: &namespaces[1].Name})
					require.NoError(t, err)
					assert.NotZero(t, tuples[0].Subject.(*relationtuple.SubjectID).ID)
					assert.NotZero(t, tuples[0].Object)

					// Migrate down to before "migrate-strings-to-uuids"
					migrateDownTo(t, tm, "20220513200300000000")
					t.Log("status after down migration")
					logMigrationStatus(t, tm)

					// Assert that relationtuples have strings
					var oldRTs []*tuplesBeforeUUID
					require.NoError(t, p.Connection(ctx).
						Select("subject_id", "object").
						//lint:ignore SA1019 backwards compatibility
						//nolint:staticcheck
						Where("namespace_id = ?", namespaces[1].ID).
						All(&oldRTs))
					assert.Equalf(t, "user", oldRTs[0].SubjectID.String, "%+v", oldRTs[0])
					assert.Equal(t, "object", oldRTs[0].Object)
				})

				t.Run("paginates", func(t *testing.T) {
					ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
					defer cancel()
					require.NoError(t, tm.Down(ctx, -1))

					// migrate up to before all UUID migrations
					migrateUpTo(t, tm, "20220512151000000000")
					t.Log("status after up migration")
					logMigrationStatus(t, tm)

					oldRTs := make([]tuplesBeforeUUID, 2000)
					expected := make([]*ketoapi.RelationTuple, len(oldRTs))
					for i := range oldRTs {
						oldRTs[i] = tuplesBeforeUUID{
							NetworkID: p.NetworkID(ctx),
							//lint:ignore SA1019 backwards compatibility
							//nolint:staticcheck
							NamespaceID: namespaces[1].ID,
							Object:      "object-" + strconv.Itoa(i),
							Relation:    "pagination-works",
							SubjectID:   stdSql.NullString{String: "subject-" + strconv.Itoa(i), Valid: true},
							CommitTime:  time.Now(),
						}
						expected[i] = &ketoapi.RelationTuple{
							Namespace: namespaces[1].Name,
							Object:    oldRTs[i].Object,
							Relation:  "pagination-works",
							SubjectID: pointerx.Ptr(oldRTs[i].SubjectID.String),
						}
					}
					require.NoError(t, p.Connection(ctx).Create(oldRTs))
					require.NoError(t, tm.Up(ctx))

					newRTs := make([]*relationtuple.RelationTuple, 0, len(oldRTs))
					for nextPage := keysetpagination.NewPaginator(); !nextPage.IsLast(); {
						var rts []*relationtuple.RelationTuple
						rts, nextPage, err = p.GetRelationTuples(ctx, &relationtuple.RelationQuery{Relation: pointerx.Ptr("pagination-works")}, nextPage.ToOptions()...)
						require.NoError(t, err)
						newRTs = append(newRTs, rts...)
					}
					assert.Len(t, newRTs, len(oldRTs))
					actual, err := reg.Mapper().ToTuple(ctx, newRTs...)
					require.NoError(t, err)
					assert.ElementsMatch(t, expected, actual)
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

// migrateUpTo migrates up to the specified version (inclusive)
func migrateUpTo(t *testing.T, tm *popx.MigrationBox, version string) {
	statuses, err := tm.Status(context.Background())
	require.NoError(t, err)

	for i, status := range statuses {
		if status.Version == version {
			_, err = tm.UpTo(context.Background(), i+1)
			require.NoError(t, err)
			return
		}
	}
	t.Fatal("could not find ", version)
}

// migrateDownTo migrates down to the specified version (exclusive)
func migrateDownTo(t *testing.T, tm *popx.MigrationBox, version string) {
	statuses, err := tm.Status(context.Background())
	require.NoError(t, err)

	for i, status := range statuses {
		if status.Version == version {
			require.NoError(t, tm.Down(context.Background(), len(statuses)-i))
			return
		}
	}
	t.Fatal("could not find ", version)
}

func logMigrationStatus(t *testing.T, m *popx.MigrationBox) {
	t.Helper()

	status, err := m.Status(context.Background())
	require.NoError(t, err)
	s := strings.Builder{}
	_ = status.Write(&s)
	t.Log("Migration status:\n", s.String())
}

type tuplesBeforeUUID struct {
	ID                    uuid.UUID         `db:"shard_id"`
	NetworkID             uuid.UUID         `db:"nid"`
	NamespaceID           int32             `db:"namespace_id"`
	Object                string            `db:"object"`
	Relation              string            `db:"relation"`
	SubjectID             stdSql.NullString `db:"subject_id"`
	SubjectSetNamespaceID stdSql.NullInt32  `db:"subject_set_namespace_id"`
	SubjectSetObject      stdSql.NullString `db:"subject_set_object"`
	SubjectSetRelation    stdSql.NullString `db:"subject_set_relation"`
	CommitTime            time.Time         `db:"commit_time"`
}

func (tuplesBeforeUUID) TableName(_ context.Context) string {
	return "keto_relation_tuples"
}
