// Copyright © 2017 Aeneas Rekkas <aeneas+oss@aeneas.io>
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
)

// policiesImportCmd  represents the import command
var policiesImportCmd = &cobra.Command{
	Use:   "import <path/to/file.json> [<path/to/other/file.json>...]",
	Short: "Import policies from JSON files",
	Run:   cmdHandler.Policies.ImportPolicy,
}

func init() {
	policiesCmd.AddCommand(policiesImportCmd)

}
