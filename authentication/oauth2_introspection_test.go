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
 * @Copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 *
 */

package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOAuth2Introspection(t *testing.T) {
	h := httprouter.New()
	var cb func(w http.ResponseWriter, r *http.Request, req AuthenticationOAuth2IntrospectionRequest) *IntrospectionResponse

	h.POST("/oauth2/introspect", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req AuthenticationOAuth2IntrospectionRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		require.Nil(t, err)

		ir := cb(w, r, req)
		herodot.NewJSONWriter(logrus.New()).Write(w, r, ir)
	})
	ts := httptest.NewServer(h)

	authenticator := &OAuth2IntrospectionAuthentication{
		client:           http.DefaultClient,
		introspectionURL: ts.URL + "/oauth2/introspect",
	}

	now := time.Now().UTC().Round(time.Minute)

	for k, tc := range []struct {
		d               string
		cb              func(w http.ResponseWriter, r *http.Request, req AuthenticationOAuth2IntrospectionRequest) *IntrospectionResponse
		req             *AuthenticationOAuth2IntrospectionRequest
		expectedErr     error
		expectedSession *OAuth2Session
	}{
		{
			cb: func(w http.ResponseWriter, r *http.Request, req AuthenticationOAuth2IntrospectionRequest) *IntrospectionResponse {
				assert.Equal(t, "foo-token", req.Token)
				assert.EqualValues(t, []string{"foo-scope", "foo-scope-a"}, req.Scopes)
				return &IntrospectionResponse{Active: false}
			},
			req:         &AuthenticationOAuth2IntrospectionRequest{Token: "foo-token", Scopes: []string{"foo-scope", "foo-scope-a"}},
			expectedErr: ErrUnauthorized,
		},
		{
			cb: func(w http.ResponseWriter, r *http.Request, req AuthenticationOAuth2IntrospectionRequest) *IntrospectionResponse {
				return &IntrospectionResponse{
					Active:    true,
					Scope:     "scope",
					ClientID:  "scope-ip",
					Subject:   "subject",
					ExpiresAt: now.Unix(),
					IssuedAt:  now.Unix(),
					NotBefore: now.Unix(),
					Username:  "username",
					Audience:  "audience",
					Issuer:    "issuer",
				}
			},
			req: &AuthenticationOAuth2IntrospectionRequest{Token: "foo-token", Scopes: []string{"foo-scope", "foo-scope-a"}},
			expectedSession: &OAuth2Session{
				DefaultSession: &DefaultSession{
					Subject: "subject",
					Allowed: false,
				},
				GrantedScopes: []string{"scope"},
				ClientID:      "scope-ip",
				ExpiresAt:     now,
				IssuedAt:      now,
				NotBefore:     now,
				Username:      "username",
				Audience:      "audience",
				Issuer:        "issuer",
			},
		},
	} {
		t.Run(fmt.Sprintf("case=%d/description=%s", k, tc.d), func(t *testing.T) {
			cb = tc.cb

			out, err := json.Marshal(tc.req)
			require.NoError(t, err)

			r := &http.Request{Body: ioutil.NopCloser(bytes.NewReader(out))}

			session, err := authenticator.Authenticate(r)
			if tc.expectedErr == nil {
				if err != nil {
					require.NoError(t, err, "%+v", err.(stackTracer).StackTrace())
				}
				assert.EqualValues(t, tc.expectedSession, session)
			} else {
				if err == nil {
					require.Error(t, err)
				}
				assert.EqualError(t, err, tc.expectedErr.Error(), "%+v", err.(stackTracer).StackTrace())
			}
		})
	}
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}
