//go:build nomysql

package dbx

import (
	"testing"
)

func RunMySQL(testing.TB, string) string { return "" }
