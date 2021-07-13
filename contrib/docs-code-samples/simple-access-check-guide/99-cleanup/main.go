// +build docscodesamples

package main

import (
	"context"

	"google.golang.org/grpc"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"
)

func main() {
	wc, err := grpc.Dial("127.0.0.1:4467", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer wc.Close()

	wClient := acl.NewWriteServiceClient(wc)
	_, err = wClient.TransactRelationTuples(context.Background(), &acl.TransactRelationTuplesRequest{
		RelationTupleDeltas: []*acl.RelationTupleDelta{
			{
				RelationTuple: &acl.RelationTuple{
					Namespace: "messages",
					Object:    "02y_15_4w350m3",
					Relation:  "decypher",
					Subject:   acl.NewSubjectID("john"),
				},
				Action: acl.RelationTupleDelta_DELETE,
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
