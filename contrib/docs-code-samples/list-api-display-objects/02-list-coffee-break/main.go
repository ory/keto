// Copyright © 2022 Ory Corp

package main

import (
	"context"
	"fmt"
	"sort"

	"github.com/ory/keto/internal/x"

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
		RelationQuery: &rts.RelationQuery{
			Namespace: x.Ptr("chats"),
			Object:    x.Ptr("coffee-break"),
			Relation:  x.Ptr("member"),
		},
	})
	if err != nil {
		panic(err.Error())
	}

	ids := make([]string, len(res.RelationTuples))
	for i, rt := range res.RelationTuples {
		ids[i] = rt.Subject.Ref.(*rts.Subject_Id).Id
	}
	sort.Strings(ids)
	for _, id := range ids {
		fmt.Println(id)
	}
}
