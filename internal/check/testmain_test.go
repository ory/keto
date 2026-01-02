// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check_test

import (
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m,
		goleak.IgnoreTopFunction("github.com/dgraph-io/ristretto/v2.(*defaultPolicy[...]).processItems"),
		goleak.IgnoreTopFunction("github.com/dgraph-io/ristretto/v2.(*Cache[...]).processItems"),
		// fixed-size worker pool:
		goleak.IgnoreTopFunction("github.com/ory/keto/internal/check/checkgroup.worker"),
		goleak.IgnoreTopFunction("net/http.(*persistConn).readLoop"),
		goleak.IgnoreTopFunction("net/http.(*persistConn).writeLoop"),
		goleak.IgnoreTopFunction("database/sql.(*DB).connectionOpener"),
	)
}
