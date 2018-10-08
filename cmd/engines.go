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
	"os"

	"github.com/spf13/cobra"
)

// enginesCmd represents the engines command
var enginesCmd = &cobra.Command{
	Use:   "engines",
	Short: "Manage access control engines",
}

func init() {
	RootCmd.AddCommand(enginesCmd)
	enginesCmd.PersistentFlags().String("endpoint", os.Getenv("KETO_URL"), "URL of the ORY Keto server - defaults to environment variable KETO_URL")
}
