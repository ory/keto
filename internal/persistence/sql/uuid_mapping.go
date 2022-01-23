package sql

import (
	"context"
	"errors"

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

func (p *Persister) MappedUUID(ctx context.Context, representation string) (uuid.UUID, error) {
	p.d.Logger().Trace("looking up mapped UUID")

	m := &UUIDMapping{}
	if err := sqlcon.HandleError(p.Connection(ctx).Where("string_representation = ?", representation).First(m)); err != nil {
		if errors.Is(err, sqlcon.ErrNoRows) {
			id := uuid.Must(uuid.NewV4())
			return id, p.AddUUIDMapping(ctx, id, representation)
		}
		return uuid.Nil, err
	}

	return m.ID, nil
}
