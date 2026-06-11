// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package dbx

import (
	"testing"

	"github.com/ory/x/dbal"
)

func GetSqlite(t testing.TB) *DsnT {
	return &DsnT{
		MigrateUp:   true,
		MigrateDown: false,
		Conn:        dbal.NewSQLiteTestDatabase(t),
		Name:        "sqlite",
	}
}
