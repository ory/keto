// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"sync"
	"testing"

	"github.com/ory/x/configx"
	"github.com/ory/x/tlsx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/ory/keto/ketoctx"

	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/x/dbx"

	"github.com/sirupsen/logrus"

	"github.com/pkg/errors"

	"github.com/ory/x/logrusx"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver/config"
)

func NewDefaultRegistry(ctx context.Context, flags *pflag.FlagSet, withoutNetwork bool, opts []ketoctx.Option) (Registry, error) {
	reg, ok := ctx.Value(RegistryContextKey).(Registry)
	if ok {
		return reg, nil
	}

	options := ketoctx.Options(opts...)

	l := options.Logger()
	if l == nil {
		l = newLogger(ctx)
	}

	c := config.New(ctx, l, nil)
	cp, err := config.NewProvider(ctx, flags, c)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize config provider")
	}
	c.WithSource(options.Contextualizer().Config(ctx, cp))

	r := &RegistryDefault{
		c:                         c,
		l:                         l,
		ctxer:                     options.Contextualizer(),
		defaultUnaryInterceptors:  options.GRPCUnaryInterceptors(),
		defaultStreamInterceptors: options.GRPCStreamInterceptors(),
		defaultHttpMiddlewares:    options.HTTPMiddlewares(),
		defaultMigrationOptions:   options.MigrationOptions(),
	}

	init := r.Init
	if withoutNetwork {
		init = r.InitWithoutNetworkID
	}
	if err := init(ctx); err != nil {
		return nil, errors.Wrap(err, "unable to initialize service registry")
	}

	return r, nil
}

func NewSqliteTestRegistry(t testing.TB, debugOnDisk bool, opts ...TestRegistryOption) *RegistryDefault {
	mode := dbx.SQLiteMemory
	if debugOnDisk {
		mode = dbx.SQLiteDebug
	}
	return NewTestRegistry(t, dbx.GetSqlite(t, mode), opts...)
}

type TestRegistryOption func(t testing.TB, r *RegistryDefault)

func WithNamespaces(namespaces []*namespace.Namespace) TestRegistryOption {
	return func(t testing.TB, r *RegistryDefault) {
		require.NoError(t, r.c.Set(config.KeyNamespaces, namespaces))
	}
}
func WithGRPCUnaryInterceptors(i ...grpc.UnaryServerInterceptor) TestRegistryOption {
	return func(_ testing.TB, r *RegistryDefault) {
		r.defaultUnaryInterceptors = i
	}
}

func WithGRPCStreamInterceptors(i ...grpc.StreamServerInterceptor) TestRegistryOption {
	return func(_ testing.TB, r *RegistryDefault) {
		r.defaultStreamInterceptors = i
	}
}

type selfSignedCert struct {
	once sync.Once
	cert *tls.Certificate
	err  error
}

var sharedTestCert selfSignedCert

func (s *selfSignedCert) generate() {
	s.once.Do(func() {
		key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			s.err = fmt.Errorf("could not create key: %v", err)
			return
		}

		s.cert, err = tlsx.CreateSelfSignedTLSCertificate(key)
		if err != nil {
			s.err = fmt.Errorf("could not create TLS certificate: %v", err)
		}
	})
}

func WithSelfsignedTransportCredentials() TestRegistryOption {
	return func(t testing.TB, r *RegistryDefault) {
		sharedTestCert.generate()
		if sharedTestCert.err != nil {
			t.Error(sharedTestCert.err)
			return
		}

		r.grpcTransportCredentials = credentials.NewServerTLSFromCert(sharedTestCert.cert)
	}
}

func NewTestRegistry(t testing.TB, dsn *dbx.DsnT, opts ...TestRegistryOption) *RegistryDefault {
	l := logrusx.New("Ory Keto", "testing")
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	ctx = configx.ContextWithConfigOptions(ctx, configx.WithValues(map[string]interface{}{
		config.KeyDSN:        dsn.Conn,
		"log.level":          "debug",
		config.KeyNamespaces: []*namespace.Namespace{},
	}))
	c, err := config.NewDefault(ctx, nil, l)
	require.NoError(t, err)

	r := &RegistryDefault{
		c:     c,
		l:     l,
		ctxer: &ketoctx.DefaultContextualizer{},
	}

	for _, opt := range opts {
		opt(t, r)
	}

	if dsn.MigrateUp {
		require.NoError(t, r.MigrateUp(ctx))
	}

	require.NoError(t, r.Init(ctx))

	return r
}

func newLogger(ctx context.Context) *logrusx.Logger {
	hook, ok := ctx.Value(LogrusHookContextKey).(logrus.Hook)

	var opts []logrusx.Option
	if ok {
		opts = append(opts, logrusx.WithHook(hook))
	}

	return logrusx.New("Ory Keto", config.Version, opts...)
}
