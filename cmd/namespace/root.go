package namespace

import (
	"github.com/ory/keto/ketoctx"
	"github.com/spf13/cobra"
)

func NewNamespaceCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "namespace",
		Short: "Read and manipulate namespaces",
	}
}

func RegisterCommandsRecursive(parent *cobra.Command, _ []ketoctx.Option) {
	rootCmd := NewNamespaceCmd()
	rootCmd.AddCommand(NewValidateCmd())

	parent.AddCommand(rootCmd)
}
