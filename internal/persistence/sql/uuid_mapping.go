package sql

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/ory/x/sqlcon"
)

type (
	UUIDMapping struct {
		ID                   uuid.UUID `db:"id"`
		StringRepresentation string    `db:"string_representation"`
	}
	UUIDMappings []*UUIDMapping
)

func (UUIDMappings) TableName() string {
	return "keto_uuid_mappings"
}

func (UUIDMapping) TableName() string {
	return "keto_uuid_mappings"
}

func (p *Persister) AddUUIDMapping(ctx context.Context, id uuid.UUID, representation string) error {
	m := &UUIDMapping{
		ID:                   id,
		StringRepresentation: representation,
	}
	p.d.Logger().Trace("adding UUID mapping")

	return sqlcon.HandleError(p.Connection(ctx).Create(m))
}

func (p *Persister) LookupUUID(ctx context.Context, id uuid.UUID) (rep string, err error) {
	p.d.Logger().Trace("looking up UUID")

	m := &UUIDMapping{}
	if err := sqlcon.HandleError(p.Connection(ctx).Find(m, id)); err != nil {
		return "", err
	}

	return m.StringRepresentation, nil
}
