// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/cmd/helpers"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/ketoapi"
)

func TestDeleteCmd(t *testing.T) {
	nspace := &namespace.Namespace{Name: t.Name()}
	nspaceUser := &namespace.Namespace{Name: "User"}

	newCmd := func() *cobra.Command {
		cmd := &cobra.Command{
			Use: "keto",
		}
		RegisterCommandsRecursive(cmd)
		return cmd
	}

	otherNspace := &namespace.Namespace{Name: "other"}
	ts := client.NewTestServer(t, []*namespace.Namespace{nspace, nspaceUser, otherNspace}, newCmd)
	defer ts.Shutdown(t)

	// createTuple is a helper that creates a tuple via the create command so we can then delete it.
	createTuple := func(t *testing.T, tuple *ketoapi.RelationTuple) {
		t.Helper()
		tmpFile := filepath.Join(t.TempDir(), "tuple.json")
		data, err := json.Marshal(tuple)
		require.NoError(t, err)
		require.NoError(t, os.WriteFile(tmpFile, data, 0o600))
		ts.Cmd.ExecNoErr(t, "relation-tuple", "create", "-f="+tmpFile)
	}

	type getTuple func() *ketoapi.RelationTuple

	fns := map[string]getTuple{
		"subjectID": func() *ketoapi.RelationTuple {
			return helpers.RandomTupleWithSubjectID(nspace.Name)
		},
		"subjectSet": func() *ketoapi.RelationTuple {
			return helpers.RandomTupleWithSubjectSet(nspace.Name, nspaceUser.Name)
		},
	}

	for name, makeTuple := range fns {
		t.Run("case="+name, func(t *testing.T) {
			t.Run("case=deletes single tuple from file", func(t *testing.T) {
				tuple := makeTuple()
				createTuple(t, tuple)

				tmpFile := filepath.Join(t.TempDir(), "tuple.json")
				data, err := json.Marshal(tuple)
				require.NoError(t, err)
				require.NoError(t, os.WriteFile(tmpFile, data, 0o600))

				stdOut := ts.Cmd.ExecNoErr(t, "relation-tuple", "delete", "-f="+tmpFile)
				require.Equal(t, renderTable(NewAPICollection([]*ketoapi.RelationTuple{tuple})), stdOut)

				// assert that the tuple is gone
				res := ts.Cmd.ExecNoErr(t, "relation-tuple", "get", "--namespace", tuple.Namespace, "--object", tuple.Object)
				require.Equal(t, renderTable(&responseOutput{
					RelationTuples: NewAPICollection(nil),
					IsLastPage:     true,
				}), res)
			})

			t.Run("case=deletes tuples from stdin", func(t *testing.T) {
				tuple := makeTuple()
				createTuple(t, tuple)

				data, err := json.Marshal(tuple)
				require.NoError(t, err)

				stdOut, stdErr, err := ts.Cmd.Exec(bytes.NewReader(data), "relation-tuple", "delete", "-f=-")
				require.NoError(t, err)
				require.Empty(t, stdErr)
				require.Equal(t, renderTable(NewAPICollection([]*ketoapi.RelationTuple{tuple})), stdOut)
			})

			t.Run("case=deletes tuples from directory", func(t *testing.T) {
				dir := t.TempDir()

				tuple1 := makeTuple()
				tuple2 := makeTuple()
				createTuple(t, tuple1)
				createTuple(t, tuple2)

				data1, err := json.Marshal(tuple1)
				require.NoError(t, err)
				require.NoError(t, os.WriteFile(filepath.Join(dir, "tuple1.json"), data1, 0o600))

				data2, err := json.Marshal(tuple2)
				require.NoError(t, err)
				require.NoError(t, os.WriteFile(filepath.Join(dir, "tuple2.json"), data2, 0o600))

				stdOut := ts.Cmd.ExecNoErr(t, "relation-tuple", "delete", "-f="+dir)
				require.Equal(t, renderTable(NewAPICollection([]*ketoapi.RelationTuple{tuple1, tuple2})), stdOut)
			})

			t.Run("case=deletes inline tuple", func(t *testing.T) {
				tuple := makeTuple()
				createTuple(t, tuple)

				subject := ""
				if tuple.SubjectID != nil {
					subject = *tuple.SubjectID
				} else {
					subject = tuple.SubjectSet.String()
				}

				stdOut := ts.Cmd.ExecNoErr(t, "relation-tuple", "delete", subject, tuple.Relation, tuple.Namespace+":"+tuple.Object)
				require.Equal(t, renderTable(NewAPICollection([]*ketoapi.RelationTuple{tuple})), stdOut)
			})
		})
	}

	t.Run("case=fails on invalid JSON", func(t *testing.T) {
		tmpFile := filepath.Join(t.TempDir(), "invalid.json")
		require.NoError(t, os.WriteFile(tmpFile, []byte("not valid json"), 0o600))

		_, stdErr, err := ts.Cmd.Exec(nil, "relation-tuple", "delete", "-f="+tmpFile)
		require.Error(t, err)
		require.Contains(t, stdErr, "could not decode")
	})

	t.Run("case=fails on nonexistent file", func(t *testing.T) {
		stdErr := ts.Cmd.ExecExpectedErr(t, "relation-tuple", "delete", "-f=/nonexistent/file.json")
		require.Contains(t, stdErr, "error getting stats")
	})

	t.Run("case=fails on json file arguments without -f", func(t *testing.T) {
		_, stdErr, err := ts.Cmd.Exec(nil, "relation-tuple", "delete", "file1.json", "file2.json")
		require.Error(t, err)
		require.NotEmpty(t, stdErr)
	})

	t.Run("case=missing colon between ns and obj", func(t *testing.T) {
		_, stdErr, err := ts.Cmd.Exec(nil, "relation-tuple", "delete", "sns:so1", "rel", "nsobj")
		require.Error(t, err)
		require.Contains(t, stdErr, "expected <object_namespace>:<object_id> format, got \"nsobj\"")
	})

	t.Run("case=handles JSON with leading whitespace and BOM", func(t *testing.T) {
		// Test with various prefixes: whitespace, BOM, and combination
		prefixes := [][]byte{
			{'\n', '\n', ' ', ' ', '\t'},
			{0xEF, 0xBB, 0xBF},
			{0xEF, 0xBB, 0xBF, '\n', ' ', ' '},
		}

		for _, prefix := range prefixes {
			// Test single tuple
			tuple := helpers.RandomTupleWithSubjectID(nspace.Name)
			createTuple(t, tuple)

			tmpFile := filepath.Join(t.TempDir(), "tuple.json")
			data, err := json.Marshal(tuple)
			require.NoError(t, err)
			dataWithPrefix := append(prefix, data...)
			require.NoError(t, os.WriteFile(tmpFile, dataWithPrefix, 0o600))

			stdOut := ts.Cmd.ExecNoErr(t, "relation-tuple", "delete", "-f="+tmpFile)
			require.Equal(t, renderTable(NewAPICollection([]*ketoapi.RelationTuple{tuple})), stdOut)

			// Test array
			tuple1 := helpers.RandomTupleWithSubjectID(nspace.Name)
			tuple2 := helpers.RandomTupleWithSubjectID(nspace.Name)
			createTuple(t, tuple1)
			createTuple(t, tuple2)

			tmpFile2 := filepath.Join(t.TempDir(), "tuples.json")
			data2, err := json.Marshal([]*ketoapi.RelationTuple{tuple1, tuple2})
			require.NoError(t, err)
			dataWithPrefix2 := append(prefix, data2...)
			require.NoError(t, os.WriteFile(tmpFile2, dataWithPrefix2, 0o600))

			stdOut2 := ts.Cmd.ExecNoErr(t, "relation-tuple", "delete", "-f="+tmpFile2)
			require.Equal(t, renderTable(NewAPICollection([]*ketoapi.RelationTuple{tuple1, tuple2})), stdOut2)
		}
	})
}

func TestReadTuplesFromDir(t *testing.T) {
	writeTuple := func(t *testing.T, dir string, name string) *ketoapi.RelationTuple {
		t.Helper()
		tuple := helpers.RandomTupleWithSubjectID(t.Name())
		data, err := json.Marshal(tuple)
		require.NoError(t, err)
		require.NoError(t, os.WriteFile(filepath.Join(dir, name), data, 0o600))
		return tuple
	}

	mkdirs := func(t *testing.T, root string, depth int) string {
		t.Helper()
		dir := root
		for i := range depth {
			dir = filepath.Join(dir, fmt.Sprintf("d%d", i))
		}
		require.NoError(t, os.MkdirAll(dir, 0o750))
		return dir
	}

	t.Run("case=succeeds at MaxDirDepth nested directories", func(t *testing.T) {
		root := t.TempDir()
		deepest := mkdirs(t, root, MaxDirDepth)

		// Place a file inside the second-to-last directory so its
		// path-component depth equals MaxDirDepth.
		tuple := writeTuple(t, filepath.Dir(deepest), "tuple.json")

		tuples, err := readTuplesFromDir(root, decodeTuples)
		require.NoError(t, err)
		require.Equal(t, []*ketoapi.RelationTuple{tuple}, tuples)
	})

	t.Run("case=fails at MaxDirDepth+1 nested directories", func(t *testing.T) {
		root := t.TempDir()
		mkdirs(t, root, MaxDirDepth+1)

		_, err := readTuplesFromDir(root, decodeTuples)
		require.ErrorContains(t, err, "maximum directory depth")
	})
}
