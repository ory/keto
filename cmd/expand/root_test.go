// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package expand

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/x/cmdx"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/cmd/helpers"
	"github.com/ory/keto/cmd/relationtuple"
	"github.com/ory/keto/internal/namespace"
)

func TestExpandCommand(t *testing.T) {
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

	tuple := helpers.RandomTupleWithSubjectSet(nspace.Name, nspaceUser.Name)

	nsObj := tuple.Namespace + ":" + tuple.Object
	rel := tuple.Relation
	subjectSet := tuple.SubjectSet.String()

	// create tuple with subjectSet
	ts.Cmd.ExecNoErr(t, "relation-tuple", "create", subjectSet, rel, nsObj)

	t.Run("case=expands existing tuple as JSON", func(t *testing.T) {
		stdOut := ts.Cmd.ExecNoErr(t, "expand", rel, nsObj,
			"--"+cmdx.FlagFormat, string(cmdx.FormatJSON))
		assert.Contains(t, stdOut, tuple.SubjectSet.Object)
	})

	t.Run("case=unknown tuple returns null JSON", func(t *testing.T) {
		stdOut := ts.Cmd.ExecNoErr(t, "expand", rel, nspace.Name+":unknown-obj",
			"--"+cmdx.FlagFormat, string(cmdx.FormatJSON))
		assert.Equal(t, "null\n", stdOut)
	})

	t.Run("case=unknown tuple prints empty tree in default format", func(t *testing.T) {
		stdOut := ts.Cmd.ExecNoErr(t, "expand", rel, nspace.Name+":unknown-obj")
		assert.Contains(t, stdOut, "empty tree")
	})

	t.Run("case=3-arg format still works with deprecation warning", func(t *testing.T) {
		stdOut, stdErr, err := ts.Cmd.Exec(nil, "expand", rel, nspace.Name, tuple.Object,
			"--"+cmdx.FlagFormat, string(cmdx.FormatJSON))
		require.NoError(t, err)
		assert.Contains(t, stdOut, tuple.SubjectSet.Object)
		require.Contains(t, stdErr, "deprecated")
	})

	t.Run("case=errors on invalid namespace:object format", func(t *testing.T) {
		_, stdErr, err := ts.Cmd.Exec(nil, "expand", rel, "no-colon-here")
		require.Error(t, err)
		require.Contains(t, stdErr, "expected <object_namespace>:<object_id> format")
	})
}
