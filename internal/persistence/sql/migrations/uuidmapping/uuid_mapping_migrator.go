package uuidmapping

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/ory/x/popx"
	"github.com/ory/x/sqlcon"
	"golang.org/x/exp/maps"
)

// We copy the definitions of RelationTuple and UUIDMapping here so that the
// migration will always work on the same definitions.
type (
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
	NewRelationTuple struct {
		ID                    uuid.UUID      `db:"shard_id"`
		NetworkID             uuid.UUID      `db:"nid"`
		NamespaceID           int32          `db:"namespace_id"`
		Object                uuid.UUID      `db:"object"`
		Relation              string         `db:"relation"`
		SubjectID             uuid.NullUUID  `db:"subject_id"`
		SubjectSetNamespaceID sql.NullInt32  `db:"subject_set_namespace_id"`
		SubjectSetObject      uuid.NullUUID  `db:"subject_set_object"`
		SubjectSetRelation    sql.NullString `db:"subject_set_relation"`
		CommitTime            time.Time      `db:"commit_time"`
	}
	UUIDMapping struct {
		ID                   uuid.UUID `db:"id"`
		StringRepresentation string    `db:"string_representation"`
	}
	UUIDMappings   []*UUIDMapping
	ColumnProvider interface{ dbCols() []any }
)

const (
	RelationTupleTableName     = "keto_relation_tuples"
	RelationTupleUUIDTableName = "keto_relation_tuples_uuid"
	UUIDMappingTableName       = "keto_uuid_mappings"
	MigrationVersion           = "20220513200500000000"
)

func (RelationTuple) TableName() string    { return RelationTupleTableName }
func (NewRelationTuple) TableName() string { return RelationTupleUUIDTableName }

func (rt *RelationTuple) dbCols() []any {
	return []any{rt.ID, rt.NetworkID, rt.NamespaceID, rt.Object, rt.Relation, rt.SubjectID, rt.SubjectSetNamespaceID, rt.SubjectSetObject, rt.SubjectSetRelation, rt.CommitTime}
}

func (rt *NewRelationTuple) dbCols() []any {
	return []any{rt.ID, rt.NetworkID, rt.NamespaceID, rt.Object, rt.Relation, rt.SubjectID, rt.SubjectSetNamespaceID, rt.SubjectSetObject, rt.SubjectSetRelation, rt.CommitTime}
}

func (UUIDMappings) TableName() string { return UUIDMappingTableName }
func (UUIDMapping) TableName() string  { return UUIDMappingTableName }

func (m *UUIDMapping) dbCols() []any {
	return []any{m.ID, m.StringRepresentation}
}

func (rt *RelationTuple) ToUUID(s string) uuid.UUID {
	return uuid.NewV5(rt.NetworkID, s)
}

func (rt *RelationTuple) ToNew() (newRT *NewRelationTuple, objectMapping *UUIDMapping, subjectMapping *UUIDMapping) {
	newRT = &NewRelationTuple{
		ID:          rt.ID,
		NetworkID:   rt.NetworkID,
		NamespaceID: rt.NamespaceID,
		Object:      uuid.NewV5(rt.NetworkID, rt.Object),
		Relation:    rt.Relation,
		SubjectID: uuid.NullUUID{
			Valid: rt.SubjectID.Valid,
			UUID:  uuid.NewV5(rt.NetworkID, rt.SubjectID.String),
		},
		SubjectSetNamespaceID: rt.SubjectSetNamespaceID,
		SubjectSetObject: uuid.NullUUID{
			Valid: rt.SubjectSetObject.Valid,
			UUID:  uuid.NewV5(rt.NetworkID, rt.SubjectSetObject.String),
		},
		SubjectSetRelation: rt.SubjectSetRelation,
		CommitTime:         rt.CommitTime,
	}
	objectMapping = &UUIDMapping{
		ID:                   newRT.Object,
		StringRepresentation: rt.Object,
	}
	switch {
	case rt.SubjectID.Valid:
		subjectMapping = &UUIDMapping{
			ID:                   newRT.SubjectID.UUID,
			StringRepresentation: rt.SubjectID.String,
		}
	case rt.SubjectSetObject.Valid:
		subjectMapping = &UUIDMapping{
			ID:                   newRT.SubjectSetObject.UUID,
			StringRepresentation: rt.SubjectSetObject.String,
		}
	}
	return
}

var (
	name       = "migrate-strings-to-uuids"
	Migrations = popx.Migrations{
		// The "up" migration will add the UUID mappings to the database and
		// replace the strings with UUIDs.
		{
			Version:   MigrationVersion,
			Name:      name,
			Path:      name,
			Direction: "up",
			DBType:    "all",
			Type:      "go",
			Runner: func(_ popx.Migration, conn *pop.Connection, _ *pop.Tx) error {
				for lastID := uuid.Nil; ; {
					relationTuples, hasNext, err := GetRelationTuples[RelationTuple](conn, lastID)
					if err != nil {
						return fmt.Errorf("could not get relation tuples: %w", err)
					}

					mappings := make([]*UUIDMapping, len(relationTuples)*2)
					newTuples := make([]*NewRelationTuple, len(relationTuples))
					for i := range relationTuples {
						newTuples[i], mappings[i*2], mappings[i*2+1] = relationTuples[i].ToNew()
					}

					if err := BatchWriteMappings(conn, mappings); err != nil {
						return fmt.Errorf("could not write mappings: %w", err)
					}
					if err := BatchInsertTuples(conn, newTuples); err != nil {
						return fmt.Errorf("could not insert new tuples: %w", err)
					}
					if !hasNext {
						break
					}
					lastID = relationTuples[len(relationTuples)-1].ID
				}

				return nil
			},
		},
		// The "down" migration will replace all UUIDs with strings from the
		// mapping table.
		{
			Version:   MigrationVersion,
			Name:      name,
			Path:      name,
			Direction: "down",
			DBType:    "all",
			Type:      "go",
			Runner: func(_ popx.Migration, conn *pop.Connection, _ *pop.Tx) error {
				for lastID := uuid.Nil; ; {
					relationTuples, hasNext, err := GetRelationTuples[NewRelationTuple](conn, lastID)
					if err != nil {
						return fmt.Errorf("could not get relation tuples: %w", err)
					}

					mappings := make(map[uuid.UUID][]*string, len(relationTuples)*2)
					oldTuples := make([]*RelationTuple, len(relationTuples))
					for i, rt := range relationTuples {
						ot := &RelationTuple{
							ID:          rt.ID,
							NetworkID:   rt.NetworkID,
							NamespaceID: rt.NamespaceID,
							Relation:    rt.Relation,
							SubjectID: sql.NullString{
								Valid: rt.SubjectID.Valid,
							},
							SubjectSetNamespaceID: rt.SubjectSetNamespaceID,
							SubjectSetObject: sql.NullString{
								Valid: rt.SubjectSetObject.Valid,
							},
							SubjectSetRelation: rt.SubjectSetRelation,
							CommitTime:         rt.CommitTime,
						}
						mappings[rt.Object] = append(mappings[rt.Object], &ot.Object)
						switch {
						case rt.SubjectID.Valid:
							mappings[rt.SubjectID.UUID] = append(mappings[rt.SubjectID.UUID], &ot.SubjectID.String)
						case rt.SubjectSetObject.Valid:
							mappings[rt.SubjectSetObject.UUID] = append(mappings[rt.SubjectSetObject.UUID], &ot.SubjectSetObject.String)
						}
						oldTuples[i] = ot
					}
					if err := BatchReplaceUUIDs(conn, mappings); err != nil {
						return fmt.Errorf("could not replace UUIDs: %w", err)
					}

					if err := BatchInsertTuples(conn, oldTuples); err != nil {
						return fmt.Errorf("could not insert old tuples: %w", err)
					}

					if !hasNext {
						break
					}
					lastID = relationTuples[len(relationTuples)-1].ID
				}

				return nil
			},
		},
	}
)

func ConstructArgs[T ColumnProvider](nCols int, items []T) (string, []interface{}) {
	placeholderRow := "(" + strings.Repeat("?,", nCols-1) + "?)"

	q := &strings.Builder{}
	q.Grow(len(items) * (len(placeholderRow) + 1))

	args := make([]interface{}, 0, len(items)*nCols)

	q.WriteString(placeholderRow)
	args = append(args, items[0].dbCols()...)

	for _, item := range items[1:] {
		q.WriteRune(',')
		q.WriteString(placeholderRow)
		args = append(args, item.dbCols()...)
	}

	return q.String(), args
}

func GetRelationTuples[RT pop.TableNameAble](conn *pop.Connection, lastID uuid.UUID) (
	res []RT, hasNext bool, err error,
) {
	const pageSize = 500
	q := conn.Where("shard_id > ?", lastID).Order("shard_id").Limit(pageSize + 1)

	if err := q.All(&res); err != nil {
		return nil, false, sqlcon.HandleError(err)
	}
	if len(res) > pageSize {
		return res[:pageSize], true, nil
	}
	return res, false, nil
}

func BatchWriteMappings(conn *pop.Connection, mappings []*UUIDMapping) (err error) {
	if len(mappings) == 0 {
		// Nothing to do.
		return nil
	}

	placeholders, args := ConstructArgs(2, mappings)

	// We need to write manual SQL here because the INSERT should not fail if
	// the UUID already exists, but we still want to return an error if anything
	// else goes wrong.
	var query string
	switch d := conn.Dialect.Name(); d {
	case "mysql":
		query = `INSERT IGNORE INTO keto_uuid_mappings (id, string_representation) VALUES ` + placeholders
	default:
		query = `
			INSERT INTO keto_uuid_mappings (id, string_representation)
			VALUES ` + placeholders + `
			ON CONFLICT (id) DO NOTHING`
	}

	if err = sqlcon.HandleError(conn.RawQuery(query, args...).Exec()); err != nil {
		return err
	}

	return nil
}

func BatchReplaceUUIDs(conn *pop.Connection, uuidToTargets map[uuid.UUID][]*string) error {
	if len(uuidToTargets) == 0 {
		return nil
	}

	ids := maps.Keys(uuidToTargets)

	mappings := &[]UUIDMapping{}
	query := conn.Where("id in (?)", ids)
	if err := sqlcon.HandleError(query.All(mappings)); err != nil {
		return err
	}

	// Write the representation to the correct pointer(s).
	for _, m := range *mappings {
		for _, target := range uuidToTargets[m.ID] {
			*target = m.StringRepresentation
		}
	}

	return nil
}

func BatchInsertTuples[RT interface {
	pop.TableNameAble
	ColumnProvider
}](conn *pop.Connection, rts []RT) error {
	if len(rts) == 0 {
		return nil
	}

	placeholders, args := ConstructArgs(10, rts)
	query := fmt.Sprintf("INSERT INTO %s (shard_id, nid, namespace_id, object, relation, subject_id, subject_set_namespace_id, subject_set_object, subject_set_relation, commit_time) VALUES %s", rts[0].TableName(), placeholders)

	return sqlcon.HandleError(conn.RawQuery(query, args...).Exec())
}
