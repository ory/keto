//go:build go_mod_indirect_pins
// +build go_mod_indirect_pins

package main

import (
	_ "github.com/mattn/goveralls"
	_ "golang.org/x/tools/cmd/goimports"

	_ "github.com/ory/go-acc"

	// Protobuf and gRPC related tools
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"

	_ "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc"
)
