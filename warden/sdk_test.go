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
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package warden_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/fosite"
	"github.com/ory/herodot"
	"github.com/ory/keto/authentication"
	keto "github.com/ory/keto/sdk/go/keto/swagger"
	"github.com/ory/keto/warden"
	"github.com/ory/ladon"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func setupMockOAuth2Introspection(t *testing.T) *httptest.Server {
	h := herodot.NewJSONWriter(nil)
	router := httprouter.New()
	router.POST("/oauth2/token", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if u, p, ok := r.BasicAuth(); !ok || u != "client" || p != "secret" {
			h.WriteError(w, r, errors.New("Basic auth failed"))
			return
		}
		h.Write(w, r, oauth2.Token{
			AccessToken: "access_token",
			TokenType:   "bearer",
			Expiry:      time.Now().Add(time.Hour),
		})
	})

	router.POST("/oauth2/introspect", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if r.Header.Get("Authorization") != "Bearer access_token" {
			h.WriteError(w, r, errors.Errorf("Auth failed: %s", r.Header.Get("Authorization")))
			return
		}

		if err := r.ParseForm(); err != nil {
			h.WriteError(w, r, err)
			return
		}

		if r.PostForm.Get("token") != "alice_token" && r.PostForm.Get("scope") != "fooscope" {
			h.WriteError(w, r, errors.New("Req failed"))
			return
		}

		h.Write(w, r, authentication.IntrospectionResponse{
			Active:  true,
			Subject: "alice",
			Scope:   "fooscope",
		})
	})
	return httptest.NewServer(router)
}

func TestWardenSDK(t *testing.T) {
	oauth2Server := setupMockOAuth2Introspection(t)

	router := httprouter.New()

	handler := warden.NewHandler(herodot.NewJSONWriter(nil), wardens["local"], map[string]authentication.Authenticator{
		"subjects": authentication.NewPlaintextAuthentication(),
		"oauth2/access-tokens": authentication.NewOAuth2IntrospectionAuthentication(
			"client",
			"secret",
			oauth2Server.URL+"/oauth2/token",
			oauth2Server.URL+"/oauth2/introspect",
			[]string{""},
			fosite.HierarchicScopeStrategy,
		),
		"oauth2/clients": authentication.NewOAuth2ClientCredentialsAuthentication(
			oauth2Server.URL + "/oauth2/token",
		),
	})
	handler.SetRoutes(router)
	server := httptest.NewServer(router)

	client := keto.NewWardenApiWithBasePath(server.URL)

	t.Run("IsSubjectAuthorized", func(t *testing.T) {
		for k, c := range accessRequestTestCases {
			t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
				result, response, err := client.IsSubjectAuthorized(keto.WardenSubjectAuthorizationRequest{
					Action:   c.req.Action,
					Resource: c.req.Resource,
					Subject:  c.req.Subject,
					Context:  c.req.Context,
				})

				require.NoError(t, err, "%s", response.Payload)
				require.Equal(t, http.StatusOK, response.StatusCode)
				assert.Equal(t, !c.expectErr, result.Allowed)
			})
		}
	})

	t.Run("IsOAuth2AccessTokenAuthorized", func(t *testing.T) {
		result, response, err := client.IsOAuth2AccessTokenAuthorized(keto.WardenOAuth2AccessTokenAuthorizationRequest{
			Resource: "matrix",
			Action:   "create",
			Context:  ladon.Context{},
			Token:    "alice_token",
			Scope:    []string{"fooscope"},
		})

		require.NoError(t, err, "%s", response.Payload)
		require.Equal(t, http.StatusOK, response.StatusCode, "%s", response.Payload)
		assert.True(t, result.Allowed)
		assert.EqualValues(t, "alice", result.Sub)
	})

	t.Run("IsOAuth2ClientAuthorized", func(t *testing.T) {
		result, response, err := client.IsOAuth2ClientAuthorized(keto.WardenOAuth2ClientAuthorizationRequest{
			Resource:     "matrix",
			Action:       "create",
			ClientId:     "client",
			ClientSecret: "secret",
			Context:      ladon.Context{},
			Scope:        []string{"fooscope"},
		})

		require.NoError(t, err, "%s", response.Payload)
		require.Equal(t, http.StatusOK, response.StatusCode, "%s", response.Payload)
		assert.True(t, result.Allowed)
		assert.EqualValues(t, "client", result.Sub)
	})
}
