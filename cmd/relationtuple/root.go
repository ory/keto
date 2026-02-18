// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/ory/x/cmdx"

	"github.com/ory/keto/cmd/client"
)

func newRelationCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "relation-tuple",
		Short: "Read and manipulate relationships",
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

func registerFileFlag(flags *pflag.FlagSet) {
	flags.StringSliceP(FlagFile, "f", nil, "Read relationships from JSON files or directories (use the special filename - for stdin)")
}
