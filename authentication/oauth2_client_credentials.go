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
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/oauth2/clientcredentials"
)

// swagger:model authenticationOAuth2ClientCredentialsSession
type OAuth2ClientCredentialsSession struct {
	// Here, it's subject
	*DefaultSession
}

type OAuth2ClientCredentialsAuthentication struct {
	tokenURL string
}

// swagger:model AuthenticationOAuth2ClientCredentialsRequest
type AuthenticationOAuth2ClientCredentialsRequest struct {
	// Token is the token to introspect.
	ClientID string `json:"client_id"`

	ClientSecret string `json:"client_secret"`

	// Scope is an array of scopes that are required.
	Scopes []string `json:"scope"`
}

func NewOAuth2ClientCredentialsSession() *OAuth2ClientCredentialsSession {
	return &OAuth2ClientCredentialsSession{
		DefaultSession: new(DefaultSession),
	}
}

func NewOAuth2ClientCredentialsAuthentication(tokenURL string) *OAuth2ClientCredentialsAuthentication {
	return &OAuth2ClientCredentialsAuthentication{
		tokenURL: tokenURL,
	}
}

func (a *OAuth2ClientCredentialsAuthentication) Authenticate(r *http.Request) (Session, error) {
	var auth AuthenticationOAuth2ClientCredentialsRequest

	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil {
		return nil, errors.WithStack(err)
	}

	c := &clientcredentials.Config{
		TokenURL:     a.tokenURL,
		ClientID:     auth.ClientID,
		ClientSecret: auth.ClientSecret,
		Scopes:       auth.Scopes,
	}

	token, err := c.Token(context.Background())
	if err != nil {
		return nil, errors.WithStack(ErrUnauthorized)
	} else if token.AccessToken == "" {
		return nil, errors.WithStack(ErrUnauthorized)
	}

	return &OAuth2ClientCredentialsSession{
		DefaultSession: &DefaultSession{
			Subject: auth.ClientID,
		},
	}, nil
}
