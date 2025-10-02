// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package expand

import (
	"context"

	"google.golang.org/grpc"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoapi"
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
)

const RouteBase = "/relation-tuples/expand"

func NewHandler(d handlerDependencies) *handler {
	return &handler{d: d}
}

func (h *handler) RegisterReadGRPC(s *grpc.Server) {
	rts.RegisterExpandServiceServer(s, h)
}

func (h *handler) Expand(ctx context.Context, req *rts.ExpandRequest) (*rts.ExpandResponse, error) {
	var subSet *ketoapi.SubjectSet

	if req.Subject == nil {
		subSet = &ketoapi.SubjectSet{
			Namespace: req.Namespace,
			Object:    req.Object,
			Relation:  req.Relation,
		}
	} else {
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
