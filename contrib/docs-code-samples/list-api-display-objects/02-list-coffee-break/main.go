package main

import (
	"context"
	"fmt"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:4466", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err.Error())
	}

	client := rts.NewReadServiceClient(conn)

	res, err := client.ListRelationTuples(context.Background(), &rts.ListRelationTuplesRequest{
		Query: &rts.ListRelationTuplesRequest_Query{
			Namespace: "chats",
			Object:    "coffee-break",
			Relation:  "member",
		},
	})
	if err != nil {
		panic(err.Error())
	}

	for _, rt := range res.RelationTuples {
		fmt.Println(rt.Subject.Ref.(*rts.Subject_Id).Id)
	}
}
