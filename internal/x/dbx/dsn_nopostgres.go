// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

//go:build nopostgres

package dbx

import (
	"testing"
)

func RunPostgres(testing.TB, string) string { return "" }
