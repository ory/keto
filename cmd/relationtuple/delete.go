package relationtuple

import (
	"fmt"
	"github.com/ory/keto/ketoapi"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/client"
)

func newDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <relation-tuple.json> [<relation-tuple-dir>]",
		Short: "Delete relation tuples defined in JSON files",
		Long: "Delete relation tuples defined in the given JSON files.\n" +
			"A directory will be traversed and all relation tuples will be deleted.\n" +
			"Pass the special filename `-` to read from STD_IN.",
		Args: cobra.MinimumNArgs(1),
		RunE: transactRelationTuples(rts.RelationTupleDelta_ACTION_DELETE),
	}
	registerPackageFlags(cmd.Flags())

	return cmd
}

func transactRelationTuples(action rts.RelationTupleDelta_Action) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		conn, err := client.GetWriteConn(cmd)
		if err != nil {
			return err
		}

		var tuples []*ketoapi.RelationTuple
		var deltas []*rts.RelationTupleDelta
		for _, fn := range args {
			tuple, err := readTuplesFromArg(cmd, fn)
			if err != nil {
				return err
			}
			for _, t := range tuple {
				tuples = append(tuples, t)
				deltas = append(deltas, &rts.RelationTupleDelta{
					Action:        action,
					RelationTuple: t.ToProto(),
				})
			}
		}

		cl := rts.NewWriteServiceClient(conn)

		_, err = cl.TransactRelationTuples(cmd.Context(), &rts.TransactRelationTuplesRequest{
			RelationTupleDeltas: deltas,
		})
		if err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Error doing the request: %s\n", err)
			return cmdx.FailSilently(cmd)
		}

		cmdx.PrintTable(cmd, NewAPICollection(tuples))
		return nil
	}
}
