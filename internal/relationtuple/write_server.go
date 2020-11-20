package relationtuple

import (
	"context"

	acl "github.com/ory/keto/api/keto/acl/v1alpha1"
)

var _ acl.WriteServiceServer = &GRPCServer{}

func tuplesWithAction(deltas []*acl.RelationTupleWriteDelta, action acl.RelationTupleWriteDelta_Action) (filtered []*InternalRelationTuple) {
	for _, d := range deltas {
		if d.Action == action {
			filtered = append(
				filtered,
				(&InternalRelationTuple{}).FromGRPC(d.RelationTuple),
			)
		}
	}
	return
}

func (s *GRPCServer) WriteRelationTuples(ctx context.Context, req *acl.WriteRelationTuplesRequest) (*acl.WriteRelationTuplesResponse, error) {
	insertTuples := tuplesWithAction(req.RelationTupleDeltas, acl.RelationTupleWriteDelta_INSERT)

	err := s.d.RelationTupleManager().WriteRelationTuples(ctx, insertTuples...)
	if err != nil {
		return nil, err
	}

	snaptokens := make([]string, len(insertTuples))
	for i := range insertTuples {
		snaptokens[i] = "not yet implemented"
	}
	return &acl.WriteRelationTuplesResponse{
		Snaptokens: snaptokens,
	}, nil
}
