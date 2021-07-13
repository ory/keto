package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/ory/keto/internal/expand"

	"google.golang.org/grpc"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:4466", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := acl.NewExpandServiceClient(conn)

	res, err := client.Expand(context.Background(), &acl.ExpandRequest{
		Subject:  acl.NewSubjectSet("files", "/photos/beach.jpg", "access"),
		MaxDepth: 3,
	})
	if err != nil {
		panic(err)
	}

	tree, err := expand.TreeFromProto(res.Tree)
	if err != nil {
		panic(err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(tree); err != nil {
		panic(err.Error())
	}
}
