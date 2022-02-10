package client

import (
	"context"
	"net"
	"testing"

	"github.com/ory/keto/internal/driver/config"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/namespace"
)

type (
	TestServer struct {
		Reg              driver.Registry
		Namespace        *namespace.Namespace
		Addr, FlagRemote string
		Cmd              *cmdx.CommandExecuter
		Server           *grpc.Server
		NewServer        func(ctx context.Context) *grpc.Server

		errG *errgroup.Group
	}
	ServerType string
)

const (
	WriteServer ServerType = "write"
	ReadServer  ServerType = "read"
)

func NewTestServer(t *testing.T, rw ServerType, nspaces []*namespace.Namespace, newCmd func() *cobra.Command) *TestServer {
	ctx := context.Background()
	ts := &TestServer{
		Reg: driver.NewSqliteTestRegistry(t, false),
	}
	require.NoError(t, ts.Reg.Config(ctx).Set(config.KeyNamespaces, nspaces))

	switch rw {
	case ReadServer:
		ts.NewServer = ts.Reg.ReadGRPCServer
		ts.FlagRemote = FlagReadRemote
	case WriteServer:
		ts.NewServer = ts.Reg.WriteGRPCServer
		ts.FlagRemote = FlagWriteRemote
	default:
		t.Logf("Got unknown server type %s", rw)
		t.FailNow()
	}

	ts.Server = ts.NewServer(ctx)

	l, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	ts.Addr = l.Addr().String()

	ts.errG = &errgroup.Group{}
	ts.errG.Go(func() error {
		return ts.Server.Serve(l)
	})

	ts.Cmd = &cmdx.CommandExecuter{
		New:            newCmd,
		PersistentArgs: []string{"--" + ts.FlagRemote, ts.Addr},
		Ctx:            ctx,
	}

	return ts
}

func (ts *TestServer) Shutdown(t *testing.T) {
	ts.Server.GracefulStop()
	require.NoError(t, ts.errG.Wait())
}
