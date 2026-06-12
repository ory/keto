// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"connectrpc.com/connect"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/ory/x/httprouterx"
	"github.com/ory/x/httpx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/otelx"

	rts "github.com/ory/keto/gen/go/ory/keto/relation_tuples/v1alpha2"
	"github.com/ory/keto/gen/go/ory/keto/relation_tuples/v1alpha2/relationtuplesconnect"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/x"
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
		x.HandlerOptionsProvider
	}
	ReadHandler struct {
		relationtuplesconnect.UnimplementedReadServiceHandler
		d handlerDeps
	}
	WriteHandler struct {
		relationtuplesconnect.UnimplementedWriteServiceHandler
		d handlerDeps
	}
)

const (
	ReadRouteBase  = "/relation-tuples"
	WriteRouteBase = "/admin/relation-tuples"
)

func NewReadHandler(d handlerDeps) *ReadHandler   { return &ReadHandler{d: d} }
func NewWriteHandler(d handlerDeps) *WriteHandler { return &WriteHandler{d: d} }

func (h *ReadHandler) RegisterReadRoutes(r *httprouterx.RouterPublic) {
	r.GET(ReadRouteBase, h.getRelations)

	listRelationTuplesHandler := connect.NewUnaryHandler(
		relationtuplesconnect.ReadServiceListRelationTuplesProcedure,
		h.ListRelationTuples,
		connect.WithSchema(rts.File_ory_keto_relation_tuples_v1alpha2_read_service_proto.
			Services().ByName("ReadService").
			Methods().ByName("ListRelationTuples")),
		connect.WithHandlerOptions(h.d.HandlerOptions()...),
	)
	r.Handle(relationtuplesconnect.ReadServiceListRelationTuplesProcedure, listRelationTuplesHandler)
}

func (h *WriteHandler) RegisterWriteRoutes(r *httprouterx.RouterAdmin) {
	r.PUT(WriteRouteBase, h.createRelation)
	r.DELETE(WriteRouteBase, h.deleteRelations)
	r.PATCH(WriteRouteBase, h.patchRelationTuples)

	writeServiceMethods := rts.File_ory_keto_relation_tuples_v1alpha2_write_service_proto.Services().ByName("WriteService").Methods()
	transactRelationTuplesHandler := connect.NewUnaryHandler(
		relationtuplesconnect.WriteServiceTransactRelationTuplesProcedure,
		h.TransactRelationTuples,
		connect.WithSchema(writeServiceMethods.ByName("TransactRelationTuples")),
		connect.WithHandlerOptions(h.d.HandlerOptions()...),
	)
	r.Handle(relationtuplesconnect.WriteServiceTransactRelationTuplesProcedure, transactRelationTuplesHandler)
	deleteRelationTuplesHandler := connect.NewUnaryHandler(
		relationtuplesconnect.WriteServiceDeleteRelationTuplesProcedure,
		h.DeleteRelationTuples,
		connect.WithSchema(writeServiceMethods.ByName("DeleteRelationTuples")),
		connect.WithHandlerOptions(h.d.HandlerOptions()...),
	)
	r.Handle(relationtuplesconnect.WriteServiceDeleteRelationTuplesProcedure, deleteRelationTuplesHandler)
}

func (h *ReadHandler) ProtoFiles() []protoreflect.FileDescriptor {
	return []protoreflect.FileDescriptor{
		rts.File_ory_keto_relation_tuples_v1alpha2_read_service_proto,
		rts.File_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto,
	}
}

func (h *WriteHandler) ProtoFiles() []protoreflect.FileDescriptor {
	return []protoreflect.FileDescriptor{
		rts.File_ory_keto_relation_tuples_v1alpha2_write_service_proto,
		rts.File_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto,
	}
}
