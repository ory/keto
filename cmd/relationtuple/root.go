// Copyright © 2022 Ory Corp

package relationtuple

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/ory/keto/cmd/client"

	"github.com/ory/x/cmdx"
)

func newRelationCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "relation-tuple",
		Short: "Read and manipulate relation tuples",
	}
}

func RegisterCommandsRecursive(parent *cobra.Command) {
	relationCmd := newRelationCmd()

	parent.AddCommand(relationCmd)

	relationCmd.AddCommand(NewGetCmd(), NewCreateCmd(), NewDeleteCmd(), NewDeleteAllCmd(), NewParseCmd())
}

func registerPackageFlags(flags *pflag.FlagSet) {
	client.RegisterRemoteURLFlags(flags)
	cmdx.RegisterFormatFlags(flags)
}
