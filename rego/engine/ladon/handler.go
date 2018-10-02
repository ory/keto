package ladon

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/open-policy-agent/opa/rego"
	"github.com/ory/go-convenience/stringslice"
	"github.com/ory/go-convenience/stringsx"
	"github.com/ory/herodot"
	"github.com/ory/keto/rego/engine"
	kstorage "github.com/ory/keto/rego/storage"
	"github.com/pkg/errors"
	"net/http"
)

type Engine struct {
	sh     *kstorage.Handler
	engine *engine.Engine
	s      kstorage.Manager
	h      herodot.Writer
}

const (
	basePath = "/engines/lacp/:flavor"
	schema   = `{
	"store": {
		"ladon": {
			"regex": {
				"policies": [],
				"roles": []
			},
			"regex": {
				"policies": [],
				"roles": []
			}
		}
	}
}`
)

func policyCollection(f string) (string) {
	return fmt.Sprintf("/store/ladon/%s/policies", f)
}

func roleCollection(f string) (string) {
	return fmt.Sprintf("/store/ladon/%s/roles", f)
}

func NewEngine(store kstorage.Manager, sh *kstorage.Handler, e *engine.Engine, h herodot.Writer) *Engine {
	return &Engine{
		s:      store,
		h:      h,
		sh:     sh,
		engine: e,
	}
}

func (e *Engine) Register(r *httprouter.Router) {
	r.POST(basePath+"/allowed", e.engine.Evaluate(e.eval))

	r.POST(basePath+"/policies", e.sh.Upsert(e.policiesCreate))
	r.GET(basePath+"/policies", e.sh.List(e.policiesList))
	r.GET(basePath+"/policies/:id", e.sh.Get(e.policiesGet))
	r.DELETE(basePath+"/policies/:id", e.sh.Delete(e.policiesDelete))

	r.GET(basePath+"/roles", e.sh.List(e.rolesList))
	r.GET(basePath+"/roles/:id", e.sh.Get(e.rolesGet))
	r.DELETE(basePath+"/roles/:id", e.sh.Delete(e.policiesDelete))
	r.PUT(basePath+"/roles/:id/members", e.sh.Upsert(e.rolesMembersAdd))
	r.DELETE(basePath+"/roles/:id/members/:member", e.sh.Upsert(e.rolesMembersRemove))
}

func (e *Engine) rolesList(ctx context.Context, r *http.Request, ps httprouter.Params) (*kstorage.ListRequest, error) {
	var p Roles

	f, err := flavor(ps)
	if err != nil {
		return nil, err
	}

	return &kstorage.ListRequest{
		Collection: roleCollection(f),
		Value:      &p,
	}, nil
}

func (e *Engine) rolesGet(ctx context.Context, r *http.Request, ps httprouter.Params) (*kstorage.GetRequest, error) {
	var p Role

	f, err := flavor(ps)
	if err != nil {
		return nil, err
	}

	return &kstorage.GetRequest{
		Collection: roleCollection(f),
		Key:        ps.ByName("id"),
		Value:      &p,
	}, nil
}

func (e *Engine) rolesDelete(ctx context.Context, r *http.Request, ps httprouter.Params) (*kstorage.DeleteRequest, error) {
	f, err := flavor(ps)
	if err != nil {
		return nil, err
	}

	return &kstorage.DeleteRequest{
		Collection: roleCollection(f),
		Key:        ps.ByName("id"),
	}, nil
}

func (e *Engine) rolesMembersAdd(ctx context.Context, r *http.Request, ps httprouter.Params) (*kstorage.UpsertRequest, error) {
	f, err := flavor(ps)
	if err != nil {
		return nil, err
	}

	var i Role
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		return nil, errors.WithStack(err)
	}

	var ro Role
	if err := e.s.Get(ctx, roleCollection(f), ps.ByName("id"), &ro); errors.Cause(err) == &herodot.ErrorNotFound {
		i.ID = ps.ByName("id")
		ro = i
	} else if err != nil {
		return nil, err
	} else {
		ro.Members = stringsx.Unique(append(ro.Members, i.Members...))
	}

	return &kstorage.UpsertRequest{
		Collection: roleCollection(f),
		Key:        ro.ID,
		Value:      &ro,
	}, nil

}
func (e *Engine) rolesMembersRemove(ctx context.Context, r *http.Request, ps httprouter.Params) (*kstorage.UpsertRequest, error) {
	f, err := flavor(ps)
	if err != nil {
		return nil, err
	}

	var ro Role
	if err := e.s.Get(ctx, roleCollection(f), ps.ByName("id"), &ro); err != nil {
		return nil, err
	}

	ro.Members = stringslice.Filter(ro.Members, func(s string) bool {
		return s == ps.ByName("member")
	})

	return &kstorage.UpsertRequest{
		Collection: roleCollection(f),
		Key:        ro.ID,
		Value:      &ro,
	}, nil
}

func (e *Engine) policiesCreate(ctx context.Context, r *http.Request, ps httprouter.Params) (*kstorage.UpsertRequest, error) {
	var p Policy
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		return nil, errors.WithStack(err)
	}

	f, err := flavor(ps)
	if err != nil {
		return nil, err
	}

	return &kstorage.UpsertRequest{
		Collection: policyCollection(f),
		Key:        p.ID,
		Value:      &p,
	}, nil
}

func (e *Engine) policiesList(ctx context.Context, r *http.Request, ps httprouter.Params) (*kstorage.ListRequest, error) {
	var p Policies

	f, err := flavor(ps)
	if err != nil {
		return nil, err
	}

	return &kstorage.ListRequest{
		Collection: policyCollection(f),
		Value:      &p,
	}, nil
}

func (e *Engine) policiesDelete(ctx context.Context, r *http.Request, ps httprouter.Params) (*kstorage.DeleteRequest, error) {
	f, err := flavor(ps)
	if err != nil {
		return nil, err
	}

	return &kstorage.DeleteRequest{
		Collection: policyCollection(f),
		Key:        ps.ByName("id"),
	}, nil
}

func (e *Engine) policiesGet(ctx context.Context, r *http.Request, ps httprouter.Params) (*kstorage.GetRequest, error) {
	var p Policy

	f, err := flavor(ps)
	if err != nil {
		return nil, err
	}

	return &kstorage.GetRequest{
		Collection: policyCollection(f),
		Key:        ps.ByName("id"),
		Value:      &p,
	}, nil
}

type input struct {
	// Resource is the resource that access is requested to.
	Resource string `json:"resource"`

	// Action is the action that is requested on the resource.
	Action string `json:"action"`

	// Subject is the subject that is requesting access.
	Subject string `json:"subject"`

	// Context is the request's environmental context.
	Context map[string]interface{} `json:"context"`
}

func flavor(ps httprouter.Params) (string, error) {
	t := ps.ByName("flavor")
	if t != "regex" && t != "exact" {
		return "", errors.WithStack(&herodot.ErrorNotFound)
	}

	return t, nil
}

func (e *Engine) eval(ctx context.Context, r *http.Request, ps httprouter.Params) ([]func(*rego.Rego), error) {
	f, err := flavor(ps)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("data.ladon.%s.allow", f)
	store, err := e.s.Storage(ctx, schema, []string{policyCollection(f), roleCollection(f)})
	if err != nil {
		return nil, err
	}

	var i input
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&i); err != nil {
		return nil, errors.WithStack(err)
	}

	return []func(*rego.Rego){
		rego.Query(query),
		rego.Store(store),
		rego.Input(&i),
	}, nil
}
