package main

import (
	"context"
	"fmt"
	"sort"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:4466", grpc.WithInsecure())
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
