// +build go_mod_indirect_pins

package main

import (
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/mattn/goveralls"
	_ "github.com/sqs/goreturns"
	_ "golang.org/x/tools/cmd/cover"
	_ "golang.org/x/tools/cmd/goimports"

	// FIXME pins websocket to 1.4.2
	// FIXME See https://github.com/gobuffalo/buffalo/pull/1999
	_ "github.com/gorilla/websocket"

	_ "github.com/ory/cli"
	_ "github.com/ory/go-acc"
	_ "github.com/ory/x/tools/listx"

	// Protobuf and gRPC related tools
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/golang/mock/mockgen"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"

	_ "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc"

	_ "github.com/goreleaser/godownloader"
)
