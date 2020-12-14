package migrate

import (
	"context"
	"fmt"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/logrusx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver"
)

func newUpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "up",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			reg := driver.NewDefaultRegistry(ctx, logrusx.New("keto", "test"), cmd.Flags(), "test", "adf", "today")
			if err := reg.Migrator().MigrateUp(ctx); err != nil {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Could not apply migrations: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}
			return nil
		},
	}
	return cmd
}
