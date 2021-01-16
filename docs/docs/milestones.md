---
id: milestones
title: Milestones and Roadmap
---

## [Next Gen Keto - first working version](https://github.com/ory/keto/milestone/4)

Goals:

- check/expand/read/write APIs
- SQL persistence (only local database)
- operation using one node in one data center
- namespace configuration including subject set rewrites

Non-goals:

- watch API
- caching
- fan-out
- Leopard indexing system
- fancy query features

### [Bug](https://github.com/ory/keto/labels/bug)

Something is not working.

#### Issues

- [ ] CLI remote flag should be required
      ([keto#287](https://github.com/ory/keto/issues/287)) -
      [@Patrik](https://github.com/zepatrik)
- [ ] REST Relations API returns null instead of `[]`
      ([keto#289](https://github.com/ory/keto/issues/289)) -
      [@Patrik](https://github.com/zepatrik)
- [ ] REST API create relation should mirror payload in 201 OK response
      ([keto#290](https://github.com/ory/keto/issues/290)) -
      [@Patrik](https://github.com/zepatrik)
- [ ] Relations are not unique
      ([keto#292](https://github.com/ory/keto/issues/292)) -
      [@Patrik](https://github.com/zepatrik)
- [ ] REST API create and subsequent get relation does not properly persist
      fields ([keto#291](https://github.com/ory/keto/issues/291)) -
      [@Patrik](https://github.com/zepatrik)
- [ ] Unable to create relations using REST API and string notation
      ([keto#293](https://github.com/ory/keto/issues/293)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Replace in-memory persister with SQLite schema
      ([keto#294](https://github.com/ory/keto/issues/294)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Write GRPC handler and tests for check engine
      ([keto#296](https://github.com/ory/keto/issues/296)) -
      [@Patrik](https://github.com/zepatrik),
      [@Robin Brämer](https://github.com/robinbraemer)

### [Feat](https://github.com/ory/keto/labels/feat)

New feature or request.

#### Issues

- [ ] [next-gen] Allow defining userset rewrites
      ([keto#263](https://github.com/ory/keto/issues/263)) -
      [@Patrik](https://github.com/zepatrik)
- [ ] Ensure telemetry is running for both GRPC and HTTP
      ([keto#298](https://github.com/ory/keto/issues/298)) -
      [@hackerman](https://github.com/aeneasr),
      [@Patrik](https://github.com/zepatrik),
      [@Robin Brämer](https://github.com/robinbraemer)
- [ ] Namespace configuration
      ([keto#303](https://github.com/ory/keto/issues/303)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Write relationtuple tests for http and grpc handlers
      ([keto#297](https://github.com/ory/keto/issues/297)) -
      [@Patrik](https://github.com/zepatrik),
      [@Robin Brämer](https://github.com/robinbraemer)
- [x] Define and architect SQL schema and queries for querying relations
      ([keto#300](https://github.com/ory/keto/issues/300)) -
      [@Patrik](https://github.com/zepatrik)

### [Blocking](https://github.com/ory/keto/labels/blocking)

Blocks milestones or other issues or pulls.

#### Issues

- [ ] Ensure telemetry is running for both GRPC and HTTP
      ([keto#298](https://github.com/ory/keto/issues/298)) -
      [@hackerman](https://github.com/aeneasr),
      [@Patrik](https://github.com/zepatrik),
      [@Robin Brämer](https://github.com/robinbraemer)
- [ ] Namespace configuration
      ([keto#303](https://github.com/ory/keto/issues/303)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Define and architect SQL schema and queries for querying relations
      ([keto#300](https://github.com/ory/keto/issues/300)) -
      [@Patrik](https://github.com/zepatrik)

#### Pull Requests

- [x] feat: first API draft and generation
      ([keto#315](https://github.com/ory/keto/pull/315)) -
      [@Patrik](https://github.com/zepatrik),
      [@Robin Brämer](https://github.com/robinbraemer)

### [Rfc](https://github.com/ory/keto/labels/rfc)

A request for comments to discuss and share ideas.

#### Issues

- [ ] Consider rename WriteRelationTuples in WriteService
      ([keto#351](https://github.com/ory/keto/issues/351)) -
      [@Patrik](https://github.com/zepatrik),
      [@Robin Brämer](https://github.com/robinbraemer)

#### Pull Requests

- [x] feat: first API draft and generation
      ([keto#315](https://github.com/ory/keto/pull/315)) -
      [@Patrik](https://github.com/zepatrik),
      [@Robin Brämer](https://github.com/robinbraemer)

## [Next Gen Keto - next milestone](https://github.com/ory/keto/milestone/3)

Tracks all the issues contributing to next gen Keto.

### [Bug](https://github.com/ory/keto/labels/bug)

Something is not working.

#### Issues

- [x] Move models package to relations
      ([keto#295](https://github.com/ory/keto/issues/295)) -
      [@Patrik](https://github.com/zepatrik)

### [Feat](https://github.com/ory/keto/labels/feat)

New feature or request.

#### Issues

- [ ] Define naming conventions for objects, relations, namespaces
      ([keto#288](https://github.com/ory/keto/issues/288)) -
      [@Patrik](https://github.com/zepatrik)
- [ ] Write benchmark tests for relationtuple package
      ([keto#301](https://github.com/ory/keto/issues/301)) -
      [@Patrik](https://github.com/zepatrik)
- [ ] Database sharding ([keto#306](https://github.com/ory/keto/issues/306)) -
      [@Patrik](https://github.com/zepatrik)
- [ ] Integrate Next Gen Keto with wider policy ecosystems / Open Policy Agent
      ([keto#318](https://github.com/ory/keto/issues/318)) -
      [@hackerman](https://github.com/aeneasr)
- [ ] Allow defining ABAC relation tuples
      ([keto#319](https://github.com/ory/keto/issues/319))
- [ ] Allow narrowed ACL evaluation in check requests
      ([keto#323](https://github.com/ory/keto/issues/323))
- [ ] Allow modifying relation tuples with consistency guarantees
      ([keto#328](https://github.com/ory/keto/issues/328)) -
      [@Robin Brämer](https://github.com/robinbraemer)
- [ ] Add TTL support to relation tuple
      ([keto#346](https://github.com/ory/keto/issues/346))
- [x] Remove support for AND/OR/XOR in GetRelationTuples
      ([keto#299](https://github.com/ory/keto/issues/299)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Rewrite check engine to search the graph in the other direction
      ([keto#302](https://github.com/ory/keto/issues/302)) -
      [@Patrik](https://github.com/zepatrik)

### [Docs](https://github.com/ory/keto/labels/docs)

Affects documentation.

#### Issues

- [ ] Integrate Next Gen Keto with wider policy ecosystems / Open Policy Agent
      ([keto#318](https://github.com/ory/keto/issues/318)) -
      [@hackerman](https://github.com/aeneasr)

### [Blocking](https://github.com/ory/keto/labels/blocking)

Blocks milestones or other issues or pulls.

#### Issues

- [ ] Protobuf APIs & tooling
      ([keto#311](https://github.com/ory/keto/issues/311)) -
      [@Patrik](https://github.com/zepatrik),
      [@Robin Brämer](https://github.com/robinbraemer)
- [ ] Write benchmark tests for relationtuple package
      ([keto#301](https://github.com/ory/keto/issues/301)) -
      [@Patrik](https://github.com/zepatrik)

#### Pull Requests

- [x] chore: make all go packages internal
      ([keto#313](https://github.com/ory/keto/pull/313))

### [Rfc](https://github.com/ory/keto/labels/rfc)

A request for comments to discuss and share ideas.

#### Issues

- [ ] Distributed cache: loading and coordination
      ([keto#312](https://github.com/ory/keto/issues/312))
- [ ] Integrate Next Gen Keto with wider policy ecosystems / Open Policy Agent
      ([keto#318](https://github.com/ory/keto/issues/318)) -
      [@hackerman](https://github.com/aeneasr)
- [ ] Allow defining ABAC relation tuples
      ([keto#319](https://github.com/ory/keto/issues/319))
- [ ] Allow narrowed ACL evaluation in check requests
      ([keto#323](https://github.com/ory/keto/issues/323))
- [ ] Allow modifying relation tuples with consistency guarantees
      ([keto#328](https://github.com/ory/keto/issues/328)) -
      [@Robin Brämer](https://github.com/robinbraemer)
- [x] Design decisions, Clarifications and Proposals
      ([keto#307](https://github.com/ory/keto/issues/307)) -
      [@hackerman](https://github.com/aeneasr),
      [@Patrik](https://github.com/zepatrik)

## [next](https://github.com/ory/keto/milestone/2)

_This milestone does not have a description._

### [Bug](https://github.com/ory/keto/labels/bug)

Something is not working.

#### Issues

- [ ] Slash (/) in role or policy id causes 404 error for GET and DELETE
      ([keto#140](https://github.com/ory/keto/issues/140))
- [ ] Keto is posting plain text passwords when there is an issue with DSN
      ([keto#237](https://github.com/ory/keto/issues/237))

### [Feat](https://github.com/ory/keto/labels/feat)

New feature or request.

#### Issues

- [ ] Evaluate queries needs to get the entire database in cache to works
      ([keto#187](https://github.com/ory/keto/issues/187))
- [ ] Roles and Policies Filter by using flavor strategy
      ([keto#186](https://github.com/ory/keto/issues/186))
- [x] Add a description attribute to access control policy role
      ([keto#213](https://github.com/ory/keto/issues/213))

### [Rfc](https://github.com/ory/keto/labels/rfc)

A request for comments to discuss and share ideas.

#### Issues

- [ ] IDs being limited to varchar(64) reduces the usefulness of URNs when you
      are using them globally
      ([keto#197](https://github.com/ory/keto/issues/197))
