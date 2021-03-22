package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:4467", grpc.WithInsecure())
	if err != nil {
		panic("Encountered error: " + err.Error())
	}

	client := acl.NewWriteServiceClient(conn)

	var tupleDeltas []*acl.RelationTupleDelta
	// memes
	for _, user := range []string{"PM", "Vincent", "Julia"} {
		tupleDeltas = append(tupleDeltas, &acl.RelationTupleDelta{
			Action: acl.RelationTupleDelta_INSERT,
			RelationTuple: &acl.RelationTuple{
				Namespace: "chats",
				Object:    "memes",
				Relation:  "member",
				Subject: &acl.Subject{Ref: &acl.Subject_Id{
					Id: user,
				}},
			},
		})
	}
	// cars
	for _, user := range []string{"PM", "Julia"} {
		tupleDeltas = append(tupleDeltas, &acl.RelationTupleDelta{
			Action: acl.RelationTupleDelta_INSERT,
			RelationTuple: &acl.RelationTuple{
				Namespace: "chats",
				Object:    "cars",
				Relation:  "member",
				Subject: &acl.Subject{Ref: &acl.Subject_Id{
					Id: user,
				}},
			},
		})
	}
	// coffee-break
	for _, user := range []string{"PM", "Vincent", "Julia", "Patrik"} {
		tupleDeltas = append(tupleDeltas, &acl.RelationTupleDelta{
			Action: acl.RelationTupleDelta_INSERT,
			RelationTuple: &acl.RelationTuple{
				Namespace: "chats",
				Object:    "coffee-break",
				Relation:  "member",
				Subject: &acl.Subject{Ref: &acl.Subject_Id{
					Id: user,
				}},
			},
		})
	}

	_, err = client.TransactRelationTuples(context.Background(), &acl.TransactRelationTuplesRequest{
		RelationTupleDeltas: tupleDeltas,
	})
	if err != nil {
		panic("Encountered error: " + err.Error())
	}

	fmt.Println("Successfully created tuples")
}
