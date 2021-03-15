package migrate

import (
	"fmt"
	"strconv"

	"github.com/ory/x/flagx"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver"
)

func newDownCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "down <steps>",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			steps, err := strconv.ParseInt(args[0], 0, 0)
			if err != nil {
				// return this error so it gets printed along the usage
				return fmt.Errorf("malformed argument %s for <steps>: %+v", args[0], err)
			}

			reg, err := driver.NewDefaultRegistry(ctx, cmd.Flags())
			if err != nil {
				return err
			}

			s, err := reg.Migrator().MigrationStatus(ctx)
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not get migration status: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}
			cmdx.PrintTable(cmd, s)

			if !flagx.MustGetBool(cmd, FlagYes) && !cmdx.AskForConfirmation("Do you really want to migrate down? This will delete data.", cmd.InOrStdin(), cmd.OutOrStdout()) {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Migration aborted.")
				return nil
			}

			if err := reg.Migrator().MigrateDown(ctx, int(steps)); err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could apply down migrations: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}

			s, err = reg.Migrator().MigrationStatus(ctx)
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not get migration status: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}
			cmdx.PrintTable(cmd, s)

			return nil
		},
	}

	registerYesFlag(cmd.Flags())
	cmdx.RegisterFormatFlags(cmd.Flags())

	return cmd
}
