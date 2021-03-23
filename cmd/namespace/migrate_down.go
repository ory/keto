package namespace

import (
	"fmt"
	"strconv"

	"github.com/ory/x/flagx"

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

			if !flagx.MustGetBool(cmd, YesFlag) && !cmdx.AskForConfirmation(fmt.Sprintf("Do you really want to delete namespace %s? This will irrecoverably delete all relation tuples within the namespace.", n.Name), cmd.InOrStdin(), cmd.OutOrStdout()) {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Migration of namespace \"%s\" aborted.\n", n.Name)
				return nil
			}

			if err := reg.NamespaceMigrator().MigrateNamespaceDown(ctx, n, int(steps)); err != nil {
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
