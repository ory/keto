// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
)

type (
	TestServer struct {
		Reg driver.Registry
		Cmd *cmdx.CommandExecuter
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
	reg := driver.NewSqliteTestRegistry(t, append(registryOpts, driver.WithSelfSignedTransportCredentials())...)

	require.NoError(t, reg.Config(t.Context()).Set(config.KeyNamespaces, nspaces))

	ts := &TestServer{
		Reg: reg,
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
		s := httptest.NewUnstartedServer(ts.NewHandler(t, st))
		s.EnableHTTP2 = true
		s.StartTLS()
		t.Cleanup(s.Close)

		ts.Cmd.PersistentArgs = append(ts.Cmd.PersistentArgs, "--"+st.FlagName(), s.Listener.Addr().String())
	}
}

func (ts *TestServer) NewHandler(t *testing.T, st ServerType) http.Handler {
	switch st {
	case ReadServer:
		return ts.Reg.ReadRouter(t.Context())
	case WriteServer:
		return ts.Reg.WriteRouter(t.Context())
	case OplServer:
		return ts.Reg.OPLSyntaxRouter(t.Context())
	default:
		panic(fmt.Sprintf("unknown ServerType %s", st))
	}
}
