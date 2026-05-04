// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package dbx

import (
	"crypto/rand"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/sjson"
	"golang.org/x/sync/errgroup"

	"github.com/ory/pop/v6"
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

func dbName() string {
	var buf [20]byte
	_, _ = rand.Read(buf[:])
	return fmt.Sprintf("testdb_%x", buf)
}

func openAndPing(url string) (*pop.Connection, error) {
	conn, err := pop.NewConnection(&pop.ConnectionDetails{URL: url})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %q: %w", url, err)
	}
	for range 120 {
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
		_ = conn.Close()
		return nil, fmt.Errorf("failed to ping: %w", err)
	}
	return conn, nil
}

func createDB(t testing.TB, url string, dbName string) (err error) {
	var conn *pop.Connection

	if conn, err = openAndPing(url); err != nil {
		return fmt.Errorf("failed to connect to %q: %w", url, err)
	}
	t.Cleanup(func() {
		if err := conn.Close(); err != nil {
			t.Logf("could not close connection for %q in %q: %v", dbName, url, err)
		}
	})

	if err := conn.RawQuery("CREATE DATABASE " + dbName).Exec(); err != nil {
		return fmt.Errorf("failed to create db %q in %q: %w", dbName, url, err)
	}

	t.Cleanup(func() {
		if err := conn.RawQuery("DROP DATABASE " + dbName).Exec(); err != nil {
			t.Logf("could not drop database %q in %q: %v", dbName, url, err)
		}
	})

	return

}

func GetDSNs(t testing.TB) []*DsnT {
	dsns := []*DsnT{GetSqlite(t)}

	if !testing.Short() {
		var mysql, postgres, cockroach string
		testDB := dbName()

		var eg errgroup.Group
		eg.Go(func() (err error) {
			postgres, err = RunPostgres(t, testDB)
			return
		})
		eg.Go(func() (err error) {
			mysql, err = RunMySQL(t, testDB)
			return
		})
		eg.Go(func() (err error) {
			cockroach, err = RunCockroach(t, testDB)
			return
		})
		require.NoError(t, eg.Wait())

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
	}

	require.NotZero(t, len(dsns), "expected to run against at least one database")
	return dsns
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

	require.NoError(t, os.WriteFile(fn, c, 0600))

	return fn
}

type pinger interface{ Ping() error }

func Ping(conn *pop.Connection) error {
	p, ok := conn.Store.(pinger)
	if !ok {
		return fmt.Errorf("connection not opened: Store is %T", conn.Store)
	}
	return p.Ping()
}
