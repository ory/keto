// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package step

import (
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m,
		goleak.IgnoreTopFunction("github.com/dgraph-io/ristretto/v2.(*defaultPolicy[...]).processItems"),
		goleak.IgnoreTopFunction("github.com/dgraph-io/ristretto/v2.(*Cache[...]).processItems"),
	)
}
