// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package migrate

import (
	"fmt"
	"time"

	"github.com/ory/x/popx"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/ketoctx"
)

func newStatusCmd(opts []ketoctx.Option) *cobra.Command {
	block := false
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get the current migration status",
		Long: "Get the current migration status.\n" +
			"This does not affect namespaces. Use `keto namespace migrate status` for migrating namespaces.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()

			reg, err := driver.NewDefaultRegistry(ctx, cmd.Flags(), true, opts)
			if err != nil {
				return err
			}

			mb, err := reg.MigrationBox(ctx)
			if err != nil {
				return err
			}

			s, err := mb.Status(ctx)
			if err != nil {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Could not get migration status: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}

			for block && s.HasPending() {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Waiting for migrations to finish...\n")
				for _, m := range s {
					if m.State == popx.Pending {
						_, _ = fmt.Fprintf(cmd.OutOrStdout(), " - %s\n", m.Name)
					}
				}
				select {
				case <-ctx.Done():
					_, _ = fmt.Fprintln(cmd.ErrOrStderr(), "Context was canceled, exiting...")
					return cmdx.FailSilently(cmd)
				case <-time.After(time.Second):
				}
				s, err = mb.Status(ctx)
				if err != nil {
					_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Could not get migration status: %+v\n", err)
					return cmdx.FailSilently(cmd)
				}
			}

			cmdx.PrintTable(cmd, s)
			return nil
		},
	}

	cmdx.RegisterFormatFlags(cmd.Flags())
	cmd.Flags().BoolVar(&block, "block", false, "Block until all migrations have been applied")

	return cmd
}
