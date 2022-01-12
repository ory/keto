package relationtuple

import (
	"fmt"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/flagx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/internal/relationtuple"
	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"
)

const FlagNamespace = "namespace"

func newDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [<relation-tuple.json>] [<relation-tuple-dir>]",
		Short: "Delete relation tuples defined in JSON files",
		Long: "Delete relation tuples defined in the given JSON files.\n" +
			"A directory will be traversed and all relation tuples will be deleted.\n" +
			"Pass the special filename `-` to read from STD_IN.",
		Args: cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return deleteRelationTuplesFromQuery()(cmd, args)
			}
			return transactRelationTuples(acl.RelationTupleDelta_DELETE)(cmd, args)
		},
	}
	cmd.Flags().AddFlagSet(packageFlags)
	registerRelationTupleFlags(cmd.Flags())
	cmd.Flags().String(FlagNamespace, "", "Set the requested namespace")

	return cmd
}

func deleteRelationTuplesFromQuery() func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, s []string) error {
		if cmd.Flags().Changed(FlagSubject) {
			return fmt.Errorf("usage of --%s is not supported anymore, use --%s or --%s respectively", FlagSubject, FlagSubjectID, FlagSubjectSet)
		}

		var namespace string
		if cmd.Flags().Changed(FlagNamespace) {
			namespace = flagx.MustGetString(cmd, FlagNamespace)

		}
		query, err := readQueryFromFlags(cmd, namespace)
		if err != nil {
			return err
		}

		conn, err := client.GetWriteConn(cmd)
		if err != nil {
			return err
		}
		defer conn.Close()
		cl := acl.NewWriteServiceClient(conn)
		_, err = cl.DeleteRelationTuples(cmd.Context(), &acl.DeleteRelationTuplesRequest{
			Query: (*acl.DeleteRelationTuplesRequest_Query)(query),
		})
		if err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not make request: %s\n", err)
			return cmdx.FailSilently(cmd)
		}

		return nil
	}
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
