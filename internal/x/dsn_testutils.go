package x

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/sjson"

	"github.com/ory/keto/internal/driver/config"

	"github.com/ory/x/sqlcon/dockertest"
)

type DsnT struct {
	Name string
	Conn string
}

func GetDSNs(t testing.TB) []*DsnT {
	// we use a slice of structs here to always have the same execution order
	dsns := []*DsnT{
		{
			Name: "memory",
			Conn: config.DSNMemory,
		},
	}
	if !testing.Short() {
		dsns = append(dsns,
			&DsnT{
				Name: "mysql",
				Conn: dockertest.RunTestMySQL(t),
			},
			&DsnT{
				Name: "postgres",
				Conn: dockertest.RunTestPostgreSQL(t),
			},
			&DsnT{
				Name: "cockroach",
				Conn: dockertest.RunTestCockroachDB(t),
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
