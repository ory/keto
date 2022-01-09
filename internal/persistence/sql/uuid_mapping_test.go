package sql_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/persistence/sql"
	"github.com/ory/keto/internal/x/dbx"
)

func TestUUIDMapping(t *testing.T) {
	for _, dsn := range dbx.GetDSNs(t, false) {
		t.Run("dsn="+dsn.Name, func(t *testing.T) {
			reg := driver.NewTestRegistry(t, dsn)
			c, err := reg.PopConnection()
			require.NoError(t, err)

			for _, tc := range []struct {
				desc      string
				mappings  interface{}
				shouldErr bool
			}{{
				desc:      "empty should fail on constraint",
				mappings:  &sql.UUIDMapping{},
				shouldErr: true,
			}, {
				desc:      "single with string rep should succeed",
				mappings:  &sql.UUIDMapping{StringRepresentation: "foo"},
				shouldErr: false,
			}, {
				desc: "two with same rep should fail on constraint",
				mappings: sql.UUIDMappings{
					&sql.UUIDMapping{StringRepresentation: "bar"},
					&sql.UUIDMapping{StringRepresentation: "bar"},
				},
				shouldErr: true,
			}} {
				t.Run("case="+tc.desc, func(t *testing.T) {
					err = c.Create(tc.mappings)
					if tc.shouldErr {
						assert.Error(t, err)
					} else {
						assert.NoError(t, err)
					}
				})
			}
		})
	}
}
