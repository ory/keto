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

func (s *Server) mustEmbedUnimplementedRelationTupleServiceServer() {
	panic("implement me")
}

func (s *Server) WriteRelationTuple(ctx context.Context, r *models.WriteRelationTupleRequest) (*models.WriteRelationTupleResponse, error) {
	return &models.WriteRelationTupleResponse{}, s.d.RelationTupleManager().WriteRelationTuple(ctx, (&models.InternalRelationTuple{}).FromGRPC(r.Tuple))
}

func NewServer(d serverDependencies) *Server {
	return &Server{
		d: d,
	}
}

func (s *Server) ReadRelationTuples(ctx context.Context, req *models.ReadRelationTuplesRequest) (*models.ReadRelationTuplesResponse, error) {
	queries := make([]*models.RelationQuery, len(req.TupleSets))
	for i, tupleset := range req.TupleSets {
		queries[i] = (&models.RelationQuery{}).FromGRPC(tupleset)
	}

	normalRels, _ := s.d.RelationTupleManager().GetRelationTuples(ctx, queries, req.Page, req.PerPage)

	rpcRels := make([]*models.RelationTuple, len(normalRels))
	for i, tupleset := range normalRels {
		rpcRels[i] = (&models.RelationTuple{}).FromInternal(tupleset)
	}

	return &models.ReadRelationTuplesResponse{Tuples: rpcRels}, nil
}
