package namespace

import (
	"fmt"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/migrate"
	"github.com/ory/keto/internal/driver"
)

func NewMigrateStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status <namespace-name>",
		Short: "Get the current namespace migration status",
		Long:  "Get the current namespace migration status",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			reg, err := driver.NewDefaultRegistry(cmd.Context(), cmd.Flags())
			if err != nil {
				return err
			}

			nm, err := reg.Config().NamespaceManager()
			if err != nil {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Could not initialize the namespace manager: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}

			n, err := nm.GetNamespace(cmd.Context(), args[0])
			if err != nil {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Could not find the namespace with name \"%s\": %+v\n", args[0], err)
				return cmdx.FailSilently(cmd)
			}

			mb, err := reg.NamespaceMigrator().NamespaceMigrationBox(cmd.Context(), n)
			if err != nil {
				return err
			}

			return migrate.BoxStatus(cmd, mb, "[namespace="+n.Name+"] ")
		},
	}

	cmdx.RegisterFormatFlags(cmd.Flags())

	return cmd
}
