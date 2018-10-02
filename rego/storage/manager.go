package storage

import (
	"context"
	"github.com/open-policy-agent/opa/storage"
)

type Manager interface {
	Get(ctx context.Context, collection string, key string, value interface{}) error
	List(ctx context.Context, collection string, value interface{}, limit, offset int) error
	Upsert(ctx context.Context, collection string, key string, value interface{}) error
	Delete(ctx context.Context, collection string, key string) error
	Storage(ctx context.Context, schema string, collections []string) (storage.Store, error)
}
