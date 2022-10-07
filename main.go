// Copyright © 2022 Ory Corp

package main

import (
	"github.com/ory/x/profilex"

	"github.com/ory/keto/cmd"
)

func main() {
	defer profilex.Profile().Stop()
	cmd.Execute()
}
