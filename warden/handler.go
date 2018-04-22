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

package warden

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bytes"
	"io/ioutil"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"github.com/ory/keto/authentication"
	"github.com/pkg/errors"
)

const (
	AuthenticatorHandlerPath = "/warden/%s/authorize"

	// AllowedHandlerPath points to the access request validation endpoint.
	//AllowedHandlerPath = "/warden/oauth2/access-tokens/authorize"
	//AllowedHandlerPath = "/warden/oauth2/clients/authorize"
	//AllowedHandlerPath = "/warden/subjects/authorize"
	//AllowedHandlerPath = "/warden/jwt/authorize"
	//AllowedHandlerPath = "/warden/saml/authorize"
)

var notAllowed = struct {
	Allowed bool `json:"allowed"`
}{Allowed: false}

// Handler is capable of handling HTTP request and validating access tokens and access requests.
type Handler struct {
	H      herodot.Writer
	Warden Firewall

	ResourcePrefix string
	authenticators map[string]authentication.Authenticator
}

func NewHandler(writer herodot.Writer, warden Firewall, authenticators map[string]authentication.Authenticator) *Handler {
	h := &Handler{
		H:              writer,
		Warden:         warden,
		authenticators: authenticators,
	}

	return h
}

func (h *Handler) SetRoutes(r *httprouter.Router) {
	for k, a := range h.authenticators {
		r.POST(fmt.Sprintf(AuthenticatorHandlerPath, k), h.authorized(a))
	}
}

func (h *Handler) authorized(authenticator authentication.Authenticator) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		all, err := ioutil.ReadAll(r.Body)
		if err != nil {
			h.H.WriteError(w, r, errors.WithStack(err))
			return
		}

		var ctx = r.Context()
		r.Body = ioutil.NopCloser(bytes.NewReader(all))
		session, err := authenticator.Authenticate(r)
		if err != nil && errors.Cause(err).Error() == authentication.ErrUnauthorized.Error() {
			h.H.Write(w, r, &notAllowed)
			return
		} else if err != nil {
			h.H.WriteError(w, r, err)
			return
		}

		var access AccessRequest
		if err := json.Unmarshal(all, &access); err != nil {
			h.H.WriteError(w, r, errors.WithStack(err))
			return
		}

		access.Subject = session.GetSubject()
		if err := h.Warden.IsAllowed(ctx, &access); err != nil {
			h.H.Write(w, r, &notAllowed)
			return
		}

		session.GrantAccess()
		h.H.Write(w, r, session)
	}
}
