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
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/httpclient/client/engines"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/x/cmdx"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <flavor> <id> [<id-2>, [<...>]]",
	Short: "Delete an ORY Access Control Policy Role",
	Run: func(cmd *cobra.Command, args []string) {
		cmdx.MinArgs(cmd, args, 2)
		client.CheckLadonFlavor(args[0])

		c := client.NewClient(cmd)
		for _, id := range args[1:] {
			_, err := c.Engines.DeleteOryAccessControlPolicyRole(engines.NewDeleteOryAccessControlPolicyRoleParams().WithFlavor(args[0]).WithID(id))
			cmdx.Must(err, "Unable to delete ORY Access Control Policy Role: %s", err)
		}
	},
}

func init() {
	enginesAcpOryRolesCmd.AddCommand(deleteCmd)
}
