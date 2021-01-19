package check

import (
	"fmt"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"

	acl "github.com/ory/keto/api/keto/acl/v1alpha1"
	"github.com/ory/keto/cmd/client"
)

func newCheckCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "check <subject> <relation> <namespace> <object>",
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := client.GetReadConn(cmd)
			if err != nil {
				return err
			}
			defer conn.Close()

			cl := acl.NewCheckServiceClient(conn)
			resp, err := cl.Check(cmd.Context(), &acl.CheckRequest{
				Subject: &acl.Subject{
					Ref: &acl.Subject_Id{Id: args[0]},
				},
				Relation:  args[1],
				Namespace: args[2],
				Object:    args[3],
			})
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not make request: %s\n", err)
				return err
			}

			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%v", resp.Allowed)
			return nil
		},
	}

	client.RegisterRemoteURLFlags(cmd.Flags())
	cmdx.RegisterFormatFlags(cmd.Flags())

	return cmd
}

func RegisterCommandsRecursive(parent *cobra.Command) {
	parent.AddCommand(newCheckCmd())
}
