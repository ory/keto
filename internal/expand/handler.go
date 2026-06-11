// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package expand

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/ory/x/httprouterx"
	"github.com/ory/x/httpx"
	"github.com/ory/x/logrusx"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/ory/keto/gen/go/ory/keto/relation_tuples/v1alpha2/relationtuplesconnect"

	"github.com/ory/keto/ketoapi"

	"github.com/ory/herodot"

	rts "github.com/ory/keto/gen/go/ory/keto/relation_tuples/v1alpha2"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

type (
	handlerDependencies interface {
		EngineProvider
		relationtuple.ManagerProvider
		relationtuple.MapperProvider
		logrusx.Provider
		httpx.WriterProvider
		x.HandlerOptionsProvider
	}
	Handler struct {
		relationtuplesconnect.UnimplementedExpandServiceHandler
		d handlerDependencies
	}
)

var (
	_ relationtuplesconnect.ExpandServiceHandler = (*Handler)(nil)
	_ *expandPermissions                         = nil
)

const RouteBase = "/relation-tuples/expand"

func NewHandler(d handlerDependencies) *Handler { return &Handler{d: d} }

func (h *Handler) RegisterReadRoutes(r *httprouterx.RouterPublic) {
	r.GET(RouteBase, h.getExpand)

	expandHandler := connect.NewUnaryHandler(
		relationtuplesconnect.ExpandServiceExpandProcedure,
		h.Expand,
		connect.WithSchema(rts.File_ory_keto_relation_tuples_v1alpha2_expand_service_proto.
			Services().ByName("ExpandService").
			Methods().ByName("Expand")),
		connect.WithHandlerOptions(h.d.HandlerOptions()...),
	)
	r.Handle(relationtuplesconnect.ExpandServiceExpandProcedure, expandHandler)
}

func (h *Handler) ProtoFiles() []protoreflect.FileDescriptor {
	return []protoreflect.FileDescriptor{
		rts.File_ory_keto_relation_tuples_v1alpha2_expand_service_proto,
		rts.File_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto,
	}
}

// Expand Permissions Request Parameters
//
// swagger:parameters expandPermissions
type expandPermissions struct {
	// in:query
	MaxDepth int `json:"max-depth"`
	// in:query
	ketoapi.SubjectSet
}

// swagger:route GET /relation-tuples/expand permission expandPermissions
//
// # Expand a Relationship into permissions.
//
// Use this endpoint to expand a relationship tuple into permissions.
//
//	Consumes:
//	-  application/x-www-form-urlencoded
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Responses:
//	  200: expandedPermissionTree
//	  400: errorGeneric
//	  404: errorGeneric
//	  default: errorGeneric
//
//	Extensions:
//	  x-ory-ratelimit-bucket: keto-admin-medium
func (h *Handler) getExpand(w http.ResponseWriter, r *http.Request) {
	maxDepth, err := x.GetMaxDepthFromQuery(r.URL.Query())
	if err != nil {
		h.d.Writer().WriteError(w, r, herodot.ErrBadRequest().WithError(err.Error()))
		return
	}

	subSet := (&ketoapi.SubjectSet{}).FromURLQuery(r.URL.Query())
	internal, err := h.d.ReadOnlyMapper().FromSubjectSet(r.Context(), subSet)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	res, err := h.d.ExpandEngine().BuildTree(r.Context(), internal, maxDepth)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}
	if res == nil {
		h.d.Writer().Write(w, r, herodot.ErrNotFound().WithError("no relation tuple found"))
		return
	}

	tree, err := h.d.ReadOnlyMapper().ToTree(r.Context(), res)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	h.d.Writer().Write(w, r, tree)
}

func (h *Handler) Expand(ctx context.Context, req *connect.Request[rts.ExpandRequest]) (*connect.Response[rts.ExpandResponse], error) {
	var subSet *ketoapi.SubjectSet

	switch sub := req.Msg.Subject.Ref.(type) {
	case *rts.Subject_Id:
		return connect.NewResponse(&rts.ExpandResponse{
			Tree: &rts.SubjectTree{
				NodeType: rts.NodeType_NODE_TYPE_LEAF,
				Subject:  rts.NewSubjectID(sub.Id),
			},
		}), nil
	case *rts.Subject_Set:
		subSet = &ketoapi.SubjectSet{
			Namespace: sub.Set.Namespace,
			Object:    sub.Set.Object,
			Relation:  sub.Set.Relation,
		}
	}

	internal, err := h.d.ReadOnlyMapper().FromSubjectSet(ctx, subSet)
	if err != nil {
		return nil, err
	}
	res, err := h.d.ExpandEngine().BuildTree(ctx, internal, int(req.Msg.MaxDepth))
	if err != nil {
		return nil, err
	}
	if res == nil {
		return connect.NewResponse(&rts.ExpandResponse{}), nil
	}

	tree, err := h.d.ReadOnlyMapper().ToTree(ctx, res)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&rts.ExpandResponse{Tree: tree.ToProto()}), nil
}
