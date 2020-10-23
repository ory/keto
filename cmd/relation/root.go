package relation

import (
	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/client"
)

var relationCmd = &cobra.Command{
	Use: "relation",
}

func RegisterCommandRecursive(parent *cobra.Command) {
	parent.AddCommand(relationCmd)

	relationCmd.AddCommand(getByUserRelationCmd)

	client.RegisterRemoteURLFlag(relationCmd.PersistentFlags())
}
