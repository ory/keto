// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package expand

import (
	"fmt"

	"github.com/ory/keto/ketoapi"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/ory/x/flagx"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/client"
)

const FlagMaxDepth = "max-depth"

func NewExpandCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "expand <relation> <namespace> <object>",
		Short: "Expand a subject set",
		Long:  "Expand a subject set into a tree of subjects.",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := client.GetReadConn(cmd)
			if err != nil {
				return err
			}
			defer conn.Close()

			maxDepth, err := cmd.Flags().GetInt32(FlagMaxDepth)
			if err != nil {
				return err
			}

			cl := rts.NewExpandServiceClient(conn)
			resp, err := cl.Expand(cmd.Context(), &rts.ExpandRequest{
				Subject:  rts.NewSubjectSet(args[1], args[2], args[0]),
				MaxDepth: maxDepth,
			})
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Error making the request: %s\n", err.Error())
				return cmdx.FailSilently(cmd)
			}

			var tree *ketoapi.Tree[*ketoapi.RelationTuple]
			if resp.Tree != nil {
				tree = ketoapi.TreeFromProto[*ketoapi.RelationTuple](resp.Tree)
			}

			cmdx.PrintJSONAble(cmd, tree)
			switch flagx.MustGetString(cmd, cmdx.FlagFormat) {
			case string(cmdx.FormatDefault), "":
				if tree == nil && !flagx.MustGetBool(cmd, cmdx.FlagQuiet) {
					_, _ = fmt.Fprint(cmd.OutOrStdout(), "Got an empty tree. This probably means that the requested relation tuple is not present in Keto.")
				}
				_, _ = fmt.Fprintln(cmd.OutOrStdout())
			}
			return nil
		},
	}

	client.RegisterRemoteURLFlags(cmd.Flags())
	cmdx.RegisterJSONFormatFlags(cmd.Flags())
	cmdx.RegisterNoiseFlags(cmd.Flags())
	cmd.Flags().Int32P(FlagMaxDepth, "d", 0, "Maximum depth of the tree to be returned. If the value is less than 1 or greater than the global max-depth then the global max-depth will be used instead.")

	return cmd
}

func RegisterCommandsRecursive(parent *cobra.Command) {
	parent.AddCommand(NewExpandCmd())
}
