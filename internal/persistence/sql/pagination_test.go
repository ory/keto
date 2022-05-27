package sql

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/internal/x"
)

func TestPaginationToken(t *testing.T) {
	t.Parallel()

	for i, tc := range []struct {
		size            int
		token           string
		expectedErr     error
		expectedPage    int
		expectedPerPage int
	}{
		{
			size:            10,
			token:           "10",
			expectedPage:    10,
			expectedPerPage: 10,
		},
		{
			size:            0,
			token:           "15",
			expectedPage:    15,
			expectedPerPage: defaultPageSize,
		},
		{
			size:            0,
			token:           "-15",
			expectedErr:     persistence.ErrMalformedPageToken,
			expectedPerPage: defaultPageSize,
		},
	} {
		t.Run(fmt.Sprintf("case=%d/size:%d token:%s", i, tc.size, tc.token), func(t *testing.T) {
			pagination, err := internalPaginationFromOptions(x.WithSize(tc.size), x.WithToken(tc.token))

			assert.True(t, errors.Is(err, tc.expectedErr))
			assert.Equal(t, tc.expectedPerPage, pagination.PerPage)
			assert.Equal(t, tc.expectedPage, pagination.Page)
		})
	}
}
