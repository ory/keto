package relationtuple

import (
	"context"

	acl "github.com/ory/keto/api/keto/acl/v1alpha1"
	"github.com/ory/keto/internal/x"
)

var _ acl.ReadServiceServer = &grpcHandler{}

func (s *grpcHandler) ListRelationTuples(ctx context.Context, req *acl.ListRelationTuplesRequest) (*acl.ListRelationTuplesResponse, error) {
	rels, nextPage, err := s.d.RelationTupleManager().GetRelationTuples(ctx,
		&RelationQuery{
			Namespace: req.Query.Namespace,
			Object:    req.Query.Object,
			Relation:  req.Query.Relation,
			Subject:   SubjectFromProto(req.Query.Subject),
		},
		x.WithSize(uint(req.PageSize)),
		x.WithToken(req.PageToken),
	)
	if err != nil {
		return nil, err
	}

	resp := &acl.ListRelationTuplesResponse{
		RelationTuples: make([]*acl.RelationTuple, len(rels)),
		NextPageToken:  nextPage,
	}
	for i, r := range rels {
		resp.RelationTuples[i] = r.ToProto()
	}

	return resp, nil
}
