/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author        Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @Copyright     2017-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license     Apache-2.0
 */

package health

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"github.com/ory/keto/sdk/go/keto/swagger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealth(t *testing.T) {
	alive := errors.New("not alive")
	handler := &Handler{
		H:             herodot.NewJSONWriter(nil),
		VersionString: "test version",
		ReadyChecks: map[string]ReadyChecker{
			"test": func() error {
				return alive
			},
		},
	}
	router := httprouter.New()
	handler.SetRoutes(router)
	ts := httptest.NewServer(router)

	healthClient := swagger.NewHealthApiWithBasePath(ts.URL)

	body, response, err := healthClient.IsInstanceAlive()
	require.NoError(t, err)
	require.EqualValues(t, http.StatusOK, response.StatusCode)
	assert.EqualValues(t, "ok", body.Status)

	versionClient := swagger.NewVersionApiWithBasePath(ts.URL)
	version, response, err := versionClient.GetVersion()
	require.NoError(t, err)
	require.EqualValues(t, http.StatusOK, response.StatusCode)
	require.EqualValues(t, version.Version, handler.VersionString)

	_, response, err = healthClient.IsInstanceReady()
	require.NoError(t, err)
	require.EqualValues(t, http.StatusServiceUnavailable, response.StatusCode)
	assert.Equal(t, `{"errors":{"test":"not alive"}}`, string(response.Payload))

	alive = nil
	body, response, err = healthClient.IsInstanceReady()
	require.NoError(t, err)
	require.EqualValues(t, http.StatusOK, response.StatusCode)
	assert.EqualValues(t, "ok", body.Status)
}
