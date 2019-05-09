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
	var wg sync.WaitGroup
	const workers = 25

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(w int) {
			defer wg.Done()
			for i := 0; i < 30000/workers; i++ {
				createPolicy(strconv.Itoa(i)+strconv.Itoa(w))
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("Policies created.")

	start := time.Now()

	l.Infof("Took %.8fs to make cache hot", check().Seconds())

	wg.Add(workers)
	// Create some workers to simulate humans
	for i := 0; i < workers; i++ {
		go func(w int) {
			defer wg.Done()
			var count int64
			var total time.Duration

			for time.Now().Sub(start).Seconds() < 10 {
				count++
				total += check()
				time.Sleep(time.Millisecond)
			}

			l.Infof("Took %.8fs to on average for worker %d", total.Seconds()/float64(count), w)
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

	req, err := http.NewRequest("POST", "http://localhost:4466/engines/acp/ory/exact/allowed", &b)
	cmdx.Panic(err)
	req.Header.Set("Content-Type", "application/json")

	start := time.Now()
	res, err := http.DefaultClient.Do(req)
	end := time.Now()

	if err != nil {
		return end.Sub(start)
	}

	cmdx.Panic(err)
	defer res.Body.Close()
	if res.StatusCode != 200 && res.StatusCode != 403 {
		panic(fmt.Sprintf("not 200 / 403 but %d", res.StatusCode))
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

	req, err := http.NewRequest("PUT", "http://localhost:4466/engines/acp/ory/exact/policies", &b)
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
