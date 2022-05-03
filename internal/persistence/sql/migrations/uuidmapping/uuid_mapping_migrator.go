package uuidmapping

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/ory/x/popx"
	"github.com/ory/x/sqlcon"
)

type (
	// We copy the definitions of RelationTuple and UUIDMapping here so that the
	// migration will always work on the same definitions.
	RelationTuple struct {
		// An ID field is required to make pop happy. The actual ID is a
		// composite primary key.
		ID                    uuid.UUID      `db:"shard_id"`
		NetworkID             uuid.UUID      `db:"nid"`
		NamespaceID           int32          `db:"namespace_id"`
		Object                string         `db:"object"`
		Relation              string         `db:"relation"`
		SubjectID             sql.NullString `db:"subject_id"`
		SubjectSetNamespaceID sql.NullInt32  `db:"subject_set_namespace_id"`
		SubjectSetObject      sql.NullString `db:"subject_set_object"`
		SubjectSetRelation    sql.NullString `db:"subject_set_relation"`
		CommitTime            time.Time      `db:"commit_time"`
	}
	UUIDMapping struct {
		ID                   uuid.UUID `db:"id"`
		StringRepresentation string    `db:"string_representation"`
	}
	UUIDMappings []*UUIDMapping
)

func (RelationTuple) TableName() string { return "keto_relation_tuples" }
func (UUIDMappings) TableName() string  { return "keto_uuid_mappings" }
func (UUIDMapping) TableName() string   { return "keto_uuid_mappings" }

var (
	name       = "migrate-strings-to-uuids"
	version    = "20220513210000000000"
	Migrations = popx.Migrations{
		{
			Version:   version,
			Name:      name,
			Path:      name,
			Direction: "up",
			DBType:    "all",
			Type:      "go",
			Runner: func(_ popx.Migration, conn *pop.Connection, _ *pop.Tx) error {
				for page := 1; ; page++ {
					relationTuples, hasNext, err := getRelationTuples(conn, page)
					if err != nil {
						return fmt.Errorf("could not get relation tuples: %w", err)
					}

					for _, rt := range relationTuples {
						if err = migrateSubjectID(conn, &rt); err != nil {
							return fmt.Errorf("could not migrate subject ID: %w", err)
						}
						if err = migrateSubjectSetObject(conn, &rt); err != nil {
							return fmt.Errorf("could not migrate subject set object: %w", err)
						}
						if err = migrateObject(conn, &rt); err != nil {
							return fmt.Errorf("could not migrate object: %w", err)
						}
						if err = conn.Update(&rt); err != nil {
							return fmt.Errorf("failed to update relation tuple: %w", err)
						}
					}

					if !hasNext {
						break
					}
				}

				return nil
			},
		},
		// We have to specify the "down" migration even if it is a no-op, since
		// the migrator will still manipulate the version table in the database.
		{
			Version:   version,
			Name:      name,
			Path:      name,
			Direction: "down",
			DBType:    "all",
			Type:      "go",
			Runner: func(_ popx.Migration, _ *pop.Connection, _ *pop.Tx) error {
				return nil
			},
		},
	}
)

func hasMapping(conn *pop.Connection, id string) (bool, error) {
	found, err := conn.Where("id = ?", id).Exists(&UUIDMapping{})
	if err != nil {
		return false, nil
	}
	return found, nil
}

func migrateSubjectID(conn *pop.Connection, rt *RelationTuple) error {
	if !rt.SubjectID.Valid || rt.SubjectID.String == "" {
		return nil
	}
	skip, err := hasMapping(conn, rt.SubjectID.String)
	if err != nil {
		return err
	}
	if skip {
		return nil
	}

	rt.SubjectID.String, err = addUUIDMapping(conn, rt.NetworkID, rt.SubjectID.String)
	return err
}

func migrateSubjectSetObject(conn *pop.Connection, rt *RelationTuple) error {
	if !rt.SubjectSetObject.Valid || rt.SubjectSetObject.String == "" {
		return nil
	}
	skip, err := hasMapping(conn, rt.SubjectSetObject.String)
	if err != nil {
		return err
	}
	if skip {
		return nil
	}

	rt.SubjectSetObject.String, err = addUUIDMapping(conn, rt.NetworkID, rt.SubjectSetObject.String)
	return err
}

func migrateObject(conn *pop.Connection, rt *RelationTuple) error {
	if rt.Object == "" {
		return nil
	}
	skip, err := hasMapping(conn, rt.Object)
	if err != nil {
		return err
	}
	if skip {
		return nil
	}

	rt.Object, err = addUUIDMapping(conn, rt.NetworkID, rt.Object)
	return err
}

func addUUIDMapping(conn *pop.Connection, networkID uuid.UUID, value string) (uid string, err error) {
	uid = uuid.NewV5(networkID, value).String()

	// We need to write manual SQL here because the INSERT should not fail if
	// the UUID already exists, but we still want to return an error if anything
	// else goes wrong.
	var query string
	switch d := conn.Dialect.Name(); d {
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

	err = sqlcon.HandleError(conn.RawQuery(query, uid, value).Exec())
	if err != nil {
		return "", fmt.Errorf("failed to add UUID mapping: %w", err)
	}
	return
}

func getRelationTuples(conn *pop.Connection, page int) (
	res []RelationTuple, hasNext bool, err error,
) {
	q := conn.Order("nid, shard_id").Paginate(page, 100)

	if err := q.All(&res); err != nil {
		return nil, false, sqlcon.HandleError(err)
	}
	return res, q.Paginator.Page < q.Paginator.TotalPages, nil
}
