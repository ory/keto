// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/helpers"
	"github.com/ory/keto/ketoctx"
)

// serveCmd represents the serve command
func newServe(opts []ketoctx.Option) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Starts the server and serves the HTTP REST and gRPC APIs",
		Long: `This command opens the network ports and listens to HTTP and gRPC API requests.

## Configuration

ORY Keto can be configured using environment variables as well as a configuration file. For more information
on configuration options, open the configuration documentation:

>> https://www.ory.sh/keto/docs/reference/configuration <<`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			reg, err := helpers.NewRegistry(cmd, opts)
			if err != nil {
				return err
			}

			return reg.ServeAllSQA(cmd)
		},
	}

	cmd.Flags().Bool("sqa-opt-out", false, "Disable anonymized telemetry reports - for more information please visit https://www.ory.sh/docs/ecosystem/sqa")

	return cmd
}

func RegisterCommandsRecursive(parent *cobra.Command, opts []ketoctx.Option) {
	parent.AddCommand(newServe(opts))
}
