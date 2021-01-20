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
	for i, tc := range []struct {
		size           int
		token          string
		expectedErr    error
		expectedOffset int
		expectedLimit  int
	}{
		{
			size:           10,
			token:          "10",
			expectedOffset: 10,
			expectedLimit:  10,
		},
		{
			size:           0,
			token:          "15",
			expectedOffset: 15,
			expectedLimit:  defaultPageSize,
		},
		{
			size:          0,
			token:         "-15",
			expectedErr:   persistence.ErrMalformedPageToken,
			expectedLimit: defaultPageSize,
		},
	} {
		t.Run(fmt.Sprintf("case=%d/size:%d token:%s", i, tc.size, tc.token), func(t *testing.T) {
			pagination, err := internalPaginationFromOptions(x.WithSize(tc.size), x.WithToken(tc.token))

			assert.True(t, errors.Is(err, tc.expectedErr))
			assert.Equal(t, tc.expectedLimit, pagination.Limit)
			assert.Equal(t, tc.expectedOffset, pagination.Offset)
		})
	}
}
