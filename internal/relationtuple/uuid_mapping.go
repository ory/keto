package relationtuple

import (
	"context"
	"github.com/ory/keto/ketoapi"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type (
	// UUIDMappable is an interface for objects that have fields that can be
	// mapped to and from UUIDs.
	UUIDMappable interface{ UUIDMappableFields() []*string }

	UUIDMappingManager interface {
		MapStringsToUUIDs(ctx context.Context, s ...string) ([]uuid.UUID, error)

		// MapFieldsToUUID maps all fields of the given object to UUIDs.
		MapFieldsToUUID(ctx context.Context, m UUIDMappable) error

		// MapFieldsFromUUID maps all fields of the given object from UUIDs to
		// their string value.
		MapFieldsFromUUID(ctx context.Context, m UUIDMappable) error
	}
)

func preferNil[V any, PtrV interface{ *V }, F any](v PtrV, fallback F) (res F) {
	if v == nil {
		return
	}
	return fallback
}

func InternalRelationQuery(ctx context.Context, m UUIDMappingManager, q *ketoapi.RelationQuery) (*RelationQuery, error) {
	iq := &RelationQuery{
		Namespace: q.Namespace,
		Relation:  q.Relation,
	}

	if q.SubjectID != nil {
		mappings, err := m.MapStringsToUUIDs(ctx, q.Object, *q.SubjectID)
		if err != nil {
			return nil, err
		}
		iq.Object = mappings[0]
		iq.SubjectID = &mappings[1]
	}
	if q.SubjectSet != nil {
		mappings, err := m.MapStringsToUUIDs(ctx, q.Object, q.SubjectSet.Object)
		if err != nil {
			return nil, err
		}
		iq.Object = mappings[0]
		iq.SubjectSet = &SubjectSet{
			Namespace: q.SubjectSet.Namespace,
			Object:    mappings[1],
			Relation:  q.SubjectSet.Relation,
		}
	}

	return iq, nil
}

func UUIDMappingManagerTest(t *testing.T, m UUIDMappingManager) {
	ctx := context.Background()

	t.Run("case=ToUUID_FromUUID", func(t *testing.T) {
		s1 := SubjectID{"rep1"}
		err := m.MapFieldsToUUID(ctx, &s1)
		require.NoError(t, err)

		s2 := SubjectID{s1.ID}
		err = m.MapFieldsFromUUID(ctx, &s2)
		assert.NoError(t, err)
		assert.Equal(t, "rep1", s2.ID)
	})

	t.Run("case=Idempotent_ToUUID", func(t *testing.T) {
		s1 := SubjectID{"string"}
		s2 := SubjectID{"string"}
		assert.NoError(t, m.MapFieldsToUUID(ctx, &s1))
		assert.NoError(t, m.MapFieldsToUUID(ctx, &s2))
		assert.Equal(t, s1.ID, s2.ID)
		assert.NotEqual(t, "string", s1.ID)
	})

	t.Run("case=batch to UUID", func(t *testing.T) {
		rt := InternalRelationTuple{Object: "object", Subject: &SubjectID{"subject"}}
		assert.NoError(t, m.MapFieldsToUUID(ctx, &rt))
		objectUUID, err := uuid.FromString(rt.Object)
		assert.NoError(t, err)
		subjectUUID, err := uuid.FromString(rt.Subject.String())
		assert.NoError(t, err)

		rt2 := InternalRelationTuple{Object: "object", Subject: &SubjectID{"another subject"}}
		assert.NoError(t, m.MapFieldsToUUID(ctx, &rt2))
		assert.Equal(t, objectUUID, uuid.Must(uuid.FromString(rt2.Object)))
		assert.NotEqual(t, subjectUUID, uuid.Must(uuid.FromString(rt2.Subject.String())))

		rt3 := InternalRelationTuple{Object: "another object", Subject: &SubjectID{"subject"}}
		assert.NoError(t, m.MapFieldsToUUID(ctx, &rt3))
		assert.NotEqual(t, objectUUID, uuid.Must(uuid.FromString(rt3.Object)))
		assert.Equal(t, subjectUUID, uuid.Must(uuid.FromString(rt3.Subject.String())))
	})

	t.Run("case=IdempotentMapFieldsToAndFromUUIDs", func(t *testing.T) {
		tc := []struct {
			name string
			obj  UUIDMappable
			copy UUIDMappable
		}{
			{
				name: "RelationTuple",
				obj:  &InternalRelationTuple{Namespace: "n", Relation: "r", Object: "Object", Subject: &SubjectID{ID: "Subject"}},
				copy: &InternalRelationTuple{Namespace: "n", Relation: "r", Object: "Object", Subject: &SubjectID{ID: "Subject"}},
			}, {
				name: "SubjectID",
				obj:  &SubjectID{ID: "sub"},
				copy: &SubjectID{ID: "sub"},
			},
		}
		for _, tt := range tc {
			t.Run("type="+tt.name, func(t *testing.T) {
				assert.NoError(t, m.MapFieldsToUUID(ctx, tt.obj))
				assert.NoError(t, m.MapFieldsFromUUID(ctx, tt.obj))
				assert.Equal(t, tt.copy, tt.obj)
			})
		}
	})
}
