package namespace

import (
	"fmt"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver"
)

func NewMigrateDownCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "down <namespace-name> <steps>",
		Short: "Migrate a namespace down.",
		Args:  cobra.ExactArgs(2),
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

			if err := reg.NamespaceMigrator().MigrateNamespaceDown(ctx, n, 0); err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not apply namespace migration: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}

			return nil
		},
	}

	registerYesFlag(cmd.Flags())
	registerPackageFlags(cmd.Flags())

	return cmd
}
