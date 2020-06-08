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

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ory/x/viperx"

	"github.com/ory/x/logrusx"

	"github.com/ory/keto/cmd/server"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the server and serves the HTTP REST API",
	Long: `This command opens a network port and listens to HTTP/2 API requests.

## Configuration

ORY Keto can be configured using environment variables as well as a configuration file. For more information
on configuration options, open the configuration documentation:

>> https://github.com/ory/keto/blob/` + Version + `/docs/config.yaml <<`,
	Run: server.RunServe(logger, Version, Commit, Date),
}

func init() {
	RootCmd.AddCommand(serveCmd)

	disableTelemetryEnv := viperx.GetBool(logrusx.New("ORY Keto", Version), "sqa.opt_out", false, "DISABLE_TELEMETRY")
	serveCmd.PersistentFlags().Bool("disable-telemetry", disableTelemetryEnv, "Disable anonymized telemetry reports - for more information please visit https://www.ory.sh/docs/ecosystem/sqa")
	serveCmd.PersistentFlags().Bool("sqa-opt-out", disableTelemetryEnv, "Disable anonymized telemetry reports - for more information please visit https://www.ory.sh/docs/ecosystem/sqa")
}
