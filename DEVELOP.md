# Development

This document explains how to develop Ory Keto, run tests, and work with the tooling around it.

## Upgrading and changelog

New releases might introduce breaking changes. To help you identify and incorporate those changes, we document these changes in [UPGRADE.md](./UPGRADE.md) and [CHANGELOG.md](./CHANGELOG.md).

## Command line documentation

To see available commands and flags, run:

```shell
keto -h
# or
keto help
```

## Contribution guidelines

We encourage all contributions. Before opening a pull request, read the [contribution guidelines](./CONTRIBUTING.md).

## Prerequisites

You need Go 1.19+ and, for the test suites:

- Docker and Docker Compose
- GNU Make 4.3
- Node.js and npm >= v7

It is possible to develop Ory Keto on Windows, but please be aware that all guides assume a Unix shell like bash or zsh.

## Install from source

To install Keto from source:

```shell
make install
```

## Formatting code

Format all code using:

```shell
make format
```

The continuous integration pipeline checks code formatting.

## Running tests

There are two types of tests:

- Short tests that do not require a SQL database like PostgreSQL
- Regular tests that require PostgreSQL, MySQL, and CockroachDB

### Short tests

Short tests run fairly quickly and use SQLite in-memory.

Run all short tests:

```shell
go test -short -tags sqlite ./...
```

Run short tests in a specific module:

```shell
go test -tags sqlite -short ./internal/check/...
```

### Regular tests

Regular tests require a database setup.

The test suite can work with docker directly using [ory/dockertest](https://github.com/ory/dockertest), but we encourage using the script instead. Using dockertest can bloat the number of Docker Images on your system and starting them on each run is quite slow.

Run the full test suite:

```shell
source ./scripts/test-resetdb.sh
go test -tags sqlite ./...
```

### End-to-end tests

The e2e tests are part of the normal `go test`. To only run the e2e test, use:

```shell
source ./scripts/test-resetdb.sh
go test -tags sqlite ./internal/e2e/...
```

or add the `-short` tag to only test against sqlite in-memory.

## Build Docker image

To build a development Docker image:

```shell
make docker
```
