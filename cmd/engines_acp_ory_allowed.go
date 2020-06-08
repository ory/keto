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

	"github.com/ory/keto/internal/httpclient/client/engines"
	"github.com/ory/keto/internal/httpclient/models"

	"github.com/spf13/cobra"

	"github.com/ory/x/cmdx"

	"github.com/ory/keto/cmd/client"
)

// enginesAcpOryAllowedCmd represents the roles command
var enginesAcpOryAllowedCmd = &cobra.Command{
	Use:   "allowed <flavor> <subject> <resource> <action>",
	Short: "Check if a request should be allowed or not",
	Run: func(cmd *cobra.Command, args []string) {
		cmdx.MinArgs(cmd, args, 4)
		client.CheckLadonFlavor(args[0])

		c := client.NewClient(cmd)
		res, err := c.Engines.DoOryAccessControlPoliciesAllow(
			engines.NewDoOryAccessControlPoliciesAllowParams().
				WithFlavor(args[0]).
				WithBody(&models.OryAccessControlPolicyAllowedInput{
					Subject:  args[1],
					Resource: args[2],
					Action:   args[3],
				}),
		)
		if err != nil {
			switch d := err.(type) {
			case *engines.DoOryAccessControlPoliciesAllowForbidden:
				fmt.Println(cmdx.FormatResponse(&d.Payload))
				return
			default:
				cmdx.Must(err, "Unable to call ORY Access Control Policy allowed endpoint: %s")
			}
		}
		fmt.Println(cmdx.FormatResponse(&res.Payload))
	},
}

func init() {
	enginesAcpOryCmd.AddCommand(enginesAcpOryAllowedCmd)
}
