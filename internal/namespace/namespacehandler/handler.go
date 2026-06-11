// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package namespacehandler

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/ory/herodot"
	"github.com/ory/x/httprouterx"
	"github.com/ory/x/httpx"
	"github.com/ory/x/logrusx"

	rts "github.com/ory/keto/gen/go/ory/keto/relation_tuples/v1alpha2"
	"github.com/ory/keto/gen/go/ory/keto/relation_tuples/v1alpha2/relationtuplesconnect"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/x"
)

type (
	handlerDeps interface {
		logrusx.Provider
		httpx.WriterProvider
		config.Provider
		x.HandlerOptionsProvider
	}
	Handler struct {
		relationtuplesconnect.UnimplementedNamespacesServiceHandler
		d handlerDeps
	}
)

const (
	RouteBase = "/namespaces"
)

func New(d handlerDeps) *Handler { return &Handler{d: d} }

func (h *Handler) RegisterReadRoutes(r *httprouterx.RouterPublic) {
	r.GET(RouteBase, h.getNamespaces)

	listNamespacesHandler := connect.NewUnaryHandler(
		relationtuplesconnect.NamespacesServiceListNamespacesProcedure,
		h.ListNamespaces,
		connect.WithSchema(rts.File_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto.
			Services().ByName("NamespacesService").
			Methods().ByName("ListNamespaces")),
		connect.WithHandlerOptions(h.d.HandlerOptions()...),
	)
	r.Handle(relationtuplesconnect.NamespacesServiceListNamespacesProcedure, listNamespacesHandler)
}

func (h *Handler) ProtoFiles() []protoreflect.FileDescriptor {
	return []protoreflect.FileDescriptor{rts.File_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto}
}

func (h *Handler) ListNamespaces(ctx context.Context, _ *connect.Request[rts.ListNamespacesRequest]) (*connect.Response[rts.ListNamespacesResponse], error) {
	m, err := h.d.Config(ctx).NamespaceManager()
	if err != nil {
		h.d.Logger().WithError(err).Errorf("could not get namespace manager")
		return nil, herodot.ErrInternalServerError()
	}
	namespaces, err := m.Namespaces(ctx)
	if err != nil {
		h.d.Logger().WithError(err).Errorf("could not get namespaces")
		return nil, herodot.ErrInternalServerError()
	}
	apiNamespaces := make([]*rts.Namespace, len(namespaces))
	for i, n := range namespaces {
		apiNamespaces[i] = &rts.Namespace{Name: n.Name}
	}
	return connect.NewResponse(&rts.ListNamespacesResponse{Namespaces: apiNamespaces}), nil
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
//
//	Extensions:
//	  x-ory-ratelimit-bucket: keto-admin-low
func (h *Handler) getNamespaces(w http.ResponseWriter, r *http.Request) {
	res, err := h.ListNamespaces(r.Context(), nil)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}
	h.d.Writer().Write(w, r, res.Msg)
}
