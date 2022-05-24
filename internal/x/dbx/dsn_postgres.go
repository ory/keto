//go:build !nopostgres

package dbx

import (
	"github.com/ory/x/sqlcon/dockertest"
	"testing"
)

func RunPostgres(t testing.TB, testDB string) string {
	url := dockertest.RunTestPostgreSQL(t)
	if err := createDB(t, url, testDB); err != nil {
		t.Fatal(err)
	}
	return withDbName(url, testDB)
}
