package namespace

import (
	"github.com/ory/keto/cmd/migrate"

	"github.com/spf13/cobra"
)

func NewMigrateDownCmd() *cobra.Command {
	cmd := &cobra.Command{
		Deprecated: "This step is not necessary anymore, see TODO",
		Use:        "down <namespace-name> <steps>",
		Short:      "Migrate a namespace down",
		Long: "Migrate a namespace down.\n" +
			"Pass 0 steps to fully migrate down.",
		Args: cobra.ExactArgs(2),
		RunE: func(_ *cobra.Command, _ []string) error {
			return nil
		},
	}

	migrate.RegisterYesFlag(cmd.Flags())
	registerPackageFlags(cmd.Flags())

	return cmd
}
