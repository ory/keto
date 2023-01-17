// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package expand

import (
	"context"
	"net/http"

	"github.com/ory/keto/ketoapi"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"google.golang.org/grpc"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type (
	handlerDependencies interface {
		EngineProvider
		relationtuple.ManagerProvider
		relationtuple.MapperProvider
		x.LoggerProvider
		x.WriterProvider
	}
	handler struct {
		d handlerDependencies
	}
)

var (
	_ rts.ExpandServiceServer = (*handler)(nil)
	_ *expandPermissions      = nil
)

const RouteBase = "/relation-tuples/expand"

func NewHandler(d handlerDependencies) *handler {
	return &handler{d: d}
}

func (h *handler) RegisterReadRoutes(r *x.ReadRouter) {
	r.GET(RouteBase, h.getExpand)
}

func (h *handler) RegisterReadGRPC(s *grpc.Server) {
	rts.RegisterExpandServiceServer(s, h)
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
func (h *handler) getExpand(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	maxDepth, err := x.GetMaxDepthFromQuery(r.URL.Query())
	if err != nil {
		h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithError(err.Error()))
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
		h.d.Writer().Write(w, r, herodot.ErrNotFound.WithError("no relation tuple found"))
		return
	}

	tree, err := h.d.ReadOnlyMapper().ToTree(r.Context(), res)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	h.d.Writer().Write(w, r, tree)
}

func (h *handler) Expand(ctx context.Context, req *rts.ExpandRequest) (*rts.ExpandResponse, error) {
	var subSet *ketoapi.SubjectSet

	switch sub := req.Subject.Ref.(type) {
	case *rts.Subject_Id:
		return &rts.ExpandResponse{
			Tree: &rts.SubjectTree{
				NodeType: rts.NodeType_NODE_TYPE_LEAF,
				Subject:  rts.NewSubjectID(sub.Id),
			},
		}, nil
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
	res, err := h.d.ExpandEngine().BuildTree(ctx, internal, int(req.MaxDepth))
	if err != nil {
		return nil, err
	}
	if res == nil {
		return &rts.ExpandResponse{}, nil
	}

	tree, err := h.d.ReadOnlyMapper().ToTree(ctx, res)
	if err != nil {
		return nil, err
	}

	return &rts.ExpandResponse{Tree: tree.ToProto()}, nil
}
