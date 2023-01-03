// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package graph

import (
	"context"
	"sync"
	"testing"

	"github.com/gofrs/uuid"

	"github.com/ory/keto/internal/relationtuple"

	"github.com/stretchr/testify/assert"
)

func TestEngineUtilsProvider_CheckVisited(t *testing.T) {
	a, b, c, d, e := uuid.Must(uuid.NewV4()), uuid.Must(uuid.NewV4()), uuid.Must(uuid.NewV4()), uuid.Must(uuid.NewV4()), uuid.Must(uuid.NewV4())
	t.Run("case=finds cycle", func(t *testing.T) {
		linkedList := []relationtuple.SubjectSet{{
			Namespace: "1",
			Object:    a,
			Relation:  "connected",
		}, {
			Namespace: "1",
			Object:    b,
			Relation:  "connected",
		}, {
			Namespace: "1",
			Object:    c,
			Relation:  "connected",
		}, {
			Namespace: "1",
			Object:    b,
			Relation:  "connected",
		}, {
			Namespace: "1",
			Object:    d,
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
			Namespace: "1",
			Object:    a,
			Relation:  "connected",
		}, {
			Namespace: "1",
			Object:    b,
			Relation:  "connected",
		}, {
			Namespace: "1",
			Object:    c,
			Relation:  "connected",
		}, {
			Namespace: "1",
			Object:    d,
			Relation:  "connected",
		}, {
			Namespace: "1",
			Object:    e,
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

	t.Run("case=no race condition during adding", func(t *testing.T) {
		racyObj := uuid.Must(uuid.NewV4())
		otherObj := uuid.Must(uuid.NewV4())
		// we repeat this test a few times to ensure we don't have a race condition
		// the race detector alone was not able to catch it
		for i := 0; i < 500; i++ {
			subject := &relationtuple.SubjectSet{
				Namespace: "default",
				Object:    racyObj,
				Relation:  "connected",
			}

			ctx, _ := CheckAndAddVisited(
				context.Background(),
				&relationtuple.SubjectSet{Object: otherObj},
			)
			var wg sync.WaitGroup
			var aCycle, bCycle bool
			var aCtx, bCtx context.Context

			wg.Add(2)
			go func() {
				aCtx, aCycle = CheckAndAddVisited(ctx, subject)
				wg.Done()
			}()
			go func() {
				bCtx, bCycle = CheckAndAddVisited(ctx, subject)
				wg.Done()
			}()

			wg.Wait()
			// one should be true, and one false
			assert.False(t, aCycle && bCycle)
			assert.True(t, aCycle || bCycle)
			assert.Equal(t, aCtx, bCtx)
		}
	})
}
