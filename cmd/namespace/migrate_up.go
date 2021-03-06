package namespace

import (
	"fmt"

	"github.com/ory/keto/cmd/migrate"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver"
)

func NewMigrateUpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up <namespace-name>",
		Short: "Migrate a namespace up",
		Long:  "Migrate a namespace up to the most recent migration.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			reg, err := driver.NewDefaultRegistry(ctx, cmd.Flags())
			if err != nil {
				return err
			}

			nm, err := reg.Config().NamespaceManager()
			if err != nil {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Could not initialize the namespace manager: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}

			n, err := nm.GetNamespace(ctx, args[0])
			if err != nil {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Could not find the namespace with name \"%s\": %+v\n", args[0], err)
				return cmdx.FailSilently(cmd)
			}

			mb, err := reg.NamespaceMigrator().NamespaceMigrationBox(ctx, n)
			if err != nil {
				return err
			}

			return migrate.BoxUp(cmd, mb, "[namespace="+n.Name+"] ")
		},
	}

	migrate.RegisterYesFlag(cmd.Flags())
	registerPackageFlags(cmd.Flags())

	return cmd
}
