// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"fmt"
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
		Reg     driver.Registry
		Cmd     *cmdx.CommandExecuter
		Servers []*grpc.Server
		errG    *errgroup.Group
	}
	ServerType string
)

const (
	WriteServer ServerType = "write"
	ReadServer  ServerType = "read"
	OplServer   ServerType = "opl"
)

func (st ServerType) FlagName() string {
	switch st {
	case WriteServer:
		return FlagWriteRemote
	case ReadServer:
		return FlagReadRemote
	case OplServer:
		return FlagOplRemote
	default:
		panic(fmt.Sprintf("unknown ServerType %s", st))
	}
}

func NewTestServer(t *testing.T, nspaces []*namespace.Namespace, newCmd func() *cobra.Command, registryOpts ...driver.TestRegistryOption) *TestServer {
	reg := driver.NewSqliteTestRegistry(t, false, append(registryOpts, driver.WithSelfsignedTransportCredentials())...)

	require.NoError(t, reg.Config(t.Context()).Set(config.KeyNamespaces, nspaces))

	ts := &TestServer{
		Reg:  reg,
		errG: &errgroup.Group{},
		Cmd: &cmdx.CommandExecuter{
			New:            newCmd,
			PersistentArgs: []string{"--insecure-skip-hostname-verification=true"},
			Ctx:            t.Context(),
		},
	}
	ts.serve(t)
	return ts
}

func (ts *TestServer) serve(t *testing.T, serverTypes ...ServerType) {
	if len(serverTypes) == 0 {
		serverTypes = []ServerType{WriteServer, ReadServer}
	}

	for _, st := range serverTypes {
		l, err := net.Listen("tcp", "127.0.0.1:0")
		require.NoError(t, err)

		server := ts.NewServer(t, st)
		ts.errG.Go(func() error {
			return server.Serve(l)
		})

		ts.Servers = append(ts.Servers, server)
		ts.Cmd.PersistentArgs = append(ts.Cmd.PersistentArgs, "--"+st.FlagName(), l.Addr().String())
	}
}

func (ts *TestServer) NewServer(t *testing.T, st ServerType) *grpc.Server {
	switch st {
	case ReadServer:
		return ts.Reg.ReadGRPCServer(t.Context())
	case WriteServer:
		return ts.Reg.WriteGRPCServer(t.Context())
	case OplServer:
		return ts.Reg.OplGRPCServer(t.Context())
	default:
		panic(fmt.Sprintf("unknown ServerType %s", st))
	}
}

func (ts *TestServer) Shutdown(t *testing.T) {
	for _, s := range ts.Servers {
		s.GracefulStop()
	}
	require.NoError(t, ts.errG.Wait())
}
