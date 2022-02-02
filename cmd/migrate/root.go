package migrate

import "github.com/spf13/cobra"

func newMigrateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Commands to migrate the database",
		Long: "Commands to migrate the database.\n" +
			"This does not affect namespaces. Use `keto namespace migrate` for migrating namespaces.",
	}
	cmd.AddCommand(
		newStatusCmd(),
		newUpCmd(),
		newDownCmd(),
		newMigrateUUIDMappingCmd(),
	)
	return cmd
}

func RegisterCommandsRecursive(parent *cobra.Command) {
	parent.AddCommand(newMigrateCmd())
}
