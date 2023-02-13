// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ghodss/yaml"
	"github.com/ory/x/logrusx"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/namespace"
)

func TestNamespaceProvider(t *testing.T) {
	setup := func(t *testing.T, target string) (*NamespaceWatcher, *test.Hook) {
		ctx, cancel := context.WithCancel(context.Background())
		t.Cleanup(cancel)
		hook := &test.Hook{}
		l := logrusx.New("", "", logrusx.WithHook(hook))

		nw, err := NewNamespaceWatcher(ctx, l, target)
		require.NoError(t, err)

		return nw, hook
	}

	writeNamespace := func(t *testing.T, n *namespace.Namespace, fn string) {
		var marshal func(interface{}) ([]byte, error)
		switch ext := filepath.Ext(fn); ext {
		case ".yaml", ".yml":
			marshal = yaml.Marshal
		case ".toml":
			marshal = toml.Marshal
		case ".json":
			marshal = json.Marshal
		default:
			t.Logf("got unexpected file extension %s", ext)
			t.FailNow()
		}
		nEnc, err := marshal(n)
		require.NoError(t, err)
		require.NoError(t, os.WriteFile(fn, nEnc, 0600))
	}

	writeDir := func(t *testing.T, dir string, files map[string]interface{}) {
		for fn, fc := range files {
			fp := filepath.Join(dir, fn)
			require.NoError(t, os.MkdirAll(filepath.Dir(fp), 0700))

			if n, ok := fc.(*namespace.Namespace); ok {
				writeNamespace(t, n, fp)
				continue
			}

			s, ok := fc.(string)
			if !ok {
				t.Logf("expected file content to be string or *namespace.Namespace but got %T", fc)
				t.FailNow()
			}

			require.NoError(t, os.WriteFile(fp, []byte(s), 0600))
		}
	}

	writeJsonNamespace := func(t *testing.T) (string, *namespace.Namespace) {
		dir := t.TempDir()
		n := &namespace.Namespace{
			Name: "test namespace 1",
		}
		fn := filepath.Join(dir, "n.json")

		writeNamespace(t, n, fn)

		return fn, n
	}

	t.Run("case=loads JSON namespace file", func(t *testing.T) {
		fn, n := writeJsonNamespace(t)

		ws, _ := setup(t, "file://"+fn)

		nspaces, err := ws.Namespaces(context.Background())
		require.NoError(t, err)

		assert.Equal(t, []*namespace.Namespace{n}, nspaces)
	})

	t.Run("case=reads namespace files from directory", func(t *testing.T) {
		dir := t.TempDir()
		files := map[string]interface{}{
			"b.yml": &namespace.Namespace{
				Name: "b",
			},
			"a.toml": &namespace.Namespace{
				Name: "a",
			},
			"c.json": &namespace.Namespace{
				Name: "c",
			},
		}

		writeDir(t, dir, files)

		nw, _ := setup(t, "file://"+dir)

		nspaces, err := nw.Namespaces(context.Background())
		require.NoError(t, err)

		for _, n := range files {
			assert.Contains(t, nspaces, n.(*namespace.Namespace))
		}

		nsfs := nw.NamespaceFiles()
		assert.Equal(t, len(nspaces), len(nsfs))
		for _, n := range nsfs {
			assert.NotNil(t, n.Contents)
			assert.NotNil(t, n.Parser)
		}
	})

	t.Run("case=ignores but warns about unsupported file extensions", func(t *testing.T) {
		dir := t.TempDir()

		n := &namespace.Namespace{
			Name: "some name",
		}
		nJson, err := json.Marshal(n)
		require.NoError(t, err)

		writeDir(t, dir, map[string]interface{}{
			"unsupported.file": "foo bar\n",
			"supported.json":   string(nJson),
		})

		nw, hook := setup(t, "file://"+dir)

		require.Len(t, hook.Entries, 1)

		assert.Equal(t, logrus.WarnLevel, hook.Entries[0].Level, "%+v", hook.Entries[0])
		assert.True(t, strings.HasSuffix(hook.Entries[0].Data["file_name"].(string), "unsupported.file"))

		namespaces, err := nw.Namespaces(context.Background())
		require.NoError(t, err)

		require.Len(t, namespaces, 1)
		assert.Equal(t, n, namespaces[0])
	})

	t.Run("case=still returns successful namespace if one errors", func(t *testing.T) {
		dir := t.TempDir()

		n := &namespace.Namespace{
			Name: "some name",
		}
		nJson, err := json.Marshal(n)
		require.NoError(t, err)

		writeDir(t, dir, map[string]interface{}{
			"malformed.yml": "foo bar\n",
			"correct.json":  string(nJson),
		})

		nw, hook := setup(t, "file://"+dir)

		require.Len(t, hook.Entries, 1)

		assert.Equal(t, logrus.ErrorLevel, hook.Entries[0].Level, "%+v", hook.Entries[0])
		assert.True(t, strings.HasSuffix(hook.Entries[0].Data["file_name"].(string), "malformed.yml"))

		namespaces, err := nw.Namespaces(context.Background())
		require.NoError(t, err)

		require.Len(t, namespaces, 1)
		assert.Equal(t, n, namespaces[0])

		// files are included even if ns is unparsable
		nsfs := nw.NamespaceFiles()
		assert.Equal(t, 2, len(nsfs))
	})

	t.Run("method=should reload", func(t *testing.T) {
		nw := &NamespaceWatcher{
			target: "foo",
		}
		assert.False(t, nw.ShouldReload("foo"))
		assert.True(t, nw.ShouldReload("bar"))
		assert.True(t, nw.ShouldReload([]*namespace.Namespace{}))
	})
}
