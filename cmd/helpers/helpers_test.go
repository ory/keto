// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

func TestParseSubject(t *testing.T) {
	for _, tc := range []struct {
		input    string
		expected *rts.Subject
	}{
		{
			input:    "nspace:obj#rel",
			expected: rts.NewSubjectSet("nspace", "obj", "rel"),
		},
		{
			input:    "someid",
			expected: rts.NewSubjectID("someid"),
		},
		{
			input:    "nspace:id",
			expected: rts.NewSubjectSet("nspace", "id", ""),
		},
	} {
		t.Run("subject="+tc.input, func(t *testing.T) {
			actual, err := ParseSubject(tc.input)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
