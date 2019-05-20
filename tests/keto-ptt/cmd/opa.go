// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/flagx"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/ory/keto/engine/ladon"
)

// opaCmd represents the opa command
var opaCmd = &cobra.Command{
	Use:  "opa <flavor>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		workers := flagx.MustGetInt(cmd, "workers")
		policies := flagx.MustGetInt(cmd, "policies")
		runFor := flagx.MustGetInt(cmd, "run-for")

		l.Infof("Creating %d policies...", policies)
		createOPAPolicies(policies, args[0])
		l.Infof("Created %d policies", policies)


		var wg sync.WaitGroup
		wg.Add(workers)
		start := time.Now()
		// Create some workers to simulate humans
		for i := 0; i < workers; i++ {
			go func(w int) {
				defer wg.Done()
				var count int64
				var total time.Duration

				for time.Now().Sub(start).Seconds() < float64(runFor) {
					count++
					total += checkOPA(args[0])
				}

				l.Infof("Took %.8fs to on average for worker %d", total.Seconds()/float64(count), w)
			}(i)
		}
		wg.Wait()
		l.Infof("Done")
	},
}

func init() {
	rootCmd.AddCommand(opaCmd)

	opaCmd.Flags().Int("workers", 1, "the number of concurrent policy creation and allowed check workers")
	opaCmd.Flags().Int("policies", 30000, "the number of policies to create")
	opaCmd.Flags().Int("run-for", 10, "how long the benchmark should run in seconds")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// opaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// opaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkOPA(flavor string) time.Duration {
	randInt := strconv.Itoa(rand.Int())

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(&struct {
		Input *ladon.Input `json:"input"`
	}{
		Input: &ladon.Input{
			Subject:  "tenant:" + randInt + ":user:" + randInt,
			Action:   "check",
			Resource: "tenant:" + randInt + ":thing0:" + randInt,
		},
	})
	cmdx.Must(err, "%v", err)

	req, err := http.NewRequest("POST", "http://localhost:8181/v1/data/ory/"+flavor+"/allow", &b)
	cmdx.Must(err, "%v", err)
	req.Header.Set("Content-Type", "application/json")

	start := time.Now()
	res, err := http.DefaultClient.Do(req)
	end := time.Now()

	if err != nil {
		return end.Sub(start)
	}

	cmdx.Must(err, "%v", err)
	defer res.Body.Close()
	if res.StatusCode != 200 && res.StatusCode != 403 {
		err := errors.Errorf("expected status code 200 or 403 but got: %d", res.StatusCode)
		cmdx.Must(err, "%s", err)
	}

	var decision struct{ Result bool `json:"result"` }
	cmdx.Must(json.NewDecoder(res.Body).Decode(&decision), "")

	return end.Sub(start)
}
func createOPAPolicies(amount int, flavor string) {

	// (cd tests/opa; curl -X PUT --data-binary @data.json -H 'Content-Type: application/json' localhost:8181/v1/data/store/ory/exact)
	// (cd tests/opa; curl -X PUT --data-binary @data.json -H 'Content-Type: application/json' localhost:8181/v1/data/store/ory/regex)
	// (cd tests/opa; curl -X PUT --data-binary @data.json -H 'Content-Type: application/json' localhost:8181/v1/data/store/ory/glob)

	req, err := http.NewRequest("DELETE", "http://localhost:8181/v1/data/store", nil)
	cmdx.Must(err, "%v", err)
	res, err := http.DefaultClient.Do(req)
	cmdx.Must(err, "%v", err)
	defer res.Body.Close()

	policies := make([]*ladon.Policy, amount)
	for i := 0; i < amount; i++ {
		id := strconv.Itoa(i)
		policies[i] = &ladon.Policy{
			ID:       id,
			Subjects: []string{"tenant:<.*>:user:<.*>"},
			Resources: []string{
				"tenant:<.*>:thing" + id + ":<.*>",
				"tenant:<.*>:foo" + id + ":<.*>",
				"tenant:<.*>:bar" + id + ":<.*>",
				"tenant:<.*>:baz" + id + ":<.*>",
				"tenant:<.*>:boo" + id + ":<.*>",
				"tenant:<.*>:bam" + id + ":<.*>",
				"tenant:<.*>:bag" + id + ":<.*>",
				"tenant:<.*>:bad" + id + ":<.*>",
			},
			Actions: []string{"check"},
			Effect:  "allow",
		}
	}

	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(&struct {
		Policies []*ladon.Policy `json:"policies"`
		Roles    []string        `json:"roles"`
	}{
		Policies: policies,
		Roles:    []string{},
	})
	cmdx.Must(err, "%v", err)

	req, err = http.NewRequest("PUT", "http://localhost:8181/v1/data/store/ory/"+flavor, &b)
	cmdx.Must(err, "%v", err)
	req.Header.Set("Content-Type", "application/json")

	start := time.Now()
	res, err = http.DefaultClient.Do(req)
	end := time.Now()

	cmdx.Must(err, "%v", err)
	defer res.Body.Close()
	if res.StatusCode != 200 && res.StatusCode != 204 {
		panic(fmt.Sprintf("not 200 but %d", res.StatusCode))
	}

	l.Debugf("Took %.8fms to create %d policies", end.Sub(start).Seconds(), policies)
}
