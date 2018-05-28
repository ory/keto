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

package client

import (
	"encoding/json"
	"fmt"
	"os"

	"strings"

	keto "github.com/ory/keto/sdk/go/keto/swagger"
	"github.com/spf13/cobra"
)

func getBasePath(cmd *cobra.Command) string {
	location, err := cmd.Flags().GetString("endpoint")
	if err != nil || location == "" {
		fmt.Println(cmd.UsageString())
		fatalf("Please set the location of ORY Keto by using the --endpoint flag or the KETO_URL environment variable.")
	}
	return strings.TrimRight(location, "/")
}

func must(err error, message string, args ...interface{}) {
	if err == nil {
		return
	}

	fmt.Fprintf(os.Stderr, message+"\n", args...)
	os.Exit(1)
}

func checkResponse(response *keto.APIResponse, err error, expectedStatusCode int) {
	must(err, "Command failed because error \"%s\" occurred.\n", err)

	if response.StatusCode != expectedStatusCode {
		fmt.Fprintf(os.Stderr, "Command failed because status code %d was expeceted but code %d was received.\n", expectedStatusCode, response.StatusCode)
		os.Exit(1)
		return
	}
}

func formatResponse(response interface{}) string {
	out, err := json.MarshalIndent(response, "", "\t")
	must(err, `Command failed because an error ("%s") occurred while prettifying output.`, err)
	return string(out)
}

func fatalf(message string, args ...interface{}) {
	if len(args) > 0 {
		fmt.Printf(message+"\n", args)
	} else {
		fmt.Println(message)
	}
	os.Exit(1)
}
