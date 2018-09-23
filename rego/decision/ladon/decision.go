package ladon

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"encoding/json"
	"github.com/open-policy-agent/opa/storage"
	"context"
	kstorage "github.com/ory/keto/rego/storage"
)

type Decider struct {
	store kstorage.Manager
}

type input struct {
	// Resource is the resource that access is requested to.
	Resource string `json:"resource"`

	// Action is the action that is requested on the resource.
	Action string `json:"action"`

	// Subejct is the subject that is requesting access.
	Subject string `json:"subject"`

	// Context is the request's environmental context.
	Context map[string]interface{} `json:"context"`
}

func (h *Decider) Exact(ctx context.Context, r *http.Request, ps httprouter.Params) (string, storage.Store, interface{}, error) {
	return h.eval(ctx, "exact", r, ps)
}

func (h *Decider) Regex(ctx context.Context, r *http.Request, ps httprouter.Params) (string, storage.Store, interface{}, error) {
	return h.eval(ctx, "regex", r, ps)
}

func (h *Decider) eval(ctx context.Context, t string, r *http.Request, ps httprouter.Params) (string, storage.Store, interface{}, error) {
	query := "data.ladon." + t + ".allow"
	store, err := h.store.Storage(ctx, "data.ladon."+t+".policies", "data.ladon."+t+".roles")
	if err != nil {
		return "", nil, nil, err
	}

	var i input
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&i); err != nil {
		return "", nil, nil, errors.WithStack(err)
	}

	return query, store, &i, err
}
