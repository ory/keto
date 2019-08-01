package storage

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/ory/herodot"
	"github.com/ory/x/pagination"
)

type Handler struct {
	s Manager
	h herodot.Writer
}

func NewHandler(s Manager, h herodot.Writer) *Handler {
	return &Handler{
		s: s,
		h: h,
	}
}

type GetRequest struct {
	Collection string
	Key        string
	Value      interface{}
}

func (h *Handler) Get(factory func(context.Context, *http.Request, httprouter.Params) (*GetRequest, error)) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := r.Context()
		d, err := factory(ctx, r, ps)
		if err != nil {
			h.h.WriteError(w, r, err)
			return
		}

		if err := h.s.Get(ctx, d.Collection, d.Key, d.Value); err != nil {
			h.h.WriteError(w, r, err)
			return
		}

		h.h.Write(w, r, d.Value)
	}
}

type DeleteRequest struct {
	Collection string
	Key        string
}

func (h *Handler) Delete(factory func(context.Context, *http.Request, httprouter.Params) (*DeleteRequest, error)) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := r.Context()
		d, err := factory(ctx, r, ps)
		if err != nil {
			h.h.WriteError(w, r, err)
			return
		}

		if err := h.s.Delete(ctx, d.Collection, d.Key); err != nil {
			h.h.WriteError(w, r, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

type ListRequest struct {
	Collection string
	Value      interface{}
	FilterFunc func(*ListRequest, map[string][]string)
}

func (l *ListRequest) Filter(m map[string][]string) *ListRequest {
	if l.FilterFunc != nil {
		l.FilterFunc(l, m)
	}
	return l
}

func ListByQuery(l *ListRequest, m map[string][]string) {
	switch val := l.Value.(type) {
	case *Roles:
		var res Roles
		for _, role := range *val {
			filteredRole := role.withMembers(m["members"]).withIDs(m["id"])
			if filteredRole != nil {
				res = append(res, *filteredRole)
			}
		}
		l.Value = &res
	case *Policies:
		var res Policies
		for _, policy := range *val {
			filteredPolicy := policy.withSubjects(m["subjects"]).withResources(m["resources"]).withActions(m["actions"]).withIDs(m["id"])
			if filteredPolicy != nil {
				res = append(res, *filteredPolicy)
			}
		}
		l.Value = &res
	default:
		panic("storage:unable to cast list request to a known type!")
	}
}

func (h *Handler) List(factory func(context.Context, *http.Request, httprouter.Params) (*ListRequest, error)) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := r.Context()
		l, err := factory(ctx, r, ps)
		if err != nil {
			h.h.WriteError(w, r, err)
			return
		}
		limit, offset := pagination.Parse(r, 100, 0, 500)

		if err := h.s.List(ctx, l.Collection, l.Value, limit, offset); err != nil {
			h.h.WriteError(w, r, err)
			return
		}

		m := r.URL.Query()
		h.h.Write(w, r, l.Filter(m).Value)
	}
}

type UpsertRequest struct {
	Collection string
	Key        string
	Value      interface{}
}

func (h *Handler) Upsert(factory func(context.Context, *http.Request, httprouter.Params) (*UpsertRequest, error)) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := r.Context()
		u, err := factory(ctx, r, ps)
		if err != nil {
			h.h.WriteError(w, r, err)
			return
		}

		if err := h.s.Upsert(ctx, u.Collection, u.Key, u.Value); err != nil {
			h.h.WriteError(w, r, err)
			return
		}

		h.h.Write(w, r, u.Value)
	}
}
