package namespace

import (
	"context"
	"errors"
	"fmt"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/flagx"
	"github.com/ory/x/logrusx"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/persistence"
)

func NewMigrateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate <namespaces.yml>",
		Short: "Migrate a namespace up.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			d := driver.NewDefaultDriver(logrusx.New("keto", "master"), "master", "local", "today")

			n, err := validateNamespaceFile(cmd, args[0])
			if err != nil {
				return err
			}

			status, err := d.Registry().NamespaceMigrator().NamespaceStatus(context.Background(), n.ID)
			if err != nil {
				if !errors.Is(err, persistence.ErrNamespaceUnknown) {
					_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not get status for namespace \"%s\": %+v\n", n.Name, err)
					return cmdx.FailSilently(cmd)
				}
			} else {
				if status.CurrentVersion == status.NextVersion {
					_, _ = fmt.Fprintln(cmd.OutOrStdout(), "The namespace is already migrated up to the most recent version, there is noting to do.")
					return nil
				}

				cmdx.PrintRow(cmd, status)

				if !flagx.MustGetBool(cmd, YesFlag) {
					if !cmdx.AskForConfirmation("Are you sure that you want to apply this migration? Make sure to check the CHANGELOG and UPGRADE for breaking changes beforehand.", cmd.InOrStdin(), cmd.OutOrStdout()) {
						_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Migration of namespace \"%s\" aborted.\n", n.Name)
						return nil
					}
				}
			}

			if err := d.Registry().NamespaceMigrator().MigrateNamespaceUp(context.Background(), n); err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not apply namespace migration: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}

			status, err = d.Registry().NamespaceMigrator().NamespaceStatus(context.Background(), n.ID)
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not get status for namespace \"%s\": %+v\n", n.Name, err)
				return cmdx.FailSilently(cmd)
			}

			cmdx.PrintRow(cmd, status)

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
