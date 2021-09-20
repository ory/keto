package namespace

import (
	"fmt"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/flagx"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/migrate"
	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/internal/persistence/sql/migrations"
)

func NewMigrateLegacyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "legacy [<namespace-name>]",
		Short: "Migrate a namespace from v0.6.x to v0.7.x and later.",
		Long: "Migrate a legacy namespaces from v0.6.x to the v0.7.x and later.\n" +
			"This step only has to be executed once.\n" +
			"If no namespace is specified, all legacy namespaces will be migrated.\n" +
			"Please ensure that namespace IDs did not change in the config file and you have a backup in case something goes wrong!",
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			reg, err := driver.NewDefaultRegistry(cmd.Context(), cmd.Flags(), false)
			if errors.Is(err, persistence.ErrNetworkMigrationsMissing) {
				_, _ = fmt.Fprintln(cmd.ErrOrStderr(), "Migrations were not applied yet, please apply them first using `keto migrate up`.")
				return cmdx.FailSilently(cmd)
			} else if err != nil {
				return err
			}

			migrator := migrations.NewToSingleTableMigrator(reg)

			var nn []*namespace.Namespace
			if len(args) == 1 {
				nm, err := reg.Config().NamespaceManager()
				if err != nil {
					_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "There seems to be a problem with the config: %s\n", err.Error())
					return cmdx.FailSilently(cmd)
				}
				n, err := nm.GetNamespaceByName(cmd.Context(), args[0])
				if err != nil {
					_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "There seems to be a problem with the config: %s\n", err.Error())
					return cmdx.FailSilently(cmd)
				}

				nn = []*namespace.Namespace{n}

				if !flagx.MustGetBool(cmd, migrate.FlagYes) &&
					!cmdx.AskForConfirmation(
						fmt.Sprintf("Are you sure that you want to migrate the namespace '%s'?", args[0]),
						cmd.InOrStdin(), cmd.OutOrStdout()) {
					_, _ = fmt.Fprintln(cmd.OutOrStdout(), "OK, aborting.")
					return nil
				}
			} else {
				nn, err = migrator.LegacyNamespaces(cmd.Context())
				if err != nil {
					_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not get legacy namespaces: %s\n", err.Error())
					return cmdx.FailSilently(cmd)
				}

				if len(nn) == 0 {
					_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Could not find legacy namespaces, there seems nothing to be done.")
					return nil
				}

				var names string
				for _, n := range nn {
					names += "  " + n.Name + "\n"
				}
				if !flagx.MustGetBool(cmd, migrate.FlagYes) &&
					!cmdx.AskForConfirmation(
						fmt.Sprintf("I found the following legacy namespaces:\n%sDo you want to migrate all of them?", names),
						cmd.InOrStdin(), cmd.OutOrStdout()) {
					_, _ = fmt.Fprintln(cmd.OutOrStdout(), "OK, aborting.")
					return nil
				}
			}

			for _, n := range nn {
				if err := migrator.MigrateNamespace(cmd.Context(), n); err != nil {
					_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Encountered error while migrating: %s\nAborting.\n", err.Error())
					return cmdx.FailSilently(cmd)
				}
				if err := migrator.MigrateDown(cmd.Context(), n); err != nil {
					_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not migrate down: %s\n", err.Error())
					return cmdx.FailSilently(cmd)
				}
			}

			return nil
		},
	}

	migrate.RegisterYesFlag(cmd.Flags())
	registerPackageFlags(cmd.Flags())

	return cmd
}
