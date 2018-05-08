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
	"github.com/ory/keto/cmd/server"
	"github.com/spf13/cobra"
)

// migrateHydraCmd represents the hydra command
var migrateHydraCmd = &cobra.Command{
	Use:   "hydra <database-url>",
	Short: "Applies SQL migration plans that migrate groups and policies from ORY Hydra < v1.0.0",
	Long: `It is recommended to run this command close to the SQL instance (e.g. same subnet) instead of over the public internet.
This decreases risk of failure and decreases time required.

### WARNING ###

Before running this command on an existing database, create a back up!
`,
	Run: server.RunMigrateHydra(logger),
}

func init() {
	migrateCmd.AddCommand(migrateHydraCmd)

	migrateHydraCmd.Flags().Bool("read-from-env", false, "Instead of reading the database URL from the command line arguments, the value of environment variable DATABASE_URL will be used.")
}
