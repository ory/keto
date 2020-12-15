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
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"

	acl "github.com/ory/keto/api/keto/acl/v1alpha1"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/graceful"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/relationtuple"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the server and serves the HTTP REST API",
	Long: `This command opens a network port and listens to HTTP/2 API requests.

## Configuration

ORY Keto can be configured using environment variables as well as a configuration file. For more information
on configuration options, open the configuration documentation:

>> https://github.com/ory/keto/blob/` + config.Version + `/docs/config.yaml <<`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		reg := driver.NewDefaultRegistry(ctx, cmd.Flags())

		wg := &sync.WaitGroup{}
		wg.Add(2)

		go func() {
			defer wg.Done()

			lis, err := net.Listen("tcp", reg.Config().GRPCListenOn())
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "%+v\n", err)
				return
			}

			s := grpc.NewServer()
			relS := relationtuple.NewGRPCServer(reg)
			acl.RegisterReadServiceServer(s, relS)

			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Serving GRPC on %s\n", lis.Addr().String())
			if err := s.Serve(lis); err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "%+v\n", err)
			}
		}()

		go func() {
			defer wg.Done()

			router := httprouter.New()
			relationtuple.NewHandler(reg).RegisterPublicRoutes(router)
			check.NewHandler(reg).RegisterPublicRoutes(router)
			expand.NewHandler(reg).RegisterPublicRoutes(router)

			server := graceful.WithDefaults(&http.Server{
				Addr:    reg.Config().RESTListenOn(),
				Handler: router,
			})

			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Serving REST on %s\n", server.Addr)
			if err := server.ListenAndServe(); err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "%+v\n", err)
			}
		}()

		wg.Wait()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	disableTelemetry, err := strconv.ParseBool(os.Getenv("DISABLE_TELEMETRY"))
	if err != nil {
		disableTelemetry = true
	}
	sqaOptOut, err := strconv.ParseBool(os.Getenv("SQA_OPT_OUT"))
	if err != nil {
		sqaOptOut = true
	}

	serveCmd.PersistentFlags().Bool("disable-telemetry", disableTelemetry, "Disable anonymized telemetry reports - for more information please visit https://www.ory.sh/docs/ecosystem/sqa")
	serveCmd.PersistentFlags().Bool("sqa-opt-out", sqaOptOut, "Disable anonymized telemetry reports - for more information please visit https://www.ory.sh/docs/ecosystem/sqa")
}
