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

package client

import (
	"fmt"
	"net/http"

	keto "github.com/ory/keto/sdk/go/keto/swagger"
	"github.com/spf13/cobra"
)

type WardenHandler struct{}

func (h *WardenHandler) newWardenManager(cmd *cobra.Command) *keto.WardenApi {
	c := keto.NewWardenApiWithBasePath(getBasePath(cmd))

	if token, err := cmd.Flags().GetString("bearer-token"); err == nil && token != "" {
		c.Configuration.DefaultHeader["Authorization"] = "Bearer " + token
	}

	if term, _ := cmd.Flags().GetBool("fake-tls-termination"); term {
		c.Configuration.DefaultHeader["X-Forwarded-Proto"] = "https"
	}
	return c
}

func newWardenHandler() *WardenHandler {
	return &WardenHandler{}
}

func (h *WardenHandler) IsOAuth2AccessTokenAuthorized(cmd *cobra.Command, args []string) {
	token, _ := cmd.Flags().GetString("token")
	scope, _ := cmd.Flags().GetStringArray("scope")
	action, _ := cmd.Flags().GetString("action")
	resource, _ := cmd.Flags().GetString("resource")

	m := h.newWardenManager(cmd)
	_, response, err := m.IsOAuth2AccessTokenAuthorized(keto.WardenOAuth2AccessTokenAuthorizationRequest{
		Token:    token,
		Scope:    scope,
		Action:   action,
		Resource: resource,
	})
	checkResponse(response, err, http.StatusOK)
	fmt.Printf("%s\n", response.Payload)
}

func (h *WardenHandler) IsSubjectAuthorized(cmd *cobra.Command, args []string) {
	subject, _ := cmd.Flags().GetString("subject")
	action, _ := cmd.Flags().GetString("action")
	resource, _ := cmd.Flags().GetString("resource")

	m := h.newWardenManager(cmd)
	_, response, err := m.IsSubjectAuthorized(keto.WardenSubjectAuthorizationRequest{
		Action:   action,
		Subject:  subject,
		Resource: resource,
	})
	checkResponse(response, err, http.StatusOK)
	fmt.Printf("%s\n", response.Payload)
}
