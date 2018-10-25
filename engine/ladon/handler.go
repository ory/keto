package ladon

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/open-policy-agent/opa/rego"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"

	"github.com/ory/go-convenience/stringslice"
	"github.com/ory/herodot"
	"github.com/ory/keto/engine"
	kstorage "github.com/ory/keto/storage"
)

// swagger:ignore
type Engine struct {
	sh     *kstorage.Handler
	engine *engine.Engine
	s      kstorage.Manager
	h      herodot.Writer
}

var EnabledFlavors = []string{"exact", "regex"}

const (
	BasePath = "/engines/acp/ory/:flavor"
	schema   = `{
	"store": {
		"ory": {
			"regex": {
				"policies": [],
				"roles": []
			},
			"exact": {
				"policies": [],
				"roles": []
			}
		}
	}
}`
)

func RoutesToObserve() []string {
	var r []string

	for _, f := range []string{"exact", "regex"} {
		for _, p := range []string{"policies", "roles", "allowed"} {
			r = append(r,
				fmt.Sprintf(strings.Replace(BasePath, ":flavor", "%s", 1)+"/%s", f, p),
			)
		}
	}

	return r
}

func policyCollection(f string) string {
	return fmt.Sprintf("/store/ory/%s/policies", f)
}

func roleCollection(f string) string {
	return fmt.Sprintf("/store/ory/%s/roles", f)
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
	// swagger:route GET /engines/ory/{flavor}/allowed engines doOryAccessControlPoliciesAllow
	//
	// Check if a request is allowed
	//
	// Use this endpoint to check if a request is allowed or not.
	//
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: authorizationResult
	//       500: genericError
	r.POST(BasePath+"/allowed", e.engine.Evaluate(e.eval))

	// swagger:route PUT /engines/ory/{flavor}/policies engines upsertOryAccessControlPolicy
	//
	// Upsert an ORY Access Control Policy
	//
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: policy
	//       500: genericError
	r.PUT(BasePath+"/policies", e.sh.Upsert(e.policiesCreate))

	// swagger:route GET /engines/ory/{flavor}/policies engines listOryAccessControlPolicies
	//
	// List ORY Access Control Policies
	//
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: oryAccessControlPolicies
	//       500: genericError
	r.GET(BasePath+"/policies", e.sh.List(e.policiesList))

	// swagger:route GET /engines/ory/{flavor}/policies/{id} engines getOryAccessControlPolicy
	//
	// Get an ORY Access Control Policy
	//
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: oryAccessControlPolicy
	//       404: genericError
	//       500: genericError
	r.GET(BasePath+"/policies/:id", e.sh.Get(e.policiesGet))

	// swagger:route DELETE /engines/ory/{flavor}/policies/{id} engines deleteOryAccessControlPolicy
	//
	// Delete an ORY Access Control Policy
	//
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       201: emptyResponse
	//       500: genericError
	r.DELETE(BasePath+"/policies/:id", e.sh.Delete(e.policiesDelete))

	// swagger:route GET /engines/ory/{flavor}/roles engines listOryAccessControlPolicyRoles
	//
	// List ORY Access Control Policy Roles
	//
	// Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID
	// as subject in the OACP.
	//
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: oryAccessControlPolicyRoles
	//       500: genericError
	r.GET(BasePath+"/roles", e.sh.List(e.rolesList))

	// swagger:route GET /engines/ory/{flavor}/roles/{id} engines getOryAccessControlPolicyRole
	//
	// Get an ORY Access Control Policy Role
	//
	// Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID
	// as subject in the OACP.
	//
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: oryAccessControlPolicyRole
	//       404: genericError
	//       500: genericError
	r.GET(BasePath+"/roles/:id", e.sh.Get(e.rolesGet))

	// swagger:route PUT /engines/ory/{flavor}/roles engines upsertOryAccessControlPolicyRole
	//
	// Upsert an ORY Access Control Policy Role
	//
	// Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID
	// as subject in the OACP.
	//
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: oryAccessControlPolicyRole
	//       500: genericError
	r.PUT(BasePath+"/roles", e.sh.Upsert(e.rolesUpsert))

	// swagger:route DELETE /engines/ory/{flavor}/roles/{id} engines deleteOryAccessControlPolicyRole
	//
	// Delete an ORY Access Control Policy Role
	//
	// Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID
	// as subject in the OACP.
	//
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       201: emptyResponse
	//       500: genericError
	r.DELETE(BasePath+"/roles/:id", e.sh.Delete(e.rolesDelete))

	// swagger:route PUT /engines/ory/{flavor}/roles/{id}/members engines addOryAccessControlPolicyRoleMembers
	//
	// Add a member to an ORY Access Control Policy Role
	//
	// Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID
	// as subject in the OACP.
	//
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: oryAccessControlPolicyRole
	//       500: genericError
	r.PUT(BasePath+"/roles/:id/members", e.sh.Upsert(e.rolesMembersAdd))

	// swagger:route DELETE /engines/ory/{flavor}/roles/{id}/members engines removeOryAccessControlPolicyRoleMembers
	//
	// Remove a member from an ORY Access Control Policy Role
	//
	// Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID
	// as subject in the OACP.
	//
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       201: emptyResponse
	//       500: genericError
	r.DELETE(BasePath+"/roles/:id/members/:member", e.sh.Upsert(e.rolesMembersRemove))
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

func (e *Engine) rolesUpsert(ctx context.Context, r *http.Request, ps httprouter.Params) (*kstorage.UpsertRequest, error) {
	var p Role
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		return nil, errors.WithStack(err)
	}

	if p.ID == "" {
		p.ID = uuid.New()
	}

	f, err := flavor(ps)
	if err != nil {
		return nil, err
	}

	return &kstorage.UpsertRequest{
		Collection: roleCollection(f),
		Key:        p.ID,
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
		ro.Members = stringslice.Unique(append(ro.Members, i.Members...))
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

	p, err := validatePolicy(p)
	if err != nil {
		return nil, err
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

// swagger:model oryAccessControlPolicyAllowedInput
type Input struct {
	// Resource is the resource that access is requested to.
	Resource string `json:"resource"`

	// Action is the action that is requested on the resource.
	Action string `json:"action"`

	// Subject is the subject that is requesting access.
	Subject string `json:"subject"`

	// Context is the request's environmental context.
	Context Context `json:"context"`
}

func flavor(ps httprouter.Params) (string, error) {
	t := ps.ByName("flavor")
	if !stringslice.Has(EnabledFlavors, t) {
		return "", errors.WithStack(&herodot.ErrorNotFound)
	}

	return t, nil
}

func (e *Engine) eval(ctx context.Context, r *http.Request, ps httprouter.Params) ([]func(*rego.Rego), error) {
	f, err := flavor(ps)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("data.ory.%s.allow", f)
	store, err := e.s.Storage(ctx, schema, []string{policyCollection(f), roleCollection(f)})
	if err != nil {
		return nil, err
	}

	var i Input
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
