package namespace

import (
	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/ory/keto/cmd/client"
)

func NewNamespaceCmd() *cobra.Command {
	return &cobra.Command{
		Use: "namespace",
	}
}

func NewMigrateCmd() *cobra.Command {
	return &cobra.Command{
		Use: "migrate",
	}
}

func RegisterCommandsRecursive(parent *cobra.Command) {
	rootCmd := NewNamespaceCmd()
	migrateCmd := NewMigrateCmd()
	migrateCmd.AddCommand(NewMigrateUpCmd(), NewMigrateDownCmd())

	rootCmd.AddCommand(migrateCmd, NewValidateCmd())

	parent.AddCommand(rootCmd)
}

func registerPackageFlags(flags *pflag.FlagSet) {
	client.RegisterRemoteURLFlags(flags)
	cmdx.RegisterFormatFlags(flags)
}
