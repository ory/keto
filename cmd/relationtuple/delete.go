package relationtuple

import (
	"fmt"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/internal/relationtuple"
	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"
)

func newDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete <relation-tuple.json> [<relation-tuple-dir>]",
		Args: cobra.MinimumNArgs(1),
		RunE: transactRelationTuples(acl.RelationTupleDelta_DELETE),
	}
	cmd.Flags().AddFlagSet(packageFlags)

	return cmd
}

func transactRelationTuples(action acl.RelationTupleDelta_Action) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		conn, err := client.GetWriteConn(cmd)
		if err != nil {
			return err
		}

		var tuples []*relationtuple.InternalRelationTuple
		var deltas []*acl.RelationTupleDelta
		for _, fn := range args {
			tuple, err := readTuplesFromArg(cmd, fn)
			if err != nil {
				return err
			}
			for _, t := range tuple {
				tuples = append(tuples, t)
				deltas = append(deltas, &acl.RelationTupleDelta{
					Action:        action,
					RelationTuple: t.ToProto(),
				})
			}
		}

		cl := acl.NewWriteServiceClient(conn)

		_, err = cl.TransactRelationTuples(cmd.Context(), &acl.TransactRelationTuplesRequest{
			RelationTupleDeltas: deltas,
		})
		if err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Error doing the request: %s\n", err)
			return cmdx.FailSilently(cmd)
		}

		cmdx.PrintTable(cmd, relationtuple.NewRelationCollection(tuples))
		return nil
	}
}
