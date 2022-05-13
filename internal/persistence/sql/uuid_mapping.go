package sql

import (
	"context"
	"fmt"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/ory/x/sqlcon"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
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

func (p *Persister) FromUUID(ctx context.Context, ids []uuid.UUID, opts ...x.PaginationOptionSetter) (res []string, err error) {
	p.d.Logger().Trace("looking up UUIDs")

	// We need to paginate on the ids, because we want to get the exact chunk of
	// string representations for the given ids.
	pagination, _ := internalPaginationFromOptions(opts...)
	pageSize := pagination.PerPage

	// Build a map from UUID -> indices in the result.
	idIdx := make(map[uuid.UUID][]int)
	for i, id := range ids {
		if ids, ok := idIdx[id]; ok {
			idIdx[id] = append(ids, i)
		} else {
			idIdx[id] = []int{i}
		}
	}

	res = make([]string, len(ids))

	for i := 0; i < len(ids); i += pageSize {
		end := i + pageSize
		if end > len(ids) {
			end = len(ids)
		}
		idsToLookup := ids[i:end]
		mappings := &[]UUIDMapping{}
		query := p.Connection(ctx).Where("id in (?)", idsToLookup)
		if err := sqlcon.HandleError(query.All(mappings)); err != nil {
			return []string{}, err
		}

		// Write the representation to the correct index.
		for _, m := range *mappings {
			for _, idx := range idIdx[m.ID] {
				res[idx] = m.StringRepresentation
			}
		}
	}

	return
}

func (p *Persister) replaceWithUUID(ctx context.Context, s *string) error {
	if s == nil {
		return nil
	}
	uuid, err := p.ToUUID(ctx, *s)
	if err != nil {
		return err
	}
	*s = uuid.String()

	return nil
}

func (p *Persister) MapFieldsToUUID(ctx context.Context, m relationtuple.UUIDMappable) error {
	return p.Transaction(ctx, func(ctx context.Context, _ *pop.Connection) error {
		for _, s := range m.UUIDMappableFields() {
			if s == nil || *s == "" {
				continue
			}
			if err := p.replaceWithUUID(ctx, s); err != nil {
				p.d.Logger().WithError(err).WithField("string", s).Error("got an error while mapping string to UUID")
				return err
			}
		}
		return nil
	})
}

func (p *Persister) MapFieldsFromUUID(ctx context.Context, m relationtuple.UUIDMappable) error {
	ids := make([]uuid.UUID, len(m.UUIDMappableFields()))
	for i, field := range m.UUIDMappableFields() {
		if field == nil {
			continue
		}
		id, err := uuid.FromString(*field)
		if err != nil {
			p.d.Logger().WithError(err).WithField("UUID", *field).Error("could not parse as UUID")
			return err
		}
		ids[i] = id
	}
	reps, err := p.FromUUID(ctx, ids)
	if err != nil {
		p.d.Logger().WithError(err).WithField("UUIDs", ids).Error("could fetch string mappings from DB")
		return err
	}
	for i, field := range m.UUIDMappableFields() {
		if field == nil {
			continue
		}
		if reps[i] == "" {
			p.d.Logger().WithError(err).WithField("string", reps[i]).Error("could not find the corresponding UUID")
			return fmt.Errorf("failed to map %s", ids[i])
		}
		*field = reps[i]
	}
	return nil
}
