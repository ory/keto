package driver

import (
	"github.com/ory/herodot"
	"github.com/ory/x/logrusx"

	"github.com/ory/keto/persistence/memory"
	"github.com/ory/keto/relation"
	"github.com/ory/keto/x"
)

var _ relation.ManagerProvider = &RegistryDefault{}
var _ x.WriterProvider = &RegistryDefault{}
var _ x.LoggerProvider = &RegistryDefault{}

type RegistryDefault struct {
	p *memory.Persister
	l *logrusx.Logger
	w herodot.Writer
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

func (r *RegistryDefault) RelationManager() relation.Manager {
	if r.p == nil {
		r.p = memory.NewPersister()
	}
	return r.p
}
