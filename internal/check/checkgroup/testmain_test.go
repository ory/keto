// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package checkgroup_test

import (
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m, goleak.IgnoreCurrent())
}
