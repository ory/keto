// +build go_mod_indirect_pins

package main

import (
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/gobuffalo/packr/packr"
	_ "github.com/sqs/goreturns"
	_ "golang.org/x/tools/cmd/cover"
	_ "golang.org/x/tools/cmd/goimports"

	// FIXME pins websocket to 1.4.2
	// FIXME See https://github.com/gobuffalo/buffalo/pull/1999
	_ "github.com/gorilla/websocket"

	_ "github.com/ory/cli"
	_ "github.com/ory/go-acc"
	_ "github.com/ory/x/tools/listx"
)
