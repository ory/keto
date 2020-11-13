package relationtuple

import (
	"context"

	"github.com/ory/keto/internal/x"
)

var _ RelationTupleServiceServer = &Server{}

type (
	serverDependencies interface {
		ManagerProvider
	}
	Server struct {
		UnimplementedRelationTupleServiceServer

		d serverDependencies
	}
)

func NewServer(d serverDependencies) *Server {
	return &Server{
		d: d,
	}
}

func (s *Server) WriteRelationTuple(ctx context.Context, r *WriteRelationTupleRequest) (*WriteRelationTupleResponse, error) {
	return &WriteRelationTupleResponse{}, s.d.RelationTupleManager().WriteRelationTuples(ctx, (&InternalRelationTuple{}).FromGRPC(r.Tuple))
}

func (s *Server) ReadRelationTuples(ctx context.Context, req *ReadRelationTuplesRequest) (*ReadRelationTuplesResponse, error) {
	query := (&RelationQuery{}).FromGRPC(req.Query)

	normalRels, _ := s.d.RelationTupleManager().GetRelationTuples(ctx, query, x.WithPage(int(req.Page)), x.WithPerPage(int(req.PerPage)))

	rpcRels := make([]*RelationTuple, len(normalRels))
	for i, tupleset := range normalRels {
		rpcRels[i] = (&RelationTuple{}).FromInternal(tupleset)
	}

	return &ReadRelationTuplesResponse{Tuples: rpcRels}, nil
}
