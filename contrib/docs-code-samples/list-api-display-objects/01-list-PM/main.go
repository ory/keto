package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:4466", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}

	client := acl.NewReadServiceClient(conn)

	res, err := client.ListRelationTuples(context.Background(), &acl.ListRelationTuplesRequest{
		Query: &acl.ListRelationTuplesRequest_Query{
			Namespace: "chats",
			Relation:  "member",
			Subject: &acl.Subject{Ref: &acl.Subject_Id{
				Id: "PM",
			}},
		},
	})
	if err != nil {
		panic(err.Error())
	}

	for _, rt := range res.RelationTuples {
		fmt.Println(rt.Object)
	}
}
