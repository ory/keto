package sql_test

import (
	"context"
	stdSql "database/sql"
	"strings"
	"testing"
	"time"

	"github.com/ory/x/networkx"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/persistence/sql"
	"github.com/ory/keto/internal/x/dbx"
)

func rt(nw *networkx.Network, setSID, setNID, setO, setR bool) *sql.RelationTuple {
	return &sql.RelationTuple{
		ID:        uuid.Must(uuid.NewV4()),
		NetworkID: nw.ID,
		SubjectID: stdSql.NullString{
			Valid: setSID,
		},
		SubjectSetNamespaceID: stdSql.NullInt32{
			Valid: setNID,
		},
		SubjectSetObject: stdSql.NullString{
			Valid: setO,
		},
		SubjectSetRelation: stdSql.NullString{
			Valid: setR,
		},
		CommitTime: time.Now(),
	}
}

func TestRelationTupleSubjectTypeCheck(t *testing.T) {
	t.Parallel()

	for _, dsn := range dbx.GetDSNs(t, false) {
		dsn := dsn
		t.Run("dsn="+dsn.Name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			reg := driver.NewTestRegistry(t, dsn)
			c, err := reg.PopConnection(context.Background())
			require.NoError(t, err)
			nw, err := reg.DetermineNetwork(ctx)
			require.NoError(t, err)

			for _, tc := range []struct {
				desc                                string
				setSID, setNID, setO, setR, success bool
			}{
				{
					desc:    "all",
					setSID:  true,
					setNID:  true,
					setO:    true,
					setR:    true,
					success: false,
				},
				{
					desc:    "nothing",
					setSID:  false,
					setNID:  false,
					setO:    false,
					setR:    false,
					success: false,
				},
				{
					desc:    "subject set",
					setSID:  false,
					setNID:  true,
					setO:    true,
					setR:    true,
					success: true,
				},
				{
					desc:    "subject ID",
					setSID:  true,
					setNID:  false,
					setO:    false,
					setR:    false,
					success: true,
				},
				{
					desc:    "incomplete subject set",
					setSID:  false,
					setNID:  true,
					setO:    true,
					setR:    false,
					success: false,
				},
			} {
				tc := tc
				t.Run("case="+tc.desc, func(t *testing.T) {
					t.Parallel()
					err = c.Create(rt(nw, tc.setSID, tc.setNID, tc.setO, tc.setR))

					if tc.success {
						assert.NoError(t, err)
					} else {
						require.Error(t, err)
						assert.True(t,
							strings.Contains(err.Error(), "chk_keto_rt_subject_type") || // <- normal databases
								strings.Contains(err.Error(), "SQLSTATE 23514")) // <- mysql
					}
				})
			}
		})
	}
}
