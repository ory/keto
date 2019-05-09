package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ory/keto/engine"
	"github.com/ory/keto/engine/ladon"
	"github.com/ory/x/cmdx"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var l = logrus.New()

func main() {
	for i := 0; i < 30000; i++ {
		createPolicy(strconv.Itoa(i))
	}
	fmt.Println("Policies created.")

	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(15)
	// Create some workers to simulate humans
	for i := 0; i < 15; i++ {
		go func(w int) {
			defer wg.Done()
			var count int64
			var total time.Duration

			for time.Now().Sub(start).Seconds() < 10 {
				count++
				total += check()
				time.Sleep(time.Second)
			}

			l.Infof("Took %.8fs to on average for worker %d", total.Seconds(), w)
		}(i)
	}
	wg.Wait()
}

func check() time.Duration {
	randInt := strconv.Itoa(rand.Int())

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(&ladon.Input{
		Subject:  "tenant:" + randInt + ":user:" + randInt,
		Action:   "check",
		Resource: "tenant:" + randInt + ":thing0:" + randInt,
	})
	cmdx.Must(err, "%", err)

	req, err := http.NewRequest("POST", "http://localhost:4466/engines/acp/ory/regex/allowed", &b)
	cmdx.Panic(err)
	req.Header.Set("Content-Type", "application/json")

	start := time.Now()
	res, err := http.DefaultClient.Do(req)
	end := time.Now()

	cmdx.Panic(err)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic(fmt.Sprintf("not 200 but %d", res.StatusCode))
	}

	var decision engine.AuthorizationResult
	cmdx.Must(json.NewDecoder(res.Body).Decode(&decision), "")

	return end.Sub(start)
}

func createPolicy(id string) {
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
	cmdx.Panic(err)

	req, err := http.NewRequest("PUT", "http://localhost:4466/engines/acp/ory/regex/policies", &b)
	cmdx.Panic(err)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	cmdx.Panic(err)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic(fmt.Sprintf("not 200 but %d", res.StatusCode))
	}

	//l.Infof("Took %.8fms to create policy with id: %s", end.Sub(start).Seconds(), id)
}
