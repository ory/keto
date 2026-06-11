// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	rts "github.com/ory/keto/gen/go/ory/keto/relation_tuples/v1alpha2"
	"github.com/ory/keto/ketoapi"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:4466", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := rts.NewExpandServiceClient(conn)

	res, err := client.Expand(context.Background(), &rts.ExpandRequest{
		Subject:  rts.NewSubjectSet("files", "/photos/beach.jpg", "access"),
		MaxDepth: 3,
	})
	if err != nil {
		panic(err)
	}

	tree := ketoapi.TreeFromProto[*ketoapi.RelationTuple](res.Tree)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(tree); err != nil {
		panic(err.Error())
	}
}
