// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:4467", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Encountered error: " + err.Error())
	}

	client := rts.NewWriteServiceClient(conn)

	var tuples []*rts.RelationTuple
	// memes
	for _, user := range []string{"PM", "Vincent", "Julia"} {
		tuples = append(tuples, &rts.RelationTuple{
			Namespace: "chats",
			Object:    "memes",
			Relation:  "member",
			Subject:   rts.NewSubjectID(user),
		})
	}
	// cars
	for _, user := range []string{"PM", "Julia"} {
		tuples = append(tuples, &rts.RelationTuple{
			Namespace: "chats",
			Object:    "cars",
			Relation:  "member",
			Subject:   rts.NewSubjectID(user),
		})
	}
	// coffee-break
	for _, user := range []string{"PM", "Vincent", "Julia", "Patrik"} {
		tuples = append(tuples, &rts.RelationTuple{
			Namespace: "chats",
			Object:    "coffee-break",
			Relation:  "member",
			Subject:   rts.NewSubjectID(user),
		})
	}

	_, err = client.TransactRelationTuples(context.Background(), &rts.TransactRelationTuplesRequest{
		RelationTupleDeltas: rts.RelationTupleToDeltas(tuples, rts.RelationTupleDelta_ACTION_INSERT),
	})
	if err != nil {
		panic("Encountered error: " + err.Error())
	}

	fmt.Println("Successfully created tuples")
}
