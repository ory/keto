package migrate

import (
	"context"
	"fmt"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver"
)

func newStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "status",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			reg := driver.NewDefaultRegistry(ctx, cmd.Flags())
			if err := reg.Migrator().MigrationStatus(ctx, cmd.OutOrStdout()); err != nil {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Could not get migration status: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}
			return nil
		},
	}
	return cmd
}
