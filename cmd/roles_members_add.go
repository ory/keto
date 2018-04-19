/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @Copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 *
 */

package cmd

import (
	"github.com/spf13/cobra"
)

var rolesMembersAdd = &cobra.Command{
	Use:   "add <role> <member> [<member>...]",
	Short: "Add members to a role",
	Long: `This command adds members to a role.

Example:
  keto roles members add my-group peter julia
`,
	Run: cmdHandler.Roles.RoleAddMembers,
}

func init() {
	rolesMembersCmd.AddCommand(rolesMembersAdd)
}
