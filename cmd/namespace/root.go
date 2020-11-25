package namespace

import (
	"github.com/ory/keto/cmd/client"
	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewNamespaceCmd() *cobra.Command {
	return &cobra.Command{
		Use: "namespace",
	}
}

func RegisterCommandsRecursive(parent *cobra.Command) {
	rootCmd := NewNamespaceCmd()
	rootCmd.AddCommand(NewMigrateCmd())

	parent.AddCommand(rootCmd)
}

func registerPackageFlags(flags *pflag.FlagSet) {
	client.RegisterRemoteURLFlag(flags)
	cmdx.RegisterFormatFlags(flags)
}
