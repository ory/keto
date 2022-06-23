package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
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
			Relation:  "member",
			Subject:   rts.NewSubjectID("PM"),
		},
	})
	if err != nil {
		panic(err.Error())
	}

	for _, rt := range res.RelationTuples {
		fmt.Println(rt.Object)
	}
}
