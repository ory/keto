package namespace

import (
	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/ory/keto/ketoctx"

	"github.com/ory/keto/cmd/client"
)

func NewNamespaceCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "namespace",
		Short: "Read and manipulate namespaces",
	}
}

func NewMigrateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Migrate a namespace",
	}
}

func RegisterCommandsRecursive(parent *cobra.Command, opts []ketoctx.Option) {
	rootCmd := NewNamespaceCmd()
	migrateCmd := NewMigrateCmd()
	migrateCmd.AddCommand(NewMigrateUpCmd(), NewMigrateDownCmd(), NewMigrateStatusCmd(), NewMigrateLegacyCmd(opts))

	rootCmd.AddCommand(migrateCmd, NewValidateCmd())

	parent.AddCommand(rootCmd)
}

func registerPackageFlags(flags *pflag.FlagSet) {
	client.RegisterRemoteURLFlags(flags)
	cmdx.RegisterFormatFlags(flags)
}
