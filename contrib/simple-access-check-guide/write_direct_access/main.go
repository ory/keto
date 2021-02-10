package main

import (
	"context"
	"fmt"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4467")
	if err != nil {
		panic(err.Error())
	}

	client := acl.NewWriteServiceClient(conn)

	resp, err := client.TransactRelationTuples(context.Background(), &acl.TransactRelationTuplesRequest{
		RelationTupleDeltas: []*acl.RelationTupleDelta{
			{
				Action: acl.RelationTupleDelta_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "messages",
					Object:    "02y_15_4w350m3",
					Relation:  "decypher",
					Subject: &acl.Subject{Ref: &acl.Subject_Id{
						Id: "john",
					}},
				},
			},
		},
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", resp)
}
