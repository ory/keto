//go:build docscodesamples
// +build docscodesamples

package main

import (
	"context"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func main() {
	rc, err := grpc.Dial("127.0.0.1:4466", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer rc.Close()

	rClient := rts.NewReadServiceClient(rc)
	resp, err := rClient.ListRelationTuples(context.Background(), &rts.ListRelationTuplesRequest{
		Query: &rts.ListRelationTuplesRequest_Query{
			Namespace: "chats",
		},
	})
	if err != nil {
		panic(err)
	}

	wc, err := grpc.Dial("127.0.0.1:4467", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer wc.Close()

	deltas := make([]*rts.RelationTupleDelta, len(resp.RelationTuples))
	for i, rt := range resp.RelationTuples {
		deltas[i] = &rts.RelationTupleDelta{
			Action:        rts.RelationTupleDelta_ACTION_DELETE,
			RelationTuple: proto.Clone(rt).(*rts.RelationTuple),
		}
	}

	wClient := rts.NewWriteServiceClient(wc)
	_, err = wClient.TransactRelationTuples(context.Background(), &rts.TransactRelationTuplesRequest{
		RelationTupleDeltas: deltas,
	})
	if err != nil {
		panic(err)
	}
}
