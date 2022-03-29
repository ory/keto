//go:build docscodesamples
// +build docscodesamples

package main

import (
	"context"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"google.golang.org/grpc"
)

func main() {
	wc, err := grpc.Dial("127.0.0.1:4467", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer wc.Close()

	wClient := rts.NewWriteServiceClient(wc)
	_, err = wClient.TransactRelationTuples(context.Background(), &rts.TransactRelationTuplesRequest{
		RelationTupleDeltas: []*rts.RelationTupleDelta{
			{
				RelationTuple: &rts.RelationTuple{
					Namespace: "messages",
					Object:    "02y_15_4w350m3",
					Relation:  "decypher",
					Subject:   rts.NewSubjectID("john"),
				},
				Action: rts.RelationTupleDelta_ACTION_DELETE,
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
