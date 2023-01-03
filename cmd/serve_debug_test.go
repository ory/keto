// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/spf13/cobra"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/configx"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/internal/driver"
)

// This test is intended for debug purposes. Just comment out t.SkipNow() and run it through your IDE debugger with -tags sqlite
func Test_DebugOnly(t *testing.T) {
	t.SkipNow()

	ctx := context.Background()
	err := cmdx.ExecBackgroundCtx(
		ctx,
		NewRootCmd(),
		os.Stdin,
		os.Stdout,
		os.Stderr,
		"serve",
		"-c",
		"../contrib/cat-videos-example/keto.yml",
	).Wait()

	t.Logf("got error %s", err)
}

// This benchmark is intended for profiling specific operations
//
// Summary namespace manager reload:
//
//	Memory manager:
//	  Mem: 4.35% of all while github.com/ory/x/configx.(*Provider).reload takes 93.56%
//	  CPU: 1.86% of all while github.com/ory/x/configx.(*Provider).reload takes 67.28%
//	File manager:
//	  Mem: 0.83% of all while github.com/ory/x/configx.(*Provider).reload takes 96.34%
//	  CPU: 3.84% of all while github.com/ory/x/configx.(*Provider).reload takes 61.53%
func BenchmarkServe(b *testing.B) {
	b.SkipNow()

	b.StopTimer()
	h := test.Hook{}

	rCtx, cancel := context.WithCancel(context.WithValue(context.Background(),
		driver.LogrusHookContextKey, &h))
	defer cancel()

	f, err := os.Create(filepath.Join(b.TempDir(), "keto.yml"))
	require.NoError(b, err)
	_, err = f.WriteString(`
log:
  level: debug

namespaces:
  - id: 0
    name: videos

dsn: memory`)
	require.NoError(b, err)
	require.NoError(b, f.Sync())

	flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
	configx.RegisterFlags(flags)
	require.NoError(b, flags.Set(configx.FlagConfig, f.Name()))

	reg, err := driver.NewDefaultRegistry(rCtx, flags, false, nil)

	// setting env vars instead of flags bc flags are not defined on every command
	require.NoError(b, os.Setenv(client.EnvReadRemote, "127.0.0.1:4466"))
	require.NoError(b, os.Setenv(client.EnvWriteRemote, "127.0.0.1:4467"))

	c := cmdx.CommandExecuter{
		New: func() *cobra.Command {
			return NewRootCmd(nil)
		},
		Ctx: context.WithValue(rCtx, driver.RegistryContextKey, reg),
	}
	stdOut, stdErr := &bytes.Buffer{}, &bytes.Buffer{}
	eg := c.ExecBackground(nil, stdOut, stdErr, "serve", "-c", f.Name())
	b.Cleanup(func() {
		require.NoError(b, eg.Wait())
	})
	require.Equal(b, "SERVING\n", c.ExecNoErr(b, "status", "-b"))
	time.Sleep(100 * time.Millisecond)

	b.StartTimer()

	b.Run("memory namespace manager", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err = f.Seek(0, io.SeekStart)
			require.NoError(b, err)
			_, err = fmt.Fprintf(f, `
log:
  level: debug

dsn: memory

namespaces:
  - id: %d
    name: videos`, i)
			require.NoError(b, err)

			var le *logrus.Entry

			for i := 0; i < 1000; i++ {
				le = h.LastEntry()
				if le != nil && le.Message == "Configuration change processed successfully." {
					break
				}
				time.Sleep(1 * time.Millisecond)
			}
			require.NotNil(b, le, "iteration %d", i)
			require.Equal(b, "Configuration change processed successfully.", le.Message)

			h.Reset()
		}
	})

	b.Run("file namespace manager", func(b *testing.B) {
		nn := []string{filepath.Join(b.TempDir(), "namespace.yml"), filepath.Join(b.TempDir(), "namespace.yml")}
		for _, n := range nn {
			require.NoError(b, os.WriteFile(n, []byte(`
id: 0
name: foo`), 0600))
		}

		for i := 0; i < b.N; i++ {
			_, err = f.Seek(0, io.SeekStart)
			require.NoError(b, err)
			_, err = fmt.Fprintf(f, `
log:
  level: debug

dsn: memory

namespaces: file://%s

# foo bar this is only for ensuring the old values are not around`, nn[i%2])
			require.NoError(b, err)
			require.NoError(b, f.Sync())

			var le *logrus.Entry

			for i := 0; i < 1000; i++ {
				le = h.LastEntry()
				if le != nil && le.Message == "Configuration change processed successfully." {
					break
				}
				time.Sleep(1 * time.Millisecond)
			}
			require.NotNil(b, le)
			require.Equal(b, "Configuration change processed successfully.", le.Message)

			h.Reset()
		}
	})
}
