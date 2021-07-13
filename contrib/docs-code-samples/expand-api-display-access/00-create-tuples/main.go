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

	//directories:/photos#owner@maureen
	//files:/photos/beach.jpg#owner@maureen
	//files:/photos/mountains.jpg#owner@laura
	//directories:/photos#access@laura
	//directories:/photos#access@(directories:/photos#owner)
	//files:/photos/beach.jpg#access@(files:/photos/beach.jpg#owner)
	//files:/photos/beach.jpg#access@(directories:/photos#access)
	//files:/photos/mountains.jpg#access@(files:/photos/mountains.jpg#owner)
	//files:/photos/mountains.jpg#access@(directories:/photos#access)

	tuples := []*acl.RelationTuple{
		// ownership
		{
			Namespace: "directories",
			Object:    "/photos",
			Relation:  "owner",
			Subject:   acl.NewSubjectID("maureen"),
		},
		{
			Namespace: "files",
			Object:    "/photos/beach.jpg",
			Relation:  "owner",
			Subject:   acl.NewSubjectID("maureen"),
		},
		{
			Namespace: "files",
			Object:    "/photos/mountains.jpg",
			Relation:  "owner",
			Subject:   acl.NewSubjectID("laura"),
		},
		// granted access
		{
			Namespace: "directories",
			Object:    "/photos",
			Relation:  "access",
			Subject:   acl.NewSubjectID("laura"),
		},
	}
	// should be subject set rewrite
	// owners have access
	for _, o := range []struct{ n, o string }{
		{"files", "/photos/beach.jpg"},
		{"files", "/photos/mountains.jpg"},
		{"directories", "/photos"},
	} {
		tuples = append(tuples, &acl.RelationTuple{
			Namespace: o.n,
			Object:    o.o,
			Relation:  "access",
			Subject:   acl.NewSubjectSet(o.n, o.o, "owner"),
		})
	}
	// should be subject set rewrite
	// access on parent means access on child
	for _, obj := range []string{"/photos/beach.jpg", "/photos/mountains.jpg"} {
		tuples = append(tuples, &acl.RelationTuple{
			Namespace: "files",
			Object:    obj,
			Relation:  "access",
			Subject:   acl.NewSubjectSet("directories", "/photos", "access"),
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
