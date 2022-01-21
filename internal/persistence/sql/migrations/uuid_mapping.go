package migrations

import (
	"context"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"

	"github.com/ory/keto/internal/persistence/sql"
	"github.com/ory/x/sqlcon"
)

type (
	toUUIDMappingMigrator struct {
		d       dependencies
		perPage int
	}
)

// NewToUUIDMappingMigrator creates a new UUID mapping migrator.
func NewToUUIDMappingMigrator(d dependencies) *toUUIDMappingMigrator {
	return &toUUIDMappingMigrator{d: d, perPage: 100}
}

// MigrateUUIDMappings migrates to UUID-mapped subject IDs for all relation
// tuples in the database.
func (m *toUUIDMappingMigrator) MigrateUUIDMappings(ctx context.Context) error {
	p, ok := m.d.Persister().(*sql.Persister)
	if !ok {
		panic("got unexpected persister")
	}

	return p.Transaction(ctx, func(ctx context.Context, c *pop.Connection) error {
		for page := 1; ; page++ {
			relationTuples, hasNext, err := m.getRelationTuples(ctx, page)
			if err != nil {
				return err
			}

			if err := p.Connection(ctx).All(&relationTuples); err != nil {
				return err
			}

			for _, rt := range relationTuples {
				_, err := uuid.FromString(rt.SubjectID.String)
				if err == nil || !rt.SubjectID.Valid || rt.SubjectID.String == "" {
					continue
				}

				id := uuid.Must(uuid.NewV4())
				if err := p.AddUUIDMapping(ctx, id, rt.SubjectID.String); err != nil {
					return err
				}

				rt.SubjectID.String = id.String()
				if err := c.Update(rt); err != nil {
					return err
				}
			}

			if !hasNext {
				break
			}
		}
		return nil
	})
}

func (m *toUUIDMappingMigrator) getRelationTuples(ctx context.Context, page int) (res []*sql.RelationTuple, hasNext bool, err error) {
	q := m.d.Persister().Connection(ctx).
		Order("nid, shard_id").
		Paginate(page, m.perPage)

	if err := q.All(&res); err != nil {
		return nil, false, sqlcon.HandleError(err)
	}
	return res, q.Paginator.Page < q.Paginator.TotalPages, nil
}
