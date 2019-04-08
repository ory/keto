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
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/server"
	"github.com/ory/x/cmdx"
	"github.com/ory/x/corsx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/profilex"
	"github.com/ory/x/sqlcon"
	"github.com/ory/x/tlsx"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the server and serves the HTTP REST API",
	Long: cmdx.EnvVarExamplesHelpMessage("keto") + `

All possible controls are listed below.

CORE CONTROLS
=============

` + sqlcon.HelpMessage() + `

` + logrusx.HelpMessage() + `

HTTP(S) CONTROLS
==============
- HOST: The host to listen on. Defaults to listening on all interfaces.

	Example:
		$ export HOST=127.0.0.1

- PORT: The port to listen on. Defaults to port 4466.

	Example:
		$ export PORT=4466

` + tlsx.HTTPSCertificateHelpMessage() + `

` + corsx.HelpMessage() + `


DEBUG CONTROLS
==============

` + profilex.HelpMessage() + `

`,
	Run: server.RunServe(logger, Version, Commit, Date),
}

func init() {
	RootCmd.AddCommand(serveCmd)

	disableTelemetryEnv, _ := strconv.ParseBool(os.Getenv("DISABLE_TELEMETRY")) // #nosec
	serveCmd.Flags().Bool("disable-telemetry", disableTelemetryEnv, "Disable anonymized telemetry reports - for more information please visit https://www.ory.sh/docs/ecosystem/sqa")
}
