// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package sql

import (
	"bytes"
	"context"
	"iter"
	"maps"
	"slices"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/ory/x/otelx"
	"github.com/ory/x/sqlcon"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

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
	idIdx := make(map[uuid.UUID][]int, len(ids))
	for i, id := range ids {
		if ids, ok := idIdx[id]; ok {
			idIdx[id] = append(ids, i)
		} else {
			idIdx[id] = []int{i}
		}
	}
	nextID, stop := iter.Pull(maps.Keys(idIdx))
	defer stop()

	res = make([]string, len(ids))

	idsToLookup := make([]uuid.UUID, 0, pageSize)
	mappings := make([]UUIDMapping, 0, pageSize)
	for i := 0; i < len(idIdx); i += pageSize {
		idsToLookup = idsToLookup[:0]
		mappings = mappings[:0]

		for range pageSize {
			id, ok := nextID()
			if !ok {
				break
			}
			idsToLookup = append(idsToLookup, id)
		}
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

func (p *Persister) MapStringsToUUIDs(ctx context.Context, values ...string) (uuids []uuid.UUID, err error) {
	if len(values) == 0 {
		return
	}

	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.MapStringsToUUIDs",
		trace.WithAttributes(attribute.Int("num_values", len(values))))
	defer otelx.End(span, &err)

	uuids, err = p.MapStringsToUUIDsReadOnly(ctx, values...)
	if err != nil {
		return nil, err
	}

	p.d.Logger().WithField("values", values).WithField("UUIDs", uuids).Trace("adding UUID mappings")

	mappings := make([]UUIDMapping, len(values))
	for i := range values {
		mappings[i] = UUIDMapping{
			ID:                   uuids[i],
			StringRepresentation: values[i],
		}
	}
	slices.SortFunc(mappings, func(a, b UUIDMapping) int {
		return bytes.Compare(a.ID[:], b.ID[:])
	})
	mappings = slices.CompactFunc(mappings, func(a, b UUIDMapping) bool {
		return a.ID == b.ID
	})

	span.SetAttributes(attribute.Int("num_mappings", len(mappings)))

	err = p.Transaction(ctx, func(ctx context.Context) error {
		for chunk := range slices.Chunk(mappings, chunkSizeInsertUUIDMappings) {
			query, args := buildInsertUUIDs(chunk, p.conn.Dialect.Name())
			if err := p.Connection(ctx).RawQuery(query, args...).Exec(); err != nil {
				return sqlcon.HandleError(err)
			}
		}
		return nil
	})

	return uuids, err
}

func (p *Persister) MapStringsToUUIDsReadOnly(ctx context.Context, ss ...string) (uuids []uuid.UUID, err error) {
	// This function doesn't talk to the database or do anything interesting, so we don't need to trace it.
	// ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.MapStringsToUUIDsReadOnly")
	// defer otelx.End(span, &err)

	uuids = make([]uuid.UUID, len(ss))
	for i := range ss {
		uuids[i] = uuid.NewV5(p.NetworkID(ctx), ss[i])
	}
	return uuids, nil
}

func (p *Persister) MapUUIDsToStrings(ctx context.Context, u ...uuid.UUID) (_ []string, err error) {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.MapUUIDsToStrings")
	defer otelx.End(span, &err)

	return p.batchFromUUIDs(ctx, u)
}

func buildInsertUUIDs(values []UUIDMapping, dialect string) (query string, args []any) {
	if len(values) == 0 {
		return "", nil
	}

	const placeholder = "(?,?)"
	const separator = ","

	var q strings.Builder
	args = make([]any, 0, len(values)*2)

	if dialect == "mysql" {
		q.WriteString("INSERT IGNORE INTO keto_uuid_mappings (id, string_representation) VALUES ")
	} else {
		q.WriteString("INSERT INTO keto_uuid_mappings (id, string_representation) VALUES ")
	}

	q.Grow(len(values)*(len(placeholder)+len(separator)) + 100)

	for i, val := range values {
		if i > 0 {
			q.WriteString(separator)
		}
		q.WriteString(placeholder)
		args = append(args, val.ID, val.StringRepresentation)
	}

	if dialect == "mysql" {
		// nothing
	} else {
		q.WriteString(" ON CONFLICT (id) DO NOTHING")
	}

	return q.String(), args
}
