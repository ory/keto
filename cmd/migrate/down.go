package migrate

import (
	"fmt"
	"strconv"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/flagx"
	"github.com/ory/x/popx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/ketoctx"
)

func newDownCmd(opts []ketoctx.Option) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "down <steps>",
		Short: "Migrate the database down",
		Long: "Migrate the database down a specific amount of steps.\n" +
			"Pass 0 steps to fully migrate down.",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			steps, err := strconv.ParseInt(args[0], 0, 0)
			if err != nil {
				// return this error so it gets printed along the usage
				return fmt.Errorf("malformed argument %s for <steps>: %+v", args[0], err)
			}

			reg, err := driver.NewDefaultRegistry(cmd.Context(), cmd.Flags(), true, opts...)
			if err != nil {
				return err
			}

			mb, err := reg.MigrationBox(cmd.Context())
			if err != nil {
				return err
			}

			return BoxDown(cmd, mb, int(steps))
		},
	}

	RegisterYesFlag(cmd.Flags())
	cmdx.RegisterFormatFlags(cmd.Flags())

	return cmd
}

func BoxDown(cmd *cobra.Command, mb *popx.MigrationBox, steps int) error {
	s, err := mb.Status(cmd.Context())
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not get migration status: %+v\n", err)
		return cmdx.FailSilently(cmd)
	}
	cmdx.PrintTable(cmd, s)

	if !flagx.MustGetBool(cmd, FlagYes) && !cmdx.AskForConfirmation("Do you really want to migrate down? This will delete data.", cmd.InOrStdin(), cmd.OutOrStdout()) {
		_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Migration aborted.")
		return nil
	}

	if err := mb.Down(cmd.Context(), steps); err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could apply down migrations: %+v\n", err)
		return cmdx.FailSilently(cmd)
	}

	s, err = mb.Status(cmd.Context())
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not get migration status: %+v\n", err)
		return cmdx.FailSilently(cmd)
	}
	cmdx.PrintTable(cmd, s)
	return nil
}
