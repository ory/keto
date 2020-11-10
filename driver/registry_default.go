package driver

import (
	"github.com/ory/herodot"
	"github.com/ory/x/logrusx"

	"github.com/ory/keto/persistence"

	"github.com/ory/keto/expand"

	"github.com/ory/keto/check"

	"github.com/ory/keto/persistence/memory"
	"github.com/ory/keto/relationtuple"
	"github.com/ory/keto/x"
)

var _ relationtuple.ManagerProvider = &RegistryDefault{}
var _ x.WriterProvider = &RegistryDefault{}
var _ x.LoggerProvider = &RegistryDefault{}

type RegistryDefault struct {
	p  persistence.Persister
	l  *logrusx.Logger
	w  herodot.Writer
	ce *check.Engine
	ee *expand.Engine
}

func (r *RegistryDefault) Logger() *logrusx.Logger {
	if r.l == nil {
		r.l = logrusx.New("keto", "dev")
	}
	return r.l
}

func (r *RegistryDefault) Writer() herodot.Writer {
	if r.w == nil {
		r.w = herodot.NewJSONWriter(r.Logger())
	}
	return r.w
}

func (r *RegistryDefault) RelationTupleManager() relationtuple.Manager {
	if r.p == nil {
		r.p = memory.NewPersister()
	}
	return r.p
}

func (r *RegistryDefault) PermissionEngine() *check.Engine {
	if r.ce == nil {
		r.ce = check.NewEngine(r)
	}
	return r.ce
}

func (r *RegistryDefault) ExpandEngine() *expand.Engine {
	if r.ee == nil {
		r.ee = expand.NewEngine(r)
	}
	return r.ee
}
