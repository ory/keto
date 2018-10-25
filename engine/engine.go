package engine

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/topdown"
	"github.com/pkg/errors"

	"github.com/ory/herodot"
)

// swagger:ignore
type Engine struct {
	compiler *ast.Compiler
	h        herodot.Writer
}

func NewEngine(
	compiler *ast.Compiler,
	h herodot.Writer,
) *Engine {
	return &Engine{
		compiler: compiler,
		h:        h,
	}
}

// swagger:ignore
type evaluator func(ctx context.Context, r *http.Request, ps httprouter.Params) ([]func(*rego.Rego), error)

func (h *Engine) Evaluate(e evaluator) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := r.Context()

		rs, err := e(ctx, r, ps)
		if err != nil {
			h.h.WriteError(w, r, err)
			return
		}

		allowed, err := h.eval(ctx, rs)
		if err != nil {
			h.h.WriteError(w, r, err)
			return
		}

		if err := json.NewEncoder(w).Encode(&AuthorizationResult{Allowed: allowed}); err != nil {
			h.h.WriteError(w, r, errors.WithStack(err))
			return
		}
	}
}

func (h *Engine) eval(ctx context.Context, options []func(*rego.Rego)) (bool, error) {
	tracer := topdown.NewBufferTracer()
	r := rego.New(
		append(
			options,
			rego.Compiler(h.compiler),
			rego.Tracer(tracer),
		)...,
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
