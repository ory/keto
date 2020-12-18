package migrate

import "github.com/spf13/cobra"

func newMigrateCmd() *cobra.Command {
	return &cobra.Command{
		Use: "migrate",
	}
}

func RegisterCommandsRecursive(parent *cobra.Command) {
	migrateCmd := newMigrateCmd()

	migrateCmd.AddCommand(newStatusCmd(), newUpCmd())

	parent.AddCommand(migrateCmd)
}
