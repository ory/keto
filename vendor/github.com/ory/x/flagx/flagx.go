package flagx

import (
	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"
)

func MustGetBool(cmd *cobra.Command, name string) bool {
	ok, err := cmd.Flags().GetBool(name)
	if err != nil {
		cmdx.Fatalf(err.Error())
	}
	return ok
}

func MustGetString(cmd *cobra.Command, name string) string {
	ok, err := cmd.Flags().GetString(name)
	if err != nil {
		cmdx.Fatalf(err.Error())
	}
	return ok
}
