// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ory/keto/ketoctx"

	"github.com/ory/keto/cmd/status"

	"github.com/ory/keto/cmd/expand"

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
func NewRootCmd(opts ...ketoctx.Option) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keto",
		Short: "Global and consistent permission and authorization server",
	}

	configx.RegisterConfigFlag(cmd.PersistentFlags(), []string{filepath.Join(userHomeDir(), "keto.yml")})

	relationtuple.RegisterCommandsRecursive(cmd)
	namespace.RegisterCommandsRecursive(cmd, opts)
	migrate.RegisterCommandsRecursive(cmd, opts)
	server.RegisterCommandsRecursive(cmd, opts)
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
