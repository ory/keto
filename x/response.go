package x

import (
	"bytes"
	"github.com/ory/keto/sdk/go/keto/swagger"
	"github.com/ory/x/cmdx"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"testing"
)

func CheckResponse(err error, expectedStatusCode int, response *swagger.APIResponse) {
	var r *http.Response
	if response != nil {
		r = response.Response
		r.Body = ioutil.NopCloser(bytes.NewBuffer(response.Payload))
	}

	cmdx.CheckResponse(err, expectedStatusCode, r)
}

func CheckResponseTest(t *testing.T, err error, expectedStatusCode int, response *swagger.APIResponse) {
	require.NoError(t, err, "%s %s: %s", response.Request.Method, response.RequestURL, response.Payload)
	require.Equal(t, expectedStatusCode, response.StatusCode, "%s %s (%d): %s", response.Request.Method, response.RequestURL, response.StatusCode, response.Payload)
}
