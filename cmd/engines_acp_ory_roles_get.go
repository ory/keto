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

	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/x/cmdx"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <flavor> [<id-2>, [<...>]]",
	Short: "Get an ORY Access Control Policy",
	Run: func(cmd *cobra.Command, args []string) {
		cmdx.MinArgs(cmd, args, 2)
		client.CheckLadonFlavor(args[0])

		c := client.NewClient(cmd)
		for _, id := range args[1:] {
			r, err := c.Engines.GetOryAccessControlPolicyRole(engines.NewGetOryAccessControlPolicyRoleParams().WithFlavor(args[0]).WithID(id))
			cmdx.Must(err, "Unable to get ORY Access Control Policy Role: %s", err)
			fmt.Println(cmdx.FormatResponse(r.Payload))
		}
	},
}

func init() {
	enginesAcpOryRolesCmd.AddCommand(getCmd)
}
