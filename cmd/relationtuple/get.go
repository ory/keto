package relationtuple

import (
	"fmt"
	"strconv"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"

	"github.com/ory/x/flagx"

	"github.com/ory/keto/internal/relationtuple"

	"github.com/spf13/pflag"

	"github.com/ory/x/cmdx"

	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/client"
)

const (
	FlagSubject   = "subject"
	FlagRelation  = "relation"
	FlagObject    = "object"
	FlagPageSize  = "page-size"
	FlagPageToken = "page-token"
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
	var (
		pageSize  int32
		pageToken string
	)

	cmd := &cobra.Command{
		Use:  "get <namespace>",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := client.GetReadConn(cmd)
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
				Query:     query,
				PageSize:  pageSize,
				PageToken: pageToken,
			})
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not make request: %s\n", err)
				return cmdx.FailSilently(cmd)
			}

			cmdx.PrintTable(cmd, &responseOutput{
				RelationTuples: relationtuple.NewProtoRelationCollection(resp.RelationTuples),
				IsLastPage:     resp.IsLastPage,
				NextPageToken:  resp.NextPageToken,
			})
			return nil
		},
	}

	cmd.Flags().AddFlagSet(packageFlags)
	registerRelationTupleFlags(cmd.Flags())

	cmd.Flags().StringVar(&pageToken, FlagPageToken, "", "page token acquired from a previous response")
	cmd.Flags().Int32Var(&pageSize, FlagPageSize, 100, "maximum number of items to return")

	return cmd
}

type responseOutput struct {
	RelationTuples *relationtuple.RelationCollection `json:"relation_tuples"`
	IsLastPage     bool                              `json:"is_last_page"`
	NextPageToken  string                            `json:"next_page_token"`
}

func (r *responseOutput) Header() []string {
	return r.RelationTuples.Header()
}

func (r *responseOutput) Table() [][]string {
	return append(
		r.RelationTuples.Table(),
		[]string{},
		[]string{"NEXT PAGE TOKEN", r.NextPageToken},
		[]string{"IS LAST PAGE", strconv.FormatBool(r.IsLastPage)},
	)
}

func (r *responseOutput) Interface() interface{} {
	return r
}

func (r *responseOutput) Len() int {
	return r.RelationTuples.Len() + 3
}

func (r *responseOutput) IDs() []string {
	return r.RelationTuples.IDs()
}

var _ cmdx.Table = (*responseOutput)(nil)
