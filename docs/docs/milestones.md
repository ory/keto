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

- [x] CLI remote flag should be required
      ([keto#287](https://github.com/ory/keto/issues/287)) -
      [@Patrik](https://github.com/zepatrik)
- [x] REST Relations API returns null instead of `[]`
      ([keto#289](https://github.com/ory/keto/issues/289)) -
      [@Patrik](https://github.com/zepatrik)
- [x] REST API create relation should mirror payload in 201 OK response
      ([keto#290](https://github.com/ory/keto/issues/290)) -
      [@Patrik](https://github.com/zepatrik)
- [x] REST API create and subsequent get relation does not properly persist
      fields ([keto#291](https://github.com/ory/keto/issues/291)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Unable to create relations using REST API and string notation
      ([keto#293](https://github.com/ory/keto/issues/293)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Replace in-memory persister with SQLite schema
      ([keto#294](https://github.com/ory/keto/issues/294)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Write GRPC handler and tests for check engine
      ([keto#296](https://github.com/ory/keto/issues/296)) -
      [@Patrik](https://github.com/zepatrik),
      [@Robin Brämer](https://github.com/robinbraemer)
- [x] Remove `config.Provider` interface
      ([keto#403](https://github.com/ory/keto/issues/403)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Abort early in subjectIsAllowed
      ([keto#405](https://github.com/ory/keto/issues/405)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Check REST API should return JSON Object
      ([keto#406](https://github.com/ory/keto/issues/406)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Clarify where handlers are tested
      ([keto#407](https://github.com/ory/keto/issues/407)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Write relationtuple tests
      ([keto#408](https://github.com/ory/keto/issues/408)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Write dedicated persistence tests
      ([keto#409](https://github.com/ory/keto/issues/409)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Persistence errors are not properly handled and wrapped
      ([keto#432](https://github.com/ory/keto/issues/432)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Non-nil return values despite errors
      ([keto#433](https://github.com/ory/keto/issues/433)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Don't use fmt.Sprintf to construct queries
      ([keto#434](https://github.com/ory/keto/issues/434)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Make errors debuggable and understandable
      ([keto#438](https://github.com/ory/keto/issues/438)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Command status --block is not blocking until healthy
      ([keto#456](https://github.com/ory/keto/issues/456)) -
      [@Patrik](https://github.com/zepatrik)

### [Feat](https://github.com/ory/keto/labels/feat)

New feature or request.

#### Issues

- [ ] Ensure telemetry is running for both GRPC and HTTP
      ([keto#298](https://github.com/ory/keto/issues/298)) -
      [@hackerman](https://github.com/aeneasr),
      [@Patrik](https://github.com/zepatrik),
      [@Robin Brämer](https://github.com/robinbraemer)
- [ ] Add tracing and metrics (prometheus/jaeger/...) capabilities
      ([keto#463](https://github.com/ory/keto/issues/463)) -
      [@Andreas Bucksteeg](https://github.com/tricky42),
      [@Piotr Mścichowski](https://github.com/piotrmsc)
- [x] Write relationtuple tests for http and grpc handlers
      ([keto#297](https://github.com/ory/keto/issues/297)) -
      [@Patrik](https://github.com/zepatrik),
      [@Robin Brämer](https://github.com/robinbraemer)
- [x] Define and architect SQL schema and queries for querying relations
      ([keto#300](https://github.com/ory/keto/issues/300)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Move buf configs to project root
      ([keto#399](https://github.com/ory/keto/issues/399)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Follow buf style guide
      ([keto#400](https://github.com/ory/keto/issues/400)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Remove stale protobuf definitions
      ([keto#401](https://github.com/ory/keto/issues/401)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Relation Tuple Migrations should not be hardcoded but use prefixed tables
      ([keto#404](https://github.com/ory/keto/issues/404)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Use `errors.WithStack` where ever non-keto code returns errors.
      ([keto#437](https://github.com/ory/keto/issues/437)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Allow deletion of relation tuples
      ([keto#452](https://github.com/ory/keto/issues/452)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Namespace migrate status command
      ([keto#502](https://github.com/ory/keto/issues/502)) -
      [@Patrik](https://github.com/zepatrik)

### [Docs](https://github.com/ory/keto/labels/docs)

Affects documentation.

#### Issues

- [ ] Next Gen Documentation
      ([keto#420](https://github.com/ory/keto/issues/420)) -
      [@Patrik](https://github.com/zepatrik)
- [ ] Do not use latest tags in docker-compose examples
      ([keto#481](https://github.com/ory/keto/issues/481)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Improve documentation of Expand API's max_depth parameter
      ([keto#419](https://github.com/ory/keto/issues/419)) -
      [@Patrik](https://github.com/zepatrik)

### [Ci](https://github.com/ory/keto/labels/ci)

Affects Continuous Integration (CI).

#### Issues

- [x] Add golangci-lint with gosec linter
      ([keto#435](https://github.com/ory/keto/issues/435)) -
      [@Patrik](https://github.com/zepatrik)

### [Blocking](https://github.com/ory/keto/labels/blocking)

Blocks milestones or other issues or pulls.

#### Issues

- [ ] Ensure telemetry is running for both GRPC and HTTP
      ([keto#298](https://github.com/ory/keto/issues/298)) -
      [@hackerman](https://github.com/aeneasr),
      [@Patrik](https://github.com/zepatrik),
      [@Robin Brämer](https://github.com/robinbraemer)
- [ ] Ensure goreleaser config is still working
      ([keto#410](https://github.com/ory/keto/issues/410)) -
      [@hackerman](https://github.com/aeneasr)
- [ ] Next Gen Documentation
      ([keto#420](https://github.com/ory/keto/issues/420)) -
      [@Patrik](https://github.com/zepatrik)
- [ ] Require relations to be defined in the namespace config
      ([keto#509](https://github.com/ory/keto/issues/509)) -
      [@Patrik](https://github.com/zepatrik)
- [ ] Use fizz migrations to generate SQL migrations
      ([keto#448](https://github.com/ory/keto/issues/448)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Define and architect SQL schema and queries for querying relations
      ([keto#300](https://github.com/ory/keto/issues/300)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Add health check and version endpoints to read/write REST (and gRPC?) API
      ([keto#422](https://github.com/ory/keto/issues/422)) -
      [@Patrik](https://github.com/zepatrik)
- [x] SQL: make strings variable length
      ([keto#430](https://github.com/ory/keto/issues/430))
- [x] Persistence errors are not properly handled and wrapped
      ([keto#432](https://github.com/ory/keto/issues/432)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Non-nil return values despite errors
      ([keto#433](https://github.com/ory/keto/issues/433)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Don't use fmt.Sprintf to construct queries
      ([keto#434](https://github.com/ory/keto/issues/434)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Add golangci-lint with gosec linter
      ([keto#435](https://github.com/ory/keto/issues/435)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Use `errors.WithStack` where ever non-keto code returns errors.
      ([keto#437](https://github.com/ory/keto/issues/437)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Make errors debuggable and understandable
      ([keto#438](https://github.com/ory/keto/issues/438)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Allow deletion of relation tuples
      ([keto#452](https://github.com/ory/keto/issues/452)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Command status --block is not blocking until healthy
      ([keto#456](https://github.com/ory/keto/issues/456)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Namespace migrate status command
      ([keto#502](https://github.com/ory/keto/issues/502)) -
      [@Patrik](https://github.com/zepatrik)

#### Pull Requests

- [x] feat: first API draft and generation
      ([keto#315](https://github.com/ory/keto/pull/315)) -
      [@Patrik](https://github.com/zepatrik),
      [@Robin Brämer](https://github.com/robinbraemer)

### [Rfc](https://github.com/ory/keto/labels/rfc)

A request for comments to discuss and share ideas.

#### Issues

- [x] Consider rename WriteRelationTuples in WriteService
      ([keto#351](https://github.com/ory/keto/issues/351)) -
      [@Patrik](https://github.com/zepatrik),
      [@Robin Brämer](https://github.com/robinbraemer)
- [x] Allow deletion of relation tuples
      ([keto#452](https://github.com/ory/keto/issues/452)) -
      [@Patrik](https://github.com/zepatrik)

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

- [ ] Relations are not unique
      ([keto#292](https://github.com/ory/keto/issues/292)) -
      [@Patrik](https://github.com/zepatrik)
- [ ] Handle circular relation tuples
      ([keto#411](https://github.com/ory/keto/issues/411)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Move models package to relations
      ([keto#295](https://github.com/ory/keto/issues/295)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Listing relation tuples where none match cause SQL error
      ([keto#439](https://github.com/ory/keto/issues/439)) -
      [@Patrik](https://github.com/zepatrik)

### [Feat](https://github.com/ory/keto/labels/feat)

New feature or request.

#### Issues

- [ ] Define naming conventions for objects, relations, namespaces
      ([keto#288](https://github.com/ory/keto/issues/288)) -
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
- [ ] [next-gen] Allow defining userset rewrites
      ([keto#263](https://github.com/ory/keto/issues/263)) -
      [@Patrik](https://github.com/zepatrik)
- [ ] Namespace configuration
      ([keto#303](https://github.com/ory/keto/issues/303)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Remove support for AND/OR/XOR in GetRelationTuples
      ([keto#299](https://github.com/ory/keto/issues/299)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Write benchmark tests for relationtuple package
      ([keto#301](https://github.com/ory/keto/issues/301)) -
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

- [ ] Namespace configuration
      ([keto#303](https://github.com/ory/keto/issues/303)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Write benchmark tests for relationtuple package
      ([keto#301](https://github.com/ory/keto/issues/301)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Protobuf APIs & tooling
      ([keto#311](https://github.com/ory/keto/issues/311)) -
      [@Patrik](https://github.com/zepatrik),
      [@Robin Brämer](https://github.com/robinbraemer)
- [x] Listing relation tuples where none match cause SQL error
      ([keto#439](https://github.com/ory/keto/issues/439)) -
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
- [ ] Handle circular relation tuples
      ([keto#411](https://github.com/ory/keto/issues/411)) -
      [@Patrik](https://github.com/zepatrik)
- [x] Design decisions, Clarifications and Proposals
      ([keto#307](https://github.com/ory/keto/issues/307)) -
      [@hackerman](https://github.com/aeneasr),
      [@Patrik](https://github.com/zepatrik)

## [next](https://github.com/ory/keto/milestone/2)

_This milestone does not have a description._

### [Bug](https://github.com/ory/keto/labels/bug)

Something is not working.

#### Issues

- [x] Slash (/) in role or policy id causes 404 error for GET and DELETE
      ([keto#140](https://github.com/ory/keto/issues/140))
- [x] Keto is posting plain text passwords when there is an issue with DSN
      ([keto#237](https://github.com/ory/keto/issues/237))

### [Feat](https://github.com/ory/keto/labels/feat)

New feature or request.

#### Issues

- [x] Roles and Policies Filter by using flavor strategy
      ([keto#186](https://github.com/ory/keto/issues/186))
- [x] Evaluate queries needs to get the entire database in cache to works
      ([keto#187](https://github.com/ory/keto/issues/187))
- [x] Add a description attribute to access control policy role
      ([keto#213](https://github.com/ory/keto/issues/213))

### [Rfc](https://github.com/ory/keto/labels/rfc)

A request for comments to discuss and share ideas.

#### Issues

- [x] IDs being limited to varchar(64) reduces the usefulness of URNs when you
      are using them globally
      ([keto#197](https://github.com/ory/keto/issues/197))
