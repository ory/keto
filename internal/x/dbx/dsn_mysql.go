//go:build !nomysql

package dbx

import (
	"github.com/ory/x/sqlcon/dockertest"
	"testing"
	"time"
)

func RunMySQL(t testing.TB, testDB string) string {
	url := dockertest.RunTestMySQL(t)
	time.Sleep(1 * time.Second)
	if err := createDB(t, url, testDB); err != nil {
		t.Fatal(err)
	}
	return withDbName(url, testDB)
}
