package migrate

import (
	"fmt"

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
		Use: "up",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()

			reg, err := driver.NewDefaultRegistry(ctx, cmd.Flags())
			if err != nil {
				return err
			}

			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Current status:")

			s, err := reg.Migrator().MigrationStatus(ctx)
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

			if err := reg.Migrator().MigrateUp(ctx); err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not apply migrations: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}

			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Successfully applied all migrations:")

			s, err = reg.Migrator().MigrationStatus(ctx)
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not get migration status: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}
			cmdx.PrintTable(cmd, s)

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
				s, err := reg.NamespaceMigrator().NamespaceStatus(cmd.Context(), nspace)
				if err != nil {
					_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not get migration status for namespace %s: %+v\n", nspace.Name, err)
					return cmdx.FailSilently(cmd)
				}
				cmdx.PrintTable(cmd, s)

				if !s.HasPending() {
					_, _ = fmt.Fprintf(cmd.OutOrStdout(), "All migrations are already applied for namespace %s, there is nothing to do.\n", nspace.Name)
					continue
				}

				if !flagx.MustGetBool(cmd, FlagYes) && !cmdx.AskForConfirmation(fmt.Sprintf("Do you want to apply above planned migrations for namespace %s?", nspace.Name), cmd.InOrStdin(), cmd.OutOrStdout()) {
					_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Skipping namespace %s\n", nspace.Name)
					continue
				}

				if err := reg.NamespaceMigrator().MigrateNamespaceUp(cmd.Context(), nspace); err != nil {
					_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not apply namespace migrations for namespace %s: %+v\n", nspace.Name, err)
					return cmdx.FailSilently(cmd)
				}

				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Successfully migrated namespace %s\n.", nspace.Name)
			}

			return nil
		},
	}

	registerYesFlag(cmd.Flags())
	cmd.Flags().BoolVar(&allNamespaces, FlagAllNamespace, false, "migrate all pending namespaces as well")

	cmdx.RegisterFormatFlags(cmd.Flags())

	return cmd
}

func registerYesFlag(flags *pflag.FlagSet) {
	flags.BoolP(FlagYes, "y", false, "yes to all questions, no user input required")
}
