// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ory/x/cmdx"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/cmd/helpers"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type checkOutput struct {
	Allowed bool `json:"allowed"`
}

func (o *checkOutput) String() string {
	if o.Allowed {
		return "Allowed\n"
	}
	return "Denied\n"
}

const FlagMaxDepth = "max-depth"

func NewCheckCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check <subject_namespace>:<subject_id> <relation> <object_namespace>:<object_id>",
		Short: "Check whether a subject has a relation on an object",
		Long: "Check whether a subject has a relation on an object.\n\n" +
			"Example:\n" +
			"	keto check User:Alice view Doc:readme\n\n" +
			"Legacy format 'keto check <subject_namespace>:<subject_id> <relation> <object_namespace> <object_id>' is deprecated.",
		Args: cobra.RangeArgs(3, 4),
		RunE: check,
	}

	client.RegisterRemoteURLFlags(cmd.Flags())
	cmdx.RegisterFormatFlags(cmd.Flags())
	cmd.Flags().Int32P(FlagMaxDepth, "d", 0, "Maximum depth of the search tree. If the value is less than 1 or greater than the global max-depth then the global max-depth will be used instead.")

	return cmd
}

func check(cmd *cobra.Command, args []string) error {
	namespace, object, err := helpers.ParseNamespaceObject(cmd, args[2:])
	if err != nil {
		return err
	}

	sub, err := helpers.ParseSubject(args[0])
	if err != nil {
		return fmt.Errorf("could not parse subject %q: %w", args[0], err)
	}
	conn, err := client.GetReadConn(cmd)
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	maxDepth, err := cmd.Flags().GetInt32(FlagMaxDepth)
	if err != nil {
		return err
	}

	cl := rts.NewCheckServiceClient(conn)
	resp, err := cl.Check(cmd.Context(), &rts.CheckRequest{
		Tuple: &rts.RelationTuple{
			Namespace: namespace,
			Object:    object,
			Relation:  args[1],
			Subject:   sub,
		},
		MaxDepth: maxDepth,
	})
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not make request: %s\n", err)
		return cmdx.FailSilently(cmd)
	}

	cmdx.PrintJSONAble(cmd, &checkOutput{Allowed: resp.Allowed})
	return nil
}

func RegisterCommandsRecursive(parent *cobra.Command) {
	parent.AddCommand(NewCheckCmd())
}
