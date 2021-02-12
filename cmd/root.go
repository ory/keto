// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/configx"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/cmd/check"
	"github.com/ory/keto/cmd/expand"
	"github.com/ory/keto/cmd/migrate"
	"github.com/ory/keto/cmd/namespace"
	"github.com/ory/keto/cmd/relationtuple"
	"github.com/ory/keto/cmd/server"
	"github.com/ory/keto/cmd/status"
	"github.com/ory/keto/internal/driver/config"
)

// RootCmd represents the base command when called without any subcommands
func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "keto",
	}

	configx.RegisterConfigFlag(cmd.PersistentFlags(), []string{filepath.Join(userHomeDir(), "keto.yml")})

	relationtuple.RegisterCommandsRecursive(cmd)
	namespace.RegisterCommandsRecursive(cmd)
	migrate.RegisterCommandsRecursive(cmd)
	server.RegisterCommandsRecursive(cmd)
	check.RegisterCommandsRecursive(cmd)
	expand.RegisterCommandsRecursive(cmd)
	status.RegisterCommandRecursive(cmd)

	cmd.AddCommand(cmdx.Version(&config.Version, &config.Commit, &config.Date))

	return cmd
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := NewRootCmd().ExecuteContext(ctx); err != nil {
		if !errors.Is(err, cmdx.ErrNoPrintButFail) {
			fmt.Println(err)
		}
		os.Exit(-1)
	}
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

// CommandExecuter returns a cmdx.CommandExecutor that is initialized with the
// given config. The config can either be the string path to a config file or a
// map[string]interface{} specifying all values.
//
// Example test config:
//
// ports, err := freeport.GetFreePorts(2)
// require.NoError(t, err)
// cmd.CommandExecuter(t, map[string]interface{}{
//		"dsn":		 		"memory",
//		"log.level": 		"trace",
//		"serve.read.port":	ports[0],
//		"serve.write.port": ports[1],
//		"namespaces": 		[]map[string]interface{}{{
//			"name": "foo",
//			"id":	1,
// 		}},
// })
func CommandExecuter(t *testing.T, config interface{}) *cmdx.CommandExecuter {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	persistentArgs := []string{"-c", ""}

	switch c := config.(type) {
	case string:
		persistentArgs[1] = c
	case map[string]interface{}:
		fn := filepath.Join(t.TempDir(), "empty_keto.yml")
		_, err := os.Create(fn)
		require.NoError(t, err)
		persistentArgs[1] = fn

		opts := make([]configx.OptionModifier, 0, len(c))
		for k, v := range c {
			opts = append(opts, configx.WithValue(k, v))
		}

		ctx = configx.ContextWithConfigOptions(ctx, opts...)
	}

	return &cmdx.CommandExecuter{
		New:            NewRootCmd,
		Ctx:            ctx,
		PersistentArgs: persistentArgs,
	}
}
