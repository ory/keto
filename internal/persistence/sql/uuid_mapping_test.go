// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package sql_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/persistence/sql"
	"github.com/ory/keto/internal/x/dbx"
)

func assertCheckErr(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
	t.(*testing.T).Helper()
	if err == nil {
		return assert.Fail(t, "Did not receive an error", msgAndArgs...)
	}

	if strings.Contains(err.Error(), "keto_uuid_mappings") || // <- normal databases
		strings.Contains(err.Error(), "string_representation") || // <- sqlite
		strings.Contains(err.Error(), "SQLSTATE 23505") || // <- cockroach
		strings.Contains(err.Error(), "SQLSTATE 23514") { // <- mysql
		return true
	}
	return assert.Fail(t, fmt.Sprintf("Did not receive check error, got:\n%+v", err), msgAndArgs...)
}

func TestUUIDMapping(t *testing.T) {
	t.Parallel()

	for _, dsn := range dbx.GetDSNs(t, false) {
		dsn := dsn
		t.Run("dsn="+dsn.Name, func(t *testing.T) {
			t.Parallel()
			reg := driver.NewTestRegistry(t, dsn)
			c, err := reg.PopConnection(context.Background())
			require.NoError(t, err)

			testUUID := uuid.Must(uuid.NewV4())

			for _, tc := range []struct {
				desc      string
				mappings  interface{}
				assertErr assert.ErrorAssertionFunc
			}{{
				desc:      "empty should not fail on constraint",
				mappings:  &sql.UUIDMapping{},
				assertErr: assert.NoError,
			}, {
				desc:      "empty strings should not fail on constraint",
				mappings:  &sql.UUIDMapping{ID: uuid.Nil},
				assertErr: assert.NoError,
			}, {
				desc:      "single with string rep should succeed",
				mappings:  &sql.UUIDMapping{StringRepresentation: "foo"},
				assertErr: assert.NoError,
			}, {
				desc: "two with same uuid should fail on constraint",
				mappings: &[]sql.UUIDMapping{
					{ID: testUUID, StringRepresentation: "foo"},
					{ID: testUUID, StringRepresentation: "bar"},
				},
				assertErr: assertCheckErr,
			}, {
				desc: "two with same rep should succeed",
				mappings: &[]sql.UUIDMapping{
					{ID: uuid.Must(uuid.NewV4()), StringRepresentation: "bar"},
					{ID: uuid.Must(uuid.NewV4()), StringRepresentation: "bar"},
				},
				assertErr: assert.NoError,
			}} {
				tc := tc
				t.Run("case="+tc.desc, func(t *testing.T) {
					t.Parallel()
					err := c.Create(tc.mappings)
					tc.assertErr(t, err)
				})
			}
		})
	}
}
