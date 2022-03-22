package config

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/ory/jsonschema/v3"
	"github.com/ory/x/cmdx"
	"github.com/ory/x/otelx"

	"github.com/ory/keto/embedx"

	"github.com/ory/herodot"

	"github.com/ory/x/watcherx"

	"github.com/ory/keto/internal/namespace"

	_ "github.com/ory/jsonschema/v3/httploader"
	"github.com/ory/x/configx"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/spf13/pflag"

	"github.com/ory/x/logrusx"
)

const (
	KeyDSN = "dsn"

	KeyLimitMaxReadDepth = "limit.max_read_depth"
	KeyReadAPIHost       = "serve.read.host"
	KeyReadAPIPort       = "serve.read.port"

	KeyWriteAPIHost = "serve.write.host"
	KeyWriteAPIPort = "serve.write.port"

	KeyMetricsHost = "serve.metrics.host"
	KeyMetricsPort = "serve.metrics.port"

	KeyNamespaces = "namespaces"

	DSNMemory = "sqlite://file::memory:?_fk=true&cache=shared"
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
			configx.AttachWatcher(config.watcher),
		)...,
	)
	if validationErr := new(jsonschema.ValidationError); errors.As(err, &validationErr) {
		// the configx provider already printed the validation error
		return nil, cmdx.ErrNoPrintButFail
	} else if err != nil {
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

func (k *Config) watcher(_ watcherx.Event, err error) {
	if err != nil {
		return
	}
	nm, err := k.NamespaceManager()
	if err != nil {
		k.l.WithError(err).Error("got internal error in config watcher: could not get namespace manager")
		return
	}

	nn, err := k.getNamespaces()
	if err != nil {
		k.l.WithError(err).Error("could not get namespaces from config")
		return
	}
	if nm.ShouldReload(nn) {
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

func (k *Config) Set(key string, v interface{}) error {
	if err := k.p.Set(key, v); err != nil {
		return err
	}

	if key == KeyNamespaces {
		k.resetNamespaceManager()
	}
	return nil
}

func (k *Config) ReadAPIListenOn() string {
	return fmt.Sprintf(
		"%s:%d",
		k.p.StringF(KeyReadAPIHost, ""),
		k.p.IntF(KeyReadAPIPort, 4466),
	)
}

func (k *Config) MaxReadDepth() int {
	return k.p.Int(KeyLimitMaxReadDepth)
}

func (k *Config) WriteAPIListenOn() string {
	return fmt.Sprintf(
		"%s:%d",
		k.p.StringF(KeyWriteAPIHost, ""),
		k.p.IntF(KeyWriteAPIPort, 4467),
	)
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

func (k *Config) TracingServiceName() string {
	return k.p.StringF("tracing.service_name", "Ory Keto")
}

func (k *Config) TracingProvider() string {
	return k.p.StringF("tracing.provider", "")
}

func (k *Config) TracingConfig() *otelx.Config {
	return k.p.TracingConfigOtel("Ory Keto")
}

func (k *Config) NamespaceManager() (namespace.Manager, error) {
	k.nmLock.Lock()
	defer k.nmLock.Unlock()

	if k.nm == nil {
		var ctx context.Context
		ctx, k.cancelNamespaceManager = context.WithCancel(k.ctx)

		nn, err := k.getNamespaces()
		if err != nil {
			return nil, err
		}

		switch nTyped := nn.(type) {
		case string:
			var err error
			k.nm, err = NewNamespaceWatcher(ctx, k.l, nTyped)
			if err != nil {
				return nil, err
			}
		case []*namespace.Namespace:
			k.nm = NewMemoryNamespaceManager(nTyped...)
		default:
			return nil, errors.WithStack(herodot.ErrInternalServerError.WithReasonf("got unexpected namespaces type %T", nn))
		}
	}

	return k.nm, nil
}

// getNamespaces returns string or []*namespace.Namespace
func (k *Config) getNamespaces() (interface{}, error) {
	switch nTyped := k.p.GetF(KeyNamespaces, "file://./keto_namespaces").(type) {
	case string:
		return nTyped, nil
	case []*namespace.Namespace:
		return nTyped, nil
	case []interface{}:
		nEnc, err := json.Marshal(nTyped)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		nn := make([]*namespace.Namespace, len(nTyped))

		if err := json.Unmarshal(nEnc, &nn); err != nil {
			return nil, errors.WithStack(err)
		}

		return nn, nil
	default:
		return nil, errors.WithStack(herodot.ErrInternalServerError.WithReasonf("could not infer namespaces for type %T", nTyped))
	}
}

func (k *Config) MetricsListenOn() string {
	return fmt.Sprintf(
		"%s:%d",
		k.p.StringF(KeyMetricsHost, ""),
		k.p.IntF(KeyMetricsPort, 4468),
	)
}
