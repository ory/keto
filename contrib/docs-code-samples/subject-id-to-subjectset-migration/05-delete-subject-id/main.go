// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

// subjectIdTuples are the SubjectID tuples collected and persisted during migration step.
// In production, load this list from the file or database written during migration so that
// you can target deletion of the tuples that were already migrated, and resume if interrupted.
var subjectIdTuples = []*rts.RelationTuple{
	{Namespace: "File", Object: "data.txt", Relation: "viewer", Subject: rts.NewSubjectID("user_alice")},
	{Namespace: "File", Object: "data.txt", Relation: "viewer", Subject: rts.NewSubjectID("user_bob")},
	{Namespace: "File", Object: "data.txt", Relation: "viewer", Subject: rts.NewSubjectID("user_charlie")},
	{Namespace: "File", Object: "data.txt", Relation: "viewer", Subject: rts.NewSubjectID("apikey_ci-bot")},
}

func main() {
	conn, err := grpc.NewClient("127.0.0.1:4467", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := rts.NewWriteServiceClient(conn)

	deltas := make([]*rts.RelationTupleDelta, len(subjectIdTuples))

	for i, t := range subjectIdTuples {
		deltas[i] = &rts.RelationTupleDelta{
			Action:        rts.RelationTupleDelta_ACTION_DELETE,
			RelationTuple: t,
		}
	}

	_, err = client.TransactRelationTuples(context.Background(), &rts.TransactRelationTuplesRequest{
		RelationTupleDeltas: deltas,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully deleted SubjectID tuples")
}
