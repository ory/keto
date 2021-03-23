package check

import (
	"fmt"

	"github.com/ory/keto/internal/check"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/client"
)

type checkOutput check.RESTResponse

func (o *checkOutput) String() string {
	if o.Allowed {
		return "Allowed\n"
	}
	return "Denied\n"
}

func newCheckCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check <subject> <relation> <namespace> <object>",
		Short: "Check whether a subject has a relation on an object",
		Long:  "Check whether a subject has a relation on an object. This method resolves subject sets and subject set rewrites.",
		Args:  cobra.ExactArgs(4),
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

			cmdx.PrintJSONAble(cmd, &checkOutput{Allowed: resp.Allowed})
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
