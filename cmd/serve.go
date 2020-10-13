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
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/ory/graceful"
	"github.com/ory/keto/driver"
	"github.com/ory/keto/relation"
	"github.com/ory/keto/relation/read"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/ory/x/viperx"

	"github.com/ory/x/logrusx"
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
	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", ":4467")
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "%+v\n", err)
			os.Exit(1)
		}

		reg := &driver.RegistryDefault{}

		wg := &sync.WaitGroup{}
		wg.Add(2)

		go func() {
			defer wg.Done()

			s := grpc.NewServer()
			read.RegisterRelationReaderServer(s, relation.NewServer(reg))
			if err := s.Serve(lis); err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "%+v\n", err)
			}
		}()

		go func() {
			defer wg.Done()

			router := httprouter.New()
			h := relation.NewHandler(reg)
			h.RegisterPublicRoutes(router)

			server := graceful.WithDefaults(&http.Server{
				Addr:    ":4466",
				Handler: router,
			})

			if err := server.ListenAndServe(); err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "%+v\n", err)
			}
		}()

		wg.Wait()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	disableTelemetryEnv := viperx.GetBool(logrusx.New("ORY Keto", Version), "sqa.opt_out", false, "DISABLE_TELEMETRY")
	serveCmd.PersistentFlags().Bool("disable-telemetry", disableTelemetryEnv, "Disable anonymized telemetry reports - for more information please visit https://www.ory.sh/docs/ecosystem/sqa")
	serveCmd.PersistentFlags().Bool("sqa-opt-out", disableTelemetryEnv, "Disable anonymized telemetry reports - for more information please visit https://www.ory.sh/docs/ecosystem/sqa")
}
