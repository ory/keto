package namespace

import (
	"github.com/ory/keto/cmd/migrate"

	"github.com/spf13/cobra"
)

func NewMigrateUpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Deprecated: "This step is not necessary anymore, see TODO",
		Use:        "up <namespace-name>",
		Short:      "Migrate a namespace up",
		Long:       "Migrate a namespace up to the most recent migration.",
		Args:       cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, _ []string) error {
			return nil
		},
	}

	migrate.RegisterYesFlag(cmd.Flags())
	registerPackageFlags(cmd.Flags())

	return cmd
}
