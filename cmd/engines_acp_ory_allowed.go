// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/sdk/go/keto/swagger"
	"github.com/ory/x/cmdx"
)

// enginesAcpOryAllowedCmd represents the roles command
var enginesAcpOryAllowedCmd = &cobra.Command{
	Use:   "allowed <flavor> <subject> <resource> <action>",
	Short: "Check if a request should be allowed or not",
	Run: func(cmd *cobra.Command, args []string) {
		cmdx.MinArgs(cmd, args, 4)
		client.CheckLadonFlavor(args[0])

		c := swagger.NewEnginesApiWithBasePath(client.EndpointURL(cmd))
		a, res, err := c.DoOryAccessControlPoliciesAllow(args[0], swagger.OryAccessControlPolicyAllowedInput{
			Subject:  args[1],
			Resource: args[2],
			Action:   args[3],
		})
		cmdx.Must(err, "Command failed because error occurred: %s", err)

		if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusForbidden {
			cmdx.Fatalf("Expected status code %d or %d but got: %d", http.StatusOK, http.StatusForbidden, res.StatusCode)
		}

		fmt.Println(cmdx.FormatResponse(&a))
	},
}

func init() {
	enginesAcpOryCmd.AddCommand(enginesAcpOryAllowedCmd)
}
