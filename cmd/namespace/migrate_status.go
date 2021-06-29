package namespace

import (
	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"
)

func NewMigrateStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Deprecated: "This step is not necessary anymore, see TODO",
		Use:        "status <namespace-name>",
		Short:      "Get the current namespace migration status",
		Long:       "Get the current migration status of one specific namespace.\nDoes not apply any changes.",
		Args:       cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, _ []string) error {
			return nil
		},
	}

	cmdx.RegisterFormatFlags(cmd.Flags())

	return cmd
}
