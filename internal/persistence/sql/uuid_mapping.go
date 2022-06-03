package sql

import (
	"context"
	"fmt"
	"strings"

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

func (p *Persister) batchToUUIDs(ctx context.Context, values []string) (uuids []uuid.UUID, err error) {
	if len(values) == 0 {
		return
	}

	uuids = make([]uuid.UUID, len(values))
	placeholderArray := make([]string, len(values))
	args := make([]interface{}, 0, len(values)*2)
	for i, val := range values {
		uuids[i] = uuid.NewV5(p.NetworkID(ctx), val)
		placeholderArray[i] = "(?, ?)"
		args = append(args, uuids[i].String(), val)
	}
	placeholders := strings.Join(placeholderArray, ", ")

	p.d.Logger().WithField("values", values).WithField("UUIDs", uuids).Trace("adding UUID mappings")

	// We need to write manual SQL here because the INSERT should not fail if
	// the UUID already exists, but we still want to return an error if anything
	// else goes wrong.
	var query string
	switch d := p.Connection(ctx).Dialect.Name(); d {
	case "mysql":
		query = `
			INSERT IGNORE INTO keto_uuid_mappings (id, string_representation) VALUES ` + placeholders
	default:
		query = `
			INSERT INTO keto_uuid_mappings (id, string_representation)
			VALUES ` + placeholders + `
			ON CONFLICT (id) DO NOTHING`
	}

	return uuids, sqlcon.HandleError(
		p.Connection(ctx).RawQuery(query, args...).Exec(),
	)
}

func (p *Persister) batchFromUUIDs(ctx context.Context, ids []uuid.UUID, opts ...x.PaginationOptionSetter) (res []string, err error) {
	if len(ids) == 0 {
		return
	}

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

func filterFields(fields []*string) []*string {
	res := make([]*string, 0, len(fields))
	for _, field := range fields {
		if field != nil && *field != "" {
			res = append(res, field)
		}
	}
	return res
}

func (p *Persister) MapStringsToUUIDs(ctx context.Context, s ...string) ([]uuid.UUID, error) {
	return p.batchToUUIDs(ctx, s)
}

func (p *Persister) MapUUIDsToStrings(ctx context.Context, u ...uuid.UUID) ([]string, error) {
	return p.batchFromUUIDs(ctx, u)
}

func (p *Persister) MapFieldsToUUID(ctx context.Context, m relationtuple.UUIDMappable) error {
	fields := filterFields(m.UUIDMappableFields())
	values := make([]string, len(fields))

	for i, field := range fields {
		values[i] = *field
	}
	ids, err := p.batchToUUIDs(ctx, values)
	if err != nil {
		p.d.Logger().WithError(err).WithField("values", values).Error("could insert UUID mappings")
		return err
	}
	for i, field := range fields {
		*field = ids[i].String()
	}
	return nil
}

func (p *Persister) MapFieldsFromUUID(ctx context.Context, m relationtuple.UUIDMappable) error {
	fields := filterFields(m.UUIDMappableFields())
	ids := make([]uuid.UUID, len(fields))
	for i, field := range fields {
		id, err := uuid.FromString(*field)
		if err != nil {
			p.d.Logger().WithError(err).WithField("UUID", *field).Error("could not parse as UUID")
			return err
		}
		ids[i] = id
	}
	reps, err := p.batchFromUUIDs(ctx, ids)
	if err != nil {
		p.d.Logger().WithError(err).WithField("UUIDs", ids).Error("could fetch string mappings from DB")
		return err
	}
	for i, field := range fields {
		if reps[i] == "" {
			p.d.Logger().WithError(err).WithField("string", reps[i]).Error("could not find the corresponding UUID")
			return fmt.Errorf("failed to map %s", ids[i])
		}
		*field = reps[i]
	}
	return nil
}
