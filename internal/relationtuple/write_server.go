package relationtuple

import (
	"context"

	acl "github.com/ory/keto/api/keto/acl/v1alpha1"
)

var _ acl.WriteServiceServer = &grpcHandler{}

func tuplesWithAction(deltas []*acl.RelationTupleDelta, action acl.RelationTupleDelta_Action) (filtered []*InternalRelationTuple) {
	for _, d := range deltas {
		if d.Action == action {
			filtered = append(
				filtered,
				(&InternalRelationTuple{}).FromDataProvider(d.RelationTuple),
			)
		}
	}
	return
}

func (s *grpcHandler) TransactRelationTuples(ctx context.Context, req *acl.TransactRelationTuplesRequest) (*acl.TransactRelationTuplesResponse, error) {
	insertTuples := tuplesWithAction(req.RelationTupleDeltas, acl.RelationTupleDelta_INSERT)

	err := s.d.RelationTupleManager().WriteRelationTuples(ctx, insertTuples...)
	if err != nil {
		return nil, err
	}

	snaptokens := make([]string, len(insertTuples))
	for i := range insertTuples {
		snaptokens[i] = "not yet implemented"
	}
	return &acl.TransactRelationTuplesResponse{
		Snaptokens: snaptokens,
	}, nil
}
