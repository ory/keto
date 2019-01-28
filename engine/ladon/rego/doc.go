// Package rego is a placeholder package to force download this package and its subpackages when we run `go mod vendor`
package rego

import (
	_ "github.com/ory/keto/engine/ladon/rego/condition"
	_ "github.com/ory/keto/engine/ladon/rego/core"
	_ "github.com/ory/keto/engine/ladon/rego/exact"
	_ "github.com/ory/keto/engine/ladon/rego/glob"
	_ "github.com/ory/keto/engine/ladon/rego/regex"
)
