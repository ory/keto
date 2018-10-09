# sqlcon

[![CircleCI](https://circleci.com/gh/ory/sqlcon.svg?style=shield)](https://circleci.com/gh/ory/sqlcon)

## Dockertest

This library also helps with setting up Dockertest for SQL databases:

```go
func TestMain(m *testing.M) {
    // Registers a docker pool and sets up a list of resources.
	runner := dockertest.Register()

    // We must parse the flags in order to check if `testing.Short()` is true.
	flag.Parse()

	// Connects to the databases in parallel.
    dockertest.Parallel([]func(){
        connectToPostgres,
        connectToMySQL,
    })

    // Kills and removes the docker containers.
	runner.Exit(m.Run())
}

func connectToMySQL() {
	db, err := dockertest.ConnectToTestMySQL()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

    // ...
}

func connectToPostgres() {
	db, err := dockertest.ConnectToTestPostgres()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

    // ...
}
```
