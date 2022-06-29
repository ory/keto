package main

import (
	"context"
	"fmt"
	"sort"

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

	objects := []string{}
	for _, rt := range res.RelationTuples {
		objects = append(objects, rt.Object)
	}
	sort.Strings(objects)
	for _, o := range objects {
		fmt.Println(o)
	}
}
