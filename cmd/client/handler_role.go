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

type RoleHandler struct {
}

func (h *RoleHandler) newPolicyManager(cmd *cobra.Command) *keto.RoleApi {
	c := keto.NewRoleApiWithBasePath(getBasePath(cmd))

	if token, err := cmd.Flags().GetString("bearer-token"); err == nil && token != "" {
		c.Configuration.DefaultHeader["Authorization"] = "Bearer " + token
	}

	if term, _ := cmd.Flags().GetBool("fake-tls-termination"); term {
		c.Configuration.DefaultHeader["X-Forwarded-Proto"] = "https"
	}
	return c
}

func newRoleHandler() *RoleHandler {
	return &RoleHandler{}
}

func (h *RoleHandler) CreateRole(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Print(cmd.UsageString())
		return
	}
	m := h.newPolicyManager(cmd)

	_, response, err := m.CreateRole(keto.Role{Id: args[0]})
	checkResponse(response, err, http.StatusCreated)
	fmt.Printf("Group %s created.\n", args[0])
}

func (h *RoleHandler) DeleteRole(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Print(cmd.UsageString())
		return
	}

	m := h.newPolicyManager(cmd)
	response, err := m.DeleteRole(args[0])
	checkResponse(response, err, http.StatusNoContent)
	fmt.Printf("Group %s deleted.\n", args[0])
}

func (h *RoleHandler) RoleAddMembers(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Print(cmd.UsageString())
		return
	}

	m := h.newPolicyManager(cmd)
	response, err := m.AddMembersToRole(args[0], keto.RoleMembers{Members: args[1:]})
	checkResponse(response, err, http.StatusNoContent)
	fmt.Printf("Members %v added to group %s.\n", args[1:], args[0])
}

func (h *RoleHandler) RoleRemoveMembers(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Print(cmd.UsageString())
		return
	}

	m := h.newPolicyManager(cmd)
	response, err := m.RemoveMembersFromRole(args[0], keto.RoleMembers{Members: args[1:]})
	checkResponse(response, err, http.StatusNoContent)
	fmt.Printf("Members %v removed from group %s.\n", args[1:], args[0])
}

func (h *RoleHandler) FindRoles(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Print(cmd.UsageString())
		return
	}

	m := h.newPolicyManager(cmd)
	groups, response, err := m.ListRoles(args[0], 500, 0)
	checkResponse(response, err, http.StatusOK)
	formatResponse(groups)
}

func (h *RoleHandler) ListRoles(cmd *cobra.Command, args []string) {
	if len(args) != 0 {
		fmt.Print(cmd.UsageString())
		return
	}

	m := h.newPolicyManager(cmd)
	groups, response, err := m.ListRoles("", 500, 0)
	checkResponse(response, err, http.StatusOK)
	formatResponse(groups)
}

func (h *RoleHandler) GetRole(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Print(cmd.UsageString())
		return
	}

	m := h.newPolicyManager(cmd)
	groups, response, err := m.GetRole(args[0])
	checkResponse(response, err, http.StatusOK)
	formatResponse(groups)
}
