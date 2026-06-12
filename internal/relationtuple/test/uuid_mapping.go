// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/relationtuple"
)

func MappingManagerTest(t *testing.T, m relationtuple.MappingManager) {
	t.Run("case=str -> uuid -> str", func(t *testing.T) {
		const s = "rep1"
		u, err := m.MapStringsToUUIDs(t.Context(), s)
		require.NoError(t, err)

		actual, err := m.MapUUIDsToStrings(t.Context(), u[0])
		require.NoError(t, err)
		require.Len(t, actual, 1)
		assert.Equal(t, s, actual[0])

		t.Run("case=batch", func(t *testing.T) {
			s := []string{"rep1", "rep2", "rep3"}

			u, err := m.MapStringsToUUIDs(t.Context(), s...)
			require.NoError(t, err)
			require.Len(t, u, len(s))

			assert.NotContains(t, u, uuid.Nil)

			actual, err := m.MapUUIDsToStrings(t.Context(), u...)
			require.NoError(t, err)
			require.Len(t, actual, len(s))
			assert.Equal(t, s, actual)
		})
	})

	t.Run("case=deterministic MapStringsToUUIDs", func(t *testing.T) {
		const s = "some string"

		u0, err := m.MapStringsToUUIDs(t.Context(), s)
		require.NoError(t, err)
		u1, err := m.MapStringsToUUIDs(t.Context(), s)
		require.NoError(t, err)

		assert.Equal(t, u0, u1)
	})
}
