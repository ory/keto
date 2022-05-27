package dbx

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"net/url"
	"path/filepath"
	"strings"
	"testing"
	"time"

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

func openAndPing(url string) (conn *pop.Connection, err error) {
	if conn, err = pop.NewConnection(&pop.ConnectionDetails{URL: url}); err != nil {
		return nil, fmt.Errorf("failed to connect to %q: %w", url, err)
	}
	for i := 0; i < 120; i++ {
		fmt.Println("trying to open connection to", url)
		if err := conn.Open(); err != nil {
			// return nil, fmt.Errorf("failed to open connection: %w", err)
			time.Sleep(1 * time.Second)
			continue
		}
		if err := Ping(conn); err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if err := Ping(conn); err != nil {
		return nil, fmt.Errorf("failed to ping: %w", err)
	}
	return conn, nil
}

func createDB(t testing.TB, url string, dbName string) (err error) {
	var conn *pop.Connection

	if conn, err = openAndPing(url); err != nil {
		return fmt.Errorf("failed to connect to %q: %w", url, err)
	}

	if err := conn.RawQuery("CREATE DATABASE " + dbName).Exec(); err != nil {
		return fmt.Errorf("failed to create db %q in %q: %w", dbName, url, err)
	}

	t.Cleanup(func() {
		if err := conn.RawQuery("DROP DATABASE " + dbName).Exec(); err != nil {
			t.Logf("could not drop database %q in %q: %v", dbName, url, err)
		}
		conn.Close()
	})

	return

}

func GetDSNs(t testing.TB, debugSqliteOnDisk bool) []*DsnT {
	dsns := allSqlite(t, debugSqliteOnDisk)

	if !testing.Short() {
		var mysql, postgres, cockroach string
		testDB := dbName(t.Name())

		dockertest.Parallel([]func(){
			func() {
				mysql = RunMySQL(t, testDB)
			},
			func() {
				postgres = RunPostgres(t, testDB)
			},
			func() {
				cockroach = RunCockroach(t, testDB)
			},
		})

		if mysql != "" {
			dsns = append(dsns, &DsnT{
				Name:        "mysql",
				Conn:        mysql,
				MigrateUp:   true,
				MigrateDown: true,
			})
		}
		if postgres != "" {
			dsns = append(dsns, &DsnT{
				Name:        "postgres",
				Conn:        postgres,
				MigrateUp:   true,
				MigrateDown: true,
			})
		}
		if cockroach != "" {
			dsns = append(dsns, &DsnT{
				Name:        "cockroach",
				Conn:        cockroach,
				MigrateUp:   true,
				MigrateDown: true,
			})
		}

		t.Cleanup(dockertest.KillAllTestDatabases)
	}

	require.NotZero(t, len(dsns), "expected to run against at least one database")
	return dsns
}

type sqliteMode = int

const (
	SQLiteMemory = iota
	SQLiteFile
	SQLiteDebug
)

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

type pinger interface{ Ping() error }

func Ping(conn *pop.Connection) error { return conn.Store.(pinger).Ping() }
