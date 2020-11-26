package configuration

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/ory/viper"
	"github.com/ory/x/logrusx"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/namespace"
)

func setupNamespaceTest(t *testing.T, files map[string]string) ([]*namespace.Namespace, *test.Hook) {
	dir, err := ioutil.TempDir("", "")
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, os.RemoveAll(dir))
	})

	for fn, fc := range files {
		require.NoError(t,
			ioutil.WriteFile(filepath.Join(dir, fn), []byte(fc), 0600),
		)
	}

	hook := &test.Hook{}
	config := NewViperProvider(logrusx.New("test-keto", "testing", logrusx.WithHook(hook)))

	viper.Set(ViperKeyNamespacePath, dir)

	namespaces := config.Namespaces()

	return namespaces, hook
}

func TestViperProvider(t *testing.T) {
	t.Run("case=reads namespace files", func(t *testing.T) {
		namespaces, hook := setupNamespaceTest(t, map[string]string{
			"videos.yml": `
name: videos
id: 1
`,
		})

		require.Len(t, hook.Entries, 0)

		assert.Len(t, namespaces, 1)
		assert.Equal(t, &namespace.Namespace{
			ID:   1,
			Name: "videos",
		}, namespaces[0])
	})

	t.Run("case=ignores but warns about unsupported file extensions", func(t *testing.T) {
		namespaces, hook := setupNamespaceTest(t, map[string]string{
			"unsupported.file": "foo bar\n",
			"supported.json": `{
	"name": "namespace name",
	"id": 2
}
`,
		})

		require.Len(t, hook.Entries, 1)

		assert.Equal(t, logrus.InfoLevel, hook.Entries[0].Level)
		assert.Contains(t, hook.Entries[0].Message, "unsupported.file")

		assert.Len(t, namespaces, 1)
		assert.Equal(t, &namespace.Namespace{
			Name: "namespace name",
			ID:   2,
		}, namespaces[0])
	})

	t.Run("case=still returns successful namespace if one errors", func(t *testing.T) {
		namespaces, hook := setupNamespaceTest(t, map[string]string{
			"malformed.yml": "foo bar\n",
			"correct.yml":   "id: 1\nname: some name\n",
		})

		require.Len(t, hook.Entries, 1)

		assert.Equal(t, logrus.ErrorLevel, hook.Entries[0].Level)
		assert.Contains(t, hook.Entries[0].Message, "malformed.yml")

		assert.Len(t, namespaces, 1)
		assert.Equal(t, &namespace.Namespace{
			Name: "some name",
			ID:   1,
		}, namespaces[0])
	})
}
