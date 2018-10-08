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
	"strings"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/engine/ladon"
	"github.com/ory/urlx"
	"github.com/ory/x/cmdx"
	"github.com/ory/x/flagx"
	"github.com/spf13/cobra"
)

// enginesAcpOryPoliciesDeleteCmd represents the delete command
var enginesAcpOryPoliciesDeleteCmd = &cobra.Command{
	Use:   "delete <flavor> <id> [<id-2>, [<...>]]",
	Short: "Delete an ORY Access Control Policy",
	Run: func(cmd *cobra.Command, args []string) {
		cmdx.MinArgs(cmd, args, 2)
		client.CheckLadonFlavor(args[0])
		for _, id := range args[1:] {
			client.Delete(urlx.JoinURLStrings(
				flagx.MustGetString(cmd, "endpoint"),
				fmt.Sprintf(strings.Replace(ladon.BasePath, ":flavor", "%s", 1), args[0]),
				"policies",
				id,
			))
		}
	},
}

func init() {
	enginesAcpOryPoliciesCmd.AddCommand(enginesAcpOryPoliciesDeleteCmd)
}
