package driver

import (
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/graceful"
	"github.com/pkg/errors"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func (r *RegistryDefault) ServeAll(ctx context.Context) error {
	eg := &errgroup.Group{}

	eg.Go(r.ServeRead(ctx))
	eg.Go(r.ServeWrite(ctx))

	return eg.Wait()
}

func (r *RegistryDefault) ServeRead(ctx context.Context) func() error {
	rt, s := r.ReadRouter().Router, r.ReadGRPCServer()

	return func() error {
		return multiplexPort(ctx, r.Config().ReadAPIListenOn(), rt, s)
	}
}

func (r *RegistryDefault) ServeWrite(ctx context.Context) func() error {
	rt, s := r.WriteRouter().Router, r.WriteGRPCServer()

	return func() error {
		return multiplexPort(ctx, r.Config().WriteAPIListenOn(), rt, s)
	}
}

func multiplexPort(ctx context.Context, addr string, router *httprouter.Router, grpcS *grpc.Server) error {
	l, err := net.Listen("tcp", addr)
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
	ctx, cancel := context.WithCancel(ctx)
	serversDone := make(chan struct{}, 2)

	eg.Go(func() error {
		defer func() {
			serversDone <- struct{}{}
		}()
		return errors.WithStack(grpcS.Serve(grpcL))
	})

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
		for i := 0; i < 2; i++ {
			<-serversDone
		}

		// we have to stop the servers as well as they might still be running (for whatever reason I could not figure out)
		grpcS.GracefulStop()

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
