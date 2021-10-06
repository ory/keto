---
id: milestones
title: Milestones and Roadmap
---

## [v0.7.0-alpha.0](https://github.com/ory/keto/milestone/5)

The biggest change for the next release will be the new SQL table structure
(https://github.com/ory/keto/pull/638). The main goal is to improve the QoS, big
features are planned for the next release.

### [Bug](https://github.com/ory/keto/labels/bug)

Something is not working.

#### Issues

- [x] Check valid relation-tuple characters on insert
      ([keto#661](https://github.com/ory/keto/issues/661))
- [x] Keto version API does not work in REST API and CLI.
      ([keto#696](https://github.com/ory/keto/issues/696)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Keto patch relation tuple REST API panic rather than return bad request
      for wrong input. ([keto#706](https://github.com/ory/keto/issues/706)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Change REST API to not work with encoded subjects
      ([keto#708](https://github.com/ory/keto/issues/708))
- [x] Config schema: replace `ory://*` references with something actually
      resolvable ([keto#719](https://github.com/ory/keto/issues/719)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Docker files and docker-compose.yml do not work on master
      ([keto#725](https://github.com/ory/keto/issues/725)) -
      [@Patrik](https://github.com/zepatrik)

### [Feat](https://github.com/ory/keto/labels/feat)

New feature or request.

#### Issues

- [x] Ensure telemetry is running for GRPC
      ([keto#298](https://github.com/ory/keto/issues/298)) -
      [@hackerman](https://github.com/aeneasr),
      [@Patrik](https://github.com/zepatrik),
      [@Robin Brämer](https://github.com/robinbraemer)
- [x] Single table migration: run as one big transaction
      ([keto#734](https://github.com/ory/keto/issues/734)) -
      [@Patrik](https://github.com/zepatrik)

### [Docs](https://github.com/ory/keto/labels/docs)

Affects documentation.

#### Issues

- [x] Document and improve go gRPC client import
      ([keto#635](https://github.com/ory/keto/issues/635))
- [x] Guide for data migration to new single table schema
      ([keto#733](https://github.com/ory/keto/issues/733)) -
      [@Patrik](https://github.com/zepatrik)

### [Ci](https://github.com/ory/keto/labels/ci)

Affects Continuous Integration (CI).

#### Issues

- [x] Document and improve go gRPC client import
      ([keto#635](https://github.com/ory/keto/issues/635))
- [x] Keto version API does not work in REST API and CLI.
      ([keto#696](https://github.com/ory/keto/issues/696)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Config schema: replace `ory://*` references with something actually
      resolvable ([keto#719](https://github.com/ory/keto/issues/719)) -
      [@Patrik](https://github.com/zepatrik)

### [Blocking](https://github.com/ory/keto/labels/blocking)

Blocks milestones or other issues or pulls.

#### Issues

- [x] Ensure telemetry is running for GRPC
      ([keto#298](https://github.com/ory/keto/issues/298)) -
      [@hackerman](https://github.com/aeneasr),
      [@Patrik](https://github.com/zepatrik),
      [@Robin Brämer](https://github.com/robinbraemer)
- [x] Config schema: replace `ory://*` references with something actually
      resolvable ([keto#719](https://github.com/ory/keto/issues/719)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Docker files and docker-compose.yml do not work on master
      ([keto#725](https://github.com/ory/keto/issues/725)) -
      [@Patrik](https://github.com/zepatrik)
