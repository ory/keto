package uuidmapping

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gofrs/uuid"
)

type (
	ManagerProvider interface {
		UUIDMappingManager() Manager
	}
	Manager interface {
		// ToUUID returns the mapped UUID for the given string representation.
		// If the string representation is not mapped, a new UUID will be
		// created automatically.
		ToUUID(ctx context.Context, representation string) (uuid.UUID, error)

		// FromUUID returns the text representation for the given UUID.
		FromUUID(ctx context.Context, id uuid.UUID) (text string, err error)
	}
)

func ManagerTest(t *testing.T, m Manager) {
	ctx := context.Background()

	t.Run("case=ToUUID_FromUUID", func(t *testing.T) {
		rep1 := "foo"
		id, err := m.ToUUID(ctx, rep1)
		require.NoError(t, err)

		rep2, err := m.FromUUID(ctx, id)
		assert.NoError(t, err)
		assert.Equal(t, rep1, rep2)
	})

	t.Run("case=FromUUID", func(t *testing.T) {
		id1, err := m.ToUUID(ctx, "string")
		assert.NoError(t, err)
		id2, err := m.ToUUID(ctx, "string")
		assert.NoError(t, err)
		assert.Equal(t, id1, id2)
	})
}
