// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"context"
	"testing"

	"github.com/ory/keto/ketoapi"

	"github.com/stretchr/testify/require"
)

// MapAndWriteTuples is a test helper to write relationships to the database
// while mapping all strings to UUIDs.
func MapAndWriteTuples(t *testing.T, m interface {
	MapperProvider
	ManagerProvider
}, tuples ...*ketoapi.RelationTuple) {
	t.Helper()

	its, err := m.Mapper().FromTuple(context.Background(), tuples...)
	require.NoError(t, err)
	require.NoError(t, m.RelationTupleManager().WriteRelationTuples(context.Background(), its...))
}
