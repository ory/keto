package relationtuple

import (
	"fmt"
	"github.com/ory/keto/internal/x"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/ory/x/pointerx"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/flagx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/client"
)

const (
	FlagForce = "force"
)

func newDeleteAllCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-all",
		Short: "Delete ALL relation tuples matching the relation query.",
		Long: "Delete all relation tuples matching the relation query.\n" +
			"It is recommended to first run the command without the `--force` flag to verify that the operation is safe.",
		Args: cobra.ExactArgs(0),
		RunE: deleteRelationTuplesFromQuery,
	}
	registerPackageFlags(cmd.Flags())
	registerRelationTupleFlags(cmd.Flags())
	cmd.Flags().Bool(FlagForce, false, "Force the deletion of relation tuples")

	return cmd
}

func deleteRelationTuplesFromQuery(cmd *cobra.Command, args []string) error {
	if !flagx.MustGetBool(cmd, FlagForce) {
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "WARNING: This operation is not reversible. Please use the `--%s` flag to proceed.\nThese tuples would be deleted:\n\n", FlagForce)
		return getTuples(x.Ptr(int32(100)), pointerx.String(""))(cmd, args)
	}

	if cmd.Flags().Changed(FlagSubject) {
		return fmt.Errorf("usage of --%s is not supported anymore, use --%s or --%s respectively", FlagSubject, FlagSubjectID, FlagSubjectSet)
	}

	query, err := readQueryFromFlags(cmd)
	if err != nil {
		return err
	}

	conn, err := client.GetWriteConn(cmd)
	if err != nil {
		return err
	}
	defer conn.Close()
	cl := rts.NewWriteServiceClient(conn)
	_, err = cl.DeleteRelationTuples(cmd.Context(), &rts.DeleteRelationTuplesRequest{
		Query: (*rts.DeleteRelationTuplesRequest_Query)(query),
	})
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not make request: %s\n", err)
		return cmdx.FailSilently(cmd)
	}

	return nil
}
