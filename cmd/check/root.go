// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check

import (
	"fmt"
	"strings"

	"github.com/ory/keto/ketoapi"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/ory/keto/internal/check"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/client"
)

type checkOutput check.CheckPermissionResult

func (o *checkOutput) String() string {
	if o.Allowed {
		return "Allowed\n"
	}
	return "Denied\n"
}

const FlagMaxDepth = "max-depth"

func NewCheckCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check <subject> <relation> <namespace> <object>",
		Short: "Check whether a subject has a relation on an object",
		Long:  "Check whether a subject has a relation on an object. This method resolves subject sets and subject set rewrites.",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := client.GetReadConn(cmd)
			if err != nil {
				return err
			}
			defer conn.Close()

			maxDepth, err := cmd.Flags().GetInt32(FlagMaxDepth)
			if err != nil {
				return err
			}

			cl := rts.NewCheckServiceClient(conn)

			sub, err := parseSubject(args[0])
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not parse subject %q: %s\n", args[0], err)
				return err
			}
			resp, err := cl.Check(cmd.Context(), &rts.CheckRequest{
				Tuple: &rts.RelationTuple{
					Namespace: args[2],
					Object:    args[3],
					Relation:  args[1],
					Subject:   sub,
				},
				MaxDepth: maxDepth,
			})
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not make request: %s\n", err)
				return err
			}

			cmdx.PrintJSONAble(cmd, &checkOutput{Allowed: resp.Allowed})
			return nil
		},
	}

	client.RegisterRemoteURLFlags(cmd.Flags())
	cmdx.RegisterFormatFlags(cmd.Flags())
	cmd.Flags().Int32P(FlagMaxDepth, "d", 0, "Maximum depth of the search tree. If the value is less than 1 or greater than the global max-depth then the global max-depth will be used instead.")

	return cmd
}

func RegisterCommandsRecursive(parent *cobra.Command) {
	parent.AddCommand(NewCheckCmd())
}

func parseSubject(s string) (*rts.Subject, error) {
	if strings.Contains(s, ":") {
		su, err := (&ketoapi.SubjectSet{}).FromString(s)
		if err != nil {
			return nil, err
		}

		return rts.NewSubjectSet(su.Namespace, su.Object, su.Relation), nil
	}
	return rts.NewSubjectID(s), nil
}
