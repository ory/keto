package relationtuple

import (
	"fmt"
	"strconv"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/ory/x/flagx"

	"github.com/ory/keto/internal/relationtuple"

	"github.com/spf13/pflag"

	"github.com/ory/x/cmdx"

	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/client"
)

const (
	FlagNamespace  = "namespace"
	FlagSubject    = "subject"
	FlagSubjectID  = "subject-id"
	FlagSubjectSet = "subject-set"
	FlagRelation   = "relation"
	FlagObject     = "object"
	FlagPageSize   = "page-size"
	FlagPageToken  = "page-token"
)

func registerRelationTupleFlags(flags *pflag.FlagSet) {
	flags.String(FlagNamespace, "", "Set the requested namespace")
	flags.String(FlagSubjectID, "", "Set the requested subject ID")
	flags.String(FlagSubjectSet, "", `Set the requested subject set; format: "namespace:object#relation"`)
	flags.String(FlagRelation, "", "Set the requested relation")
	flags.String(FlagObject, "", "Set the requested object")

	flags.String(FlagSubject, "", "")
	if err := flags.MarkHidden(FlagSubject); err != nil {
		panic(err.Error())
	}
}

func readQueryFromFlags(cmd *cobra.Command) (*rts.ListRelationTuplesRequest_Query, error) {
	query := &rts.ListRelationTuplesRequest_Query{
		Namespace: flagx.MustGetString(cmd, FlagNamespace),
		Object:    flagx.MustGetString(cmd, FlagObject),
		Relation:  flagx.MustGetString(cmd, FlagRelation),
	}

	switch flags := cmd.Flags(); {
	case flags.Changed(FlagSubjectID) && flags.Changed(FlagSubjectSet):
		return nil, relationtuple.ErrDuplicateSubject
	case flags.Changed(FlagSubjectID):
		query.Subject = (&relationtuple.SubjectID{ID: flagx.MustGetString(cmd, FlagSubjectID)}).ToProto()
	case flags.Changed(FlagSubjectSet):
		s, err := (&relationtuple.SubjectSet{}).FromString(flagx.MustGetString(cmd, FlagSubjectSet))
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
		Use:   "get",
		Short: "Get relation tuples",
		Long: "Get relation tuples matching the given partial tuple.\n" +
			"Returns paginated results.",
		Args: cobra.ExactArgs(0),
		RunE: getTuples(&pageSize, &pageToken),
	}

	cmd.Flags().AddFlagSet(packageFlags)
	registerRelationTupleFlags(cmd.Flags())

	cmd.Flags().StringVar(&pageToken, FlagPageToken, "", "page token acquired from a previous response")
	cmd.Flags().Int32Var(&pageSize, FlagPageSize, 100, "maximum number of items to return")

	return cmd
}

func getTuples(pageSize *int32, pageToken *string) func(cmd *cobra.Command, _ []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Changed(FlagSubject) {
			return fmt.Errorf("usage of --%s is not supported anymore, use --%s or --%s respectively", FlagSubject, FlagSubjectID, FlagSubjectSet)
		}

		conn, err := client.GetReadConn(cmd)
		if err != nil {
			return err
		}
		defer conn.Close()

		cl := rts.NewReadServiceClient(conn)
		query, err := readQueryFromFlags(cmd)
		if err != nil {
			return err
		}

		resp, err := cl.ListRelationTuples(cmd.Context(), &rts.ListRelationTuplesRequest{
			Query:     query,
			PageSize:  *pageSize,
			PageToken: *pageToken,
		})
		if err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not make request: %s\n", err)
			return cmdx.FailSilently(cmd)
		}

		cmdx.PrintTable(cmd, &responseOutput{
			RelationTuples: relationtuple.NewProtoRelationCollection(resp.RelationTuples),
			IsLastPage:     resp.NextPageToken == "",
			NextPageToken:  resp.NextPageToken,
		})
		return nil
	}
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
