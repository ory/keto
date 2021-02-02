package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/ory/x/watcherx"

	"github.com/ory/keto/internal/namespace"

	"github.com/markbates/pkger"
	_ "github.com/ory/jsonschema/v3/httploader"
	"github.com/ory/x/configx"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/spf13/pflag"

	"github.com/ory/x/logrusx"
	"github.com/ory/x/tracing"
)

const (
	KeyDSN = "dsn"

	KeyReadAPIHost = "serve.read.host"
	KeyReadAPIPort = "serve.read.port"

	KeyWriteAPIHost = "serve.write.host"
	KeyWriteAPIPort = "serve.write.port"

	KeyNamespaces = "namespaces"

	DSNMemory = "sqlite://file::memory:?_fk=true&cache=shared"
)

type (
	Provider struct {
		p                      *configx.Provider
		l                      *logrusx.Logger
		ctx                    context.Context
		nm                     namespace.Manager
		cancelNamespaceManager context.CancelFunc
		nmLock                 sync.RWMutex
	}
)

func New(ctx context.Context, flags *pflag.FlagSet, l *logrusx.Logger) (*Provider, error) {
	sf, err := pkger.Open("github.com/ory/keto:/.schema/config.schema.json")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	schema, err := ioutil.ReadAll(sf)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	kp := &Provider{
		l:   l,
		ctx: ctx,
	}

	kp.p, err = configx.New(
		schema,
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
	)
	if err != nil {
		return nil, err
	}
	l.UseConfig(kp.p)

	return kp, nil
}

func (k *Provider) resetNamespaceManager() {
	k.nmLock.Lock()
	defer k.nmLock.Unlock()

	if k.nm == nil {
		return
	}

	// cancel and delete old namespace manager
	// the next read request will result in a new one being created
	k.cancelNamespaceManager()
	k.nm = nil
}

func (k *Provider) Set(key string, v interface{}) error {
	if err := k.p.Set(key, v); err != nil {
		return err
	}

	if key == KeyNamespaces {
		k.resetNamespaceManager()
	}
	return nil
}

func (k *Provider) ReadAPIListenOn() string {
	return fmt.Sprintf(
		"%s:%d",
		k.p.StringF(KeyReadAPIHost, ""),
		k.p.IntF(KeyReadAPIPort, 4466),
	)
}

func (k *Provider) WriteAPIListenOn() string {
	return fmt.Sprintf(
		"%s:%d",
		k.p.StringF(KeyWriteAPIHost, ""),
		k.p.IntF(KeyWriteAPIPort, 4467),
	)
}

func (k *Provider) CORS() (cors.Options, bool) {
	return k.p.CORS("serve", cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})
}

func (k *Provider) DSN() string {
	dsn := k.p.String(KeyDSN)
	if dsn == "memory" {
		return DSNMemory
	}
	return dsn
}

func (k *Provider) TracingServiceName() string {
	return k.p.StringF("tracing.service_name", "ORY Keto")
}

func (k *Provider) TracingProvider() string {
	return k.p.StringF("tracing.provider", "")
}

func (k *Provider) TracingConfig() *tracing.Config {
	return k.p.TracingConfig("ORY Keto")
}

func (k *Provider) NamespaceManager() (namespace.Manager, error) {
	if k.nm == nil {
		k.nmLock.Lock()
		defer k.nmLock.Unlock()

		var ctx context.Context
		ctx, k.cancelNamespaceManager = context.WithCancel(k.ctx)

		switch nTyped := k.p.GetF(KeyNamespaces, "file://./keto_namespaces").(type) {
		case string:
			var err error
			k.nm, err = NewNamespaceWatcher(ctx, k.l, nTyped)
			if err != nil {
				return nil, err
			}
		case []*namespace.Namespace:
			k.nm = NewMemoryNamespaceManager(nTyped...)
		case []interface{}:
			nEnc, err := json.Marshal(nTyped)
			if err != nil {
				return nil, errors.WithStack(err)
			}

			nn := make([]*namespace.Namespace, len(nTyped))

			if err := json.Unmarshal(nEnc, &nn); err != nil {
				return nil, errors.WithStack(err)
			}

			k.nm = NewMemoryNamespaceManager(nn...)
		default:
			return nil, errors.Errorf("could not create namespace manager from %#v, this indicates an error in the JSON schema that should be reported", nTyped)
		}

		// return here to properly unlock
		return k.nm, nil
	}

	k.nmLock.RLock()
	defer k.nmLock.RUnlock()

	return k.nm, nil
}
