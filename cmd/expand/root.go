package expand

import (
	"fmt"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"

	acl "github.com/ory/keto/api/keto/acl/v1alpha1"
	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/internal/expand"
)

const FlagMaxDepth = "max-depth"

func NewExpandCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "expand <relation> <namespace> <object>",
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := client.GetBasicConn(cmd)
			if err != nil {
				return nil
			}
			defer conn.Close()

			maxDepth, err := cmd.Flags().GetInt32(FlagMaxDepth)
			if err != nil {
				return err
			}

			cl := acl.NewExpandServiceClient(conn)
			resp, err := cl.Expand(cmd.Context(), &acl.ExpandRequest{
				Subject: &acl.Subject{
					Ref: &acl.Subject_Set{
						Set: &acl.SubjectSet{
							Relation:  args[0],
							Namespace: args[1],
							Object:    args[2],
						},
					},
				},
				MaxDepth: maxDepth,
			})
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Error making the request: %s\n", err.Error())
				return cmdx.FailSilently(cmd)
			}

			cmdx.PrintJSONAble(cmd, expand.TreeFromProto(resp.Tree))
			return nil
		},
	}

	client.RegisterRemoteURLFlags(cmd.Flags())
	cmdx.RegisterJSONFormatFlags(cmd.Flags())
	cmd.Flags().Int32P(FlagMaxDepth, "d", 100, "maximum depth of the tree")

	return cmd
}

func RegisterCommandsRecursive(parent *cobra.Command) {
	parent.AddCommand(NewExpandCmd())
}
