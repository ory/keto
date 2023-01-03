// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package uuidmapping_test

import (
	"strings"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/internal/persistence/sql/migrations/uuidmapping"
)

func TestConstructArgs(t *testing.T) {
	createItems := func(n int) []uuidmapping.ColumnProvider {
		items := make([]uuidmapping.ColumnProvider, n)
		for i := range items {
			items[i] = &uuidmapping.UUIDMapping{
				ID:                   uuid.Must(uuid.NewV4()),
				StringRepresentation: "foo",
			}
		}
		return items
	}
	for _, tc := range []struct {
		desc  string
		items []uuidmapping.ColumnProvider
		nCols int
	}{
		{
			desc:  "only one item",
			items: createItems(1),
			nCols: 2,
		},
		{
			desc:  "multiple items",
			items: createItems(10),
			nCols: 2,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			placeholders, args := uuidmapping.ConstructArgs(tc.nCols, tc.items)
			assert.Len(t, args, tc.nCols*len(tc.items))
			assert.Len(t, args, strings.Count(placeholders, "?"))
		})
	}
}
