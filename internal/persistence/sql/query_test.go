package sql

import (
	"database/sql"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/ory/x/uuidx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/relationtuple"
)

func TestBuildDelete(t *testing.T) {
	t.Parallel()
	nid := uuidx.NewV4()

	q, args, err := buildDelete(nid, nil)
	assert.Error(t, err)
	assert.Empty(t, q)
	assert.Empty(t, args)

	obj1, obj2, sub1, obj3 := uuidx.NewV4(), uuidx.NewV4(), uuidx.NewV4(), uuidx.NewV4()

	q, args, err = buildDelete(nid, []*relationtuple.RelationTuple{
		{
			Namespace: "ns1",
			Object:    obj1,
			Relation:  "rel1",
			Subject: &relationtuple.SubjectID{
				ID: sub1,
			},
		},
		{
			Namespace: "ns2",
			Object:    obj2,
			Relation:  "rel2",
			Subject: &relationtuple.SubjectSet{
				Namespace: "ns3",
				Object:    obj3,
				Relation:  "rel3",
			},
		},
	})
	require.NoError(t, err)

	// parentheses are important here
	assert.Equal(t, q, "DELETE FROM keto_relation_tuples WHERE ((namespace = ? AND object = ? AND relation = ? AND subject_id = ? AND subject_set_namespace IS NULL AND subject_set_object IS NULL AND subject_set_relation IS NULL) OR (namespace = ? AND object = ? AND relation = ? AND subject_id IS NULL AND subject_set_namespace = ? AND subject_set_object = ? AND subject_set_relation = ?)) AND nid = ?")
	assert.Equal(t, []any{"ns1", obj1, "rel1", sub1, "ns2", obj2, "rel2", "ns3", obj3, "rel3", nid}, args)
}

func TestBuildInsert(t *testing.T) {
	t.Parallel()
	nid := uuidx.NewV4()

	q, args, err := buildInsert(time.Now(), nid, nil)
	assert.Error(t, err)
	assert.Empty(t, q)
	assert.Empty(t, args)

	obj1, obj2, sub1, obj3 := uuidx.NewV4(), uuidx.NewV4(), uuidx.NewV4(), uuidx.NewV4()

	now := time.Now()

	q, args, err = buildInsert(now, nid, []*relationtuple.RelationTuple{
		{
			Namespace: "ns1",
			Object:    obj1,
			Relation:  "rel1",
			Subject: &relationtuple.SubjectID{
				ID: sub1,
			},
		},
		{
			Namespace: "ns2",
			Object:    obj2,
			Relation:  "rel2",
			Subject: &relationtuple.SubjectSet{
				Namespace: "ns3",
				Object:    obj3,
				Relation:  "rel3",
			},
		},
	})
	require.NoError(t, err)

	assert.Equal(t, q, "INSERT INTO keto_relation_tuples (shard_id, nid, namespace, object, relation, subject_id, subject_set_namespace, subject_set_object, subject_set_relation, commit_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?), (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	assert.Equal(t, []any{
		args[0], // this is kind of cheating but we generate the shard id in the buildInsert function
		nid,
		"ns1",
		obj1,
		"rel1",
		uuid.NullUUID{sub1, true},
		sql.NullString{}, uuid.NullUUID{}, sql.NullString{},
		now,

		args[10], // again, cheating
		nid,
		"ns2",
		obj2,
		"rel2",
		uuid.NullUUID{},
		sql.NullString{"ns3", true}, uuid.NullUUID{obj3, true}, sql.NullString{"rel3", true},
		now,
	}, args)
}

func TestBuildInsertUUIDs(t *testing.T) {
	t.Parallel()

	foo, bar, baz := uuidx.NewV4(), uuidx.NewV4(), uuidx.NewV4()
	uuids := []UUIDMapping{
		{foo, "foo"},
		{bar, "bar"},
		{baz, "baz"},
	}

	q, args := buildInsertUUIDs(uuids, "mysql")
	assert.Equal(t, "INSERT IGNORE INTO keto_uuid_mappings (id, string_representation) VALUES (?,?),(?,?),(?,?)", q)
	assert.Equal(t, []any{foo, "foo", bar, "bar", baz, "baz"}, args)

	q, args = buildInsertUUIDs(uuids, "anything else")
	assert.Equal(t, "INSERT INTO keto_uuid_mappings (id, string_representation) VALUES (?,?),(?,?),(?,?) ON CONFLICT (id) DO NOTHING", q)
	assert.Equal(t, []any{foo, "foo", bar, "bar", baz, "baz"}, args)
}
