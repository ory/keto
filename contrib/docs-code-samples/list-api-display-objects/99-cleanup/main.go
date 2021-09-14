//go:build docscodesamples
// +build docscodesamples

package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"
)

func main() {
	rc, err := grpc.Dial("127.0.0.1:4466", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer rc.Close()

	rClient := acl.NewReadServiceClient(rc)
	resp, err := rClient.ListRelationTuples(context.Background(), &acl.ListRelationTuplesRequest{
		Query: &acl.ListRelationTuplesRequest_Query{
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

	deltas := make([]*acl.RelationTupleDelta, len(resp.RelationTuples))
	for i, rt := range resp.RelationTuples {
		deltas[i] = &acl.RelationTupleDelta{
			Action:        acl.RelationTupleDelta_DELETE,
			RelationTuple: proto.Clone(rt).(*acl.RelationTuple),
		}
	}

	wClient := acl.NewWriteServiceClient(wc)
	_, err = wClient.TransactRelationTuples(context.Background(), &acl.TransactRelationTuplesRequest{
		RelationTupleDeltas: deltas,
	})
	if err != nil {
		panic(err)
	}
}
