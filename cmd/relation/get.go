package relation

import (
	"context"
	"fmt"
	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/relation/read"
	"github.com/spf13/cobra"
)

var getByUserRelationCmd = &cobra.Command{
	Use:  "get-by-user <id>",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := client.GetGRPCConn(cmd)
		if err != nil {
			return err
		}
		defer conn.Close()

		cl := read.NewRelationReaderClient(conn)
		resp, err := cl.RelationsByUser(context.Background(), &read.RelationsRequest{
			Page:    0,
			PerPage: 100,
			Id:      args[0],
		})
		if err != nil {
			return err
		}

		fmt.Printf("Got %d relations for user %s\n", len(resp.Relations), args[0])
		for _, r := range resp.Relations {
			fmt.Printf("%+v\n", r)
		}
		return nil
	},
}
