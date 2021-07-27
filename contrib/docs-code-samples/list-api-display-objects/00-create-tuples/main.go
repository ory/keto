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

	var tuples []*acl.RelationTuple
	// memes
	for _, user := range []string{"PM", "Vincent", "Julia"} {
		tuples = append(tuples, &acl.RelationTuple{
			Namespace: "chats",
			Object:    "memes",
			Relation:  "member",
			Subject:   acl.NewSubjectID(user),
		})
	}
	// cars
	for _, user := range []string{"PM", "Julia"} {
		tuples = append(tuples, &acl.RelationTuple{
			Namespace: "chats",
			Object:    "cars",
			Relation:  "member",
			Subject:   acl.NewSubjectID(user),
		})
	}
	// coffee-break
	for _, user := range []string{"PM", "Vincent", "Julia", "Patrik"} {
		tuples = append(tuples, &acl.RelationTuple{
			Namespace: "chats",
			Object:    "coffee-break",
			Relation:  "member",
			Subject:   acl.NewSubjectID(user),
		})
	}

	_, err = client.TransactRelationTuples(context.Background(), &acl.TransactRelationTuplesRequest{
		RelationTupleDeltas: acl.RelationTupleToDeltas(tuples, acl.RelationTupleDelta_INSERT),
	})
	if err != nil {
		panic("Encountered error: " + err.Error())
	}

	fmt.Println("Successfully created tuples")
}
