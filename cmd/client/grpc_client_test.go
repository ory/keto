// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"bytes"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRemote(t *testing.T) {
	setup := func() (cmd *cobra.Command, stdOut, stdErr *bytes.Buffer) {
		cmd = &cobra.Command{}
		stdOut, stdErr = &bytes.Buffer{}, &bytes.Buffer{}
		cmd.SetOut(stdOut)
		cmd.SetErr(stdErr)
		RegisterRemoteURLFlags(cmd.Flags())
		return
	}

	setEnv := func(t *testing.T, env, val string) {
		require.NoError(t, os.Setenv(env, val))
		t.Cleanup(func() {
			require.NoError(t, os.Unsetenv(EnvReadRemote))
		})
	}

	t.Run("case=prefers flag value", func(t *testing.T) {
		cmd, stdOut, stdErr := setup()

		expectedRemote := "ketotest.oryapis.com"
		require.NoError(t, cmd.Flags().Set(FlagReadRemote, "ketotest.oryapis.com"))
		setEnv(t, EnvReadRemote, "not"+expectedRemote)

		assert.Equal(t, expectedRemote, getRemote(cmd, FlagReadRemote, EnvReadRemote))
		assert.Equal(t, 0, stdOut.Len())
		assert.Equal(t, 0, stdErr.Len())
	})

	t.Run("case=uses env value", func(t *testing.T) {
		cmd, stdOut, stdErr := setup()

		expectedRemote := "ketotest.oryapis.com"
		setEnv(t, EnvReadRemote, expectedRemote)

		assert.Equal(t, expectedRemote, getRemote(cmd, FlagReadRemote, EnvReadRemote))
		assert.Equal(t, 0, stdOut.Len())
		assert.Equal(t, 0, stdErr.Len())
	})

	t.Run("case=falls back to flag default and prints warning", func(t *testing.T) {
		cmd, stdOut, stdErr := setup()

		expectedRemote := cmd.Flags().Lookup(FlagReadRemote).DefValue

		assert.Equal(t, expectedRemote, getRemote(cmd, FlagReadRemote, EnvReadRemote))
		assert.Equal(t, 0, stdOut.Len())
		assert.Contains(t, stdErr.String(), "falling back")
	})
}
