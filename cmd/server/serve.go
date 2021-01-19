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
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/graceful"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
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
			reg, err := driver.NewDefaultRegistry(cmd.Context(), cmd.Flags())
			if err != nil {
				return err
			}

			eg := &errgroup.Group{}

			basicRouter, privilegedRouter := reg.BasicRouter().Router, reg.PrivilegedRouter().Router
			basicGRPC, privilegedGRPC := reg.BasicGRPCServer(), reg.PrivilegedGRPCServer()

			// basic port
			eg.Go(func() error {
				return multiplexPort(cmd.Context(), reg.Config().BasicListenOn(), basicRouter, basicGRPC)
			})

			// privileged port
			eg.Go(func() error {
				return multiplexPort(cmd.Context(), reg.Config().PrivilegedListenOn(), privilegedRouter, privilegedGRPC)
			})

			return eg.Wait()
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

func multiplexPort(ctx context.Context, addr string, router *httprouter.Router, grpcS *grpc.Server) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	m := cmux.New(l)

	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1())

	restS := graceful.WithDefaults(&http.Server{
		Handler: router,
	})

	eg := &errgroup.Group{}
	ctx, cancel := context.WithCancel(ctx)

	eg.Go(func() error {
		return errors.WithStack(grpcS.Serve(grpcL))
	})

	eg.Go(func() error {
		if err := restS.Serve(httpL); !errors.Is(err, http.ErrServerClosed) {
			// unexpected error
			return errors.WithStack(err)
		}
		return nil
	})

	eg.Go(func() error {
		err := m.Serve()
		if err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
			// unexpected error
			return errors.WithStack(err)
		}
		// trigger further shutdown
		cancel()
		return nil
	})

	eg.Go(func() error {
		<-ctx.Done()
		m.Close()
		return nil
	})

	if err := eg.Wait(); !errors.Is(err, cmux.ErrServerClosed) {
		return err
	}
	return nil
}
