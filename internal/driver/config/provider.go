package config

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/ory/herodot"

	"github.com/ory/x/watcherx"

	"github.com/ory/keto/internal/namespace"

	_ "github.com/ory/jsonschema/v3/httploader"
	"github.com/ory/x/configx"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/spf13/pflag"

	"github.com/ory/x/logrusx"
	"github.com/ory/x/tracing"
)

//go:embed config.schema.json
var Schema []byte

const (
	KeyDSN = "dsn"

	KeyReadMaxDepth = "serve.read.max-depth"
	KeyReadAPIHost  = "serve.read.host"
	KeyReadAPIPort  = "serve.read.port"

	KeyWriteAPIHost = "serve.write.host"
	KeyWriteAPIPort = "serve.write.port"

	KeyNamespaces = "namespaces"

	DSNMemory = "sqlite://file::memory:?_fk=true&cache=shared"
)

type (
	Config struct {
		p                      *configx.Provider
		l                      *logrusx.Logger
		ctx                    context.Context
		nm                     namespace.Manager
		cancelNamespaceManager context.CancelFunc
		nmLock                 sync.Mutex
	}
	Provider interface {
		Config() *Config
	}
)

func New(ctx context.Context, flags *pflag.FlagSet, l *logrusx.Logger) (*Config, error) {
	kp := &Config{
		l:   l,
		ctx: ctx,
	}

	var err error
	kp.p, err = configx.New(
		Schema,
		configx.WithFlags(flags),
		configx.WithStderrValidationReporter(),
		configx.WithImmutables(KeyDSN, "serve"),
		configx.OmitKeysFromTracing(KeyDSN),
		configx.WithLogrusWatcher(kp.l),
		configx.WithContext(ctx),
		configx.AttachWatcher(func(watcherx.Event, error) {
			// TODO this can be optimized to run only on changes related to namespace config
			kp.resetNamespaceManager()
		}),
		configx.AttachWatcher(func(watcherx.Event, error) {
			if err != nil {
				return
			}
			nm, err := kp.NamespaceManager()
			if err != nil {
				l.WithError(err).Error("got internal error in config watcher: could not get namespace manager")
				return
			}

			nn, err := kp.getNamespaces()
			if err != nil {
				l.WithError(err).Error("could not get namespaces from config")
				return
			}
			if nm.ShouldReload(nn) {
				kp.resetNamespaceManager()
			}
		}),
	)
	if err != nil {
		return nil, err
	}
	l.UseConfig(kp.p)

	return kp, nil
}

func (k *Config) Source() *configx.Provider {
	return k.p
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

func (k *Config) ReadAPIMaxDepth() int {
	return k.p.Int(KeyReadMaxDepth)
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
	case "read", "write":
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
	return k.p.StringF("tracing.service_name", "ORY Keto")
}

func (k *Config) TracingProvider() string {
	return k.p.StringF("tracing.provider", "")
}

func (k *Config) TracingConfig() *tracing.Config {
	return k.p.TracingConfig("ORY Keto")
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
