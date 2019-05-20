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

	"github.com/ory/keto/engine"
	"github.com/ory/keto/engine/ladon"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:  "run <flavor>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		workers := flagx.MustGetInt(cmd, "workers")
		policies := flagx.MustGetInt(cmd, "policies")
		runFor := flagx.MustGetInt(cmd, "run-for")

		run(workers, policies, runFor, args[0])
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().Int("workers", 25, "the number of concurrent policy creation and allowed check workers")
	runCmd.Flags().Int("policies", 5000, "the number of policies to create")
	runCmd.Flags().Int("run-for", 10, "how long the benchmark should run in seconds")
}

func run(workers, policies, runFor int, flavor string) {
	var wg sync.WaitGroup

	l.Infof("Creating %d policies...", policies)
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(w int) {
			defer wg.Done()
			for i := 0; i < policies/workers; i++ {
				createPolicy(strconv.Itoa(i)+"-"+strconv.Itoa(w), flavor)
			}
		}(i)
	}

	wg.Wait()

	l.Infof("Created %d policies", policies)

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
				total += check(flavor)
				time.Sleep(time.Millisecond)
			}

			l.Infof("Took %.8fs to on average for worker %d", total.Seconds()/float64(count), w)
		}(i)
	}
	wg.Wait()
	l.Infof("Done")
}

func check(flavor string) time.Duration {
	randInt := strconv.Itoa(rand.Int())

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(&ladon.Input{
		Subject:  "tenant:" + randInt + ":user:" + randInt,
		Action:   "check",
		Resource: "tenant:" + randInt + ":thing0:" + randInt,
	})
	cmdx.Must(err, "%v", err)

	req, err := http.NewRequest("POST", "http://localhost:4466/engines/acp/ory/"+flavor+"/allowed", &b)
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

	var decision engine.AuthorizationResult
	cmdx.Must(json.NewDecoder(res.Body).Decode(&decision), "")

	return end.Sub(start)
}

func createPolicy(id string, flavor string) {
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(&ladon.Policy{
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
	})
	cmdx.Must(err, "%v", err)

	req, err := http.NewRequest("PUT", "http://localhost:4466/engines/acp/ory/"+flavor+"/policies", &b)
	cmdx.Must(err, "%v", err)
	req.Header.Set("Content-Type", "application/json")

	start := time.Now()
	res, err := http.DefaultClient.Do(req)
	end := time.Now()

	cmdx.Must(err, "%v", err)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic(fmt.Sprintf("not 200 but %d", res.StatusCode))
	}

	l.Debugf("Took %.8fms to create policy with id: %s", end.Sub(start).Seconds(), id)
}
