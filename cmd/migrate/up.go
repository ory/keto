package migrate

import (
	"bytes"
	"fmt"
	"strings"

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
	var yes, allNamespaces bool

	cmd := &cobra.Command{
		Use: "up",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()

			reg, err := driver.NewDefaultRegistry(ctx, cmd.Flags())
			if err != nil {
				return err
			}

			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Current status:")

			status := &bytes.Buffer{}
			if err := reg.Migrator().MigrationStatus(ctx, status); err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not get migration status: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}
			_, _ = cmd.OutOrStdout().Write(status.Bytes())

			if !strings.Contains(status.String(), "Pending") {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "All migrations are already applied, there is nothing to do.")
				return nil
			}

			if !yes && !cmdx.AskForConfirmation("Do you want to apply above planned migrations?", cmd.InOrStdin(), cmd.OutOrStdout()) {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Aborting")
				return nil
			}

			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Applying migrations...")

			if err := reg.Migrator().MigrateUp(ctx); err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not apply migrations: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}

			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Successfully applied all migrations:")

			if err := reg.Migrator().MigrationStatus(ctx, cmd.OutOrStdout()); err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not get migration status: %+v\n", err)
				return cmdx.FailSilently(cmd)
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
				status := &bytes.Buffer{}
				if err := reg.NamespaceMigrator().NamespaceStatus(cmd.Context(), status, nspace); err != nil {
					_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not get migration status for namespace %s: %+v\n", nspace.Name, err)
					return cmdx.FailSilently(cmd)
				}
				_, _ = cmd.OutOrStdout().Write(status.Bytes())

				if !strings.Contains(status.String(), "Pending") {
					_, _ = fmt.Fprintf(cmd.OutOrStdout(), "All migrations are already applied for namespace %s, there is nothing to do.\n", nspace.Name)
					continue
				}

				if !yes && !cmdx.AskForConfirmation(fmt.Sprintf("Do you want to apply above planned migrations for namespace %s?", nspace.Name), cmd.InOrStdin(), cmd.OutOrStdout()) {
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

	cmd.Flags().BoolVarP(&yes, FlagYes, "y", false, "yes to all questions, no user input required")
	cmd.Flags().BoolVar(&allNamespaces, FlagAllNamespace, false, "migrate all pending namespaces as well")

	return cmd
}
