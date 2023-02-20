// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package sql

import (
	"context"
	"strings"

	"golang.org/x/exp/maps"

	"github.com/gofrs/uuid"
	"github.com/ory/x/otelx"
	"github.com/ory/x/sqlcon"

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

func (p *Persister) batchToUUIDs(ctx context.Context, values []string, readOnly bool) (uuids []uuid.UUID, err error) {
	if len(values) == 0 {
		return
	}

	uuids = make([]uuid.UUID, len(values))
	placeholderArray := make([]string, len(values))
	args := make([]interface{}, 0, len(values)*2)
	for i, val := range values {
		uuids[i] = uuid.NewV5(p.NetworkID(ctx), val)
		placeholderArray[i] = "(?, ?)"
		args = append(args, uuids[i], val)
	}
	placeholders := strings.Join(placeholderArray, ", ")

	p.d.Logger().WithField("values", values).WithField("UUIDs", uuids).Trace("adding UUID mappings")

	if !readOnly {
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
	} else {
		return uuids, nil
	}
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
	uniqueIDs := maps.Keys(idIdx)

	res = make([]string, len(ids))

	for i := 0; i < len(uniqueIDs); i += pageSize {
		end := i + pageSize
		if end > len(uniqueIDs) {
			end = len(uniqueIDs)
		}
		idsToLookup := uniqueIDs[i:end]
		var mappings []UUIDMapping
		query := p.Connection(ctx).Where("id in (?)", idsToLookup)
		if err := sqlcon.HandleError(query.All(&mappings)); err != nil {
			return []string{}, err
		}

		// Write the representation to the correct index.
		for _, m := range mappings {
			for _, idx := range idIdx[m.ID] {
				res[idx] = m.StringRepresentation
			}
		}
	}

	return
}

func (p *Persister) MapStringsToUUIDs(ctx context.Context, s ...string) (_ []uuid.UUID, err error) {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.MapStringsToUUIDs")
	defer otelx.End(span, &err)

	return p.batchToUUIDs(ctx, s, false)
}

func (p *Persister) MapStringsToUUIDsReadOnly(ctx context.Context, s ...string) (_ []uuid.UUID, err error) {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.MapStringsToUUIDsReadOnly")
	defer otelx.End(span, &err)

	return p.batchToUUIDs(ctx, s, true)
}

func (p *Persister) MapUUIDsToStrings(ctx context.Context, u ...uuid.UUID) (_ []string, err error) {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.MapUUIDsToStrings")
	defer otelx.End(span, &err)

	return p.batchFromUUIDs(ctx, u)
}
