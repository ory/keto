package relationtuple

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/x"
)

type (
	// UUIDMappable is an interface for objects that have fields that can be
	// mapped to and from UUIDs.
	UUIDMappable interface{ UUIDMappableFields() []*string }

	UUIDMappingManager interface {
		// ToUUID returns the mapped UUID for the given string representation.
		// If the string representation is not mapped, a new UUID will be
		// created automatically.
		ToUUID(ctx context.Context, representation string) (uuid.UUID, error)

		// MapFields maps all fields of the given object to UUIDs.
		MapFieldsToUUID(ctx context.Context, m UUIDMappable) error

		// MapFieldsFromUUID maps all fields of the given object from UUIDs to
		// their string value.
		MapFieldsFromUUID(ctx context.Context, m UUIDMappable) error

		// FromUUID returns the text representations for the given UUIDs, such
		// that ids[i] is mapped to reps[i].
		//
		// Of the pagination options, only the page size is considered and used
		// as a batch size.
		FromUUID(ctx context.Context, ids []uuid.UUID, opts ...x.PaginationOptionSetter) (reps []string, err error)
	}
)

func UUIDMappingManagerTest(t *testing.T, m UUIDMappingManager) {
	ctx := context.Background()

	t.Run("case=ToUUID_FromUUID", func(t *testing.T) {
		rep1 := "foo"
		id, err := m.ToUUID(ctx, rep1)
		require.NoError(t, err)

		rep2, err := m.FromUUID(ctx, []uuid.UUID{id})
		assert.NoError(t, err)
		assert.Equal(t, rep1, rep2[0])
	})

	t.Run("case=Idempotent_ToUUID", func(t *testing.T) {
		id1, err := m.ToUUID(ctx, "string")
		assert.NoError(t, err)
		id2, err := m.ToUUID(ctx, "string")
		assert.NoError(t, err)
		assert.Equal(t, id1, id2)
	})

	// Test that the batch mapping preserves ordering, i.e. id[i] is mapped to
	// rep[i].
	t.Run("case=Batch_ToUUID_Paginates", func(t *testing.T) {
		expected := []string{"foo", "foo", "bar", "baz"}
		ids := make([]uuid.UUID, len(expected))
		for i, s := range expected {
			var err error
			ids[i], err = m.ToUUID(ctx, s)
			assert.NoError(t, err)
		}

		actual, err := m.FromUUID(ctx, ids, x.WithSize(1))
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
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
