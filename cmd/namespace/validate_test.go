// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package namespace

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/configx"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const fileMode = 0644

const (
	configEmbeddedYaml = `
dsn: memory
namespaces:
  - name: testns0
    id: 0
  - name: testns1
    id: 1
`

	configEmbeddedJson = `{"dsn":"memory","namespaces":[{"name":"testns0","id":0}]}`

	configEmbeddedToml = `
[[namespaces]]
name = "testns0"
id = 0

[[namespaces]]
name = "testns1"
id = 1
`
)

const configReference = `
dsn: memory
namespaces: %s
`

const nsyaml = "name: testns0\nid: 0"
const nsjson = `{"name": "testns1", "id": 1}`

func TestValidateConfigNamespaces(t *testing.T) {
	cmd := cmdx.CommandExecuter{New: validateCommand}

	t.Run("case=read valid config with embedded namespaces", func(t *testing.T) {
		dir := t.TempDir()
		fnyaml := filepath.Join(dir, "keto.yaml")
		require.NoError(t, os.WriteFile(fnyaml, []byte(configEmbeddedYaml), fileMode))

		fnjson := filepath.Join(dir, "keto.json")
		require.NoError(t, os.WriteFile(fnjson, []byte(configEmbeddedJson), fileMode))

		fntoml := filepath.Join(dir, "keto.toml")
		require.NoError(t, os.WriteFile(fntoml, []byte(configEmbeddedToml), fileMode))

		stdOut := cmd.ExecNoErr(t, "validate", "-c", fnyaml+","+fnjson+","+fntoml)
		assert.Contains(t, stdOut, "Congrats, all files are valid!\n")
	})

	t.Run("case=supports 3 namespace file formats", func(t *testing.T) {
		dir := t.TempDir()
		nsfiles := map[string]string{
			filepath.Join(dir, "ns.yaml"): "name: testns0\nid: 0",
			filepath.Join(dir, "ns.json"): "{\"name\": \"testns0\",\"id\": 0}",
			filepath.Join(dir, "ns.toml"): "name = \"testns0\"\nid = 0",
		}
		for fn, contents := range nsfiles {
			require.NoError(t, os.WriteFile(fn, []byte(contents), fileMode))
		}

		params := append([]string{"validate"}, keys(nsfiles)...)
		cmd.ExecNoErr(t, params...)
	})

	t.Run("case=unknown namespace format gives error", func(t *testing.T) {
		fn := filepath.Join(t.TempDir(), "ns.txt")
		require.NoError(t, os.WriteFile(fn, []byte("name: ns\nid: 0"), fileMode))

		_, stdErr, err := cmd.Exec(nil, "validate", fn)
		require.ErrorIs(t, err, cmdx.ErrNoPrintButFail)
		assert.Contains(t, stdErr, "Unable to infer file type")
	})

	t.Run("case=config passed as varg fails", func(t *testing.T) {
		fn := filepath.Join(t.TempDir(), "keto.yaml")
		require.NoError(t, os.WriteFile(fn, []byte(configEmbeddedYaml), fileMode))

		// interprets config file as namespace file when `-c` flag is not passed
		_, stdErr, err := cmd.Exec(nil, "validate", fn)
		require.ErrorIs(t, err, cmdx.ErrNoPrintButFail)
		assert.Regexp(t, "additionalProperties ((\"namespaces\", \"dsn\")|(\"dsn\", \"namespaces\")) not allowed", stdErr)
	})

	t.Run("case=read config with invalid embedded namespace", func(t *testing.T) {
		fn := filepath.Join(t.TempDir(), "keto.yaml")
		require.NoError(t, os.WriteFile(fn, []byte(`{"namespaces":[{"id":"x","name":"x"}]}`), fileMode))

		_, stdErr, err := cmd.Exec(nil, "validate", "--config", fn)
		require.ErrorIs(t, err, cmdx.ErrNoPrintButFail)
		assert.Contains(t, stdErr, "not a valid namespace file")
		assert.Contains(t, stdErr, "id:")
	})

	t.Run("case=read config with namespace as dir reference", func(t *testing.T) {
		nsdir := t.TempDir()
		fn := filepath.Join(t.TempDir(), "config.yaml")
		require.NoError(t, os.WriteFile(filepath.Join(nsdir, "ns0.yaml"), []byte(nsyaml), fileMode))
		require.NoError(t, os.WriteFile(filepath.Join(nsdir, "ns1.json"), []byte(nsjson), fileMode))
		require.NoError(t, os.WriteFile(fn, []byte(fmt.Sprintf(configReference, nsdir)), fileMode))

		cmd.PersistentArgs = []string{}
		stdOut := cmd.ExecNoErr(t, "validate", "-c", fn)
		assert.Contains(t, stdOut, "Congrats, all files are valid!\n")

		stdOut = cmd.ExecNoErr(t, "validate", "-c", fmt.Sprintf("%s,%s", fn, fn))
		assert.Contains(t, stdOut, "Congrats, all files are valid!\n")
	})

	t.Run("case=read config with dir reference bad namespaces", func(t *testing.T) {
		nsdir := t.TempDir()
		fn := filepath.Join(t.TempDir(), "config.yaml")
		require.NoError(t, os.WriteFile(filepath.Join(nsdir, "ns0.yaml"), []byte("name: badId\nid: foo"), fileMode))
		require.NoError(t, os.WriteFile(fn, []byte(fmt.Sprintf(configReference, nsdir)), fileMode))

		cmd.PersistentArgs = []string{}
		_, stdErr, err := cmd.Exec(nil, "validate", "-c", fn)
		require.ErrorIs(t, err, cmdx.ErrNoPrintButFail)
		assert.Contains(t, stdErr, "contains values or keys which are invalid")
	})
}

func validateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keto",
		Short: "Global and consistent permission and authorization server",
	}
	configx.RegisterConfigFlag(cmd.PersistentFlags(), []string{})
	cmd.AddCommand(NewValidateCmd())
	return cmd
}

func keys(m map[string]string) []string {
	rv := make([]string, 0, len(m))
	for k := range m {
		rv = append(rv, k)
	}
	return rv
}
