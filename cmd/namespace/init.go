package namespace

import (
	"github.com/spf13/cobra"
)

func NewInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [<namespace-name> ...]",
		Short: "Initialize the namespace config",
		Long: `This command initializes the namespace config for the given namespaces.
A "default" namespace is created if none is specified.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				args = []string{"default"}
			}
			if err := generateConfigFiles(args, cmd.Flag(FlagOut).Value.String()); err != nil {
				return err
			}
			return nil
		},
	}
	registerOutputFlag(cmd)

	return cmd
}
