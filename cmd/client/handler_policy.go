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
	"os"

	"net/http"

	keto "github.com/ory/keto/sdk/go/keto/swagger"
	"github.com/ory/ladon"
	"github.com/spf13/cobra"
	"github.com/square/go-jose/json"
)

type PolicyHandler struct{}

func (h *PolicyHandler) newPolicyManager(cmd *cobra.Command) *keto.PolicyApi {
	c := keto.NewPolicyApiWithBasePath(getBasePath(cmd))

	if token, err := cmd.Flags().GetString("bearer-token"); err == nil && token != "" {
		c.Configuration.DefaultHeader["Authorization"] = "Bearer " + token
	}

	if term, _ := cmd.Flags().GetBool("fake-tls-termination"); term {
		c.Configuration.DefaultHeader["X-Forwarded-Proto"] = "https"
	}
	return c
}

func newPolicyHandler() *PolicyHandler {
	return &PolicyHandler{}
}

func (h *PolicyHandler) ImportPolicy(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println(cmd.UsageString())
		return
	}

	m := h.newPolicyManager(cmd)

	for _, path := range args {
		reader, err := os.Open(path)
		must(err, "Could not open file %s: %s", path, err)

		var p keto.Policy
		err = json.NewDecoder(reader).Decode(&p)
		must(err, "Could not parse JSON: %s", err)

		_, response, err := m.CreatePolicy(p)
		checkResponse(response, err, http.StatusCreated)
		fmt.Printf("Imported policy %s from %s.\n", p.Id, path)
	}

	return
}

func (h *PolicyHandler) CreatePolicy(cmd *cobra.Command, args []string) {
	m := h.newPolicyManager(cmd)

	if files, _ := cmd.Flags().GetStringSlice("files"); len(files) > 0 {
		fmt.Println("Importing policies using the -f flag is deprecated and will be removed in the future.")
		fmt.Println(`Please use "hydra policies import" instead.`)
		h.ImportPolicy(cmd, files)
		return
	}

	id, _ := cmd.Flags().GetString("id")
	description, _ := cmd.Flags().GetString("description")
	subjects, _ := cmd.Flags().GetStringSlice("subjects")
	resources, _ := cmd.Flags().GetStringSlice("resources")
	actions, _ := cmd.Flags().GetStringSlice("actions")
	isAllow, _ := cmd.Flags().GetBool("allow")

	if len(subjects) == 0 || len(resources) == 0 || len(actions) == 0 {
		fmt.Println(cmd.UsageString())
		fmt.Println("")
		fmt.Println("Got empty subject, resource or action list")
		return
	}

	effect := ladon.DenyAccess
	if isAllow {
		effect = ladon.AllowAccess
	}

	result, response, err := m.CreatePolicy(keto.Policy{
		Id:          id,
		Description: description,
		Subjects:    subjects,
		Resources:   resources,
		Actions:     actions,
		Effect:      effect,
	})
	checkResponse(response, err, http.StatusCreated)
	fmt.Printf("Created policy %s.\n", result.Id)
}

func (h *PolicyHandler) AddResourceToPolicy(cmd *cobra.Command, args []string) {
	m := h.newPolicyManager(cmd)
	if len(args) < 2 {
		fmt.Print(cmd.UsageString())
		return
	}

	p, response, err := m.GetPolicy(args[0])
	checkResponse(response, err, http.StatusOK)

	p.Resources = append(p.Resources, args[1:]...)

	_, response, err = m.UpdatePolicy(p.Id, *p)
	checkResponse(response, err, http.StatusOK)
	fmt.Printf("Added resources to policy %s", p.Id)
}

func (h *PolicyHandler) RemoveResourceFromPolicy(cmd *cobra.Command, args []string) {
	m := h.newPolicyManager(cmd)
	if len(args) < 2 {
		fmt.Print(cmd.UsageString())
		return
	}

	p, response, err := m.GetPolicy(args[0])
	checkResponse(response, err, http.StatusOK)

	resources := []string{}
	for _, r := range p.Resources {
		var filter bool
		for _, a := range args[1:] {
			if r == a {
				filter = true
			}
		}
		if !filter {
			resources = append(resources, r)
		}
	}
	p.Resources = resources

	_, response, err = m.UpdatePolicy(p.Id, *p)
	checkResponse(response, err, http.StatusOK)
	fmt.Printf("Removed resources from policy %s", p.Id)
}

func (h *PolicyHandler) AddSubjectToPolicy(cmd *cobra.Command, args []string) {
	m := h.newPolicyManager(cmd)
	if len(args) < 2 {
		fmt.Print(cmd.UsageString())
		return
	}

	p, response, err := m.GetPolicy(args[0])
	checkResponse(response, err, http.StatusOK)

	p.Subjects = append(p.Subjects, args[1:]...)

	_, response, err = m.UpdatePolicy(p.Id, *p)
	checkResponse(response, err, http.StatusOK)
	fmt.Printf("Added subjects to policy %s", p.Id)
}

func (h *PolicyHandler) RemoveSubjectFromPolicy(cmd *cobra.Command, args []string) {
	m := h.newPolicyManager(cmd)
	if len(args) < 2 {
		fmt.Print(cmd.UsageString())
		return
	}

	p, response, err := m.GetPolicy(args[0])
	checkResponse(response, err, http.StatusOK)

	subjects := []string{}
	for _, r := range p.Subjects {
		var filter bool
		for _, a := range args[1:] {
			if r == a {
				filter = true
			}
		}
		if !filter {
			subjects = append(subjects, r)
		}
	}
	p.Subjects = subjects

	_, response, err = m.UpdatePolicy(p.Id, *p)
	checkResponse(response, err, http.StatusOK)
	fmt.Printf("Removed subjects from policy %s.\n", p.Id)
}

func (h *PolicyHandler) AddActionToPolicy(cmd *cobra.Command, args []string) {
	m := h.newPolicyManager(cmd)
	if len(args) < 2 {
		fmt.Print(cmd.UsageString())
		return
	}

	p, response, err := m.GetPolicy(args[0])
	checkResponse(response, err, http.StatusOK)

	p.Actions = append(p.Actions, args[1:]...)

	_, response, err = m.UpdatePolicy(p.Id, *p)
	checkResponse(response, err, http.StatusOK)
	fmt.Printf("Added actions to policy %s.\n", p.Id)
}

func (h *PolicyHandler) RemoveActionFromPolicy(cmd *cobra.Command, args []string) {
	m := h.newPolicyManager(cmd)
	if len(args) < 2 {
		fmt.Print(cmd.UsageString())
		return
	}

	p, response, err := m.GetPolicy(args[0])
	checkResponse(response, err, http.StatusOK)

	actions := []string{}
	for _, r := range p.Actions {
		var filter bool
		for _, a := range args[1:] {
			if r == a {
				filter = true
			}
		}
		if !filter {
			actions = append(actions, r)
		}
	}
	p.Actions = actions

	_, response, err = m.UpdatePolicy(p.Id, *p)
	checkResponse(response, err, http.StatusOK)
	fmt.Printf("Removed actions from policy %s.\n", p.Id)
}

func (h *PolicyHandler) GetPolicy(cmd *cobra.Command, args []string) {
	m := h.newPolicyManager(cmd)
	if len(args) == 0 {
		fmt.Print(cmd.UsageString())
		return
	}

	p, response, err := m.GetPolicy(args[0])
	checkResponse(response, err, http.StatusOK)

	fmt.Printf("%s\n", formatResponse(p))
}

func (h *PolicyHandler) DeletePolicy(cmd *cobra.Command, args []string) {
	m := h.newPolicyManager(cmd)
	if len(args) == 0 {
		fmt.Print(cmd.UsageString())
		return
	}

	for _, arg := range args {
		response, err := m.DeletePolicy(arg)
		checkResponse(response, err, http.StatusNoContent)
		fmt.Printf("Policy %s deleted.\n", arg)
	}
}
