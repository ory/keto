// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

//go:build sqlite

package dbx

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func GetSqlite(t testing.TB, mode sqliteMode) *DsnT {
	dsn := &DsnT{
		MigrateUp:   true,
		MigrateDown: false,
	}

	fn := fmt.Sprintf("TestDB_%s_%d.sqlite", t.Name(), rand.Int31())
	// init switch
	switch mode {
	case SQLiteMemory:
		dsn.Name = "memory"
		dsn.Conn = fmt.Sprintf("sqlite://file:%s?_fk=true&cache=shared&mode=memory", fn)
		t.Cleanup(func() {
			_ = os.Remove(fn)
		})
	case SQLiteFile, SQLiteDebug:
		dsn.Name = "sqlite"
		dsn.Conn = fmt.Sprintf("sqlite://file:%s?_fk=true", fn)
	}
	// cleanup switch
	switch mode {
	case SQLiteMemory, SQLiteFile:
		t.Cleanup(func() {
			_ = os.Remove(fn)
		})
	}

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
