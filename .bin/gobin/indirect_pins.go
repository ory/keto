//go:build go_mod_indirect_pins
// +build go_mod_indirect_pins

package main

import (
	_ "github.com/josephburnett/jd"
	_ "github.com/mattn/goveralls"
	_ "github.com/ory/go-acc"
	_ "golang.org/x/tools/cmd/goimports"

	// Protobuf and gRPC related tools
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"

	_ "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc"
)
