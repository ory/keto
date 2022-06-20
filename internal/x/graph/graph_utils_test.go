package graph

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"

	"github.com/ory/keto/internal/relationtuple"

	"github.com/stretchr/testify/assert"
)

func TestEngineUtilsProvider_CheckVisited(t *testing.T) {
	a, b, c, d, e := uuid.Must(uuid.NewV4()), uuid.Must(uuid.NewV4()), uuid.Must(uuid.NewV4()), uuid.Must(uuid.NewV4()), uuid.Must(uuid.NewV4())
	t.Run("case=finds cycle", func(t *testing.T) {
		linkedList := []relationtuple.SubjectSet{{
			Namespace: 1,
			Object:    a,
			Relation:  "connected",
		}, {
			Namespace: 1,
			Object:    b,
			Relation:  "connected",
		}, {
			Namespace: 1,
			Object:    c,
			Relation:  "connected",
		}, {
			Namespace: 1,
			Object:    b,
			Relation:  "connected",
		}, {
			Namespace: 1,
			Object:    d,
			Relation:  "connected",
		}}

		ctx := context.Background()
		var isThereACycle bool
		for i := range linkedList {
			ctx, isThereACycle = CheckAndAddVisited(ctx, linkedList[i].Hash())
			if isThereACycle {
				break
			}
		}

		assert.Equal(t, isThereACycle, true)
	})

	t.Run("case=ignores if no cycle", func(t *testing.T) {
		list := []relationtuple.SubjectSet{{
			Namespace: 1,
			Object:    a,
			Relation:  "connected",
		}, {
			Namespace: 1,
			Object:    b,
			Relation:  "connected",
		}, {
			Namespace: 1,
			Object:    c,
			Relation:  "connected",
		}, {
			Namespace: 1,
			Object:    d,
			Relation:  "connected",
		}, {
			Namespace: 1,
			Object:    e,
			Relation:  "connected",
		}}

		ctx := context.Background()
		var isThereACycle bool
		for i := range list {
			ctx, isThereACycle = CheckAndAddVisited(ctx, list[i].Hash())
			if isThereACycle {
				break
			}
		}

		assert.Equal(t, isThereACycle, false)
	})
}
