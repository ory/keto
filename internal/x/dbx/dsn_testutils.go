package dbx

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/gobuffalo/pop/v6"
	"github.com/ory/x/sqlcon/dockertest"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/sjson"
)

type DsnT struct {
	Name                   string
	Conn                   string
	MigrateUp, MigrateDown bool
}

const mySQLSchema = "mysql://"

func mySQLWithDbName(dsn string, db string) string {
	cfg, err := mysql.ParseDSN(strings.TrimPrefix(dsn, mySQLSchema))
	if err != nil {
		return ""
	}
	cfg.DBName = db
	return mySQLSchema + cfg.FormatDSN()
}

func withDbName(dsn string, db string) string {
	// Special case for mysql, because their URLs are not parsable.
	if strings.HasPrefix(dsn, mySQLSchema) {
		return mySQLWithDbName(dsn, db)
	}

	u, err := url.Parse(dsn)
	if err != nil {
		return ""
	}
	u.Path = db

	return u.String()
}

// dbName returns a name for the database based on the test name.
func dbName(_ string) string {
	var buf [20]byte
	_, _ = rand.Read(buf[:])
	return fmt.Sprintf("testdb_%x", buf)
}

func createDB(t testing.TB, dsn string) (err error) {
	var conn *pop.Connection

	if conn, err = pop.NewConnection(&pop.ConnectionDetails{URL: dsn}); err != nil {
		return fmt.Errorf("failed to connect to %q: %w", dsn, err)
	}
	if err = pop.CreateDB(conn); err != nil {
		return fmt.Errorf("failed to create db in %q: %w", dsn, err)
	}
	t.Cleanup(func() {
		if err = pop.DropDB(conn); err != nil {
			t.Log(err)
		}
	})

	return
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
		db := dbName(t.Name())

		dockertest.Parallel([]func(){
			func() {
				mysql = withDbName(dockertest.RunTestMySQL(t), db)
				if err := createDB(t, mysql); err != nil {
					t.Fatal(err)
				}
			},
			func() {
				postgres = withDbName(dockertest.RunTestPostgreSQL(t), db)
				if err := createDB(t, postgres); err != nil {
					t.Fatal(err)
				}
			},
			func() {
				cockroach = withDbName(dockertest.RunTestCockroachDB(t), db)
				if err := createDB(t, cockroach); err != nil {
					t.Fatal(err)
				}
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
