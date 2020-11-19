package relationtuple

import (
	"context"
	"fmt"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/x/flagx"

	"github.com/spf13/pflag"

	"github.com/ory/x/cmdx"

	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/client"
)

const (
	FlagSubject   = "subject"
	FlagRelation  = "relation"
	FlagObjectID  = "object-id"
	FlagNamespace = "namespace"
)

func registerRelationTupleFlags(flags *pflag.FlagSet) {
	flags.String(FlagSubject, "", "Set the requested subject")
	flags.String(FlagRelation, "", "Set the requested relation")
	flags.String(FlagObjectID, "", "Set the requested object")
	flags.String(FlagNamespace, "", "Set the requested namespace")
}

func readQueryFromFlags(cmd *cobra.Command) (*relationtuple.ReadRelationTuplesRequest_Query, error) {
	subject := flagx.MustGetString(cmd, FlagSubject)
	relation := flagx.MustGetString(cmd, FlagRelation)
	objectID := flagx.MustGetString(cmd, FlagObjectID)
	namespace := flagx.MustGetString(cmd, FlagNamespace)

	query := &relationtuple.ReadRelationTuplesRequest_Query{
		Relation:  relation,
		ObjectId:  objectID,
		Namespace: namespace,
	}

	relSub := relationtuple.SubjectFromString(subject)
	switch s := relSub.(type) {
	case *relationtuple.UserID:
		query.Subject = &relationtuple.ReadRelationTuplesRequest_Query_UserId{UserId: s.ID}
	case *relationtuple.UserSet:
		query.Subject = &relationtuple.ReadRelationTuplesRequest_Query_UserSet{
			UserSet: &relationtuple.RelationUserSet{
				ObjectId:  s.ObjectID,
				Namespace: s.Namespace,
				Relation:  s.Relation,
			},
		}
	}

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

			cl := relationtuple.NewRelationTupleServiceClient(conn)
			query, err := readQueryFromFlags(cmd)
			if err != nil {
				return err
			}
			resp, err := cl.ReadRelationTuples(context.Background(), &relationtuple.ReadRelationTuplesRequest{
				Query:   query,
				Page:    0,
				PerPage: 100,
			})
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not make request: %s\n", err)
				return err
			}

			cmdx.PrintCollection(cmd, relationtuple.NewGRPCRelationCollection(resp.Tuples))
			return nil
		},
	}

	cmd.Flags().AddFlagSet(packageFlags)
	registerRelationTupleFlags(cmd.Flags())

	return cmd
}
