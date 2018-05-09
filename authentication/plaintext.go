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
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func NewPlaintextAuthentication() *PlaintextAuthentication {
	return &PlaintextAuthentication{}
}

type PlaintextAuthentication struct {
	client           *http.Client
	introspectionURL string
}

func (a *PlaintextAuthentication) Authenticate(r *http.Request) (Session, error) {
	var session struct {
		Subject string `json:"subject"`
	}

	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		return nil, errors.WithStack(err)
	}

	return &DefaultSession{Subject: session.Subject}, nil
}
