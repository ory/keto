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
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/engine"
	"github.com/ory/keto/engine/ladon"
	"github.com/ory/x/cmdx"
	"github.com/ory/x/urlx"
)

// enginesAcpOryAllowedCmd represents the roles command
var enginesAcpOryAllowedCmd = &cobra.Command{
	Use:   "allowed <flavor> <subject> <resource> <action>",
	Short: "Check if a request should be allowed or not",
	Run: func(cmd *cobra.Command, args []string) {
		cmdx.MinArgs(cmd, args, 4)
		client.CheckLadonFlavor(args[0])

		var b bytes.Buffer
		err := json.NewEncoder(&b).Encode(&ladon.Input{
			Subject:  args[1],
			Resource: args[2],
			Action:   args[3],
		})
		cmdx.Must(err, "Unable to encode input data to json: %s", err)

		res, err := http.DefaultClient.Post(
			urlx.MustJoin(
				client.LadonEndpointURL(cmd, args[0]),
				"allowed",
			),
			"application/json",
			&b,
		)
		cmdx.CheckResponse(err, http.StatusOK, res)
		defer res.Body.Close()

		var result engine.AuthorizationResult
		d := json.NewDecoder(res.Body)
		d.DisallowUnknownFields()
		err = d.Decode(&result)
		cmdx.Must(err, "Unable to decode data to json: %s", err)
		fmt.Println(cmdx.FormatResponse(&result))
	},
}

func init() {
	enginesAcpOryCmd.AddCommand(enginesAcpOryAllowedCmd)
}
