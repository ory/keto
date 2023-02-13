// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"context"
	"testing"
	"time"

	"github.com/ory/x/cmdx"
	"github.com/stretchr/testify/assert"
)

func TestDeleteAllCmd(t *testing.T) {
	t.Run("executes get instead of delete when run without force", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		cmd := NewDeleteAllCmd()
		// we will get an error here because there is no server running, but we really only care about the execution path
		stdout, _, _ := cmdx.ExecCtx(ctx, cmd, nil, "--insecure-skip-hostname-verification=true")
		assert.Contains(t, stdout, "WARNING: This operation is not reversible.")
	})
}
