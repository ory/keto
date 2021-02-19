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

	client := acl.NewCheckServiceClient(conn)

	res, err := client.Check(context.Background(), &acl.CheckRequest{
		Namespace: "messages",
		Object:    "02y_15_4w350m3",
		Relation:  "decypher",
		Subject: &acl.Subject{Ref: &acl.Subject_Id{
			Id: "john",
		}},
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Allowed: %v\n", res.Allowed)
}
