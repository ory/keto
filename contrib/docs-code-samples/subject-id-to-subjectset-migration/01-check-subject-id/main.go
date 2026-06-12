// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	rts "github.com/ory/keto/gen/go/ory/keto/relation_tuples/v1alpha2"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:4466", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := rts.NewCheckServiceClient(conn)

	res, err := client.Check(context.Background(), &rts.CheckRequest{
		Namespace: "File",
		Object:    "data.txt",
		Relation:  "viewer",
		Subject:   rts.NewSubjectID("user_alice"),
	})
	if err != nil {
		panic(err)
	}

	if res.Allowed {
		fmt.Println("Allowed")
		return
	}
	fmt.Println("Denied")
}
