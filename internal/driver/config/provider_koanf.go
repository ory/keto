package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

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

	KeyHost = "serve.host"
	KeyPort = "serve.port"

	KeyNamespaces = "namespaces"
)

type (
	KoanfProvider struct {
		p   *configx.Provider
		l   *logrusx.Logger
		ctx context.Context
		nm  namespace.Manager
	}
)

func New(flags *pflag.FlagSet, l *logrusx.Logger) (Provider, error) {
	sf, err := pkger.Open("/.schema/config.schema.json")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	schema, err := ioutil.ReadAll(sf)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ctx := context.Background()

	p, err := configx.New(
		schema,
		flags,
		configx.WithStderrValidationReporter(),
		configx.WithImmutables(KeyDSN, "serve"),
		configx.OmitKeysFromTracing(KeyDSN),
		configx.WithLogrusWatcher(l),
		configx.WithContext(ctx),
	)
	if err != nil {
		return nil, err
	}

	return &KoanfProvider{
		l:   l,
		p:   p,
		ctx: ctx,
	}, nil
}

func (k *KoanfProvider) Set(key string, v interface{}) {
	k.p.Set(key, v)
}

func (k *KoanfProvider) ListenOn() string {
	return fmt.Sprintf(
		"%s:%d",
		k.p.StringF(KeyHost, ""),
		k.p.IntF(KeyPort, 4466),
	)
}

func (k *KoanfProvider) CORS() (cors.Options, bool) {
	return k.p.CORS("serve", cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})
}

func (k *KoanfProvider) DSN() string {
	return k.p.StringF(KeyDSN, "memory")
}

func (k *KoanfProvider) TracingServiceName() string {
	return k.p.StringF("tracing.service_name", "ORY Keto")
}

func (k *KoanfProvider) TracingProvider() string {
	return k.p.StringF("tracing.provider", "")
}

func (k *KoanfProvider) TracingConfig() *tracing.Config {
	return k.p.TracingConfig("ORY Keto")
}

func (k *KoanfProvider) NamespaceManager(ctx context.Context) (namespace.Manager, error) {
	if k.nm == nil {
		switch nTyped := k.p.GetF(KeyNamespaces, "file://./keto_namespaces").(type) {
		case string:
			var err error
			k.nm, err = NewNamespaceWatcher(ctx, k.l, nTyped)
			if err != nil {
				return nil, err
			}
		case []*namespace.Namespace:
			staticManager := staticNamespaces(nTyped)
			k.nm = &staticManager
		case []interface{}:
			nEnc, err := json.Marshal(nTyped)
			if err != nil {
				return nil, errors.WithStack(err)
			}

			nn := make([]*namespace.Namespace, len(nTyped))

			if err := json.Unmarshal(nEnc, &nn); err != nil {
				return nil, errors.WithStack(err)
			}

			staticManager := staticNamespaces(nn)
			k.nm = &staticManager
		default:
			return nil, errors.Errorf("could not create namespace manager from %#v, this indicates an error in the JSON schema that should be reported", nTyped)
		}
	}

	return k.nm, nil
}
