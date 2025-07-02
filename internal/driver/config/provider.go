// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"context"
	"crypto/sha512"
	_ "embed"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/ory/herodot"
	_ "github.com/ory/jsonschema/v3/httploader"
	"github.com/ory/x/configx"
	"github.com/ory/x/fetcher"
	"github.com/ory/x/httpx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/otelx"
	"github.com/ory/x/watcherx"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/spf13/pflag"
	"go.opentelemetry.io/otel/trace/noop"

	"github.com/ory/keto/embedx"
	"github.com/ory/keto/internal/namespace"
)

type EndpointType string

const (
	EndpointRead      EndpointType = "read"
	EndpointWrite     EndpointType = "write"
	EndpointMetrics   EndpointType = "metrics"
	EndpointOPLSyntax EndpointType = "opl"

	KeyDSN = "dsn"

	KeyLimitMaxReadDepth = "limit.max_read_depth"
	KeyLimitMaxReadWidth = "limit.max_read_width"

	KeyBatchCheckMaxBatchSize         = "limit.max_batch_check_size"
	KeyBatchCheckParallelizationLimit = "limit.batch_check_max_parallelization"

	KeyReadAPIHost         = "serve." + string(EndpointRead) + ".host"
	KeyReadAPIPort         = "serve." + string(EndpointRead) + ".port"
	KeyReadAPIListenFile   = "serve." + string(EndpointRead) + ".write_listen_file"
	KeyWriteAPIHost        = "serve." + string(EndpointWrite) + ".host"
	KeyWriteAPIPort        = "serve." + string(EndpointWrite) + ".port"
	KeyWriteAPIListenFile  = "serve." + string(EndpointWrite) + ".write_listen_file"
	KeyOPLSyntaxAPIHost    = "serve." + string(EndpointOPLSyntax) + ".host"
	KeyOPLSyntaxAPIPort    = "serve." + string(EndpointOPLSyntax) + ".port"
	KeyOPLSyntaxListenFile = "serve." + string(EndpointOPLSyntax) + ".write_listen_file"
	KeyMetricsHost         = "serve." + string(EndpointMetrics) + ".host"
	KeyMetricsPort         = "serve." + string(EndpointMetrics) + ".port"
	KeyMetricsListenFile   = "serve." + string(EndpointMetrics) + ".write_listen_file"

	KeyNamespaces                       = "namespaces"
	KeyNamespacesExperimentalStrictMode = KeyNamespaces + ".experimental_strict_mode"

	DSNMemory = "sqlite://file::memory:?_fk=true&cache=shared"

	KeySecretsPagination = "secrets.pagination"
)

type (
	Config struct {
		p   *configx.Provider
		l   *logrusx.Logger
		ctx context.Context

		nm                     namespace.Manager
		cancelNamespaceManager context.CancelFunc
		nmLock                 sync.Mutex
	}
	Provider interface {
		Config(ctx context.Context) *Config
	}
)

func New(ctx context.Context, l *logrusx.Logger, p *configx.Provider) *Config {
	return &Config{
		p:   p,
		l:   l,
		ctx: ctx,
	}
}

func NewDefault(ctx context.Context, flags *pflag.FlagSet, l *logrusx.Logger, opts ...configx.OptionModifier) (*Config, error) {
	c := New(ctx, l, nil)
	cp, err := NewProvider(ctx, flags, c, opts...)
	if err != nil {
		return nil, err
	}
	c.WithSource(cp)

	return c, nil
}

func NewProvider(ctx context.Context, flags *pflag.FlagSet, config *Config, opts ...configx.OptionModifier) (*configx.Provider, error) {
	p, err := configx.New(
		ctx,
		embedx.ConfigSchema,
		append(opts,
			configx.WithFlags(flags),
			configx.WithStderrValidationReporter(),
			configx.WithImmutables(KeyDSN, "serve"),
			configx.OmitKeysFromTracing(KeyDSN),
			configx.WithLogrusWatcher(config.l),
			configx.WithContext(ctx),
			configx.AttachWatcher(config.namespaceWatcher),
		)...,
	)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (k *Config) Source() *configx.Provider {
	return k.p
}

func (k *Config) WithSource(p *configx.Provider) {
	k.p = p
	k.l.UseConfig(p)
}

func (k *Config) namespaceWatcher(_ watcherx.Event, err error) {
	if err != nil {
		return
	}
	nm, err := k.NamespaceManager()
	if err != nil {
		k.l.WithError(err).Error("got internal error in config namespace watcher: could not get namespace manager")
		return
	}

	nnCfg, err := k.namespaceConfig()
	if err != nil {
		k.l.WithError(err).Error("could not get namespaces from config")
		return
	}
	if nm.ShouldReload(nnCfg.value()) {
		k.resetNamespaceManager()
	}
}

func (k *Config) resetNamespaceManager() {
	k.nmLock.Lock()
	defer k.nmLock.Unlock()

	if k.cancelNamespaceManager == nil {
		return
	}

	// cancel and delete old namespace manager
	// the next read request will result in a new one being created
	k.cancelNamespaceManager()
	k.nm, k.cancelNamespaceManager = nil, nil
}

func (k *Config) Set(key string, v any) error {
	if err := k.p.Set(key, v); err != nil {
		return err
	}

	if key == KeyNamespaces {
		k.resetNamespaceManager()
	}
	return nil
}

func (k *Config) addressFor(endpoint EndpointType) (addr string, listenFile string) {
	return fmt.Sprintf(
		"%s:%d",
		k.p.StringF("serve."+string(endpoint)+".host", ""),
		k.p.IntF("serve."+string(endpoint)+".port", 0),
	), k.p.StringF("serve."+string(endpoint)+".write_listen_file", "")
}

func (k *Config) ReadAPIListenOn() (addr string, listenFile string) {
	return k.addressFor(EndpointRead)
}
func (k *Config) WriteAPIListenOn() (addr string, listenFile string) {
	return k.addressFor(EndpointWrite)
}
func (k *Config) MetricsListenOn() (addr string, listenFile string) {
	return k.addressFor(EndpointMetrics)
}
func (k *Config) OPLSyntaxAPIListenOn() (addr string, listenFile string) {
	return k.addressFor(EndpointOPLSyntax)
}

func (k *Config) MaxReadDepth() int {
	return k.p.Int(KeyLimitMaxReadDepth)
}
func (k *Config) MaxReadWidth() int {
	return k.p.Int(KeyLimitMaxReadWidth)
}

func (k *Config) BatchCheckMaxBatchSize() int {
	return k.p.Int(KeyBatchCheckMaxBatchSize)
}

func (k *Config) BatchCheckParallelizationLimit() int {
	return k.p.Int(KeyBatchCheckParallelizationLimit)
}

func (k *Config) CORS(iface string) (cors.Options, bool) {
	switch iface {
	case "read", "write", "metrics":
	default:
		panic("expected interface 'read' or 'write', but got unknown interface " + iface)
	}

	return k.p.CORS("serve."+iface, cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})
}

func (k *Config) DSN() string {
	dsn := k.p.String(KeyDSN)
	if dsn == "memory" {
		return DSNMemory
	}
	return dsn
}

func (k *Config) Fetcher() *fetcher.Fetcher {
	// Tracing still works correctly even though we pass a no-op tracer
	// here, because the otelhttp package will preferentially use the
	// tracer from the incoming request context over this one.
	opts := []httpx.ResilientOptions{httpx.ResilientClientWithTracer(noop.NewTracerProvider().Tracer("keto/internal/driver/config"))}
	if k.p.Bool("clients.http.disallow_private_ip_ranges") {
		opts = append(opts, httpx.ResilientClientDisallowInternalIPs())
	}
	return fetcher.NewFetcher(
		fetcher.WithClient(httpx.NewResilientClient(opts...)),
	)
}

func (k *Config) TracingServiceName() string {
	return k.p.StringF("tracing.service_name", "Ory Keto")
}

func (k *Config) TracingProvider() string {
	return k.p.StringF("tracing.provider", "")
}

func (k *Config) TracingConfig() *otelx.Config {
	return k.p.TracingConfig("Ory Keto")
}

func (k *Config) NamespaceManager() (namespace.Manager, error) {
	k.nmLock.Lock()
	defer k.nmLock.Unlock()

	if k.nm == nil {
		var ctx context.Context
		ctx, k.cancelNamespaceManager = context.WithCancel(k.ctx)

		nnCfg, err := k.namespaceConfig()
		if err != nil {
			return nil, err
		}

		k.nm, err = nnCfg.newManager()(ctx, k)
		if err != nil {
			return nil, err
		}
	}

	return k.nm, nil
}

func (k *Config) StrictMode() bool {
	return k.p.BoolF(KeyNamespacesExperimentalStrictMode, false)
}

type (
	buildNamespaceFn func(context.Context, *Config) (namespace.Manager, error)

	namespaceConfig interface {
		// newManager builds a new namespace manager.
		newManager() buildNamespaceFn
		// value returns the wrapped value (for comparing if we should reload)
		value() any
	}

	legacyURINamespaceConfig string
	literalNamespaceConfig   []*namespace.Namespace
	oplNamespaceConfig       map[string]any
)

func (uri legacyURINamespaceConfig) newManager() buildNamespaceFn {
	return func(ctx context.Context, c *Config) (namespace.Manager, error) {
		return NewNamespaceWatcher(ctx, c.l, string(uri))
	}
}
func (uri legacyURINamespaceConfig) value() any {
	return string(uri)
}

func (namespaces literalNamespaceConfig) newManager() buildNamespaceFn {
	return func(ctx context.Context, _ *Config) (namespace.Manager, error) {
		return NewMemoryNamespaceManager(namespaces...), nil
	}
}
func (namespaces literalNamespaceConfig) value() any {
	return []*namespace.Namespace(namespaces)
}

func (oplConfig oplNamespaceConfig) newManager() buildNamespaceFn {
	return func(ctx context.Context, c *Config) (namespace.Manager, error) {
		entry, ok := oplConfig["location"]
		if !ok {
			return nil, errors.New("location key not found")
		}
		target, ok := entry.(string)
		if !ok {
			return nil, fmt.Errorf("config value must be string, was %T", entry)
		}
		return newOPLConfigWatcher(ctx, c, target)
	}
}
func (oplConfig oplNamespaceConfig) value() any {
	return map[string]any(oplConfig)
}

// namespaceConfig returns a namespace config, which can be either a URI (in
// which case we want to watch that URI), or a literal list of namespaces (in
// which case we just load them into memory), or a list of URIs referencing OPL
// definitions (in which case we want to watch each URI and parse the content).
func (k *Config) namespaceConfig() (namespaceConfig, error) {
	switch nTyped := k.p.GetF(KeyNamespaces, "file://./keto_namespaces").(type) {
	case string:
		return legacyURINamespaceConfig(nTyped), nil

	case []*namespace.Namespace:
		return literalNamespaceConfig(nTyped), nil

	case []any:
		nEnc, err := json.Marshal(nTyped)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		nn := make([]*namespace.Namespace, len(nTyped))

		if err := json.Unmarshal(nEnc, &nn); err != nil {
			return nil, errors.WithStack(err)
		}
		return literalNamespaceConfig(nn), nil

	case map[string]any:
		return oplNamespaceConfig(nTyped), nil

	default:
		return nil, errors.WithStack(herodot.ErrInternalServerError.WithReasonf("could not infer namespaces for type %T", nTyped))
	}
}

func (k *Config) PaginationEncryptionKeys() [][32]byte {
	secrets := k.p.Strings(KeySecretsPagination)

	encryptionKeys := make([][32]byte, len(secrets))
	for i, key := range secrets {
		encryptionKeys[i] = sha512.Sum512_256([]byte(key))
	}

	return encryptionKeys
}
