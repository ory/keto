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

	tupleDeltas := []*acl.RelationTupleDelta{
		// ownership
		{
			Action: acl.RelationTupleDelta_INSERT,
			RelationTuple: &acl.RelationTuple{
				Namespace: "directories",
				Object:    "/photos",
				Relation:  "owner",
				Subject: &acl.Subject{Ref: &acl.Subject_Id{
					Id: "maureen",
				}},
			},
		},
		{
			Action: acl.RelationTupleDelta_INSERT,
			RelationTuple: &acl.RelationTuple{
				Namespace: "files",
				Object:    "/photos/beach.jpg",
				Relation:  "owner",
				Subject: &acl.Subject{Ref: &acl.Subject_Id{
					Id: "maureen",
				}},
			},
		},
		{
			Action: acl.RelationTupleDelta_INSERT,
			RelationTuple: &acl.RelationTuple{
				Namespace: "files",
				Object:    "/photos/mountains.jpg",
				Relation:  "owner",
				Subject: &acl.Subject{Ref: &acl.Subject_Id{
					Id: "laura",
				}},
			},
		},
		// granted access
		{
			Action: acl.RelationTupleDelta_INSERT,
			RelationTuple: &acl.RelationTuple{
				Namespace: "directories",
				Object:    "/photos",
				Relation:  "access",
				Subject: &acl.Subject{Ref: &acl.Subject_Id{
					Id: "laura",
				}},
			},
		},
	}
	// should be subject set rewrite
	// owners have access
	for _, o := range []struct{ n, o string }{
		{"files", "/photos/beach.jpg"},
		{"files", "/photos/mountains.jpg"},
		{"directories", "/photos"},
	} {
		tupleDeltas = append(tupleDeltas, &acl.RelationTupleDelta{
			Action: acl.RelationTupleDelta_INSERT,
			RelationTuple: &acl.RelationTuple{
				Namespace: o.n,
				Object:    o.o,
				Relation:  "access",
				Subject: &acl.Subject{Ref: &acl.Subject_Set{Set: &acl.SubjectSet{
					Namespace: o.n,
					Object:    o.o,
					Relation:  "owner",
				}}},
			},
		})
	}
	// should be subject set rewrite
	// access on parent means access on child
	for _, obj := range []string{"/photos/beach.jpg", "/photos/mountains.jpg"} {
		tupleDeltas = append(tupleDeltas, &acl.RelationTupleDelta{
			Action: acl.RelationTupleDelta_INSERT,
			RelationTuple: &acl.RelationTuple{
				Namespace: "files",
				Object:    obj,
				Relation:  "access",
				Subject: &acl.Subject{Ref: &acl.Subject_Set{Set: &acl.SubjectSet{
					Namespace: "directories",
					Object:    "/photos",
					Relation:  "access",
				}}},
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
