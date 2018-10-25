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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ory/go-convenience/stringslice"
	"github.com/ory/keto/engine/ladon"
	"github.com/ory/x/cmdx"
	"github.com/ory/x/flagx"
	"github.com/ory/x/urlx"
)

var client = http.DefaultClient

func ImportFile(file string, proto interface{}, f func()) {
		b, err := ioutil.ReadFile(filepath.Clean(file))
		cmdx.Must(err, "Unable to read file %s: %s", file, err)

		err = json.Unmarshal(b, proto)
		cmdx.Must(err, "Unable to decode file %s to json: %s", file, err)
		f()
}

func Get(location string, proto interface{}) {
	res, err := client.Get(location)
	cmdx.CheckResponse(err, http.StatusOK, res)
	defer res.Body.Close()

	d := json.NewDecoder(res.Body)
	d.DisallowUnknownFields()

	err = d.Decode(proto)
	cmdx.Must(err, "Unable to decode data to json: %s", err)
	fmt.Println(cmdx.FormatResponse(proto))
}

func Delete(location string) {
	req, err := http.NewRequest("DELETE", location, nil)
	cmdx.Must(err, "Unable to initialize HTTP request: %s", err)

	res, err := client.Do(req)
	cmdx.CheckResponse(err, http.StatusNoContent, res)
	err = res.Body.Close()
	cmdx.Must(err, "Unable to close body: %s", err)
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
