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
)

// subjectCmd represents the subject command
var subjectCmd = &cobra.Command{
	Use:   "subject",
	Short: "Checks if a subject is authorized to perform a certain request",
	Run:   cmdHandler.Warden.IsSubjectAuthorized,
}

func init() {
	authorizeCmd.AddCommand(subjectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subjectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	subjectCmd.Flags().StringP("subject", "s", "", "The request's subject")
	subjectCmd.Flags().StringP("action", "a", "", "The request's action")
	subjectCmd.Flags().StringP("resource", "r", "", "The request's resource")
}
