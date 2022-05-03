package relationtuple

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// MapAndWriteTuples is a test helper to write relation tuples to the database
// while mapping all strings to UUIDs.
func MapAndWriteTuples(t *testing.T, m ManagerProvider, tuples ...*InternalRelationTuple) {
	t.Helper()

	require.NoError(t, m.UUIDMappingManager().MapFieldsToUUID(context.Background(), InternalRelationTuples(tuples)))
	require.NoError(t, m.RelationTupleManager().WriteRelationTuples(context.Background(), tuples...))
	require.NoError(t, m.UUIDMappingManager().MapFieldsFromUUID(context.Background(), InternalRelationTuples(tuples)))
}
