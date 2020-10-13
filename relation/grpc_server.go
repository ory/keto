package relation

import (
	"context"
	"github.com/ory/keto/relation/read"
)

var _ read.RelationReaderServer = &Server{}

type (
	serverDependencies interface {
		ManagerProvider
	}
	Server struct {
		read.UnimplementedRelationReaderServer

		d serverDependencies
	}
)

func NewServer(d serverDependencies) *Server {
	return &Server{
		d: d,
	}
}

func (_ *Server) relationsHelper(ctx context.Context, queryID string, page, perPage int32, getterFunc func(context.Context, string, int32, int32) ([]Relation, error)) (*read.RelationsResponse, error) {
	rels, err := getterFunc(ctx, queryID, page, perPage)
	if err != nil {
		return nil, err
	}

	rpcRels := make([]*read.Relation, len(rels))
	for i, r := range rels {
		rpcRels[i] = &read.Relation{
			ObjectID:     r.ObjectID,
			RelationName: r.Name,
			UserID:       r.UserID,
		}
	}
	return &read.RelationsResponse{
		Relations: rpcRels,
	}, nil
}

func (s *Server) RelationsByObject(ctx context.Context, req *read.RelationsRequest) (*read.RelationsResponse, error) {
	return s.relationsHelper(ctx, req.Id, req.Page, req.PerPage, s.d.RelationManager().GetRelationsByObject)
}

func (s *Server) RelationsByUser(ctx context.Context, req *read.RelationsRequest) (*read.RelationsResponse, error) {
	return s.relationsHelper(ctx, req.Id, req.Page, req.PerPage, s.d.RelationManager().GetRelationsByUser)
}
