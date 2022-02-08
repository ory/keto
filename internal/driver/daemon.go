package driver

import (
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/relationtuple"

	"github.com/ory/analytics-go/v4"
	"github.com/ory/x/healthx"
	"github.com/ory/x/metricsx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver/config"

	"github.com/ory/graceful"
	"github.com/pkg/errors"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func (r *RegistryDefault) enableSqa(cmd *cobra.Command) {
	r.sqaService = metricsx.New(
		cmd,
		r.Logger(),
		r.Config().Source(),
		&metricsx.Options{
			Service:       "ory-keto",
			ClusterID:     metricsx.Hash(r.Config().DSN()),
			IsDevelopment: strings.HasPrefix(r.Config().DSN(), "sqlite"),
			WriteKey:      "qQlI6q8Q4WvkzTjKQSor4sHYOikHIvvi",
			WhitelistedPaths: []string{
				"/",
				healthx.AliveCheckPath,
				healthx.ReadyCheckPath,
				healthx.VersionPath,

				relationtuple.RouteBase,
				check.RouteBase,
				expand.RouteBase,
			},
			BuildVersion: config.Version,
			BuildHash:    config.Commit,
			BuildTime:    config.Date,
			Config: &analytics.Config{
				Endpoint: "https://sqa.ory.sh",
			},
		},
	)
}

func (r *RegistryDefault) ServeAllSQA(cmd *cobra.Command) error {
	r.enableSqa(cmd)
	return r.ServeAll(cmd.Context())
}

func (r *RegistryDefault) ServeAll(ctx context.Context) error {
	eg := &errgroup.Group{}

	eg.Go(r.ServeRead(ctx))
	eg.Go(r.ServeWrite(ctx))
	eg.Go(r.ServeMetrics(ctx))
	return eg.Wait()
}

func (r *RegistryDefault) ServeRead(ctx context.Context) func() error {
	rt, s := r.ReadRouter(), r.ReadGRPCServer()

	return func() error {
		return multiplexPort(ctx, r.Config().ReadAPIListenOn(), rt, s)
	}
}

func (r *RegistryDefault) ServeWrite(ctx context.Context) func() error {
	rt, s := r.WriteRouter(), r.WriteGRPCServer()

	return func() error {
		return multiplexPort(ctx, r.Config().WriteAPIListenOn(), rt, s)
	}
}

func (r *RegistryDefault) ServeMetrics(ctx context.Context) func() error {
	return func() error {
		graceful.WithDefaults(&http.Server{
			Handler: r.MetricsRouter(),
			Addr:    r.Config().MetricsListenOn(),
		})
	}
}

func multiplexPort(ctx context.Context, addr string, router http.Handler, grpcS *grpc.Server) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	eg := &errgroup.Group{}
	ctx, cancel := context.WithCancel(ctx)
	c := 1
	m := cmux.New(l)
	m.SetReadTimeout(graceful.DefaultReadTimeout)

	if grpcS != nil {
		grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
		serversDone := make(chan struct{}, c+1)
		eg.Go(func() error {
			defer func() {
				serversDone <- struct{}{}
			}()
			return errors.WithStack(grpcS.Serve(grpcL))
		})
	}

	httpL := m.Match(cmux.HTTP1())
	restS := graceful.WithDefaults(&http.Server{
		Handler: router,
	})
	serversDone := make(chan struct{}, c)

	eg.Go(func() error {
		defer func() {
			serversDone <- struct{}{}
		}()
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
		for i := 0; i < c; i++ {
			<-serversDone
		}
		if grpcS != nil {
			// we have to stop the servers as well as they might still be running (for whatever reason I could not figure out)
			grpcS.GracefulStop()
		}
		ctx, cancel := context.WithTimeout(context.Background(), graceful.DefaultReadTimeout)
		defer cancel()
		return restS.Shutdown(ctx)
	})

	if err := eg.Wait(); !errors.Is(err, cmux.ErrServerClosed) &&
		!errors.Is(err, cmux.ErrListenerClosed) &&
		(err != nil && !strings.Contains(err.Error(), "use of closed network connection")) {
		// unexpected error
		return err
	}
	return nil
}

func multiplexPortNoGRPC(ctx context.Context, addr string, router http.Handler) error {
	return multiplexPort(ctx, addr, router, nil)
}
