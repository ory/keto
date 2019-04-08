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

	"github.com/ory/keto/sdk/go/keto/client/engines"

	"github.com/spf13/cobra"

	"github.com/ory/x/flagx"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/x/cmdx"
)

// enginesAcpOryPoliciesListCmd represents the list command
var enginesAcpOryPoliciesListCmd = &cobra.Command{
	Use:   "list <flavor>",
	Short: "List ORY Access Control Policies",
	Run: func(cmd *cobra.Command, args []string) {
		cmdx.MinArgs(cmd, args, 1)
		client.CheckLadonFlavor(args[0])

		limit := int64(flagx.MustGetInt(cmd, "limit"))
		offset := int64(flagx.MustGetInt(cmd, "offset"))
		c := client.NewClient(cmd)
		r, err := c.Engines.ListOryAccessControlPolicies(
			engines.NewListOryAccessControlPoliciesParams().WithFlavor(args[0]).WithLimit(&limit).WithOffset(&offset),
		)
		cmdx.Must(err, "Unable to list ORY Access Control Policies: %s")
		fmt.Println(cmdx.FormatResponse(r))
	},
}

func init() {
	enginesAcpOryPoliciesCmd.AddCommand(enginesAcpOryPoliciesListCmd)
	enginesAcpOryPoliciesListCmd.Flags().Int("limit", 100, "Limit the items being fetched")
	enginesAcpOryPoliciesListCmd.Flags().Int("offset", 0, "Set the offset for fetching items")
}
