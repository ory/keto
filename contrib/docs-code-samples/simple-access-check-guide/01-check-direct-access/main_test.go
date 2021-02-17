// +build docscodesamples

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	// capture errors from main()
	defer func() {
		require.Nil(t, recover())
	}()

	f, err := os.Create(filepath.Join(t.TempDir(), "mock_output"))
	require.NoError(t, err)

	os.Stdout, f = f, os.Stdout
	main()
	os.Stdout, f = f, os.Stdout

	out, err := ioutil.ReadFile(f.Name())
	require.NoError(t, err)
	_, _ = os.Stdout.Write(out)

	assert.Equal(t, string(out), "Allowed: true\n")
}
