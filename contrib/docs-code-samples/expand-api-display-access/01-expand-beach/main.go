package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/ory/keto/ketoapi"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:4466", grpc.WithInsecure())
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

	tree := (&ketoapi.ExpandTree{}).FromProto(res.Tree)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(tree); err != nil {
		panic(err.Error())
	}
}
