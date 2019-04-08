// +build tools

package cmd

import (
	_ "golang.org/x/tools/cmd/cover"
	_ "golang.org/x/tools/cmd/goimports"

	_ "github.com/ory/go-acc"

	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/gobuffalo/packr/packr"
	_ "github.com/sqs/goreturns"

	_ "github.com/ory/x/tools/listx"
)
