// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

//go:build !sqlite

package dbx

import (
	"testing"
)

func GetSqlite(t testing.TB, mode sqliteMode) *DsnT {
	t.Fatalf("use `-tags sqlite` to enable sqlite")
	return nil
}

func allSqlite(testing.TB, bool) []*DsnT {
	return nil
}
