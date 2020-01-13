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

package cmd

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ory/viper"
)

func TestExecute(t *testing.T) {
	viper.Set("dsn", "memory")
	ep := fmt.Sprintf("http://127.0.0.1:%d", port)

	for _, c := range []struct {
		args      []string
		wait      func(t *testing.T) bool
		expectErr bool
	}{
		{
			args: []string{"serve", "--disable-telemetry"},
			wait: func(t *testing.T) bool {
				t.Logf("Trying to connect to port %d...", port)
				_, err := http.DefaultClient.Get(fmt.Sprintf("http://127.0.0.1:%d/", port))
				return err != nil
			},
		},
		{args: []string{"engines", "acp", "ory", "roles", "list", "exact"}},
		{args: []string{"engines", "acp", "ory", "roles", "import", "--endpoint", ep, "exact", "../tests/stubs/roles.json"}},
		{args: []string{"engines", "acp", "ory", "roles", "get", "--endpoint", ep, "exact", "role-1"}},
		{args: []string{"engines", "acp", "ory", "roles", "delete", "--endpoint", ep, "exact", "role-1"}},

		{args: []string{"engines", "acp", "ory", "policies", "list", "--endpoint", ep, "exact"}},
		{args: []string{"engines", "acp", "ory", "policies", "import", "--endpoint", ep, "exact", "../tests/stubs/policies.json"}},
		{args: []string{"engines", "acp", "ory", "policies", "get", "--endpoint", ep, "exact", "policy-1"}},
		{args: []string{"engines", "acp", "ory", "policies", "delete", "--endpoint", ep, "exact", "policy-1"}},

		{args: []string{"engines", "acp", "ory", "allowed", "--endpoint", ep, "exact", "peter-1", "resources-11", "actions-11"}},

		{args: []string{"help", "migrate", "sql"}},
		{args: []string{"version"}},
	} {
		RootCmd.SetArgs(c.args)

		t.Run(fmt.Sprintf("command=%v", c.args), func(t *testing.T) {
			if c.wait != nil {
				go func() {
					assert.Nil(t, RootCmd.Execute())
				}()
			}

			if c.wait != nil {
				var count = 0
				for c.wait(t) {
					t.Logf("Port not open yet, retrying attempt #%d...", count)
					count++
					if count > 30 {
						t.FailNow()
					}
					time.Sleep(time.Second)
				}
			} else {
				err := RootCmd.Execute()
				if c.expectErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			}
		})
	}
}
