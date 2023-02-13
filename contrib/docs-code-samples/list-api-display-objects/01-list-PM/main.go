// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"sort"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/ory/x/pointerx"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:4466", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err.Error())
	}

	client := rts.NewReadServiceClient(conn)

	res, err := client.ListRelationTuples(context.Background(), &rts.ListRelationTuplesRequest{
		RelationQuery: &rts.RelationQuery{
			Namespace: pointerx.Ptr("chats"),
			Relation:  pointerx.Ptr("member"),
			Subject:   rts.NewSubjectID("PM"),
		},
	})
	if err != nil {
		panic(err.Error())
	}

	objects := make([]string, len(res.RelationTuples))
	for i, rt := range res.RelationTuples {
		objects[i] = rt.Object
	}
	sort.Strings(objects)
	for _, o := range objects {
		fmt.Println(o)
	}
}
