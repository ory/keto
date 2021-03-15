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

	ts.Cmd.PersistentArgs = append(ts.Cmd.PersistentArgs, "--"+cmdx.FlagFormat, string(cmdx.FormatJSON))
	stdOut := ts.Cmd.ExecNoErr(t, "access", nspace.Name, "object")
	assert.Equal(t, "null\n", stdOut)
}
