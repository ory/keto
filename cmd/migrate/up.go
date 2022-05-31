package migrate

import (
	"fmt"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/flagx"
	"github.com/ory/x/popx"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/ketoctx"
)

const (
	FlagYes = "yes"
)

func newUpCmd(opts []ketoctx.Option) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up",
		Short: "Migrate the database up",
		Long: `Run this command on a fresh SQL installation and when you upgrade Ory Keto from version v0.7.x and later.

It is recommended to run this command close to the SQL instance (e.g. same subnet) instead of over the public internet.
This decreases risk of failure and decreases time required.

### WARNING ###

Before running this command on an existing database, create a back up!
`,
		Args: cobra.NoArgs,
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

			if err := BoxUp(cmd, mb); err != nil {
				return err
			}

			return nil
		},
	}

	RegisterYesFlag(cmd.Flags())

	cmdx.RegisterFormatFlags(cmd.Flags())

	return cmd
}

func RegisterYesFlag(flags *pflag.FlagSet) {
	flags.BoolP(FlagYes, "y", false, "yes to all questions, no user input required")
}

func BoxUp(cmd *cobra.Command, mb *popx.MigrationBox) error {
	_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Current status:")

	s, err := mb.Status(cmd.Context())
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not get migration status: %+v\n", err)
		return cmdx.FailSilently(cmd)
	}
	cmdx.PrintTable(cmd, s)

	if !s.HasPending() {
		_, _ = fmt.Fprintln(cmd.OutOrStdout(), "All migrations are already applied, there is nothing to do.")
		return nil
	}

	if !flagx.MustGetBool(cmd, FlagYes) && !cmdx.AskForConfirmation("Are you sure that you want to apply this migration? Make sure to check the CHANGELOG.md for breaking changes beforehand.", cmd.InOrStdin(), cmd.OutOrStdout()) {
		_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Aborting")
		return nil
	}

	_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Applying migrations...")

	if err := mb.Up(cmd.Context()); err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not apply migrations: %+v\n", err)
		return cmdx.FailSilently(cmd)
	}

	_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Successfully applied all migrations:")

	s, err = mb.Status(cmd.Context())
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not get migration status: %+v\n", err)
		return cmdx.FailSilently(cmd)
	}

	cmdx.PrintTable(cmd, s)
	return nil
}
