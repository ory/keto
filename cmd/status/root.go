package status

import (
	"context"
	"errors"
	"fmt"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/stringsx"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"

	cliclient "github.com/ory/keto/cmd/client"
)

const (
	FlagBlock    = "block"
	FlagEndpoint = "endpoint"
)

func newStatusCmd() *cobra.Command {
	var (
		block    bool
		endpoint string
	)

	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get the status of the upstream Keto instance",
		Long:  "Get a status report about the upstream Keto instance. Can also block until the service is healthy.",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			var connect func(*cobra.Command) (*grpc.ClientConn, error)

			switch endpoints := stringsx.SwitchExact(endpoint); {
			case endpoints.AddCase("read"):
				connect = cliclient.GetReadConn
			case endpoints.AddCase("write"):
				connect = cliclient.GetWriteConn
			default:
				return endpoints.ToUnknownCaseErr()
			}

			loudPrinter := cmdx.NewLoudOutPrinter(cmd)

			conn, err := connect(cmd)
			for block && err != nil {
				if !errors.Is(err, context.DeadlineExceeded) {
					return err
				}
				_, _ = loudPrinter.Println("Context deadline exceeded, going to retry.")
				conn, err = connect(cmd)
			}

			if errors.Is(err, context.DeadlineExceeded) {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), grpcHealthV1.HealthCheckResponse_NOT_SERVING.String())
				return nil
			} else if err != nil {
				return err
			}

			c := grpcHealthV1.NewHealthClient(conn)

			var status interface {
				GetStatus() grpcHealthV1.HealthCheckResponse_ServingStatus
			}
			if block {
				ctx, cancel := context.WithCancel(cmd.Context())
				defer cancel()

				wc, err := c.Watch(ctx, &grpcHealthV1.HealthCheckRequest{})
				if err != nil {
					_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not start watching the status: %+v\n", err)
					return cmdx.FailSilently(cmd)
				}

				for {
					select {
					case <-cmd.Context().Done():
						return nil
					default:
					}

					status, err = wc.Recv()
					if err != nil {
						_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not watch the status: %+v\n", err)
						return cmdx.FailSilently(cmd)
					}

					if status.GetStatus() == grpcHealthV1.HealthCheckResponse_SERVING {
						cancel()
						break
					}

					_, _ = loudPrinter.Printf("Got the status %s, waiting until %s.\n", status.GetStatus(), grpcHealthV1.HealthCheckResponse_SERVING)
				}
			} else {
				status, err = c.Check(cmd.Context(), &grpcHealthV1.HealthCheckRequest{})
				if err != nil {
					_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Unable to get a check response: %+v\n", err)
					return cmdx.FailSilently(cmd)
				}
			}

			_, _ = fmt.Fprintln(cmd.OutOrStdout(), status.GetStatus().String())
			return nil
		},
	}

	cliclient.RegisterRemoteURLFlags(cmd.Flags())
	cmdx.RegisterNoiseFlags(cmd.Flags())

	cmd.Flags().BoolVarP(&block, FlagBlock, "b", false, "block until the service is healthy")
	cmd.Flags().StringVar(&endpoint, FlagEndpoint, "read", "which endpoint to use; one of {read, write}")

	return cmd
}

func RegisterCommandRecursive(parent *cobra.Command) {
	parent.AddCommand(newStatusCmd())
}
