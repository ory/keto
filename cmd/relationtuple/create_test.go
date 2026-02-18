// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/randx"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/cmd/helpers"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/ketoapi"
)

func renderTable(table cmdx.Table) string {
	var buf bytes.Buffer
	cmdx.PrintTablef(&buf, table)
	return buf.String()
}

func TestCreateCmd(t *testing.T) {
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

	createFile := func(t *testing.T, tuple any) string {
		dir := t.TempDir()

		data, err := json.Marshal(tuple)
		require.NoError(t, err)

		filename := randx.MustString(10, randx.AlphaLowerNum) + ".json"
		require.NoError(t, os.WriteFile(filepath.Join(dir, filename), data, 0o600))
		return filepath.Join(dir, filename)
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

	for name, createTuple := range fns {
		t.Run("case="+name, func(t *testing.T) {
			t.Run("case=creates single tuple from files", func(t *testing.T) {
				tuple1 := createTuple()
				tuple2 := createTuple()

				tmpFile := createFile(t, tuple1)
				tmpFile2 := createFile(t, []*ketoapi.RelationTuple{tuple2})

				stdOut := ts.Cmd.ExecNoErr(t, "relation-tuple", "create", "-f", tmpFile, "-f", tmpFile2)
				require.Equal(t, renderTable(NewAPICollection([]*ketoapi.RelationTuple{tuple1, tuple2})), stdOut)

				// assert that we can `get` it
				res1 := ts.Cmd.ExecNoErr(t, "relation-tuple", "get", "--namespace", tuple1.Namespace, "--object", tuple1.Object)
				require.Equal(t, renderTable(&responseOutput{
					RelationTuples: NewAPICollection([]*ketoapi.RelationTuple{tuple1}),
					IsLastPage:     true,
				}), res1)

				// assert that we can `get` it
				res2 := ts.Cmd.ExecNoErr(t, "relation-tuple", "get", "--namespace", tuple2.Namespace, "--object", tuple2.Object)
				require.Equal(t, renderTable(&responseOutput{
					RelationTuples: NewAPICollection([]*ketoapi.RelationTuple{tuple2}),
					IsLastPage:     true,
				}), res2)
			})

			t.Run("case=creates tuples from files and stdin", func(t *testing.T) {
				tuple1 := createTuple()
				tupleStdin := createTuple()

				tupleFile := createFile(t, tuple1)
				data, err := json.Marshal(tupleStdin)
				require.NoError(t, err)

				stdOut, stdErr, err := ts.Cmd.Exec(bytes.NewReader(data), "relation-tuple", "create", "-f", "-", "-f", tupleFile)
				require.NoError(t, err)
				require.Empty(t, stdErr)
				require.Equal(t, renderTable(NewAPICollection([]*ketoapi.RelationTuple{tupleStdin, tuple1})), stdOut)
			})

			t.Run("case=creates tuples from directory", func(t *testing.T) {
				dir := t.TempDir()

				tuple1 := createTuple()
				tuple2 := createTuple()

				data1, err := json.Marshal(tuple1)
				require.NoError(t, err)
				require.NoError(t, os.WriteFile(filepath.Join(dir, "tuple1.json"), data1, 0o600))

				data2, err := json.Marshal(tuple2)
				require.NoError(t, err)
				require.NoError(t, os.WriteFile(filepath.Join(dir, "tuple2.json"), data2, 0o600))

				stdOut := ts.Cmd.ExecNoErr(t, "relation-tuple", "create", "-f="+dir)
				require.Equal(t, renderTable(NewAPICollection([]*ketoapi.RelationTuple{tuple1, tuple2})), stdOut)
			})

			t.Run("case=creates inline tuple", func(t *testing.T) {
				tuple := createTuple()

				subject := ""
				if tuple.SubjectID != nil {
					subject = *tuple.SubjectID
				} else {
					subject = tuple.SubjectSet.String()
				}

				stdOut := ts.Cmd.ExecNoErr(t, "relation-tuple", "create", subject, tuple.Relation, tuple.Namespace+":"+tuple.Object)
				require.Equal(t, renderTable(NewAPICollection([]*ketoapi.RelationTuple{tuple})), stdOut)
			})
		})
	}

	t.Run("case=fails on invalid JSON", func(t *testing.T) {
		tmpFile := filepath.Join(t.TempDir(), "invalid.json")
		require.NoError(t, os.WriteFile(tmpFile, []byte("not valid json"), 0o600))

		_, stdErr, err := ts.Cmd.Exec(nil, "relation-tuple", "create", "-f="+tmpFile)
		require.Error(t, err)
		require.Contains(t, stdErr, "could not decode")
	})

	t.Run("case=fails on nonexistent file", func(t *testing.T) {
		_, stdErr, err := ts.Cmd.Exec(nil, "relation-tuple", "create", "-f=/nonexistent/file.json")
		require.Error(t, err)
		require.Contains(t, stdErr, "error getting stats")
	})

	t.Run("case=fails on json file arguments without -f", func(t *testing.T) {
		_, stdErr, err := ts.Cmd.Exec(nil, "relation-tuple", "create", "file1.json", "file2.json")
		require.Error(t, err)
		require.Contains(t, stdErr, "expected inline arguments or JSON files")
	})

	t.Run("case=missing colon between ns and obj", func(t *testing.T) {
		_, stdErr, err := ts.Cmd.Exec(nil, "relation-tuple", "create", "sns:so1", "rel", "nsobj")
		require.Error(t, err)
		require.Contains(t, stdErr, "expected <object_namespace>:<object_id> format, got \"nsobj\"")
	})
}
