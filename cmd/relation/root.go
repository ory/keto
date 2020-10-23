package relation

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/ory/x/cmdx"

	"github.com/ory/keto/cmd/client"
)

var relationCmd = &cobra.Command{
	Use: "relation",
}

var packageFlags = pflag.NewFlagSet("relation package flags", pflag.ContinueOnError)

func RegisterCommandRecursive(parent *cobra.Command) {
	parent.AddCommand(relationCmd)

	relationCmd.AddCommand(newGetCmd())
	relationCmd.AddCommand(newCreateCmd())
}

func init() {
	client.RegisterRemoteURLFlag(packageFlags)
	cmdx.RegisterFormatFlags(packageFlags)
}
