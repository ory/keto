package migrations

import (
	"context"
	dbsql "database/sql"
	"testing"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/persistence/sql"
	"github.com/ory/keto/internal/x/dbx"
)

func TestToUUIDMappingMigrator(t *testing.T) {
	const debugOnDisk = false

	for _, dsn := range dbx.GetDSNs(t, debugOnDisk) {
		t.Run("db="+dsn.Name, func(t *testing.T) {
			ctx := context.Background()
			r := driver.NewTestRegistry(t, dsn)
			m := NewToUUIDMappingMigrator(r)
			p := m.d.Persister().(*sql.Persister)
			conn := p.Connection(ctx)

			testCases := []struct {
				name          string
				rt            *sql.RelationTuple
				expectMapping bool
			}{{
				name: "with string subject",
				rt: &sql.RelationTuple{
					ID:         uuid.Must(uuid.NewV4()),
					Object:     "object",
					SubjectID:  dbsql.NullString{String: "subject", Valid: true},
					CommitTime: time.Now(),
				},
				expectMapping: true,
			}, {
				name: "with null subject",
				rt: &sql.RelationTuple{
					ID:                    uuid.Must(uuid.NewV4()),
					SubjectID:             dbsql.NullString{String: "", Valid: false},
					SubjectSetNamespaceID: dbsql.NullInt32{Int32: 0, Valid: true},
					SubjectSetObject:      dbsql.NullString{String: "object", Valid: true},
					SubjectSetRelation:    dbsql.NullString{String: "subject_set_relation", Valid: true},
					Object:                "object",
					CommitTime:            time.Now(),
				},
				expectMapping: true,
			}, {
				name: "with UUID subject",
				rt: &sql.RelationTuple{
					ID:         uuid.Must(uuid.NewV4()),
					SubjectID:  dbsql.NullString{String: uuid.Must(uuid.NewV4()).String(), Valid: true},
					CommitTime: time.Now(),
				},
				expectMapping: false,
			}}

			for _, tc := range testCases {
				t.Run("case="+tc.name, func(t *testing.T) {
					require.NoError(t, conn.Create(tc.rt))
					require.NoError(t, m.MigrateUUIDMappings(ctx))

					newRt := &sql.RelationTuple{}
					require.NoError(t, conn.Find(newRt, tc.rt.ID))

					if tc.expectMapping {
						// Check subject mapping
						if tc.rt.SubjectID.Valid {
							assertHasMapping(t, conn, tc.rt.SubjectID.String, newRt.SubjectID.String)
						} else {
							assertHasMapping(t, conn, tc.rt.SubjectSetObject.String, newRt.SubjectSetObject.String)
							// check that both "OBJECT" strings are mapped to same UUID.
							assert.Equal(t, newRt.Object, newRt.SubjectSetObject.String)
						}
						assertHasMapping(t, conn, tc.rt.Object, newRt.Object)
					} else {
						// Nothing should have changed (ignoring commit time)
						newRt.CommitTime = tc.rt.CommitTime
						assert.Equal(t, tc.rt, newRt)
					}
				})
			}
		})
	}
}

// assertHasMapping checks that there is a mapping from the given string (value)
// to the given UUID (uid).
func assertHasMapping(t *testing.T, conn *pop.Connection, value, uid string) {
	t.Helper()
	mapping := &sql.UUIDMapping{}
	require.NoError(t, conn.Find(mapping, uid), "Could not find mapping for %q", uid)
	assert.NotEqual(t, value, uid, "value was not replaced by UUID")
	assert.Equal(t, value, mapping.StringRepresentation)
}
