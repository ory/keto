// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"google.golang.org/grpc"

	"github.com/ory/keto/internal/driver/config"

	"github.com/ory/keto/internal/x"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type (
	handlerDeps interface {
		ManagerProvider
		MapperProvider
		config.Provider
		x.LoggerProvider
		x.WriterProvider
		x.TracingProvider
		x.NetworkIDProvider
		x.TransactorProvider
	}
	handler struct {
		d handlerDeps
	}
)

const (
	ReadRouteBase  = "/relation-tuples"
	WriteRouteBase = "/admin/relation-tuples"
)

func NewHandler(d handlerDeps) *handler {
	return &handler{
		d: d,
	}
}

func (h *handler) RegisterReadRoutes(r *x.ReadRouter) {
	r.GET(ReadRouteBase, h.getRelations)
}

func (h *handler) RegisterWriteRoutes(r *x.WriteRouter) {
	r.PUT(WriteRouteBase, h.createRelation)
	r.DELETE(WriteRouteBase, h.deleteRelations)
	r.PATCH(WriteRouteBase, h.patchRelationTuples)
}

func (h *handler) RegisterReadGRPC(s *grpc.Server) {
	rts.RegisterReadServiceServer(s, h)
}

func (h *handler) RegisterWriteGRPC(s *grpc.Server) {
	rts.RegisterWriteServiceServer(s, h)
}
