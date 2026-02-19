// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/ory/keto/ketoapi"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseTuplesFromPath(t *testing.T) {
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
				SubjectID: new("sub"),
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
					SubjectID: new("sub1"),
				},
				{
					Namespace: "nspace",
					Object:    "obj2",
					Relation:  "rel",
					SubjectID: new("sub2"),
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
					SubjectID: new("sub"),
				},
				{
					Namespace: "nspace",
					Object:    "indent",
					Relation:  "rel",
					SubjectID: new("sub"),
				},
			},
		},
	} {
		t.Run("case="+tc.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.SetIn(bytes.NewBufferString(tc.input))

			actual, err := readTuplesFromPath(cmd, "-", 0, parseHumanReadable)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}

	t.Run("case=error contains line number", func(t *testing.T) {
		input := "nspace:obj#rel@subns:subid\n\n\ninvalid-line\nnspace:obj2#rel@subns:subid2\n"
		cmd := &cobra.Command{}
		cmd.SetIn(bytes.NewBufferString(input))

		_, err := readTuplesFromPath(cmd, "-", 0, parseHumanReadable)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "line 4:")
	})

	t.Run("case=reads from fs", func(t *testing.T) {
		dir := t.TempDir()
		fn := filepath.Join(dir, "test.tuples")
		require.NoError(t, os.WriteFile(fn, []byte(`
nspace:obj1#rel@sub1

nspace:obj2#rel@sub2`), 0o600))

		actual, err := readTuplesFromPath(&cobra.Command{}, fn, 0, parseHumanReadable)
		require.NoError(t, err)
		assert.Equal(t, []*ketoapi.RelationTuple{
			{
				Namespace: "nspace",
				Object:    "obj1",
				Relation:  "rel",
				SubjectID: new("sub1"),
			},
			{
				Namespace: "nspace",
				Object:    "obj2",
				Relation:  "rel",
				SubjectID: new("sub2"),
			},
		}, actual)
	})
}
