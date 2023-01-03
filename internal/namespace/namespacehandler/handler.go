// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package namespacehandler

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"google.golang.org/grpc"

	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/x"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type (
	handlerDeps interface {
		x.LoggerProvider
		x.WriterProvider
		config.Provider
	}
	handler struct {
		handlerDeps
	}
)

const (
	RouteBase = "/namespaces"
)

func New(d handlerDeps) *handler {
	return &handler{d}
}

func (h *handler) RegisterReadRoutes(r *x.ReadRouter) {
	r.GET(RouteBase, h.getNamespaces)
}

func (h *handler) RegisterReadGRPC(s *grpc.Server) {
	rts.RegisterNamespacesServiceServer(s, h)
}

func (h *handler) RegisterWriteRoutes(r *x.WriteRouter) {}
func (h *handler) RegisterWriteGRPC(s *grpc.Server)     {}

func (h *handler) ListNamespaces(ctx context.Context, _ *rts.ListNamespacesRequest) (*rts.ListNamespacesResponse, error) {
	m, err := h.Config(ctx).NamespaceManager()
	if err != nil {
		h.Logger().WithError(err).Errorf("could not get namespace manager")
		return nil, herodot.ErrInternalServerError
	}
	namespaces, err := m.Namespaces(ctx)
	if err != nil {
		h.Logger().WithError(err).Errorf("could not get namespaces")
		return nil, herodot.ErrInternalServerError
	}
	apiNamespaces := make([]*rts.Namespace, len(namespaces))
	for i, n := range namespaces {
		apiNamespaces[i] = &rts.Namespace{Name: n.Name}
	}
	return &rts.ListNamespacesResponse{Namespaces: apiNamespaces}, nil
}

// swagger:route GET /namespaces relationship listRelationshipNamespaces
//
// # Query namespaces
//
// Get all namespaces
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Responses:
//	  200: relationshipNamespaces
//	  default: errorGeneric
func (h *handler) getNamespaces(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := h.ListNamespaces(r.Context(), nil)
	if err != nil {
		h.Writer().WriteError(w, r, err)
		return
	}
	h.Writer().Write(w, r, res)
}
