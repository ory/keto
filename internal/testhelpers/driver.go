package testhelpers

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	keysetpagination "github.com/ory/x/pagination/keysetpagination_v2"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/relationtuple"
)

func CustomMapperNamespace(ctx context.Context) uuid.UUID {
	return uuid.Nil
}

func RegistryWithManagerWrapper(t *testing.T, reg driver.Registry, pageOpts ...keysetpagination.Option) (*Registry, *relationtuple.ManagerWrapper) {
	mw := relationtuple.NewManagerWrapper(t, reg, pageOpts...)
	return &Registry{mw: mw, Registry: reg}, mw
}

type Registry struct {
	mw *relationtuple.ManagerWrapper // managerProvider
	driver.Registry
}

func (d *Registry) RelationTupleManager() relationtuple.Manager {
	return d.mw
}
