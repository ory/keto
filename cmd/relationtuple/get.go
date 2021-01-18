package relationtuple

import (
	"fmt"

	"github.com/ory/x/flagx"

	acl "github.com/ory/keto/api/keto/acl/v1alpha1"
	"github.com/ory/keto/internal/relationtuple"

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

func readQueryFromFlags(cmd *cobra.Command, namespace string) (*acl.ListRelationTuplesRequest_Query, error) {
	subject := flagx.MustGetString(cmd, FlagSubject)
	relation := flagx.MustGetString(cmd, FlagRelation)
	object := flagx.MustGetString(cmd, FlagObject)

	query := &acl.ListRelationTuplesRequest_Query{
		Relation:  relation,
		Object:    object,
		Namespace: namespace,
	}

	if subject != "" {
		s, err := relationtuple.SubjectFromString(subject)
		if err != nil {
			return nil, err
		}

		query.Subject = s.ToProto()
	}

	return query, nil
}

func newGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get <namespace>",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := client.GetBasicConn(cmd)
			if err != nil {
				return err
			}
			defer conn.Close()

			cl := acl.NewReadServiceClient(conn)
			query, err := readQueryFromFlags(cmd, args[0])
			if err != nil {
				return err
			}

			resp, err := cl.ListRelationTuples(cmd.Context(), &acl.ListRelationTuplesRequest{
				Query:    query,
				PageSize: 100,
			})
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not make request: %s\n", err)
				return err
			}

			cmdx.PrintTable(cmd, relationtuple.NewProtoRelationCollection(resp.RelationTuples))
			return nil
		},
	}

	cmd.Flags().AddFlagSet(packageFlags)
	registerRelationTupleFlags(cmd.Flags())

	return cmd
}
