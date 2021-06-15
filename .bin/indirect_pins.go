// +build go_mod_indirect_pins

package _bin

import (
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "golang.org/x/tools/cmd/goimports"

	_ "github.com/ory/cli"

	// Protobuf and gRPC related tools
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"

	_ "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc"
)
