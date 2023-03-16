// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gobuffalo/httptest"
	"github.com/ory/x/configx"
	"github.com/ory/x/logrusx"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/embedx"
	"github.com/ory/keto/internal/namespace"
)

// createFile writes the content to a temporary file, returning the path.
// Good for testing config files.
func createFile(t *testing.T, content string) (path string) {
	t.Helper()

	f, err := os.CreateTemp(t.TempDir(), "config-*.yaml")
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = os.Remove(f.Name()) })

	n, err := f.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(content) {
		t.Fatal("failed to write the complete content")
	}

	return f.Name()
}

// createFileF writes the content format string with the applied args to a
// temporary file, returning the path. Good for testing config files.
func createFileF(t *testing.T, contentF string, args ...any) (path string) {
	return createFile(t, fmt.Sprintf(contentF, args...))
}

func setup(t *testing.T, configFile string) (*test.Hook, *Config) {
	t.Helper()
	hook := test.Hook{}
	l := logrusx.New("test", "today", logrusx.WithHook(&hook))

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	t.Cleanup(cancel)

	config, err := NewDefault(
		ctx,
		pflag.NewFlagSet("test", pflag.ContinueOnError),
		l,
		configx.WithConfigFiles(configFile),
	)
	require.NoError(t, err)

	return &hook, config
}

func assertNamespaces(t *testing.T, p *Config, nn ...*namespace.Namespace) {
	t.Helper()

	nm, err := p.NamespaceManager()
	require.NoError(t, err)
	actualNamespaces, err := nm.Namespaces(context.Background())
	require.NoError(t, err)
	assert.Equal(t, len(nn), len(actualNamespaces))
	assert.ElementsMatch(t, nn, actualNamespaces)
}

// The new way to configure namespaces is through the Ory Permissions Language.
// We check here that we still support enumerating the namespaces directly in
// the config or through a file reference, in which case there should be no
// rewrites configured.
func TestLegacyNamespaceConfig(t *testing.T) {
	t.Run("case=creates memory namespace manager when namespaces are set", func(t *testing.T) {
		config := createFile(t, `
dsn: memory
namespaces:
  - name: n0
  - name: n1
  - name: n2`)

		run := func(namespaces []*namespace.Namespace) func(*testing.T) {
			return func(t *testing.T) {
				_, p := setup(t, config)

				assertNamespaces(t, p, namespaces...)

				nm, err := p.NamespaceManager()
				require.NoError(t, err)
				_, ok := nm.(*memoryNamespaceManager)
				assert.True(t, ok)
				assert.False(t, p.StrictMode())
			}

		}

		nn := []*namespace.Namespace{
			{Name: "n0"},
			{Name: "n1"},
			{Name: "n2"},
		}
		nnJson, err := json.Marshal(nn)
		require.NoError(t, err)
		nnValue := make([]interface{}, 0)
		require.NoError(t, json.Unmarshal(nnJson, &nnValue))

		t.Run(
			"type=[]*namespace.Namespace",
			run(nn),
		)
	})

	t.Run("case=reloads namespace manager when namespaces are updated using Set()", func(t *testing.T) {
		_, p := setup(t, createFile(t, "dsn: memory"))

		n0 := &namespace.Namespace{
			Name: "n0",
		}
		n1 := &namespace.Namespace{
			Name: "n1",
		}

		require.NoError(t, p.Set(KeyNamespaces, []*namespace.Namespace{n0}))
		assertNamespaces(t, p, n0)

		require.NoError(t, p.Set(KeyNamespaces, []*namespace.Namespace{n1}))
		assertNamespaces(t, p, n1)
	})

	t.Run("case=creates watcher manager when namespaces is string URL", func(t *testing.T) {
		_, p := setup(t, createFileF(t, `
dsn: memory
namespaces: file://%s`,
			t.TempDir()))

		nm, err := p.NamespaceManager()
		require.NoError(t, err)
		_, ok := nm.(*NamespaceWatcher)
		assert.True(t, ok)
		assert.False(t, p.StrictMode())
	})

	t.Run("case=uses passed configx provider", func(t *testing.T) {
		ctx := context.Background()
		cp, err := configx.New(ctx, embedx.ConfigSchema, configx.WithValue(KeyDSN, "foobar"))
		require.NoError(t, err)

		p := New(ctx, logrusx.New("test", "today"), cp)
		assert.Equal(t, "foobar", p.DSN())
		assert.Same(t, cp, p.p)
	})
}

func TestInvalidNamespaceConfig(t *testing.T) {
	config := createFile(t, `
dsn: memory
namespaces:
  - name: n0
    id: 2147483648
`)
	hook := test.Hook{}
	l := logrusx.New("test", "today", logrusx.WithHook(&hook))
	_, err := NewDefault(
		context.Background(),
		pflag.NewFlagSet("test", pflag.ContinueOnError),
		l,
		configx.WithConfigFiles(config),
	)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "must be <= 2.147483647e+09 but found 2147483648")
}

// Test that the namespaces can be configured through the Ory Permission
// Language.
func TestRewritesNamespaceConfig(t *testing.T) {
	oplContent := `
class User implements Namespace {
  related: {
    manager: User[]
  }
}
  
class Group implements Namespace {
  related: {
    members: (User | Group)[]
  }
}`

	oplConfigFile := createFile(t, oplContent)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte(oplContent))
		if err != nil {
			t.Fatal(err)
		}
	}))
	t.Cleanup(func() { srv.Close() })

	cases := []struct {
		name                    string
		location                string
		disallowPrivateIPRanges bool
		expectErr               bool
	}{{
		name:     "local file",
		location: "file://" + oplConfigFile,
	}, {
		name:                    "HTTP url forbidden",
		location:                srv.URL,
		disallowPrivateIPRanges: true,
		expectErr:               true,
	}, {
		name:     "HTTP url allowed",
		location: srv.URL,
	}}

	for _, tc := range cases {
		t.Run("case="+tc.name, func(t *testing.T) {
			config := createFileF(t, `
dsn: memory
clients:
  http:
    disallow_private_ip_ranges: %v
namespaces:
  location: %s`, tc.disallowPrivateIPRanges, tc.location)

			_, p := setup(t, config)
			nm, err := p.NamespaceManager()
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			namespaces, err := nm.Namespaces(context.Background())
			require.NoError(t, err)
			require.Len(t, namespaces, 2)

			names, relationNames := []string{namespaces[0].Name, namespaces[1].Name}, []string{namespaces[0].Relations[0].Name, namespaces[1].Relations[0].Name}

			assert.False(t, p.StrictMode())
			assert.ElementsMatch(t, names, []string{"User", "Group"})
			assert.ElementsMatch(t, relationNames, []string{"manager", "members"})
		})
	}

	t.Run("case=strict_mode=true", func(t *testing.T) {
		config := createFileF(t, `
dsn: memory
namespaces:
  location: file://%s
  experimental_strict_mode: true`, oplConfigFile)

		_, p := setup(t, config)
		assert.True(t, p.StrictMode())
	})
}

func TestProvider_DefaultReadAPIListenOn(t *testing.T) {
	ctx := context.Background()
	config, err := NewDefault(
		ctx,
		pflag.NewFlagSet("test", pflag.ContinueOnError),
		logrusx.New("", ""),
		configx.WithValue("dsn", "foo"),
	)
	require.NoError(t, err)

	assert.Equal(t, ":4466", config.ReadAPIListenOn())
}
