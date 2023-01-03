// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"fmt"
	"strconv"

	"github.com/ory/keto/ketoapi"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/ory/x/flagx"
	"github.com/ory/x/pointerx"

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

func readQueryFromFlags(cmd *cobra.Command) (*rts.RelationQuery, error) {
	getStringPtr := func(flagName string) *string {
		if f := cmd.Flags().Lookup(flagName); f.Changed {
			return pointerx.Ptr(f.Value.String())
		}
		return nil
	}

	query := &ketoapi.RelationQuery{
		Namespace: getStringPtr(FlagNamespace),
		Object:    getStringPtr(FlagObject),
		Relation:  getStringPtr(FlagRelation),
		SubjectID: getStringPtr(FlagSubjectID),
	}
	if f := cmd.Flags().Lookup(FlagSubjectSet); f.Changed {
		s, err := (&ketoapi.SubjectSet{}).FromString(flagx.MustGetString(cmd, FlagSubjectSet))
		if err != nil {
			return nil, err
		}
		query.SubjectSet = s
	}

	return query.ToProto(), nil
}

func NewGetCmd() *cobra.Command {
	var (
		pageSize  int32
		pageToken string
	)

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get relationships",
		Long: "Get relationships matching the given partial tuple.\n" +
			"Returns paginated results.",
		Args: cobra.ExactArgs(0),
		RunE: getTuples(&pageSize, &pageToken),
	}

	registerPackageFlags(cmd.Flags())
	registerRelationTupleFlags(cmd.Flags())

	cmd.Flags().StringVar(&pageToken, FlagPageToken, "", "page token acquired from a previous response")
	cmd.Flags().Int32Var(&pageSize, FlagPageSize, 100, "maximum number of items to return")

	return cmd
}

func getTuples(pageSize *int32, pageToken *string) func(cmd *cobra.Command, _ []string) error {
	return func(cmd *cobra.Command, _ []string) error {
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
			RelationQuery: query,
			PageSize:      *pageSize,
			PageToken:     *pageToken,
		})
		if err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not make request: %s\n", err)
			return cmdx.FailSilently(cmd)
		}

		relationTuples, err := NewProtoCollection(resp.RelationTuples)
		if err != nil {
			return err
		}
		cmdx.PrintTable(cmd, &responseOutput{
			RelationTuples: relationTuples,
			IsLastPage:     resp.NextPageToken == "",
			NextPageToken:  resp.NextPageToken,
		})
		return nil
	}
}

type responseOutput struct {
	RelationTuples *Collection `json:"relation_tuples"`
	IsLastPage     bool        `json:"is_last_page"`
	NextPageToken  string      `json:"next_page_token"`
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
