// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package migrate

import (
	"github.com/spf13/cobra"

	"github.com/ory/keto/ketoctx"
)

func newMigrateCmd(opts []ketoctx.Option) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Commands to migrate the database",
		Long: "Commands to migrate the database.\n" +
			"This does not affect namespaces. Use `keto namespace migrate` for migrating namespaces.",
	}
	cmd.AddCommand(
		newStatusCmd(opts),
		newUpCmd(opts),
		newDownCmd(opts),
	)
	return cmd
}

func RegisterCommandsRecursive(parent *cobra.Command, opts []ketoctx.Option) {
	parent.AddCommand(newMigrateCmd(opts))
}
