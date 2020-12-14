package migrate

import "github.com/spf13/cobra"

var migrateCmd = &cobra.Command{
	Use: "migrate",
}

func RegisterCommandRecursive(parent *cobra.Command) {
	migrateCmd.AddCommand(newStatusCmd(), newUpCmd())

	parent.AddCommand(migrateCmd)
}
