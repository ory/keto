package migrations

import (
	"context"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"

	"github.com/ory/x/sqlcon"

	"github.com/ory/keto/internal/persistence/sql"
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
				if err := m.migrateSubjectID(ctx, rt); err != nil {
					return err
				}
				if err := m.migrateSubjectSetObject(ctx, rt); err != nil {
					return err
				}
				if err := m.migrateObject(ctx, rt); err != nil {
					return err
				}
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

func (m *toUUIDMappingMigrator) migrateSubjectID(ctx context.Context, rt *sql.RelationTuple) error {
	_, err := uuid.FromString(rt.SubjectID.String)
	if err == nil || !rt.SubjectID.Valid || rt.SubjectID.String == "" {
		return nil
	}

	rt.SubjectID.String, err = m.addUUIDMapping(ctx, rt.SubjectID.String)
	return err
}

func (m *toUUIDMappingMigrator) migrateSubjectSetObject(ctx context.Context, rt *sql.RelationTuple) error {
	_, err := uuid.FromString(rt.SubjectSetObject.String)
	if err == nil || !rt.SubjectSetObject.Valid || rt.SubjectSetObject.String == "" {
		return nil
	}

	rt.SubjectSetObject.String, err = m.addUUIDMapping(ctx, rt.SubjectSetObject.String)
	return err
}

func (m *toUUIDMappingMigrator) migrateObject(ctx context.Context, rt *sql.RelationTuple) error {
	_, err := uuid.FromString(rt.Object)
	if err == nil || rt.Object == "" {
		return nil
	}

	rt.Object, err = m.addUUIDMapping(ctx, rt.Object)
	return err
}

func (m *toUUIDMappingMigrator) addUUIDMapping(ctx context.Context, value string) (id string, err error) {
	uid, err := m.d.Persister().MappedUUID(ctx, value)
	return uid.String(), err
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
