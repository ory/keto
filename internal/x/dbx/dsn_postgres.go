// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

//go:build !nopostgres

package dbx

import (
	"testing"

	"github.com/ory/x/sqlcon/dockertest"
)

func RunPostgres(t testing.TB, testDB string) string {
	url := dockertest.RunTestPostgreSQL(t)
	if err := createDB(t, url, testDB); err != nil {
		t.Fatal(err)
	}
	return withDbName(url, testDB)
}
