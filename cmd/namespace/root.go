package namespace

import (
	"github.com/spf13/cobra"

	"github.com/ory/keto/ketoctx"
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
