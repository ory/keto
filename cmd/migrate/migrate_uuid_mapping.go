package migrate

import (
	"fmt"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/flagx"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/internal/persistence/sql/migrations"
)

func newMigrateUUIDMappingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uuid-mapping",
		Short: "Migrate the non-UUID subject and object names to UUIDs.",
		Long: `Migrate the non-UUID subject and object names to UUIDs.
This step only has to be executed once.
Please ensure that you have a backup in case something goes wrong!"`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			reg, err := driver.NewDefaultRegistry(cmd.Context(), cmd.Flags(), false)
			if errors.Is(err, persistence.ErrNetworkMigrationsMissing) {
				_, _ = fmt.Fprintln(cmd.ErrOrStderr(),
					"Migrations were not applied yet, please apply them first using `keto migrate up`.")
				return cmdx.FailSilently(cmd)
			} else if err != nil {
				return err
			}

			if !flagx.MustGetBool(cmd, FlagYes) &&
				!cmdx.AskForConfirmation(
					"Are you sure you want to migrate the subject and object names to UUIDs?",
					cmd.InOrStdin(), cmd.OutOrStdout()) {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "OK, aborting.")
				return nil
			}

			migrator := migrations.NewToUUIDMappingMigrator(reg)
			return migrator.MigrateUUIDMappings(cmd.Context())
		},
	}
	RegisterYesFlag(cmd.Flags())

	return cmd
}
