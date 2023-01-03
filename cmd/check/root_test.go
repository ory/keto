// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check

import (
	"testing"

	"github.com/stretchr/testify/require"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/internal/namespace"
)

func TestCheckCommand(t *testing.T) {
	nspace := &namespace.Namespace{Name: t.Name()}
	ts := client.NewTestServer(t, client.ReadServer, []*namespace.Namespace{nspace}, NewCheckCmd)
	defer ts.Shutdown(t)

	stdOut := ts.Cmd.ExecNoErr(t, "subject", "access", nspace.Name, "object",
		"--insecure-skip-hostname-verification",
	)
	assert.Equal(t, "Denied\n", stdOut)
}

func TestParseSubject(t *testing.T) {
	for _, tc := range []struct {
		input    string
		expected *rts.Subject
	}{
		{
			input:    "nspace:obj#rel",
			expected: rts.NewSubjectSet("nspace", "obj", "rel"),
		},
		{
			input:    "someid",
			expected: rts.NewSubjectID("someid"),
		},
	} {
		t.Run("subject="+tc.input, func(t *testing.T) {
			actual, err := parseSubject(tc.input)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
