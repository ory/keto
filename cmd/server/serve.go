// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"os"
	"strconv"

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
		RunE: func(cmd *cobra.Command, args []string) error {
			reg, err := helpers.NewRegistry(cmd, opts)
			if err != nil {
				return err
			}

			return reg.ServeAllSQA(cmd)
		},
	}
	disableTelemetry, err := strconv.ParseBool(os.Getenv("DISABLE_TELEMETRY"))
	if err != nil {
		disableTelemetry = true
	}
	sqaOptOut, err := strconv.ParseBool(os.Getenv("SQA_OPT_OUT"))
	if err != nil {
		sqaOptOut = true
	}

	cmd.Flags().Bool("disable-telemetry", disableTelemetry, "Disable anonymized telemetry reports - for more information please visit https://www.ory.sh/docs/ecosystem/sqa")
	cmd.Flags().Bool("sqa-opt-out", sqaOptOut, "Disable anonymized telemetry reports - for more information please visit https://www.ory.sh/docs/ecosystem/sqa")

	return cmd
}

func RegisterCommandsRecursive(parent *cobra.Command, opts []ketoctx.Option) {
	parent.AddCommand(newServe(opts))
}
