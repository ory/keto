package x

import (
	"testing"

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
