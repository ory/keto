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
		// MappedUUID returns the mapped UUID for the given string
		// representation. If the string representation is not mapped, a new
		// UUID will be created automatically.
		MappedUUID(ctx context.Context, representation string) (uuid.UUID, error)

		// AddUUIDMapping adds a new mapping between a UUID and a string.
		AddUUIDMapping(ctx context.Context, id uuid.UUID, representation string) error

		// LookupUUID returns the string representation for the given UUID.
		LookupUUID(ctx context.Context, id uuid.UUID) (rep string, err error)
	}
)

func ManagerTest(t *testing.T, m Manager) {
	ctx := context.Background()

	t.Run("case=add_lookup", func(t *testing.T) {
		id := uuid.Must(uuid.NewV4())
		rep1 := "foo"
		require.NoError(t, m.AddUUIDMapping(ctx, id, rep1))

		rep2, err := m.LookupUUID(ctx, id)
		assert.NoError(t, err)
		assert.Equal(t, rep1, rep2)
	})

	t.Run("case=MappedUUID", func(t *testing.T) {
		id1, err := m.MappedUUID(ctx, "string")
		require.NoError(t, err)
		id2, err := m.MappedUUID(ctx, "string")
		require.NoError(t, err)
		assert.Equal(t, id1, id2)
	})
}
