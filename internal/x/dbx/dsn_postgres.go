// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

//go:build !nopostgres

package dbx

import (
	"testing"

	"github.com/ory/x/sqlcon/dockertest"
)

func RunPostgres(t testing.TB, testDB string) (string, error) {
	url := dockertest.RunTestPostgreSQLWithVersion(t, "16")
	if err := createDB(t, url, testDB); err != nil {
		return "", err
	}
	return withDbName(url, testDB), nil
}
