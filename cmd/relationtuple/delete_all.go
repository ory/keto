// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"fmt"

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

func NewDeleteAllCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-all",
		Short: "Delete ALL relationships matching the relation query.",
		Long: "Delete all relationships matching the relation query.\n" +
			"It is recommended to first run the command without the `--force` flag to verify that the operation is safe.",
		Args: cobra.ExactArgs(0),
		RunE: deleteRelationTuplesFromQuery,
	}
	registerPackageFlags(cmd.Flags())
	registerRelationTupleFlags(cmd.Flags())
	cmd.Flags().Bool(FlagForce, false, "Force the deletion of relationships")

	return cmd
}

func deleteRelationTuplesFromQuery(cmd *cobra.Command, args []string) error {
	if !flagx.MustGetBool(cmd, FlagForce) {
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "WARNING: This operation is not reversible. Please use the `--%s` flag to proceed.\nThese tuples would be deleted:\n\n", FlagForce)
		return getTuples(pointerx.Ptr(int32(100)), pointerx.String(""))(cmd, args)
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
		RelationQuery: query,
	})
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not make request: %s\n", err)
		return cmdx.FailSilently(cmd)
	}

	return nil
}
