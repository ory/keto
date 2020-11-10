package relationtuple

import (
	"context"

	"github.com/ory/keto/models"
)

var _ models.RelationTupleServiceServer = &Server{}

type (
	serverDependencies interface {
		ManagerProvider
	}
	Server struct {
		models.UnimplementedRelationTupleServiceServer

		d serverDependencies
	}
)

func NewServer(d serverDependencies) *Server {
	return &Server{
		d: d,
	}
}

func (s *Server) WriteRelationTuple(ctx context.Context, r *models.WriteRelationTupleRequest) (*models.WriteRelationTupleResponse, error) {
	return &models.WriteRelationTupleResponse{}, s.d.RelationTupleManager().WriteRelationTuples(ctx, (&models.InternalRelationTuple{}).FromGRPC(r.Tuple))
}

func (s *Server) ReadRelationTuples(ctx context.Context, req *models.ReadRelationTuplesRequest) (*models.ReadRelationTuplesResponse, error) {
	query := (&models.RelationQuery{}).FromGRPC(req.Query)

	normalRels, _ := s.d.RelationTupleManager().GetRelationTuples(ctx, query, WithPage(int(req.Page)), WithPerPage(int(req.PerPage)))

	rpcRels := make([]*models.RelationTuple, len(normalRels))
	for i, tupleset := range normalRels {
		rpcRels[i] = (&models.RelationTuple{}).FromInternal(tupleset)
	}

	return &models.ReadRelationTuplesResponse{Tuples: rpcRels}, nil
}
