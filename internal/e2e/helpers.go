// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/ory/x/configx"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x/dbx"
)

type namespaceTestManager struct {
	reg     driver.Registry
	ctx     context.Context
	nspaces []*namespace.Namespace
}

func (m *namespaceTestManager) add(t *testing.T, nn ...*namespace.Namespace) {
	m.nspaces = append(m.nspaces, nn...)

	require.NoError(t, m.reg.Config(m.ctx).Set(config.KeyNamespaces, m.nspaces))

	t.Cleanup(func() {
		for _, n := range nn {
			require.NoError(t, m.reg.RelationTupleManager().DeleteAllRelationTuples(m.ctx, &relationtuple.RelationQuery{
				Namespace: &n.Name,
			}))
		}
	})
}

func (m *namespaceTestManager) remove(t *testing.T, name string) {
	newNamespaces := make([]*namespace.Namespace, 0, len(m.nspaces))
	for _, n := range m.nspaces {
		if n.Name != name {
			newNamespaces = append(newNamespaces, n)
		}
	}
	m.nspaces = newNamespaces
	require.NoError(t, m.reg.Config(m.ctx).Set(config.KeyNamespaces, m.nspaces))
}

func newInitializedReg(t testing.TB, dsn *dbx.DsnT, cfgOverwrites map[string]interface{}) (context.Context, driver.Registry, *namespaceTestManager, driver.GetAddr) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(func() {
		cancel()
	})

	flags := pflag.NewFlagSet("", pflag.ContinueOnError)
	configx.RegisterConfigFlag(flags, nil)

	cfgValues := map[string]interface{}{
		config.KeyDSN:               dsn.Conn,
		"log.level":                 "debug",
		"log.leak_sensitive_values": true,
		config.KeyReadAPIHost:       "127.0.0.1",
		config.KeyWriteAPIHost:      "127.0.0.1",
		config.KeyOPLSyntaxAPIHost:  "127.0.0.1",
		config.KeyMetricsHost:       "127.0.0.1",
		config.KeyNamespaces:        []*namespace.Namespace{},
		config.KeySecretsPagination: []string{"test pagination secret"},
	}
	for k, v := range cfgOverwrites {
		cfgValues[k] = v
	}

	cf := dbx.ConfigFile(t, cfgValues)
	require.NoError(t, flags.Parse([]string{"--" + configx.FlagConfig, cf}))

	reg, err := driver.NewDefaultRegistry(ctx, flags, true, nil)
	require.NoError(t, err)

	getAddr := driver.UseDynamicPorts(ctx, t, reg)

	require.NoError(t, reg.MigrateUp(ctx))
	assertMigrated(ctx, t, reg)

	return ctx, reg, &namespaceTestManager{
		reg:     reg,
		ctx:     ctx,
		nspaces: []*namespace.Namespace{},
	}, getAddr
}

func assertMigrated(ctx context.Context, t testing.TB, r driver.Registry) {
	mb, err := r.MigrationBox(ctx)
	require.NoError(t, err)
	s, err := mb.Status(ctx)
	require.NoError(t, err)
	assert.False(t, s.HasPending())
}

func startServer(ctx context.Context, t testing.TB, reg driver.Registry) func() {
	// Start the server
	serverCtx, serverCancel := context.WithCancel(ctx)
	serverErr := make(chan error)
	go func() {
		serverErr <- reg.ServeAll(serverCtx)
	}()

	// defer this close function to make sure it is shutdown on test failure as well
	return func() {
		// stop the server
		serverCancel()
		// wait for it to stop
		require.NoError(t, <-serverErr)
	}
}

// convert the struct in `from` to the pointer in `to` using JSON mashal and unmarshal. from and toPtr must have the
// same json field tags.
func convert(from, toPtr any) error {
	raw, err := json.Marshal(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, toPtr)
}

type (
	paginationOptions struct {
		Token string `json:"page_token"`
		Size  int    `json:"page_size"`
	}
	paginationOptionSetter func(*paginationOptions) *paginationOptions
)

func withToken(t string) paginationOptionSetter {
	return func(opts *paginationOptions) *paginationOptions {
		opts.Token = t
		return opts
	}
}

func withSize(size int) paginationOptionSetter {
	return func(opts *paginationOptions) *paginationOptions {
		opts.Size = size
		return opts
	}
}

func getPaginationOptions(modifiers ...paginationOptionSetter) *paginationOptions {
	opts := &paginationOptions{}
	for _, f := range modifiers {
		opts = f(opts)
	}
	return opts
}
