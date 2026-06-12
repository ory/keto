// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	rts "github.com/ory/keto/gen/go/ory/keto/relation_tuples/v1alpha2"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:4467", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := rts.NewWriteServiceClient(conn)

	// During migration, write every new tuple as both SubjectID and SubjectSet.
	// This keeps existing SubjectID checks working while the backfill runs.
	// Here we simulate an incoming request granting Charlie viewer access.
	_, err = client.TransactRelationTuples(context.Background(), &rts.TransactRelationTuplesRequest{
		RelationTupleDeltas: []*rts.RelationTupleDelta{
			{
				Action: rts.RelationTupleDelta_ACTION_INSERT,
				RelationTuple: &rts.RelationTuple{
					Namespace: "File", Object: "data.txt", Relation: "viewer",
					Subject: rts.NewSubjectID("user_charlie"),
				},
			},
			{
				Action: rts.RelationTupleDelta_ACTION_INSERT,
				RelationTuple: &rts.RelationTuple{
					Namespace: "File", Object: "data.txt", Relation: "viewer",
					Subject: rts.NewSubjectSet("User", "charlie", ""),
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully created tuples")
}
