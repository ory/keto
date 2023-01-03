// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/ory/x/pointerx"

	"github.com/ory/keto/ketoapi"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// the command delegates most of the functionality to the parseFile helper, so we test that
func TestParseCmdParseFile(t *testing.T) {
	for _, tc := range []struct {
		input, name string
		expected    []*ketoapi.RelationTuple
	}{
		{
			name:  "single basic tuple",
			input: "nspace:obj#rel@sub\n",
			expected: []*ketoapi.RelationTuple{{
				Namespace: "nspace",
				Object:    "obj",
				Relation:  "rel",
				SubjectID: pointerx.Ptr("sub"),
			}},
		},
		{
			name: "multiple tuples",
			input: `nspace:obj1#rel@sub1
nspace:obj2#rel@sub2
nspace:obj2#rel@(nspace:obj2#rel)`,
			expected: []*ketoapi.RelationTuple{
				{
					Namespace: "nspace",
					Object:    "obj1",
					Relation:  "rel",
					SubjectID: pointerx.Ptr("sub1"),
				},
				{
					Namespace: "nspace",
					Object:    "obj2",
					Relation:  "rel",
					SubjectID: pointerx.Ptr("sub2"),
				},
				{
					Namespace: "nspace",
					Object:    "obj2",
					Relation:  "rel",
					SubjectSet: &ketoapi.SubjectSet{
						Namespace: "nspace",
						Object:    "obj2",
						Relation:  "rel",
					},
				},
			},
		},
		{
			name: "crap around tuples",
			input: `// foo comment
nspace:obj#rel@sub

  // also indentation and trailing spaces
     nspace:indent#rel@sub  `,
			expected: []*ketoapi.RelationTuple{
				{
					Namespace: "nspace",
					Object:    "obj",
					Relation:  "rel",
					SubjectID: pointerx.Ptr("sub"),
				},
				{
					Namespace: "nspace",
					Object:    "indent",
					Relation:  "rel",
					SubjectID: pointerx.Ptr("sub"),
				},
			},
		},
	} {
		t.Run("case="+tc.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.SetIn(bytes.NewBufferString(tc.input))

			actual, err := parseFile(cmd, "-")
			require.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}

	t.Run("case=reads from fs", func(t *testing.T) {
		dir := t.TempDir()
		fn := filepath.Join(dir, "test.tuples")
		require.NoError(t, os.WriteFile(fn, []byte(`
nspace:obj1#rel@sub1

nspace:obj2#rel@sub2`), 0600))

		actual, err := parseFile(&cobra.Command{}, fn)
		require.NoError(t, err)
		assert.Equal(t, []*ketoapi.RelationTuple{
			{
				Namespace: "nspace",
				Object:    "obj1",
				Relation:  "rel",
				SubjectID: pointerx.Ptr("sub1"),
			},
			{
				Namespace: "nspace",
				Object:    "obj2",
				Relation:  "rel",
				SubjectID: pointerx.Ptr("sub2"),
			},
		}, actual)
	})
}
