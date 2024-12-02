// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package migrate

import (
	"github.com/ory/x/configx"
	"github.com/ory/x/popx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver"

	"github.com/ory/keto/ketoctx"
)

func RegisterCommandsRecursive(parent *cobra.Command, opts []ketoctx.Option) {
	parent.AddCommand(newMigrateCmd(opts))
}

// migrateSqlCmd represents the sql command
func newMigrateCmd(opts []ketoctx.Option) *cobra.Command {
	c := &cobra.Command{
		Use:   "migrate",
		Short: "Create SQL schemas and apply migration plans",
	}

	configx.RegisterFlags(c.PersistentFlags())
	c.PersistentFlags().BoolP("read-from-env", "e", false, "If set, reads the database connection string from the environment variable DSN or config file key dsn.")
	c.Flags().BoolP("yes", "y", false, "If set all confirmation requests are accepted without user interaction.")

	c.AddCommand(NewMigrateSQLStatusCmd(opts))
	c.AddCommand(NewMigrateSQLUpCmd(opts))
	c.AddCommand(NewMigrateSQLDownCmd(opts))

	c.AddCommand(newMigrateSqlCmd(opts))

	return c
}

// migrateSqlCmd represents the sql command
func newMigrateSqlCmd(opts []ketoctx.Option) *cobra.Command {
	c := &cobra.Command{
		Use:   "sql",
		Short: "Create SQL schemas and apply migration plans",
	}

	configx.RegisterFlags(c.PersistentFlags())
	c.PersistentFlags().BoolP("read-from-env", "e", false, "If set, reads the database connection string from the environment variable DSN or config file key dsn.")
	c.Flags().BoolP("yes", "y", false, "If set all confirmation requests are accepted without user interaction.")

	c.AddCommand(NewMigrateSQLStatusCmd(opts))
	c.AddCommand(NewMigrateSQLUpCmd(opts))
	c.AddCommand(NewMigrateSQLDownCmd(opts))

	return c
}

func NewMigrateSQLDownCmd(opts []ketoctx.Option) *cobra.Command {
	return popx.NewMigrateSQLDownCmd("keto", func(cmd *cobra.Command, args []string) error {
		reg, err := driver.NewDefaultRegistry(cmd.Context(), cmd.Flags(), true, opts)
		if err != nil {
			return err
		}

		return popx.MigrateSQLDown(cmd, reg.Persister())
	})
}

func NewMigrateSQLUpCmd(opts []ketoctx.Option) *cobra.Command {
	return popx.NewMigrateSQLUpCmd("keto", func(cmd *cobra.Command, args []string) error {
		reg, err := driver.NewDefaultRegistry(cmd.Context(), cmd.Flags(), true, opts)
		if err != nil {
			return err
		}

		return popx.MigrateSQLUp(cmd, reg.Persister())
	})
}

func NewMigrateSQLStatusCmd(opts []ketoctx.Option) *cobra.Command {
	return popx.NewMigrateSQLStatusCmd("keto", func(cmd *cobra.Command, args []string) error {
		reg, err := driver.NewDefaultRegistry(cmd.Context(), cmd.Flags(), true, opts)
		if err != nil {
			return err
		}

		return popx.MigrateStatus(cmd, reg.Persister())
	})
}
