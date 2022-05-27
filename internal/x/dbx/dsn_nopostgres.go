//go:build nopostgres

package dbx

import (
	"testing"
)

func RunPostgres(testing.TB, string) string { return "" }
