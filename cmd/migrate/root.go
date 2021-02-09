package migrate

import "github.com/spf13/cobra"

func newMigrateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "migrate",
	}
	cmd.AddCommand(
		newStatusCmd(),
		newUpCmd(),
		newDownCmd(),
	)
	return cmd
}

func RegisterCommandsRecursive(parent *cobra.Command) {
	parent.AddCommand(newMigrateCmd())
}
