// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/gofrs/uuid"

	"github.com/ory/x/configx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/otelx"
	"github.com/ory/x/tlsx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/x/dbx"
	"github.com/ory/keto/ketoctx"
)

// createFile writes the content to a temporary file, returning the path.
// Good for testing config files.
func createFile(t testing.TB, content string) (path string) {
	t.Helper()

	f, err := os.CreateTemp(t.TempDir(), "config-*.yaml")
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = os.Remove(f.Name()) })

	n, err := f.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(content) {
		t.Fatal("failed to write the complete content")
	}

	return f.Name()
}

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
		tracerWrapper:             options.TracerWrapper,
		ctxer:                     options.Contextualizer(),
		defaultUnaryInterceptors:  options.GRPCUnaryInterceptors(),
		defaultStreamInterceptors: options.GRPCStreamInterceptors(),
		defaultGRPCServerOptions:  options.GRPCServerOptions(),
		defaultHttpMiddlewares:    options.HTTPMiddlewares(),
		extraMigrations:           options.ExtraMigrations(),
		defaultMigrationOptions:   options.MigrationOptions(),
		healthReadyCheckers:       options.ReadyCheckers(),
	}

	init := r.Init
	if withoutNetwork {
		init = r.InitWithoutNetworkID
	}
	if err := init(ctx); err != nil {
		return nil, errors.Wrap(err, "unable to initialize service registry")
	}

	if inspect := options.Inspect(); inspect != nil {
		if err := inspect(r.Persister().Connection(ctx)); err != nil {
			return nil, errors.Wrap(err, "inspect")
		}
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

func NewCRDBTestRegistry(t testing.TB) *RegistryDefault {
	var buf [20]byte
	_, _ = rand.Read(buf[:])
	testdb := fmt.Sprintf("testdb_%x", buf)
	return NewTestRegistry(t, &dbx.DsnT{
		Name:        "cockroach",
		Conn:        dbx.RunCockroach(t, testdb),
		MigrateUp:   true,
		MigrateDown: true,
	})
}

type TestRegistryOption func(t testing.TB, r *RegistryDefault)

func WithConfig(key string, value any) TestRegistryOption {
	return func(t testing.TB, r *RegistryDefault) {
		require.NoError(t, r.c.Set(key, value))
	}
}
func WithNamespaces(namespaces []*namespace.Namespace) TestRegistryOption {
	return func(t testing.TB, r *RegistryDefault) {
		require.NoError(t, r.c.Set(config.KeyNamespaces, namespaces))
	}
}
func WithOPL(opl string) TestRegistryOption {
	return func(t testing.TB, r *RegistryDefault) {
		f := createFile(t, opl)
		require.NoError(t, r.c.Set(config.KeyNamespaces+".location", "file://"+f))
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
func WithTracer(tracer trace.Tracer) TestRegistryOption {
	return func(_ testing.TB, r *RegistryDefault) {
		r.tracer = new(otelx.Tracer).WithOTLP(tracer)
	}
}
func WithLogLevel(level string) TestRegistryOption {
	return func(t testing.TB, r *RegistryDefault) {
		require.NoError(t, r.c.Set("log.level", level))
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
		config.KeyDSN:               dsn.Conn,
		"log.level":                 "info",
		config.KeyNamespaces:        []*namespace.Namespace{},
		config.KeySecretsPagination: []string{uuid.Must(uuid.NewV4()).String()},
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
