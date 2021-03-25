package namespace

import (
	"fmt"
	"strconv"

	"github.com/ory/keto/cmd/migrate"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver"
)

func NewMigrateDownCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "down <namespace-name> <steps>",
		Short: "Migrate a namespace down",
		Long: "Migrate a namespace down.\n" +
			"Pass 0 steps to fully migrate down.",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			steps, err := strconv.ParseInt(args[1], 0, 0)
			if err != nil {
				// return this error so it gets printed along the usage
				return fmt.Errorf("malformed argument %s for <steps>: %+v", args[0], err)
			}

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

			return migrate.BoxDown(cmd, mb, int(steps), "[namespace="+n.Name+"] ")
		},
	}

	migrate.RegisterYesFlag(cmd.Flags())
	registerPackageFlags(cmd.Flags())

	return cmd
}
