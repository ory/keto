// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"github.com/ory/x/httprouterx"
	"github.com/ory/x/httpx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/otelx"
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
		logrusx.Provider
		httpx.WriterProvider
		otelx.Provider
		x.NetworkIDProvider
		x.TransactorProvider
	}
	Handler struct {
		d handlerDeps
	}
)

const (
	ReadRouteBase  = "/relation-tuples"
	WriteRouteBase = "/admin/relation-tuples"
)

func NewHandler(d handlerDeps) *Handler {
	return &Handler{
		d: d,
	}
}

func (h *Handler) RegisterReadRoutes(r *httprouterx.RouterPublic) {
	r.GET(ReadRouteBase, h.getRelations)
}

func (h *Handler) RegisterWriteRoutes(r *httprouterx.RouterAdmin) {
	r.PUT(WriteRouteBase, h.createRelation)
	r.DELETE(WriteRouteBase, h.deleteRelations)
	r.PATCH(WriteRouteBase, h.patchRelationTuples)
}

func (h *Handler) RegisterReadGRPC(s *grpc.Server) {
	rts.RegisterReadServiceServer(s, h)
}

func (h *Handler) RegisterWriteGRPC(s *grpc.Server) {
	rts.RegisterWriteServiceServer(s, h)
}
