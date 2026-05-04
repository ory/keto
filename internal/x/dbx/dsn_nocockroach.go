// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

//go:build nocrdb || nocockroach

package dbx

import (
	"testing"
)

func RunCockroach(testing.TB, string) (string, error) { return "", nil }
