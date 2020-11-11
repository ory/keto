package relationtuple

import (
	"context"
	"fmt"
	"strings"

	"github.com/ory/x/flagx"

	"github.com/ory/keto/relationtuple"

	"github.com/spf13/pflag"

	"github.com/ory/x/cmdx"

	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/client"
)

const (
	FlagSubject   = "subject"
	FlagRelation  = "relation"
	FlagObjectID  = "object"
	FlagNamespace = "namespace"
)

func registerRelationTupleFlags(flags *pflag.FlagSet) {
	flags.String(FlagSubject, "", "Set the requested subject")
	flags.String(FlagRelation, "", "Set the requested relation")
	flags.String(FlagObjectID, "", "Set the requested object ID")
	flags.String(FlagNamespace, "", "Set the requested namespace")
}

func readQueryFromFlags(cmd *cobra.Command) (*relationtuple.ReadRelationTuplesRequest_Query, error) {
	subject, err := cmd.Flags().GetString(FlagSubject)
	if err != nil {
		return nil, err
	}
	relation, err := cmd.Flags().GetString(FlagRelation)
	if err != nil {
		return nil, err
	}
	objectID, err := cmd.Flags().GetString(FlagObjectID)
	if err != nil {
		return nil, err
	}

	query := &relationtuple.ReadRelationTuplesRequest_Query{
		Relation:  relation,
		ObjectId:  objectID,
		Namespace: flagx.MustGetString(cmd, FlagNamespace),
	}

	subjectParts := strings.Split(subject, "#")
	if len(subjectParts) == 2 {
		query.Subject = &relationtuple.ReadRelationTuplesRequest_Query_UserSet{
			UserSet: &relationtuple.RelationUserSet{
				Object:   (&relationtuple.RelationObject{}).FromString(subjectParts[0]),
				Relation: subjectParts[1],
			},
		}
	} else {
		query.Subject = &relationtuple.ReadRelationTuplesRequest_Query_UserId{
			UserId: subject,
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
