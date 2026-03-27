// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package dbx

import (
	"testing"

	"github.com/ory/x/dbal"
)

func GetSqlite(t testing.TB, mode sqliteMode) *DsnT {
	dsn := &DsnT{
		MigrateUp:   true,
		MigrateDown: false,
	}

	if mode == SQLiteMemory {
		dsn.Name = "sqlite-memory"
		dsn.Conn = dbal.NewSQLiteTestDatabase(t) + "&mode=memory"
		return dsn
	}

	dsn.Name = "sqlite-file"
	dsn.Conn = dbal.NewSQLiteTestDatabase(t)
	return dsn
}

func allSqlite(t testing.TB, debugSqliteOnDisk bool) []*DsnT {
	sqliteMode := SQLiteFile
	if debugSqliteOnDisk {
		sqliteMode = SQLiteDebug
	}

	return []*DsnT{
		GetSqlite(t, sqliteMode),
		GetSqlite(t, SQLiteMemory),
	}
}
