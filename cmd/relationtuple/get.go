package relationtuple

import (
	"context"
	"fmt"
	acl "github.com/ory/keto/api/keto/acl/v1alpha1"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/x/flagx"

	"github.com/spf13/pflag"

	"github.com/ory/x/cmdx"

	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/client"
)

const (
	FlagSubject  = "subject"
	FlagRelation = "relation"
	FlagObject   = "object"
)

func registerRelationTupleFlags(flags *pflag.FlagSet) {
	flags.String(FlagSubject, "", "Set the requested subject")
	flags.String(FlagRelation, "", "Set the requested relation")
	flags.String(FlagObject, "", "Set the requested object")
}

func readQueryFromFlags(cmd *cobra.Command) (*acl.ListRelationTuplesRequest_Query, error) {
	subject := flagx.MustGetString(cmd, FlagSubject)
	relation := flagx.MustGetString(cmd, FlagRelation)
	object := flagx.MustGetString(cmd, FlagObject)

	query := &acl.ListRelationTuplesRequest_Query{
		Relation: relation,
		Object:   object,
	}

	s, err := relationtuple.SubjectFromString(subject)
	if err != nil {
		return nil, err
	}

	query.Subject = s.ToGRPC()

	return query, nil
}

func newGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := client.GetGRPCConn(cmd)
			if err != nil {
				return err
			}
			defer conn.Close()

			cl := acl.NewReadServiceClient(conn)
			query, err := readQueryFromFlags(cmd)
			if err != nil {
				return err
			}

			resp, err := cl.ListRelationTuples(context.Background(), &acl.ListRelationTuplesRequest{
				Query:    query,
				PageSize: 100,
			})
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not make request: %s\n", err)
				return err
			}

			cmdx.PrintCollection(cmd, relationtuple.NewGRPCRelationCollection(resp.RelationTuples))
			return nil
		},
	}

	cmd.Flags().AddFlagSet(packageFlags)
	registerRelationTupleFlags(cmd.Flags())

	return cmd
}
