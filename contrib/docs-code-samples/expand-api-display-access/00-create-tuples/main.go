// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:4467", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Encountered error: " + err.Error())
	}

	client := rts.NewWriteServiceClient(conn)

	//directories:/photos#owner@maureen
	//files:/photos/beach.jpg#owner@maureen
	//files:/photos/mountains.jpg#owner@laura
	//directories:/photos#access@laura
	//directories:/photos#access@(directories:/photos#owner)
	//files:/photos/beach.jpg#access@(files:/photos/beach.jpg#owner)
	//files:/photos/beach.jpg#access@(directories:/photos#access)
	//files:/photos/mountains.jpg#access@(files:/photos/mountains.jpg#owner)
	//files:/photos/mountains.jpg#access@(directories:/photos#access)

	tuples := []*rts.RelationTuple{
		// ownership
		{
			Namespace: "directories",
			Object:    "/photos",
			Relation:  "owner",
			Subject:   rts.NewSubjectID("maureen"),
		},
		{
			Namespace: "files",
			Object:    "/photos/beach.jpg",
			Relation:  "owner",
			Subject:   rts.NewSubjectID("maureen"),
		},
		{
			Namespace: "files",
			Object:    "/photos/mountains.jpg",
			Relation:  "owner",
			Subject:   rts.NewSubjectID("laura"),
		},
		// granted access
		{
			Namespace: "directories",
			Object:    "/photos",
			Relation:  "access",
			Subject:   rts.NewSubjectID("laura"),
		},
	}
	// should be subject set rewrite
	// owners have access
	for _, o := range []struct{ n, o string }{
		{"files", "/photos/beach.jpg"},
		{"files", "/photos/mountains.jpg"},
		{"directories", "/photos"},
	} {
		tuples = append(tuples, &rts.RelationTuple{
			Namespace: o.n,
			Object:    o.o,
			Relation:  "access",
			Subject:   rts.NewSubjectSet(o.n, o.o, "owner"),
		})
	}
	// should be subject set rewrite
	// access on parent means access on child
	for _, obj := range []string{"/photos/beach.jpg", "/photos/mountains.jpg"} {
		tuples = append(tuples, &rts.RelationTuple{
			Namespace: "files",
			Object:    obj,
			Relation:  "access",
			Subject:   rts.NewSubjectSet("directories", "/photos", "access"),
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
