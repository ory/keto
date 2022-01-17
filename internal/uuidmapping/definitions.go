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
		AddUUIDMapping(ctx context.Context, id uuid.UUID, representation string) error
		LookupUUID(ctx context.Context, id uuid.UUID) (rep string, err error)
	}
)

func ManagerTest(t *testing.T, m Manager) {
	ctx := context.Background()

	id := uuid.Must(uuid.NewV4())
	rep1 := "foo"
	require.NoError(t, m.AddUUIDMapping(ctx, id, rep1))

	rep2, err := m.LookupUUID(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, rep1, rep2)
}
