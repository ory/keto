package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ory/x/healthx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	keto "github.com/ory/keto/cmd"
)

func TestExample(t *testing.T) {
	c := keto.CommandExecuter(t, "../keto.yml")

	serverErr := make(chan error)
	go func() {
		stdOut, stdErr, err := c.Exec(nil, "serve")
		t.Logf("server StdOut: %s\n\nserver StdErr: %s\n\n", stdOut, stdErr)
		serverErr <- err
	}()

	// capture errors from main()
	defer func() {
		require.Nil(t, recover())
	}()

	var healthReady = func() bool {
		resp, err := http.DefaultClient.Get("http://localhost:4466" + healthx.ReadyCheckPath)
		if err != nil {
			return false
		}
		return resp.StatusCode == http.StatusOK
	}
	// wait for /health/ready
	for !healthReady() {
		select {
		case <-time.After(10 * time.Millisecond):
		case err := <-serverErr:
			require.NoError(t, err)
		}
	}

	f, err := os.Create(filepath.Join(t.TempDir(), "mock_output"))
	require.NoError(t, err)

	os.Stdout, f = f, os.Stdout
	main()
	os.Stdout, f = f, os.Stdout

	out, err := ioutil.ReadFile(f.Name())
	require.NoError(t, err)
	_, _ = os.Stdout.Write(out)

	assert.Equal(t, string(out), "Successfully created tuple.\n")
}
