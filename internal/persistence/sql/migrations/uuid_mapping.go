package migrations

import (
	"context"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"

	"github.com/ory/keto/internal/persistence/sql"
)

type (
	toUUIDMappingMigrator struct {
		d dependencies
	}
)

// NewToUUIDMappingMigrator creates a new UUID mapping migrator.
func NewToUUIDMappingMigrator(d dependencies) *toUUIDMappingMigrator {
	return &toUUIDMappingMigrator{d: d}
}

// MigrateUUIDMappings migrates to UUID-mapped subject IDs for all relation
// tuples in the database.
func (m *toUUIDMappingMigrator) MigrateUUIDMappings(ctx context.Context) error {
	p, ok := m.d.Persister().(*sql.Persister)
	if !ok {
		panic("got unexpected persister")
	}

	return p.Transaction(ctx, func(ctx context.Context, c *pop.Connection) error {
		var relationTuples []*sql.RelationTuple

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
		return nil
	})
}
