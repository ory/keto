package migrate

import (
	"fmt"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/popx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/ketoctx"
)

func newStatusCmd(opts []ketoctx.Option) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get the current migration status",
		Long: "Get the current migration status.\n" +
			"This does not affect namespaces. Use `keto namespace migrate status` for migrating namespaces.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()

			reg, err := driver.NewDefaultRegistry(ctx, cmd.Flags(), true, opts...)
			if err != nil {
				return err
			}

			mb, err := reg.MigrationBox(ctx)
			if err != nil {
				return err
			}

			return BoxStatus(cmd, mb, "")
		},
	}

	cmdx.RegisterFormatFlags(cmd.Flags())

	return cmd
}

func BoxStatus(cmd *cobra.Command, mb *popx.MigrationBox, msgPrefix string) error {
	s, err := mb.Status(cmd.Context())
	if err != nil {
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%sCould not get migration status: %+v\n", msgPrefix, err)
		return cmdx.FailSilently(cmd)
	}

	cmdx.PrintTable(cmd, s)
	return nil
}
