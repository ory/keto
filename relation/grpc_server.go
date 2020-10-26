package relation

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

func (s *Server) WriteRelation(ctx context.Context, r *models.RelationTuple) (*models.WriteRelationTupleResponse, error) {
	return &models.WriteRelationTupleResponse{}, s.d.RelationManager().WriteRelation(ctx, (&models.Relation{}).ImportFromGRPC(r))
}

func NewServer(d serverDependencies) *Server {
	return &Server{
		d: d,
	}
}

func (s *Server) ReadTuples(ctx context.Context, req *models.ReadRelationTuplesRequest) (*models.ReadRelationTuplesResponse, error) {
	queries := make([]*models.RelationQuery, len(req.Tuplesets))
	for i, tupleset := range req.Tuplesets {
		queries[i] = (&models.RelationQuery{}).ImportFromGRPC(tupleset)
	}

	normalRels, _ := s.d.RelationManager().GetRelations(ctx, queries, req.Page, req.PerPage)

	rpcRels := make([]*models.RelationTuple, len(req.Tuplesets))
	for i, tupleset := range normalRels {
		rpcRels[i] = (&models.RelationTuple{}).ImportFromNormal(tupleset)
	}

	return &models.ReadRelationTuplesResponse{Tuples: rpcRels}, nil
}
