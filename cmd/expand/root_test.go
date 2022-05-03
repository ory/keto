package expand

import (
	"testing"

	"github.com/ory/x/cmdx"
	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/internal/namespace"
)

func TestExpandCommand(t *testing.T) {
	nspace := &namespace.Namespace{Name: t.Name()}
	ts := client.NewTestServer(t, client.ReadServer, []*namespace.Namespace{nspace}, NewExpandCmd)
	defer ts.Shutdown(t)

	t.Run("case=unknown tuple", func(t *testing.T) {
		t.Run("format=JSON", func(t *testing.T) {
			stdOut := ts.Cmd.ExecNoErr(t,
				"access", nspace.Name,
				"object", "--"+cmdx.FlagFormat, string(cmdx.FormatJSON))
			assert.Equal(t, "null\n", stdOut)
		})

		t.Run("format=default", func(t *testing.T) {
			stdOut := ts.Cmd.ExecNoErr(t,
				"access", nspace.Name,
				"object", "--"+cmdx.FlagFormat, string(cmdx.FormatDefault))
			assert.Contains(t, stdOut, "empty tree")
		})
	})
}
