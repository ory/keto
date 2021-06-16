package graph

import (
	"context"
	"testing"

	"github.com/ory/keto/internal/relationtuple"

	"github.com/stretchr/testify/assert"
)

func TestEngineUtilsProvider_CheckVisited(t *testing.T) {
	t.Run("case=finds cycle", func(t *testing.T) {

		linkedList := []relationtuple.SubjectSet{{
			Namespace: "default",
			Object:    "A",
			Relation:  "connected",
		}, {
			Namespace: "default",
			Object:    "B",
			Relation:  "connected",
		}, {
			Namespace: "default",
			Object:    "C",
			Relation:  "connected",
		}, {
			Namespace: "default",
			Object:    "B",
			Relation:  "connected",
		}, {
			Namespace: "default",
			Object:    "D",
			Relation:  "connected",
		}}

		ctx := context.Background()
		var isThereACycle bool
		for i := range linkedList {
			ctx, isThereACycle = CheckAndAddVisited(ctx, &linkedList[i])
			if isThereACycle {
				break
			}
		}

		assert.Equal(t, isThereACycle, true)
	})

	t.Run("case=ignores if no cycle", func(t *testing.T) {

		list := []relationtuple.SubjectSet{{
			Namespace: "default",
			Object:    "A",
			Relation:  "connected",
		}, {
			Namespace: "default",
			Object:    "B",
			Relation:  "connected",
		}, {
			Namespace: "default",
			Object:    "C",
			Relation:  "connected",
		}, {
			Namespace: "default",
			Object:    "D",
			Relation:  "connected",
		}, {
			Namespace: "default",
			Object:    "E",
			Relation:  "connected",
		}}

		ctx := context.Background()
		var isThereACycle bool
		for i := range list {
			ctx, isThereACycle = CheckAndAddVisited(ctx, &list[i])
			if isThereACycle {
				break
			}
		}

		assert.Equal(t, isThereACycle, false)
	})
}
