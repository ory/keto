package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"sync"

	"github.com/open-policy-agent/opa/storage"
	"github.com/pkg/errors"

	"github.com/ory/herodot"
	"github.com/ory/x/pagination"
)

type MemoryManager struct {
	sync.RWMutex
	items map[string][]memoryItem
}

type memoryItem struct {
	Key  string
	Data json.RawMessage
}

func NewMemoryManager() *MemoryManager {
	return &MemoryManager{
		items: map[string][]memoryItem{},
	}
}

func (m *MemoryManager) collection(collection string) []memoryItem {
	m.RLock()
	v, ok := m.items[collection]
	m.RUnlock()
	if !ok {
		m.Lock()
		v = []memoryItem{}
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
	defer m.Unlock()

	var found bool
	for k, i := range m.items[collection] {
		if i.Key == key {
			m.items[collection][k].Data = b.Bytes()
			found = true
			break
		}
	}
	if !found {
		m.items[collection] = append(m.items[collection], memoryItem{Key: key, Data: b.Bytes()})
	}

	return nil
}

func (m *MemoryManager) List(ctx context.Context, collection string, value interface{}, limit, offset int) error {
	c := m.collection(collection)
	start, end := pagination.Index(limit, offset, len(c))
	items := m.list(ctx, collection)[start:end]
	return roundTrip(&items, value)
}

func (m *MemoryManager) list(ctx context.Context, collection string) []json.RawMessage {
	c := m.collection(collection)
	items := make([]json.RawMessage, len(c))

	m.RLock()
	for k, i := range c {
		items[k] = i.Data
	}
	m.RUnlock()

	return items
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
		return errors.WithStack(&herodot.ErrNotFound)
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

func (m *MemoryManager) Storage(ctx context.Context, schema string, collections []string) (storage.Store, error) {
	return toRegoStore(ctx, schema, collections, func(i context.Context, s string) ([]json.RawMessage, error) {
		return m.list(i, s), nil
	})
}