// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package namespace_test

import (
	"bytes"
	"testing"

	"github.com/ory/x/snapshotx"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/cmd/namespace"
)

func TestGenerateOPLConfig(t *testing.T) {
	cases := []struct {
		name       string
		namespaces []string
	}{{
		name:       "empty",
		namespaces: []string{},
	}, {
		name:       "one",
		namespaces: []string{"one"},
	}, {
		name:       "many",
		namespaces: []string{"one", "two", "three"},
	}}

	for _, tc := range cases {
		t.Run("case="+tc.name, func(t *testing.T) {
			var out bytes.Buffer
			require.NoError(t, namespace.GenerateOPLConfig(tc.namespaces, &out))
			snapshotx.SnapshotT(t, out.String())
		})
	}
}
