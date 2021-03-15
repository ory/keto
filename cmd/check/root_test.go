package check

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/internal/namespace"
)

func TestCheckCommand(t *testing.T) {
	nspace := &namespace.Namespace{Name: t.Name()}
	ts := client.NewTestServer(t, client.ReadServer, []*namespace.Namespace{nspace}, newCheckCmd)
	defer ts.Shutdown(t)

	stdOut := ts.Cmd.ExecNoErr(t, "subject", "access", nspace.Name, "object")
	assert.Equal(t, "Denied\n", stdOut)
}
