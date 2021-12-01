package dbx

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/sjson"

	"github.com/ory/x/sqlcon/dockertest"
)

type DsnT struct {
	Name                   string
	Conn                   string
	MigrateUp, MigrateDown bool
}

func GetDSNs(t testing.TB, debugSqliteOnDisk bool) []*DsnT {
	sqliteMode := SQLiteFile
	if debugSqliteOnDisk {
		sqliteMode = SQLiteDebug
	}

	// we use a slice of structs here to always have the same execution order
	dsns := []*DsnT{
		GetSqlite(t, sqliteMode),
		GetSqlite(t, SQLiteMemory),
	}

	if !testing.Short() {
		var mysql, postgres, cockroach string

		dockertest.Parallel([]func(){
			func() {
				mysql = dockertest.RunTestMySQL(t)
			},
			func() {
				postgres = dockertest.RunTestPostgreSQL(t)
			},
			func() {
				cockroach = dockertest.RunTestCockroachDB(t)
			},
		})

		dsns = append(dsns,
			&DsnT{
				Name:        "mysql",
				Conn:        mysql,
				MigrateUp:   true,
				MigrateDown: true,
			},
			&DsnT{
				Name:        "postgres",
				Conn:        postgres,
				MigrateUp:   true,
				MigrateDown: true,
			},
			&DsnT{
				Name:        "cockroach",
				Conn:        cockroach,
				MigrateUp:   true,
				MigrateDown: true,
			},
		)

		t.Cleanup(dockertest.KillAllTestDatabases)
	}

	return dsns
}

type sqliteMode = int

const (
	SQLiteMemory = iota
	SQLiteFile
	SQLiteDebug
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
			_ = os.Remove("TestDB.sqlite")
		})
		fallthrough
	case SQLiteDebug:
		dsn.Name = "sqlite"
		dsn.Conn = "sqlite://file:TestDB.sqlite?_fk=true"
	}

	return dsn
}

func ConfigFile(t testing.TB, values map[string]interface{}) string {
	dir := t.TempDir()
	fn := filepath.Join(dir, "keto.json")

	c := []byte("{}")
	for key, val := range values {
		var err error
		c, err = sjson.SetBytes(c, key, val)
		require.NoError(t, err)
	}

	require.NoError(t, ioutil.WriteFile(fn, c, 0600))

	return fn
}
