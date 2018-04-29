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

// oauth2Cmd represents the oauth2 command
var oauth2Cmd = &cobra.Command{
	Use:   "oauth2",
	Short: "Checks if an OAuth 2.0 Access Token is authorized to perform a certain request",
	Run:   cmdHandler.Warden.IsOAuth2AccessTokenAuthorized,
}

func init() {
	authorizeCmd.AddCommand(oauth2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// oauth2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// oauth2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	oauth2Cmd.Flags().String("token", "", "The request's bearer token")
	oauth2Cmd.Flags().StringArray("scope", []string{}, "The request's required scope")
	oauth2Cmd.Flags().String("action", "", "The request's action")
	oauth2Cmd.Flags().String("resource", "", "The request's resource")
}
