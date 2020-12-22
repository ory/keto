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

	"github.com/ory/keto/cmd/check"

	"github.com/ory/keto/cmd/server"
	"github.com/ory/keto/internal/driver/config"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/configx"

	"github.com/ory/keto/cmd/migrate"
	"github.com/ory/keto/cmd/namespace"
	"github.com/ory/keto/cmd/relationtuple"

	"github.com/spf13/cobra"
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
