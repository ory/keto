package namespace

import (
	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"
)

func NewFromLegacy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "from-legacy",
		Short: "Convert legacy namespace configs to OPL configs",
		Long:  `This command converts legacy namespace configs to OPL configs.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmdx.RegisterFormatFlags()
			return nil
		},
	}

	return cmd
}
