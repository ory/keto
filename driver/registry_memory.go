package driver

import (
	"github.com/ory/keto/storage"
	"github.com/ory/x/dbal"
)

type RegistryMemory struct {
	*RegistryBase

	sm storage.Manager
}

var _ Registry = new(RegistryMemory)

func init() {
	dbal.RegisterDriver(NewRegistryMemory())
}

func NewRegistryMemory() *RegistryMemory {
	r := &RegistryMemory{
		RegistryBase: new(RegistryBase),
	}
	r.RegistryBase.with(r)
	return r
}

func (m *RegistryMemory) Init() error {
	return nil
}

func (m *RegistryMemory) CanHandle(dsn string) bool {
	return dsn == "memory"
}

func (m *RegistryMemory) Ping() error {
	return nil
}

func (m *RegistryMemory) StorageManager() storage.Manager {
	if m.sm == nil {
		m.sm = storage.NewMemoryManager()
	}
	return m.sm
}
