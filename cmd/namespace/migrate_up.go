package namespace

import (
	"errors"
	"fmt"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/flagx"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/persistence"
)

func NewMigrateUpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up <namespace-name>",
		Short: "Migrate a namespace up.",
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

			s, err := reg.NamespaceMigrator().NamespaceStatus(ctx, n)
			if err != nil {
				if !errors.Is(err, persistence.ErrNamespaceUnknown) {
					_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not get status for namespace \"%s\": %+v\n", n.Name, err)
					return cmdx.FailSilently(cmd)
				}
			} else {
				cmdx.PrintTable(cmd, s)

				if !s.HasPending() {
					_, _ = fmt.Fprintln(cmd.OutOrStdout(), "The namespace is already migrated up to the most recent version, there is noting to do.")
					return nil
				}

				if !flagx.MustGetBool(cmd, YesFlag) {
					if !cmdx.AskForConfirmation("Are you sure that you want to apply this migration? Make sure to check the CHANGELOG.md for breaking changes beforehand.", cmd.InOrStdin(), cmd.OutOrStdout()) {
						_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Migration of namespace \"%s\" aborted.\n", n.Name)
						return nil
					}
				}
			}

			if err := reg.NamespaceMigrator().MigrateNamespaceUp(ctx, n); err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not apply namespace migration: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}

			s, err = reg.NamespaceMigrator().NamespaceStatus(ctx, n)
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not get status for namespace \"%s\": %+v\n", n.Name, err)
				return cmdx.FailSilently(cmd)
			}
			cmdx.PrintTable(cmd, s)

			return nil
		},
	}

	registerYesFlag(cmd.Flags())
	registerPackageFlags(cmd.Flags())

	return cmd
}

const YesFlag = "yes"

func registerYesFlag(flags *pflag.FlagSet) {
	flags.BoolP(YesFlag, YesFlag[:1], false, "answer all questions with yes")
}
