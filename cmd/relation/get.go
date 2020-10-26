package relation

import (
	"context"
	"fmt"

	"github.com/ory/x/cmdx"

	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/models"
)

func newGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get-by-user <id>",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := client.GetGRPCConn(cmd)
			if err != nil {
				return err
			}
			defer conn.Close()

			cl := models.NewRelationTupleServiceClient(conn)
			resp, err := cl.ReadRelationTuples(context.Background(), &models.ReadRelationTuplesRequest{
				Tuplesets: []*models.ReadRelationTuplesRequest_Query{
					{
						User: &models.ReadRelationTuplesRequest_Query_UserId{
							UserId: args[0],
						},
					},
				},
				Page:    0,
				PerPage: 100,
			})
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not make request: %s\n", err)
				return err
			}

			cmdx.PrintCollection(cmd, models.NewGRPCRelationCollection(resp.Tuples))
			return nil
		},
	}

	cmd.Flags().AddFlagSet(packageFlags)

	return cmd
}
