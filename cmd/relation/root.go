package relation

import (
	"github.com/ory/keto/cmd/client"
	"github.com/spf13/cobra"
)

var relationCmd = &cobra.Command{
	Use: "relation",
}

func RegisterCommandRecursive(parent *cobra.Command) {
	parent.AddCommand(relationCmd)

	relationCmd.AddCommand(getByUserRelationCmd)

	client.RegisterRemoteURLFlag(relationCmd.PersistentFlags())
}
