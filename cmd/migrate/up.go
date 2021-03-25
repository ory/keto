package migrate

import (
	"fmt"

	"github.com/ory/x/popx"

	"github.com/ory/x/flagx"

	"github.com/spf13/pflag"

	"github.com/pkg/errors"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver"
)

const (
	FlagYes          = "yes"
	FlagAllNamespace = "all-namespaces"
)

func newUpCmd() *cobra.Command {
	var allNamespaces bool

	cmd := &cobra.Command{
		Use:   "up",
		Short: "Migrate the database up",
		Long: "Migrate the database up.\n" +
			"This does not affect namespaces. Use `keto namespace migrate up` for migrating namespaces.",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()

			reg, err := driver.NewDefaultRegistry(ctx, cmd.Flags())
			if err != nil {
				return err
			}

			mb, err := reg.Migrator().MigrationBox(ctx)
			if err != nil {
				return err
			}

			if err := BoxUp(cmd, mb, ""); err != nil {
				return err
			}

			if !allNamespaces {
				// everything is done already
				return nil
			}

			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "\nGoing to migrate namespaces.")

			nm, err := reg.Config().NamespaceManager()
			if err != nil {
				return errors.Wrap(err, "could not get the namespace manager")
			}

			nspaces, err := nm.Namespaces(cmd.Context())
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not get namespaces: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}

			for _, nspace := range nspaces {
				mb, err := reg.NamespaceMigrator().NamespaceMigrationBox(ctx, nspace)
				if err != nil {
					return err
				}

				if err := BoxUp(cmd, mb, "[namespace="+nspace.Name+"] "); err != nil {
					return err
				}
			}

			return nil
		},
	}

	RegisterYesFlag(cmd.Flags())
	cmd.Flags().BoolVar(&allNamespaces, FlagAllNamespace, false, "migrate all pending namespaces as well")

	cmdx.RegisterFormatFlags(cmd.Flags())

	return cmd
}

func RegisterYesFlag(flags *pflag.FlagSet) {
	flags.BoolP(FlagYes, "y", false, "yes to all questions, no user input required")
}

func BoxUp(cmd *cobra.Command, mb *popx.MigrationBox, msgPrefix string) error {
	_, _ = fmt.Fprintln(cmd.OutOrStdout(), msgPrefix+"Current status:")

	s, err := mb.Status(cmd.Context())
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "%sCould not get migration status: %+v\n", msgPrefix, err)
		return cmdx.FailSilently(cmd)
	}
	cmdx.PrintTable(cmd, s)

	if !s.HasPending() {
		_, _ = fmt.Fprintln(cmd.OutOrStdout(), msgPrefix+"All migrations are already applied, there is nothing to do.")
		return nil
	}

	if !flagx.MustGetBool(cmd, FlagYes) && !cmdx.AskForConfirmation(msgPrefix+"Are you sure that you want to apply this migration? Make sure to check the CHANGELOG.md for breaking changes beforehand.", cmd.InOrStdin(), cmd.OutOrStdout()) {
		_, _ = fmt.Fprintln(cmd.OutOrStdout(), msgPrefix+"Aborting")
		return nil
	}

	_, _ = fmt.Fprintln(cmd.OutOrStdout(), msgPrefix+"Applying migrations...")

	if err := mb.Up(cmd.Context()); err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "%sCould not apply migrations: %+v\n", msgPrefix, err)
		return cmdx.FailSilently(cmd)
	}

	_, _ = fmt.Fprintln(cmd.OutOrStdout(), msgPrefix+"Successfully applied all migrations:")

	s, err = mb.Status(cmd.Context())
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "%sCould not get migration status: %+v\n", msgPrefix, err)
		return cmdx.FailSilently(cmd)
	}

	cmdx.PrintTable(cmd, s)
	return nil
}
