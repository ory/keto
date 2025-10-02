// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:4466", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	marshaler := protojson.MarshalOptions{EmitUnpopulated: true}
	bs, err := marshaler.Marshal(res.Tree)
	if err != nil {
		panic(err)
	}
	if _, err = os.Stdout.Write(bs); err != nil {
		panic(err)
	}
}
