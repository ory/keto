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
	"github.com/ory/keto/sdk/go/keto/client/engines"
	"github.com/ory/keto/sdk/go/keto/models"

	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/x/cmdx"
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

		c := client.NewClient(cmd)
		for _, file := range args[1:] {
			var p []models.OryAccessControlPolicy
			client.ImportFile(
				file,
				&p,
				func() {
					for _, pp := range p {
						_, err := c.Engines.UpsertOryAccessControlPolicy(engines.NewUpsertOryAccessControlPolicyParams().WithFlavor(args[0]).WithBody(&pp))
						cmdx.Must(err, "Unable to import ORY Access Control Policy: %s", err)
					}
				},
			)
		}
	},
}

func init() {
	enginesAcpOryPoliciesCmd.AddCommand(enginesAcpOryPoliciesImportCmd)
}
