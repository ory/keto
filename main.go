// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/ory/x/profilex"

	"github.com/ory/keto/cmd"
)

func main() {
	defer profilex.Profile().Stop()
	cmd.Execute()
}
