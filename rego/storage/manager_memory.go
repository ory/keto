package storage

import (
	"sync"
	"github.com/ory/herodot"
	"github.com/pkg/errors"
	"encoding/json"
	"bytes"
	"context"
	"github.com/ory/pagination"
	"fmt"
	"github.com/open-policy-agent/opa/storage"
	"github.com/open-policy-agent/opa/storage/inmem"
)

type MemoryManager struct {
	sync.RWMutex
	items map[string][]Item
}

type Item struct {
	Key  string
	Data json.RawMessage
}

func NewMemoryManager() *MemoryManager {
	return &MemoryManager{
		items: map[string][]Item{},
	}
}

func (m *MemoryManager) collection(collection string) []Item {
	m.RLock()
	v, ok := m.items[collection]
	m.RUnlock()
	if !ok {
		m.Lock()
		v = []Item{}
		m.items[collection] = v
		m.Unlock()
	}
	return v
}

func (m *MemoryManager) Upsert(_ context.Context, collection, key string, value interface{}) error {
	b := bytes.NewBuffer(nil)
	if err := json.NewEncoder(b).Encode(value); err != nil {
		return errors.WithStack(err)
	}

	// no need to evaluate, just create collection if necessary.
	m.collection(collection)
	m.Lock()
	m.items[collection] = append(m.items[collection], Item{Key: key, Data: b.Bytes()})
	m.Unlock()

	return nil
}

func (m *MemoryManager) List(ctx context.Context, collection string, value interface{}, limit, offset int) error {
	c := m.collection(collection)
	start, end := pagination.Index(limit, offset, len(c))
	b := bytes.NewBuffer(nil)
	enc := json.NewEncoder(b)
	dec := json.NewDecoder(b)
	dec.DisallowUnknownFields()

	items := make([]json.RawMessage, end-start)

	m.RLock()
	for k, i := range c[start:end] {
		items[k] = i.Data
	}
	m.RUnlock()

	if err := enc.Encode(&items); err != nil {
		return errors.WithStack(err)
	}

	fmt.Printf("Got: %s", b.String())

	if err := dec.Decode(value); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (m *MemoryManager) Get(_ context.Context, collection, key string, value interface{}) error {
	c := m.collection(collection)

	m.RLock()
	defer m.RUnlock()

	var v []byte
	for _, i := range c {
		if i.Key == key {
			v = i.Data
			break
		}
	}

	if len(v) == 0 {
		return errors.WithStack(&herodot.ErrorNotFound)
	}

	b := bytes.NewBuffer(v)
	d := json.NewDecoder(b)
	d.DisallowUnknownFields()
	if err := d.Decode(value); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (m *MemoryManager) Delete(_ context.Context, collection, key string) error {
	// no need to evaluate, just create collection if necessary.
	m.collection(collection)

	m.Lock()
	for k, i := range m.items[collection] {
		if i.Key == key {
			m.items[collection] = append(m.items[collection][:k], m.items[collection][k+1:]...)
			break
		}
	}
	m.Unlock()

	return nil
}

func (m *MemoryManager) Storage(data interface{}) (storage.Store, error) {
	db := inmem.New()
	ctx := context.Background()
	txn, err := db.NewTransaction(ctx, storage.WriteParams)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if err := db.Write(ctx, txn, storage.AddOp, storage.Path{}, data); err != nil {
		return nil, errors.WithStack(err)
	}
	if err := db.Commit(ctx, txn); err != nil {
		return nil, errors.WithStack(err)
	}
	return db, nil
}