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

	"strings"
	"time"

	"net/url"

	"fmt"

	"github.com/ory/fosite"
	"github.com/pkg/errors"
	"golang.org/x/oauth2/clientcredentials"
)

// swagger:model authenticationOAuth2Session
type OAuth2Session struct {
	// Here, it's subject
	*DefaultSession

	// GrantedScopes is a list of scopes that the subject authorized when asked for consent.
	GrantedScopes []string `json:"grantedScopes"`

	// Issuer is the id of the issuer, typically an hydra instance.
	Issuer string `json:"issuer"`

	// ClientID is the id of the OAuth2 client that requested the token.
	ClientID string `json:"clientId"`

	// IssuedAt is the token creation time stamp.
	IssuedAt time.Time `json:"issuedAt"`

	// ExpiresAt is the expiry timestamp.
	ExpiresAt time.Time `json:"expiresAt"`

	NotBefore time.Time `json:"notBefore,omitempty"`
	Username  string    `json:"username,omitempty"`
	Audience  string    `json:"audience,omitempty"`

	// Extra represents arbitrary session data.
	Extra map[string]interface{} `json:"accessTokenExtra"`
}

type IntrospectionResponse struct {
	Active   bool   `json:"active"`
	Scope    string `json:"scope,omitempty"`
	ClientID string `json:"client_id,omitempty"`
	// Here, it's sub
	Subject   string `json:"sub,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Username  string `json:"username,omitempty"`
	Audience  string `json:"aud,omitempty"`
	Issuer    string `json:"iss,omitempty"`
}

type OAuth2IntrospectionAuthentication struct {
	client           *http.Client
	introspectionURL string
	scopeStrategy    fosite.ScopeStrategy
}

// swagger:model AuthenticationOAuth2IntrospectionRequest
type AuthenticationOAuth2IntrospectionRequest struct {
	// Token is the token to introspect.
	Token string `json:"token"`

	// Scopes is an array of scopes that are required.
	Scopes []string `json:"scopes"`
}

func NewOAuth2Session() *OAuth2Session {
	return &OAuth2Session{
		DefaultSession: new(DefaultSession),
	}
}

func NewOAuth2IntrospectionAuthentication(clientID, clientSecret, tokenURL, introspectionURL string, scopes []string, strategy fosite.ScopeStrategy) *OAuth2IntrospectionAuthentication {
	c := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
		Scopes:       scopes,
	}

	return &OAuth2IntrospectionAuthentication{
		client:           c.Client(context.Background()),
		introspectionURL: introspectionURL,
		scopeStrategy:    strategy,
	}
}

func (a *OAuth2IntrospectionAuthentication) Authenticate(r *http.Request) (Session, error) {
	var token AuthenticationOAuth2IntrospectionRequest

	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		return nil, errors.WithStack(err)
	}

	body := url.Values{"token": {token.Token}, "scope": {strings.Join(token.Scopes, " ")}}
	resp, err := a.client.Post(a.introspectionURL, "application/x-www-form-urlencoded", strings.NewReader(body.Encode()))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("Introspection returned status code %d but expected %d", resp.StatusCode, http.StatusOK)
	}

	var ir IntrospectionResponse
	if err := json.NewDecoder(resp.Body).Decode(&ir); err != nil {
		return nil, errors.WithStack(err)
	}

	if !ir.Active {
		return nil, errors.WithStack(ErrUnauthorized.WithReason("Access token introspection says token is not active"))
	}

	for _, scope := range token.Scopes {
		if !a.scopeStrategy(strings.Split(ir.Scope, " "), scope) {
			return nil, errors.WithStack(ErrUnauthorized.WithReason(fmt.Sprintf("Scope %s was not granted", scope)))
		}
	}

	return &OAuth2Session{
		DefaultSession: &DefaultSession{
			Subject: ir.Subject,
		},
		GrantedScopes: strings.Split(ir.Scope, " "),
		ClientID:      ir.ClientID,
		ExpiresAt:     time.Unix(ir.ExpiresAt, 0).UTC(),
		IssuedAt:      time.Unix(ir.IssuedAt, 0).UTC(),
		NotBefore:     time.Unix(ir.NotBefore, 0).UTC(),
		Username:      ir.Username,
		Audience:      ir.Audience,
		Issuer:        ir.Issuer,
	}, nil
}
