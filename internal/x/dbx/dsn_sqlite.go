//go:build sqlite

package dbx

import (
	"fmt"
	"os"
	"testing"
)

func GetSqlite(t testing.TB, mode sqliteMode) *DsnT {
	dsn := &DsnT{
		MigrateUp:   true,
		MigrateDown: false,
	}

	switch mode {
	case SQLiteMemory:
		dsn.Name = "memory"
		dsn.Conn = fmt.Sprintf("sqlite://file:%s?_fk=true&cache=shared&mode=memory", t.Name())
		t.Cleanup(func() {
			_ = os.Remove(t.Name())
		})
	case SQLiteFile:
		t.Cleanup(func() {
			_ = os.Remove(fmt.Sprintf("TestDB_%s.sqlite", t.Name()))
		})
		fallthrough
	case SQLiteDebug:
		dsn.Name = "sqlite"
		dsn.Conn = fmt.Sprintf("sqlite://file:TestDB_%s.sqlite?_fk=true", t.Name())
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
