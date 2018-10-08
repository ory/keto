// Copyright © 2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
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

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ory/x/urlx"

	"github.com/ory/go-convenience/stringslice"
	"github.com/ory/keto/engine/ladon"
	"github.com/ory/x/cmdx"
	"github.com/ory/x/flagx"
	"github.com/spf13/cobra"
)

var client = http.DefaultClient

func Import(method, location string, files []string) {
	for _, file := range files {
		var data []interface{}
		fmt.Printf("Importing file %s to %s...\n", file, location)
		b, err := ioutil.ReadFile(file)
		cmdx.Must(err, "%s", err)

		err = json.Unmarshal(b, &data)
		cmdx.Must(err, "%s", err)

		for _, d := range data {
			var b bytes.Buffer
			err := json.NewEncoder(&b).Encode(d)
			cmdx.Must(err, "%s", err)

			req, err := http.NewRequest(method, location, &b)
			cmdx.Must(err, "%s", err)

			res, err := client.Do(req)
			cmdx.CheckResponse(err, http.StatusOK, res.StatusCode)

			res.Body.Close()
			fmt.Printf("Data from file %s successfully imported!\n", file)
		}
	}
}

func Get(location string, proto interface{}) {
	res, err := client.Get(location)
	cmdx.CheckResponse(err, http.StatusOK, res.StatusCode)
	defer res.Body.Close()

	d := json.NewDecoder(res.Body)
	d.DisallowUnknownFields()

	err = d.Decode(proto)
	cmdx.Must(err, "%s", err)
	cmdx.FormatResponse(proto)
}

func Delete(location string) {
	req, err := http.NewRequest("DELETE", location, nil)
	cmdx.Must(err, "%s", err)

	res, err := client.Do(req)
	cmdx.CheckResponse(err, http.StatusNoContent, res.StatusCode)
	res.Body.Close()
	fmt.Printf("Resource at location %s was deleted successfully!", location)
}

func CheckLadonFlavor(flavor string) {
	if !stringslice.Has(ladon.EnabledFlavors, flavor) {
		cmdx.Fatalf("Flavor %s is not supported, please choose one of: %v", flavor, ladon.EnabledFlavors)
	}
}

func EndpointURL(cmd *cobra.Command) string {
	e := flagx.MustGetString(cmd, "endpoint")
	if e == "" {
		fmt.Println(cmd.UsageString())
		cmdx.Fatalf("Please set the location of the ORY Keto server by using the --endpoint flag or the KETO_URL environment variable.")
	}
	return e
}

func LadonEndpointURL(cmd *cobra.Command, flavor string) string {
	return urlx.MustJoin(
		flagx.MustGetString(cmd, "endpoint"),
		fmt.Sprintf(strings.Replace(ladon.BasePath, ":flavor", "%s", 1), flavor),
	)
}
