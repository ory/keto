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

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/engine/ladon"
	"github.com/ory/x/cmdx"
	"github.com/ory/x/urlx"
)

// enginesAcpOryPoliciesListCmd represents the list command
var enginesAcpOryPoliciesListCmd = &cobra.Command{
	Use:   "list <flavor>",
	Short: "List ORY Access Control Policies",
	Run: func(cmd *cobra.Command, args []string) {
		var proto ladon.Policies
		cmdx.MinArgs(cmd, args, 1)
		client.CheckLadonFlavor(args[0])
		client.Get(
			urlx.MustJoin(
				client.LadonEndpointURL(cmd, args[0]),
				"policies",
			),
			&proto,
		)
	},
}

func init() {
	enginesAcpOryPoliciesCmd.AddCommand(enginesAcpOryPoliciesListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
