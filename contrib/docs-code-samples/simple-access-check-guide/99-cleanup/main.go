// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	wc, err := grpc.Dial("127.0.0.1:4467", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
