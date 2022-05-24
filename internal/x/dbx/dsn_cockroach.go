//go:build !nocrdb && !nocockroach

package dbx

import (
	"github.com/ory/x/sqlcon/dockertest"
	"testing"
)

func RunCockroach(t testing.TB, testDB string) string {
	url := dockertest.RunTestCockroachDB(t)
	if err := createDB(t, url, testDB); err != nil {
		t.Fatal(err)
	}
	return withDbName(url, testDB)
}
