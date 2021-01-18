package relationtuple

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	acl "github.com/ory/keto/api/keto/acl/v1alpha1"

	"github.com/ory/keto/internal/relationtuple"

	"github.com/spf13/cobra"

	"github.com/ory/x/cmdx"

	"github.com/ory/keto/cmd/client"
)

func newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create <relation-tuple.json>",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := client.GetPrivilegedConn(cmd)
			if err != nil {
				return err
			}

			var f io.Reader
			if args[0] == "-" {
				f = cmd.InOrStdin()
			} else {
				f, err = os.Open(args[0])
				if err != nil {
					_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could open %s: %s\n", args[0], err)
					return cmdx.FailSilently(cmd)
				}
			}

			var r relationtuple.InternalRelationTuple
			err = json.NewDecoder(f).Decode(&r)
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not decode: %s\n", err)
				return cmdx.FailSilently(cmd)
			}

			cl := acl.NewWriteServiceClient(conn)

			_, err = cl.WriteRelationTuples(cmd.Context(), &acl.WriteRelationTuplesRequest{
				RelationTupleDeltas: []*acl.RelationTupleWriteDelta{
					{
						Action:        acl.RelationTupleWriteDelta_INSERT,
						RelationTuple: r.ToGRPC(),
					},
				},
			})
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Error doing the request: %s\n", err)
				return cmdx.FailSilently(cmd)
			}

			cmdx.PrintRow(cmd, &r)
			return nil
		},
	}
	cmd.Flags().AddFlagSet(packageFlags)

	return cmd
}
