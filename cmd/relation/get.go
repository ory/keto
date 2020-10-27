package relation

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/pflag"

	"github.com/ory/x/cmdx"

	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/models"
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

func readQueryFromFlags(cmd *cobra.Command) (*models.ReadRelationTuplesRequest_Query, error) {
	subject, err := cmd.Flags().GetString(FlagSubject)
	if err != nil {
		return nil, err
	}
	relation, err := cmd.Flags().GetString(FlagRelation)
	if err != nil {
		return nil, err
	}
	object, err := cmd.Flags().GetString(FlagObject)
	if err != nil {
		return nil, err
	}

	query := &models.ReadRelationTuplesRequest_Query{
		Relation: relation,
		Object:   (&models.RelationObject{}).FromString(object),
	}

	subjectParts := strings.Split(subject, "#")
	if len(subjectParts) == 2 {
		query.Subject = &models.ReadRelationTuplesRequest_Query_UserSet{
			UserSet: &models.RelationUserSet{
				Object:   (&models.RelationObject{}).FromString(subjectParts[0]),
				Relation: subjectParts[1],
			},
		}
	} else {
		query.Subject = &models.ReadRelationTuplesRequest_Query_UserId{
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

			cl := models.NewRelationTupleServiceClient(conn)
			query, err := readQueryFromFlags(cmd)
			if err != nil {
				return err
			}
			resp, err := cl.ReadRelationTuples(context.Background(), &models.ReadRelationTuplesRequest{
				TupleSets: []*models.ReadRelationTuplesRequest_Query{query},
				Page:      0,
				PerPage:   100,
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
	registerRelationTupleFlags(cmd.Flags())

	return cmd
}
