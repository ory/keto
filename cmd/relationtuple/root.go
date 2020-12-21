package relationtuple

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/ory/x/cmdx"

	"github.com/ory/keto/cmd/client"
)

func newRelationCmd() *cobra.Command {
	return &cobra.Command{
		Use: "relation-tuple",
	}
}

var packageFlags = pflag.NewFlagSet("relation package flags", pflag.ContinueOnError)

func RegisterCommandsRecursive(parent *cobra.Command) {
	relationCmd := newRelationCmd()

	parent.AddCommand(relationCmd)

	relationCmd.AddCommand(newGetCmd())
	relationCmd.AddCommand(newCreateCmd())
}

func init() {
	client.RegisterRemoteURLFlag(packageFlags)
	cmdx.RegisterFormatFlags(packageFlags)
}
