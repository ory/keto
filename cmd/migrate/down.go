package migrate

import (
	"fmt"
	"strconv"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver"
)

func newDownCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "down <steps>",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			steps, err := strconv.ParseInt(args[0], 0, 0)
			if err != nil {
				// return this error so it gets printed along the usage
				return fmt.Errorf("malformed argument %s for <steps>: %+v", args[0], err)
			}

			reg, err := driver.NewDefaultRegistry(ctx, cmd.Flags())
			if err != nil {
				return err
			}
			if err := reg.Migrator().MigrateDown(ctx, int(steps)); err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could apply down migrations: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}

			return nil
		},
	}

	return cmd
}
