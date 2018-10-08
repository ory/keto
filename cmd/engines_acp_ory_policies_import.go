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
	"github.com/ory/keto/cmd/client"
	"github.com/ory/urlx"
	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"
)

// enginesAcpOryPoliciesImportCmd represents the import command
var enginesAcpOryPoliciesImportCmd = &cobra.Command{
	Use:   "import <flavor> <file.json> [<file-2.json>, [<file-3.json>, [...]]",
	Short: "Import an ORY Access Control Policy",
	Long: `This command imports one or more json files into the ORY Access Control Policy store.

The json file(s) have to be formatted as arrays:

[
	{"id": "1", "subjects": [...], ...},
	{"id": "2", "subjects": [...], ...},
]`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdx.MinArgs(cmd, args, 2)
		client.CheckLadonFlavor(args[0])
		client.Import(
			"PUT",
			urlx.JoinURLStrings(
				client.LadonEndpointURL(cmd, args[0]),
				"policies",
			),
			args[1:],
		)
	},
}

func init() {
	enginesAcpOryPoliciesCmd.AddCommand(enginesAcpOryPoliciesImportCmd)
}
