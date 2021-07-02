package dbx

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/ory/keto/internal/driver"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/sjson"

	"github.com/ory/x/sqlcon/dockertest"
)

type DsnT struct {
	Name string
	Conn string
}

func GetDSNs(t testing.TB, debugSqliteOnDisk bool) []*DsnT {
	// we use a slice of structs here to always have the same execution order
	dsns := []*DsnT{
		{
			Name: "memory",
			Conn: driver.SqliteTestDSN(t, debugSqliteOnDisk),
		},
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
				Name: "mysql",
				Conn: mysql,
			},
			&DsnT{
				Name: "postgres",
				Conn: postgres,
			},
			&DsnT{
				Name: "cockroach",
				Conn: cockroach,
			},
		)

		t.Cleanup(dockertest.KillAllTestDatabases)
	}

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

	require.NoError(t, ioutil.WriteFile(fn, c, 0600))

	return fn
}
