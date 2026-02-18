// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/cmd/helpers"
	"github.com/ory/keto/cmd/relationtuple"
	"github.com/ory/keto/internal/namespace"
)

func TestCheckCommand(t *testing.T) {
	nspace := &namespace.Namespace{Name: t.Name()}
	nspaceUser := &namespace.Namespace{Name: "User"}

	newCmd := func() *cobra.Command {
		cmd := &cobra.Command{
			Use: "keto",
		}
		RegisterCommandsRecursive(cmd)
		relationtuple.RegisterCommandsRecursive(cmd)
		return cmd
	}

	ts := client.NewTestServer(t, []*namespace.Namespace{nspace, nspaceUser}, newCmd)
	defer ts.Shutdown(t)

	tuple1 := helpers.RandomTupleWithSubjectID(nspace.Name)
	tuple2 := helpers.RandomTupleWithSubjectSet(nspace.Name, nspaceUser.Name)
	ts.Cmd.ExecNoErr(t, "relation-tuple", "create", tuple2.SubjectSet.String(), tuple2.Relation, tuple2.Namespace+":"+tuple2.Object)
	ts.Cmd.ExecNoErr(t, "relation-tuple", "create", *tuple1.SubjectID, tuple1.Relation, tuple1.Namespace+":"+tuple1.Object)

	t.Run("case=SubjectSet", func(t *testing.T) {
		subject := tuple2.SubjectSet.String()
		rel := tuple2.Relation
		nsObj := tuple2.Namespace + ":" + tuple2.Object

		noPermSubject := tuple2.SubjectSet.Namespace + ":no-perm-subject"

		t.Run("case=allowed when tuple exists", func(t *testing.T) {
			stdOut := ts.Cmd.ExecNoErr(t, "check", subject, rel, nsObj)
			require.Equal(t, "Allowed\n", stdOut)
		})

		t.Run("case=denied for unrelated subject", func(t *testing.T) {
			stdOut := ts.Cmd.ExecNoErr(t, "check", noPermSubject, rel, nsObj)
			require.Equal(t, "Denied\n", stdOut)
		})

		t.Run("case=4-arg format still works with deprecation warning", func(t *testing.T) {
			stdOut, stdErr, err := ts.Cmd.Exec(nil, "check", subject, rel, tuple2.Namespace, tuple2.Object)
			require.NoError(t, err)
			require.Equal(t, "Allowed\n", stdOut)
			require.Contains(t, stdErr, "deprecated")
		})

		t.Run("case=errors on invalid namespace:object format", func(t *testing.T) {
			_, stdErr, err := ts.Cmd.Exec(nil, "check", subject, rel, "no-colon-here")
			require.Error(t, err)
			require.Contains(t, stdErr, "expected <object_namespace>:<object_id> format")
		})
	})

	t.Run("case=SubjectID", func(t *testing.T) {
		subject := *tuple1.SubjectID
		rel := tuple1.Relation
		nsObj := tuple1.Namespace + ":" + tuple1.Object

		noPermSubject := "no-perm-subject"

		t.Run("case=SubjectID, allowed when tuple exists", func(t *testing.T) {
			stdOut := ts.Cmd.ExecNoErr(t, "check", subject, rel, nsObj)
			require.Equal(t, "Allowed\n", stdOut)
		})

		t.Run("case=denied for unrelated object", func(t *testing.T) {
			stdOut := ts.Cmd.ExecNoErr(t, "check", noPermSubject, rel, nspace.Name+":nope")
			require.Equal(t, "Denied\n", stdOut)
		})

		t.Run("case=4-arg format still works with deprecation warning", func(t *testing.T) {
			stdOut, stdErr, err := ts.Cmd.Exec(nil, "check", subject, rel, tuple1.Namespace, tuple1.Object)
			require.NoError(t, err)
			require.Equal(t, "Allowed\n", stdOut)
			require.Contains(t, stdErr, "deprecated")
		})
	})
}
