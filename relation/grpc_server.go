package relation

import (
	"context"

	"github.com/ory/keto/models"
)

var _ models.GRPCRelationReaderServer = &Server{}
var _ models.GRPCRelationWriterServer = &Server{}

type (
	serverDependencies interface {
		ManagerProvider
	}
	Server struct {
		models.UnimplementedGRPCRelationReaderServer
		models.UnimplementedGRPCRelationWriterServer

		d serverDependencies
	}
)

func (s *Server) WriteRelation(ctx context.Context, r *models.GRPCRelation) (*models.GRPCRelationsWriteResponse, error) {
	return &models.GRPCRelationsWriteResponse{}, s.d.RelationManager().WriteRelation(ctx, (*models.Relation)(r))
}

func NewServer(d serverDependencies) *Server {
	return &Server{
		d: d,
	}
}

func (_ *Server) relationsHelper(ctx context.Context, queryID string, page, perPage int32, getterFunc func(context.Context, string, int32, int32) ([]*models.Relation, error)) (*models.GRPCRelationsReadResponse, error) {
	rels, err := getterFunc(ctx, queryID, page, perPage)
	if err != nil {
		return nil, err
	}

	rpcRels := make([]*models.GRPCRelation, len(rels))
	for i := range rels {
		rpcRels[i] = (&models.GRPCRelation{}).ImportFromNormal(rels[i])
	}
	return &models.GRPCRelationsReadResponse{
		Relations: rpcRels,
	}, nil
}

func (s *Server) RelationsByObject(ctx context.Context, req *models.GRPCRelationsReadRequest) (*models.GRPCRelationsReadResponse, error) {
	return s.relationsHelper(ctx, req.Id, req.Page, req.PerPage, s.d.RelationManager().GetRelationsByObject)
}

func (s *Server) RelationsByUser(ctx context.Context, req *models.GRPCRelationsReadRequest) (*models.GRPCRelationsReadResponse, error) {
	return s.relationsHelper(ctx, req.Id, req.Page, req.PerPage, s.d.RelationManager().GetRelationsByUser)
}
