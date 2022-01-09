package uuidmapping

import (
	"context"

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
