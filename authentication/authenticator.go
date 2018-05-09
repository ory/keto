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
	"net/http"

	"github.com/ory/herodot"
)

var ErrUnauthorized = &herodot.DefaultError{
	CodeField:  http.StatusUnauthorized,
	ErrorField: "The provided credentials are invalid, expired, or are not authorized to use the requested scope",
}

type Session interface {
	GrantAccess()
	DenyAccess()
	GetSubject() string
}

// swagger:model authenticationDefaultSession
type DefaultSession struct {
	// Subject is the identity that authorized issuing the token, for example a user or an OAuth2 app.
	// This is usually a uuid but you can choose a urn or some other id too.
	Subject string `json:"sub"`

	// Allowed is true if the request is allowed and false otherwise.
	Allowed bool `json:"allowed"`
}

func (s *DefaultSession) GrantAccess() {
	s.Allowed = true
}

func (s *DefaultSession) DenyAccess() {
	s.Allowed = false
}

func (s *DefaultSession) GetSubject() string {
	return s.Subject
}

type Authenticator interface {
	Authenticate(r *http.Request) (Session, error)
}
