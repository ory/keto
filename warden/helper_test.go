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
	"os"
	"testing"

	"github.com/ory/keto/role"
	"github.com/ory/keto/warden"
	"github.com/ory/ladon"
	"github.com/ory/ladon/manager/memory"
	"github.com/sirupsen/logrus"
)

var (
	accessRequestTestCases = []struct {
		req       *warden.AccessRequest
		expectErr bool
	}{
		{
			req: &warden.AccessRequest{
				Subject:  "alice",
				Resource: "other-thing",
				Action:   "create",
				Context:  ladon.Context{},
			},
			expectErr: true,
		},
		{
			req: &warden.AccessRequest{
				Subject:  "alice",
				Resource: "matrix",
				Action:   "delete",
				Context:  ladon.Context{},
			},
			expectErr: true,
		},
		{
			req: &warden.AccessRequest{
				Subject:  "alice",
				Resource: "matrix",
				Action:   "create",
				Context:  ladon.Context{},
			},
			expectErr: false,
		},
		{
			req: &warden.AccessRequest{
				Subject:  "ken",
				Resource: "forbidden_matrix",
				Action:   "create",
				Context:  ladon.Context{},
			},
			expectErr: true,
		},
		{
			req: &warden.AccessRequest{
				Subject:  "ken",
				Resource: "allowed_matrix",
				Action:   "create",
				Context:  ladon.Context{},
			},
			expectErr: false,
		},
	}
	wardens     = map[string]warden.Firewall{}
	ladonWarden = &ladon.Ladon{
		Manager: &memory.MemoryManager{
			Policies: map[string]ladon.Policy{
				"1": &ladon.DefaultPolicy{
					ID:        "1",
					Subjects:  []string{"alice", "group1", "client"},
					Resources: []string{"matrix", "forbidden_matrix", "rn:hydra:token<.*>"},
					Actions:   []string{"create", "decide"},
					Effect:    ladon.AllowAccess,
				},
				"2": &ladon.DefaultPolicy{
					ID:        "2",
					Subjects:  []string{"siri"},
					Resources: []string{"<.*>"},
					Actions:   []string{"decide"},
					Effect:    ladon.AllowAccess,
				},
				"3": &ladon.DefaultPolicy{
					ID:        "3",
					Subjects:  []string{"group1"},
					Resources: []string{"forbidden_matrix", "rn:hydra:token<.*>"},
					Actions:   []string{"create", "decide"},
					Effect:    ladon.DenyAccess,
				},
				"4": &ladon.DefaultPolicy{
					ID:        "4",
					Subjects:  []string{"group1"},
					Resources: []string{"allowed_matrix", "rn:hydra:token<.*>"},
					Actions:   []string{"create", "decide"},
					Effect:    ladon.AllowAccess,
				},
			},
		},
	}
)

func TestMain(m *testing.M) {
	wardens["local"] = &warden.Warden{
		L:      logrus.New(),
		Warden: ladonWarden,
		Roles: &role.MemoryManager{
			Roles: map[string]role.Role{
				"group1": {
					ID:      "group1",
					Members: []string{"ken"},
				},
			},
		},
	}

	os.Exit(m.Run())
}
