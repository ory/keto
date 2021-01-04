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
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/pkg/errors"

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
func newServe() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Starts the server and serves the HTTP REST API",
		Long: `This command opens a network port and listens to HTTP/2 API requests.

## Configuration

ORY Keto can be configured using environment variables as well as a configuration file. For more information
on configuration options, open the configuration documentation:

>> https://github.com/ory/keto/blob/` + config.Version + `/docs/config.yaml <<`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			reg, err := driver.NewDefaultRegistry(ctx, cmd.Flags())
			if err != nil {
				return err
			}

			wg := &sync.WaitGroup{}
			// the two servers + the ctx.Done listener go routine
			wg.Add(3)

			var grpcErr, restErr error
			var grpcServer *grpc.Server
			go func() {
				defer wg.Done()

				lis, err := net.Listen("tcp", reg.Config().GRPCListenOn())
				if err != nil {
					_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "%+v\n", err)
					grpcErr = err
					return
				}

				grpcServer = grpc.NewServer()
				relS := relationtuple.NewGRPCServer(reg)
				acl.RegisterReadServiceServer(grpcServer, relS)
				acl.RegisterWriteServiceServer(grpcServer, relS)

				checkS := check.NewGRPCServer(reg)
				acl.RegisterCheckServiceServer(grpcServer, checkS)

				expandS := expand.NewGRPCServer(reg)
				acl.RegisterExpandServiceServer(grpcServer, expandS)

				reg.Logger().WithField("addr", lis.Addr().String()).Info("serving GRPC")
				if err := grpcServer.Serve(lis); err != nil {
					_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "%+v\n", err)
					grpcErr = err
				}
			}()

			var restServer *http.Server
			go func() {
				defer wg.Done()

				router := httprouter.New()
				relationtuple.NewHandler(reg).RegisterPublicRoutes(router)
				check.NewHandler(reg).RegisterPublicRoutes(router)
				expand.NewHandler(reg).RegisterPublicRoutes(router)
				reg.HealthHandler().SetRoutes(router, false)

				restServer = graceful.WithDefaults(&http.Server{
					Addr:    reg.Config().RESTListenOn(),
					Handler: router,
				})

				reg.Logger().WithField("addr", restServer.Addr).Info("serving REST")
				if err := restServer.ListenAndServe(); err != nil {
					if errors.Is(err, http.ErrServerClosed) {
						// this means the server got closed and should not be reported
						return
					}

					_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "%+v\n", err)
					restErr = err
				}
			}()

			go func() {
				defer wg.Done()

				<-ctx.Done()
				// ctx is done already, so we need a new context
				err := restServer.Shutdown(context.Background())

				if restErr != nil {
					if err != nil {
						restErr = errors.Wrap(restErr, err.Error())
					}
				} else {
					restErr = err
				}

				grpcServer.GracefulStop()
			}()

			wg.Wait()

			if grpcErr == nil && restErr == nil {
				return nil
			}

			return fmt.Errorf("GRPC Error: %+v\nREST Error: %+v", grpcErr, restErr)
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

func RegisterCommandsRecursive(parent *cobra.Command) {
	parent.AddCommand(newServe())
}
