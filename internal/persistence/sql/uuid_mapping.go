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

func (p *Persister) ToUUID(ctx context.Context, text string) (uuid.UUID, error) {
	id := uuid.NewV5(p.NetworkID(ctx), text)
	p.d.Logger().Trace("adding UUID mapping")

	// We need to write manual SQL here because the INSERT should not fail if
	// the UUID already exists, but we still want to return an error if anything
	// else goes wrong.
	var query string
	switch d := p.Connection(ctx).Dialect.Name(); d {
	case "mysql":
		query = `
			INSERT IGNORE INTO keto_uuid_mappings (id, string_representation)
			VALUES (?, ?)`
	default:
		query = `
			INSERT INTO keto_uuid_mappings (id, string_representation)
			VALUES (?, ?)
			ON CONFLICT (id) DO NOTHING`
	}

	return id, sqlcon.HandleError(
		p.Connection(ctx).RawQuery(query, id, text).Exec(),
	)
}

func (p *Persister) FromUUID(ctx context.Context, id uuid.UUID) (rep string, err error) {
	p.d.Logger().Trace("looking up UUID")

	m := &UUIDMapping{}
	if err := sqlcon.HandleError(p.Connection(ctx).Find(m, id)); err != nil {
		return "", err
	}

	return m.StringRepresentation, nil
}
