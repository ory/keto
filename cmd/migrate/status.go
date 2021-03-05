package migrate

import (
	"fmt"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver"
)

func newStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "status",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()

			reg, err := driver.NewDefaultRegistry(ctx, cmd.Flags())
			if err != nil {
				return err
			}

			s, err := reg.Migrator().MigrationStatus(ctx)
			if err != nil {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Could not get migration status: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}

			cmdx.PrintTable(cmd, s)
			return nil
		},
	}

	cmdx.RegisterFormatFlags(cmd.Flags())

	return cmd
}
