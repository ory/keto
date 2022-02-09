package driver

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ory/x/logrusx"

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
	innerCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	doneShutdown := make(chan struct{}, 3)

	go func() {
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

		select {
		case <-osSignals:
			cancel()
		case <-innerCtx.Done():
		}

		ctx, cancel := context.WithTimeout(context.Background(), graceful.DefaultShutdownTimeout)
		defer cancel()

		nWaitingForShutdown := cap(doneShutdown)
		select {
		case <-ctx.Done():
			return
		case <-doneShutdown:
			nWaitingForShutdown--
			if nWaitingForShutdown == 0 {
				// graceful shutdown done
				return
			}
		}
	}()

	eg := &errgroup.Group{}

	eg.Go(r.serveRead(innerCtx, doneShutdown))
	eg.Go(r.serveWrite(innerCtx, doneShutdown))
	eg.Go(r.serveMetrics(innerCtx, doneShutdown))

	return eg.Wait()
}

func (r *RegistryDefault) serveRead(ctx context.Context, done chan<- struct{}) func() error {
	rt, s := r.ReadRouter(), r.ReadGRPCServer()

	return func() error {
		return multiplexPort(ctx, r.Logger().WithField("endpoint", "read"), r.Config().ReadAPIListenOn(), rt, s, done)
	}
}

func (r *RegistryDefault) serveWrite(ctx context.Context, done chan<- struct{}) func() error {
	rt, s := r.WriteRouter(), r.WriteGRPCServer()

	return func() error {
		return multiplexPort(ctx, r.Logger().WithField("endpoint", "write"), r.Config().WriteAPIListenOn(), rt, s, done)
	}
}

func (r *RegistryDefault) serveMetrics(ctx context.Context, done chan<- struct{}) func() error {
	return func() error {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		eg := &errgroup.Group{}
		s := graceful.WithDefaults(&http.Server{
			Handler: r.MetricsRouter(),
			Addr:    r.Config().MetricsListenOn(),
		})

		eg.Go(func() error {
			if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				return errors.WithStack(err)
			}
			return nil
		})
		eg.Go(func() (err error) {
			defer func() {
				l := r.Logger().WithField("endpoint", "metrics")
				if err != nil {
					l.WithError(err).Error("graceful shutdown failed")
				} else {
					l.Info("gracefully shutdown server")
				}
				done <- struct{}{}
			}()

			<-ctx.Done()
			ctx, cancel := context.WithTimeout(context.Background(), graceful.DefaultShutdownTimeout)
			defer cancel()
			return s.Shutdown(ctx)
		})

		return eg.Wait()
	}
}

func multiplexPort(ctx context.Context, log *logrusx.Logger, addr string, router http.Handler, grpcS *grpc.Server, done chan<- struct{}) error {
	l, err := (&net.ListenConfig{}).Listen(ctx, "tcp", addr)
	if err != nil {
		return err
	}

	m := cmux.New(l)
	m.SetReadTimeout(graceful.DefaultReadTimeout)

	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1())

	restS := graceful.WithDefaults(&http.Server{
		Handler: router,
	})

	eg := &errgroup.Group{}

	eg.Go(func() error {
		if err := grpcS.Serve(grpcL); !errors.Is(err, cmux.ErrServerClosed) {
			return errors.WithStack(err)
		}
		return nil
	})

	eg.Go(func() error {
		if err := restS.Serve(httpL); !errors.Is(err, http.ErrServerClosed) && !errors.Is(err, cmux.ErrServerClosed) {
			return errors.WithStack(err)
		}
		return nil
	})

	eg.Go(func() error {
		err := m.Serve()
		if err != nil && !errors.Is(err, net.ErrClosed) {
			// unexpected error
			return errors.WithStack(err)
		}
		return nil
	})

	eg.Go(func() (err error) {
		defer func() {
			if err != nil {
				log.WithError(err).Error("graceful shutdown failed")
			} else {
				log.Info("gracefully shutdown server")
			}
			done <- struct{}{}
		}()

		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), graceful.DefaultShutdownTimeout)
		defer cancel()

		shutdownEg := errgroup.Group{}
		shutdownEg.Go(func() error {
			// we ignore net.ErrClosed, because a cmux listener's close func is actually the one of the root listener (which is closed in a racy fashion)
			if err := restS.Shutdown(ctx); !(err == nil || errors.Is(err, http.ErrServerClosed) || errors.Is(err, net.ErrClosed)) {
				// unexpected error
				return errors.WithStack(err)
			}
			return nil
		})
		shutdownEg.Go(func() error {
			gracefulDone := make(chan struct{})
			go func() {
				grpcS.GracefulStop()
				close(gracefulDone)
			}()
			select {
			case <-gracefulDone:
				return nil
			case <-ctx.Done():
				grpcS.Stop()
				return errors.New("graceful stop of gRPC server canceled, had to force it")
			}
		})

		return shutdownEg.Wait()
	})

	return eg.Wait()
}
