// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

//go:build !nocrdb && !nocockroach

package dbx

import (
	"testing"

	"github.com/ory/x/sqlcon/dockertest"
)

func RunCockroach(t testing.TB, testDB string) string {
	url := dockertest.RunTestCockroachDB(t)
	if err := createDB(t, url, testDB); err != nil {
		t.Fatal(err)
	}
	return withDbName(url, testDB)
}
