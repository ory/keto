// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package sql

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gofrs/uuid"

	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/internal/x"
)

func TestPaginationToken(t *testing.T) {
	t.Parallel()

	ids := x.UUIDs(3)
	for i, tc := range []struct {
		size            int
		token           string
		expectedErr     error
		expectedLastID  uuid.UUID
		expectedPerPage int
	}{
		{
			size:            10,
			token:           ids[0].String(),
			expectedLastID:  ids[0],
			expectedPerPage: 10,
		},
		{
			size:            0,
			token:           ids[1].String(),
			expectedLastID:  ids[1],
			expectedPerPage: defaultPageSize,
		},
		{
			size:            0,
			token:           "foobar",
			expectedErr:     persistence.ErrMalformedPageToken,
			expectedPerPage: defaultPageSize,
		},
	} {
		t.Run(fmt.Sprintf("case=%d/size:%d token:%s", i, tc.size, tc.token), func(t *testing.T) {
			pagination, err := internalPaginationFromOptions(x.WithSize(tc.size), x.WithToken(tc.token))

			assert.True(t, errors.Is(err, tc.expectedErr))
			assert.Equal(t, tc.expectedPerPage, pagination.PerPage)
			assert.Equal(t, tc.expectedLastID, pagination.LastID)
		})
	}
}
