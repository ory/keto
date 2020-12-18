package migrate

import (
	"fmt"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/flagx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver"
)

const FlagYes = "yes"

func newUpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "up",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()

			reg, err := driver.NewDefaultRegistry(ctx, cmd.Flags())
			if err != nil {
				return err
			}
			if err := reg.Migrator().MigrationStatus(ctx, cmd.OutOrStdout()); err != nil {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Could not get migration status: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}

			if !flagx.MustGetBool(cmd, FlagYes) && !cmdx.AskForConfirmation("Do you want to apply above planned migrations?", cmd.InOrStdin(), cmd.OutOrStdout()) {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Aborting")
				return nil
			}

			if err := reg.Migrator().MigrateUp(ctx); err != nil {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Could not apply migrations: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}
			return nil
		},
	}

	cmd.Flags().BoolP(FlagYes, "y", false, "yes to all questions, no user input required")

	return cmd
}
