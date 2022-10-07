// Copyright © 2022 Ory Corp

//go:build nocrdb || nocockroach

package dbx

import (
	"testing"
)

func RunCockroach(t testing.TB, testDB string) string { return "" }
