package decision

import (
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage"
	"github.com/pkg/errors"
	"context"
	"github.com/open-policy-agent/opa/ast"
	"github.com/ory/herodot"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

type Decider func(ctx context.Context, r *http.Request, ps httprouter.Params) (string, storage.Store, interface{}, error)

type Evaluator struct {
	compiler *ast.Compiler
	h        herodot.Writer
	deciders map[string]Decider
}

type Result struct {
	Allowed bool `json:"allowed"`
}

func (h *Evaluator) handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	d, ok := h.deciders[r.URL.Path]
	if !ok {
		h.h.WriteError(w, r, errors.WithStack(&herodot.ErrorNotFound))
		return
	}

	query, store, input, err := d(ctx, r, ps)
	if err != nil {
		h.h.WriteError(w, r, err)
		return
	}

	result, err := h.eval(ctx, query,store,input)
	if err != nil {
		h.h.WriteError(w, r, err)
		return
	}

	if err := json.NewEncoder(w).Encode(&Result{Allowed: result}); err != nil {
		h.h.WriteError(w, r, err)
		return
	}
}

func (h *Evaluator) eval(ctx context.Context, query string, store storage.Store, input interface{}) (bool, error) {
	r := rego.New(
		rego.Query(query),
		rego.Compiler(h.compiler),
		rego.Store(store),
		// rego.Tracer(tracer),
		rego.Input(input),
	)

	rs, err := r.Eval(ctx)
	if err != nil {
		return false, errors.WithStack(err)
	}

	if len(rs) != 1 || len(rs[0].Expressions) != 1 {
		return false, errors.Errorf("expected one evaluation result but got %d results instead", len(rs))
	}

	result, ok := rs[0].Expressions[0].Value.(bool)
	if !ok {
		return false, errors.Errorf("expected evaluation result to be of type bool but got %T instead", rs[0].Expressions[0].Value)
	}

	return result, nil
}
