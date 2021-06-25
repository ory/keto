---
id: changelog
title: Changelog
custom_edit_url: null
---

# [Unreleased](https://github.com/ory/keto/compare/v0.6.0-alpha.3...8e301198298858fd7f387ef63a7abf4fa55ea240) (2021-06-18)

### Bug Fixes

- Add missing tracers ([#600](https://github.com/ory/keto/issues/600))
  ([aa263be](https://github.com/ory/keto/commit/aa263be9a7830e3c769d7698d36137555ca230bc)),
  closes [#593](https://github.com/ory/keto/issues/593)
- Handle relation tuple cycles in expand and check engine
  ([#623](https://github.com/ory/keto/issues/623))
  ([8e30119](https://github.com/ory/keto/commit/8e301198298858fd7f387ef63a7abf4fa55ea240))
- Log all database connection errors
  ([#588](https://github.com/ory/keto/issues/588))
  ([2b0fad8](https://github.com/ory/keto/commit/2b0fad897e61400bd2a6cdf47f33ff4301e9c5f8))
- Move gRPC client module root up
  ([#620](https://github.com/ory/keto/issues/620))
  ([3b881f6](https://github.com/ory/keto/commit/3b881f6015a93b382b3fbbca4be9259622038b6a)):
  BREAKING: The npm package `@ory/keto-grpc-client` from now on includes all API
  versions. Because of that, the import paths changed. For migrating to the new
  client package, change the import path according to the following example:
  ```diff
  - import acl from '@ory/keto-grpc-client/acl_pb.js'
  + // from the latest version
  + import { acl } from '@ory/keto-grpc-client'
  + // or a specific one
  + import acl from '@ory/keto-grpc-client/ory/keto/acl/v1alpha1/acl_pb.js'
  ```
- Update docker-compose.yml version
  ([#595](https://github.com/ory/keto/issues/595))
  ([7fa4dca](https://github.com/ory/keto/commit/7fa4dca4182a1fa024f9cef0a04163f2cbd882aa)),
  closes [#549](https://github.com/ory/keto/issues/549)

### Documentation

- Fix example not following best practice
  ([#582](https://github.com/ory/keto/issues/582))
  ([a015818](https://github.com/ory/keto/commit/a0158182c5f87cfd4767824e1c5d6cbb8094a4e6))
- Update NPM links due to organisation move
  ([#616](https://github.com/ory/keto/issues/616))
  ([6355bea](https://github.com/ory/keto/commit/6355beae5b5b28c3eee19fdee85b9875cbc165c3))

### Features

- Make generated gRPC client its own module
  ([#583](https://github.com/ory/keto/issues/583))
  ([f0fbb64](https://github.com/ory/keto/commit/f0fbb64b3358e9800854295cebc9ec8b8e56c87a))
- Max_idle_conn_time ([#605](https://github.com/ory/keto/issues/605))
  ([50a8623](https://github.com/ory/keto/commit/50a862338e17f86900ca162da7f3467f55f9f954)),
  closes [#523](https://github.com/ory/keto/issues/523)

# [0.6.0-alpha.3](https://github.com/ory/keto/compare/v0.6.0-alpha.2...v0.6.0-alpha.3) (2021-04-29)

Resolves CRDB and build issues.

### Code Generation

- Pin v0.6.0-alpha.3 release commit
  ([d766968](https://github.com/ory/keto/commit/d766968419d10a68fd843df45316e3436b68d61d))

# [0.6.0-alpha.2](https://github.com/ory/keto/compare/v0.6.0-alpha.1...v0.6.0-alpha.2) (2021-04-29)

This release improves stability and documentation.

### Bug Fixes

- Add npm run format to make format
  ([7d844a8](https://github.com/ory/keto/commit/7d844a8e6412ae561963b97ac26d4682411095d4))
- Makefile target
  ([0e6f612](https://github.com/ory/keto/commit/0e6f6122de7bdbb691ad7cc236b6bc9a3601d39e))
- Move swagger to spec dir
  ([7f6a061](https://github.com/ory/keto/commit/7f6a061aafda275d278bf60f16e90039da45bc57))
- Resolve clidoc issues
  ([ef12b4e](https://github.com/ory/keto/commit/ef12b4e267f34fbf9709fe26023f9b7ae6670c24))
- Update install.sh ([#568](https://github.com/ory/keto/issues/568))
  ([86ab245](https://github.com/ory/keto/commit/86ab24531d608df0b5391ee8ec739291b9a90e20))
- Use correct id
  ([5e02902](https://github.com/ory/keto/commit/5e029020b5ba3931f15d343cf6a9762b064ffd45))
- Use correct id for api
  ([32a6b04](https://github.com/ory/keto/commit/32a6b04609054cba84f7b56ebbe92341ec5dcd98))
- Use sqlite image versions ([#544](https://github.com/ory/keto/issues/544))
  ([ec6cc5e](https://github.com/ory/keto/commit/ec6cc5ed528f1a097ea02669d059e060b7eff824))

### Code Generation

- Pin v0.6.0-alpha.2 release commit
  ([470b2c6](https://github.com/ory/keto/commit/470b2c61c649fe5fcf638c84d4418212ff0330a5))

### Documentation

- Add gRPC client README.md ([#559](https://github.com/ory/keto/issues/559))
  ([9dc3596](https://github.com/ory/keto/commit/9dc35969ada8b0d4d73dee9089c4dc61cd9ea657))
- Change forum to discussions readme
  ([#539](https://github.com/ory/keto/issues/539))
  ([ea2999d](https://github.com/ory/keto/commit/ea2999d4963316810a8d8634fcd123bda31eaa8f))
- Fix cat videos example docker compose
  ([#549](https://github.com/ory/keto/issues/549))
  ([b25a711](https://github.com/ory/keto/commit/b25a7114631957935c71ac6a020ab6bd0c244cd7))
- Fix typo ([#538](https://github.com/ory/keto/issues/538))
  ([99a9693](https://github.com/ory/keto/commit/99a969373497792bb4cd8ff62bf5245087517737))
- Include namespace in olymp library example
  ([#540](https://github.com/ory/keto/issues/540))
  ([135e814](https://github.com/ory/keto/commit/135e8145c383a76b494b469253c949c38f4414a7))
- Update install from source steps to actually work
  ([#548](https://github.com/ory/keto/issues/548))
  ([e662256](https://github.com/ory/keto/commit/e6622564f58b7612b13b11b54e75a7350f52d6de))

### Features

- Global docs sidebar and added cloud pages
  ([c631c82](https://github.com/ory/keto/commit/c631c82b7ff3d12734869ac22730b52e73dcf287))
- Support retryable CRDB transactions
  ([833147d](https://github.com/ory/keto/commit/833147dae40e9ac5bdf220f8aa3f01abd444f791))

# [0.6.0-alpha.1](https://github.com/ory/keto/compare/v0.5.6-alpha.1...v0.6.0-alpha.1) (2021-04-07)

We are extremely happy to announce next-gen Ory Keto which implements
[Zanzibar: Google’s Consistent, Global Authorization System](https://research.google/pubs/pub48190/):

> Zanzibar provides a uniform data model and configuration language for
> expressing a wide range of access control policies from hundreds of client
> services at Google, including Calendar, Cloud, Drive, Maps, Photos, and
> YouTube. Its authorization decisions respect causal ordering of user actions
> and thus provide external consistency amid changes to access control lists and
> object contents. Zanzibar scales to trillions of access control lists and
> millions of authorization requests per second to support services used by
> billions of people. It has maintained 95th-percentile latency of less than 10
> milliseconds and availability of greater than 99.999% over 3 years of
> production use.

Ory Keto is the first open source planet-scale authorization system built with
cloud native technologies (Go, gRPC, newSQL) and architecture. It is also the
first open source implementation of Google Zanzibar :tada:!

Many concepts developer by Google Zanzibar are implemented in Ory Keto already.
Let's take a look!

As of this release, Ory Keto knows how to interpret and operate on the basic
access control lists known as relation tuples. They encode relations between
objects and subjects. One simple example of such a relation tuple could encode
"`user1` has access to file `/foo`", a more complex one could encode "everyone
who has write access on `/foo` has read access on `/foo`".

Ory Keto comes with all the basic APIs as described in the Zanzibar paper. All
of them are available over gRPC and REST.

1. List: query relation tuples
2. Check: determine whether a subject has a relation on an object
3. Expand: get a tree of all subjects who have a relation on an object
4. Change: create, update, and delete relation tuples

For all details, head over to the
[documentation](https://www.ory.sh/keto/docs/concepts/api-overview).

With this release we officially move the "old" Keto to the
[legacy-0.5 branch](https://github.com/ory/keto/tree/legacy-0.5). We will only
provide security fixes from now on. A migration path to v0.6 is planned but not
yet implemented, as the architectures are vastly different. Please refer to
[the issue](https://github.com/ory/keto/issues/318).

We are keen to bring more features and performance improvements. The next
features we will tackle are:

- Subject Set rewrites
- Native ABAC & RBAC Support
- Integration with other policy servers
- Latency reduction through aggressive caching
- Cluster mode that fans out requests over all Keto instances

So stay tuned, :star: this repo, :eyes: releases, and
[subscribe to our newsletter :email:](https://ory.us10.list-manage.com/subscribe?u=ffb1a878e4ec6c0ed312a3480&id=f605a41b53&MERGE0=&group[17097][32]=1).

### Bug Fixes

- Add description attribute to access control policy role
  ([#215](https://github.com/ory/keto/issues/215))
  ([831eba5](https://github.com/ory/keto/commit/831eba59f810ca68561dd584c9df7684df10b843))
- Add leak_sensitive_values to config schema
  ([2b21d2b](https://github.com/ory/keto/commit/2b21d2bdf4ca9523d16159c5f73c4429b692e17d))
- Bump CLI
  ([80c82d0](https://github.com/ory/keto/commit/80c82d026cbfbab8fbb84d850d8980866ecf88df))
- Bump deps and replace swagutil
  ([#212](https://github.com/ory/keto/issues/212))
  ([904258d](https://github.com/ory/keto/commit/904258d23959c3fa96b6d8ccfdb79f6788c106ec))
- Check engine overwrote result in some cases
  ([#412](https://github.com/ory/keto/issues/412))
  ([3404492](https://github.com/ory/keto/commit/3404492002ca5c3f017ef25486e377e911987aa4))
- Check health status in status command
  ([21c64d4](https://github.com/ory/keto/commit/21c64d45f21a505744b9f70d780f9b3079d3822c))
- Check REST API returns JSON object
  ([#460](https://github.com/ory/keto/issues/460))
  ([501dcff](https://github.com/ory/keto/commit/501dcff4427f76902671f6d5733f28722bd51fa7)),
  closes [#406](https://github.com/ory/keto/issues/406)
- Empty relationtuple list should not error
  ([#440](https://github.com/ory/keto/issues/440))
  ([fbcb3e1](https://github.com/ory/keto/commit/fbcb3e1f337b5114d7697fa512ded92b5f409ef4))
- Ensure nil subject is not allowed
  ([#449](https://github.com/ory/keto/issues/449))
  ([7a0fcfc](https://github.com/ory/keto/commit/7a0fcfc4fe83776fa09cf78ee11f407610554d04)):
  The nodejs gRPC client was a great fuzzer and pointed me to some nil pointer
  dereference panics. This adds some input validation to prevent panics.
- Ensure persister errors are handled by sqlcon
  ([#473](https://github.com/ory/keto/issues/473))
  ([4343c4a](https://github.com/ory/keto/commit/4343c4acd8f917fb7ae131e67bca6855d4d61694))
- Handle pagination and errors in the check/expand engines
  ([#398](https://github.com/ory/keto/issues/398))
  ([5eb1a7d](https://github.com/ory/keto/commit/5eb1a7d49af6b43707c122de8727cbd72285cb5c))
- Ignore dist
  ([ba816ea](https://github.com/ory/keto/commit/ba816ea2ca39962f02c08e0c7b75cfe3cf1d963d))
- Ignore x/net false positives
  ([d8b36cb](https://github.com/ory/keto/commit/d8b36cb1812abf7265ac15c29780222be025186b))
- Improve CLI remote sourcing ([#474](https://github.com/ory/keto/issues/474))
  ([a85f4d7](https://github.com/ory/keto/commit/a85f4d7470ac3744476e82e5889b97d5a0680473))
- Improve handlers and add tests
  ([#470](https://github.com/ory/keto/issues/470))
  ([ca5ccb9](https://github.com/ory/keto/commit/ca5ccb9c237fdcc4db031ec97a75616a859cbf8f))
- Insert relation tuples without fmt.Sprintf
  ([#443](https://github.com/ory/keto/issues/443))
  ([fe507bb](https://github.com/ory/keto/commit/fe507bb4ea719780e732d098291aa190d6b1c441))
- Minor bugfixes ([#371](https://github.com/ory/keto/issues/371))
  ([185ee1e](https://github.com/ory/keto/commit/185ee1e51bc4bcdee028f71fcaf207b7e342313b))
- Move dockerfile to where it belongs
  ([f087843](https://github.com/ory/keto/commit/f087843ac8f24e741bf39fe65ee5b0a9adf9a5bb))
- Namespace migrator ([#417](https://github.com/ory/keto/issues/417))
  ([ea79300](https://github.com/ory/keto/commit/ea7930064f490b063a712b4e18521f8996931a13)),
  closes [#404](https://github.com/ory/keto/issues/404)
- Remove SQL logging ([#455](https://github.com/ory/keto/issues/455))
  ([d8e2a86](https://github.com/ory/keto/commit/d8e2a869db2a9cfb44423b434330536036b2f421))
- Rename /relationtuple endpoint to /relation-tuples
  ([#519](https://github.com/ory/keto/issues/519))
  ([8eb55f6](https://github.com/ory/keto/commit/8eb55f6269399f2bc5f000b8a768bcdf356c756f))
- Resolve gitignore build
  ([6f04bbb](https://github.com/ory/keto/commit/6f04bbb6057779b4d73d3f94677cea365843f7ac))
- Resolve goreleaser issues
  ([d32767f](https://github.com/ory/keto/commit/d32767f32856cf5bd48514c5d61746417fbed6f5))
- Resolve windows build issues
  ([8bcdfbf](https://github.com/ory/keto/commit/8bcdfbfbdb0b10c03ff93838e8fe6e778236e96d))
- Rewrite check engine to search starting at the object
  ([#310](https://github.com/ory/keto/issues/310))
  ([7d99694](https://github.com/ory/keto/commit/7d9969414ebc8cf6ef5d211ad34f8ae01bd3b4ee)),
  closes [#302](https://github.com/ory/keto/issues/302)
- Secure query building ([#442](https://github.com/ory/keto/issues/442))
  ([c7d2770](https://github.com/ory/keto/commit/c7d2770ed570238fd1262bcc4e5b4afa6c12d80e))
- Strict version enforcement in docker
  ([e45b28f](https://github.com/ory/keto/commit/e45b28fec626db35f1bd4580e5b11c9c94a02669))
- Update dd-trace to fix build issues
  ([2ad489f](https://github.com/ory/keto/commit/2ad489f0d9cae3191718d36823fe25df58ab95e6))
- Update docker to go 1.16 and alpine
  ([c63096c](https://github.com/ory/keto/commit/c63096cb53d2171f22f4a0d4a9ac3c9bfac89d01))
- Use errors.WithStack everywhere
  ([#462](https://github.com/ory/keto/issues/462))
  ([5f25bce](https://github.com/ory/keto/commit/5f25bceea35179c67d24dd95f698dc57b789d87a)),
  closes [#437](https://github.com/ory/keto/issues/437): Fixed all occurrences
  found using the search pattern `return .*, err\n`.
- Use package name in pkger
  ([6435939](https://github.com/ory/keto/commit/6435939ad7e5899505cd0e6261f5dfc819c9ca42))
- **schema:** Add trace level to logger
  ([a5a1402](https://github.com/ory/keto/commit/a5a1402c61e1a37b1a9a349ad5736eaca66bd6a4))
- Use make() to initialize slices
  ([#250](https://github.com/ory/keto/issues/250))
  ([84f028d](https://github.com/ory/keto/commit/84f028dc35665174542e103c0aefc635bb6d3e52)),
  closes [#217](https://github.com/ory/keto/issues/217)

### Build System

- Pin dependency versions of buf and protoc plugins
  ([#338](https://github.com/ory/keto/issues/338))
  ([5a2fd1c](https://github.com/ory/keto/commit/5a2fd1cc8dff02aa7017771adc0d9101f6c86775))

### Code Generation

- Pin v0.6.0-alpha.1 release commit
  ([875af25](https://github.com/ory/keto/commit/875af25f89b813455148e58884dcdf1cd3600b86))

### Code Refactoring

- Data structures ([#279](https://github.com/ory/keto/issues/279))
  ([1316077](https://github.com/ory/keto/commit/131607762d0006e4cf4f93e8731ef7648348b2ec))

### Documentation

- Add check- and expand-API guides
  ([#493](https://github.com/ory/keto/issues/493))
  ([09a25b4](https://github.com/ory/keto/commit/09a25b4063abcfdcd4c0de315a2ef088d6d4e72e))
- Add current features overview ([#505](https://github.com/ory/keto/issues/505))
  ([605afa0](https://github.com/ory/keto/commit/605afa029794ad115bba02e004e1596cea038e8e))
- Add missing pages ([#518](https://github.com/ory/keto/issues/518))
  ([43cbaa9](https://github.com/ory/keto/commit/43cbaa9140cfa0ea3c72f699f6bb34f5ed31d8dd))
- Add namespace and relation naming conventions
  ([#510](https://github.com/ory/keto/issues/510))
  ([dd31865](https://github.com/ory/keto/commit/dd318653178cd45da47f3e7cef507b42708363ef))
- Add performance page ([#413](https://github.com/ory/keto/issues/413))
  ([6fe0639](https://github.com/ory/keto/commit/6fe0639d36087b5ecd555eb6fe5ce949f3f6f0d7)):
  This also refactored the server startup. Functionality did not change.
- Add production guide
  ([a9163c7](https://github.com/ory/keto/commit/a9163c7690c55c8191650c4dfb464b75ea02446b))
- Add zanzibar overview to README.md
  ([#265](https://github.com/ory/keto/issues/265))
  ([15a95b2](https://github.com/ory/keto/commit/15a95b28e745592353e4656d42a9d0bd20ce468f))
- API overview ([#501](https://github.com/ory/keto/issues/501))
  ([05fe03b](https://github.com/ory/keto/commit/05fe03b5bf7a3f790aa6c9c1d3fcdb31304ef6af))
- Concepts ([#429](https://github.com/ory/keto/issues/429))
  ([2f2c885](https://github.com/ory/keto/commit/2f2c88527b3f6d1d46a5c287d8aca0874d18a28d))
- Delete old redirect homepage
  ([c0a3784](https://github.com/ory/keto/commit/c0a378448f8c7723bae68f7b52a019b697b25863))
- Document gRPC SKDs
  ([7583fe8](https://github.com/ory/keto/commit/7583fe8933f6676b4e37477098b1d43d12819b8b))
- Fix grammatical error ([#222](https://github.com/ory/keto/issues/222))
  ([256a0d2](https://github.com/ory/keto/commit/256a0d2e53fe1eb859e41fc539870ae1d5a493d2))
- Fix regression issues
  ([9697bb4](https://github.com/ory/keto/commit/9697bb43dd23c0d1fae74ea42e848883c45dae77))
- Generate gRPC reference page ([#488](https://github.com/ory/keto/issues/488))
  ([93ebe6d](https://github.com/ory/keto/commit/93ebe6db7e887d708503a54c5ec943254e37ca43))
- Improve CLI documentation ([#503](https://github.com/ory/keto/issues/503))
  ([be9327f](https://github.com/ory/keto/commit/be9327f7b28152a78f731043acf83b7092e42e29))
- Minor fixes ([#532](https://github.com/ory/keto/issues/532))
  ([638342e](https://github.com/ory/keto/commit/638342eb9519d9bf609926fb87558071e2815fb3))
- Move development section
  ([9ff393f](https://github.com/ory/keto/commit/9ff393f6cba1fb0a33918377ce505455c34d9dfc))
- Move to json sidebar
  ([257bf96](https://github.com/ory/keto/commit/257bf96044df37c3d7af8a289fb67098d48da1a3))
- Remove duplicate "is"
  ([ca3277d](https://github.com/ory/keto/commit/ca3277d82c1508797bc8c663963407d2e4d9112f))
- Remove duplicate template
  ([1d3b38e](https://github.com/ory/keto/commit/1d3b38e4045b0b874bb1186ea628f5a37353a2e6))
- Remove old documentation ([#426](https://github.com/ory/keto/issues/426))
  ([eb76913](https://github.com/ory/keto/commit/eb7691306018678e024211b51627a1c27e780a6b))
- Replace TODO links ([#512](https://github.com/ory/keto/issues/512))
  ([ad8e20b](https://github.com/ory/keto/commit/ad8e20b3bef2bc46b3a32c2c9ccb6e16e4bad22c))
- Resolve broken links
  ([0d0a50b](https://github.com/ory/keto/commit/0d0a50b3f4112893f32c81adc8edd137b5a62541))
- Simple access check guide ([#451](https://github.com/ory/keto/issues/451))
  ([e0485af](https://github.com/ory/keto/commit/e0485afc46a445868580aa541e962e80cbea0670)):
  This also enables gRPC go, gRPC nodejs, cURL, and Keto CLI code samples to be
  tested.
- Update comment in write response
  ([#329](https://github.com/ory/keto/issues/329))
  ([4ca0baf](https://github.com/ory/keto/commit/4ca0baf62e34402e749e870fe8c0cc893684192c))
- Update install instructions
  ([d2e4123](https://github.com/ory/keto/commit/d2e4123f3e2e58da8be181a0a542e3dcc1313e16))
- Update introduction
  ([5f71d73](https://github.com/ory/keto/commit/5f71d73e2ee95d02abc4cd42a76c98a35942df0c))
- Update README ([#515](https://github.com/ory/keto/issues/515))
  ([18d3cd6](https://github.com/ory/keto/commit/18d3cd61b0a79400170dc0f89860b4614cc4a543)):
  Also format all markdown files in the root.
- Update repository templates
  ([db505f9](https://github.com/ory/keto/commit/db505f9e10755bc20c4623c4f5f99f33283dffda))
- Update repository templates
  ([6c056bb](https://github.com/ory/keto/commit/6c056bb2043af6e82f06fdfa509ab3fa0d5e5d06))
- Update SDK links ([#514](https://github.com/ory/keto/issues/514))
  ([f920fbf](https://github.com/ory/keto/commit/f920fbfc8dcc6711ad9e046578a4506179952be7))
- Update swagger documentation for REST endpoints
  ([c363de6](https://github.com/ory/keto/commit/c363de61edf912fef85acc6bcdac6e1c15c48f4f))
- Use mdx for api reference
  ([340f3a3](https://github.com/ory/keto/commit/340f3a3dd20c82c743e7b3ad6aaf06a4c118b5a1))
- Various improvements and updates
  ([#486](https://github.com/ory/keto/issues/486))
  ([a812ace](https://github.com/ory/keto/commit/a812ace2303214e0e7acb2e283efa1cff0d5d279))

### Features

- Add .dockerignore
  ([8b0ff06](https://github.com/ory/keto/commit/8b0ff066b2508ef2f3629f9a3e2fce601b8dcce1))
- Add and automate version schema
  ([b01eef8](https://github.com/ory/keto/commit/b01eef8d4d5834b5888cb369ecf01ee01b40c24c))
- Add check engine ([#277](https://github.com/ory/keto/issues/277))
  ([396c1ae](https://github.com/ory/keto/commit/396c1ae33b777031f8d59549d9de4a88e3f6b10a))
- Add gRPC health status ([#427](https://github.com/ory/keto/issues/427))
  ([51c4223](https://github.com/ory/keto/commit/51c4223d6cb89a9bfbc115ef20db8350aeb2e8af))
- Add is_last_page to list response
  ([#425](https://github.com/ory/keto/issues/425))
  ([b73d91f](https://github.com/ory/keto/commit/b73d91f061ab155c53d802263c0263aa39e64bdf))
- Add POST REST handler for policy check
  ([7d89860](https://github.com/ory/keto/commit/7d89860bc4a790a69f5bea5b0dbe4a2938c6729f))
- Add relation write API ([#275](https://github.com/ory/keto/issues/275))
  ([f2ddb9d](https://github.com/ory/keto/commit/f2ddb9d884ed71037b5371c00bb10b63d25d47c0))
- Add REST and gRPC logger middlewares
  ([#436](https://github.com/ory/keto/issues/436))
  ([615eb0b](https://github.com/ory/keto/commit/615eb0bec3bdc0fd26abc7af0b8990269b0cbedd))
- Add SQA telemetry ([#535](https://github.com/ory/keto/issues/535))
  ([9f6472b](https://github.com/ory/keto/commit/9f6472b0c996505d41058e9b55afa8fd6b9bb2d5))
- Add sql persister ([#350](https://github.com/ory/keto/issues/350))
  ([d595d52](https://github.com/ory/keto/commit/d595d52dabb8f4953b5c23d3a8154cac13d00306))
- Add tracing ([#536](https://github.com/ory/keto/issues/536))
  ([b57a144](https://github.com/ory/keto/commit/b57a144e0a7ec39d5831dbb79840c2b25c044e6a))
- Allow to apply namespace migrations together with regular migrations
  ([#441](https://github.com/ory/keto/issues/441))
  ([57e2bbc](https://github.com/ory/keto/commit/57e2bbc5eaebe43834f2432eb1ee2820d9cb2988))
- Delete relation tuples ([#457](https://github.com/ory/keto/issues/457))
  ([3ec8afa](https://github.com/ory/keto/commit/3ec8afa68c5b5ddc26609b9afd17cc0d06cd82bf)),
  closes [#452](https://github.com/ory/keto/issues/452)
- Dockerfile and docker compose example
  ([#390](https://github.com/ory/keto/issues/390))
  ([10cd0b3](https://github.com/ory/keto/commit/10cd0b39c12ef96710bda6ff013f7c5eeae97118))
- Expand API ([#285](https://github.com/ory/keto/issues/285))
  ([a3ca0b8](https://github.com/ory/keto/commit/a3ca0b8a109b63f443e359cd8ff18a7b3e489f84))
- Expand GPRC service and CLI ([#383](https://github.com/ory/keto/issues/383))
  ([acf2154](https://github.com/ory/keto/commit/acf21546d3e135deb77c853b751a3da3a7b16f00))
- First API draft and generation
  ([#315](https://github.com/ory/keto/issues/315))
  ([bda5d8b](https://github.com/ory/keto/commit/bda5d8b7e90d749600f5b5e169df8a6ec3705b22))
- GRPC status codes and improved error messages
  ([#467](https://github.com/ory/keto/issues/467))
  ([4a4f8c6](https://github.com/ory/keto/commit/4a4f8c6b323664329414b61e7d80d7838face730))
- GRPC version API ([#475](https://github.com/ory/keto/issues/475))
  ([89cc46f](https://github.com/ory/keto/commit/89cc46fe4a13b062693d3db4f803834ba37f4e48))
- Implement goreleaser pipeline
  ([888ac43](https://github.com/ory/keto/commit/888ac43e6f706f619b2f1b58271dd027094c9ae9)),
  closes [#410](https://github.com/ory/keto/issues/410)
- Incorporate new GRPC API structure
  ([#331](https://github.com/ory/keto/issues/331))
  ([e0916ad](https://github.com/ory/keto/commit/e0916ad00c81b24177cfe45faf77b93d2c33dc1f))
- Koanf and namespace configuration
  ([#367](https://github.com/ory/keto/issues/367))
  ([3ad32bc](https://github.com/ory/keto/commit/3ad32bc13a4d96135be8031eb6fe4c15868272ca))
- Namespace configuration ([#324](https://github.com/ory/keto/issues/324))
  ([b94f50d](https://github.com/ory/keto/commit/b94f50d1800c47a43561df5009cb38b44ccd0088))
- Namespace migrate status CLI ([#508](https://github.com/ory/keto/issues/508))
  ([e3f7ad9](https://github.com/ory/keto/commit/e3f7ad91585b616e97f85ce0f55c76406b6c4d0a)):
  This also refactors the current `migrate` and `namespace migrate` commands.
- Nodejs gRPC definitions ([#447](https://github.com/ory/keto/issues/447))
  ([3b5c313](https://github.com/ory/keto/commit/3b5c31326645adb2d5b14ced901771a7ba00fd1c)):
  Includes Typescript definitions.
- Read API ([#269](https://github.com/ory/keto/issues/269))
  ([de5119a](https://github.com/ory/keto/commit/de5119a6e3c7563cfc2e1ada12d47b27ebd7faaa)):
  This is a first draft of the read API. It is reachable by REST and gRPC calls.
  The main purpose of this PR is to establish the basic repository structure and
  define the API.
- Relationtuple parse command ([#490](https://github.com/ory/keto/issues/490))
  ([91a3cf4](https://github.com/ory/keto/commit/91a3cf47fbdb8203b799cf7c69bcf3dbbfb98b3a)):
  This command parses the relation tuple format used in the docs. It greatly
  improves the experience when copying something from the documentation. It can
  especially be used to pipe relation tuples into other commands, e.g.:
  ```shell
  echo "messages:02y_15_4w350m3#decypher@john" | \
    keto relation-tuple parse - --format json | \
    keto relation-tuple create -
  ```
- REST patch relation tuples ([#491](https://github.com/ory/keto/issues/491))
  ([d38618a](https://github.com/ory/keto/commit/d38618a9e647902ce019396ff1c33973020bf797)):
  The new PATCH handler allows transactional changes similar to the already
  existing gRPC service.
- Separate and multiplex ports based on read/write privilege
  ([#397](https://github.com/ory/keto/issues/397))
  ([6918ac3](https://github.com/ory/keto/commit/6918ac3bfa355cbd551e44376c214f412e3414e4))
- Swagger SDK ([#476](https://github.com/ory/keto/issues/476))
  ([011888c](https://github.com/ory/keto/commit/011888c2b7e2d0f7b8923c994c70e62d374a2830))

### Tests

- Add command tests ([#487](https://github.com/ory/keto/issues/487))
  ([61c28e4](https://github.com/ory/keto/commit/61c28e48a5c3f623e5cc133e69ba368c5103f414))
- Add dedicated persistence tests
  ([#416](https://github.com/ory/keto/issues/416))
  ([4e98906](https://github.com/ory/keto/commit/4e9890605edf3ea26134917a95bfa6fbb176565e))
- Add handler tests ([#478](https://github.com/ory/keto/issues/478))
  ([9315a77](https://github.com/ory/keto/commit/9315a77820d50e400b78f2f019a871be022a9887))
- Add initial e2e test ([#380](https://github.com/ory/keto/issues/380))
  ([dc5d3c9](https://github.com/ory/keto/commit/dc5d3c9d02604fddbfa56ac5ebbc1fef56a881d9))
- Add relationtuple definition tests
  ([#415](https://github.com/ory/keto/issues/415))
  ([2e3dcb2](https://github.com/ory/keto/commit/2e3dcb200a7769dc8710d311ca08a7515012fbdd))
- Enable GRPC client in e2e test
  ([#382](https://github.com/ory/keto/issues/382))
  ([4e5c6ae](https://github.com/ory/keto/commit/4e5c6aed56e5a449003956ec114ec131be068aaf))
- Improve docs sample tests ([#461](https://github.com/ory/keto/issues/461))
  ([6e0e5e6](https://github.com/ory/keto/commit/6e0e5e6184916e894fd4694cfa3a158f11fae11f))

# [0.5.6-alpha.1](https://github.com/ory/keto/compare/v0.5.5-alpha.1...v0.5.6-alpha.1) (2020-05-28)

This release bumps vulnerable transient dependencies (those are not actually
used in ORY Keto) and updates several documentation pages and improves
structured logging output. Additionally, ORY Keto now uses the updated release
pipeline!

### Bug Fixes

- Update install script
  ([21e1bf0](https://github.com/ory/keto/commit/21e1bf05177576a9d743bd11744ef6a42be50b8d))

### Chores

- Pin v0.5.6-alpha.1 release commit
  ([ed0da08](https://github.com/ory/keto/commit/ed0da08a03a910660358fc56c568692325749b6d))

# [0.5.5-alpha.1](https://github.com/ory/keto/compare/v0.5.4-alpha.1...v0.5.5-alpha.1) (2020-05-28)

This release bumps vulnerable transient dependencies (those are not actually
used in ORY Keto) and updates several documentation pages and improves
structured logging output. Additionally, ORY Keto now uses the updated release
pipeline!

### Bug Fixes

- Move deps to go_mod_indirect_pins
  ([dd3e971](https://github.com/ory/keto/commit/dd3e971ac418baf10c1b33005acc7e6f66fb0d85))
- Resolve test issues
  ([9bd9956](https://github.com/ory/keto/commit/9bd9956e33731f1619c32e1e6b7c78f42e7c47c3))
- Update install.sh script
  ([f64d320](https://github.com/ory/keto/commit/f64d320b6424fe3256eb7fad1c94dcc1ef0bf487))
- Use semver-regex replacer func
  ([2cc3bbb](https://github.com/ory/keto/commit/2cc3bbb2d75ba5fa7a3653d7adcaa712ff38c603))

### Chores

- Pin v0.5.5-alpha.1 release commit
  ([4666a0f](https://github.com/ory/keto/commit/4666a0f258f253d19a14eca34f4b7049f2d0afa2))

### Documentation

- Add missing colon in docker run command
  ([#193](https://github.com/ory/keto/issues/193))
  ([383063d](https://github.com/ory/keto/commit/383063d260d995665da4c02c9a7bac7e06a2c8d3))
- Update github templates ([#182](https://github.com/ory/keto/issues/182))
  ([72ea09b](https://github.com/ory/keto/commit/72ea09bbbf9925d7705842703b32826376f636e4))
- Update github templates ([#184](https://github.com/ory/keto/issues/184))
  ([ed546b7](https://github.com/ory/keto/commit/ed546b7a2b9ee690284a48c641edd1570464d71f))
- Update github templates ([#188](https://github.com/ory/keto/issues/188))
  ([ebd75b2](https://github.com/ory/keto/commit/ebd75b2f6545ff4372773f6370300c7b2ca71c51))
- Update github templates ([#189](https://github.com/ory/keto/issues/189))
  ([fd4c0b1](https://github.com/ory/keto/commit/fd4c0b17bcb1c281baac1772ab94e305ec8c5c86))
- Update github templates ([#195](https://github.com/ory/keto/issues/195))
  ([ba0943c](https://github.com/ory/keto/commit/ba0943c45d36ef10bdf1169f0aeef439a3a67d28))
- Update linux install guide ([#191](https://github.com/ory/keto/issues/191))
  ([7d8b24b](https://github.com/ory/keto/commit/7d8b24bddb9c92feb78c7b65f39434d538773b58))
- Update repository templates
  ([ea65b5c](https://github.com/ory/keto/commit/ea65b5c5ada0a7453326fa755aa914306f1b1851))
- Use central banner repo for README
  ([0d95d97](https://github.com/ory/keto/commit/0d95d97504df4d0ab57d18dc6d0a824a3f8f5896))
- Use correct banner
  ([c6dfe28](https://github.com/ory/keto/commit/c6dfe280fd962169c424834cea040a408c1bc83f))
- Use correct version
  ([5f7030c](https://github.com/ory/keto/commit/5f7030c9069fe392200be72f8ce1a93890fbbba8)),
  closes [#200](https://github.com/ory/keto/issues/200)
- Use correct versions in install docs
  ([52e6c34](https://github.com/ory/keto/commit/52e6c34780ed41c169504d71c39459898b5d14f9))

# [0.5.4-alpha.1](https://github.com/ory/keto/compare/v0.5.3-alpha.3...v0.5.4-alpha.1) (2020-04-07)

fix: resolve panic when executing migrations (#178)

Closes #177

### Bug Fixes

- Resolve panic when executing migrations
  ([#178](https://github.com/ory/keto/issues/178))
  ([7e83fee](https://github.com/ory/keto/commit/7e83feefaad041c60f09232ac44ed8b7240c6558)),
  closes [#177](https://github.com/ory/keto/issues/177)

# [0.5.3-alpha.3](https://github.com/ory/keto/compare/v0.5.3-alpha.2...v0.5.3-alpha.3) (2020-04-06)

autogen(docs): regenerate and update changelog

### Code Generation

- **docs:** Regenerate and update changelog
  ([769cef9](https://github.com/ory/keto/commit/769cef90f27ba9c203d3faf47272287ab17dc7eb))

### Code Refactoring

- Move docs to this repository ([#172](https://github.com/ory/keto/issues/172))
  ([312480d](https://github.com/ory/keto/commit/312480de3cefc5b72916ba95d8287443cf3ccb3d))

### Documentation

- Regenerate and update changelog
  ([dda79b1](https://github.com/ory/keto/commit/dda79b106a18bc33d70ae60e352118b0d288d26b))
- Regenerate and update changelog
  ([9048dd8](https://github.com/ory/keto/commit/9048dd8d8a0f0654072b3d4b77261fe947a34ece))
- Regenerate and update changelog
  ([806f68c](https://github.com/ory/keto/commit/806f68c603781742e0177ec0b2deecaf64c5b721))
- Regenerate and update changelog
  ([8905ee7](https://github.com/ory/keto/commit/8905ee74d4ec394af92240e180cc5d7f6493cb2f))
- Regenerate and update changelog
  ([203c1cc](https://github.com/ory/keto/commit/203c1cc659a72f81a370d7b9b7fbda60e7c96c9e))
- Regenerate and update changelog
  ([8875a95](https://github.com/ory/keto/commit/8875a95b35df57668acb27820a3aff1cdfbe8b30))
- Regenerate and update changelog
  ([28ddd3e](https://github.com/ory/keto/commit/28ddd3e1483afe8571b3d2bf9afcc31386d85f7f))
- Regenerate and update changelog
  ([927c4ed](https://github.com/ory/keto/commit/927c4edc4a770133bcb34bc044dd5c5e0eb3ffb7))
- Updates issue and pull request templates
  ([#168](https://github.com/ory/keto/issues/168))
  ([29a38a8](https://github.com/ory/keto/commit/29a38a85c61ec2c8d0ad2ce6d5a0f9e9d74b52f7))
- Updates issue and pull request templates
  ([#169](https://github.com/ory/keto/issues/169))
  ([99b7d5d](https://github.com/ory/keto/commit/99b7d5de24fed1aed746c4447a390d084632f89a))
- Updates issue and pull request templates
  ([#171](https://github.com/ory/keto/issues/171))
  ([7a9876b](https://github.com/ory/keto/commit/7a9876b8ed4282f50f886a025033641bd027a0e2))

# [0.5.3-alpha.1](https://github.com/ory/keto/compare/v0.5.2...v0.5.3-alpha.1) (2020-04-03)

chore: move to ory analytics fork (#167)

### Chores

- Move to ory analytics fork ([#167](https://github.com/ory/keto/issues/167))
  ([f824011](https://github.com/ory/keto/commit/f824011b4d19058504b3a43ed53a420619444a51))

# [0.5.2](https://github.com/ory/keto/compare/v0.5.1-alpha.1...v0.5.2) (2020-04-02)

docs: Regenerate and update changelog

### Documentation

- Regenerate and update changelog
  ([1e52100](https://github.com/ory/keto/commit/1e521001a43a0a13e2224e1a44956442ac6ffbc7))
- Regenerate and update changelog
  ([e4d32a6](https://github.com/ory/keto/commit/e4d32a62c1ae96115ea50bb471f5ff2ce2f2c4b9))

# [0.5.0](https://github.com/ory/keto/compare/v0.4.5-alpha.1...v0.5.0) (2020-04-02)

docs: use real json bool type in swagger (#162)

Closes #160

### Bug Fixes

- Move to ory sqa service ([#159](https://github.com/ory/keto/issues/159))
  ([c3bf1b1](https://github.com/ory/keto/commit/c3bf1b1964a14be4cc296aae98d0739e65917e18))
- Use correct response mode for removeOryAccessControlPolicyRoleMe…
  ([#161](https://github.com/ory/keto/issues/161))
  ([17543cf](https://github.com/ory/keto/commit/17543cfef63a1d040a2234bd63b210fb9c4f6015))

### Documentation

- Regenerate and update changelog
  ([6a77f75](https://github.com/ory/keto/commit/6a77f75d66e89420f2daf2fae011d31bcfa34008))
- Regenerate and update changelog
  ([c8c9d29](https://github.com/ory/keto/commit/c8c9d29e77ef53e1196cc6fe600c53d93376229b))
- Regenerate and update changelog
  ([fe8327d](https://github.com/ory/keto/commit/fe8327d951394084df7785166c9a9578c1ab0643))
- Regenerate and update changelog
  ([b5b1d66](https://github.com/ory/keto/commit/b5b1d66a4b933df8789337cce3f6d6bf391b617b))
- Update forum and chat links
  ([e96d7ba](https://github.com/ory/keto/commit/e96d7ba3dcc693c22eb983b3f58a05c9c6adbda7))
- Updates issue and pull request templates
  ([#158](https://github.com/ory/keto/issues/158))
  ([ab14cfa](https://github.com/ory/keto/commit/ab14cfa51ce195b26a83c050452530a5008589d7))
- Use real json bool type in swagger
  ([#162](https://github.com/ory/keto/issues/162))
  ([5349e7f](https://github.com/ory/keto/commit/5349e7f910ad22558a01b76be62db2136b5eb301)),
  closes [#160](https://github.com/ory/keto/issues/160)

# [0.4.5-alpha.1](https://github.com/ory/keto/compare/v0.4.4-alpha.1...v0.4.5-alpha.1) (2020-02-29)

docs: Regenerate and update changelog

### Bug Fixes

- **driver:** Extract scheme from DSN using sqlcon.GetDriverName
  ([#156](https://github.com/ory/keto/issues/156))
  ([187e289](https://github.com/ory/keto/commit/187e289f1a235b5cacf2a0b7ca5e98c384fa7a14)),
  closes [#145](https://github.com/ory/keto/issues/145)

### Documentation

- Regenerate and update changelog
  ([41513da](https://github.com/ory/keto/commit/41513da35ea038f3c4cc2d98b9796cee5b5a8b92))

# [0.4.4-alpha.1](https://github.com/ory/keto/compare/v0.4.3-alpha.2...v0.4.4-alpha.1) (2020-02-14)

docs: Regenerate and update changelog

### Bug Fixes

- **goreleaser:** Update brew section
  ([0918ff3](https://github.com/ory/keto/commit/0918ff3032eeecd26c67d6249c7e28e71ee110af))

### Documentation

- Prepare ecosystem automation
  ([2e39be7](https://github.com/ory/keto/commit/2e39be79ebad1cec021ae3ee4b0a75ffea4b7424))
- Regenerate and update changelog
  ([009c4c4](https://github.com/ory/keto/commit/009c4c4e4fd4c5607cc30cc9622fd0f82e3891f3))
- Regenerate and update changelog
  ([49f3c4b](https://github.com/ory/keto/commit/49f3c4ba34df5879d8f48cc96bf0df9dad820362))
- Updates issue and pull request templates
  ([#153](https://github.com/ory/keto/issues/153))
  ([7fb7521](https://github.com/ory/keto/commit/7fb752114e1e2a91ab96fdb546835de8aee4926b))

### Features

- **ci:** Add nancy vuln scanner
  ([#152](https://github.com/ory/keto/issues/152))
  ([c19c2b9](https://github.com/ory/keto/commit/c19c2b9efe8299b8878cc8099fe314d8dcda3a08))

### Unclassified

- Update CHANGELOG [ci skip]
  ([63fe513](https://github.com/ory/keto/commit/63fe513d22ec3747a95cdb8f797ba1eba5ca344f))
- Update CHANGELOG [ci skip]
  ([7b7c3ac](https://github.com/ory/keto/commit/7b7c3ac6c06c072fea1b64624ea79a3fd406b09c))
- Update CHANGELOG [ci skip]
  ([8886392](https://github.com/ory/keto/commit/8886392b39fb46ad338c8284866d4dae64ad1826))
- Update CHANGELOG [ci skip]
  ([5bbc284](https://github.com/ory/keto/commit/5bbc2844c49b0a68ba3bd8b003d91f87e2aed9e2))

# [0.4.3-alpha.2](https://github.com/ory/keto/compare/v0.4.3-alpha.1...v0.4.3-alpha.2) (2020-01-31)

Update README.md

### Unclassified

- Update README.md
  ([0ab9c6f](https://github.com/ory/keto/commit/0ab9c6f372a1538a958a68b34315c9167b5a9093))
- Update CHANGELOG [ci skip]
  ([f0a1428](https://github.com/ory/keto/commit/f0a1428f4b99ceb35ff4f1e839bc5237e19db628))

# [0.4.3-alpha.1](https://github.com/ory/keto/compare/v0.4.2-alpha.1...v0.4.3-alpha.1) (2020-01-23)

Disable access logging for health endpoints (#151)

Closes #150

### Unclassified

- Disable access logging for health endpoints (#151)
  ([6ca0c09](https://github.com/ory/keto/commit/6ca0c09b5618122762475cffdc9c32adf28456a1)),
  closes [#151](https://github.com/ory/keto/issues/151)
  [#150](https://github.com/ory/keto/issues/150)

# [0.4.2-alpha.1](https://github.com/ory/keto/compare/v0.4.1-beta.1...v0.4.2-alpha.1) (2020-01-14)

Update CHANGELOG [ci skip]

### Unclassified

- Update CHANGELOG [ci skip]
  ([afaabde](https://github.com/ory/keto/commit/afaabde63affcf568e3090e55b4b957edff2890c))

# [0.4.1-beta.1](https://github.com/ory/keto/compare/v0.4.0-sandbox...v0.4.1-beta.1) (2020-01-13)

Update CHANGELOG [ci skip]

### Unclassified

- Update CHANGELOG [ci skip]
  ([e3ca5a7](https://github.com/ory/keto/commit/e3ca5a7d8b9827ffc7b31a8b5e459db3e912a590))
- Update SDK
  ([5dd6237](https://github.com/ory/keto/commit/5dd623755d4832f33c3dcefb778a9a70eace7b52))

# [0.4.0-alpha.1](https://github.com/ory/keto/compare/v0.3.9-sandbox...v0.4.0-alpha.1) (2020-01-13)

Move to new SDK generators (#146)

### Unclassified

- Move to new SDK generators (#146)
  ([4f51a09](https://github.com/ory/keto/commit/4f51a0948723efc092f1887b111d1e6dd590a075)),
  closes [#146](https://github.com/ory/keto/issues/146)
- Fix typos in the README (#144)
  ([85d838c](https://github.com/ory/keto/commit/85d838c0872c73eb70b5bfff1ccb175b07f6b1e4)),
  closes [#144](https://github.com/ory/keto/issues/144)

# [0.3.9-sandbox](https://github.com/ory/keto/compare/v0.3.8-sandbox...v0.3.9-sandbox) (2019-12-16)

Update go modules

### Unclassified

- Update go modules
  ([1151e07](https://github.com/ory/keto/commit/1151e0755c974b0aea86be5aaeae365ea9aef094))

# [0.3.7-sandbox](https://github.com/ory/keto/compare/v0.3.6-sandbox...v0.3.7-sandbox) (2019-12-11)

Update documentation banner image (#143)

### Unclassified

- Update documentation banner image (#143)
  ([e444755](https://github.com/ory/keto/commit/e4447552031a4f26ec21a336071b0bb19843df61)),
  closes [#143](https://github.com/ory/keto/issues/143)
- Revert incorrect license changes
  ([094c4f3](https://github.com/ory/keto/commit/094c4f30184d77a05044087c13e71ce4adb4d735))
- Fix invalid pseudo version ([#138](https://github.com/ory/keto/issues/138))
  ([79b4457](https://github.com/ory/keto/commit/79b4457f0162197ba267edbb8c0031c47e03bade))

# [0.3.6-sandbox](https://github.com/ory/keto/compare/v0.3.5-sandbox...v0.3.6-sandbox) (2019-10-16)

Resolve issues with mysql tests (#137)

### Unclassified

- Resolve issues with mysql tests (#137)
  ([ef5aec8](https://github.com/ory/keto/commit/ef5aec8e493199c46b78e8f1257aa41df9545f28)),
  closes [#137](https://github.com/ory/keto/issues/137)

# [0.3.5-sandbox](https://github.com/ory/keto/compare/v0.3.4-sandbox...v0.3.5-sandbox) (2019-08-21)

Implement roles and policies filter (#124)

### Documentation

- Incorporates changes from version v0.3.3-sandbox
  ([57686d2](https://github.com/ory/keto/commit/57686d2e30b229cae33e717eb8b3db9da3bdaf0a))
- README grammar fixes ([#114](https://github.com/ory/keto/issues/114))
  ([e592736](https://github.com/ory/keto/commit/e5927360300d8c4fbea841c1c2fb92b48b77885e))
- Updates issue and pull request templates
  ([#110](https://github.com/ory/keto/issues/110))
  ([80c8516](https://github.com/ory/keto/commit/80c8516efbcf33902d8a45f1dc7dbafff2aab8b1))
- Updates issue and pull request templates
  ([#111](https://github.com/ory/keto/issues/111))
  ([22305d0](https://github.com/ory/keto/commit/22305d0a9b5114de8125c16030bbcd1de695ae9b))
- Updates issue and pull request templates
  ([#112](https://github.com/ory/keto/issues/112))
  ([dccada9](https://github.com/ory/keto/commit/dccada9a2189bbd899c5c4a18665a92113fe6cd7))
- Updates issue and pull request templates
  ([#125](https://github.com/ory/keto/issues/125))
  ([15f373a](https://github.com/ory/keto/commit/15f373a16b8cfbd6cdad2bda5f161e171c566137))
- Updates issue and pull request templates
  ([#128](https://github.com/ory/keto/issues/128))
  ([eaf8e33](https://github.com/ory/keto/commit/eaf8e33f3904484635924bdac190c8dc7b60f939))
- Updates issue and pull request templates
  ([#130](https://github.com/ory/keto/issues/130))
  ([a440d14](https://github.com/ory/keto/commit/a440d142275a7a91a0a6bb487fe47d22247f4988))
- Updates issue and pull request templates
  ([#131](https://github.com/ory/keto/issues/131))
  ([dbf2cb2](https://github.com/ory/keto/commit/dbf2cb23c5b6f0f1ee0be5c0b5a58fb0c3dbefd1))
- Updates issue and pull request templates
  ([#132](https://github.com/ory/keto/issues/132))
  ([e121048](https://github.com/ory/keto/commit/e121048d10627ed32a07e26455efd69248f1bd95))
- Updates issue and pull request templates
  ([#133](https://github.com/ory/keto/issues/133))
  ([1b7490a](https://github.com/ory/keto/commit/1b7490abc1d5d0501b66595eb2d92834b6fb0345))

### Unclassified

- Implement roles and policies filter (#124)
  ([db94481](https://github.com/ory/keto/commit/db9448103621a6a8cd086a4cef6c6a22398e621f)),
  closes [#124](https://github.com/ory/keto/issues/124)
- Add adopters placeholder ([#129](https://github.com/ory/keto/issues/129))
  ([b814838](https://github.com/ory/keto/commit/b8148388b8bea97d1f1b4b54de2f0b8ef6b8b6c7))
- Improve documentation (#126)
  ([aabb04d](https://github.com/ory/keto/commit/aabb04d5f283d3c73eb3f3531b4e470ae716db5e)),
  closes [#126](https://github.com/ory/keto/issues/126)
- Create FUNDING.yml
  ([571b447](https://github.com/ory/keto/commit/571b447ed3a02f43623ef5c5adc09682b5f379bd))
- Use non-root user in image ([#116](https://github.com/ory/keto/issues/116))
  ([a493e55](https://github.com/ory/keto/commit/a493e550a8bb86d99164f4ea76dbcecf76c9c2c1))
- Remove binary license (#117)
  ([6e85f7c](https://github.com/ory/keto/commit/6e85f7c6f430e88fb4117a131f57bd69466a8ca1)),
  closes [#117](https://github.com/ory/keto/issues/117)

# [0.3.3-sandbox](https://github.com/ory/keto/compare/v0.3.1-sandbox...v0.3.3-sandbox) (2019-05-18)

ci: Resolve goreleaser issues (#108)

### Continuous Integration

- Resolve goreleaser issues ([#108](https://github.com/ory/keto/issues/108))
  ([5753f27](https://github.com/ory/keto/commit/5753f27a9e89ccdda7c02969217c253aa72cb94b))

### Documentation

- Incorporates changes from version v0.3.1-sandbox
  ([b8a0029](https://github.com/ory/keto/commit/b8a002937483a0f71fe5aba26bb18beb41886249))
- Updates issue and pull request templates
  ([#106](https://github.com/ory/keto/issues/106))
  ([54a5a27](https://github.com/ory/keto/commit/54a5a27f24a90ab3c5f9915f36582b85eecd0d62))

# [0.3.1-sandbox](https://github.com/ory/keto/compare/v0.3.0-sandbox...v0.3.1-sandbox) (2019-04-29)

ci: Use image that includes bash/sh for release docs (#103)

Signed-off-by: aeneasr <aeneas@ory.sh>

### Continuous Integration

- Use image that includes bash/sh for release docs
  ([#103](https://github.com/ory/keto/issues/103))
  ([e9d3027](https://github.com/ory/keto/commit/e9d3027fc62b20f28cd7a023222390e24d565eb1))

### Documentation

- Incorporates changes from version v0.3.0-sandbox
  ([605d2f4](https://github.com/ory/keto/commit/605d2f43621b806b750edc81d439edc92cfb7c38))

### Unclassified

- Allow configuration files and update UPGRADE guide. (#102)
  ([3934dc6](https://github.com/ory/keto/commit/3934dc6e690822358067b43920048d45a4b7799b)),
  closes [#102](https://github.com/ory/keto/issues/102)

# [0.3.0-sandbox](https://github.com/ory/keto/compare/v0.2.3-sandbox+oryOS.10...v0.3.0-sandbox) (2019-04-29)

docker: Remove full tag from build pipeline (#101)

Signed-off-by: aeneasr <aeneas@ory.sh>

### Documentation

- Update patrons
  ([c8dc7cd](https://github.com/ory/keto/commit/c8dc7cdc68676970328b55648b8d6e469c77fbfd))

### Unclassified

- Improve naming for ory policies
  ([#100](https://github.com/ory/keto/issues/100))
  ([b39703d](https://github.com/ory/keto/commit/b39703d362d333213fcb7d3782e363d09b6dabbd))
- Remove full tag from build pipeline
  ([#101](https://github.com/ory/keto/issues/101))
  ([602a273](https://github.com/ory/keto/commit/602a273dc5a0c29e80a22f04adb937ab385c4512))
- Remove duplicate code in Makefile (#99)
  ([04f5223](https://github.com/ory/keto/commit/04f52231509dd0f3a57d745918fc43fff7c595ff)),
  closes [#99](https://github.com/ory/keto/issues/99)
- Add tracing support and general improvements (#98)
  ([63b3946](https://github.com/ory/keto/commit/63b3946e0ae1fa23c6a359e9a64b296addff868c)),
  closes [#98](https://github.com/ory/keto/issues/98): This patch improves the
  internal configuration and service management. It adds support for distributed
  tracing and resolves several issues in the release pipeline and CLI.
  Additionally, composable docker-compose configuration files have been added.
  Several bugs have been fixed in the release management pipeline.
- Add content-type in the response of allowed
  ([#90](https://github.com/ory/keto/issues/90))
  ([39a1486](https://github.com/ory/keto/commit/39a1486dc53456189d30380460a9aeba198fa9e9))
- Fix disable-telemetry check ([#85](https://github.com/ory/keto/issues/85))
  ([38b5383](https://github.com/ory/keto/commit/38b538379973fa34bd2bf24dcb2e6dbedf324e1e))
- Fix remove member from role ([#87](https://github.com/ory/keto/issues/87))
  ([698e161](https://github.com/ory/keto/commit/698e161989331ca5a3a0769301d9694ef805a876)),
  closes [#74](https://github.com/ory/keto/issues/74)
- Fix the type of conditions in the policy
  ([#86](https://github.com/ory/keto/issues/86))
  ([fc1ced6](https://github.com/ory/keto/commit/fc1ced63bd39c9fbf437e419dfc384343e36e0ee))
- Move Go SDK generation to go-swagger
  ([#94](https://github.com/ory/keto/issues/94))
  ([9f48a95](https://github.com/ory/keto/commit/9f48a95187a7b6160108cd7d0301590de2e58f07)),
  closes [#92](https://github.com/ory/keto/issues/92)
- Send 403 when authorization result is negative
  ([#93](https://github.com/ory/keto/issues/93))
  ([de806d8](https://github.com/ory/keto/commit/de806d892819db63c1abc259ab06ee08d87895dc)),
  closes [#75](https://github.com/ory/keto/issues/75)
- Update dependencies ([#91](https://github.com/ory/keto/issues/91))
  ([4d44174](https://github.com/ory/keto/commit/4d4417474ebf8cc69d01e5ac82633b966cdefbc7))
- storage/memory: Fix upsert with pre-existing key will causes duplicate records
  (#88)
  ([1cb8a36](https://github.com/ory/keto/commit/1cb8a36a08883b785d9bb0a4be1ddc00f1f9d358)),
  closes [#88](https://github.com/ory/keto/issues/88)
  [#80](https://github.com/ory/keto/issues/80)

# [0.2.3-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.2.2-sandbox+oryOS.10...v0.2.3-sandbox+oryOS.10) (2019-02-05)

dist: Fix packr build pipeline (#84)

Closes #73 Closes #81

Signed-off-by: aeneasr <aeneas@ory.sh>

### Documentation

- Add documentation for glob matching
  ([5c8babb](https://github.com/ory/keto/commit/5c8babbfbae01a78f30cfbff92d8e9c3a6b09027))
- Incorporates changes from version v0.2.2-sandbox+oryOS.10
  ([ed7af3f](https://github.com/ory/keto/commit/ed7af3fa4e5d1d0d03b5366f4cf865a5b82ec293))
- Properly generate api.swagger.json
  ([18e3f84](https://github.com/ory/keto/commit/18e3f84cdeee317f942d61753399675c98886e5d))

### Unclassified

- Add placeholder go file for rego inclusion
  ([6a6f64d](https://github.com/ory/keto/commit/6a6f64d8c59b496f6cf360f55eba1e16bf5380f1))
- Add support for glob matching
  ([bb76c6b](https://github.com/ory/keto/commit/bb76c6bebe522fc25448c4f4e4d1ef7c530a725f))
- Ex- and import rego subdirectories for `go get`
  [#77](https://github.com/ory/keto/issues/77)
  ([59cc053](https://github.com/ory/keto/commit/59cc05328f068fc3046b2dbc022a562fd5d67960)),
  closes [#73](https://github.com/ory/keto/issues/73)
- Fix packr build pipeline ([#84](https://github.com/ory/keto/issues/84))
  ([65a87d5](https://github.com/ory/keto/commit/65a87d564d237bc979bb5962beff7d3703d9689f)),
  closes [#73](https://github.com/ory/keto/issues/73)
  [#81](https://github.com/ory/keto/issues/81)
- Import glob in rego/doc.go
  ([7798442](https://github.com/ory/keto/commit/7798442553cfe7989a23d2c389c8c63a24013543))
- Properly handle dbal error
  ([6811607](https://github.com/ory/keto/commit/6811607ea79c8f3155a17bc1aea566e9e4680616))
- Properly handle TLS certificates if set
  ([36399f0](https://github.com/ory/keto/commit/36399f09261d4f3cb5e053679eee3cb15da2df19)),
  closes [#73](https://github.com/ory/keto/issues/73)

# [0.2.2-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.2.1-sandbox+oryOS.10...v0.2.2-sandbox+oryOS.10) (2018-12-13)

ci: Fix docker push arguments in publish task

Signed-off-by: aeneasr <aeneas@ory.sh>

### Continuous Integration

- Fix docker push arguments in publish task
  ([f03c77c](https://github.com/ory/keto/commit/f03c77c6b7461ab81cb03265cbec909ac45c2259))

# [0.2.1-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.2.0-sandbox+oryOS.10...v0.2.1-sandbox+oryOS.10) (2018-12-13)

ci: Fix docker release task

Signed-off-by: aeneasr <aeneas@ory.sh>

### Continuous Integration

- Fix docker release task
  ([7a0414f](https://github.com/ory/keto/commit/7a0414f614b6cc8b1d78cfbb773a2f0192d00d23))

# [0.2.0-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.0.1...v0.2.0-sandbox+oryOS.10) (2018-12-13)

all: gofmt

Signed-off-by: aeneasr <aeneas@ory.sh>

### Documentation

- Adds banner
  ([0ec1d8f](https://github.com/ory/keto/commit/0ec1d8f5e843465d17ac4c8f91f18e5badf16900))
- Adds GitHub templates & code of conduct
  ([#31](https://github.com/ory/keto/issues/31))
  ([a11e898](https://github.com/ory/keto/commit/a11e8980f2af528f1357659966123d0cbf7d43db))
- Adds link to examples repository
  ([#32](https://github.com/ory/keto/issues/32))
  ([7061a2a](https://github.com/ory/keto/commit/7061a2aa31652a9e0c2d449facb1201bfa11fd3f))
- Adds security console image
  ([fd27fc9](https://github.com/ory/keto/commit/fd27fc9cce50beb3d0189e0a93300879fd7149db))
- Changes hydra to keto in readme
  ([9dab531](https://github.com/ory/keto/commit/9dab531744cf5b0ae98862945d44b07535595781))
- Deprecate old versions in logs
  ([955d647](https://github.com/ory/keto/commit/955d647307a48ee7cf2d3f9fb4263072adf42299))
- Incorporates changes from version
  ([85c4d81](https://github.com/ory/keto/commit/85c4d81a192e92f874c106b91cfa6fb404d9a34a))
- Incorporates changes from version v0.0.0-testrelease.1
  ([6062dd4](https://github.com/ory/keto/commit/6062dd4a894607f5f1ead119af20cc8bdbe15bef))
- Incorporates changes from version v0.0.1-1-g85c4d81
  ([f4606fc](https://github.com/ory/keto/commit/f4606fce0326bece2a89dadc029bc5ce9778df18))
- Incorporates changes from version v0.0.1-11-g114914f
  ([92a4dca](https://github.com/ory/keto/commit/92a4dca7a41dcf3a88c4800bf6d2217f33cfcdd1))
- Incorporates changes from version v0.0.1-16-g7d8a8ad
  ([2b76a83](https://github.com/ory/keto/commit/2b76a83755153b3f8a2b8d28c5b0029d96d567b6))
- Incorporates changes from version v0.0.1-18-g099e7e0
  ([70b12ad](https://github.com/ory/keto/commit/70b12adf5bcc0e890d6707e11e891e6cedfb3d87))
- Incorporates changes from version v0.0.1-20-g97ccbe6
  ([b21d56e](https://github.com/ory/keto/commit/b21d56e599c7eb4c1769bc18878f7d5818b73023))
- Incorporates changes from version v0.0.1-30-gaf2c3b5
  ([a1d0dcc](https://github.com/ory/keto/commit/a1d0dcc78a9506260f86df00e4dff8ab02909ce1))
- Incorporates changes from version v0.0.1-32-gedb5a60
  ([a5c369a](https://github.com/ory/keto/commit/a5c369a90da67c96bbde60e673c67f50b841fadd))
- Incorporates changes from version v0.0.1-6-g570783e
  ([0fcbbcb](https://github.com/ory/keto/commit/0fcbbcb02f1d748f9c733c86368b223b2ee4c6e2))
- Incorporates changes from version v0.0.1-7-g0fcbbcb
  ([c0141a8](https://github.com/ory/keto/commit/c0141a8ec22ea1260bf2d45d72dfe06737ec0115))
- Incorporates changes from version v0.1.0-sandbox
  ([9ee0664](https://github.com/ory/keto/commit/9ee06646d2cfb2d69abdcc411e31d14957437a1e))
- Incorporates changes from version v1.0.0-beta.1-1-g162d7b8
  ([647c5a9](https://github.com/ory/keto/commit/647c5a9e1bc8d9d635bf6f2511c3faa9a9daefef))
- Incorporates changes from version v1.0.0-beta.2-11-g2b280bb
  ([936889d](https://github.com/ory/keto/commit/936889d760f04a03d498f65331d653cbad3702d0))
- Incorporates changes from version v1.0.0-beta.2-13-g382e1d3
  ([883df44](https://github.com/ory/keto/commit/883df44a922f3daee86597af467072555cadc7e7))
- Incorporates changes from version v1.0.0-beta.2-15-g74450da
  ([48dd9f1](https://github.com/ory/keto/commit/48dd9f1ffbeaa99ac8dc27085c5a50f9244bf9c3))
- Incorporates changes from version v1.0.0-beta.2-3-gf623c52
  ([b6b90e5](https://github.com/ory/keto/commit/b6b90e5b2180921f78064a60666704b4e72679b6))
- Incorporates changes from version v1.0.0-beta.2-5-g3852be5
  ([3f09090](https://github.com/ory/keto/commit/3f09090a2f82f3f29154c19217cea0a10d65ea3a))
- Incorporates changes from version v1.0.0-beta.2-9-gc785187
  ([4c30a3c](https://github.com/ory/keto/commit/4c30a3c0ad83ba80e1857b41211e7ddade06c4cf))
- Incorporates changes from version v1.0.0-beta.3-1-g06adbf1
  ([0ba3c06](https://github.com/ory/keto/commit/0ba3c0674832b641ef5e0c3f0d60d81ed3a647b2))
- Incorporates changes from version v1.0.0-beta.3-10-g9994967
  ([d2345ca](https://github.com/ory/keto/commit/d2345ca3beb354d6ee7c7926c1a5ddb425d6b405))
- Incorporates changes from version v1.0.0-beta.3-12-gc28b521
  ([b4d792f](https://github.com/ory/keto/commit/b4d792f74055853f05ca46c67625ffd432fc74fd))
- Incorporates changes from version v1.0.0-beta.3-3-g9e16605
  ([c43bf2b](https://github.com/ory/keto/commit/c43bf2b5232bed9106dd47d7eb53d2f93bfe260d))
- Incorporates changes from version v1.0.0-beta.3-5-ga11e898
  ([b9d9b8e](https://github.com/ory/keto/commit/b9d9b8ee33ab957f43f99c427a88ade847e79ed0))
- Incorporates changes from version v1.0.0-beta.3-8-g7061a2a
  ([d76ff9d](https://github.com/ory/keto/commit/d76ff9dc9a4c8a8f1286eeb139d8f5af9617f421))
- Incorporates changes from version v1.0.0-beta.5
  ([0dc314c](https://github.com/ory/keto/commit/0dc314c7888020b40e12eb59fd77135044fd063b))
- Incorporates changes from version v1.0.0-beta.6-1-g5e97104
  ([f14c8ed](https://github.com/ory/keto/commit/f14c8edd7204a811e333ea84429cf837b4e7d27b))
- Incorporates changes from version v1.0.0-beta.8
  ([5045b59](https://github.com/ory/keto/commit/5045b59e2a83d6ab047b1b95c581d7c34e96a2e0))
- Incorporates changes from version v1.0.0-beta.9
  ([be2f035](https://github.com/ory/keto/commit/be2f03524721ef47ecb1c9aec57c2696174e0657))
- Properly sets up changelog TOC
  ([e0acd67](https://github.com/ory/keto/commit/e0acd670ab19c0d6fd36733fea164e2b0414597d))
- Puts toc in the right place
  ([114914f](https://github.com/ory/keto/commit/114914fa354f784b310bc9dfd232a011e0d98d99))
- Revert changes from test release
  ([ab3a64d](https://github.com/ory/keto/commit/ab3a64d3d41292364c5947db98c4d27a8223853e))
- Update documentation links ([#67](https://github.com/ory/keto/issues/67))
  ([d22d413](https://github.com/ory/keto/commit/d22d413c7a001ccaa96b4c013665153f41831614))
- Update link to security console
  ([846ce4b](https://github.com/ory/keto/commit/846ce4baa9da5954bd30996f489885a026c48185))
- Update migration guide
  ([3c44b58](https://github.com/ory/keto/commit/3c44b58613e46ed39d42463537773fe9d95a54da))
- Update to latest changes
  ([1625123](https://github.com/ory/keto/commit/1625123ed342f019d5e7ab440eb37da310570842))
- Updates copyright notice
  ([9dd5578](https://github.com/ory/keto/commit/9dd557825dfd3b9d589c9db2ccb201638debbaae))
- Updates installation guide
  ([f859645](https://github.com/ory/keto/commit/f859645f230f405cfabed0c1b9a2b67b1a3841d3))
- Updates issue and pull request templates
  ([#52](https://github.com/ory/keto/issues/52))
  ([941cae6](https://github.com/ory/keto/commit/941cae6fee058f68eabbbf4dd9cafad4760e108f))
- Updates issue and pull request templates
  ([#53](https://github.com/ory/keto/issues/53))
  ([7b222d2](https://github.com/ory/keto/commit/7b222d285e74c0db482136b23f37072216b3acb0))
- Updates issue and pull request templates
  ([#54](https://github.com/ory/keto/issues/54))
  ([f098639](https://github.com/ory/keto/commit/f098639b5e748151810848fdd3173e0246bc03dc))
- Updates link to guide and header
  ([437c255](https://github.com/ory/keto/commit/437c255ecfff4127fb586cc069e07f86988ad1ba))
- Updates link to open collective
  ([382e1d3](https://github.com/ory/keto/commit/382e1d34c7da0ba0447b78506a749bd7f0085f48))
- Updates links to docs
  ([d84be3b](https://github.com/ory/keto/commit/d84be3b6a8e5eb284ec3fb137ee774ba5ee0d529))
- Updates newsletter link in README
  ([2dc36b2](https://github.com/ory/keto/commit/2dc36b21c8af8e3e39f093198715ea24b65d65af))

### Unclassified

- Add Go SDK factory
  ([99db7e6](https://github.com/ory/keto/commit/99db7e6d4edac88794266a01ddfab9cd0632e95a))
- Add go SDK interface
  ([3dd5f7d](https://github.com/ory/keto/commit/3dd5f7d61bb460c34744b84a34755bfb8219b304))
- Add health handlers
  ([bddb949](https://github.com/ory/keto/commit/bddb949459d05002b0f8882d981e4f63fdddf25f))
- Add policy list handler
  ([a290619](https://github.com/ory/keto/commit/a290619d01d15eb8e3b4e33ede1058d316ee807a))
- Add role iterator in list handler
  ([a3eb696](https://github.com/ory/keto/commit/a3eb6961783f7b562f0a0d0a7e2819bffebce5b8))
- Add SDK generation to circle ci
  ([9b37165](https://github.com/ory/keto/commit/9b37165873bcb0cc5dc60d2514d9824a073466a1))
- Adds ability to update a role using PUT
  ([#14](https://github.com/ory/keto/issues/14))
  ([97ccbe6](https://github.com/ory/keto/commit/97ccbe6d808823c56901ad237878aa6d53cddeeb)):

  - transfer UpdateRoleMembers from https://github.com/ory/hydra/pull/768 to
    keto

  - fix tests by using right http method & correcting sql request

  - Change behavior to overwrite the whole role instead of just the members.

  * small sql migration fix

- Adds log message when telemetry is active
  ([f623c52](https://github.com/ory/keto/commit/f623c52655ff85b7f7209eb73e94eb66a297c5b7))
- Clean up vendor dependencies
  ([9a33c23](https://github.com/ory/keto/commit/9a33c23f4d37ab88b4d643fd79204334d73404c6))
- Do not split empty scope ([#45](https://github.com/ory/keto/issues/45))
  ([b29cf8c](https://github.com/ory/keto/commit/b29cf8cc92607e13457dba8331f5c9286054c8c1))
- Fix typo in help command in env var name
  ([#39](https://github.com/ory/keto/issues/39))
  ([8a5016c](https://github.com/ory/keto/commit/8a5016cd75be78bb42a9a38bfd453ad5722db9db)),
  closes [#25](https://github.com/ory/keto/issues/25)
- Fixes environment variable typos
  ([566d588](https://github.com/ory/keto/commit/566d588e4fca12399966718b725fe4461a28e51e))
- Fixes typo in help command
  ([74450da](https://github.com/ory/keto/commit/74450da18a27513820328c28f72203653c664367)),
  closes [#25](https://github.com/ory/keto/issues/25)
- Format code
  ([637c78c](https://github.com/ory/keto/commit/637c78cba697682b544473a3af9b6ae7715561aa))
- Gofmt
  ([a8d7f9f](https://github.com/ory/keto/commit/a8d7f9f546ae2f3b8c3fa643d8e19b68ca26cc67))
- Improve compose documentation
  ([6870443](https://github.com/ory/keto/commit/68704435f3c299b853f4ff5cacae285b09ada3b5))
- Improves usage of metrics middleware
  ([726c4be](https://github.com/ory/keto/commit/726c4bedfc3f02fdac380930e32f37c251e51aa4))
- Improves usage of metrics middleware
  ([301f386](https://github.com/ory/keto/commit/301f38605af66abae4d28ed0cac90d0b82b655c4))
- Introduce docker-compose file for testing
  ([ba857e3](https://github.com/ory/keto/commit/ba857e3859966e857c5a741825411575e17446de))
- Introduces health and version endpoints
  ([6a9da74](https://github.com/ory/keto/commit/6a9da74f693ee6c15a775ab8d652582aea093601))
- List roles from keto_role table ([#28](https://github.com/ory/keto/issues/28))
  ([9e16605](https://github.com/ory/keto/commit/9e166054b8d474fbce6983d5d00eeeb062fc79b1))
- Properly names flags
  ([af2c3b5](https://github.com/ory/keto/commit/af2c3b5bc96e95fb31b1db5c7fe6dfd6b6fc5b20))
- Require explicit CORS enabling ([#42](https://github.com/ory/keto/issues/42))
  ([9a45107](https://github.com/ory/keto/commit/9a45107af304b2a8e663a532e4f6e4536f15888c))
- Update dependencies
  ([663d8b1](https://github.com/ory/keto/commit/663d8b13e99694a57752cd60a68342b81b041c66))
- Switch to rego as policy decision engine (#48)
  ([ee9bcf2](https://github.com/ory/keto/commit/ee9bcf2719178e5a8dccca083a90313947a8a63b)),
  closes [#48](https://github.com/ory/keto/issues/48)
- Update hydra to v1.0.0-beta.6 ([#35](https://github.com/ory/keto/issues/35))
  ([5e97104](https://github.com/ory/keto/commit/5e971042afff06e2a6ee3b54d2fea31687203623))
- Update npm package registry
  ([a53d3d2](https://github.com/ory/keto/commit/a53d3d23e11fde5dcfbb27a2add1049f4d8e10e6))
- Enable TLS option to serve API (#46)
  ([2f62063](https://github.com/ory/keto/commit/2f620632d0375bf9c7e58dbfb49627c02c66abf3)),
  closes [#46](https://github.com/ory/keto/issues/46)
- Make introspection authorization optional
  ([e5460ad](https://github.com/ory/keto/commit/e5460ad884cd018cd6177324b949cd66bfd53bc7))
- Properly output telemetry information
  ([#33](https://github.com/ory/keto/issues/33))
  ([9994967](https://github.com/ory/keto/commit/9994967b0ca54a62b8b0088fe02be9e890d9574b))
- Remove ORY Hydra dependency ([#44](https://github.com/ory/keto/issues/44))
  ([d487344](https://github.com/ory/keto/commit/d487344fe7e07cb6370371c6b0b6cf3cca767ed1))
- Resolves an issue with the hydra migrate command
  ([2b280bb](https://github.com/ory/keto/commit/2b280bb57c9073a9c8384cde0b14a6991cfacdb6)),
  closes [#23](https://github.com/ory/keto/issues/23)
- Upgrade superagent version ([#41](https://github.com/ory/keto/issues/41))
  ([9c80dbc](https://github.com/ory/keto/commit/9c80dbcc1cc63243839b58ca56ac9be104797887))
- gofmt
  ([777b1be](https://github.com/ory/keto/commit/777b1be1378d314e7cfde0c34450afcce7e590a5))
- Updates README.md (#34)
  ([c28b521](https://github.com/ory/keto/commit/c28b5219fd64314a75ee3c848a80a0c5974ebb7d)),
  closes [#34](https://github.com/ory/keto/issues/34)
- Properly parses cors options
  ([edb5a60](https://github.com/ory/keto/commit/edb5a600f2ce16c0847ee5ef399fa5a41b1e736a))
- Removes additional output if no args are passed
  ([703e124](https://github.com/ory/keto/commit/703e1246ce0fd89066b497c45f0c6cadeb06c331))
- Resolves broken role test
  ([b6c7f9c](https://github.com/ory/keto/commit/b6c7f9c33c4c1f43164d6da0ec7f2553f1f4c598))
- Resolves minor typos and updates install guide
  ([3852be5](https://github.com/ory/keto/commit/3852be56cb81df966a85d4c828de0397d9e74768))
- Updates to latest sqlcon
  ([2c9f643](https://github.com/ory/keto/commit/2c9f643042ff4edffae8bd41834d2a57c923871c))
- Use roles in warden decision
  ([c785187](https://github.com/ory/keto/commit/c785187e31fc7a4b8b762a5e27fac66dcaa97513)),
  closes [#21](https://github.com/ory/keto/issues/21)
  [#19](https://github.com/ory/keto/issues/19)
- authn/client: Payload is now prefixed with client
  ([8584d94](https://github.com/ory/keto/commit/8584d94cfb18deb37ae32ae601f4cd15c14067e7))

# [0.0.1](https://github.com/ory/keto/compare/4f00bc96ece3180a888718ec3c41c69106c86f56...v0.0.1) (2018-05-20)

authn: Checks token_type is "access_token", if set

Closes #1

### Documentation

- Incorporates changes from version
  ([b5445a0](https://github.com/ory/keto/commit/b5445a0fc5b6f813cd1731b20c8c5c79d7c4cdf8))
- Incorporates changes from version
  ([295ff99](https://github.com/ory/keto/commit/295ff998af55777823b04f423e365fd58e61753b))
- Incorporates changes from version
  ([bd44d41](https://github.com/ory/keto/commit/bd44d41b2781e33353082397c47390a27f749e16))
- Updates readme and upgrades
  ([0f95dbb](https://github.com/ory/keto/commit/0f95dbb967fd17b607caa999ae30453f5f599739))
- Uses keto repo for changelog
  ([14c0b2a](https://github.com/ory/keto/commit/14c0b2a2bd31566f2b9048831f894aba05c5b15d))

### Unclassified

- Adds migrate commands to the proper parent command
  ([231c70d](https://github.com/ory/keto/commit/231c70d816b0736a51eddc1fa0445bac672b1b2f))
- Checks token_type is "access_token", if set
  ([d2b8f5d](https://github.com/ory/keto/commit/d2b8f5d313cce597566bd18e4f3bea4a423a62ee)),
  closes [#1](https://github.com/ory/keto/issues/1)
- Removes old test
  ([07b733b](https://github.com/ory/keto/commit/07b733bfae4b733e3e2124545b92c537dabbdcf0))
- Renames subject to sub in response payloads
  ([ca4d540](https://github.com/ory/keto/commit/ca4d5408000be2b896d38eaaf5e67a3fc0a566da))
- Tells linguist to ignore SDK files
  ([f201eb9](https://github.com/ory/keto/commit/f201eb95f3309a60ac50f42cfba0bae2e38e8d13))
- Retries SQL connection on migrate commands
  ([3d33d73](https://github.com/ory/keto/commit/3d33d73c009077c5bf30ae4b03802904bfb5d5b2)):
  This patch also introduces a fatal error if migrations fail
- cmd/server: Resolves DBAL not handling postgres properly
  ([dedc32a](https://github.com/ory/keto/commit/dedc32ab218923243b1955ce5bcbbdc5cc416953))
- cmd/server: Improves error message in migrate command
  ([4b17ce8](https://github.com/ory/keto/commit/4b17ce8848113cae807840182d1a318190c2a9b3))
- Resolves travis and docker issues
  ([6f4779c](https://github.com/ory/keto/commit/6f4779cc51bf4f2ee5b97541fb77d8f882497710))
- Adds OAuth2 Client Credentials authenticator and warden endpoint
  ([c55139b](https://github.com/ory/keto/commit/c55139b51e636834759706499a2aec1451f4fbd9))
- Adds SDK helpers
  ([a1c2608](https://github.com/ory/keto/commit/a1c260801d9366fccf4bfb4fc64b2c67fc594565))
- Resolves SDK and test issues (#4)
  ([2d4cd98](https://github.com/ory/keto/commit/2d4cd9805af3081bbcbea3f806ca066d35385a4b)),
  closes [#4](https://github.com/ory/keto/issues/4)
- Initial project commit
  ([a592e51](https://github.com/ory/keto/commit/a592e5126f130f8b673fff6c894fdbd9fb56f81c))
- Initial commit
  ([4f00bc9](https://github.com/ory/keto/commit/4f00bc96ece3180a888718ec3c41c69106c86f56))

---

id: changelog title: Changelog custom_edit_url: null

---

# [Unreleased](https://github.com/ory/keto/compare/v0.6.0-alpha.3...8e301198298858fd7f387ef63a7abf4fa55ea240) (2021-06-22)

### Bug Fixes

- Add missing tracers ([#600](https://github.com/ory/keto/issues/600))
  ([aa263be](https://github.com/ory/keto/commit/aa263be9a7830e3c769d7698d36137555ca230bc)),
  closes [#593](https://github.com/ory/keto/issues/593)
- Handle relation tuple cycles in expand and check engine
  ([#623](https://github.com/ory/keto/issues/623))
  ([8e30119](https://github.com/ory/keto/commit/8e301198298858fd7f387ef63a7abf4fa55ea240))
- Log all database connection errors
  ([#588](https://github.com/ory/keto/issues/588))
  ([2b0fad8](https://github.com/ory/keto/commit/2b0fad897e61400bd2a6cdf47f33ff4301e9c5f8))
- Move gRPC client module root up
  ([#620](https://github.com/ory/keto/issues/620))
  ([3b881f6](https://github.com/ory/keto/commit/3b881f6015a93b382b3fbbca4be9259622038b6a)):

  BREAKING: The npm package `@ory/keto-grpc-client` from now on includes all API
  versions. Because of that, the import paths changed. For migrating to the new
  client package, change the import path according to the following example:

  ```diff
  - import acl from '@ory/keto-grpc-client/acl_pb.js'
  + // from the latest version
  + import { acl } from '@ory/keto-grpc-client'
  + // or a specific one
  + import acl from '@ory/keto-grpc-client/ory/keto/acl/v1alpha1/acl_pb.js'
  ```

- Update docker-compose.yml version
  ([#595](https://github.com/ory/keto/issues/595))
  ([7fa4dca](https://github.com/ory/keto/commit/7fa4dca4182a1fa024f9cef0a04163f2cbd882aa)),
  closes [#549](https://github.com/ory/keto/issues/549)

### Documentation

- Fix example not following best practice
  ([#582](https://github.com/ory/keto/issues/582))
  ([a015818](https://github.com/ory/keto/commit/a0158182c5f87cfd4767824e1c5d6cbb8094a4e6))
- Update NPM links due to organisation move
  ([#616](https://github.com/ory/keto/issues/616))
  ([6355bea](https://github.com/ory/keto/commit/6355beae5b5b28c3eee19fdee85b9875cbc165c3))

### Features

- Make generated gRPC client its own module
  ([#583](https://github.com/ory/keto/issues/583))
  ([f0fbb64](https://github.com/ory/keto/commit/f0fbb64b3358e9800854295cebc9ec8b8e56c87a))
- Max_idle_conn_time ([#605](https://github.com/ory/keto/issues/605))
  ([50a8623](https://github.com/ory/keto/commit/50a862338e17f86900ca162da7f3467f55f9f954)),
  closes [#523](https://github.com/ory/keto/issues/523)

# [0.6.0-alpha.3](https://github.com/ory/keto/compare/v0.6.0-alpha.2...v0.6.0-alpha.3) (2021-04-29)

Resolves CRDB and build issues.

### Code Generation

- Pin v0.6.0-alpha.3 release commit
  ([d766968](https://github.com/ory/keto/commit/d766968419d10a68fd843df45316e3436b68d61d))

# [0.6.0-alpha.2](https://github.com/ory/keto/compare/v0.6.0-alpha.1...v0.6.0-alpha.2) (2021-04-29)

This release improves stability and documentation.

### Bug Fixes

- Add npm run format to make format
  ([7d844a8](https://github.com/ory/keto/commit/7d844a8e6412ae561963b97ac26d4682411095d4))
- Makefile target
  ([0e6f612](https://github.com/ory/keto/commit/0e6f6122de7bdbb691ad7cc236b6bc9a3601d39e))
- Move swagger to spec dir
  ([7f6a061](https://github.com/ory/keto/commit/7f6a061aafda275d278bf60f16e90039da45bc57))
- Resolve clidoc issues
  ([ef12b4e](https://github.com/ory/keto/commit/ef12b4e267f34fbf9709fe26023f9b7ae6670c24))
- Update install.sh ([#568](https://github.com/ory/keto/issues/568))
  ([86ab245](https://github.com/ory/keto/commit/86ab24531d608df0b5391ee8ec739291b9a90e20))
- Use correct id
  ([5e02902](https://github.com/ory/keto/commit/5e029020b5ba3931f15d343cf6a9762b064ffd45))
- Use correct id for api
  ([32a6b04](https://github.com/ory/keto/commit/32a6b04609054cba84f7b56ebbe92341ec5dcd98))
- Use sqlite image versions ([#544](https://github.com/ory/keto/issues/544))
  ([ec6cc5e](https://github.com/ory/keto/commit/ec6cc5ed528f1a097ea02669d059e060b7eff824))

### Code Generation

- Pin v0.6.0-alpha.2 release commit
  ([470b2c6](https://github.com/ory/keto/commit/470b2c61c649fe5fcf638c84d4418212ff0330a5))

### Documentation

- Add gRPC client README.md ([#559](https://github.com/ory/keto/issues/559))
  ([9dc3596](https://github.com/ory/keto/commit/9dc35969ada8b0d4d73dee9089c4dc61cd9ea657))
- Change forum to discussions readme
  ([#539](https://github.com/ory/keto/issues/539))
  ([ea2999d](https://github.com/ory/keto/commit/ea2999d4963316810a8d8634fcd123bda31eaa8f))
- Fix cat videos example docker compose
  ([#549](https://github.com/ory/keto/issues/549))
  ([b25a711](https://github.com/ory/keto/commit/b25a7114631957935c71ac6a020ab6bd0c244cd7))
- Fix typo ([#538](https://github.com/ory/keto/issues/538))
  ([99a9693](https://github.com/ory/keto/commit/99a969373497792bb4cd8ff62bf5245087517737))
- Include namespace in olymp library example
  ([#540](https://github.com/ory/keto/issues/540))
  ([135e814](https://github.com/ory/keto/commit/135e8145c383a76b494b469253c949c38f4414a7))
- Update install from source steps to actually work
  ([#548](https://github.com/ory/keto/issues/548))
  ([e662256](https://github.com/ory/keto/commit/e6622564f58b7612b13b11b54e75a7350f52d6de))

### Features

- Global docs sidebar and added cloud pages
  ([c631c82](https://github.com/ory/keto/commit/c631c82b7ff3d12734869ac22730b52e73dcf287))
- Support retryable CRDB transactions
  ([833147d](https://github.com/ory/keto/commit/833147dae40e9ac5bdf220f8aa3f01abd444f791))

# [0.6.0-alpha.1](https://github.com/ory/keto/compare/v0.5.6-alpha.1...v0.6.0-alpha.1) (2021-04-07)

We are extremely happy to announce next-gen Ory Keto which implements
[Zanzibar: Google’s Consistent, Global Authorization System](https://research.google/pubs/pub48190/):

> Zanzibar provides a uniform data model and configuration language for
> expressing a wide range of access control policies from hundreds of client
> services at Google, including Calendar, Cloud, Drive, Maps, Photos, and
> YouTube. Its authorization decisions respect causal ordering of user actions
> and thus provide external consistency amid changes to access control lists and
> object contents. Zanzibar scales to trillions of access control lists and
> millions of authorization requests per second to support services used by
> billions of people. It has maintained 95th-percentile latency of less than 10
> milliseconds and availability of greater than 99.999% over 3 years of
> production use.

Ory Keto is the first open source planet-scale authorization system built with
cloud native technologies (Go, gRPC, newSQL) and architecture. It is also the
first open source implementation of Google Zanzibar :tada:!

Many concepts developer by Google Zanzibar are implemented in Ory Keto already.
Let's take a look!

As of this release, Ory Keto knows how to interpret and operate on the basic
access control lists known as relation tuples. They encode relations between
objects and subjects. One simple example of such a relation tuple could encode
"`user1` has access to file `/foo`", a more complex one could encode "everyone
who has write access on `/foo` has read access on `/foo`".

Ory Keto comes with all the basic APIs as described in the Zanzibar paper. All
of them are available over gRPC and REST.

1. List: query relation tuples
2. Check: determine whether a subject has a relation on an object
3. Expand: get a tree of all subjects who have a relation on an object
4. Change: create, update, and delete relation tuples

For all details, head over to the
[documentation](https://www.ory.sh/keto/docs/concepts/api-overview).

With this release we officially move the "old" Keto to the
[legacy-0.5 branch](https://github.com/ory/keto/tree/legacy-0.5). We will only
provide security fixes from now on. A migration path to v0.6 is planned but not
yet implemented, as the architectures are vastly different. Please refer to
[the issue](https://github.com/ory/keto/issues/318).

We are keen to bring more features and performance improvements. The next
features we will tackle are:

- Subject Set rewrites
- Native ABAC & RBAC Support
- Integration with other policy servers
- Latency reduction through aggressive caching
- Cluster mode that fans out requests over all Keto instances

So stay tuned, :star: this repo, :eyes: releases, and
[subscribe to our newsletter :email:](https://ory.us10.list-manage.com/subscribe?u=ffb1a878e4ec6c0ed312a3480&id=f605a41b53&MERGE0=&group[17097][32]=1).

### Bug Fixes

- Add description attribute to access control policy role
  ([#215](https://github.com/ory/keto/issues/215))
  ([831eba5](https://github.com/ory/keto/commit/831eba59f810ca68561dd584c9df7684df10b843))
- Add leak_sensitive_values to config schema
  ([2b21d2b](https://github.com/ory/keto/commit/2b21d2bdf4ca9523d16159c5f73c4429b692e17d))
- Bump CLI
  ([80c82d0](https://github.com/ory/keto/commit/80c82d026cbfbab8fbb84d850d8980866ecf88df))
- Bump deps and replace swagutil
  ([#212](https://github.com/ory/keto/issues/212))
  ([904258d](https://github.com/ory/keto/commit/904258d23959c3fa96b6d8ccfdb79f6788c106ec))
- Check engine overwrote result in some cases
  ([#412](https://github.com/ory/keto/issues/412))
  ([3404492](https://github.com/ory/keto/commit/3404492002ca5c3f017ef25486e377e911987aa4))
- Check health status in status command
  ([21c64d4](https://github.com/ory/keto/commit/21c64d45f21a505744b9f70d780f9b3079d3822c))
- Check REST API returns JSON object
  ([#460](https://github.com/ory/keto/issues/460))
  ([501dcff](https://github.com/ory/keto/commit/501dcff4427f76902671f6d5733f28722bd51fa7)),
  closes [#406](https://github.com/ory/keto/issues/406)
- Empty relationtuple list should not error
  ([#440](https://github.com/ory/keto/issues/440))
  ([fbcb3e1](https://github.com/ory/keto/commit/fbcb3e1f337b5114d7697fa512ded92b5f409ef4))
- Ensure nil subject is not allowed
  ([#449](https://github.com/ory/keto/issues/449))
  ([7a0fcfc](https://github.com/ory/keto/commit/7a0fcfc4fe83776fa09cf78ee11f407610554d04)):

  The nodejs gRPC client was a great fuzzer and pointed me to some nil pointer
  dereference panics. This adds some input validation to prevent panics.

- Ensure persister errors are handled by sqlcon
  ([#473](https://github.com/ory/keto/issues/473))
  ([4343c4a](https://github.com/ory/keto/commit/4343c4acd8f917fb7ae131e67bca6855d4d61694))
- Handle pagination and errors in the check/expand engines
  ([#398](https://github.com/ory/keto/issues/398))
  ([5eb1a7d](https://github.com/ory/keto/commit/5eb1a7d49af6b43707c122de8727cbd72285cb5c))
- Ignore dist
  ([ba816ea](https://github.com/ory/keto/commit/ba816ea2ca39962f02c08e0c7b75cfe3cf1d963d))
- Ignore x/net false positives
  ([d8b36cb](https://github.com/ory/keto/commit/d8b36cb1812abf7265ac15c29780222be025186b))
- Improve CLI remote sourcing ([#474](https://github.com/ory/keto/issues/474))
  ([a85f4d7](https://github.com/ory/keto/commit/a85f4d7470ac3744476e82e5889b97d5a0680473))
- Improve handlers and add tests
  ([#470](https://github.com/ory/keto/issues/470))
  ([ca5ccb9](https://github.com/ory/keto/commit/ca5ccb9c237fdcc4db031ec97a75616a859cbf8f))
- Insert relation tuples without fmt.Sprintf
  ([#443](https://github.com/ory/keto/issues/443))
  ([fe507bb](https://github.com/ory/keto/commit/fe507bb4ea719780e732d098291aa190d6b1c441))
- Minor bugfixes ([#371](https://github.com/ory/keto/issues/371))
  ([185ee1e](https://github.com/ory/keto/commit/185ee1e51bc4bcdee028f71fcaf207b7e342313b))
- Move dockerfile to where it belongs
  ([f087843](https://github.com/ory/keto/commit/f087843ac8f24e741bf39fe65ee5b0a9adf9a5bb))
- Namespace migrator ([#417](https://github.com/ory/keto/issues/417))
  ([ea79300](https://github.com/ory/keto/commit/ea7930064f490b063a712b4e18521f8996931a13)),
  closes [#404](https://github.com/ory/keto/issues/404)
- Remove SQL logging ([#455](https://github.com/ory/keto/issues/455))
  ([d8e2a86](https://github.com/ory/keto/commit/d8e2a869db2a9cfb44423b434330536036b2f421))
- Rename /relationtuple endpoint to /relation-tuples
  ([#519](https://github.com/ory/keto/issues/519))
  ([8eb55f6](https://github.com/ory/keto/commit/8eb55f6269399f2bc5f000b8a768bcdf356c756f))
- Resolve gitignore build
  ([6f04bbb](https://github.com/ory/keto/commit/6f04bbb6057779b4d73d3f94677cea365843f7ac))
- Resolve goreleaser issues
  ([d32767f](https://github.com/ory/keto/commit/d32767f32856cf5bd48514c5d61746417fbed6f5))
- Resolve windows build issues
  ([8bcdfbf](https://github.com/ory/keto/commit/8bcdfbfbdb0b10c03ff93838e8fe6e778236e96d))
- Rewrite check engine to search starting at the object
  ([#310](https://github.com/ory/keto/issues/310))
  ([7d99694](https://github.com/ory/keto/commit/7d9969414ebc8cf6ef5d211ad34f8ae01bd3b4ee)),
  closes [#302](https://github.com/ory/keto/issues/302)
- Secure query building ([#442](https://github.com/ory/keto/issues/442))
  ([c7d2770](https://github.com/ory/keto/commit/c7d2770ed570238fd1262bcc4e5b4afa6c12d80e))
- Strict version enforcement in docker
  ([e45b28f](https://github.com/ory/keto/commit/e45b28fec626db35f1bd4580e5b11c9c94a02669))
- Update dd-trace to fix build issues
  ([2ad489f](https://github.com/ory/keto/commit/2ad489f0d9cae3191718d36823fe25df58ab95e6))
- Update docker to go 1.16 and alpine
  ([c63096c](https://github.com/ory/keto/commit/c63096cb53d2171f22f4a0d4a9ac3c9bfac89d01))
- Use errors.WithStack everywhere
  ([#462](https://github.com/ory/keto/issues/462))
  ([5f25bce](https://github.com/ory/keto/commit/5f25bceea35179c67d24dd95f698dc57b789d87a)),
  closes [#437](https://github.com/ory/keto/issues/437):

  Fixed all occurrences found using the search pattern `return .*, err\n`.

- Use package name in pkger
  ([6435939](https://github.com/ory/keto/commit/6435939ad7e5899505cd0e6261f5dfc819c9ca42))
- **schema:** Add trace level to logger
  ([a5a1402](https://github.com/ory/keto/commit/a5a1402c61e1a37b1a9a349ad5736eaca66bd6a4))
- Use make() to initialize slices
  ([#250](https://github.com/ory/keto/issues/250))
  ([84f028d](https://github.com/ory/keto/commit/84f028dc35665174542e103c0aefc635bb6d3e52)),
  closes [#217](https://github.com/ory/keto/issues/217)

### Build System

- Pin dependency versions of buf and protoc plugins
  ([#338](https://github.com/ory/keto/issues/338))
  ([5a2fd1c](https://github.com/ory/keto/commit/5a2fd1cc8dff02aa7017771adc0d9101f6c86775))

### Code Generation

- Pin v0.6.0-alpha.1 release commit
  ([875af25](https://github.com/ory/keto/commit/875af25f89b813455148e58884dcdf1cd3600b86))

### Code Refactoring

- Data structures ([#279](https://github.com/ory/keto/issues/279))
  ([1316077](https://github.com/ory/keto/commit/131607762d0006e4cf4f93e8731ef7648348b2ec))

### Documentation

- Add check- and expand-API guides
  ([#493](https://github.com/ory/keto/issues/493))
  ([09a25b4](https://github.com/ory/keto/commit/09a25b4063abcfdcd4c0de315a2ef088d6d4e72e))
- Add current features overview ([#505](https://github.com/ory/keto/issues/505))
  ([605afa0](https://github.com/ory/keto/commit/605afa029794ad115bba02e004e1596cea038e8e))
- Add missing pages ([#518](https://github.com/ory/keto/issues/518))
  ([43cbaa9](https://github.com/ory/keto/commit/43cbaa9140cfa0ea3c72f699f6bb34f5ed31d8dd))
- Add namespace and relation naming conventions
  ([#510](https://github.com/ory/keto/issues/510))
  ([dd31865](https://github.com/ory/keto/commit/dd318653178cd45da47f3e7cef507b42708363ef))
- Add performance page ([#413](https://github.com/ory/keto/issues/413))
  ([6fe0639](https://github.com/ory/keto/commit/6fe0639d36087b5ecd555eb6fe5ce949f3f6f0d7)):

  This also refactored the server startup. Functionality did not change.

- Add production guide
  ([a9163c7](https://github.com/ory/keto/commit/a9163c7690c55c8191650c4dfb464b75ea02446b))
- Add zanzibar overview to README.md
  ([#265](https://github.com/ory/keto/issues/265))
  ([15a95b2](https://github.com/ory/keto/commit/15a95b28e745592353e4656d42a9d0bd20ce468f))
- API overview ([#501](https://github.com/ory/keto/issues/501))
  ([05fe03b](https://github.com/ory/keto/commit/05fe03b5bf7a3f790aa6c9c1d3fcdb31304ef6af))
- Concepts ([#429](https://github.com/ory/keto/issues/429))
  ([2f2c885](https://github.com/ory/keto/commit/2f2c88527b3f6d1d46a5c287d8aca0874d18a28d))
- Delete old redirect homepage
  ([c0a3784](https://github.com/ory/keto/commit/c0a378448f8c7723bae68f7b52a019b697b25863))
- Document gRPC SKDs
  ([7583fe8](https://github.com/ory/keto/commit/7583fe8933f6676b4e37477098b1d43d12819b8b))
- Fix grammatical error ([#222](https://github.com/ory/keto/issues/222))
  ([256a0d2](https://github.com/ory/keto/commit/256a0d2e53fe1eb859e41fc539870ae1d5a493d2))
- Fix regression issues
  ([9697bb4](https://github.com/ory/keto/commit/9697bb43dd23c0d1fae74ea42e848883c45dae77))
- Generate gRPC reference page ([#488](https://github.com/ory/keto/issues/488))
  ([93ebe6d](https://github.com/ory/keto/commit/93ebe6db7e887d708503a54c5ec943254e37ca43))
- Improve CLI documentation ([#503](https://github.com/ory/keto/issues/503))
  ([be9327f](https://github.com/ory/keto/commit/be9327f7b28152a78f731043acf83b7092e42e29))
- Minor fixes ([#532](https://github.com/ory/keto/issues/532))
  ([638342e](https://github.com/ory/keto/commit/638342eb9519d9bf609926fb87558071e2815fb3))
- Move development section
  ([9ff393f](https://github.com/ory/keto/commit/9ff393f6cba1fb0a33918377ce505455c34d9dfc))
- Move to json sidebar
  ([257bf96](https://github.com/ory/keto/commit/257bf96044df37c3d7af8a289fb67098d48da1a3))
- Remove duplicate "is"
  ([ca3277d](https://github.com/ory/keto/commit/ca3277d82c1508797bc8c663963407d2e4d9112f))
- Remove duplicate template
  ([1d3b38e](https://github.com/ory/keto/commit/1d3b38e4045b0b874bb1186ea628f5a37353a2e6))
- Remove old documentation ([#426](https://github.com/ory/keto/issues/426))
  ([eb76913](https://github.com/ory/keto/commit/eb7691306018678e024211b51627a1c27e780a6b))
- Replace TODO links ([#512](https://github.com/ory/keto/issues/512))
  ([ad8e20b](https://github.com/ory/keto/commit/ad8e20b3bef2bc46b3a32c2c9ccb6e16e4bad22c))
- Resolve broken links
  ([0d0a50b](https://github.com/ory/keto/commit/0d0a50b3f4112893f32c81adc8edd137b5a62541))
- Simple access check guide ([#451](https://github.com/ory/keto/issues/451))
  ([e0485af](https://github.com/ory/keto/commit/e0485afc46a445868580aa541e962e80cbea0670)):

  This also enables gRPC go, gRPC nodejs, cURL, and Keto CLI code samples to be
  tested.

- Update comment in write response
  ([#329](https://github.com/ory/keto/issues/329))
  ([4ca0baf](https://github.com/ory/keto/commit/4ca0baf62e34402e749e870fe8c0cc893684192c))
- Update install instructions
  ([d2e4123](https://github.com/ory/keto/commit/d2e4123f3e2e58da8be181a0a542e3dcc1313e16))
- Update introduction
  ([5f71d73](https://github.com/ory/keto/commit/5f71d73e2ee95d02abc4cd42a76c98a35942df0c))
- Update README ([#515](https://github.com/ory/keto/issues/515))
  ([18d3cd6](https://github.com/ory/keto/commit/18d3cd61b0a79400170dc0f89860b4614cc4a543)):

  Also format all markdown files in the root.

- Update repository templates
  ([db505f9](https://github.com/ory/keto/commit/db505f9e10755bc20c4623c4f5f99f33283dffda))
- Update repository templates
  ([6c056bb](https://github.com/ory/keto/commit/6c056bb2043af6e82f06fdfa509ab3fa0d5e5d06))
- Update SDK links ([#514](https://github.com/ory/keto/issues/514))
  ([f920fbf](https://github.com/ory/keto/commit/f920fbfc8dcc6711ad9e046578a4506179952be7))
- Update swagger documentation for REST endpoints
  ([c363de6](https://github.com/ory/keto/commit/c363de61edf912fef85acc6bcdac6e1c15c48f4f))
- Use mdx for api reference
  ([340f3a3](https://github.com/ory/keto/commit/340f3a3dd20c82c743e7b3ad6aaf06a4c118b5a1))
- Various improvements and updates
  ([#486](https://github.com/ory/keto/issues/486))
  ([a812ace](https://github.com/ory/keto/commit/a812ace2303214e0e7acb2e283efa1cff0d5d279))

### Features

- Add .dockerignore
  ([8b0ff06](https://github.com/ory/keto/commit/8b0ff066b2508ef2f3629f9a3e2fce601b8dcce1))
- Add and automate version schema
  ([b01eef8](https://github.com/ory/keto/commit/b01eef8d4d5834b5888cb369ecf01ee01b40c24c))
- Add check engine ([#277](https://github.com/ory/keto/issues/277))
  ([396c1ae](https://github.com/ory/keto/commit/396c1ae33b777031f8d59549d9de4a88e3f6b10a))
- Add gRPC health status ([#427](https://github.com/ory/keto/issues/427))
  ([51c4223](https://github.com/ory/keto/commit/51c4223d6cb89a9bfbc115ef20db8350aeb2e8af))
- Add is_last_page to list response
  ([#425](https://github.com/ory/keto/issues/425))
  ([b73d91f](https://github.com/ory/keto/commit/b73d91f061ab155c53d802263c0263aa39e64bdf))
- Add POST REST handler for policy check
  ([7d89860](https://github.com/ory/keto/commit/7d89860bc4a790a69f5bea5b0dbe4a2938c6729f))
- Add relation write API ([#275](https://github.com/ory/keto/issues/275))
  ([f2ddb9d](https://github.com/ory/keto/commit/f2ddb9d884ed71037b5371c00bb10b63d25d47c0))
- Add REST and gRPC logger middlewares
  ([#436](https://github.com/ory/keto/issues/436))
  ([615eb0b](https://github.com/ory/keto/commit/615eb0bec3bdc0fd26abc7af0b8990269b0cbedd))
- Add SQA telemetry ([#535](https://github.com/ory/keto/issues/535))
  ([9f6472b](https://github.com/ory/keto/commit/9f6472b0c996505d41058e9b55afa8fd6b9bb2d5))
- Add sql persister ([#350](https://github.com/ory/keto/issues/350))
  ([d595d52](https://github.com/ory/keto/commit/d595d52dabb8f4953b5c23d3a8154cac13d00306))
- Add tracing ([#536](https://github.com/ory/keto/issues/536))
  ([b57a144](https://github.com/ory/keto/commit/b57a144e0a7ec39d5831dbb79840c2b25c044e6a))
- Allow to apply namespace migrations together with regular migrations
  ([#441](https://github.com/ory/keto/issues/441))
  ([57e2bbc](https://github.com/ory/keto/commit/57e2bbc5eaebe43834f2432eb1ee2820d9cb2988))
- Delete relation tuples ([#457](https://github.com/ory/keto/issues/457))
  ([3ec8afa](https://github.com/ory/keto/commit/3ec8afa68c5b5ddc26609b9afd17cc0d06cd82bf)),
  closes [#452](https://github.com/ory/keto/issues/452)
- Dockerfile and docker compose example
  ([#390](https://github.com/ory/keto/issues/390))
  ([10cd0b3](https://github.com/ory/keto/commit/10cd0b39c12ef96710bda6ff013f7c5eeae97118))
- Expand API ([#285](https://github.com/ory/keto/issues/285))
  ([a3ca0b8](https://github.com/ory/keto/commit/a3ca0b8a109b63f443e359cd8ff18a7b3e489f84))
- Expand GPRC service and CLI ([#383](https://github.com/ory/keto/issues/383))
  ([acf2154](https://github.com/ory/keto/commit/acf21546d3e135deb77c853b751a3da3a7b16f00))
- First API draft and generation
  ([#315](https://github.com/ory/keto/issues/315))
  ([bda5d8b](https://github.com/ory/keto/commit/bda5d8b7e90d749600f5b5e169df8a6ec3705b22))
- GRPC status codes and improved error messages
  ([#467](https://github.com/ory/keto/issues/467))
  ([4a4f8c6](https://github.com/ory/keto/commit/4a4f8c6b323664329414b61e7d80d7838face730))
- GRPC version API ([#475](https://github.com/ory/keto/issues/475))
  ([89cc46f](https://github.com/ory/keto/commit/89cc46fe4a13b062693d3db4f803834ba37f4e48))
- Implement goreleaser pipeline
  ([888ac43](https://github.com/ory/keto/commit/888ac43e6f706f619b2f1b58271dd027094c9ae9)),
  closes [#410](https://github.com/ory/keto/issues/410)
- Incorporate new GRPC API structure
  ([#331](https://github.com/ory/keto/issues/331))
  ([e0916ad](https://github.com/ory/keto/commit/e0916ad00c81b24177cfe45faf77b93d2c33dc1f))
- Koanf and namespace configuration
  ([#367](https://github.com/ory/keto/issues/367))
  ([3ad32bc](https://github.com/ory/keto/commit/3ad32bc13a4d96135be8031eb6fe4c15868272ca))
- Namespace configuration ([#324](https://github.com/ory/keto/issues/324))
  ([b94f50d](https://github.com/ory/keto/commit/b94f50d1800c47a43561df5009cb38b44ccd0088))
- Namespace migrate status CLI ([#508](https://github.com/ory/keto/issues/508))
  ([e3f7ad9](https://github.com/ory/keto/commit/e3f7ad91585b616e97f85ce0f55c76406b6c4d0a)):

  This also refactors the current `migrate` and `namespace migrate` commands.

- Nodejs gRPC definitions ([#447](https://github.com/ory/keto/issues/447))
  ([3b5c313](https://github.com/ory/keto/commit/3b5c31326645adb2d5b14ced901771a7ba00fd1c)):

  Includes Typescript definitions.

- Read API ([#269](https://github.com/ory/keto/issues/269))
  ([de5119a](https://github.com/ory/keto/commit/de5119a6e3c7563cfc2e1ada12d47b27ebd7faaa)):

  This is a first draft of the read API. It is reachable by REST and gRPC calls.
  The main purpose of this PR is to establish the basic repository structure and
  define the API.

- Relationtuple parse command ([#490](https://github.com/ory/keto/issues/490))
  ([91a3cf4](https://github.com/ory/keto/commit/91a3cf47fbdb8203b799cf7c69bcf3dbbfb98b3a)):

  This command parses the relation tuple format used in the docs. It greatly
  improves the experience when copying something from the documentation. It can
  especially be used to pipe relation tuples into other commands, e.g.:

  ```shell
  echo "messages:02y_15_4w350m3#decypher@john" | \
    keto relation-tuple parse - --format json | \
    keto relation-tuple create -
  ```

- REST patch relation tuples ([#491](https://github.com/ory/keto/issues/491))
  ([d38618a](https://github.com/ory/keto/commit/d38618a9e647902ce019396ff1c33973020bf797)):

  The new PATCH handler allows transactional changes similar to the already
  existing gRPC service.

- Separate and multiplex ports based on read/write privilege
  ([#397](https://github.com/ory/keto/issues/397))
  ([6918ac3](https://github.com/ory/keto/commit/6918ac3bfa355cbd551e44376c214f412e3414e4))
- Swagger SDK ([#476](https://github.com/ory/keto/issues/476))
  ([011888c](https://github.com/ory/keto/commit/011888c2b7e2d0f7b8923c994c70e62d374a2830))

### Tests

- Add command tests ([#487](https://github.com/ory/keto/issues/487))
  ([61c28e4](https://github.com/ory/keto/commit/61c28e48a5c3f623e5cc133e69ba368c5103f414))
- Add dedicated persistence tests
  ([#416](https://github.com/ory/keto/issues/416))
  ([4e98906](https://github.com/ory/keto/commit/4e9890605edf3ea26134917a95bfa6fbb176565e))
- Add handler tests ([#478](https://github.com/ory/keto/issues/478))
  ([9315a77](https://github.com/ory/keto/commit/9315a77820d50e400b78f2f019a871be022a9887))
- Add initial e2e test ([#380](https://github.com/ory/keto/issues/380))
  ([dc5d3c9](https://github.com/ory/keto/commit/dc5d3c9d02604fddbfa56ac5ebbc1fef56a881d9))
- Add relationtuple definition tests
  ([#415](https://github.com/ory/keto/issues/415))
  ([2e3dcb2](https://github.com/ory/keto/commit/2e3dcb200a7769dc8710d311ca08a7515012fbdd))
- Enable GRPC client in e2e test
  ([#382](https://github.com/ory/keto/issues/382))
  ([4e5c6ae](https://github.com/ory/keto/commit/4e5c6aed56e5a449003956ec114ec131be068aaf))
- Improve docs sample tests ([#461](https://github.com/ory/keto/issues/461))
  ([6e0e5e6](https://github.com/ory/keto/commit/6e0e5e6184916e894fd4694cfa3a158f11fae11f))

# [0.5.6-alpha.1](https://github.com/ory/keto/compare/v0.5.5-alpha.1...v0.5.6-alpha.1) (2020-05-28)

This release bumps vulnerable transient dependencies (those are not actually
used in ORY Keto) and updates several documentation pages and improves
structured logging output. Additionally, ORY Keto now uses the updated release
pipeline!

### Bug Fixes

- Update install script
  ([21e1bf0](https://github.com/ory/keto/commit/21e1bf05177576a9d743bd11744ef6a42be50b8d))

### Chores

- Pin v0.5.6-alpha.1 release commit
  ([ed0da08](https://github.com/ory/keto/commit/ed0da08a03a910660358fc56c568692325749b6d))

# [0.5.5-alpha.1](https://github.com/ory/keto/compare/v0.5.4-alpha.1...v0.5.5-alpha.1) (2020-05-28)

This release bumps vulnerable transient dependencies (those are not actually
used in ORY Keto) and updates several documentation pages and improves
structured logging output. Additionally, ORY Keto now uses the updated release
pipeline!

### Bug Fixes

- Move deps to go_mod_indirect_pins
  ([dd3e971](https://github.com/ory/keto/commit/dd3e971ac418baf10c1b33005acc7e6f66fb0d85))
- Resolve test issues
  ([9bd9956](https://github.com/ory/keto/commit/9bd9956e33731f1619c32e1e6b7c78f42e7c47c3))
- Update install.sh script
  ([f64d320](https://github.com/ory/keto/commit/f64d320b6424fe3256eb7fad1c94dcc1ef0bf487))
- Use semver-regex replacer func
  ([2cc3bbb](https://github.com/ory/keto/commit/2cc3bbb2d75ba5fa7a3653d7adcaa712ff38c603))

### Chores

- Pin v0.5.5-alpha.1 release commit
  ([4666a0f](https://github.com/ory/keto/commit/4666a0f258f253d19a14eca34f4b7049f2d0afa2))

### Documentation

- Add missing colon in docker run command
  ([#193](https://github.com/ory/keto/issues/193))
  ([383063d](https://github.com/ory/keto/commit/383063d260d995665da4c02c9a7bac7e06a2c8d3))
- Update github templates ([#182](https://github.com/ory/keto/issues/182))
  ([72ea09b](https://github.com/ory/keto/commit/72ea09bbbf9925d7705842703b32826376f636e4))
- Update github templates ([#184](https://github.com/ory/keto/issues/184))
  ([ed546b7](https://github.com/ory/keto/commit/ed546b7a2b9ee690284a48c641edd1570464d71f))
- Update github templates ([#188](https://github.com/ory/keto/issues/188))
  ([ebd75b2](https://github.com/ory/keto/commit/ebd75b2f6545ff4372773f6370300c7b2ca71c51))
- Update github templates ([#189](https://github.com/ory/keto/issues/189))
  ([fd4c0b1](https://github.com/ory/keto/commit/fd4c0b17bcb1c281baac1772ab94e305ec8c5c86))
- Update github templates ([#195](https://github.com/ory/keto/issues/195))
  ([ba0943c](https://github.com/ory/keto/commit/ba0943c45d36ef10bdf1169f0aeef439a3a67d28))
- Update linux install guide ([#191](https://github.com/ory/keto/issues/191))
  ([7d8b24b](https://github.com/ory/keto/commit/7d8b24bddb9c92feb78c7b65f39434d538773b58))
- Update repository templates
  ([ea65b5c](https://github.com/ory/keto/commit/ea65b5c5ada0a7453326fa755aa914306f1b1851))
- Use central banner repo for README
  ([0d95d97](https://github.com/ory/keto/commit/0d95d97504df4d0ab57d18dc6d0a824a3f8f5896))
- Use correct banner
  ([c6dfe28](https://github.com/ory/keto/commit/c6dfe280fd962169c424834cea040a408c1bc83f))
- Use correct version
  ([5f7030c](https://github.com/ory/keto/commit/5f7030c9069fe392200be72f8ce1a93890fbbba8)),
  closes [#200](https://github.com/ory/keto/issues/200)
- Use correct versions in install docs
  ([52e6c34](https://github.com/ory/keto/commit/52e6c34780ed41c169504d71c39459898b5d14f9))

# [0.5.4-alpha.1](https://github.com/ory/keto/compare/v0.5.3-alpha.3...v0.5.4-alpha.1) (2020-04-07)

fix: resolve panic when executing migrations (#178)

Closes #177

### Bug Fixes

- Resolve panic when executing migrations
  ([#178](https://github.com/ory/keto/issues/178))
  ([7e83fee](https://github.com/ory/keto/commit/7e83feefaad041c60f09232ac44ed8b7240c6558)),
  closes [#177](https://github.com/ory/keto/issues/177)

# [0.5.3-alpha.3](https://github.com/ory/keto/compare/v0.5.3-alpha.2...v0.5.3-alpha.3) (2020-04-06)

autogen(docs): regenerate and update changelog

### Code Generation

- **docs:** Regenerate and update changelog
  ([769cef9](https://github.com/ory/keto/commit/769cef90f27ba9c203d3faf47272287ab17dc7eb))

### Code Refactoring

- Move docs to this repository ([#172](https://github.com/ory/keto/issues/172))
  ([312480d](https://github.com/ory/keto/commit/312480de3cefc5b72916ba95d8287443cf3ccb3d))

### Documentation

- Regenerate and update changelog
  ([dda79b1](https://github.com/ory/keto/commit/dda79b106a18bc33d70ae60e352118b0d288d26b))
- Regenerate and update changelog
  ([9048dd8](https://github.com/ory/keto/commit/9048dd8d8a0f0654072b3d4b77261fe947a34ece))
- Regenerate and update changelog
  ([806f68c](https://github.com/ory/keto/commit/806f68c603781742e0177ec0b2deecaf64c5b721))
- Regenerate and update changelog
  ([8905ee7](https://github.com/ory/keto/commit/8905ee74d4ec394af92240e180cc5d7f6493cb2f))
- Regenerate and update changelog
  ([203c1cc](https://github.com/ory/keto/commit/203c1cc659a72f81a370d7b9b7fbda60e7c96c9e))
- Regenerate and update changelog
  ([8875a95](https://github.com/ory/keto/commit/8875a95b35df57668acb27820a3aff1cdfbe8b30))
- Regenerate and update changelog
  ([28ddd3e](https://github.com/ory/keto/commit/28ddd3e1483afe8571b3d2bf9afcc31386d85f7f))
- Regenerate and update changelog
  ([927c4ed](https://github.com/ory/keto/commit/927c4edc4a770133bcb34bc044dd5c5e0eb3ffb7))
- Updates issue and pull request templates
  ([#168](https://github.com/ory/keto/issues/168))
  ([29a38a8](https://github.com/ory/keto/commit/29a38a85c61ec2c8d0ad2ce6d5a0f9e9d74b52f7))
- Updates issue and pull request templates
  ([#169](https://github.com/ory/keto/issues/169))
  ([99b7d5d](https://github.com/ory/keto/commit/99b7d5de24fed1aed746c4447a390d084632f89a))
- Updates issue and pull request templates
  ([#171](https://github.com/ory/keto/issues/171))
  ([7a9876b](https://github.com/ory/keto/commit/7a9876b8ed4282f50f886a025033641bd027a0e2))

# [0.5.3-alpha.1](https://github.com/ory/keto/compare/v0.5.2...v0.5.3-alpha.1) (2020-04-03)

chore: move to ory analytics fork (#167)

### Chores

- Move to ory analytics fork ([#167](https://github.com/ory/keto/issues/167))
  ([f824011](https://github.com/ory/keto/commit/f824011b4d19058504b3a43ed53a420619444a51))

# [0.5.2](https://github.com/ory/keto/compare/v0.5.1-alpha.1...v0.5.2) (2020-04-02)

docs: Regenerate and update changelog

### Documentation

- Regenerate and update changelog
  ([1e52100](https://github.com/ory/keto/commit/1e521001a43a0a13e2224e1a44956442ac6ffbc7))
- Regenerate and update changelog
  ([e4d32a6](https://github.com/ory/keto/commit/e4d32a62c1ae96115ea50bb471f5ff2ce2f2c4b9))

# [0.5.0](https://github.com/ory/keto/compare/v0.4.5-alpha.1...v0.5.0) (2020-04-02)

docs: use real json bool type in swagger (#162)

Closes #160

### Bug Fixes

- Move to ory sqa service ([#159](https://github.com/ory/keto/issues/159))
  ([c3bf1b1](https://github.com/ory/keto/commit/c3bf1b1964a14be4cc296aae98d0739e65917e18))
- Use correct response mode for removeOryAccessControlPolicyRoleMe…
  ([#161](https://github.com/ory/keto/issues/161))
  ([17543cf](https://github.com/ory/keto/commit/17543cfef63a1d040a2234bd63b210fb9c4f6015))

### Documentation

- Regenerate and update changelog
  ([6a77f75](https://github.com/ory/keto/commit/6a77f75d66e89420f2daf2fae011d31bcfa34008))
- Regenerate and update changelog
  ([c8c9d29](https://github.com/ory/keto/commit/c8c9d29e77ef53e1196cc6fe600c53d93376229b))
- Regenerate and update changelog
  ([fe8327d](https://github.com/ory/keto/commit/fe8327d951394084df7785166c9a9578c1ab0643))
- Regenerate and update changelog
  ([b5b1d66](https://github.com/ory/keto/commit/b5b1d66a4b933df8789337cce3f6d6bf391b617b))
- Update forum and chat links
  ([e96d7ba](https://github.com/ory/keto/commit/e96d7ba3dcc693c22eb983b3f58a05c9c6adbda7))
- Updates issue and pull request templates
  ([#158](https://github.com/ory/keto/issues/158))
  ([ab14cfa](https://github.com/ory/keto/commit/ab14cfa51ce195b26a83c050452530a5008589d7))
- Use real json bool type in swagger
  ([#162](https://github.com/ory/keto/issues/162))
  ([5349e7f](https://github.com/ory/keto/commit/5349e7f910ad22558a01b76be62db2136b5eb301)),
  closes [#160](https://github.com/ory/keto/issues/160)

# [0.4.5-alpha.1](https://github.com/ory/keto/compare/v0.4.4-alpha.1...v0.4.5-alpha.1) (2020-02-29)

docs: Regenerate and update changelog

### Bug Fixes

- **driver:** Extract scheme from DSN using sqlcon.GetDriverName
  ([#156](https://github.com/ory/keto/issues/156))
  ([187e289](https://github.com/ory/keto/commit/187e289f1a235b5cacf2a0b7ca5e98c384fa7a14)),
  closes [#145](https://github.com/ory/keto/issues/145)

### Documentation

- Regenerate and update changelog
  ([41513da](https://github.com/ory/keto/commit/41513da35ea038f3c4cc2d98b9796cee5b5a8b92))

# [0.4.4-alpha.1](https://github.com/ory/keto/compare/v0.4.3-alpha.2...v0.4.4-alpha.1) (2020-02-14)

docs: Regenerate and update changelog

### Bug Fixes

- **goreleaser:** Update brew section
  ([0918ff3](https://github.com/ory/keto/commit/0918ff3032eeecd26c67d6249c7e28e71ee110af))

### Documentation

- Prepare ecosystem automation
  ([2e39be7](https://github.com/ory/keto/commit/2e39be79ebad1cec021ae3ee4b0a75ffea4b7424))
- Regenerate and update changelog
  ([009c4c4](https://github.com/ory/keto/commit/009c4c4e4fd4c5607cc30cc9622fd0f82e3891f3))
- Regenerate and update changelog
  ([49f3c4b](https://github.com/ory/keto/commit/49f3c4ba34df5879d8f48cc96bf0df9dad820362))
- Updates issue and pull request templates
  ([#153](https://github.com/ory/keto/issues/153))
  ([7fb7521](https://github.com/ory/keto/commit/7fb752114e1e2a91ab96fdb546835de8aee4926b))

### Features

- **ci:** Add nancy vuln scanner
  ([#152](https://github.com/ory/keto/issues/152))
  ([c19c2b9](https://github.com/ory/keto/commit/c19c2b9efe8299b8878cc8099fe314d8dcda3a08))

### Unclassified

- Update CHANGELOG [ci skip]
  ([63fe513](https://github.com/ory/keto/commit/63fe513d22ec3747a95cdb8f797ba1eba5ca344f))
- Update CHANGELOG [ci skip]
  ([7b7c3ac](https://github.com/ory/keto/commit/7b7c3ac6c06c072fea1b64624ea79a3fd406b09c))
- Update CHANGELOG [ci skip]
  ([8886392](https://github.com/ory/keto/commit/8886392b39fb46ad338c8284866d4dae64ad1826))
- Update CHANGELOG [ci skip]
  ([5bbc284](https://github.com/ory/keto/commit/5bbc2844c49b0a68ba3bd8b003d91f87e2aed9e2))

# [0.4.3-alpha.2](https://github.com/ory/keto/compare/v0.4.3-alpha.1...v0.4.3-alpha.2) (2020-01-31)

Update README.md

### Unclassified

- Update README.md
  ([0ab9c6f](https://github.com/ory/keto/commit/0ab9c6f372a1538a958a68b34315c9167b5a9093))
- Update CHANGELOG [ci skip]
  ([f0a1428](https://github.com/ory/keto/commit/f0a1428f4b99ceb35ff4f1e839bc5237e19db628))

# [0.4.3-alpha.1](https://github.com/ory/keto/compare/v0.4.2-alpha.1...v0.4.3-alpha.1) (2020-01-23)

Disable access logging for health endpoints (#151)

Closes #150

### Unclassified

- Disable access logging for health endpoints (#151)
  ([6ca0c09](https://github.com/ory/keto/commit/6ca0c09b5618122762475cffdc9c32adf28456a1)),
  closes [#151](https://github.com/ory/keto/issues/151)
  [#150](https://github.com/ory/keto/issues/150)

# [0.4.2-alpha.1](https://github.com/ory/keto/compare/v0.4.1-beta.1...v0.4.2-alpha.1) (2020-01-14)

Update CHANGELOG [ci skip]

### Unclassified

- Update CHANGELOG [ci skip]
  ([afaabde](https://github.com/ory/keto/commit/afaabde63affcf568e3090e55b4b957edff2890c))

# [0.4.1-beta.1](https://github.com/ory/keto/compare/v0.4.0-sandbox...v0.4.1-beta.1) (2020-01-13)

Update CHANGELOG [ci skip]

### Unclassified

- Update CHANGELOG [ci skip]
  ([e3ca5a7](https://github.com/ory/keto/commit/e3ca5a7d8b9827ffc7b31a8b5e459db3e912a590))
- Update SDK
  ([5dd6237](https://github.com/ory/keto/commit/5dd623755d4832f33c3dcefb778a9a70eace7b52))

# [0.4.0-alpha.1](https://github.com/ory/keto/compare/v0.3.9-sandbox...v0.4.0-alpha.1) (2020-01-13)

Move to new SDK generators (#146)

### Unclassified

- Move to new SDK generators (#146)
  ([4f51a09](https://github.com/ory/keto/commit/4f51a0948723efc092f1887b111d1e6dd590a075)),
  closes [#146](https://github.com/ory/keto/issues/146)
- Fix typos in the README (#144)
  ([85d838c](https://github.com/ory/keto/commit/85d838c0872c73eb70b5bfff1ccb175b07f6b1e4)),
  closes [#144](https://github.com/ory/keto/issues/144)

# [0.3.9-sandbox](https://github.com/ory/keto/compare/v0.3.8-sandbox...v0.3.9-sandbox) (2019-12-16)

Update go modules

### Unclassified

- Update go modules
  ([1151e07](https://github.com/ory/keto/commit/1151e0755c974b0aea86be5aaeae365ea9aef094))

# [0.3.7-sandbox](https://github.com/ory/keto/compare/v0.3.6-sandbox...v0.3.7-sandbox) (2019-12-11)

Update documentation banner image (#143)

### Unclassified

- Update documentation banner image (#143)
  ([e444755](https://github.com/ory/keto/commit/e4447552031a4f26ec21a336071b0bb19843df61)),
  closes [#143](https://github.com/ory/keto/issues/143)
- Revert incorrect license changes
  ([094c4f3](https://github.com/ory/keto/commit/094c4f30184d77a05044087c13e71ce4adb4d735))
- Fix invalid pseudo version ([#138](https://github.com/ory/keto/issues/138))
  ([79b4457](https://github.com/ory/keto/commit/79b4457f0162197ba267edbb8c0031c47e03bade))

# [0.3.6-sandbox](https://github.com/ory/keto/compare/v0.3.5-sandbox...v0.3.6-sandbox) (2019-10-16)

Resolve issues with mysql tests (#137)

### Unclassified

- Resolve issues with mysql tests (#137)
  ([ef5aec8](https://github.com/ory/keto/commit/ef5aec8e493199c46b78e8f1257aa41df9545f28)),
  closes [#137](https://github.com/ory/keto/issues/137)

# [0.3.5-sandbox](https://github.com/ory/keto/compare/v0.3.4-sandbox...v0.3.5-sandbox) (2019-08-21)

Implement roles and policies filter (#124)

### Documentation

- Incorporates changes from version v0.3.3-sandbox
  ([57686d2](https://github.com/ory/keto/commit/57686d2e30b229cae33e717eb8b3db9da3bdaf0a))
- README grammar fixes ([#114](https://github.com/ory/keto/issues/114))
  ([e592736](https://github.com/ory/keto/commit/e5927360300d8c4fbea841c1c2fb92b48b77885e))
- Updates issue and pull request templates
  ([#110](https://github.com/ory/keto/issues/110))
  ([80c8516](https://github.com/ory/keto/commit/80c8516efbcf33902d8a45f1dc7dbafff2aab8b1))
- Updates issue and pull request templates
  ([#111](https://github.com/ory/keto/issues/111))
  ([22305d0](https://github.com/ory/keto/commit/22305d0a9b5114de8125c16030bbcd1de695ae9b))
- Updates issue and pull request templates
  ([#112](https://github.com/ory/keto/issues/112))
  ([dccada9](https://github.com/ory/keto/commit/dccada9a2189bbd899c5c4a18665a92113fe6cd7))
- Updates issue and pull request templates
  ([#125](https://github.com/ory/keto/issues/125))
  ([15f373a](https://github.com/ory/keto/commit/15f373a16b8cfbd6cdad2bda5f161e171c566137))
- Updates issue and pull request templates
  ([#128](https://github.com/ory/keto/issues/128))
  ([eaf8e33](https://github.com/ory/keto/commit/eaf8e33f3904484635924bdac190c8dc7b60f939))
- Updates issue and pull request templates
  ([#130](https://github.com/ory/keto/issues/130))
  ([a440d14](https://github.com/ory/keto/commit/a440d142275a7a91a0a6bb487fe47d22247f4988))
- Updates issue and pull request templates
  ([#131](https://github.com/ory/keto/issues/131))
  ([dbf2cb2](https://github.com/ory/keto/commit/dbf2cb23c5b6f0f1ee0be5c0b5a58fb0c3dbefd1))
- Updates issue and pull request templates
  ([#132](https://github.com/ory/keto/issues/132))
  ([e121048](https://github.com/ory/keto/commit/e121048d10627ed32a07e26455efd69248f1bd95))
- Updates issue and pull request templates
  ([#133](https://github.com/ory/keto/issues/133))
  ([1b7490a](https://github.com/ory/keto/commit/1b7490abc1d5d0501b66595eb2d92834b6fb0345))

### Unclassified

- Implement roles and policies filter (#124)
  ([db94481](https://github.com/ory/keto/commit/db9448103621a6a8cd086a4cef6c6a22398e621f)),
  closes [#124](https://github.com/ory/keto/issues/124)
- Add adopters placeholder ([#129](https://github.com/ory/keto/issues/129))
  ([b814838](https://github.com/ory/keto/commit/b8148388b8bea97d1f1b4b54de2f0b8ef6b8b6c7))
- Improve documentation (#126)
  ([aabb04d](https://github.com/ory/keto/commit/aabb04d5f283d3c73eb3f3531b4e470ae716db5e)),
  closes [#126](https://github.com/ory/keto/issues/126)
- Create FUNDING.yml
  ([571b447](https://github.com/ory/keto/commit/571b447ed3a02f43623ef5c5adc09682b5f379bd))
- Use non-root user in image ([#116](https://github.com/ory/keto/issues/116))
  ([a493e55](https://github.com/ory/keto/commit/a493e550a8bb86d99164f4ea76dbcecf76c9c2c1))
- Remove binary license (#117)
  ([6e85f7c](https://github.com/ory/keto/commit/6e85f7c6f430e88fb4117a131f57bd69466a8ca1)),
  closes [#117](https://github.com/ory/keto/issues/117)

# [0.3.3-sandbox](https://github.com/ory/keto/compare/v0.3.1-sandbox...v0.3.3-sandbox) (2019-05-18)

ci: Resolve goreleaser issues (#108)

### Continuous Integration

- Resolve goreleaser issues ([#108](https://github.com/ory/keto/issues/108))
  ([5753f27](https://github.com/ory/keto/commit/5753f27a9e89ccdda7c02969217c253aa72cb94b))

### Documentation

- Incorporates changes from version v0.3.1-sandbox
  ([b8a0029](https://github.com/ory/keto/commit/b8a002937483a0f71fe5aba26bb18beb41886249))
- Updates issue and pull request templates
  ([#106](https://github.com/ory/keto/issues/106))
  ([54a5a27](https://github.com/ory/keto/commit/54a5a27f24a90ab3c5f9915f36582b85eecd0d62))

# [0.3.1-sandbox](https://github.com/ory/keto/compare/v0.3.0-sandbox...v0.3.1-sandbox) (2019-04-29)

ci: Use image that includes bash/sh for release docs (#103)

Signed-off-by: aeneasr <aeneas@ory.sh>

### Continuous Integration

- Use image that includes bash/sh for release docs
  ([#103](https://github.com/ory/keto/issues/103))
  ([e9d3027](https://github.com/ory/keto/commit/e9d3027fc62b20f28cd7a023222390e24d565eb1))

### Documentation

- Incorporates changes from version v0.3.0-sandbox
  ([605d2f4](https://github.com/ory/keto/commit/605d2f43621b806b750edc81d439edc92cfb7c38))

### Unclassified

- Allow configuration files and update UPGRADE guide. (#102)
  ([3934dc6](https://github.com/ory/keto/commit/3934dc6e690822358067b43920048d45a4b7799b)),
  closes [#102](https://github.com/ory/keto/issues/102)

# [0.3.0-sandbox](https://github.com/ory/keto/compare/v0.2.3-sandbox+oryOS.10...v0.3.0-sandbox) (2019-04-29)

docker: Remove full tag from build pipeline (#101)

Signed-off-by: aeneasr <aeneas@ory.sh>

### Documentation

- Update patrons
  ([c8dc7cd](https://github.com/ory/keto/commit/c8dc7cdc68676970328b55648b8d6e469c77fbfd))

### Unclassified

- Improve naming for ory policies
  ([#100](https://github.com/ory/keto/issues/100))
  ([b39703d](https://github.com/ory/keto/commit/b39703d362d333213fcb7d3782e363d09b6dabbd))
- Remove full tag from build pipeline
  ([#101](https://github.com/ory/keto/issues/101))
  ([602a273](https://github.com/ory/keto/commit/602a273dc5a0c29e80a22f04adb937ab385c4512))
- Remove duplicate code in Makefile (#99)
  ([04f5223](https://github.com/ory/keto/commit/04f52231509dd0f3a57d745918fc43fff7c595ff)),
  closes [#99](https://github.com/ory/keto/issues/99)
- Add tracing support and general improvements (#98)
  ([63b3946](https://github.com/ory/keto/commit/63b3946e0ae1fa23c6a359e9a64b296addff868c)),
  closes [#98](https://github.com/ory/keto/issues/98):

  This patch improves the internal configuration and service management. It adds
  support for distributed tracing and resolves several issues in the release
  pipeline and CLI.

  Additionally, composable docker-compose configuration files have been added.

  Several bugs have been fixed in the release management pipeline.

- Add content-type in the response of allowed
  ([#90](https://github.com/ory/keto/issues/90))
  ([39a1486](https://github.com/ory/keto/commit/39a1486dc53456189d30380460a9aeba198fa9e9))
- Fix disable-telemetry check ([#85](https://github.com/ory/keto/issues/85))
  ([38b5383](https://github.com/ory/keto/commit/38b538379973fa34bd2bf24dcb2e6dbedf324e1e))
- Fix remove member from role ([#87](https://github.com/ory/keto/issues/87))
  ([698e161](https://github.com/ory/keto/commit/698e161989331ca5a3a0769301d9694ef805a876)),
  closes [#74](https://github.com/ory/keto/issues/74)
- Fix the type of conditions in the policy
  ([#86](https://github.com/ory/keto/issues/86))
  ([fc1ced6](https://github.com/ory/keto/commit/fc1ced63bd39c9fbf437e419dfc384343e36e0ee))
- Move Go SDK generation to go-swagger
  ([#94](https://github.com/ory/keto/issues/94))
  ([9f48a95](https://github.com/ory/keto/commit/9f48a95187a7b6160108cd7d0301590de2e58f07)),
  closes [#92](https://github.com/ory/keto/issues/92)
- Send 403 when authorization result is negative
  ([#93](https://github.com/ory/keto/issues/93))
  ([de806d8](https://github.com/ory/keto/commit/de806d892819db63c1abc259ab06ee08d87895dc)),
  closes [#75](https://github.com/ory/keto/issues/75)
- Update dependencies ([#91](https://github.com/ory/keto/issues/91))
  ([4d44174](https://github.com/ory/keto/commit/4d4417474ebf8cc69d01e5ac82633b966cdefbc7))
- storage/memory: Fix upsert with pre-existing key will causes duplicate records
  (#88)
  ([1cb8a36](https://github.com/ory/keto/commit/1cb8a36a08883b785d9bb0a4be1ddc00f1f9d358)),
  closes [#88](https://github.com/ory/keto/issues/88)
  [#80](https://github.com/ory/keto/issues/80)

# [0.2.3-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.2.2-sandbox+oryOS.10...v0.2.3-sandbox+oryOS.10) (2019-02-05)

dist: Fix packr build pipeline (#84)

Closes #73 Closes #81

Signed-off-by: aeneasr <aeneas@ory.sh>

### Documentation

- Add documentation for glob matching
  ([5c8babb](https://github.com/ory/keto/commit/5c8babbfbae01a78f30cfbff92d8e9c3a6b09027))
- Incorporates changes from version v0.2.2-sandbox+oryOS.10
  ([ed7af3f](https://github.com/ory/keto/commit/ed7af3fa4e5d1d0d03b5366f4cf865a5b82ec293))
- Properly generate api.swagger.json
  ([18e3f84](https://github.com/ory/keto/commit/18e3f84cdeee317f942d61753399675c98886e5d))

### Unclassified

- Add placeholder go file for rego inclusion
  ([6a6f64d](https://github.com/ory/keto/commit/6a6f64d8c59b496f6cf360f55eba1e16bf5380f1))
- Add support for glob matching
  ([bb76c6b](https://github.com/ory/keto/commit/bb76c6bebe522fc25448c4f4e4d1ef7c530a725f))
- Ex- and import rego subdirectories for `go get`
  [#77](https://github.com/ory/keto/issues/77)
  ([59cc053](https://github.com/ory/keto/commit/59cc05328f068fc3046b2dbc022a562fd5d67960)),
  closes [#73](https://github.com/ory/keto/issues/73)
- Fix packr build pipeline ([#84](https://github.com/ory/keto/issues/84))
  ([65a87d5](https://github.com/ory/keto/commit/65a87d564d237bc979bb5962beff7d3703d9689f)),
  closes [#73](https://github.com/ory/keto/issues/73)
  [#81](https://github.com/ory/keto/issues/81)
- Import glob in rego/doc.go
  ([7798442](https://github.com/ory/keto/commit/7798442553cfe7989a23d2c389c8c63a24013543))
- Properly handle dbal error
  ([6811607](https://github.com/ory/keto/commit/6811607ea79c8f3155a17bc1aea566e9e4680616))
- Properly handle TLS certificates if set
  ([36399f0](https://github.com/ory/keto/commit/36399f09261d4f3cb5e053679eee3cb15da2df19)),
  closes [#73](https://github.com/ory/keto/issues/73)

# [0.2.2-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.2.1-sandbox+oryOS.10...v0.2.2-sandbox+oryOS.10) (2018-12-13)

ci: Fix docker push arguments in publish task

Signed-off-by: aeneasr <aeneas@ory.sh>

### Continuous Integration

- Fix docker push arguments in publish task
  ([f03c77c](https://github.com/ory/keto/commit/f03c77c6b7461ab81cb03265cbec909ac45c2259))

# [0.2.1-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.2.0-sandbox+oryOS.10...v0.2.1-sandbox+oryOS.10) (2018-12-13)

ci: Fix docker release task

Signed-off-by: aeneasr <aeneas@ory.sh>

### Continuous Integration

- Fix docker release task
  ([7a0414f](https://github.com/ory/keto/commit/7a0414f614b6cc8b1d78cfbb773a2f0192d00d23))

# [0.2.0-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.0.1...v0.2.0-sandbox+oryOS.10) (2018-12-13)

all: gofmt

Signed-off-by: aeneasr <aeneas@ory.sh>

### Documentation

- Adds banner
  ([0ec1d8f](https://github.com/ory/keto/commit/0ec1d8f5e843465d17ac4c8f91f18e5badf16900))
- Adds GitHub templates & code of conduct
  ([#31](https://github.com/ory/keto/issues/31))
  ([a11e898](https://github.com/ory/keto/commit/a11e8980f2af528f1357659966123d0cbf7d43db))
- Adds link to examples repository
  ([#32](https://github.com/ory/keto/issues/32))
  ([7061a2a](https://github.com/ory/keto/commit/7061a2aa31652a9e0c2d449facb1201bfa11fd3f))
- Adds security console image
  ([fd27fc9](https://github.com/ory/keto/commit/fd27fc9cce50beb3d0189e0a93300879fd7149db))
- Changes hydra to keto in readme
  ([9dab531](https://github.com/ory/keto/commit/9dab531744cf5b0ae98862945d44b07535595781))
- Deprecate old versions in logs
  ([955d647](https://github.com/ory/keto/commit/955d647307a48ee7cf2d3f9fb4263072adf42299))
- Incorporates changes from version
  ([85c4d81](https://github.com/ory/keto/commit/85c4d81a192e92f874c106b91cfa6fb404d9a34a))
- Incorporates changes from version v0.0.0-testrelease.1
  ([6062dd4](https://github.com/ory/keto/commit/6062dd4a894607f5f1ead119af20cc8bdbe15bef))
- Incorporates changes from version v0.0.1-1-g85c4d81
  ([f4606fc](https://github.com/ory/keto/commit/f4606fce0326bece2a89dadc029bc5ce9778df18))
- Incorporates changes from version v0.0.1-11-g114914f
  ([92a4dca](https://github.com/ory/keto/commit/92a4dca7a41dcf3a88c4800bf6d2217f33cfcdd1))
- Incorporates changes from version v0.0.1-16-g7d8a8ad
  ([2b76a83](https://github.com/ory/keto/commit/2b76a83755153b3f8a2b8d28c5b0029d96d567b6))
- Incorporates changes from version v0.0.1-18-g099e7e0
  ([70b12ad](https://github.com/ory/keto/commit/70b12adf5bcc0e890d6707e11e891e6cedfb3d87))
- Incorporates changes from version v0.0.1-20-g97ccbe6
  ([b21d56e](https://github.com/ory/keto/commit/b21d56e599c7eb4c1769bc18878f7d5818b73023))
- Incorporates changes from version v0.0.1-30-gaf2c3b5
  ([a1d0dcc](https://github.com/ory/keto/commit/a1d0dcc78a9506260f86df00e4dff8ab02909ce1))
- Incorporates changes from version v0.0.1-32-gedb5a60
  ([a5c369a](https://github.com/ory/keto/commit/a5c369a90da67c96bbde60e673c67f50b841fadd))
- Incorporates changes from version v0.0.1-6-g570783e
  ([0fcbbcb](https://github.com/ory/keto/commit/0fcbbcb02f1d748f9c733c86368b223b2ee4c6e2))
- Incorporates changes from version v0.0.1-7-g0fcbbcb
  ([c0141a8](https://github.com/ory/keto/commit/c0141a8ec22ea1260bf2d45d72dfe06737ec0115))
- Incorporates changes from version v0.1.0-sandbox
  ([9ee0664](https://github.com/ory/keto/commit/9ee06646d2cfb2d69abdcc411e31d14957437a1e))
- Incorporates changes from version v1.0.0-beta.1-1-g162d7b8
  ([647c5a9](https://github.com/ory/keto/commit/647c5a9e1bc8d9d635bf6f2511c3faa9a9daefef))
- Incorporates changes from version v1.0.0-beta.2-11-g2b280bb
  ([936889d](https://github.com/ory/keto/commit/936889d760f04a03d498f65331d653cbad3702d0))
- Incorporates changes from version v1.0.0-beta.2-13-g382e1d3
  ([883df44](https://github.com/ory/keto/commit/883df44a922f3daee86597af467072555cadc7e7))
- Incorporates changes from version v1.0.0-beta.2-15-g74450da
  ([48dd9f1](https://github.com/ory/keto/commit/48dd9f1ffbeaa99ac8dc27085c5a50f9244bf9c3))
- Incorporates changes from version v1.0.0-beta.2-3-gf623c52
  ([b6b90e5](https://github.com/ory/keto/commit/b6b90e5b2180921f78064a60666704b4e72679b6))
- Incorporates changes from version v1.0.0-beta.2-5-g3852be5
  ([3f09090](https://github.com/ory/keto/commit/3f09090a2f82f3f29154c19217cea0a10d65ea3a))
- Incorporates changes from version v1.0.0-beta.2-9-gc785187
  ([4c30a3c](https://github.com/ory/keto/commit/4c30a3c0ad83ba80e1857b41211e7ddade06c4cf))
- Incorporates changes from version v1.0.0-beta.3-1-g06adbf1
  ([0ba3c06](https://github.com/ory/keto/commit/0ba3c0674832b641ef5e0c3f0d60d81ed3a647b2))
- Incorporates changes from version v1.0.0-beta.3-10-g9994967
  ([d2345ca](https://github.com/ory/keto/commit/d2345ca3beb354d6ee7c7926c1a5ddb425d6b405))
- Incorporates changes from version v1.0.0-beta.3-12-gc28b521
  ([b4d792f](https://github.com/ory/keto/commit/b4d792f74055853f05ca46c67625ffd432fc74fd))
- Incorporates changes from version v1.0.0-beta.3-3-g9e16605
  ([c43bf2b](https://github.com/ory/keto/commit/c43bf2b5232bed9106dd47d7eb53d2f93bfe260d))
- Incorporates changes from version v1.0.0-beta.3-5-ga11e898
  ([b9d9b8e](https://github.com/ory/keto/commit/b9d9b8ee33ab957f43f99c427a88ade847e79ed0))
- Incorporates changes from version v1.0.0-beta.3-8-g7061a2a
  ([d76ff9d](https://github.com/ory/keto/commit/d76ff9dc9a4c8a8f1286eeb139d8f5af9617f421))
- Incorporates changes from version v1.0.0-beta.5
  ([0dc314c](https://github.com/ory/keto/commit/0dc314c7888020b40e12eb59fd77135044fd063b))
- Incorporates changes from version v1.0.0-beta.6-1-g5e97104
  ([f14c8ed](https://github.com/ory/keto/commit/f14c8edd7204a811e333ea84429cf837b4e7d27b))
- Incorporates changes from version v1.0.0-beta.8
  ([5045b59](https://github.com/ory/keto/commit/5045b59e2a83d6ab047b1b95c581d7c34e96a2e0))
- Incorporates changes from version v1.0.0-beta.9
  ([be2f035](https://github.com/ory/keto/commit/be2f03524721ef47ecb1c9aec57c2696174e0657))
- Properly sets up changelog TOC
  ([e0acd67](https://github.com/ory/keto/commit/e0acd670ab19c0d6fd36733fea164e2b0414597d))
- Puts toc in the right place
  ([114914f](https://github.com/ory/keto/commit/114914fa354f784b310bc9dfd232a011e0d98d99))
- Revert changes from test release
  ([ab3a64d](https://github.com/ory/keto/commit/ab3a64d3d41292364c5947db98c4d27a8223853e))
- Update documentation links ([#67](https://github.com/ory/keto/issues/67))
  ([d22d413](https://github.com/ory/keto/commit/d22d413c7a001ccaa96b4c013665153f41831614))
- Update link to security console
  ([846ce4b](https://github.com/ory/keto/commit/846ce4baa9da5954bd30996f489885a026c48185))
- Update migration guide
  ([3c44b58](https://github.com/ory/keto/commit/3c44b58613e46ed39d42463537773fe9d95a54da))
- Update to latest changes
  ([1625123](https://github.com/ory/keto/commit/1625123ed342f019d5e7ab440eb37da310570842))
- Updates copyright notice
  ([9dd5578](https://github.com/ory/keto/commit/9dd557825dfd3b9d589c9db2ccb201638debbaae))
- Updates installation guide
  ([f859645](https://github.com/ory/keto/commit/f859645f230f405cfabed0c1b9a2b67b1a3841d3))
- Updates issue and pull request templates
  ([#52](https://github.com/ory/keto/issues/52))
  ([941cae6](https://github.com/ory/keto/commit/941cae6fee058f68eabbbf4dd9cafad4760e108f))
- Updates issue and pull request templates
  ([#53](https://github.com/ory/keto/issues/53))
  ([7b222d2](https://github.com/ory/keto/commit/7b222d285e74c0db482136b23f37072216b3acb0))
- Updates issue and pull request templates
  ([#54](https://github.com/ory/keto/issues/54))
  ([f098639](https://github.com/ory/keto/commit/f098639b5e748151810848fdd3173e0246bc03dc))
- Updates link to guide and header
  ([437c255](https://github.com/ory/keto/commit/437c255ecfff4127fb586cc069e07f86988ad1ba))
- Updates link to open collective
  ([382e1d3](https://github.com/ory/keto/commit/382e1d34c7da0ba0447b78506a749bd7f0085f48))
- Updates links to docs
  ([d84be3b](https://github.com/ory/keto/commit/d84be3b6a8e5eb284ec3fb137ee774ba5ee0d529))
- Updates newsletter link in README
  ([2dc36b2](https://github.com/ory/keto/commit/2dc36b21c8af8e3e39f093198715ea24b65d65af))

### Unclassified

- Add Go SDK factory
  ([99db7e6](https://github.com/ory/keto/commit/99db7e6d4edac88794266a01ddfab9cd0632e95a))
- Add go SDK interface
  ([3dd5f7d](https://github.com/ory/keto/commit/3dd5f7d61bb460c34744b84a34755bfb8219b304))
- Add health handlers
  ([bddb949](https://github.com/ory/keto/commit/bddb949459d05002b0f8882d981e4f63fdddf25f))
- Add policy list handler
  ([a290619](https://github.com/ory/keto/commit/a290619d01d15eb8e3b4e33ede1058d316ee807a))
- Add role iterator in list handler
  ([a3eb696](https://github.com/ory/keto/commit/a3eb6961783f7b562f0a0d0a7e2819bffebce5b8))
- Add SDK generation to circle ci
  ([9b37165](https://github.com/ory/keto/commit/9b37165873bcb0cc5dc60d2514d9824a073466a1))
- Adds ability to update a role using PUT
  ([#14](https://github.com/ory/keto/issues/14))
  ([97ccbe6](https://github.com/ory/keto/commit/97ccbe6d808823c56901ad237878aa6d53cddeeb)):

  - transfer UpdateRoleMembers from https://github.com/ory/hydra/pull/768 to
    keto

  - fix tests by using right http method & correcting sql request

  - Change behavior to overwrite the whole role instead of just the members.

  * small sql migration fix

- Adds log message when telemetry is active
  ([f623c52](https://github.com/ory/keto/commit/f623c52655ff85b7f7209eb73e94eb66a297c5b7))
- Clean up vendor dependencies
  ([9a33c23](https://github.com/ory/keto/commit/9a33c23f4d37ab88b4d643fd79204334d73404c6))
- Do not split empty scope ([#45](https://github.com/ory/keto/issues/45))
  ([b29cf8c](https://github.com/ory/keto/commit/b29cf8cc92607e13457dba8331f5c9286054c8c1))
- Fix typo in help command in env var name
  ([#39](https://github.com/ory/keto/issues/39))
  ([8a5016c](https://github.com/ory/keto/commit/8a5016cd75be78bb42a9a38bfd453ad5722db9db)),
  closes [#25](https://github.com/ory/keto/issues/25)
- Fixes environment variable typos
  ([566d588](https://github.com/ory/keto/commit/566d588e4fca12399966718b725fe4461a28e51e))
- Fixes typo in help command
  ([74450da](https://github.com/ory/keto/commit/74450da18a27513820328c28f72203653c664367)),
  closes [#25](https://github.com/ory/keto/issues/25)
- Format code
  ([637c78c](https://github.com/ory/keto/commit/637c78cba697682b544473a3af9b6ae7715561aa))
- Gofmt
  ([a8d7f9f](https://github.com/ory/keto/commit/a8d7f9f546ae2f3b8c3fa643d8e19b68ca26cc67))
- Improve compose documentation
  ([6870443](https://github.com/ory/keto/commit/68704435f3c299b853f4ff5cacae285b09ada3b5))
- Improves usage of metrics middleware
  ([726c4be](https://github.com/ory/keto/commit/726c4bedfc3f02fdac380930e32f37c251e51aa4))
- Improves usage of metrics middleware
  ([301f386](https://github.com/ory/keto/commit/301f38605af66abae4d28ed0cac90d0b82b655c4))
- Introduce docker-compose file for testing
  ([ba857e3](https://github.com/ory/keto/commit/ba857e3859966e857c5a741825411575e17446de))
- Introduces health and version endpoints
  ([6a9da74](https://github.com/ory/keto/commit/6a9da74f693ee6c15a775ab8d652582aea093601))
- List roles from keto_role table ([#28](https://github.com/ory/keto/issues/28))
  ([9e16605](https://github.com/ory/keto/commit/9e166054b8d474fbce6983d5d00eeeb062fc79b1))
- Properly names flags
  ([af2c3b5](https://github.com/ory/keto/commit/af2c3b5bc96e95fb31b1db5c7fe6dfd6b6fc5b20))
- Require explicit CORS enabling ([#42](https://github.com/ory/keto/issues/42))
  ([9a45107](https://github.com/ory/keto/commit/9a45107af304b2a8e663a532e4f6e4536f15888c))
- Update dependencies
  ([663d8b1](https://github.com/ory/keto/commit/663d8b13e99694a57752cd60a68342b81b041c66))
- Switch to rego as policy decision engine (#48)
  ([ee9bcf2](https://github.com/ory/keto/commit/ee9bcf2719178e5a8dccca083a90313947a8a63b)),
  closes [#48](https://github.com/ory/keto/issues/48)
- Update hydra to v1.0.0-beta.6 ([#35](https://github.com/ory/keto/issues/35))
  ([5e97104](https://github.com/ory/keto/commit/5e971042afff06e2a6ee3b54d2fea31687203623))
- Update npm package registry
  ([a53d3d2](https://github.com/ory/keto/commit/a53d3d23e11fde5dcfbb27a2add1049f4d8e10e6))
- Enable TLS option to serve API (#46)
  ([2f62063](https://github.com/ory/keto/commit/2f620632d0375bf9c7e58dbfb49627c02c66abf3)),
  closes [#46](https://github.com/ory/keto/issues/46)
- Make introspection authorization optional
  ([e5460ad](https://github.com/ory/keto/commit/e5460ad884cd018cd6177324b949cd66bfd53bc7))
- Properly output telemetry information
  ([#33](https://github.com/ory/keto/issues/33))
  ([9994967](https://github.com/ory/keto/commit/9994967b0ca54a62b8b0088fe02be9e890d9574b))
- Remove ORY Hydra dependency ([#44](https://github.com/ory/keto/issues/44))
  ([d487344](https://github.com/ory/keto/commit/d487344fe7e07cb6370371c6b0b6cf3cca767ed1))
- Resolves an issue with the hydra migrate command
  ([2b280bb](https://github.com/ory/keto/commit/2b280bb57c9073a9c8384cde0b14a6991cfacdb6)),
  closes [#23](https://github.com/ory/keto/issues/23)
- Upgrade superagent version ([#41](https://github.com/ory/keto/issues/41))
  ([9c80dbc](https://github.com/ory/keto/commit/9c80dbcc1cc63243839b58ca56ac9be104797887))
- gofmt
  ([777b1be](https://github.com/ory/keto/commit/777b1be1378d314e7cfde0c34450afcce7e590a5))
- Updates README.md (#34)
  ([c28b521](https://github.com/ory/keto/commit/c28b5219fd64314a75ee3c848a80a0c5974ebb7d)),
  closes [#34](https://github.com/ory/keto/issues/34)
- Properly parses cors options
  ([edb5a60](https://github.com/ory/keto/commit/edb5a600f2ce16c0847ee5ef399fa5a41b1e736a))
- Removes additional output if no args are passed
  ([703e124](https://github.com/ory/keto/commit/703e1246ce0fd89066b497c45f0c6cadeb06c331))
- Resolves broken role test
  ([b6c7f9c](https://github.com/ory/keto/commit/b6c7f9c33c4c1f43164d6da0ec7f2553f1f4c598))
- Resolves minor typos and updates install guide
  ([3852be5](https://github.com/ory/keto/commit/3852be56cb81df966a85d4c828de0397d9e74768))
- Updates to latest sqlcon
  ([2c9f643](https://github.com/ory/keto/commit/2c9f643042ff4edffae8bd41834d2a57c923871c))
- Use roles in warden decision
  ([c785187](https://github.com/ory/keto/commit/c785187e31fc7a4b8b762a5e27fac66dcaa97513)),
  closes [#21](https://github.com/ory/keto/issues/21)
  [#19](https://github.com/ory/keto/issues/19)
- authn/client: Payload is now prefixed with client
  ([8584d94](https://github.com/ory/keto/commit/8584d94cfb18deb37ae32ae601f4cd15c14067e7))

# [0.0.1](https://github.com/ory/keto/compare/4f00bc96ece3180a888718ec3c41c69106c86f56...v0.0.1) (2018-05-20)

authn: Checks token_type is "access_token", if set

Closes #1

### Documentation

- Incorporates changes from version
  ([b5445a0](https://github.com/ory/keto/commit/b5445a0fc5b6f813cd1731b20c8c5c79d7c4cdf8))
- Incorporates changes from version
  ([295ff99](https://github.com/ory/keto/commit/295ff998af55777823b04f423e365fd58e61753b))
- Incorporates changes from version
  ([bd44d41](https://github.com/ory/keto/commit/bd44d41b2781e33353082397c47390a27f749e16))
- Updates readme and upgrades
  ([0f95dbb](https://github.com/ory/keto/commit/0f95dbb967fd17b607caa999ae30453f5f599739))
- Uses keto repo for changelog
  ([14c0b2a](https://github.com/ory/keto/commit/14c0b2a2bd31566f2b9048831f894aba05c5b15d))

### Unclassified

- Adds migrate commands to the proper parent command
  ([231c70d](https://github.com/ory/keto/commit/231c70d816b0736a51eddc1fa0445bac672b1b2f))
- Checks token_type is "access_token", if set
  ([d2b8f5d](https://github.com/ory/keto/commit/d2b8f5d313cce597566bd18e4f3bea4a423a62ee)),
  closes [#1](https://github.com/ory/keto/issues/1)
- Removes old test
  ([07b733b](https://github.com/ory/keto/commit/07b733bfae4b733e3e2124545b92c537dabbdcf0))
- Renames subject to sub in response payloads
  ([ca4d540](https://github.com/ory/keto/commit/ca4d5408000be2b896d38eaaf5e67a3fc0a566da))
- Tells linguist to ignore SDK files
  ([f201eb9](https://github.com/ory/keto/commit/f201eb95f3309a60ac50f42cfba0bae2e38e8d13))
- Retries SQL connection on migrate commands
  ([3d33d73](https://github.com/ory/keto/commit/3d33d73c009077c5bf30ae4b03802904bfb5d5b2)):

  This patch also introduces a fatal error if migrations fail

- cmd/server: Resolves DBAL not handling postgres properly
  ([dedc32a](https://github.com/ory/keto/commit/dedc32ab218923243b1955ce5bcbbdc5cc416953))
- cmd/server: Improves error message in migrate command
  ([4b17ce8](https://github.com/ory/keto/commit/4b17ce8848113cae807840182d1a318190c2a9b3))
- Resolves travis and docker issues
  ([6f4779c](https://github.com/ory/keto/commit/6f4779cc51bf4f2ee5b97541fb77d8f882497710))
- Adds OAuth2 Client Credentials authenticator and warden endpoint
  ([c55139b](https://github.com/ory/keto/commit/c55139b51e636834759706499a2aec1451f4fbd9))
- Adds SDK helpers
  ([a1c2608](https://github.com/ory/keto/commit/a1c260801d9366fccf4bfb4fc64b2c67fc594565))
- Resolves SDK and test issues (#4)
  ([2d4cd98](https://github.com/ory/keto/commit/2d4cd9805af3081bbcbea3f806ca066d35385a4b)),
  closes [#4](https://github.com/ory/keto/issues/4)
- Initial project commit
  ([a592e51](https://github.com/ory/keto/commit/a592e5126f130f8b673fff6c894fdbd9fb56f81c))
- Initial commit
  ([4f00bc9](https://github.com/ory/keto/commit/4f00bc96ece3180a888718ec3c41c69106c86f56))

---

id: changelog title: Changelog custom_edit_url: null

---

# [Unreleased](https://github.com/ory/keto/compare/v0.6.0-alpha.3...8e301198298858fd7f387ef63a7abf4fa55ea240) (2021-06-22)

### Bug Fixes

- Add missing tracers ([#600](https://github.com/ory/keto/issues/600))
  ([aa263be](https://github.com/ory/keto/commit/aa263be9a7830e3c769d7698d36137555ca230bc)),
  closes [#593](https://github.com/ory/keto/issues/593)
- Handle relation tuple cycles in expand and check engine
  ([#623](https://github.com/ory/keto/issues/623))
  ([8e30119](https://github.com/ory/keto/commit/8e301198298858fd7f387ef63a7abf4fa55ea240))
- Log all database connection errors
  ([#588](https://github.com/ory/keto/issues/588))
  ([2b0fad8](https://github.com/ory/keto/commit/2b0fad897e61400bd2a6cdf47f33ff4301e9c5f8))
- Move gRPC client module root up
  ([#620](https://github.com/ory/keto/issues/620))
  ([3b881f6](https://github.com/ory/keto/commit/3b881f6015a93b382b3fbbca4be9259622038b6a)):

  BREAKING: The npm package `@ory/keto-grpc-client` from now on includes all API
  versions. Because of that, the import paths changed. For migrating to the new
  client package, change the import path according to the following example:

  ```diff
  - import acl from '@ory/keto-grpc-client/acl_pb.js'
  + // from the latest version
  + import { acl } from '@ory/keto-grpc-client'
  + // or a specific one
  + import acl from '@ory/keto-grpc-client/ory/keto/acl/v1alpha1/acl_pb.js'
  ```

- Update docker-compose.yml version
  ([#595](https://github.com/ory/keto/issues/595))
  ([7fa4dca](https://github.com/ory/keto/commit/7fa4dca4182a1fa024f9cef0a04163f2cbd882aa)),
  closes [#549](https://github.com/ory/keto/issues/549)

### Documentation

- Fix example not following best practice
  ([#582](https://github.com/ory/keto/issues/582))
  ([a015818](https://github.com/ory/keto/commit/a0158182c5f87cfd4767824e1c5d6cbb8094a4e6))
- Update NPM links due to organisation move
  ([#616](https://github.com/ory/keto/issues/616))
  ([6355bea](https://github.com/ory/keto/commit/6355beae5b5b28c3eee19fdee85b9875cbc165c3))

### Features

- Make generated gRPC client its own module
  ([#583](https://github.com/ory/keto/issues/583))
  ([f0fbb64](https://github.com/ory/keto/commit/f0fbb64b3358e9800854295cebc9ec8b8e56c87a))
- Max_idle_conn_time ([#605](https://github.com/ory/keto/issues/605))
  ([50a8623](https://github.com/ory/keto/commit/50a862338e17f86900ca162da7f3467f55f9f954)),
  closes [#523](https://github.com/ory/keto/issues/523)

# [0.6.0-alpha.3](https://github.com/ory/keto/compare/v0.6.0-alpha.2...v0.6.0-alpha.3) (2021-04-29)

Resolves CRDB and build issues.

### Code Generation

- Pin v0.6.0-alpha.3 release commit
  ([d766968](https://github.com/ory/keto/commit/d766968419d10a68fd843df45316e3436b68d61d))

# [0.6.0-alpha.2](https://github.com/ory/keto/compare/v0.6.0-alpha.1...v0.6.0-alpha.2) (2021-04-29)

This release improves stability and documentation.

### Bug Fixes

- Add npm run format to make format
  ([7d844a8](https://github.com/ory/keto/commit/7d844a8e6412ae561963b97ac26d4682411095d4))
- Makefile target
  ([0e6f612](https://github.com/ory/keto/commit/0e6f6122de7bdbb691ad7cc236b6bc9a3601d39e))
- Move swagger to spec dir
  ([7f6a061](https://github.com/ory/keto/commit/7f6a061aafda275d278bf60f16e90039da45bc57))
- Resolve clidoc issues
  ([ef12b4e](https://github.com/ory/keto/commit/ef12b4e267f34fbf9709fe26023f9b7ae6670c24))
- Update install.sh ([#568](https://github.com/ory/keto/issues/568))
  ([86ab245](https://github.com/ory/keto/commit/86ab24531d608df0b5391ee8ec739291b9a90e20))
- Use correct id
  ([5e02902](https://github.com/ory/keto/commit/5e029020b5ba3931f15d343cf6a9762b064ffd45))
- Use correct id for api
  ([32a6b04](https://github.com/ory/keto/commit/32a6b04609054cba84f7b56ebbe92341ec5dcd98))
- Use sqlite image versions ([#544](https://github.com/ory/keto/issues/544))
  ([ec6cc5e](https://github.com/ory/keto/commit/ec6cc5ed528f1a097ea02669d059e060b7eff824))

### Code Generation

- Pin v0.6.0-alpha.2 release commit
  ([470b2c6](https://github.com/ory/keto/commit/470b2c61c649fe5fcf638c84d4418212ff0330a5))

### Documentation

- Add gRPC client README.md ([#559](https://github.com/ory/keto/issues/559))
  ([9dc3596](https://github.com/ory/keto/commit/9dc35969ada8b0d4d73dee9089c4dc61cd9ea657))
- Change forum to discussions readme
  ([#539](https://github.com/ory/keto/issues/539))
  ([ea2999d](https://github.com/ory/keto/commit/ea2999d4963316810a8d8634fcd123bda31eaa8f))
- Fix cat videos example docker compose
  ([#549](https://github.com/ory/keto/issues/549))
  ([b25a711](https://github.com/ory/keto/commit/b25a7114631957935c71ac6a020ab6bd0c244cd7))
- Fix typo ([#538](https://github.com/ory/keto/issues/538))
  ([99a9693](https://github.com/ory/keto/commit/99a969373497792bb4cd8ff62bf5245087517737))
- Include namespace in olymp library example
  ([#540](https://github.com/ory/keto/issues/540))
  ([135e814](https://github.com/ory/keto/commit/135e8145c383a76b494b469253c949c38f4414a7))
- Update install from source steps to actually work
  ([#548](https://github.com/ory/keto/issues/548))
  ([e662256](https://github.com/ory/keto/commit/e6622564f58b7612b13b11b54e75a7350f52d6de))

### Features

- Global docs sidebar and added cloud pages
  ([c631c82](https://github.com/ory/keto/commit/c631c82b7ff3d12734869ac22730b52e73dcf287))
- Support retryable CRDB transactions
  ([833147d](https://github.com/ory/keto/commit/833147dae40e9ac5bdf220f8aa3f01abd444f791))

# [0.6.0-alpha.1](https://github.com/ory/keto/compare/v0.5.6-alpha.1...v0.6.0-alpha.1) (2021-04-07)

We are extremely happy to announce next-gen Ory Keto which implements
[Zanzibar: Google’s Consistent, Global Authorization System](https://research.google/pubs/pub48190/):

> Zanzibar provides a uniform data model and configuration language for
> expressing a wide range of access control policies from hundreds of client
> services at Google, including Calendar, Cloud, Drive, Maps, Photos, and
> YouTube. Its authorization decisions respect causal ordering of user actions
> and thus provide external consistency amid changes to access control lists and
> object contents. Zanzibar scales to trillions of access control lists and
> millions of authorization requests per second to support services used by
> billions of people. It has maintained 95th-percentile latency of less than 10
> milliseconds and availability of greater than 99.999% over 3 years of
> production use.

Ory Keto is the first open source planet-scale authorization system built with
cloud native technologies (Go, gRPC, newSQL) and architecture. It is also the
first open source implementation of Google Zanzibar :tada:!

Many concepts developer by Google Zanzibar are implemented in Ory Keto already.
Let's take a look!

As of this release, Ory Keto knows how to interpret and operate on the basic
access control lists known as relation tuples. They encode relations between
objects and subjects. One simple example of such a relation tuple could encode
"`user1` has access to file `/foo`", a more complex one could encode "everyone
who has write access on `/foo` has read access on `/foo`".

Ory Keto comes with all the basic APIs as described in the Zanzibar paper. All
of them are available over gRPC and REST.

1. List: query relation tuples
2. Check: determine whether a subject has a relation on an object
3. Expand: get a tree of all subjects who have a relation on an object
4. Change: create, update, and delete relation tuples

For all details, head over to the
[documentation](https://www.ory.sh/keto/docs/concepts/api-overview).

With this release we officially move the "old" Keto to the
[legacy-0.5 branch](https://github.com/ory/keto/tree/legacy-0.5). We will only
provide security fixes from now on. A migration path to v0.6 is planned but not
yet implemented, as the architectures are vastly different. Please refer to
[the issue](https://github.com/ory/keto/issues/318).

We are keen to bring more features and performance improvements. The next
features we will tackle are:

- Subject Set rewrites
- Native ABAC & RBAC Support
- Integration with other policy servers
- Latency reduction through aggressive caching
- Cluster mode that fans out requests over all Keto instances

So stay tuned, :star: this repo, :eyes: releases, and
[subscribe to our newsletter :email:](https://ory.us10.list-manage.com/subscribe?u=ffb1a878e4ec6c0ed312a3480&id=f605a41b53&MERGE0=&group[17097][32]=1).

### Bug Fixes

- Add description attribute to access control policy role
  ([#215](https://github.com/ory/keto/issues/215))
  ([831eba5](https://github.com/ory/keto/commit/831eba59f810ca68561dd584c9df7684df10b843))
- Add leak_sensitive_values to config schema
  ([2b21d2b](https://github.com/ory/keto/commit/2b21d2bdf4ca9523d16159c5f73c4429b692e17d))
- Bump CLI
  ([80c82d0](https://github.com/ory/keto/commit/80c82d026cbfbab8fbb84d850d8980866ecf88df))
- Bump deps and replace swagutil
  ([#212](https://github.com/ory/keto/issues/212))
  ([904258d](https://github.com/ory/keto/commit/904258d23959c3fa96b6d8ccfdb79f6788c106ec))
- Check engine overwrote result in some cases
  ([#412](https://github.com/ory/keto/issues/412))
  ([3404492](https://github.com/ory/keto/commit/3404492002ca5c3f017ef25486e377e911987aa4))
- Check health status in status command
  ([21c64d4](https://github.com/ory/keto/commit/21c64d45f21a505744b9f70d780f9b3079d3822c))
- Check REST API returns JSON object
  ([#460](https://github.com/ory/keto/issues/460))
  ([501dcff](https://github.com/ory/keto/commit/501dcff4427f76902671f6d5733f28722bd51fa7)),
  closes [#406](https://github.com/ory/keto/issues/406)
- Empty relationtuple list should not error
  ([#440](https://github.com/ory/keto/issues/440))
  ([fbcb3e1](https://github.com/ory/keto/commit/fbcb3e1f337b5114d7697fa512ded92b5f409ef4))
- Ensure nil subject is not allowed
  ([#449](https://github.com/ory/keto/issues/449))
  ([7a0fcfc](https://github.com/ory/keto/commit/7a0fcfc4fe83776fa09cf78ee11f407610554d04)):

  The nodejs gRPC client was a great fuzzer and pointed me to some nil pointer
  dereference panics. This adds some input validation to prevent panics.

- Ensure persister errors are handled by sqlcon
  ([#473](https://github.com/ory/keto/issues/473))
  ([4343c4a](https://github.com/ory/keto/commit/4343c4acd8f917fb7ae131e67bca6855d4d61694))
- Handle pagination and errors in the check/expand engines
  ([#398](https://github.com/ory/keto/issues/398))
  ([5eb1a7d](https://github.com/ory/keto/commit/5eb1a7d49af6b43707c122de8727cbd72285cb5c))
- Ignore dist
  ([ba816ea](https://github.com/ory/keto/commit/ba816ea2ca39962f02c08e0c7b75cfe3cf1d963d))
- Ignore x/net false positives
  ([d8b36cb](https://github.com/ory/keto/commit/d8b36cb1812abf7265ac15c29780222be025186b))
- Improve CLI remote sourcing ([#474](https://github.com/ory/keto/issues/474))
  ([a85f4d7](https://github.com/ory/keto/commit/a85f4d7470ac3744476e82e5889b97d5a0680473))
- Improve handlers and add tests
  ([#470](https://github.com/ory/keto/issues/470))
  ([ca5ccb9](https://github.com/ory/keto/commit/ca5ccb9c237fdcc4db031ec97a75616a859cbf8f))
- Insert relation tuples without fmt.Sprintf
  ([#443](https://github.com/ory/keto/issues/443))
  ([fe507bb](https://github.com/ory/keto/commit/fe507bb4ea719780e732d098291aa190d6b1c441))
- Minor bugfixes ([#371](https://github.com/ory/keto/issues/371))
  ([185ee1e](https://github.com/ory/keto/commit/185ee1e51bc4bcdee028f71fcaf207b7e342313b))
- Move dockerfile to where it belongs
  ([f087843](https://github.com/ory/keto/commit/f087843ac8f24e741bf39fe65ee5b0a9adf9a5bb))
- Namespace migrator ([#417](https://github.com/ory/keto/issues/417))
  ([ea79300](https://github.com/ory/keto/commit/ea7930064f490b063a712b4e18521f8996931a13)),
  closes [#404](https://github.com/ory/keto/issues/404)
- Remove SQL logging ([#455](https://github.com/ory/keto/issues/455))
  ([d8e2a86](https://github.com/ory/keto/commit/d8e2a869db2a9cfb44423b434330536036b2f421))
- Rename /relationtuple endpoint to /relation-tuples
  ([#519](https://github.com/ory/keto/issues/519))
  ([8eb55f6](https://github.com/ory/keto/commit/8eb55f6269399f2bc5f000b8a768bcdf356c756f))
- Resolve gitignore build
  ([6f04bbb](https://github.com/ory/keto/commit/6f04bbb6057779b4d73d3f94677cea365843f7ac))
- Resolve goreleaser issues
  ([d32767f](https://github.com/ory/keto/commit/d32767f32856cf5bd48514c5d61746417fbed6f5))
- Resolve windows build issues
  ([8bcdfbf](https://github.com/ory/keto/commit/8bcdfbfbdb0b10c03ff93838e8fe6e778236e96d))
- Rewrite check engine to search starting at the object
  ([#310](https://github.com/ory/keto/issues/310))
  ([7d99694](https://github.com/ory/keto/commit/7d9969414ebc8cf6ef5d211ad34f8ae01bd3b4ee)),
  closes [#302](https://github.com/ory/keto/issues/302)
- Secure query building ([#442](https://github.com/ory/keto/issues/442))
  ([c7d2770](https://github.com/ory/keto/commit/c7d2770ed570238fd1262bcc4e5b4afa6c12d80e))
- Strict version enforcement in docker
  ([e45b28f](https://github.com/ory/keto/commit/e45b28fec626db35f1bd4580e5b11c9c94a02669))
- Update dd-trace to fix build issues
  ([2ad489f](https://github.com/ory/keto/commit/2ad489f0d9cae3191718d36823fe25df58ab95e6))
- Update docker to go 1.16 and alpine
  ([c63096c](https://github.com/ory/keto/commit/c63096cb53d2171f22f4a0d4a9ac3c9bfac89d01))
- Use errors.WithStack everywhere
  ([#462](https://github.com/ory/keto/issues/462))
  ([5f25bce](https://github.com/ory/keto/commit/5f25bceea35179c67d24dd95f698dc57b789d87a)),
  closes [#437](https://github.com/ory/keto/issues/437):

  Fixed all occurrences found using the search pattern `return .*, err\n`.

- Use package name in pkger
  ([6435939](https://github.com/ory/keto/commit/6435939ad7e5899505cd0e6261f5dfc819c9ca42))
- **schema:** Add trace level to logger
  ([a5a1402](https://github.com/ory/keto/commit/a5a1402c61e1a37b1a9a349ad5736eaca66bd6a4))
- Use make() to initialize slices
  ([#250](https://github.com/ory/keto/issues/250))
  ([84f028d](https://github.com/ory/keto/commit/84f028dc35665174542e103c0aefc635bb6d3e52)),
  closes [#217](https://github.com/ory/keto/issues/217)

### Build System

- Pin dependency versions of buf and protoc plugins
  ([#338](https://github.com/ory/keto/issues/338))
  ([5a2fd1c](https://github.com/ory/keto/commit/5a2fd1cc8dff02aa7017771adc0d9101f6c86775))

### Code Generation

- Pin v0.6.0-alpha.1 release commit
  ([875af25](https://github.com/ory/keto/commit/875af25f89b813455148e58884dcdf1cd3600b86))

### Code Refactoring

- Data structures ([#279](https://github.com/ory/keto/issues/279))
  ([1316077](https://github.com/ory/keto/commit/131607762d0006e4cf4f93e8731ef7648348b2ec))

### Documentation

- Add check- and expand-API guides
  ([#493](https://github.com/ory/keto/issues/493))
  ([09a25b4](https://github.com/ory/keto/commit/09a25b4063abcfdcd4c0de315a2ef088d6d4e72e))
- Add current features overview ([#505](https://github.com/ory/keto/issues/505))
  ([605afa0](https://github.com/ory/keto/commit/605afa029794ad115bba02e004e1596cea038e8e))
- Add missing pages ([#518](https://github.com/ory/keto/issues/518))
  ([43cbaa9](https://github.com/ory/keto/commit/43cbaa9140cfa0ea3c72f699f6bb34f5ed31d8dd))
- Add namespace and relation naming conventions
  ([#510](https://github.com/ory/keto/issues/510))
  ([dd31865](https://github.com/ory/keto/commit/dd318653178cd45da47f3e7cef507b42708363ef))
- Add performance page ([#413](https://github.com/ory/keto/issues/413))
  ([6fe0639](https://github.com/ory/keto/commit/6fe0639d36087b5ecd555eb6fe5ce949f3f6f0d7)):

  This also refactored the server startup. Functionality did not change.

- Add production guide
  ([a9163c7](https://github.com/ory/keto/commit/a9163c7690c55c8191650c4dfb464b75ea02446b))
- Add zanzibar overview to README.md
  ([#265](https://github.com/ory/keto/issues/265))
  ([15a95b2](https://github.com/ory/keto/commit/15a95b28e745592353e4656d42a9d0bd20ce468f))
- API overview ([#501](https://github.com/ory/keto/issues/501))
  ([05fe03b](https://github.com/ory/keto/commit/05fe03b5bf7a3f790aa6c9c1d3fcdb31304ef6af))
- Concepts ([#429](https://github.com/ory/keto/issues/429))
  ([2f2c885](https://github.com/ory/keto/commit/2f2c88527b3f6d1d46a5c287d8aca0874d18a28d))
- Delete old redirect homepage
  ([c0a3784](https://github.com/ory/keto/commit/c0a378448f8c7723bae68f7b52a019b697b25863))
- Document gRPC SKDs
  ([7583fe8](https://github.com/ory/keto/commit/7583fe8933f6676b4e37477098b1d43d12819b8b))
- Fix grammatical error ([#222](https://github.com/ory/keto/issues/222))
  ([256a0d2](https://github.com/ory/keto/commit/256a0d2e53fe1eb859e41fc539870ae1d5a493d2))
- Fix regression issues
  ([9697bb4](https://github.com/ory/keto/commit/9697bb43dd23c0d1fae74ea42e848883c45dae77))
- Generate gRPC reference page ([#488](https://github.com/ory/keto/issues/488))
  ([93ebe6d](https://github.com/ory/keto/commit/93ebe6db7e887d708503a54c5ec943254e37ca43))
- Improve CLI documentation ([#503](https://github.com/ory/keto/issues/503))
  ([be9327f](https://github.com/ory/keto/commit/be9327f7b28152a78f731043acf83b7092e42e29))
- Minor fixes ([#532](https://github.com/ory/keto/issues/532))
  ([638342e](https://github.com/ory/keto/commit/638342eb9519d9bf609926fb87558071e2815fb3))
- Move development section
  ([9ff393f](https://github.com/ory/keto/commit/9ff393f6cba1fb0a33918377ce505455c34d9dfc))
- Move to json sidebar
  ([257bf96](https://github.com/ory/keto/commit/257bf96044df37c3d7af8a289fb67098d48da1a3))
- Remove duplicate "is"
  ([ca3277d](https://github.com/ory/keto/commit/ca3277d82c1508797bc8c663963407d2e4d9112f))
- Remove duplicate template
  ([1d3b38e](https://github.com/ory/keto/commit/1d3b38e4045b0b874bb1186ea628f5a37353a2e6))
- Remove old documentation ([#426](https://github.com/ory/keto/issues/426))
  ([eb76913](https://github.com/ory/keto/commit/eb7691306018678e024211b51627a1c27e780a6b))
- Replace TODO links ([#512](https://github.com/ory/keto/issues/512))
  ([ad8e20b](https://github.com/ory/keto/commit/ad8e20b3bef2bc46b3a32c2c9ccb6e16e4bad22c))
- Resolve broken links
  ([0d0a50b](https://github.com/ory/keto/commit/0d0a50b3f4112893f32c81adc8edd137b5a62541))
- Simple access check guide ([#451](https://github.com/ory/keto/issues/451))
  ([e0485af](https://github.com/ory/keto/commit/e0485afc46a445868580aa541e962e80cbea0670)):

  This also enables gRPC go, gRPC nodejs, cURL, and Keto CLI code samples to be
  tested.

- Update comment in write response
  ([#329](https://github.com/ory/keto/issues/329))
  ([4ca0baf](https://github.com/ory/keto/commit/4ca0baf62e34402e749e870fe8c0cc893684192c))
- Update install instructions
  ([d2e4123](https://github.com/ory/keto/commit/d2e4123f3e2e58da8be181a0a542e3dcc1313e16))
- Update introduction
  ([5f71d73](https://github.com/ory/keto/commit/5f71d73e2ee95d02abc4cd42a76c98a35942df0c))
- Update README ([#515](https://github.com/ory/keto/issues/515))
  ([18d3cd6](https://github.com/ory/keto/commit/18d3cd61b0a79400170dc0f89860b4614cc4a543)):

  Also format all markdown files in the root.

- Update repository templates
  ([db505f9](https://github.com/ory/keto/commit/db505f9e10755bc20c4623c4f5f99f33283dffda))
- Update repository templates
  ([6c056bb](https://github.com/ory/keto/commit/6c056bb2043af6e82f06fdfa509ab3fa0d5e5d06))
- Update SDK links ([#514](https://github.com/ory/keto/issues/514))
  ([f920fbf](https://github.com/ory/keto/commit/f920fbfc8dcc6711ad9e046578a4506179952be7))
- Update swagger documentation for REST endpoints
  ([c363de6](https://github.com/ory/keto/commit/c363de61edf912fef85acc6bcdac6e1c15c48f4f))
- Use mdx for api reference
  ([340f3a3](https://github.com/ory/keto/commit/340f3a3dd20c82c743e7b3ad6aaf06a4c118b5a1))
- Various improvements and updates
  ([#486](https://github.com/ory/keto/issues/486))
  ([a812ace](https://github.com/ory/keto/commit/a812ace2303214e0e7acb2e283efa1cff0d5d279))

### Features

- Add .dockerignore
  ([8b0ff06](https://github.com/ory/keto/commit/8b0ff066b2508ef2f3629f9a3e2fce601b8dcce1))
- Add and automate version schema
  ([b01eef8](https://github.com/ory/keto/commit/b01eef8d4d5834b5888cb369ecf01ee01b40c24c))
- Add check engine ([#277](https://github.com/ory/keto/issues/277))
  ([396c1ae](https://github.com/ory/keto/commit/396c1ae33b777031f8d59549d9de4a88e3f6b10a))
- Add gRPC health status ([#427](https://github.com/ory/keto/issues/427))
  ([51c4223](https://github.com/ory/keto/commit/51c4223d6cb89a9bfbc115ef20db8350aeb2e8af))
- Add is_last_page to list response
  ([#425](https://github.com/ory/keto/issues/425))
  ([b73d91f](https://github.com/ory/keto/commit/b73d91f061ab155c53d802263c0263aa39e64bdf))
- Add POST REST handler for policy check
  ([7d89860](https://github.com/ory/keto/commit/7d89860bc4a790a69f5bea5b0dbe4a2938c6729f))
- Add relation write API ([#275](https://github.com/ory/keto/issues/275))
  ([f2ddb9d](https://github.com/ory/keto/commit/f2ddb9d884ed71037b5371c00bb10b63d25d47c0))
- Add REST and gRPC logger middlewares
  ([#436](https://github.com/ory/keto/issues/436))
  ([615eb0b](https://github.com/ory/keto/commit/615eb0bec3bdc0fd26abc7af0b8990269b0cbedd))
- Add SQA telemetry ([#535](https://github.com/ory/keto/issues/535))
  ([9f6472b](https://github.com/ory/keto/commit/9f6472b0c996505d41058e9b55afa8fd6b9bb2d5))
- Add sql persister ([#350](https://github.com/ory/keto/issues/350))
  ([d595d52](https://github.com/ory/keto/commit/d595d52dabb8f4953b5c23d3a8154cac13d00306))
- Add tracing ([#536](https://github.com/ory/keto/issues/536))
  ([b57a144](https://github.com/ory/keto/commit/b57a144e0a7ec39d5831dbb79840c2b25c044e6a))
- Allow to apply namespace migrations together with regular migrations
  ([#441](https://github.com/ory/keto/issues/441))
  ([57e2bbc](https://github.com/ory/keto/commit/57e2bbc5eaebe43834f2432eb1ee2820d9cb2988))
- Delete relation tuples ([#457](https://github.com/ory/keto/issues/457))
  ([3ec8afa](https://github.com/ory/keto/commit/3ec8afa68c5b5ddc26609b9afd17cc0d06cd82bf)),
  closes [#452](https://github.com/ory/keto/issues/452)
- Dockerfile and docker compose example
  ([#390](https://github.com/ory/keto/issues/390))
  ([10cd0b3](https://github.com/ory/keto/commit/10cd0b39c12ef96710bda6ff013f7c5eeae97118))
- Expand API ([#285](https://github.com/ory/keto/issues/285))
  ([a3ca0b8](https://github.com/ory/keto/commit/a3ca0b8a109b63f443e359cd8ff18a7b3e489f84))
- Expand GPRC service and CLI ([#383](https://github.com/ory/keto/issues/383))
  ([acf2154](https://github.com/ory/keto/commit/acf21546d3e135deb77c853b751a3da3a7b16f00))
- First API draft and generation
  ([#315](https://github.com/ory/keto/issues/315))
  ([bda5d8b](https://github.com/ory/keto/commit/bda5d8b7e90d749600f5b5e169df8a6ec3705b22))
- GRPC status codes and improved error messages
  ([#467](https://github.com/ory/keto/issues/467))
  ([4a4f8c6](https://github.com/ory/keto/commit/4a4f8c6b323664329414b61e7d80d7838face730))
- GRPC version API ([#475](https://github.com/ory/keto/issues/475))
  ([89cc46f](https://github.com/ory/keto/commit/89cc46fe4a13b062693d3db4f803834ba37f4e48))
- Implement goreleaser pipeline
  ([888ac43](https://github.com/ory/keto/commit/888ac43e6f706f619b2f1b58271dd027094c9ae9)),
  closes [#410](https://github.com/ory/keto/issues/410)
- Incorporate new GRPC API structure
  ([#331](https://github.com/ory/keto/issues/331))
  ([e0916ad](https://github.com/ory/keto/commit/e0916ad00c81b24177cfe45faf77b93d2c33dc1f))
- Koanf and namespace configuration
  ([#367](https://github.com/ory/keto/issues/367))
  ([3ad32bc](https://github.com/ory/keto/commit/3ad32bc13a4d96135be8031eb6fe4c15868272ca))
- Namespace configuration ([#324](https://github.com/ory/keto/issues/324))
  ([b94f50d](https://github.com/ory/keto/commit/b94f50d1800c47a43561df5009cb38b44ccd0088))
- Namespace migrate status CLI ([#508](https://github.com/ory/keto/issues/508))
  ([e3f7ad9](https://github.com/ory/keto/commit/e3f7ad91585b616e97f85ce0f55c76406b6c4d0a)):

  This also refactors the current `migrate` and `namespace migrate` commands.

- Nodejs gRPC definitions ([#447](https://github.com/ory/keto/issues/447))
  ([3b5c313](https://github.com/ory/keto/commit/3b5c31326645adb2d5b14ced901771a7ba00fd1c)):

  Includes Typescript definitions.

- Read API ([#269](https://github.com/ory/keto/issues/269))
  ([de5119a](https://github.com/ory/keto/commit/de5119a6e3c7563cfc2e1ada12d47b27ebd7faaa)):

  This is a first draft of the read API. It is reachable by REST and gRPC calls.
  The main purpose of this PR is to establish the basic repository structure and
  define the API.

- Relationtuple parse command ([#490](https://github.com/ory/keto/issues/490))
  ([91a3cf4](https://github.com/ory/keto/commit/91a3cf47fbdb8203b799cf7c69bcf3dbbfb98b3a)):

  This command parses the relation tuple format used in the docs. It greatly
  improves the experience when copying something from the documentation. It can
  especially be used to pipe relation tuples into other commands, e.g.:

  ```shell
  echo "messages:02y_15_4w350m3#decypher@john" | \
    keto relation-tuple parse - --format json | \
    keto relation-tuple create -
  ```

- REST patch relation tuples ([#491](https://github.com/ory/keto/issues/491))
  ([d38618a](https://github.com/ory/keto/commit/d38618a9e647902ce019396ff1c33973020bf797)):

  The new PATCH handler allows transactional changes similar to the already
  existing gRPC service.

- Separate and multiplex ports based on read/write privilege
  ([#397](https://github.com/ory/keto/issues/397))
  ([6918ac3](https://github.com/ory/keto/commit/6918ac3bfa355cbd551e44376c214f412e3414e4))
- Swagger SDK ([#476](https://github.com/ory/keto/issues/476))
  ([011888c](https://github.com/ory/keto/commit/011888c2b7e2d0f7b8923c994c70e62d374a2830))

### Tests

- Add command tests ([#487](https://github.com/ory/keto/issues/487))
  ([61c28e4](https://github.com/ory/keto/commit/61c28e48a5c3f623e5cc133e69ba368c5103f414))
- Add dedicated persistence tests
  ([#416](https://github.com/ory/keto/issues/416))
  ([4e98906](https://github.com/ory/keto/commit/4e9890605edf3ea26134917a95bfa6fbb176565e))
- Add handler tests ([#478](https://github.com/ory/keto/issues/478))
  ([9315a77](https://github.com/ory/keto/commit/9315a77820d50e400b78f2f019a871be022a9887))
- Add initial e2e test ([#380](https://github.com/ory/keto/issues/380))
  ([dc5d3c9](https://github.com/ory/keto/commit/dc5d3c9d02604fddbfa56ac5ebbc1fef56a881d9))
- Add relationtuple definition tests
  ([#415](https://github.com/ory/keto/issues/415))
  ([2e3dcb2](https://github.com/ory/keto/commit/2e3dcb200a7769dc8710d311ca08a7515012fbdd))
- Enable GRPC client in e2e test
  ([#382](https://github.com/ory/keto/issues/382))
  ([4e5c6ae](https://github.com/ory/keto/commit/4e5c6aed56e5a449003956ec114ec131be068aaf))
- Improve docs sample tests ([#461](https://github.com/ory/keto/issues/461))
  ([6e0e5e6](https://github.com/ory/keto/commit/6e0e5e6184916e894fd4694cfa3a158f11fae11f))

# [0.5.6-alpha.1](https://github.com/ory/keto/compare/v0.5.5-alpha.1...v0.5.6-alpha.1) (2020-05-28)

This release bumps vulnerable transient dependencies (those are not actually
used in ORY Keto) and updates several documentation pages and improves
structured logging output. Additionally, ORY Keto now uses the updated release
pipeline!

### Bug Fixes

- Update install script
  ([21e1bf0](https://github.com/ory/keto/commit/21e1bf05177576a9d743bd11744ef6a42be50b8d))

### Chores

- Pin v0.5.6-alpha.1 release commit
  ([ed0da08](https://github.com/ory/keto/commit/ed0da08a03a910660358fc56c568692325749b6d))

# [0.5.5-alpha.1](https://github.com/ory/keto/compare/v0.5.4-alpha.1...v0.5.5-alpha.1) (2020-05-28)

This release bumps vulnerable transient dependencies (those are not actually
used in ORY Keto) and updates several documentation pages and improves
structured logging output. Additionally, ORY Keto now uses the updated release
pipeline!

### Bug Fixes

- Move deps to go_mod_indirect_pins
  ([dd3e971](https://github.com/ory/keto/commit/dd3e971ac418baf10c1b33005acc7e6f66fb0d85))
- Resolve test issues
  ([9bd9956](https://github.com/ory/keto/commit/9bd9956e33731f1619c32e1e6b7c78f42e7c47c3))
- Update install.sh script
  ([f64d320](https://github.com/ory/keto/commit/f64d320b6424fe3256eb7fad1c94dcc1ef0bf487))
- Use semver-regex replacer func
  ([2cc3bbb](https://github.com/ory/keto/commit/2cc3bbb2d75ba5fa7a3653d7adcaa712ff38c603))

### Chores

- Pin v0.5.5-alpha.1 release commit
  ([4666a0f](https://github.com/ory/keto/commit/4666a0f258f253d19a14eca34f4b7049f2d0afa2))

### Documentation

- Add missing colon in docker run command
  ([#193](https://github.com/ory/keto/issues/193))
  ([383063d](https://github.com/ory/keto/commit/383063d260d995665da4c02c9a7bac7e06a2c8d3))
- Update github templates ([#182](https://github.com/ory/keto/issues/182))
  ([72ea09b](https://github.com/ory/keto/commit/72ea09bbbf9925d7705842703b32826376f636e4))
- Update github templates ([#184](https://github.com/ory/keto/issues/184))
  ([ed546b7](https://github.com/ory/keto/commit/ed546b7a2b9ee690284a48c641edd1570464d71f))
- Update github templates ([#188](https://github.com/ory/keto/issues/188))
  ([ebd75b2](https://github.com/ory/keto/commit/ebd75b2f6545ff4372773f6370300c7b2ca71c51))
- Update github templates ([#189](https://github.com/ory/keto/issues/189))
  ([fd4c0b1](https://github.com/ory/keto/commit/fd4c0b17bcb1c281baac1772ab94e305ec8c5c86))
- Update github templates ([#195](https://github.com/ory/keto/issues/195))
  ([ba0943c](https://github.com/ory/keto/commit/ba0943c45d36ef10bdf1169f0aeef439a3a67d28))
- Update linux install guide ([#191](https://github.com/ory/keto/issues/191))
  ([7d8b24b](https://github.com/ory/keto/commit/7d8b24bddb9c92feb78c7b65f39434d538773b58))
- Update repository templates
  ([ea65b5c](https://github.com/ory/keto/commit/ea65b5c5ada0a7453326fa755aa914306f1b1851))
- Use central banner repo for README
  ([0d95d97](https://github.com/ory/keto/commit/0d95d97504df4d0ab57d18dc6d0a824a3f8f5896))
- Use correct banner
  ([c6dfe28](https://github.com/ory/keto/commit/c6dfe280fd962169c424834cea040a408c1bc83f))
- Use correct version
  ([5f7030c](https://github.com/ory/keto/commit/5f7030c9069fe392200be72f8ce1a93890fbbba8)),
  closes [#200](https://github.com/ory/keto/issues/200)
- Use correct versions in install docs
  ([52e6c34](https://github.com/ory/keto/commit/52e6c34780ed41c169504d71c39459898b5d14f9))

# [0.5.4-alpha.1](https://github.com/ory/keto/compare/v0.5.3-alpha.3...v0.5.4-alpha.1) (2020-04-07)

fix: resolve panic when executing migrations (#178)

Closes #177

### Bug Fixes

- Resolve panic when executing migrations
  ([#178](https://github.com/ory/keto/issues/178))
  ([7e83fee](https://github.com/ory/keto/commit/7e83feefaad041c60f09232ac44ed8b7240c6558)),
  closes [#177](https://github.com/ory/keto/issues/177)

# [0.5.3-alpha.3](https://github.com/ory/keto/compare/v0.5.3-alpha.2...v0.5.3-alpha.3) (2020-04-06)

autogen(docs): regenerate and update changelog

### Code Generation

- **docs:** Regenerate and update changelog
  ([769cef9](https://github.com/ory/keto/commit/769cef90f27ba9c203d3faf47272287ab17dc7eb))

### Code Refactoring

- Move docs to this repository ([#172](https://github.com/ory/keto/issues/172))
  ([312480d](https://github.com/ory/keto/commit/312480de3cefc5b72916ba95d8287443cf3ccb3d))

### Documentation

- Regenerate and update changelog
  ([dda79b1](https://github.com/ory/keto/commit/dda79b106a18bc33d70ae60e352118b0d288d26b))
- Regenerate and update changelog
  ([9048dd8](https://github.com/ory/keto/commit/9048dd8d8a0f0654072b3d4b77261fe947a34ece))
- Regenerate and update changelog
  ([806f68c](https://github.com/ory/keto/commit/806f68c603781742e0177ec0b2deecaf64c5b721))
- Regenerate and update changelog
  ([8905ee7](https://github.com/ory/keto/commit/8905ee74d4ec394af92240e180cc5d7f6493cb2f))
- Regenerate and update changelog
  ([203c1cc](https://github.com/ory/keto/commit/203c1cc659a72f81a370d7b9b7fbda60e7c96c9e))
- Regenerate and update changelog
  ([8875a95](https://github.com/ory/keto/commit/8875a95b35df57668acb27820a3aff1cdfbe8b30))
- Regenerate and update changelog
  ([28ddd3e](https://github.com/ory/keto/commit/28ddd3e1483afe8571b3d2bf9afcc31386d85f7f))
- Regenerate and update changelog
  ([927c4ed](https://github.com/ory/keto/commit/927c4edc4a770133bcb34bc044dd5c5e0eb3ffb7))
- Updates issue and pull request templates
  ([#168](https://github.com/ory/keto/issues/168))
  ([29a38a8](https://github.com/ory/keto/commit/29a38a85c61ec2c8d0ad2ce6d5a0f9e9d74b52f7))
- Updates issue and pull request templates
  ([#169](https://github.com/ory/keto/issues/169))
  ([99b7d5d](https://github.com/ory/keto/commit/99b7d5de24fed1aed746c4447a390d084632f89a))
- Updates issue and pull request templates
  ([#171](https://github.com/ory/keto/issues/171))
  ([7a9876b](https://github.com/ory/keto/commit/7a9876b8ed4282f50f886a025033641bd027a0e2))

# [0.5.3-alpha.1](https://github.com/ory/keto/compare/v0.5.2...v0.5.3-alpha.1) (2020-04-03)

chore: move to ory analytics fork (#167)

### Chores

- Move to ory analytics fork ([#167](https://github.com/ory/keto/issues/167))
  ([f824011](https://github.com/ory/keto/commit/f824011b4d19058504b3a43ed53a420619444a51))

# [0.5.2](https://github.com/ory/keto/compare/v0.5.1-alpha.1...v0.5.2) (2020-04-02)

docs: Regenerate and update changelog

### Documentation

- Regenerate and update changelog
  ([1e52100](https://github.com/ory/keto/commit/1e521001a43a0a13e2224e1a44956442ac6ffbc7))
- Regenerate and update changelog
  ([e4d32a6](https://github.com/ory/keto/commit/e4d32a62c1ae96115ea50bb471f5ff2ce2f2c4b9))

# [0.5.0](https://github.com/ory/keto/compare/v0.4.5-alpha.1...v0.5.0) (2020-04-02)

docs: use real json bool type in swagger (#162)

Closes #160

### Bug Fixes

- Move to ory sqa service ([#159](https://github.com/ory/keto/issues/159))
  ([c3bf1b1](https://github.com/ory/keto/commit/c3bf1b1964a14be4cc296aae98d0739e65917e18))
- Use correct response mode for removeOryAccessControlPolicyRoleMe…
  ([#161](https://github.com/ory/keto/issues/161))
  ([17543cf](https://github.com/ory/keto/commit/17543cfef63a1d040a2234bd63b210fb9c4f6015))

### Documentation

- Regenerate and update changelog
  ([6a77f75](https://github.com/ory/keto/commit/6a77f75d66e89420f2daf2fae011d31bcfa34008))
- Regenerate and update changelog
  ([c8c9d29](https://github.com/ory/keto/commit/c8c9d29e77ef53e1196cc6fe600c53d93376229b))
- Regenerate and update changelog
  ([fe8327d](https://github.com/ory/keto/commit/fe8327d951394084df7785166c9a9578c1ab0643))
- Regenerate and update changelog
  ([b5b1d66](https://github.com/ory/keto/commit/b5b1d66a4b933df8789337cce3f6d6bf391b617b))
- Update forum and chat links
  ([e96d7ba](https://github.com/ory/keto/commit/e96d7ba3dcc693c22eb983b3f58a05c9c6adbda7))
- Updates issue and pull request templates
  ([#158](https://github.com/ory/keto/issues/158))
  ([ab14cfa](https://github.com/ory/keto/commit/ab14cfa51ce195b26a83c050452530a5008589d7))
- Use real json bool type in swagger
  ([#162](https://github.com/ory/keto/issues/162))
  ([5349e7f](https://github.com/ory/keto/commit/5349e7f910ad22558a01b76be62db2136b5eb301)),
  closes [#160](https://github.com/ory/keto/issues/160)

# [0.4.5-alpha.1](https://github.com/ory/keto/compare/v0.4.4-alpha.1...v0.4.5-alpha.1) (2020-02-29)

docs: Regenerate and update changelog

### Bug Fixes

- **driver:** Extract scheme from DSN using sqlcon.GetDriverName
  ([#156](https://github.com/ory/keto/issues/156))
  ([187e289](https://github.com/ory/keto/commit/187e289f1a235b5cacf2a0b7ca5e98c384fa7a14)),
  closes [#145](https://github.com/ory/keto/issues/145)

### Documentation

- Regenerate and update changelog
  ([41513da](https://github.com/ory/keto/commit/41513da35ea038f3c4cc2d98b9796cee5b5a8b92))

# [0.4.4-alpha.1](https://github.com/ory/keto/compare/v0.4.3-alpha.2...v0.4.4-alpha.1) (2020-02-14)

docs: Regenerate and update changelog

### Bug Fixes

- **goreleaser:** Update brew section
  ([0918ff3](https://github.com/ory/keto/commit/0918ff3032eeecd26c67d6249c7e28e71ee110af))

### Documentation

- Prepare ecosystem automation
  ([2e39be7](https://github.com/ory/keto/commit/2e39be79ebad1cec021ae3ee4b0a75ffea4b7424))
- Regenerate and update changelog
  ([009c4c4](https://github.com/ory/keto/commit/009c4c4e4fd4c5607cc30cc9622fd0f82e3891f3))
- Regenerate and update changelog
  ([49f3c4b](https://github.com/ory/keto/commit/49f3c4ba34df5879d8f48cc96bf0df9dad820362))
- Updates issue and pull request templates
  ([#153](https://github.com/ory/keto/issues/153))
  ([7fb7521](https://github.com/ory/keto/commit/7fb752114e1e2a91ab96fdb546835de8aee4926b))

### Features

- **ci:** Add nancy vuln scanner
  ([#152](https://github.com/ory/keto/issues/152))
  ([c19c2b9](https://github.com/ory/keto/commit/c19c2b9efe8299b8878cc8099fe314d8dcda3a08))

### Unclassified

- Update CHANGELOG [ci skip]
  ([63fe513](https://github.com/ory/keto/commit/63fe513d22ec3747a95cdb8f797ba1eba5ca344f))
- Update CHANGELOG [ci skip]
  ([7b7c3ac](https://github.com/ory/keto/commit/7b7c3ac6c06c072fea1b64624ea79a3fd406b09c))
- Update CHANGELOG [ci skip]
  ([8886392](https://github.com/ory/keto/commit/8886392b39fb46ad338c8284866d4dae64ad1826))
- Update CHANGELOG [ci skip]
  ([5bbc284](https://github.com/ory/keto/commit/5bbc2844c49b0a68ba3bd8b003d91f87e2aed9e2))

# [0.4.3-alpha.2](https://github.com/ory/keto/compare/v0.4.3-alpha.1...v0.4.3-alpha.2) (2020-01-31)

Update README.md

### Unclassified

- Update README.md
  ([0ab9c6f](https://github.com/ory/keto/commit/0ab9c6f372a1538a958a68b34315c9167b5a9093))
- Update CHANGELOG [ci skip]
  ([f0a1428](https://github.com/ory/keto/commit/f0a1428f4b99ceb35ff4f1e839bc5237e19db628))

# [0.4.3-alpha.1](https://github.com/ory/keto/compare/v0.4.2-alpha.1...v0.4.3-alpha.1) (2020-01-23)

Disable access logging for health endpoints (#151)

Closes #150

### Unclassified

- Disable access logging for health endpoints (#151)
  ([6ca0c09](https://github.com/ory/keto/commit/6ca0c09b5618122762475cffdc9c32adf28456a1)),
  closes [#151](https://github.com/ory/keto/issues/151)
  [#150](https://github.com/ory/keto/issues/150)

# [0.4.2-alpha.1](https://github.com/ory/keto/compare/v0.4.1-beta.1...v0.4.2-alpha.1) (2020-01-14)

Update CHANGELOG [ci skip]

### Unclassified

- Update CHANGELOG [ci skip]
  ([afaabde](https://github.com/ory/keto/commit/afaabde63affcf568e3090e55b4b957edff2890c))

# [0.4.1-beta.1](https://github.com/ory/keto/compare/v0.4.0-sandbox...v0.4.1-beta.1) (2020-01-13)

Update CHANGELOG [ci skip]

### Unclassified

- Update CHANGELOG [ci skip]
  ([e3ca5a7](https://github.com/ory/keto/commit/e3ca5a7d8b9827ffc7b31a8b5e459db3e912a590))
- Update SDK
  ([5dd6237](https://github.com/ory/keto/commit/5dd623755d4832f33c3dcefb778a9a70eace7b52))

# [0.4.0-alpha.1](https://github.com/ory/keto/compare/v0.3.9-sandbox...v0.4.0-alpha.1) (2020-01-13)

Move to new SDK generators (#146)

### Unclassified

- Move to new SDK generators (#146)
  ([4f51a09](https://github.com/ory/keto/commit/4f51a0948723efc092f1887b111d1e6dd590a075)),
  closes [#146](https://github.com/ory/keto/issues/146)
- Fix typos in the README (#144)
  ([85d838c](https://github.com/ory/keto/commit/85d838c0872c73eb70b5bfff1ccb175b07f6b1e4)),
  closes [#144](https://github.com/ory/keto/issues/144)

# [0.3.9-sandbox](https://github.com/ory/keto/compare/v0.3.8-sandbox...v0.3.9-sandbox) (2019-12-16)

Update go modules

### Unclassified

- Update go modules
  ([1151e07](https://github.com/ory/keto/commit/1151e0755c974b0aea86be5aaeae365ea9aef094))

# [0.3.7-sandbox](https://github.com/ory/keto/compare/v0.3.6-sandbox...v0.3.7-sandbox) (2019-12-11)

Update documentation banner image (#143)

### Unclassified

- Update documentation banner image (#143)
  ([e444755](https://github.com/ory/keto/commit/e4447552031a4f26ec21a336071b0bb19843df61)),
  closes [#143](https://github.com/ory/keto/issues/143)
- Revert incorrect license changes
  ([094c4f3](https://github.com/ory/keto/commit/094c4f30184d77a05044087c13e71ce4adb4d735))
- Fix invalid pseudo version ([#138](https://github.com/ory/keto/issues/138))
  ([79b4457](https://github.com/ory/keto/commit/79b4457f0162197ba267edbb8c0031c47e03bade))

# [0.3.6-sandbox](https://github.com/ory/keto/compare/v0.3.5-sandbox...v0.3.6-sandbox) (2019-10-16)

Resolve issues with mysql tests (#137)

### Unclassified

- Resolve issues with mysql tests (#137)
  ([ef5aec8](https://github.com/ory/keto/commit/ef5aec8e493199c46b78e8f1257aa41df9545f28)),
  closes [#137](https://github.com/ory/keto/issues/137)

# [0.3.5-sandbox](https://github.com/ory/keto/compare/v0.3.4-sandbox...v0.3.5-sandbox) (2019-08-21)

Implement roles and policies filter (#124)

### Documentation

- Incorporates changes from version v0.3.3-sandbox
  ([57686d2](https://github.com/ory/keto/commit/57686d2e30b229cae33e717eb8b3db9da3bdaf0a))
- README grammar fixes ([#114](https://github.com/ory/keto/issues/114))
  ([e592736](https://github.com/ory/keto/commit/e5927360300d8c4fbea841c1c2fb92b48b77885e))
- Updates issue and pull request templates
  ([#110](https://github.com/ory/keto/issues/110))
  ([80c8516](https://github.com/ory/keto/commit/80c8516efbcf33902d8a45f1dc7dbafff2aab8b1))
- Updates issue and pull request templates
  ([#111](https://github.com/ory/keto/issues/111))
  ([22305d0](https://github.com/ory/keto/commit/22305d0a9b5114de8125c16030bbcd1de695ae9b))
- Updates issue and pull request templates
  ([#112](https://github.com/ory/keto/issues/112))
  ([dccada9](https://github.com/ory/keto/commit/dccada9a2189bbd899c5c4a18665a92113fe6cd7))
- Updates issue and pull request templates
  ([#125](https://github.com/ory/keto/issues/125))
  ([15f373a](https://github.com/ory/keto/commit/15f373a16b8cfbd6cdad2bda5f161e171c566137))
- Updates issue and pull request templates
  ([#128](https://github.com/ory/keto/issues/128))
  ([eaf8e33](https://github.com/ory/keto/commit/eaf8e33f3904484635924bdac190c8dc7b60f939))
- Updates issue and pull request templates
  ([#130](https://github.com/ory/keto/issues/130))
  ([a440d14](https://github.com/ory/keto/commit/a440d142275a7a91a0a6bb487fe47d22247f4988))
- Updates issue and pull request templates
  ([#131](https://github.com/ory/keto/issues/131))
  ([dbf2cb2](https://github.com/ory/keto/commit/dbf2cb23c5b6f0f1ee0be5c0b5a58fb0c3dbefd1))
- Updates issue and pull request templates
  ([#132](https://github.com/ory/keto/issues/132))
  ([e121048](https://github.com/ory/keto/commit/e121048d10627ed32a07e26455efd69248f1bd95))
- Updates issue and pull request templates
  ([#133](https://github.com/ory/keto/issues/133))
  ([1b7490a](https://github.com/ory/keto/commit/1b7490abc1d5d0501b66595eb2d92834b6fb0345))

### Unclassified

- Implement roles and policies filter (#124)
  ([db94481](https://github.com/ory/keto/commit/db9448103621a6a8cd086a4cef6c6a22398e621f)),
  closes [#124](https://github.com/ory/keto/issues/124)
- Add adopters placeholder ([#129](https://github.com/ory/keto/issues/129))
  ([b814838](https://github.com/ory/keto/commit/b8148388b8bea97d1f1b4b54de2f0b8ef6b8b6c7))
- Improve documentation (#126)
  ([aabb04d](https://github.com/ory/keto/commit/aabb04d5f283d3c73eb3f3531b4e470ae716db5e)),
  closes [#126](https://github.com/ory/keto/issues/126)
- Create FUNDING.yml
  ([571b447](https://github.com/ory/keto/commit/571b447ed3a02f43623ef5c5adc09682b5f379bd))
- Use non-root user in image ([#116](https://github.com/ory/keto/issues/116))
  ([a493e55](https://github.com/ory/keto/commit/a493e550a8bb86d99164f4ea76dbcecf76c9c2c1))
- Remove binary license (#117)
  ([6e85f7c](https://github.com/ory/keto/commit/6e85f7c6f430e88fb4117a131f57bd69466a8ca1)),
  closes [#117](https://github.com/ory/keto/issues/117)

# [0.3.3-sandbox](https://github.com/ory/keto/compare/v0.3.1-sandbox...v0.3.3-sandbox) (2019-05-18)

ci: Resolve goreleaser issues (#108)

### Continuous Integration

- Resolve goreleaser issues ([#108](https://github.com/ory/keto/issues/108))
  ([5753f27](https://github.com/ory/keto/commit/5753f27a9e89ccdda7c02969217c253aa72cb94b))

### Documentation

- Incorporates changes from version v0.3.1-sandbox
  ([b8a0029](https://github.com/ory/keto/commit/b8a002937483a0f71fe5aba26bb18beb41886249))
- Updates issue and pull request templates
  ([#106](https://github.com/ory/keto/issues/106))
  ([54a5a27](https://github.com/ory/keto/commit/54a5a27f24a90ab3c5f9915f36582b85eecd0d62))

# [0.3.1-sandbox](https://github.com/ory/keto/compare/v0.3.0-sandbox...v0.3.1-sandbox) (2019-04-29)

ci: Use image that includes bash/sh for release docs (#103)

Signed-off-by: aeneasr <aeneas@ory.sh>

### Continuous Integration

- Use image that includes bash/sh for release docs
  ([#103](https://github.com/ory/keto/issues/103))
  ([e9d3027](https://github.com/ory/keto/commit/e9d3027fc62b20f28cd7a023222390e24d565eb1))

### Documentation

- Incorporates changes from version v0.3.0-sandbox
  ([605d2f4](https://github.com/ory/keto/commit/605d2f43621b806b750edc81d439edc92cfb7c38))

### Unclassified

- Allow configuration files and update UPGRADE guide. (#102)
  ([3934dc6](https://github.com/ory/keto/commit/3934dc6e690822358067b43920048d45a4b7799b)),
  closes [#102](https://github.com/ory/keto/issues/102)

# [0.3.0-sandbox](https://github.com/ory/keto/compare/v0.2.3-sandbox+oryOS.10...v0.3.0-sandbox) (2019-04-29)

docker: Remove full tag from build pipeline (#101)

Signed-off-by: aeneasr <aeneas@ory.sh>

### Documentation

- Update patrons
  ([c8dc7cd](https://github.com/ory/keto/commit/c8dc7cdc68676970328b55648b8d6e469c77fbfd))

### Unclassified

- Improve naming for ory policies
  ([#100](https://github.com/ory/keto/issues/100))
  ([b39703d](https://github.com/ory/keto/commit/b39703d362d333213fcb7d3782e363d09b6dabbd))
- Remove full tag from build pipeline
  ([#101](https://github.com/ory/keto/issues/101))
  ([602a273](https://github.com/ory/keto/commit/602a273dc5a0c29e80a22f04adb937ab385c4512))
- Remove duplicate code in Makefile (#99)
  ([04f5223](https://github.com/ory/keto/commit/04f52231509dd0f3a57d745918fc43fff7c595ff)),
  closes [#99](https://github.com/ory/keto/issues/99)
- Add tracing support and general improvements (#98)
  ([63b3946](https://github.com/ory/keto/commit/63b3946e0ae1fa23c6a359e9a64b296addff868c)),
  closes [#98](https://github.com/ory/keto/issues/98):

  This patch improves the internal configuration and service management. It adds
  support for distributed tracing and resolves several issues in the release
  pipeline and CLI.

  Additionally, composable docker-compose configuration files have been added.

  Several bugs have been fixed in the release management pipeline.

- Add content-type in the response of allowed
  ([#90](https://github.com/ory/keto/issues/90))
  ([39a1486](https://github.com/ory/keto/commit/39a1486dc53456189d30380460a9aeba198fa9e9))
- Fix disable-telemetry check ([#85](https://github.com/ory/keto/issues/85))
  ([38b5383](https://github.com/ory/keto/commit/38b538379973fa34bd2bf24dcb2e6dbedf324e1e))
- Fix remove member from role ([#87](https://github.com/ory/keto/issues/87))
  ([698e161](https://github.com/ory/keto/commit/698e161989331ca5a3a0769301d9694ef805a876)),
  closes [#74](https://github.com/ory/keto/issues/74)
- Fix the type of conditions in the policy
  ([#86](https://github.com/ory/keto/issues/86))
  ([fc1ced6](https://github.com/ory/keto/commit/fc1ced63bd39c9fbf437e419dfc384343e36e0ee))
- Move Go SDK generation to go-swagger
  ([#94](https://github.com/ory/keto/issues/94))
  ([9f48a95](https://github.com/ory/keto/commit/9f48a95187a7b6160108cd7d0301590de2e58f07)),
  closes [#92](https://github.com/ory/keto/issues/92)
- Send 403 when authorization result is negative
  ([#93](https://github.com/ory/keto/issues/93))
  ([de806d8](https://github.com/ory/keto/commit/de806d892819db63c1abc259ab06ee08d87895dc)),
  closes [#75](https://github.com/ory/keto/issues/75)
- Update dependencies ([#91](https://github.com/ory/keto/issues/91))
  ([4d44174](https://github.com/ory/keto/commit/4d4417474ebf8cc69d01e5ac82633b966cdefbc7))
- storage/memory: Fix upsert with pre-existing key will causes duplicate records
  (#88)
  ([1cb8a36](https://github.com/ory/keto/commit/1cb8a36a08883b785d9bb0a4be1ddc00f1f9d358)),
  closes [#88](https://github.com/ory/keto/issues/88)
  [#80](https://github.com/ory/keto/issues/80)

# [0.2.3-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.2.2-sandbox+oryOS.10...v0.2.3-sandbox+oryOS.10) (2019-02-05)

dist: Fix packr build pipeline (#84)

Closes #73 Closes #81

Signed-off-by: aeneasr <aeneas@ory.sh>

### Documentation

- Add documentation for glob matching
  ([5c8babb](https://github.com/ory/keto/commit/5c8babbfbae01a78f30cfbff92d8e9c3a6b09027))
- Incorporates changes from version v0.2.2-sandbox+oryOS.10
  ([ed7af3f](https://github.com/ory/keto/commit/ed7af3fa4e5d1d0d03b5366f4cf865a5b82ec293))
- Properly generate api.swagger.json
  ([18e3f84](https://github.com/ory/keto/commit/18e3f84cdeee317f942d61753399675c98886e5d))

### Unclassified

- Add placeholder go file for rego inclusion
  ([6a6f64d](https://github.com/ory/keto/commit/6a6f64d8c59b496f6cf360f55eba1e16bf5380f1))
- Add support for glob matching
  ([bb76c6b](https://github.com/ory/keto/commit/bb76c6bebe522fc25448c4f4e4d1ef7c530a725f))
- Ex- and import rego subdirectories for `go get`
  [#77](https://github.com/ory/keto/issues/77)
  ([59cc053](https://github.com/ory/keto/commit/59cc05328f068fc3046b2dbc022a562fd5d67960)),
  closes [#73](https://github.com/ory/keto/issues/73)
- Fix packr build pipeline ([#84](https://github.com/ory/keto/issues/84))
  ([65a87d5](https://github.com/ory/keto/commit/65a87d564d237bc979bb5962beff7d3703d9689f)),
  closes [#73](https://github.com/ory/keto/issues/73)
  [#81](https://github.com/ory/keto/issues/81)
- Import glob in rego/doc.go
  ([7798442](https://github.com/ory/keto/commit/7798442553cfe7989a23d2c389c8c63a24013543))
- Properly handle dbal error
  ([6811607](https://github.com/ory/keto/commit/6811607ea79c8f3155a17bc1aea566e9e4680616))
- Properly handle TLS certificates if set
  ([36399f0](https://github.com/ory/keto/commit/36399f09261d4f3cb5e053679eee3cb15da2df19)),
  closes [#73](https://github.com/ory/keto/issues/73)

# [0.2.2-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.2.1-sandbox+oryOS.10...v0.2.2-sandbox+oryOS.10) (2018-12-13)

ci: Fix docker push arguments in publish task

Signed-off-by: aeneasr <aeneas@ory.sh>

### Continuous Integration

- Fix docker push arguments in publish task
  ([f03c77c](https://github.com/ory/keto/commit/f03c77c6b7461ab81cb03265cbec909ac45c2259))

# [0.2.1-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.2.0-sandbox+oryOS.10...v0.2.1-sandbox+oryOS.10) (2018-12-13)

ci: Fix docker release task

Signed-off-by: aeneasr <aeneas@ory.sh>

### Continuous Integration

- Fix docker release task
  ([7a0414f](https://github.com/ory/keto/commit/7a0414f614b6cc8b1d78cfbb773a2f0192d00d23))

# [0.2.0-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.0.1...v0.2.0-sandbox+oryOS.10) (2018-12-13)

all: gofmt

Signed-off-by: aeneasr <aeneas@ory.sh>

### Documentation

- Adds banner
  ([0ec1d8f](https://github.com/ory/keto/commit/0ec1d8f5e843465d17ac4c8f91f18e5badf16900))
- Adds GitHub templates & code of conduct
  ([#31](https://github.com/ory/keto/issues/31))
  ([a11e898](https://github.com/ory/keto/commit/a11e8980f2af528f1357659966123d0cbf7d43db))
- Adds link to examples repository
  ([#32](https://github.com/ory/keto/issues/32))
  ([7061a2a](https://github.com/ory/keto/commit/7061a2aa31652a9e0c2d449facb1201bfa11fd3f))
- Adds security console image
  ([fd27fc9](https://github.com/ory/keto/commit/fd27fc9cce50beb3d0189e0a93300879fd7149db))
- Changes hydra to keto in readme
  ([9dab531](https://github.com/ory/keto/commit/9dab531744cf5b0ae98862945d44b07535595781))
- Deprecate old versions in logs
  ([955d647](https://github.com/ory/keto/commit/955d647307a48ee7cf2d3f9fb4263072adf42299))
- Incorporates changes from version
  ([85c4d81](https://github.com/ory/keto/commit/85c4d81a192e92f874c106b91cfa6fb404d9a34a))
- Incorporates changes from version v0.0.0-testrelease.1
  ([6062dd4](https://github.com/ory/keto/commit/6062dd4a894607f5f1ead119af20cc8bdbe15bef))
- Incorporates changes from version v0.0.1-1-g85c4d81
  ([f4606fc](https://github.com/ory/keto/commit/f4606fce0326bece2a89dadc029bc5ce9778df18))
- Incorporates changes from version v0.0.1-11-g114914f
  ([92a4dca](https://github.com/ory/keto/commit/92a4dca7a41dcf3a88c4800bf6d2217f33cfcdd1))
- Incorporates changes from version v0.0.1-16-g7d8a8ad
  ([2b76a83](https://github.com/ory/keto/commit/2b76a83755153b3f8a2b8d28c5b0029d96d567b6))
- Incorporates changes from version v0.0.1-18-g099e7e0
  ([70b12ad](https://github.com/ory/keto/commit/70b12adf5bcc0e890d6707e11e891e6cedfb3d87))
- Incorporates changes from version v0.0.1-20-g97ccbe6
  ([b21d56e](https://github.com/ory/keto/commit/b21d56e599c7eb4c1769bc18878f7d5818b73023))
- Incorporates changes from version v0.0.1-30-gaf2c3b5
  ([a1d0dcc](https://github.com/ory/keto/commit/a1d0dcc78a9506260f86df00e4dff8ab02909ce1))
- Incorporates changes from version v0.0.1-32-gedb5a60
  ([a5c369a](https://github.com/ory/keto/commit/a5c369a90da67c96bbde60e673c67f50b841fadd))
- Incorporates changes from version v0.0.1-6-g570783e
  ([0fcbbcb](https://github.com/ory/keto/commit/0fcbbcb02f1d748f9c733c86368b223b2ee4c6e2))
- Incorporates changes from version v0.0.1-7-g0fcbbcb
  ([c0141a8](https://github.com/ory/keto/commit/c0141a8ec22ea1260bf2d45d72dfe06737ec0115))
- Incorporates changes from version v0.1.0-sandbox
  ([9ee0664](https://github.com/ory/keto/commit/9ee06646d2cfb2d69abdcc411e31d14957437a1e))
- Incorporates changes from version v1.0.0-beta.1-1-g162d7b8
  ([647c5a9](https://github.com/ory/keto/commit/647c5a9e1bc8d9d635bf6f2511c3faa9a9daefef))
- Incorporates changes from version v1.0.0-beta.2-11-g2b280bb
  ([936889d](https://github.com/ory/keto/commit/936889d760f04a03d498f65331d653cbad3702d0))
- Incorporates changes from version v1.0.0-beta.2-13-g382e1d3
  ([883df44](https://github.com/ory/keto/commit/883df44a922f3daee86597af467072555cadc7e7))
- Incorporates changes from version v1.0.0-beta.2-15-g74450da
  ([48dd9f1](https://github.com/ory/keto/commit/48dd9f1ffbeaa99ac8dc27085c5a50f9244bf9c3))
- Incorporates changes from version v1.0.0-beta.2-3-gf623c52
  ([b6b90e5](https://github.com/ory/keto/commit/b6b90e5b2180921f78064a60666704b4e72679b6))
- Incorporates changes from version v1.0.0-beta.2-5-g3852be5
  ([3f09090](https://github.com/ory/keto/commit/3f09090a2f82f3f29154c19217cea0a10d65ea3a))
- Incorporates changes from version v1.0.0-beta.2-9-gc785187
  ([4c30a3c](https://github.com/ory/keto/commit/4c30a3c0ad83ba80e1857b41211e7ddade06c4cf))
- Incorporates changes from version v1.0.0-beta.3-1-g06adbf1
  ([0ba3c06](https://github.com/ory/keto/commit/0ba3c0674832b641ef5e0c3f0d60d81ed3a647b2))
- Incorporates changes from version v1.0.0-beta.3-10-g9994967
  ([d2345ca](https://github.com/ory/keto/commit/d2345ca3beb354d6ee7c7926c1a5ddb425d6b405))
- Incorporates changes from version v1.0.0-beta.3-12-gc28b521
  ([b4d792f](https://github.com/ory/keto/commit/b4d792f74055853f05ca46c67625ffd432fc74fd))
- Incorporates changes from version v1.0.0-beta.3-3-g9e16605
  ([c43bf2b](https://github.com/ory/keto/commit/c43bf2b5232bed9106dd47d7eb53d2f93bfe260d))
- Incorporates changes from version v1.0.0-beta.3-5-ga11e898
  ([b9d9b8e](https://github.com/ory/keto/commit/b9d9b8ee33ab957f43f99c427a88ade847e79ed0))
- Incorporates changes from version v1.0.0-beta.3-8-g7061a2a
  ([d76ff9d](https://github.com/ory/keto/commit/d76ff9dc9a4c8a8f1286eeb139d8f5af9617f421))
- Incorporates changes from version v1.0.0-beta.5
  ([0dc314c](https://github.com/ory/keto/commit/0dc314c7888020b40e12eb59fd77135044fd063b))
- Incorporates changes from version v1.0.0-beta.6-1-g5e97104
  ([f14c8ed](https://github.com/ory/keto/commit/f14c8edd7204a811e333ea84429cf837b4e7d27b))
- Incorporates changes from version v1.0.0-beta.8
  ([5045b59](https://github.com/ory/keto/commit/5045b59e2a83d6ab047b1b95c581d7c34e96a2e0))
- Incorporates changes from version v1.0.0-beta.9
  ([be2f035](https://github.com/ory/keto/commit/be2f03524721ef47ecb1c9aec57c2696174e0657))
- Properly sets up changelog TOC
  ([e0acd67](https://github.com/ory/keto/commit/e0acd670ab19c0d6fd36733fea164e2b0414597d))
- Puts toc in the right place
  ([114914f](https://github.com/ory/keto/commit/114914fa354f784b310bc9dfd232a011e0d98d99))
- Revert changes from test release
  ([ab3a64d](https://github.com/ory/keto/commit/ab3a64d3d41292364c5947db98c4d27a8223853e))
- Update documentation links ([#67](https://github.com/ory/keto/issues/67))
  ([d22d413](https://github.com/ory/keto/commit/d22d413c7a001ccaa96b4c013665153f41831614))
- Update link to security console
  ([846ce4b](https://github.com/ory/keto/commit/846ce4baa9da5954bd30996f489885a026c48185))
- Update migration guide
  ([3c44b58](https://github.com/ory/keto/commit/3c44b58613e46ed39d42463537773fe9d95a54da))
- Update to latest changes
  ([1625123](https://github.com/ory/keto/commit/1625123ed342f019d5e7ab440eb37da310570842))
- Updates copyright notice
  ([9dd5578](https://github.com/ory/keto/commit/9dd557825dfd3b9d589c9db2ccb201638debbaae))
- Updates installation guide
  ([f859645](https://github.com/ory/keto/commit/f859645f230f405cfabed0c1b9a2b67b1a3841d3))
- Updates issue and pull request templates
  ([#52](https://github.com/ory/keto/issues/52))
  ([941cae6](https://github.com/ory/keto/commit/941cae6fee058f68eabbbf4dd9cafad4760e108f))
- Updates issue and pull request templates
  ([#53](https://github.com/ory/keto/issues/53))
  ([7b222d2](https://github.com/ory/keto/commit/7b222d285e74c0db482136b23f37072216b3acb0))
- Updates issue and pull request templates
  ([#54](https://github.com/ory/keto/issues/54))
  ([f098639](https://github.com/ory/keto/commit/f098639b5e748151810848fdd3173e0246bc03dc))
- Updates link to guide and header
  ([437c255](https://github.com/ory/keto/commit/437c255ecfff4127fb586cc069e07f86988ad1ba))
- Updates link to open collective
  ([382e1d3](https://github.com/ory/keto/commit/382e1d34c7da0ba0447b78506a749bd7f0085f48))
- Updates links to docs
  ([d84be3b](https://github.com/ory/keto/commit/d84be3b6a8e5eb284ec3fb137ee774ba5ee0d529))
- Updates newsletter link in README
  ([2dc36b2](https://github.com/ory/keto/commit/2dc36b21c8af8e3e39f093198715ea24b65d65af))

### Unclassified

- Add Go SDK factory
  ([99db7e6](https://github.com/ory/keto/commit/99db7e6d4edac88794266a01ddfab9cd0632e95a))
- Add go SDK interface
  ([3dd5f7d](https://github.com/ory/keto/commit/3dd5f7d61bb460c34744b84a34755bfb8219b304))
- Add health handlers
  ([bddb949](https://github.com/ory/keto/commit/bddb949459d05002b0f8882d981e4f63fdddf25f))
- Add policy list handler
  ([a290619](https://github.com/ory/keto/commit/a290619d01d15eb8e3b4e33ede1058d316ee807a))
- Add role iterator in list handler
  ([a3eb696](https://github.com/ory/keto/commit/a3eb6961783f7b562f0a0d0a7e2819bffebce5b8))
- Add SDK generation to circle ci
  ([9b37165](https://github.com/ory/keto/commit/9b37165873bcb0cc5dc60d2514d9824a073466a1))
- Adds ability to update a role using PUT
  ([#14](https://github.com/ory/keto/issues/14))
  ([97ccbe6](https://github.com/ory/keto/commit/97ccbe6d808823c56901ad237878aa6d53cddeeb)):

  - transfer UpdateRoleMembers from https://github.com/ory/hydra/pull/768 to
    keto

  - fix tests by using right http method & correcting sql request

  - Change behavior to overwrite the whole role instead of just the members.

  * small sql migration fix

- Adds log message when telemetry is active
  ([f623c52](https://github.com/ory/keto/commit/f623c52655ff85b7f7209eb73e94eb66a297c5b7))
- Clean up vendor dependencies
  ([9a33c23](https://github.com/ory/keto/commit/9a33c23f4d37ab88b4d643fd79204334d73404c6))
- Do not split empty scope ([#45](https://github.com/ory/keto/issues/45))
  ([b29cf8c](https://github.com/ory/keto/commit/b29cf8cc92607e13457dba8331f5c9286054c8c1))
- Fix typo in help command in env var name
  ([#39](https://github.com/ory/keto/issues/39))
  ([8a5016c](https://github.com/ory/keto/commit/8a5016cd75be78bb42a9a38bfd453ad5722db9db)),
  closes [#25](https://github.com/ory/keto/issues/25)
- Fixes environment variable typos
  ([566d588](https://github.com/ory/keto/commit/566d588e4fca12399966718b725fe4461a28e51e))
- Fixes typo in help command
  ([74450da](https://github.com/ory/keto/commit/74450da18a27513820328c28f72203653c664367)),
  closes [#25](https://github.com/ory/keto/issues/25)
- Format code
  ([637c78c](https://github.com/ory/keto/commit/637c78cba697682b544473a3af9b6ae7715561aa))
- Gofmt
  ([a8d7f9f](https://github.com/ory/keto/commit/a8d7f9f546ae2f3b8c3fa643d8e19b68ca26cc67))
- Improve compose documentation
  ([6870443](https://github.com/ory/keto/commit/68704435f3c299b853f4ff5cacae285b09ada3b5))
- Improves usage of metrics middleware
  ([726c4be](https://github.com/ory/keto/commit/726c4bedfc3f02fdac380930e32f37c251e51aa4))
- Improves usage of metrics middleware
  ([301f386](https://github.com/ory/keto/commit/301f38605af66abae4d28ed0cac90d0b82b655c4))
- Introduce docker-compose file for testing
  ([ba857e3](https://github.com/ory/keto/commit/ba857e3859966e857c5a741825411575e17446de))
- Introduces health and version endpoints
  ([6a9da74](https://github.com/ory/keto/commit/6a9da74f693ee6c15a775ab8d652582aea093601))
- List roles from keto_role table ([#28](https://github.com/ory/keto/issues/28))
  ([9e16605](https://github.com/ory/keto/commit/9e166054b8d474fbce6983d5d00eeeb062fc79b1))
- Properly names flags
  ([af2c3b5](https://github.com/ory/keto/commit/af2c3b5bc96e95fb31b1db5c7fe6dfd6b6fc5b20))
- Require explicit CORS enabling ([#42](https://github.com/ory/keto/issues/42))
  ([9a45107](https://github.com/ory/keto/commit/9a45107af304b2a8e663a532e4f6e4536f15888c))
- Update dependencies
  ([663d8b1](https://github.com/ory/keto/commit/663d8b13e99694a57752cd60a68342b81b041c66))
- Switch to rego as policy decision engine (#48)
  ([ee9bcf2](https://github.com/ory/keto/commit/ee9bcf2719178e5a8dccca083a90313947a8a63b)),
  closes [#48](https://github.com/ory/keto/issues/48)
- Update hydra to v1.0.0-beta.6 ([#35](https://github.com/ory/keto/issues/35))
  ([5e97104](https://github.com/ory/keto/commit/5e971042afff06e2a6ee3b54d2fea31687203623))
- Update npm package registry
  ([a53d3d2](https://github.com/ory/keto/commit/a53d3d23e11fde5dcfbb27a2add1049f4d8e10e6))
- Enable TLS option to serve API (#46)
  ([2f62063](https://github.com/ory/keto/commit/2f620632d0375bf9c7e58dbfb49627c02c66abf3)),
  closes [#46](https://github.com/ory/keto/issues/46)
- Make introspection authorization optional
  ([e5460ad](https://github.com/ory/keto/commit/e5460ad884cd018cd6177324b949cd66bfd53bc7))
- Properly output telemetry information
  ([#33](https://github.com/ory/keto/issues/33))
  ([9994967](https://github.com/ory/keto/commit/9994967b0ca54a62b8b0088fe02be9e890d9574b))
- Remove ORY Hydra dependency ([#44](https://github.com/ory/keto/issues/44))
  ([d487344](https://github.com/ory/keto/commit/d487344fe7e07cb6370371c6b0b6cf3cca767ed1))
- Resolves an issue with the hydra migrate command
  ([2b280bb](https://github.com/ory/keto/commit/2b280bb57c9073a9c8384cde0b14a6991cfacdb6)),
  closes [#23](https://github.com/ory/keto/issues/23)
- Upgrade superagent version ([#41](https://github.com/ory/keto/issues/41))
  ([9c80dbc](https://github.com/ory/keto/commit/9c80dbcc1cc63243839b58ca56ac9be104797887))
- gofmt
  ([777b1be](https://github.com/ory/keto/commit/777b1be1378d314e7cfde0c34450afcce7e590a5))
- Updates README.md (#34)
  ([c28b521](https://github.com/ory/keto/commit/c28b5219fd64314a75ee3c848a80a0c5974ebb7d)),
  closes [#34](https://github.com/ory/keto/issues/34)
- Properly parses cors options
  ([edb5a60](https://github.com/ory/keto/commit/edb5a600f2ce16c0847ee5ef399fa5a41b1e736a))
- Removes additional output if no args are passed
  ([703e124](https://github.com/ory/keto/commit/703e1246ce0fd89066b497c45f0c6cadeb06c331))
- Resolves broken role test
  ([b6c7f9c](https://github.com/ory/keto/commit/b6c7f9c33c4c1f43164d6da0ec7f2553f1f4c598))
- Resolves minor typos and updates install guide
  ([3852be5](https://github.com/ory/keto/commit/3852be56cb81df966a85d4c828de0397d9e74768))
- Updates to latest sqlcon
  ([2c9f643](https://github.com/ory/keto/commit/2c9f643042ff4edffae8bd41834d2a57c923871c))
- Use roles in warden decision
  ([c785187](https://github.com/ory/keto/commit/c785187e31fc7a4b8b762a5e27fac66dcaa97513)),
  closes [#21](https://github.com/ory/keto/issues/21)
  [#19](https://github.com/ory/keto/issues/19)
- authn/client: Payload is now prefixed with client
  ([8584d94](https://github.com/ory/keto/commit/8584d94cfb18deb37ae32ae601f4cd15c14067e7))

# [0.0.1](https://github.com/ory/keto/compare/4f00bc96ece3180a888718ec3c41c69106c86f56...v0.0.1) (2018-05-20)

authn: Checks token_type is "access_token", if set

Closes #1

### Documentation

- Incorporates changes from version
  ([b5445a0](https://github.com/ory/keto/commit/b5445a0fc5b6f813cd1731b20c8c5c79d7c4cdf8))
- Incorporates changes from version
  ([295ff99](https://github.com/ory/keto/commit/295ff998af55777823b04f423e365fd58e61753b))
- Incorporates changes from version
  ([bd44d41](https://github.com/ory/keto/commit/bd44d41b2781e33353082397c47390a27f749e16))
- Updates readme and upgrades
  ([0f95dbb](https://github.com/ory/keto/commit/0f95dbb967fd17b607caa999ae30453f5f599739))
- Uses keto repo for changelog
  ([14c0b2a](https://github.com/ory/keto/commit/14c0b2a2bd31566f2b9048831f894aba05c5b15d))

### Unclassified

- Adds migrate commands to the proper parent command
  ([231c70d](https://github.com/ory/keto/commit/231c70d816b0736a51eddc1fa0445bac672b1b2f))
- Checks token_type is "access_token", if set
  ([d2b8f5d](https://github.com/ory/keto/commit/d2b8f5d313cce597566bd18e4f3bea4a423a62ee)),
  closes [#1](https://github.com/ory/keto/issues/1)
- Removes old test
  ([07b733b](https://github.com/ory/keto/commit/07b733bfae4b733e3e2124545b92c537dabbdcf0))
- Renames subject to sub in response payloads
  ([ca4d540](https://github.com/ory/keto/commit/ca4d5408000be2b896d38eaaf5e67a3fc0a566da))
- Tells linguist to ignore SDK files
  ([f201eb9](https://github.com/ory/keto/commit/f201eb95f3309a60ac50f42cfba0bae2e38e8d13))
- Retries SQL connection on migrate commands
  ([3d33d73](https://github.com/ory/keto/commit/3d33d73c009077c5bf30ae4b03802904bfb5d5b2)):

  This patch also introduces a fatal error if migrations fail

- cmd/server: Resolves DBAL not handling postgres properly
  ([dedc32a](https://github.com/ory/keto/commit/dedc32ab218923243b1955ce5bcbbdc5cc416953))
- cmd/server: Improves error message in migrate command
  ([4b17ce8](https://github.com/ory/keto/commit/4b17ce8848113cae807840182d1a318190c2a9b3))
- Resolves travis and docker issues
  ([6f4779c](https://github.com/ory/keto/commit/6f4779cc51bf4f2ee5b97541fb77d8f882497710))
- Adds OAuth2 Client Credentials authenticator and warden endpoint
  ([c55139b](https://github.com/ory/keto/commit/c55139b51e636834759706499a2aec1451f4fbd9))
- Adds SDK helpers
  ([a1c2608](https://github.com/ory/keto/commit/a1c260801d9366fccf4bfb4fc64b2c67fc594565))
- Resolves SDK and test issues (#4)
  ([2d4cd98](https://github.com/ory/keto/commit/2d4cd9805af3081bbcbea3f806ca066d35385a4b)),
  closes [#4](https://github.com/ory/keto/issues/4)
- Initial project commit
  ([a592e51](https://github.com/ory/keto/commit/a592e5126f130f8b673fff6c894fdbd9fb56f81c))
- Initial commit
  ([4f00bc9](https://github.com/ory/keto/commit/4f00bc96ece3180a888718ec3c41c69106c86f56))

---

id: changelog title: Changelog custom_edit_url: null

---

# [Unreleased](https://github.com/ory/keto/compare/v0.6.0-alpha.3...8e301198298858fd7f387ef63a7abf4fa55ea240) (2021-06-24)

### Bug Fixes

- Add missing tracers ([#600](https://github.com/ory/keto/issues/600))
  ([aa263be](https://github.com/ory/keto/commit/aa263be9a7830e3c769d7698d36137555ca230bc)),
  closes [#593](https://github.com/ory/keto/issues/593)
- Handle relation tuple cycles in expand and check engine
  ([#623](https://github.com/ory/keto/issues/623))
  ([8e30119](https://github.com/ory/keto/commit/8e301198298858fd7f387ef63a7abf4fa55ea240))
- Log all database connection errors
  ([#588](https://github.com/ory/keto/issues/588))
  ([2b0fad8](https://github.com/ory/keto/commit/2b0fad897e61400bd2a6cdf47f33ff4301e9c5f8))
- Move gRPC client module root up
  ([#620](https://github.com/ory/keto/issues/620))
  ([3b881f6](https://github.com/ory/keto/commit/3b881f6015a93b382b3fbbca4be9259622038b6a)):

  BREAKING: The npm package `@ory/keto-grpc-client` from now on includes all API
  versions. Because of that, the import paths changed. For migrating to the new
  client package, change the import path according to the following example:

  ```diff
  - import acl from '@ory/keto-grpc-client/acl_pb.js'
  + // from the latest version
  + import { acl } from '@ory/keto-grpc-client'
  + // or a specific one
  + import acl from '@ory/keto-grpc-client/ory/keto/acl/v1alpha1/acl_pb.js'
  ```

- Update docker-compose.yml version
  ([#595](https://github.com/ory/keto/issues/595))
  ([7fa4dca](https://github.com/ory/keto/commit/7fa4dca4182a1fa024f9cef0a04163f2cbd882aa)),
  closes [#549](https://github.com/ory/keto/issues/549)

### Documentation

- Fix example not following best practice
  ([#582](https://github.com/ory/keto/issues/582))
  ([a015818](https://github.com/ory/keto/commit/a0158182c5f87cfd4767824e1c5d6cbb8094a4e6))
- Update NPM links due to organisation move
  ([#616](https://github.com/ory/keto/issues/616))
  ([6355bea](https://github.com/ory/keto/commit/6355beae5b5b28c3eee19fdee85b9875cbc165c3))

### Features

- Make generated gRPC client its own module
  ([#583](https://github.com/ory/keto/issues/583))
  ([f0fbb64](https://github.com/ory/keto/commit/f0fbb64b3358e9800854295cebc9ec8b8e56c87a))
- Max_idle_conn_time ([#605](https://github.com/ory/keto/issues/605))
  ([50a8623](https://github.com/ory/keto/commit/50a862338e17f86900ca162da7f3467f55f9f954)),
  closes [#523](https://github.com/ory/keto/issues/523)

# [0.6.0-alpha.3](https://github.com/ory/keto/compare/v0.6.0-alpha.2...v0.6.0-alpha.3) (2021-04-29)

Resolves CRDB and build issues.

### Code Generation

- Pin v0.6.0-alpha.3 release commit
  ([d766968](https://github.com/ory/keto/commit/d766968419d10a68fd843df45316e3436b68d61d))

# [0.6.0-alpha.2](https://github.com/ory/keto/compare/v0.6.0-alpha.1...v0.6.0-alpha.2) (2021-04-29)

This release improves stability and documentation.

### Bug Fixes

- Add npm run format to make format
  ([7d844a8](https://github.com/ory/keto/commit/7d844a8e6412ae561963b97ac26d4682411095d4))
- Makefile target
  ([0e6f612](https://github.com/ory/keto/commit/0e6f6122de7bdbb691ad7cc236b6bc9a3601d39e))
- Move swagger to spec dir
  ([7f6a061](https://github.com/ory/keto/commit/7f6a061aafda275d278bf60f16e90039da45bc57))
- Resolve clidoc issues
  ([ef12b4e](https://github.com/ory/keto/commit/ef12b4e267f34fbf9709fe26023f9b7ae6670c24))
- Update install.sh ([#568](https://github.com/ory/keto/issues/568))
  ([86ab245](https://github.com/ory/keto/commit/86ab24531d608df0b5391ee8ec739291b9a90e20))
- Use correct id
  ([5e02902](https://github.com/ory/keto/commit/5e029020b5ba3931f15d343cf6a9762b064ffd45))
- Use correct id for api
  ([32a6b04](https://github.com/ory/keto/commit/32a6b04609054cba84f7b56ebbe92341ec5dcd98))
- Use sqlite image versions ([#544](https://github.com/ory/keto/issues/544))
  ([ec6cc5e](https://github.com/ory/keto/commit/ec6cc5ed528f1a097ea02669d059e060b7eff824))

### Code Generation

- Pin v0.6.0-alpha.2 release commit
  ([470b2c6](https://github.com/ory/keto/commit/470b2c61c649fe5fcf638c84d4418212ff0330a5))

### Documentation

- Add gRPC client README.md ([#559](https://github.com/ory/keto/issues/559))
  ([9dc3596](https://github.com/ory/keto/commit/9dc35969ada8b0d4d73dee9089c4dc61cd9ea657))
- Change forum to discussions readme
  ([#539](https://github.com/ory/keto/issues/539))
  ([ea2999d](https://github.com/ory/keto/commit/ea2999d4963316810a8d8634fcd123bda31eaa8f))
- Fix cat videos example docker compose
  ([#549](https://github.com/ory/keto/issues/549))
  ([b25a711](https://github.com/ory/keto/commit/b25a7114631957935c71ac6a020ab6bd0c244cd7))
- Fix typo ([#538](https://github.com/ory/keto/issues/538))
  ([99a9693](https://github.com/ory/keto/commit/99a969373497792bb4cd8ff62bf5245087517737))
- Include namespace in olymp library example
  ([#540](https://github.com/ory/keto/issues/540))
  ([135e814](https://github.com/ory/keto/commit/135e8145c383a76b494b469253c949c38f4414a7))
- Update install from source steps to actually work
  ([#548](https://github.com/ory/keto/issues/548))
  ([e662256](https://github.com/ory/keto/commit/e6622564f58b7612b13b11b54e75a7350f52d6de))

### Features

- Global docs sidebar and added cloud pages
  ([c631c82](https://github.com/ory/keto/commit/c631c82b7ff3d12734869ac22730b52e73dcf287))
- Support retryable CRDB transactions
  ([833147d](https://github.com/ory/keto/commit/833147dae40e9ac5bdf220f8aa3f01abd444f791))

# [0.6.0-alpha.1](https://github.com/ory/keto/compare/v0.5.6-alpha.1...v0.6.0-alpha.1) (2021-04-07)

We are extremely happy to announce next-gen Ory Keto which implements
[Zanzibar: Google’s Consistent, Global Authorization System](https://research.google/pubs/pub48190/):

> Zanzibar provides a uniform data model and configuration language for
> expressing a wide range of access control policies from hundreds of client
> services at Google, including Calendar, Cloud, Drive, Maps, Photos, and
> YouTube. Its authorization decisions respect causal ordering of user actions
> and thus provide external consistency amid changes to access control lists and
> object contents. Zanzibar scales to trillions of access control lists and
> millions of authorization requests per second to support services used by
> billions of people. It has maintained 95th-percentile latency of less than 10
> milliseconds and availability of greater than 99.999% over 3 years of
> production use.

Ory Keto is the first open source planet-scale authorization system built with
cloud native technologies (Go, gRPC, newSQL) and architecture. It is also the
first open source implementation of Google Zanzibar :tada:!

Many concepts developer by Google Zanzibar are implemented in Ory Keto already.
Let's take a look!

As of this release, Ory Keto knows how to interpret and operate on the basic
access control lists known as relation tuples. They encode relations between
objects and subjects. One simple example of such a relation tuple could encode
"`user1` has access to file `/foo`", a more complex one could encode "everyone
who has write access on `/foo` has read access on `/foo`".

Ory Keto comes with all the basic APIs as described in the Zanzibar paper. All
of them are available over gRPC and REST.

1. List: query relation tuples
2. Check: determine whether a subject has a relation on an object
3. Expand: get a tree of all subjects who have a relation on an object
4. Change: create, update, and delete relation tuples

For all details, head over to the
[documentation](https://www.ory.sh/keto/docs/concepts/api-overview).

With this release we officially move the "old" Keto to the
[legacy-0.5 branch](https://github.com/ory/keto/tree/legacy-0.5). We will only
provide security fixes from now on. A migration path to v0.6 is planned but not
yet implemented, as the architectures are vastly different. Please refer to
[the issue](https://github.com/ory/keto/issues/318).

We are keen to bring more features and performance improvements. The next
features we will tackle are:

- Subject Set rewrites
- Native ABAC & RBAC Support
- Integration with other policy servers
- Latency reduction through aggressive caching
- Cluster mode that fans out requests over all Keto instances

So stay tuned, :star: this repo, :eyes: releases, and
[subscribe to our newsletter :email:](https://ory.us10.list-manage.com/subscribe?u=ffb1a878e4ec6c0ed312a3480&id=f605a41b53&MERGE0=&group[17097][32]=1).

### Bug Fixes

- Add description attribute to access control policy role
  ([#215](https://github.com/ory/keto/issues/215))
  ([831eba5](https://github.com/ory/keto/commit/831eba59f810ca68561dd584c9df7684df10b843))
- Add leak_sensitive_values to config schema
  ([2b21d2b](https://github.com/ory/keto/commit/2b21d2bdf4ca9523d16159c5f73c4429b692e17d))
- Bump CLI
  ([80c82d0](https://github.com/ory/keto/commit/80c82d026cbfbab8fbb84d850d8980866ecf88df))
- Bump deps and replace swagutil
  ([#212](https://github.com/ory/keto/issues/212))
  ([904258d](https://github.com/ory/keto/commit/904258d23959c3fa96b6d8ccfdb79f6788c106ec))
- Check engine overwrote result in some cases
  ([#412](https://github.com/ory/keto/issues/412))
  ([3404492](https://github.com/ory/keto/commit/3404492002ca5c3f017ef25486e377e911987aa4))
- Check health status in status command
  ([21c64d4](https://github.com/ory/keto/commit/21c64d45f21a505744b9f70d780f9b3079d3822c))
- Check REST API returns JSON object
  ([#460](https://github.com/ory/keto/issues/460))
  ([501dcff](https://github.com/ory/keto/commit/501dcff4427f76902671f6d5733f28722bd51fa7)),
  closes [#406](https://github.com/ory/keto/issues/406)
- Empty relationtuple list should not error
  ([#440](https://github.com/ory/keto/issues/440))
  ([fbcb3e1](https://github.com/ory/keto/commit/fbcb3e1f337b5114d7697fa512ded92b5f409ef4))
- Ensure nil subject is not allowed
  ([#449](https://github.com/ory/keto/issues/449))
  ([7a0fcfc](https://github.com/ory/keto/commit/7a0fcfc4fe83776fa09cf78ee11f407610554d04)):

  The nodejs gRPC client was a great fuzzer and pointed me to some nil pointer
  dereference panics. This adds some input validation to prevent panics.

- Ensure persister errors are handled by sqlcon
  ([#473](https://github.com/ory/keto/issues/473))
  ([4343c4a](https://github.com/ory/keto/commit/4343c4acd8f917fb7ae131e67bca6855d4d61694))
- Handle pagination and errors in the check/expand engines
  ([#398](https://github.com/ory/keto/issues/398))
  ([5eb1a7d](https://github.com/ory/keto/commit/5eb1a7d49af6b43707c122de8727cbd72285cb5c))
- Ignore dist
  ([ba816ea](https://github.com/ory/keto/commit/ba816ea2ca39962f02c08e0c7b75cfe3cf1d963d))
- Ignore x/net false positives
  ([d8b36cb](https://github.com/ory/keto/commit/d8b36cb1812abf7265ac15c29780222be025186b))
- Improve CLI remote sourcing ([#474](https://github.com/ory/keto/issues/474))
  ([a85f4d7](https://github.com/ory/keto/commit/a85f4d7470ac3744476e82e5889b97d5a0680473))
- Improve handlers and add tests
  ([#470](https://github.com/ory/keto/issues/470))
  ([ca5ccb9](https://github.com/ory/keto/commit/ca5ccb9c237fdcc4db031ec97a75616a859cbf8f))
- Insert relation tuples without fmt.Sprintf
  ([#443](https://github.com/ory/keto/issues/443))
  ([fe507bb](https://github.com/ory/keto/commit/fe507bb4ea719780e732d098291aa190d6b1c441))
- Minor bugfixes ([#371](https://github.com/ory/keto/issues/371))
  ([185ee1e](https://github.com/ory/keto/commit/185ee1e51bc4bcdee028f71fcaf207b7e342313b))
- Move dockerfile to where it belongs
  ([f087843](https://github.com/ory/keto/commit/f087843ac8f24e741bf39fe65ee5b0a9adf9a5bb))
- Namespace migrator ([#417](https://github.com/ory/keto/issues/417))
  ([ea79300](https://github.com/ory/keto/commit/ea7930064f490b063a712b4e18521f8996931a13)),
  closes [#404](https://github.com/ory/keto/issues/404)
- Remove SQL logging ([#455](https://github.com/ory/keto/issues/455))
  ([d8e2a86](https://github.com/ory/keto/commit/d8e2a869db2a9cfb44423b434330536036b2f421))
- Rename /relationtuple endpoint to /relation-tuples
  ([#519](https://github.com/ory/keto/issues/519))
  ([8eb55f6](https://github.com/ory/keto/commit/8eb55f6269399f2bc5f000b8a768bcdf356c756f))
- Resolve gitignore build
  ([6f04bbb](https://github.com/ory/keto/commit/6f04bbb6057779b4d73d3f94677cea365843f7ac))
- Resolve goreleaser issues
  ([d32767f](https://github.com/ory/keto/commit/d32767f32856cf5bd48514c5d61746417fbed6f5))
- Resolve windows build issues
  ([8bcdfbf](https://github.com/ory/keto/commit/8bcdfbfbdb0b10c03ff93838e8fe6e778236e96d))
- Rewrite check engine to search starting at the object
  ([#310](https://github.com/ory/keto/issues/310))
  ([7d99694](https://github.com/ory/keto/commit/7d9969414ebc8cf6ef5d211ad34f8ae01bd3b4ee)),
  closes [#302](https://github.com/ory/keto/issues/302)
- Secure query building ([#442](https://github.com/ory/keto/issues/442))
  ([c7d2770](https://github.com/ory/keto/commit/c7d2770ed570238fd1262bcc4e5b4afa6c12d80e))
- Strict version enforcement in docker
  ([e45b28f](https://github.com/ory/keto/commit/e45b28fec626db35f1bd4580e5b11c9c94a02669))
- Update dd-trace to fix build issues
  ([2ad489f](https://github.com/ory/keto/commit/2ad489f0d9cae3191718d36823fe25df58ab95e6))
- Update docker to go 1.16 and alpine
  ([c63096c](https://github.com/ory/keto/commit/c63096cb53d2171f22f4a0d4a9ac3c9bfac89d01))
- Use errors.WithStack everywhere
  ([#462](https://github.com/ory/keto/issues/462))
  ([5f25bce](https://github.com/ory/keto/commit/5f25bceea35179c67d24dd95f698dc57b789d87a)),
  closes [#437](https://github.com/ory/keto/issues/437):

  Fixed all occurrences found using the search pattern `return .*, err\n`.

- Use package name in pkger
  ([6435939](https://github.com/ory/keto/commit/6435939ad7e5899505cd0e6261f5dfc819c9ca42))
- **schema:** Add trace level to logger
  ([a5a1402](https://github.com/ory/keto/commit/a5a1402c61e1a37b1a9a349ad5736eaca66bd6a4))
- Use make() to initialize slices
  ([#250](https://github.com/ory/keto/issues/250))
  ([84f028d](https://github.com/ory/keto/commit/84f028dc35665174542e103c0aefc635bb6d3e52)),
  closes [#217](https://github.com/ory/keto/issues/217)

### Build System

- Pin dependency versions of buf and protoc plugins
  ([#338](https://github.com/ory/keto/issues/338))
  ([5a2fd1c](https://github.com/ory/keto/commit/5a2fd1cc8dff02aa7017771adc0d9101f6c86775))

### Code Generation

- Pin v0.6.0-alpha.1 release commit
  ([875af25](https://github.com/ory/keto/commit/875af25f89b813455148e58884dcdf1cd3600b86))

### Code Refactoring

- Data structures ([#279](https://github.com/ory/keto/issues/279))
  ([1316077](https://github.com/ory/keto/commit/131607762d0006e4cf4f93e8731ef7648348b2ec))

### Documentation

- Add check- and expand-API guides
  ([#493](https://github.com/ory/keto/issues/493))
  ([09a25b4](https://github.com/ory/keto/commit/09a25b4063abcfdcd4c0de315a2ef088d6d4e72e))
- Add current features overview ([#505](https://github.com/ory/keto/issues/505))
  ([605afa0](https://github.com/ory/keto/commit/605afa029794ad115bba02e004e1596cea038e8e))
- Add missing pages ([#518](https://github.com/ory/keto/issues/518))
  ([43cbaa9](https://github.com/ory/keto/commit/43cbaa9140cfa0ea3c72f699f6bb34f5ed31d8dd))
- Add namespace and relation naming conventions
  ([#510](https://github.com/ory/keto/issues/510))
  ([dd31865](https://github.com/ory/keto/commit/dd318653178cd45da47f3e7cef507b42708363ef))
- Add performance page ([#413](https://github.com/ory/keto/issues/413))
  ([6fe0639](https://github.com/ory/keto/commit/6fe0639d36087b5ecd555eb6fe5ce949f3f6f0d7)):

  This also refactored the server startup. Functionality did not change.

- Add production guide
  ([a9163c7](https://github.com/ory/keto/commit/a9163c7690c55c8191650c4dfb464b75ea02446b))
- Add zanzibar overview to README.md
  ([#265](https://github.com/ory/keto/issues/265))
  ([15a95b2](https://github.com/ory/keto/commit/15a95b28e745592353e4656d42a9d0bd20ce468f))
- API overview ([#501](https://github.com/ory/keto/issues/501))
  ([05fe03b](https://github.com/ory/keto/commit/05fe03b5bf7a3f790aa6c9c1d3fcdb31304ef6af))
- Concepts ([#429](https://github.com/ory/keto/issues/429))
  ([2f2c885](https://github.com/ory/keto/commit/2f2c88527b3f6d1d46a5c287d8aca0874d18a28d))
- Delete old redirect homepage
  ([c0a3784](https://github.com/ory/keto/commit/c0a378448f8c7723bae68f7b52a019b697b25863))
- Document gRPC SKDs
  ([7583fe8](https://github.com/ory/keto/commit/7583fe8933f6676b4e37477098b1d43d12819b8b))
- Fix grammatical error ([#222](https://github.com/ory/keto/issues/222))
  ([256a0d2](https://github.com/ory/keto/commit/256a0d2e53fe1eb859e41fc539870ae1d5a493d2))
- Fix regression issues
  ([9697bb4](https://github.com/ory/keto/commit/9697bb43dd23c0d1fae74ea42e848883c45dae77))
- Generate gRPC reference page ([#488](https://github.com/ory/keto/issues/488))
  ([93ebe6d](https://github.com/ory/keto/commit/93ebe6db7e887d708503a54c5ec943254e37ca43))
- Improve CLI documentation ([#503](https://github.com/ory/keto/issues/503))
  ([be9327f](https://github.com/ory/keto/commit/be9327f7b28152a78f731043acf83b7092e42e29))
- Minor fixes ([#532](https://github.com/ory/keto/issues/532))
  ([638342e](https://github.com/ory/keto/commit/638342eb9519d9bf609926fb87558071e2815fb3))
- Move development section
  ([9ff393f](https://github.com/ory/keto/commit/9ff393f6cba1fb0a33918377ce505455c34d9dfc))
- Move to json sidebar
  ([257bf96](https://github.com/ory/keto/commit/257bf96044df37c3d7af8a289fb67098d48da1a3))
- Remove duplicate "is"
  ([ca3277d](https://github.com/ory/keto/commit/ca3277d82c1508797bc8c663963407d2e4d9112f))
- Remove duplicate template
  ([1d3b38e](https://github.com/ory/keto/commit/1d3b38e4045b0b874bb1186ea628f5a37353a2e6))
- Remove old documentation ([#426](https://github.com/ory/keto/issues/426))
  ([eb76913](https://github.com/ory/keto/commit/eb7691306018678e024211b51627a1c27e780a6b))
- Replace TODO links ([#512](https://github.com/ory/keto/issues/512))
  ([ad8e20b](https://github.com/ory/keto/commit/ad8e20b3bef2bc46b3a32c2c9ccb6e16e4bad22c))
- Resolve broken links
  ([0d0a50b](https://github.com/ory/keto/commit/0d0a50b3f4112893f32c81adc8edd137b5a62541))
- Simple access check guide ([#451](https://github.com/ory/keto/issues/451))
  ([e0485af](https://github.com/ory/keto/commit/e0485afc46a445868580aa541e962e80cbea0670)):

  This also enables gRPC go, gRPC nodejs, cURL, and Keto CLI code samples to be
  tested.

- Update comment in write response
  ([#329](https://github.com/ory/keto/issues/329))
  ([4ca0baf](https://github.com/ory/keto/commit/4ca0baf62e34402e749e870fe8c0cc893684192c))
- Update install instructions
  ([d2e4123](https://github.com/ory/keto/commit/d2e4123f3e2e58da8be181a0a542e3dcc1313e16))
- Update introduction
  ([5f71d73](https://github.com/ory/keto/commit/5f71d73e2ee95d02abc4cd42a76c98a35942df0c))
- Update README ([#515](https://github.com/ory/keto/issues/515))
  ([18d3cd6](https://github.com/ory/keto/commit/18d3cd61b0a79400170dc0f89860b4614cc4a543)):

  Also format all markdown files in the root.

- Update repository templates
  ([db505f9](https://github.com/ory/keto/commit/db505f9e10755bc20c4623c4f5f99f33283dffda))
- Update repository templates
  ([6c056bb](https://github.com/ory/keto/commit/6c056bb2043af6e82f06fdfa509ab3fa0d5e5d06))
- Update SDK links ([#514](https://github.com/ory/keto/issues/514))
  ([f920fbf](https://github.com/ory/keto/commit/f920fbfc8dcc6711ad9e046578a4506179952be7))
- Update swagger documentation for REST endpoints
  ([c363de6](https://github.com/ory/keto/commit/c363de61edf912fef85acc6bcdac6e1c15c48f4f))
- Use mdx for api reference
  ([340f3a3](https://github.com/ory/keto/commit/340f3a3dd20c82c743e7b3ad6aaf06a4c118b5a1))
- Various improvements and updates
  ([#486](https://github.com/ory/keto/issues/486))
  ([a812ace](https://github.com/ory/keto/commit/a812ace2303214e0e7acb2e283efa1cff0d5d279))

### Features

- Add .dockerignore
  ([8b0ff06](https://github.com/ory/keto/commit/8b0ff066b2508ef2f3629f9a3e2fce601b8dcce1))
- Add and automate version schema
  ([b01eef8](https://github.com/ory/keto/commit/b01eef8d4d5834b5888cb369ecf01ee01b40c24c))
- Add check engine ([#277](https://github.com/ory/keto/issues/277))
  ([396c1ae](https://github.com/ory/keto/commit/396c1ae33b777031f8d59549d9de4a88e3f6b10a))
- Add gRPC health status ([#427](https://github.com/ory/keto/issues/427))
  ([51c4223](https://github.com/ory/keto/commit/51c4223d6cb89a9bfbc115ef20db8350aeb2e8af))
- Add is_last_page to list response
  ([#425](https://github.com/ory/keto/issues/425))
  ([b73d91f](https://github.com/ory/keto/commit/b73d91f061ab155c53d802263c0263aa39e64bdf))
- Add POST REST handler for policy check
  ([7d89860](https://github.com/ory/keto/commit/7d89860bc4a790a69f5bea5b0dbe4a2938c6729f))
- Add relation write API ([#275](https://github.com/ory/keto/issues/275))
  ([f2ddb9d](https://github.com/ory/keto/commit/f2ddb9d884ed71037b5371c00bb10b63d25d47c0))
- Add REST and gRPC logger middlewares
  ([#436](https://github.com/ory/keto/issues/436))
  ([615eb0b](https://github.com/ory/keto/commit/615eb0bec3bdc0fd26abc7af0b8990269b0cbedd))
- Add SQA telemetry ([#535](https://github.com/ory/keto/issues/535))
  ([9f6472b](https://github.com/ory/keto/commit/9f6472b0c996505d41058e9b55afa8fd6b9bb2d5))
- Add sql persister ([#350](https://github.com/ory/keto/issues/350))
  ([d595d52](https://github.com/ory/keto/commit/d595d52dabb8f4953b5c23d3a8154cac13d00306))
- Add tracing ([#536](https://github.com/ory/keto/issues/536))
  ([b57a144](https://github.com/ory/keto/commit/b57a144e0a7ec39d5831dbb79840c2b25c044e6a))
- Allow to apply namespace migrations together with regular migrations
  ([#441](https://github.com/ory/keto/issues/441))
  ([57e2bbc](https://github.com/ory/keto/commit/57e2bbc5eaebe43834f2432eb1ee2820d9cb2988))
- Delete relation tuples ([#457](https://github.com/ory/keto/issues/457))
  ([3ec8afa](https://github.com/ory/keto/commit/3ec8afa68c5b5ddc26609b9afd17cc0d06cd82bf)),
  closes [#452](https://github.com/ory/keto/issues/452)
- Dockerfile and docker compose example
  ([#390](https://github.com/ory/keto/issues/390))
  ([10cd0b3](https://github.com/ory/keto/commit/10cd0b39c12ef96710bda6ff013f7c5eeae97118))
- Expand API ([#285](https://github.com/ory/keto/issues/285))
  ([a3ca0b8](https://github.com/ory/keto/commit/a3ca0b8a109b63f443e359cd8ff18a7b3e489f84))
- Expand GPRC service and CLI ([#383](https://github.com/ory/keto/issues/383))
  ([acf2154](https://github.com/ory/keto/commit/acf21546d3e135deb77c853b751a3da3a7b16f00))
- First API draft and generation
  ([#315](https://github.com/ory/keto/issues/315))
  ([bda5d8b](https://github.com/ory/keto/commit/bda5d8b7e90d749600f5b5e169df8a6ec3705b22))
- GRPC status codes and improved error messages
  ([#467](https://github.com/ory/keto/issues/467))
  ([4a4f8c6](https://github.com/ory/keto/commit/4a4f8c6b323664329414b61e7d80d7838face730))
- GRPC version API ([#475](https://github.com/ory/keto/issues/475))
  ([89cc46f](https://github.com/ory/keto/commit/89cc46fe4a13b062693d3db4f803834ba37f4e48))
- Implement goreleaser pipeline
  ([888ac43](https://github.com/ory/keto/commit/888ac43e6f706f619b2f1b58271dd027094c9ae9)),
  closes [#410](https://github.com/ory/keto/issues/410)
- Incorporate new GRPC API structure
  ([#331](https://github.com/ory/keto/issues/331))
  ([e0916ad](https://github.com/ory/keto/commit/e0916ad00c81b24177cfe45faf77b93d2c33dc1f))
- Koanf and namespace configuration
  ([#367](https://github.com/ory/keto/issues/367))
  ([3ad32bc](https://github.com/ory/keto/commit/3ad32bc13a4d96135be8031eb6fe4c15868272ca))
- Namespace configuration ([#324](https://github.com/ory/keto/issues/324))
  ([b94f50d](https://github.com/ory/keto/commit/b94f50d1800c47a43561df5009cb38b44ccd0088))
- Namespace migrate status CLI ([#508](https://github.com/ory/keto/issues/508))
  ([e3f7ad9](https://github.com/ory/keto/commit/e3f7ad91585b616e97f85ce0f55c76406b6c4d0a)):

  This also refactors the current `migrate` and `namespace migrate` commands.

- Nodejs gRPC definitions ([#447](https://github.com/ory/keto/issues/447))
  ([3b5c313](https://github.com/ory/keto/commit/3b5c31326645adb2d5b14ced901771a7ba00fd1c)):

  Includes Typescript definitions.

- Read API ([#269](https://github.com/ory/keto/issues/269))
  ([de5119a](https://github.com/ory/keto/commit/de5119a6e3c7563cfc2e1ada12d47b27ebd7faaa)):

  This is a first draft of the read API. It is reachable by REST and gRPC calls.
  The main purpose of this PR is to establish the basic repository structure and
  define the API.

- Relationtuple parse command ([#490](https://github.com/ory/keto/issues/490))
  ([91a3cf4](https://github.com/ory/keto/commit/91a3cf47fbdb8203b799cf7c69bcf3dbbfb98b3a)):

  This command parses the relation tuple format used in the docs. It greatly
  improves the experience when copying something from the documentation. It can
  especially be used to pipe relation tuples into other commands, e.g.:

  ```shell
  echo "messages:02y_15_4w350m3#decypher@john" | \
    keto relation-tuple parse - --format json | \
    keto relation-tuple create -
  ```

- REST patch relation tuples ([#491](https://github.com/ory/keto/issues/491))
  ([d38618a](https://github.com/ory/keto/commit/d38618a9e647902ce019396ff1c33973020bf797)):

  The new PATCH handler allows transactional changes similar to the already
  existing gRPC service.

- Separate and multiplex ports based on read/write privilege
  ([#397](https://github.com/ory/keto/issues/397))
  ([6918ac3](https://github.com/ory/keto/commit/6918ac3bfa355cbd551e44376c214f412e3414e4))
- Swagger SDK ([#476](https://github.com/ory/keto/issues/476))
  ([011888c](https://github.com/ory/keto/commit/011888c2b7e2d0f7b8923c994c70e62d374a2830))

### Tests

- Add command tests ([#487](https://github.com/ory/keto/issues/487))
  ([61c28e4](https://github.com/ory/keto/commit/61c28e48a5c3f623e5cc133e69ba368c5103f414))
- Add dedicated persistence tests
  ([#416](https://github.com/ory/keto/issues/416))
  ([4e98906](https://github.com/ory/keto/commit/4e9890605edf3ea26134917a95bfa6fbb176565e))
- Add handler tests ([#478](https://github.com/ory/keto/issues/478))
  ([9315a77](https://github.com/ory/keto/commit/9315a77820d50e400b78f2f019a871be022a9887))
- Add initial e2e test ([#380](https://github.com/ory/keto/issues/380))
  ([dc5d3c9](https://github.com/ory/keto/commit/dc5d3c9d02604fddbfa56ac5ebbc1fef56a881d9))
- Add relationtuple definition tests
  ([#415](https://github.com/ory/keto/issues/415))
  ([2e3dcb2](https://github.com/ory/keto/commit/2e3dcb200a7769dc8710d311ca08a7515012fbdd))
- Enable GRPC client in e2e test
  ([#382](https://github.com/ory/keto/issues/382))
  ([4e5c6ae](https://github.com/ory/keto/commit/4e5c6aed56e5a449003956ec114ec131be068aaf))
- Improve docs sample tests ([#461](https://github.com/ory/keto/issues/461))
  ([6e0e5e6](https://github.com/ory/keto/commit/6e0e5e6184916e894fd4694cfa3a158f11fae11f))

# [0.5.6-alpha.1](https://github.com/ory/keto/compare/v0.5.5-alpha.1...v0.5.6-alpha.1) (2020-05-28)

This release bumps vulnerable transient dependencies (those are not actually
used in ORY Keto) and updates several documentation pages and improves
structured logging output. Additionally, ORY Keto now uses the updated release
pipeline!

### Bug Fixes

- Update install script
  ([21e1bf0](https://github.com/ory/keto/commit/21e1bf05177576a9d743bd11744ef6a42be50b8d))

### Chores

- Pin v0.5.6-alpha.1 release commit
  ([ed0da08](https://github.com/ory/keto/commit/ed0da08a03a910660358fc56c568692325749b6d))

# [0.5.5-alpha.1](https://github.com/ory/keto/compare/v0.5.4-alpha.1...v0.5.5-alpha.1) (2020-05-28)

This release bumps vulnerable transient dependencies (those are not actually
used in ORY Keto) and updates several documentation pages and improves
structured logging output. Additionally, ORY Keto now uses the updated release
pipeline!

### Bug Fixes

- Move deps to go_mod_indirect_pins
  ([dd3e971](https://github.com/ory/keto/commit/dd3e971ac418baf10c1b33005acc7e6f66fb0d85))
- Resolve test issues
  ([9bd9956](https://github.com/ory/keto/commit/9bd9956e33731f1619c32e1e6b7c78f42e7c47c3))
- Update install.sh script
  ([f64d320](https://github.com/ory/keto/commit/f64d320b6424fe3256eb7fad1c94dcc1ef0bf487))
- Use semver-regex replacer func
  ([2cc3bbb](https://github.com/ory/keto/commit/2cc3bbb2d75ba5fa7a3653d7adcaa712ff38c603))

### Chores

- Pin v0.5.5-alpha.1 release commit
  ([4666a0f](https://github.com/ory/keto/commit/4666a0f258f253d19a14eca34f4b7049f2d0afa2))

### Documentation

- Add missing colon in docker run command
  ([#193](https://github.com/ory/keto/issues/193))
  ([383063d](https://github.com/ory/keto/commit/383063d260d995665da4c02c9a7bac7e06a2c8d3))
- Update github templates ([#182](https://github.com/ory/keto/issues/182))
  ([72ea09b](https://github.com/ory/keto/commit/72ea09bbbf9925d7705842703b32826376f636e4))
- Update github templates ([#184](https://github.com/ory/keto/issues/184))
  ([ed546b7](https://github.com/ory/keto/commit/ed546b7a2b9ee690284a48c641edd1570464d71f))
- Update github templates ([#188](https://github.com/ory/keto/issues/188))
  ([ebd75b2](https://github.com/ory/keto/commit/ebd75b2f6545ff4372773f6370300c7b2ca71c51))
- Update github templates ([#189](https://github.com/ory/keto/issues/189))
  ([fd4c0b1](https://github.com/ory/keto/commit/fd4c0b17bcb1c281baac1772ab94e305ec8c5c86))
- Update github templates ([#195](https://github.com/ory/keto/issues/195))
  ([ba0943c](https://github.com/ory/keto/commit/ba0943c45d36ef10bdf1169f0aeef439a3a67d28))
- Update linux install guide ([#191](https://github.com/ory/keto/issues/191))
  ([7d8b24b](https://github.com/ory/keto/commit/7d8b24bddb9c92feb78c7b65f39434d538773b58))
- Update repository templates
  ([ea65b5c](https://github.com/ory/keto/commit/ea65b5c5ada0a7453326fa755aa914306f1b1851))
- Use central banner repo for README
  ([0d95d97](https://github.com/ory/keto/commit/0d95d97504df4d0ab57d18dc6d0a824a3f8f5896))
- Use correct banner
  ([c6dfe28](https://github.com/ory/keto/commit/c6dfe280fd962169c424834cea040a408c1bc83f))
- Use correct version
  ([5f7030c](https://github.com/ory/keto/commit/5f7030c9069fe392200be72f8ce1a93890fbbba8)),
  closes [#200](https://github.com/ory/keto/issues/200)
- Use correct versions in install docs
  ([52e6c34](https://github.com/ory/keto/commit/52e6c34780ed41c169504d71c39459898b5d14f9))

# [0.5.4-alpha.1](https://github.com/ory/keto/compare/v0.5.3-alpha.3...v0.5.4-alpha.1) (2020-04-07)

fix: resolve panic when executing migrations (#178)

Closes #177

### Bug Fixes

- Resolve panic when executing migrations
  ([#178](https://github.com/ory/keto/issues/178))
  ([7e83fee](https://github.com/ory/keto/commit/7e83feefaad041c60f09232ac44ed8b7240c6558)),
  closes [#177](https://github.com/ory/keto/issues/177)

# [0.5.3-alpha.3](https://github.com/ory/keto/compare/v0.5.3-alpha.2...v0.5.3-alpha.3) (2020-04-06)

autogen(docs): regenerate and update changelog

### Code Generation

- **docs:** Regenerate and update changelog
  ([769cef9](https://github.com/ory/keto/commit/769cef90f27ba9c203d3faf47272287ab17dc7eb))

### Code Refactoring

- Move docs to this repository ([#172](https://github.com/ory/keto/issues/172))
  ([312480d](https://github.com/ory/keto/commit/312480de3cefc5b72916ba95d8287443cf3ccb3d))

### Documentation

- Regenerate and update changelog
  ([dda79b1](https://github.com/ory/keto/commit/dda79b106a18bc33d70ae60e352118b0d288d26b))
- Regenerate and update changelog
  ([9048dd8](https://github.com/ory/keto/commit/9048dd8d8a0f0654072b3d4b77261fe947a34ece))
- Regenerate and update changelog
  ([806f68c](https://github.com/ory/keto/commit/806f68c603781742e0177ec0b2deecaf64c5b721))
- Regenerate and update changelog
  ([8905ee7](https://github.com/ory/keto/commit/8905ee74d4ec394af92240e180cc5d7f6493cb2f))
- Regenerate and update changelog
  ([203c1cc](https://github.com/ory/keto/commit/203c1cc659a72f81a370d7b9b7fbda60e7c96c9e))
- Regenerate and update changelog
  ([8875a95](https://github.com/ory/keto/commit/8875a95b35df57668acb27820a3aff1cdfbe8b30))
- Regenerate and update changelog
  ([28ddd3e](https://github.com/ory/keto/commit/28ddd3e1483afe8571b3d2bf9afcc31386d85f7f))
- Regenerate and update changelog
  ([927c4ed](https://github.com/ory/keto/commit/927c4edc4a770133bcb34bc044dd5c5e0eb3ffb7))
- Updates issue and pull request templates
  ([#168](https://github.com/ory/keto/issues/168))
  ([29a38a8](https://github.com/ory/keto/commit/29a38a85c61ec2c8d0ad2ce6d5a0f9e9d74b52f7))
- Updates issue and pull request templates
  ([#169](https://github.com/ory/keto/issues/169))
  ([99b7d5d](https://github.com/ory/keto/commit/99b7d5de24fed1aed746c4447a390d084632f89a))
- Updates issue and pull request templates
  ([#171](https://github.com/ory/keto/issues/171))
  ([7a9876b](https://github.com/ory/keto/commit/7a9876b8ed4282f50f886a025033641bd027a0e2))

# [0.5.3-alpha.1](https://github.com/ory/keto/compare/v0.5.2...v0.5.3-alpha.1) (2020-04-03)

chore: move to ory analytics fork (#167)

### Chores

- Move to ory analytics fork ([#167](https://github.com/ory/keto/issues/167))
  ([f824011](https://github.com/ory/keto/commit/f824011b4d19058504b3a43ed53a420619444a51))

# [0.5.2](https://github.com/ory/keto/compare/v0.5.1-alpha.1...v0.5.2) (2020-04-02)

docs: Regenerate and update changelog

### Documentation

- Regenerate and update changelog
  ([1e52100](https://github.com/ory/keto/commit/1e521001a43a0a13e2224e1a44956442ac6ffbc7))
- Regenerate and update changelog
  ([e4d32a6](https://github.com/ory/keto/commit/e4d32a62c1ae96115ea50bb471f5ff2ce2f2c4b9))

# [0.5.0](https://github.com/ory/keto/compare/v0.4.5-alpha.1...v0.5.0) (2020-04-02)

docs: use real json bool type in swagger (#162)

Closes #160

### Bug Fixes

- Move to ory sqa service ([#159](https://github.com/ory/keto/issues/159))
  ([c3bf1b1](https://github.com/ory/keto/commit/c3bf1b1964a14be4cc296aae98d0739e65917e18))
- Use correct response mode for removeOryAccessControlPolicyRoleMe…
  ([#161](https://github.com/ory/keto/issues/161))
  ([17543cf](https://github.com/ory/keto/commit/17543cfef63a1d040a2234bd63b210fb9c4f6015))

### Documentation

- Regenerate and update changelog
  ([6a77f75](https://github.com/ory/keto/commit/6a77f75d66e89420f2daf2fae011d31bcfa34008))
- Regenerate and update changelog
  ([c8c9d29](https://github.com/ory/keto/commit/c8c9d29e77ef53e1196cc6fe600c53d93376229b))
- Regenerate and update changelog
  ([fe8327d](https://github.com/ory/keto/commit/fe8327d951394084df7785166c9a9578c1ab0643))
- Regenerate and update changelog
  ([b5b1d66](https://github.com/ory/keto/commit/b5b1d66a4b933df8789337cce3f6d6bf391b617b))
- Update forum and chat links
  ([e96d7ba](https://github.com/ory/keto/commit/e96d7ba3dcc693c22eb983b3f58a05c9c6adbda7))
- Updates issue and pull request templates
  ([#158](https://github.com/ory/keto/issues/158))
  ([ab14cfa](https://github.com/ory/keto/commit/ab14cfa51ce195b26a83c050452530a5008589d7))
- Use real json bool type in swagger
  ([#162](https://github.com/ory/keto/issues/162))
  ([5349e7f](https://github.com/ory/keto/commit/5349e7f910ad22558a01b76be62db2136b5eb301)),
  closes [#160](https://github.com/ory/keto/issues/160)

# [0.4.5-alpha.1](https://github.com/ory/keto/compare/v0.4.4-alpha.1...v0.4.5-alpha.1) (2020-02-29)

docs: Regenerate and update changelog

### Bug Fixes

- **driver:** Extract scheme from DSN using sqlcon.GetDriverName
  ([#156](https://github.com/ory/keto/issues/156))
  ([187e289](https://github.com/ory/keto/commit/187e289f1a235b5cacf2a0b7ca5e98c384fa7a14)),
  closes [#145](https://github.com/ory/keto/issues/145)

### Documentation

- Regenerate and update changelog
  ([41513da](https://github.com/ory/keto/commit/41513da35ea038f3c4cc2d98b9796cee5b5a8b92))

# [0.4.4-alpha.1](https://github.com/ory/keto/compare/v0.4.3-alpha.2...v0.4.4-alpha.1) (2020-02-14)

docs: Regenerate and update changelog

### Bug Fixes

- **goreleaser:** Update brew section
  ([0918ff3](https://github.com/ory/keto/commit/0918ff3032eeecd26c67d6249c7e28e71ee110af))

### Documentation

- Prepare ecosystem automation
  ([2e39be7](https://github.com/ory/keto/commit/2e39be79ebad1cec021ae3ee4b0a75ffea4b7424))
- Regenerate and update changelog
  ([009c4c4](https://github.com/ory/keto/commit/009c4c4e4fd4c5607cc30cc9622fd0f82e3891f3))
- Regenerate and update changelog
  ([49f3c4b](https://github.com/ory/keto/commit/49f3c4ba34df5879d8f48cc96bf0df9dad820362))
- Updates issue and pull request templates
  ([#153](https://github.com/ory/keto/issues/153))
  ([7fb7521](https://github.com/ory/keto/commit/7fb752114e1e2a91ab96fdb546835de8aee4926b))

### Features

- **ci:** Add nancy vuln scanner
  ([#152](https://github.com/ory/keto/issues/152))
  ([c19c2b9](https://github.com/ory/keto/commit/c19c2b9efe8299b8878cc8099fe314d8dcda3a08))

### Unclassified

- Update CHANGELOG [ci skip]
  ([63fe513](https://github.com/ory/keto/commit/63fe513d22ec3747a95cdb8f797ba1eba5ca344f))
- Update CHANGELOG [ci skip]
  ([7b7c3ac](https://github.com/ory/keto/commit/7b7c3ac6c06c072fea1b64624ea79a3fd406b09c))
- Update CHANGELOG [ci skip]
  ([8886392](https://github.com/ory/keto/commit/8886392b39fb46ad338c8284866d4dae64ad1826))
- Update CHANGELOG [ci skip]
  ([5bbc284](https://github.com/ory/keto/commit/5bbc2844c49b0a68ba3bd8b003d91f87e2aed9e2))

# [0.4.3-alpha.2](https://github.com/ory/keto/compare/v0.4.3-alpha.1...v0.4.3-alpha.2) (2020-01-31)

Update README.md

### Unclassified

- Update README.md
  ([0ab9c6f](https://github.com/ory/keto/commit/0ab9c6f372a1538a958a68b34315c9167b5a9093))
- Update CHANGELOG [ci skip]
  ([f0a1428](https://github.com/ory/keto/commit/f0a1428f4b99ceb35ff4f1e839bc5237e19db628))

# [0.4.3-alpha.1](https://github.com/ory/keto/compare/v0.4.2-alpha.1...v0.4.3-alpha.1) (2020-01-23)

Disable access logging for health endpoints (#151)

Closes #150

### Unclassified

- Disable access logging for health endpoints (#151)
  ([6ca0c09](https://github.com/ory/keto/commit/6ca0c09b5618122762475cffdc9c32adf28456a1)),
  closes [#151](https://github.com/ory/keto/issues/151)
  [#150](https://github.com/ory/keto/issues/150)

# [0.4.2-alpha.1](https://github.com/ory/keto/compare/v0.4.1-beta.1...v0.4.2-alpha.1) (2020-01-14)

Update CHANGELOG [ci skip]

### Unclassified

- Update CHANGELOG [ci skip]
  ([afaabde](https://github.com/ory/keto/commit/afaabde63affcf568e3090e55b4b957edff2890c))

# [0.4.1-beta.1](https://github.com/ory/keto/compare/v0.4.0-sandbox...v0.4.1-beta.1) (2020-01-13)

Update CHANGELOG [ci skip]

### Unclassified

- Update CHANGELOG [ci skip]
  ([e3ca5a7](https://github.com/ory/keto/commit/e3ca5a7d8b9827ffc7b31a8b5e459db3e912a590))
- Update SDK
  ([5dd6237](https://github.com/ory/keto/commit/5dd623755d4832f33c3dcefb778a9a70eace7b52))

# [0.4.0-alpha.1](https://github.com/ory/keto/compare/v0.3.9-sandbox...v0.4.0-alpha.1) (2020-01-13)

Move to new SDK generators (#146)

### Unclassified

- Move to new SDK generators (#146)
  ([4f51a09](https://github.com/ory/keto/commit/4f51a0948723efc092f1887b111d1e6dd590a075)),
  closes [#146](https://github.com/ory/keto/issues/146)
- Fix typos in the README (#144)
  ([85d838c](https://github.com/ory/keto/commit/85d838c0872c73eb70b5bfff1ccb175b07f6b1e4)),
  closes [#144](https://github.com/ory/keto/issues/144)

# [0.3.9-sandbox](https://github.com/ory/keto/compare/v0.3.8-sandbox...v0.3.9-sandbox) (2019-12-16)

Update go modules

### Unclassified

- Update go modules
  ([1151e07](https://github.com/ory/keto/commit/1151e0755c974b0aea86be5aaeae365ea9aef094))

# [0.3.7-sandbox](https://github.com/ory/keto/compare/v0.3.6-sandbox...v0.3.7-sandbox) (2019-12-11)

Update documentation banner image (#143)

### Unclassified

- Update documentation banner image (#143)
  ([e444755](https://github.com/ory/keto/commit/e4447552031a4f26ec21a336071b0bb19843df61)),
  closes [#143](https://github.com/ory/keto/issues/143)
- Revert incorrect license changes
  ([094c4f3](https://github.com/ory/keto/commit/094c4f30184d77a05044087c13e71ce4adb4d735))
- Fix invalid pseudo version ([#138](https://github.com/ory/keto/issues/138))
  ([79b4457](https://github.com/ory/keto/commit/79b4457f0162197ba267edbb8c0031c47e03bade))

# [0.3.6-sandbox](https://github.com/ory/keto/compare/v0.3.5-sandbox...v0.3.6-sandbox) (2019-10-16)

Resolve issues with mysql tests (#137)

### Unclassified

- Resolve issues with mysql tests (#137)
  ([ef5aec8](https://github.com/ory/keto/commit/ef5aec8e493199c46b78e8f1257aa41df9545f28)),
  closes [#137](https://github.com/ory/keto/issues/137)

# [0.3.5-sandbox](https://github.com/ory/keto/compare/v0.3.4-sandbox...v0.3.5-sandbox) (2019-08-21)

Implement roles and policies filter (#124)

### Documentation

- Incorporates changes from version v0.3.3-sandbox
  ([57686d2](https://github.com/ory/keto/commit/57686d2e30b229cae33e717eb8b3db9da3bdaf0a))
- README grammar fixes ([#114](https://github.com/ory/keto/issues/114))
  ([e592736](https://github.com/ory/keto/commit/e5927360300d8c4fbea841c1c2fb92b48b77885e))
- Updates issue and pull request templates
  ([#110](https://github.com/ory/keto/issues/110))
  ([80c8516](https://github.com/ory/keto/commit/80c8516efbcf33902d8a45f1dc7dbafff2aab8b1))
- Updates issue and pull request templates
  ([#111](https://github.com/ory/keto/issues/111))
  ([22305d0](https://github.com/ory/keto/commit/22305d0a9b5114de8125c16030bbcd1de695ae9b))
- Updates issue and pull request templates
  ([#112](https://github.com/ory/keto/issues/112))
  ([dccada9](https://github.com/ory/keto/commit/dccada9a2189bbd899c5c4a18665a92113fe6cd7))
- Updates issue and pull request templates
  ([#125](https://github.com/ory/keto/issues/125))
  ([15f373a](https://github.com/ory/keto/commit/15f373a16b8cfbd6cdad2bda5f161e171c566137))
- Updates issue and pull request templates
  ([#128](https://github.com/ory/keto/issues/128))
  ([eaf8e33](https://github.com/ory/keto/commit/eaf8e33f3904484635924bdac190c8dc7b60f939))
- Updates issue and pull request templates
  ([#130](https://github.com/ory/keto/issues/130))
  ([a440d14](https://github.com/ory/keto/commit/a440d142275a7a91a0a6bb487fe47d22247f4988))
- Updates issue and pull request templates
  ([#131](https://github.com/ory/keto/issues/131))
  ([dbf2cb2](https://github.com/ory/keto/commit/dbf2cb23c5b6f0f1ee0be5c0b5a58fb0c3dbefd1))
- Updates issue and pull request templates
  ([#132](https://github.com/ory/keto/issues/132))
  ([e121048](https://github.com/ory/keto/commit/e121048d10627ed32a07e26455efd69248f1bd95))
- Updates issue and pull request templates
  ([#133](https://github.com/ory/keto/issues/133))
  ([1b7490a](https://github.com/ory/keto/commit/1b7490abc1d5d0501b66595eb2d92834b6fb0345))

### Unclassified

- Implement roles and policies filter (#124)
  ([db94481](https://github.com/ory/keto/commit/db9448103621a6a8cd086a4cef6c6a22398e621f)),
  closes [#124](https://github.com/ory/keto/issues/124)
- Add adopters placeholder ([#129](https://github.com/ory/keto/issues/129))
  ([b814838](https://github.com/ory/keto/commit/b8148388b8bea97d1f1b4b54de2f0b8ef6b8b6c7))
- Improve documentation (#126)
  ([aabb04d](https://github.com/ory/keto/commit/aabb04d5f283d3c73eb3f3531b4e470ae716db5e)),
  closes [#126](https://github.com/ory/keto/issues/126)
- Create FUNDING.yml
  ([571b447](https://github.com/ory/keto/commit/571b447ed3a02f43623ef5c5adc09682b5f379bd))
- Use non-root user in image ([#116](https://github.com/ory/keto/issues/116))
  ([a493e55](https://github.com/ory/keto/commit/a493e550a8bb86d99164f4ea76dbcecf76c9c2c1))
- Remove binary license (#117)
  ([6e85f7c](https://github.com/ory/keto/commit/6e85f7c6f430e88fb4117a131f57bd69466a8ca1)),
  closes [#117](https://github.com/ory/keto/issues/117)

# [0.3.3-sandbox](https://github.com/ory/keto/compare/v0.3.1-sandbox...v0.3.3-sandbox) (2019-05-18)

ci: Resolve goreleaser issues (#108)

### Continuous Integration

- Resolve goreleaser issues ([#108](https://github.com/ory/keto/issues/108))
  ([5753f27](https://github.com/ory/keto/commit/5753f27a9e89ccdda7c02969217c253aa72cb94b))

### Documentation

- Incorporates changes from version v0.3.1-sandbox
  ([b8a0029](https://github.com/ory/keto/commit/b8a002937483a0f71fe5aba26bb18beb41886249))
- Updates issue and pull request templates
  ([#106](https://github.com/ory/keto/issues/106))
  ([54a5a27](https://github.com/ory/keto/commit/54a5a27f24a90ab3c5f9915f36582b85eecd0d62))

# [0.3.1-sandbox](https://github.com/ory/keto/compare/v0.3.0-sandbox...v0.3.1-sandbox) (2019-04-29)

ci: Use image that includes bash/sh for release docs (#103)

Signed-off-by: aeneasr <aeneas@ory.sh>

### Continuous Integration

- Use image that includes bash/sh for release docs
  ([#103](https://github.com/ory/keto/issues/103))
  ([e9d3027](https://github.com/ory/keto/commit/e9d3027fc62b20f28cd7a023222390e24d565eb1))

### Documentation

- Incorporates changes from version v0.3.0-sandbox
  ([605d2f4](https://github.com/ory/keto/commit/605d2f43621b806b750edc81d439edc92cfb7c38))

### Unclassified

- Allow configuration files and update UPGRADE guide. (#102)
  ([3934dc6](https://github.com/ory/keto/commit/3934dc6e690822358067b43920048d45a4b7799b)),
  closes [#102](https://github.com/ory/keto/issues/102)

# [0.3.0-sandbox](https://github.com/ory/keto/compare/v0.2.3-sandbox+oryOS.10...v0.3.0-sandbox) (2019-04-29)

docker: Remove full tag from build pipeline (#101)

Signed-off-by: aeneasr <aeneas@ory.sh>

### Documentation

- Update patrons
  ([c8dc7cd](https://github.com/ory/keto/commit/c8dc7cdc68676970328b55648b8d6e469c77fbfd))

### Unclassified

- Improve naming for ory policies
  ([#100](https://github.com/ory/keto/issues/100))
  ([b39703d](https://github.com/ory/keto/commit/b39703d362d333213fcb7d3782e363d09b6dabbd))
- Remove full tag from build pipeline
  ([#101](https://github.com/ory/keto/issues/101))
  ([602a273](https://github.com/ory/keto/commit/602a273dc5a0c29e80a22f04adb937ab385c4512))
- Remove duplicate code in Makefile (#99)
  ([04f5223](https://github.com/ory/keto/commit/04f52231509dd0f3a57d745918fc43fff7c595ff)),
  closes [#99](https://github.com/ory/keto/issues/99)
- Add tracing support and general improvements (#98)
  ([63b3946](https://github.com/ory/keto/commit/63b3946e0ae1fa23c6a359e9a64b296addff868c)),
  closes [#98](https://github.com/ory/keto/issues/98):

  This patch improves the internal configuration and service management. It adds
  support for distributed tracing and resolves several issues in the release
  pipeline and CLI.

  Additionally, composable docker-compose configuration files have been added.

  Several bugs have been fixed in the release management pipeline.

- Add content-type in the response of allowed
  ([#90](https://github.com/ory/keto/issues/90))
  ([39a1486](https://github.com/ory/keto/commit/39a1486dc53456189d30380460a9aeba198fa9e9))
- Fix disable-telemetry check ([#85](https://github.com/ory/keto/issues/85))
  ([38b5383](https://github.com/ory/keto/commit/38b538379973fa34bd2bf24dcb2e6dbedf324e1e))
- Fix remove member from role ([#87](https://github.com/ory/keto/issues/87))
  ([698e161](https://github.com/ory/keto/commit/698e161989331ca5a3a0769301d9694ef805a876)),
  closes [#74](https://github.com/ory/keto/issues/74)
- Fix the type of conditions in the policy
  ([#86](https://github.com/ory/keto/issues/86))
  ([fc1ced6](https://github.com/ory/keto/commit/fc1ced63bd39c9fbf437e419dfc384343e36e0ee))
- Move Go SDK generation to go-swagger
  ([#94](https://github.com/ory/keto/issues/94))
  ([9f48a95](https://github.com/ory/keto/commit/9f48a95187a7b6160108cd7d0301590de2e58f07)),
  closes [#92](https://github.com/ory/keto/issues/92)
- Send 403 when authorization result is negative
  ([#93](https://github.com/ory/keto/issues/93))
  ([de806d8](https://github.com/ory/keto/commit/de806d892819db63c1abc259ab06ee08d87895dc)),
  closes [#75](https://github.com/ory/keto/issues/75)
- Update dependencies ([#91](https://github.com/ory/keto/issues/91))
  ([4d44174](https://github.com/ory/keto/commit/4d4417474ebf8cc69d01e5ac82633b966cdefbc7))
- storage/memory: Fix upsert with pre-existing key will causes duplicate records
  (#88)
  ([1cb8a36](https://github.com/ory/keto/commit/1cb8a36a08883b785d9bb0a4be1ddc00f1f9d358)),
  closes [#88](https://github.com/ory/keto/issues/88)
  [#80](https://github.com/ory/keto/issues/80)

# [0.2.3-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.2.2-sandbox+oryOS.10...v0.2.3-sandbox+oryOS.10) (2019-02-05)

dist: Fix packr build pipeline (#84)

Closes #73 Closes #81

Signed-off-by: aeneasr <aeneas@ory.sh>

### Documentation

- Add documentation for glob matching
  ([5c8babb](https://github.com/ory/keto/commit/5c8babbfbae01a78f30cfbff92d8e9c3a6b09027))
- Incorporates changes from version v0.2.2-sandbox+oryOS.10
  ([ed7af3f](https://github.com/ory/keto/commit/ed7af3fa4e5d1d0d03b5366f4cf865a5b82ec293))
- Properly generate api.swagger.json
  ([18e3f84](https://github.com/ory/keto/commit/18e3f84cdeee317f942d61753399675c98886e5d))

### Unclassified

- Add placeholder go file for rego inclusion
  ([6a6f64d](https://github.com/ory/keto/commit/6a6f64d8c59b496f6cf360f55eba1e16bf5380f1))
- Add support for glob matching
  ([bb76c6b](https://github.com/ory/keto/commit/bb76c6bebe522fc25448c4f4e4d1ef7c530a725f))
- Ex- and import rego subdirectories for `go get`
  [#77](https://github.com/ory/keto/issues/77)
  ([59cc053](https://github.com/ory/keto/commit/59cc05328f068fc3046b2dbc022a562fd5d67960)),
  closes [#73](https://github.com/ory/keto/issues/73)
- Fix packr build pipeline ([#84](https://github.com/ory/keto/issues/84))
  ([65a87d5](https://github.com/ory/keto/commit/65a87d564d237bc979bb5962beff7d3703d9689f)),
  closes [#73](https://github.com/ory/keto/issues/73)
  [#81](https://github.com/ory/keto/issues/81)
- Import glob in rego/doc.go
  ([7798442](https://github.com/ory/keto/commit/7798442553cfe7989a23d2c389c8c63a24013543))
- Properly handle dbal error
  ([6811607](https://github.com/ory/keto/commit/6811607ea79c8f3155a17bc1aea566e9e4680616))
- Properly handle TLS certificates if set
  ([36399f0](https://github.com/ory/keto/commit/36399f09261d4f3cb5e053679eee3cb15da2df19)),
  closes [#73](https://github.com/ory/keto/issues/73)

# [0.2.2-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.2.1-sandbox+oryOS.10...v0.2.2-sandbox+oryOS.10) (2018-12-13)

ci: Fix docker push arguments in publish task

Signed-off-by: aeneasr <aeneas@ory.sh>

### Continuous Integration

- Fix docker push arguments in publish task
  ([f03c77c](https://github.com/ory/keto/commit/f03c77c6b7461ab81cb03265cbec909ac45c2259))

# [0.2.1-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.2.0-sandbox+oryOS.10...v0.2.1-sandbox+oryOS.10) (2018-12-13)

ci: Fix docker release task

Signed-off-by: aeneasr <aeneas@ory.sh>

### Continuous Integration

- Fix docker release task
  ([7a0414f](https://github.com/ory/keto/commit/7a0414f614b6cc8b1d78cfbb773a2f0192d00d23))

# [0.2.0-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.0.1...v0.2.0-sandbox+oryOS.10) (2018-12-13)

all: gofmt

Signed-off-by: aeneasr <aeneas@ory.sh>

### Documentation

- Adds banner
  ([0ec1d8f](https://github.com/ory/keto/commit/0ec1d8f5e843465d17ac4c8f91f18e5badf16900))
- Adds GitHub templates & code of conduct
  ([#31](https://github.com/ory/keto/issues/31))
  ([a11e898](https://github.com/ory/keto/commit/a11e8980f2af528f1357659966123d0cbf7d43db))
- Adds link to examples repository
  ([#32](https://github.com/ory/keto/issues/32))
  ([7061a2a](https://github.com/ory/keto/commit/7061a2aa31652a9e0c2d449facb1201bfa11fd3f))
- Adds security console image
  ([fd27fc9](https://github.com/ory/keto/commit/fd27fc9cce50beb3d0189e0a93300879fd7149db))
- Changes hydra to keto in readme
  ([9dab531](https://github.com/ory/keto/commit/9dab531744cf5b0ae98862945d44b07535595781))
- Deprecate old versions in logs
  ([955d647](https://github.com/ory/keto/commit/955d647307a48ee7cf2d3f9fb4263072adf42299))
- Incorporates changes from version
  ([85c4d81](https://github.com/ory/keto/commit/85c4d81a192e92f874c106b91cfa6fb404d9a34a))
- Incorporates changes from version v0.0.0-testrelease.1
  ([6062dd4](https://github.com/ory/keto/commit/6062dd4a894607f5f1ead119af20cc8bdbe15bef))
- Incorporates changes from version v0.0.1-1-g85c4d81
  ([f4606fc](https://github.com/ory/keto/commit/f4606fce0326bece2a89dadc029bc5ce9778df18))
- Incorporates changes from version v0.0.1-11-g114914f
  ([92a4dca](https://github.com/ory/keto/commit/92a4dca7a41dcf3a88c4800bf6d2217f33cfcdd1))
- Incorporates changes from version v0.0.1-16-g7d8a8ad
  ([2b76a83](https://github.com/ory/keto/commit/2b76a83755153b3f8a2b8d28c5b0029d96d567b6))
- Incorporates changes from version v0.0.1-18-g099e7e0
  ([70b12ad](https://github.com/ory/keto/commit/70b12adf5bcc0e890d6707e11e891e6cedfb3d87))
- Incorporates changes from version v0.0.1-20-g97ccbe6
  ([b21d56e](https://github.com/ory/keto/commit/b21d56e599c7eb4c1769bc18878f7d5818b73023))
- Incorporates changes from version v0.0.1-30-gaf2c3b5
  ([a1d0dcc](https://github.com/ory/keto/commit/a1d0dcc78a9506260f86df00e4dff8ab02909ce1))
- Incorporates changes from version v0.0.1-32-gedb5a60
  ([a5c369a](https://github.com/ory/keto/commit/a5c369a90da67c96bbde60e673c67f50b841fadd))
- Incorporates changes from version v0.0.1-6-g570783e
  ([0fcbbcb](https://github.com/ory/keto/commit/0fcbbcb02f1d748f9c733c86368b223b2ee4c6e2))
- Incorporates changes from version v0.0.1-7-g0fcbbcb
  ([c0141a8](https://github.com/ory/keto/commit/c0141a8ec22ea1260bf2d45d72dfe06737ec0115))
- Incorporates changes from version v0.1.0-sandbox
  ([9ee0664](https://github.com/ory/keto/commit/9ee06646d2cfb2d69abdcc411e31d14957437a1e))
- Incorporates changes from version v1.0.0-beta.1-1-g162d7b8
  ([647c5a9](https://github.com/ory/keto/commit/647c5a9e1bc8d9d635bf6f2511c3faa9a9daefef))
- Incorporates changes from version v1.0.0-beta.2-11-g2b280bb
  ([936889d](https://github.com/ory/keto/commit/936889d760f04a03d498f65331d653cbad3702d0))
- Incorporates changes from version v1.0.0-beta.2-13-g382e1d3
  ([883df44](https://github.com/ory/keto/commit/883df44a922f3daee86597af467072555cadc7e7))
- Incorporates changes from version v1.0.0-beta.2-15-g74450da
  ([48dd9f1](https://github.com/ory/keto/commit/48dd9f1ffbeaa99ac8dc27085c5a50f9244bf9c3))
- Incorporates changes from version v1.0.0-beta.2-3-gf623c52
  ([b6b90e5](https://github.com/ory/keto/commit/b6b90e5b2180921f78064a60666704b4e72679b6))
- Incorporates changes from version v1.0.0-beta.2-5-g3852be5
  ([3f09090](https://github.com/ory/keto/commit/3f09090a2f82f3f29154c19217cea0a10d65ea3a))
- Incorporates changes from version v1.0.0-beta.2-9-gc785187
  ([4c30a3c](https://github.com/ory/keto/commit/4c30a3c0ad83ba80e1857b41211e7ddade06c4cf))
- Incorporates changes from version v1.0.0-beta.3-1-g06adbf1
  ([0ba3c06](https://github.com/ory/keto/commit/0ba3c0674832b641ef5e0c3f0d60d81ed3a647b2))
- Incorporates changes from version v1.0.0-beta.3-10-g9994967
  ([d2345ca](https://github.com/ory/keto/commit/d2345ca3beb354d6ee7c7926c1a5ddb425d6b405))
- Incorporates changes from version v1.0.0-beta.3-12-gc28b521
  ([b4d792f](https://github.com/ory/keto/commit/b4d792f74055853f05ca46c67625ffd432fc74fd))
- Incorporates changes from version v1.0.0-beta.3-3-g9e16605
  ([c43bf2b](https://github.com/ory/keto/commit/c43bf2b5232bed9106dd47d7eb53d2f93bfe260d))
- Incorporates changes from version v1.0.0-beta.3-5-ga11e898
  ([b9d9b8e](https://github.com/ory/keto/commit/b9d9b8ee33ab957f43f99c427a88ade847e79ed0))
- Incorporates changes from version v1.0.0-beta.3-8-g7061a2a
  ([d76ff9d](https://github.com/ory/keto/commit/d76ff9dc9a4c8a8f1286eeb139d8f5af9617f421))
- Incorporates changes from version v1.0.0-beta.5
  ([0dc314c](https://github.com/ory/keto/commit/0dc314c7888020b40e12eb59fd77135044fd063b))
- Incorporates changes from version v1.0.0-beta.6-1-g5e97104
  ([f14c8ed](https://github.com/ory/keto/commit/f14c8edd7204a811e333ea84429cf837b4e7d27b))
- Incorporates changes from version v1.0.0-beta.8
  ([5045b59](https://github.com/ory/keto/commit/5045b59e2a83d6ab047b1b95c581d7c34e96a2e0))
- Incorporates changes from version v1.0.0-beta.9
  ([be2f035](https://github.com/ory/keto/commit/be2f03524721ef47ecb1c9aec57c2696174e0657))
- Properly sets up changelog TOC
  ([e0acd67](https://github.com/ory/keto/commit/e0acd670ab19c0d6fd36733fea164e2b0414597d))
- Puts toc in the right place
  ([114914f](https://github.com/ory/keto/commit/114914fa354f784b310bc9dfd232a011e0d98d99))
- Revert changes from test release
  ([ab3a64d](https://github.com/ory/keto/commit/ab3a64d3d41292364c5947db98c4d27a8223853e))
- Update documentation links ([#67](https://github.com/ory/keto/issues/67))
  ([d22d413](https://github.com/ory/keto/commit/d22d413c7a001ccaa96b4c013665153f41831614))
- Update link to security console
  ([846ce4b](https://github.com/ory/keto/commit/846ce4baa9da5954bd30996f489885a026c48185))
- Update migration guide
  ([3c44b58](https://github.com/ory/keto/commit/3c44b58613e46ed39d42463537773fe9d95a54da))
- Update to latest changes
  ([1625123](https://github.com/ory/keto/commit/1625123ed342f019d5e7ab440eb37da310570842))
- Updates copyright notice
  ([9dd5578](https://github.com/ory/keto/commit/9dd557825dfd3b9d589c9db2ccb201638debbaae))
- Updates installation guide
  ([f859645](https://github.com/ory/keto/commit/f859645f230f405cfabed0c1b9a2b67b1a3841d3))
- Updates issue and pull request templates
  ([#52](https://github.com/ory/keto/issues/52))
  ([941cae6](https://github.com/ory/keto/commit/941cae6fee058f68eabbbf4dd9cafad4760e108f))
- Updates issue and pull request templates
  ([#53](https://github.com/ory/keto/issues/53))
  ([7b222d2](https://github.com/ory/keto/commit/7b222d285e74c0db482136b23f37072216b3acb0))
- Updates issue and pull request templates
  ([#54](https://github.com/ory/keto/issues/54))
  ([f098639](https://github.com/ory/keto/commit/f098639b5e748151810848fdd3173e0246bc03dc))
- Updates link to guide and header
  ([437c255](https://github.com/ory/keto/commit/437c255ecfff4127fb586cc069e07f86988ad1ba))
- Updates link to open collective
  ([382e1d3](https://github.com/ory/keto/commit/382e1d34c7da0ba0447b78506a749bd7f0085f48))
- Updates links to docs
  ([d84be3b](https://github.com/ory/keto/commit/d84be3b6a8e5eb284ec3fb137ee774ba5ee0d529))
- Updates newsletter link in README
  ([2dc36b2](https://github.com/ory/keto/commit/2dc36b21c8af8e3e39f093198715ea24b65d65af))

### Unclassified

- Add Go SDK factory
  ([99db7e6](https://github.com/ory/keto/commit/99db7e6d4edac88794266a01ddfab9cd0632e95a))
- Add go SDK interface
  ([3dd5f7d](https://github.com/ory/keto/commit/3dd5f7d61bb460c34744b84a34755bfb8219b304))
- Add health handlers
  ([bddb949](https://github.com/ory/keto/commit/bddb949459d05002b0f8882d981e4f63fdddf25f))
- Add policy list handler
  ([a290619](https://github.com/ory/keto/commit/a290619d01d15eb8e3b4e33ede1058d316ee807a))
- Add role iterator in list handler
  ([a3eb696](https://github.com/ory/keto/commit/a3eb6961783f7b562f0a0d0a7e2819bffebce5b8))
- Add SDK generation to circle ci
  ([9b37165](https://github.com/ory/keto/commit/9b37165873bcb0cc5dc60d2514d9824a073466a1))
- Adds ability to update a role using PUT
  ([#14](https://github.com/ory/keto/issues/14))
  ([97ccbe6](https://github.com/ory/keto/commit/97ccbe6d808823c56901ad237878aa6d53cddeeb)):

  - transfer UpdateRoleMembers from https://github.com/ory/hydra/pull/768 to
    keto

  - fix tests by using right http method & correcting sql request

  - Change behavior to overwrite the whole role instead of just the members.

  * small sql migration fix

- Adds log message when telemetry is active
  ([f623c52](https://github.com/ory/keto/commit/f623c52655ff85b7f7209eb73e94eb66a297c5b7))
- Clean up vendor dependencies
  ([9a33c23](https://github.com/ory/keto/commit/9a33c23f4d37ab88b4d643fd79204334d73404c6))
- Do not split empty scope ([#45](https://github.com/ory/keto/issues/45))
  ([b29cf8c](https://github.com/ory/keto/commit/b29cf8cc92607e13457dba8331f5c9286054c8c1))
- Fix typo in help command in env var name
  ([#39](https://github.com/ory/keto/issues/39))
  ([8a5016c](https://github.com/ory/keto/commit/8a5016cd75be78bb42a9a38bfd453ad5722db9db)),
  closes [#25](https://github.com/ory/keto/issues/25)
- Fixes environment variable typos
  ([566d588](https://github.com/ory/keto/commit/566d588e4fca12399966718b725fe4461a28e51e))
- Fixes typo in help command
  ([74450da](https://github.com/ory/keto/commit/74450da18a27513820328c28f72203653c664367)),
  closes [#25](https://github.com/ory/keto/issues/25)
- Format code
  ([637c78c](https://github.com/ory/keto/commit/637c78cba697682b544473a3af9b6ae7715561aa))
- Gofmt
  ([a8d7f9f](https://github.com/ory/keto/commit/a8d7f9f546ae2f3b8c3fa643d8e19b68ca26cc67))
- Improve compose documentation
  ([6870443](https://github.com/ory/keto/commit/68704435f3c299b853f4ff5cacae285b09ada3b5))
- Improves usage of metrics middleware
  ([726c4be](https://github.com/ory/keto/commit/726c4bedfc3f02fdac380930e32f37c251e51aa4))
- Improves usage of metrics middleware
  ([301f386](https://github.com/ory/keto/commit/301f38605af66abae4d28ed0cac90d0b82b655c4))
- Introduce docker-compose file for testing
  ([ba857e3](https://github.com/ory/keto/commit/ba857e3859966e857c5a741825411575e17446de))
- Introduces health and version endpoints
  ([6a9da74](https://github.com/ory/keto/commit/6a9da74f693ee6c15a775ab8d652582aea093601))
- List roles from keto_role table ([#28](https://github.com/ory/keto/issues/28))
  ([9e16605](https://github.com/ory/keto/commit/9e166054b8d474fbce6983d5d00eeeb062fc79b1))
- Properly names flags
  ([af2c3b5](https://github.com/ory/keto/commit/af2c3b5bc96e95fb31b1db5c7fe6dfd6b6fc5b20))
- Require explicit CORS enabling ([#42](https://github.com/ory/keto/issues/42))
  ([9a45107](https://github.com/ory/keto/commit/9a45107af304b2a8e663a532e4f6e4536f15888c))
- Update dependencies
  ([663d8b1](https://github.com/ory/keto/commit/663d8b13e99694a57752cd60a68342b81b041c66))
- Switch to rego as policy decision engine (#48)
  ([ee9bcf2](https://github.com/ory/keto/commit/ee9bcf2719178e5a8dccca083a90313947a8a63b)),
  closes [#48](https://github.com/ory/keto/issues/48)
- Update hydra to v1.0.0-beta.6 ([#35](https://github.com/ory/keto/issues/35))
  ([5e97104](https://github.com/ory/keto/commit/5e971042afff06e2a6ee3b54d2fea31687203623))
- Update npm package registry
  ([a53d3d2](https://github.com/ory/keto/commit/a53d3d23e11fde5dcfbb27a2add1049f4d8e10e6))
- Enable TLS option to serve API (#46)
  ([2f62063](https://github.com/ory/keto/commit/2f620632d0375bf9c7e58dbfb49627c02c66abf3)),
  closes [#46](https://github.com/ory/keto/issues/46)
- Make introspection authorization optional
  ([e5460ad](https://github.com/ory/keto/commit/e5460ad884cd018cd6177324b949cd66bfd53bc7))
- Properly output telemetry information
  ([#33](https://github.com/ory/keto/issues/33))
  ([9994967](https://github.com/ory/keto/commit/9994967b0ca54a62b8b0088fe02be9e890d9574b))
- Remove ORY Hydra dependency ([#44](https://github.com/ory/keto/issues/44))
  ([d487344](https://github.com/ory/keto/commit/d487344fe7e07cb6370371c6b0b6cf3cca767ed1))
- Resolves an issue with the hydra migrate command
  ([2b280bb](https://github.com/ory/keto/commit/2b280bb57c9073a9c8384cde0b14a6991cfacdb6)),
  closes [#23](https://github.com/ory/keto/issues/23)
- Upgrade superagent version ([#41](https://github.com/ory/keto/issues/41))
  ([9c80dbc](https://github.com/ory/keto/commit/9c80dbcc1cc63243839b58ca56ac9be104797887))
- gofmt
  ([777b1be](https://github.com/ory/keto/commit/777b1be1378d314e7cfde0c34450afcce7e590a5))
- Updates README.md (#34)
  ([c28b521](https://github.com/ory/keto/commit/c28b5219fd64314a75ee3c848a80a0c5974ebb7d)),
  closes [#34](https://github.com/ory/keto/issues/34)
- Properly parses cors options
  ([edb5a60](https://github.com/ory/keto/commit/edb5a600f2ce16c0847ee5ef399fa5a41b1e736a))
- Removes additional output if no args are passed
  ([703e124](https://github.com/ory/keto/commit/703e1246ce0fd89066b497c45f0c6cadeb06c331))
- Resolves broken role test
  ([b6c7f9c](https://github.com/ory/keto/commit/b6c7f9c33c4c1f43164d6da0ec7f2553f1f4c598))
- Resolves minor typos and updates install guide
  ([3852be5](https://github.com/ory/keto/commit/3852be56cb81df966a85d4c828de0397d9e74768))
- Updates to latest sqlcon
  ([2c9f643](https://github.com/ory/keto/commit/2c9f643042ff4edffae8bd41834d2a57c923871c))
- Use roles in warden decision
  ([c785187](https://github.com/ory/keto/commit/c785187e31fc7a4b8b762a5e27fac66dcaa97513)),
  closes [#21](https://github.com/ory/keto/issues/21)
  [#19](https://github.com/ory/keto/issues/19)
- authn/client: Payload is now prefixed with client
  ([8584d94](https://github.com/ory/keto/commit/8584d94cfb18deb37ae32ae601f4cd15c14067e7))

# [0.0.1](https://github.com/ory/keto/compare/4f00bc96ece3180a888718ec3c41c69106c86f56...v0.0.1) (2018-05-20)

authn: Checks token_type is "access_token", if set

Closes #1

### Documentation

- Incorporates changes from version
  ([b5445a0](https://github.com/ory/keto/commit/b5445a0fc5b6f813cd1731b20c8c5c79d7c4cdf8))
- Incorporates changes from version
  ([295ff99](https://github.com/ory/keto/commit/295ff998af55777823b04f423e365fd58e61753b))
- Incorporates changes from version
  ([bd44d41](https://github.com/ory/keto/commit/bd44d41b2781e33353082397c47390a27f749e16))
- Updates readme and upgrades
  ([0f95dbb](https://github.com/ory/keto/commit/0f95dbb967fd17b607caa999ae30453f5f599739))
- Uses keto repo for changelog
  ([14c0b2a](https://github.com/ory/keto/commit/14c0b2a2bd31566f2b9048831f894aba05c5b15d))

### Unclassified

- Adds migrate commands to the proper parent command
  ([231c70d](https://github.com/ory/keto/commit/231c70d816b0736a51eddc1fa0445bac672b1b2f))
- Checks token_type is "access_token", if set
  ([d2b8f5d](https://github.com/ory/keto/commit/d2b8f5d313cce597566bd18e4f3bea4a423a62ee)),
  closes [#1](https://github.com/ory/keto/issues/1)
- Removes old test
  ([07b733b](https://github.com/ory/keto/commit/07b733bfae4b733e3e2124545b92c537dabbdcf0))
- Renames subject to sub in response payloads
  ([ca4d540](https://github.com/ory/keto/commit/ca4d5408000be2b896d38eaaf5e67a3fc0a566da))
- Tells linguist to ignore SDK files
  ([f201eb9](https://github.com/ory/keto/commit/f201eb95f3309a60ac50f42cfba0bae2e38e8d13))
- Retries SQL connection on migrate commands
  ([3d33d73](https://github.com/ory/keto/commit/3d33d73c009077c5bf30ae4b03802904bfb5d5b2)):

  This patch also introduces a fatal error if migrations fail

- cmd/server: Resolves DBAL not handling postgres properly
  ([dedc32a](https://github.com/ory/keto/commit/dedc32ab218923243b1955ce5bcbbdc5cc416953))
- cmd/server: Improves error message in migrate command
  ([4b17ce8](https://github.com/ory/keto/commit/4b17ce8848113cae807840182d1a318190c2a9b3))
- Resolves travis and docker issues
  ([6f4779c](https://github.com/ory/keto/commit/6f4779cc51bf4f2ee5b97541fb77d8f882497710))
- Adds OAuth2 Client Credentials authenticator and warden endpoint
  ([c55139b](https://github.com/ory/keto/commit/c55139b51e636834759706499a2aec1451f4fbd9))
- Adds SDK helpers
  ([a1c2608](https://github.com/ory/keto/commit/a1c260801d9366fccf4bfb4fc64b2c67fc594565))
- Resolves SDK and test issues (#4)
  ([2d4cd98](https://github.com/ory/keto/commit/2d4cd9805af3081bbcbea3f806ca066d35385a4b)),
  closes [#4](https://github.com/ory/keto/issues/4)
- Initial project commit
  ([a592e51](https://github.com/ory/keto/commit/a592e5126f130f8b673fff6c894fdbd9fb56f81c))
- Initial commit
  ([4f00bc9](https://github.com/ory/keto/commit/4f00bc96ece3180a888718ec3c41c69106c86f56))

---

id: changelog title: Changelog custom_edit_url: null

---

# [Unreleased](https://github.com/ory/keto/compare/v0.6.0-alpha.3...3bcd0e34f2270401a0b1c24b67cf2df5330584aa) (2021-06-24)

### Bug Fixes

- Add missing tracers ([#600](https://github.com/ory/keto/issues/600))
  ([aa263be](https://github.com/ory/keto/commit/aa263be9a7830e3c769d7698d36137555ca230bc)),
  closes [#593](https://github.com/ory/keto/issues/593)
- Handle relation tuple cycles in expand and check engine
  ([#623](https://github.com/ory/keto/issues/623))
  ([8e30119](https://github.com/ory/keto/commit/8e301198298858fd7f387ef63a7abf4fa55ea240))
- Log all database connection errors
  ([#588](https://github.com/ory/keto/issues/588))
  ([2b0fad8](https://github.com/ory/keto/commit/2b0fad897e61400bd2a6cdf47f33ff4301e9c5f8))
- Move gRPC client module root up
  ([#620](https://github.com/ory/keto/issues/620))
  ([3b881f6](https://github.com/ory/keto/commit/3b881f6015a93b382b3fbbca4be9259622038b6a)):

  BREAKING: The npm package `@ory/keto-grpc-client` from now on includes all API
  versions. Because of that, the import paths changed. For migrating to the new
  client package, change the import path according to the following example:

  ```diff
  - import acl from '@ory/keto-grpc-client/acl_pb.js'
  + // from the latest version
  + import { acl } from '@ory/keto-grpc-client'
  + // or a specific one
  + import acl from '@ory/keto-grpc-client/ory/keto/acl/v1alpha1/acl_pb.js'
  ```

- Update docker-compose.yml version
  ([#595](https://github.com/ory/keto/issues/595))
  ([7fa4dca](https://github.com/ory/keto/commit/7fa4dca4182a1fa024f9cef0a04163f2cbd882aa)),
  closes [#549](https://github.com/ory/keto/issues/549)

### Documentation

- Fix example not following best practice
  ([#582](https://github.com/ory/keto/issues/582))
  ([a015818](https://github.com/ory/keto/commit/a0158182c5f87cfd4767824e1c5d6cbb8094a4e6))
- Update NPM links due to organisation move
  ([#616](https://github.com/ory/keto/issues/616))
  ([6355bea](https://github.com/ory/keto/commit/6355beae5b5b28c3eee19fdee85b9875cbc165c3))

### Features

- Make generated gRPC client its own module
  ([#583](https://github.com/ory/keto/issues/583))
  ([f0fbb64](https://github.com/ory/keto/commit/f0fbb64b3358e9800854295cebc9ec8b8e56c87a))
- Max_idle_conn_time ([#605](https://github.com/ory/keto/issues/605))
  ([50a8623](https://github.com/ory/keto/commit/50a862338e17f86900ca162da7f3467f55f9f954)),
  closes [#523](https://github.com/ory/keto/issues/523)

### Tests

- De-flake status command test ([#629](https://github.com/ory/keto/issues/629))
  ([3bcd0e3](https://github.com/ory/keto/commit/3bcd0e34f2270401a0b1c24b67cf2df5330584aa)):

  Confirmed that the fix works because

  ```
  $ go test -tags sqlite -run TestStatusCmd/server_type=read/case=block -count 1000 ./cmd/status
  ```

  passed.

# [0.6.0-alpha.3](https://github.com/ory/keto/compare/v0.6.0-alpha.2...v0.6.0-alpha.3) (2021-04-29)

Resolves CRDB and build issues.

### Code Generation

- Pin v0.6.0-alpha.3 release commit
  ([d766968](https://github.com/ory/keto/commit/d766968419d10a68fd843df45316e3436b68d61d))

# [0.6.0-alpha.2](https://github.com/ory/keto/compare/v0.6.0-alpha.1...v0.6.0-alpha.2) (2021-04-29)

This release improves stability and documentation.

### Bug Fixes

- Add npm run format to make format
  ([7d844a8](https://github.com/ory/keto/commit/7d844a8e6412ae561963b97ac26d4682411095d4))
- Makefile target
  ([0e6f612](https://github.com/ory/keto/commit/0e6f6122de7bdbb691ad7cc236b6bc9a3601d39e))
- Move swagger to spec dir
  ([7f6a061](https://github.com/ory/keto/commit/7f6a061aafda275d278bf60f16e90039da45bc57))
- Resolve clidoc issues
  ([ef12b4e](https://github.com/ory/keto/commit/ef12b4e267f34fbf9709fe26023f9b7ae6670c24))
- Update install.sh ([#568](https://github.com/ory/keto/issues/568))
  ([86ab245](https://github.com/ory/keto/commit/86ab24531d608df0b5391ee8ec739291b9a90e20))
- Use correct id
  ([5e02902](https://github.com/ory/keto/commit/5e029020b5ba3931f15d343cf6a9762b064ffd45))
- Use correct id for api
  ([32a6b04](https://github.com/ory/keto/commit/32a6b04609054cba84f7b56ebbe92341ec5dcd98))
- Use sqlite image versions ([#544](https://github.com/ory/keto/issues/544))
  ([ec6cc5e](https://github.com/ory/keto/commit/ec6cc5ed528f1a097ea02669d059e060b7eff824))

### Code Generation

- Pin v0.6.0-alpha.2 release commit
  ([470b2c6](https://github.com/ory/keto/commit/470b2c61c649fe5fcf638c84d4418212ff0330a5))

### Documentation

- Add gRPC client README.md ([#559](https://github.com/ory/keto/issues/559))
  ([9dc3596](https://github.com/ory/keto/commit/9dc35969ada8b0d4d73dee9089c4dc61cd9ea657))
- Change forum to discussions readme
  ([#539](https://github.com/ory/keto/issues/539))
  ([ea2999d](https://github.com/ory/keto/commit/ea2999d4963316810a8d8634fcd123bda31eaa8f))
- Fix cat videos example docker compose
  ([#549](https://github.com/ory/keto/issues/549))
  ([b25a711](https://github.com/ory/keto/commit/b25a7114631957935c71ac6a020ab6bd0c244cd7))
- Fix typo ([#538](https://github.com/ory/keto/issues/538))
  ([99a9693](https://github.com/ory/keto/commit/99a969373497792bb4cd8ff62bf5245087517737))
- Include namespace in olymp library example
  ([#540](https://github.com/ory/keto/issues/540))
  ([135e814](https://github.com/ory/keto/commit/135e8145c383a76b494b469253c949c38f4414a7))
- Update install from source steps to actually work
  ([#548](https://github.com/ory/keto/issues/548))
  ([e662256](https://github.com/ory/keto/commit/e6622564f58b7612b13b11b54e75a7350f52d6de))

### Features

- Global docs sidebar and added cloud pages
  ([c631c82](https://github.com/ory/keto/commit/c631c82b7ff3d12734869ac22730b52e73dcf287))
- Support retryable CRDB transactions
  ([833147d](https://github.com/ory/keto/commit/833147dae40e9ac5bdf220f8aa3f01abd444f791))

# [0.6.0-alpha.1](https://github.com/ory/keto/compare/v0.5.6-alpha.1...v0.6.0-alpha.1) (2021-04-07)

We are extremely happy to announce next-gen Ory Keto which implements
[Zanzibar: Google’s Consistent, Global Authorization System](https://research.google/pubs/pub48190/):

> Zanzibar provides a uniform data model and configuration language for
> expressing a wide range of access control policies from hundreds of client
> services at Google, including Calendar, Cloud, Drive, Maps, Photos, and
> YouTube. Its authorization decisions respect causal ordering of user actions
> and thus provide external consistency amid changes to access control lists and
> object contents. Zanzibar scales to trillions of access control lists and
> millions of authorization requests per second to support services used by
> billions of people. It has maintained 95th-percentile latency of less than 10
> milliseconds and availability of greater than 99.999% over 3 years of
> production use.

Ory Keto is the first open source planet-scale authorization system built with
cloud native technologies (Go, gRPC, newSQL) and architecture. It is also the
first open source implementation of Google Zanzibar :tada:!

Many concepts developer by Google Zanzibar are implemented in Ory Keto already.
Let's take a look!

As of this release, Ory Keto knows how to interpret and operate on the basic
access control lists known as relation tuples. They encode relations between
objects and subjects. One simple example of such a relation tuple could encode
"`user1` has access to file `/foo`", a more complex one could encode "everyone
who has write access on `/foo` has read access on `/foo`".

Ory Keto comes with all the basic APIs as described in the Zanzibar paper. All
of them are available over gRPC and REST.

1. List: query relation tuples
2. Check: determine whether a subject has a relation on an object
3. Expand: get a tree of all subjects who have a relation on an object
4. Change: create, update, and delete relation tuples

For all details, head over to the
[documentation](https://www.ory.sh/keto/docs/concepts/api-overview).

With this release we officially move the "old" Keto to the
[legacy-0.5 branch](https://github.com/ory/keto/tree/legacy-0.5). We will only
provide security fixes from now on. A migration path to v0.6 is planned but not
yet implemented, as the architectures are vastly different. Please refer to
[the issue](https://github.com/ory/keto/issues/318).

We are keen to bring more features and performance improvements. The next
features we will tackle are:

- Subject Set rewrites
- Native ABAC & RBAC Support
- Integration with other policy servers
- Latency reduction through aggressive caching
- Cluster mode that fans out requests over all Keto instances

So stay tuned, :star: this repo, :eyes: releases, and
[subscribe to our newsletter :email:](https://ory.us10.list-manage.com/subscribe?u=ffb1a878e4ec6c0ed312a3480&id=f605a41b53&MERGE0=&group[17097][32]=1).

### Bug Fixes

- Add description attribute to access control policy role
  ([#215](https://github.com/ory/keto/issues/215))
  ([831eba5](https://github.com/ory/keto/commit/831eba59f810ca68561dd584c9df7684df10b843))
- Add leak_sensitive_values to config schema
  ([2b21d2b](https://github.com/ory/keto/commit/2b21d2bdf4ca9523d16159c5f73c4429b692e17d))
- Bump CLI
  ([80c82d0](https://github.com/ory/keto/commit/80c82d026cbfbab8fbb84d850d8980866ecf88df))
- Bump deps and replace swagutil
  ([#212](https://github.com/ory/keto/issues/212))
  ([904258d](https://github.com/ory/keto/commit/904258d23959c3fa96b6d8ccfdb79f6788c106ec))
- Check engine overwrote result in some cases
  ([#412](https://github.com/ory/keto/issues/412))
  ([3404492](https://github.com/ory/keto/commit/3404492002ca5c3f017ef25486e377e911987aa4))
- Check health status in status command
  ([21c64d4](https://github.com/ory/keto/commit/21c64d45f21a505744b9f70d780f9b3079d3822c))
- Check REST API returns JSON object
  ([#460](https://github.com/ory/keto/issues/460))
  ([501dcff](https://github.com/ory/keto/commit/501dcff4427f76902671f6d5733f28722bd51fa7)),
  closes [#406](https://github.com/ory/keto/issues/406)
- Empty relationtuple list should not error
  ([#440](https://github.com/ory/keto/issues/440))
  ([fbcb3e1](https://github.com/ory/keto/commit/fbcb3e1f337b5114d7697fa512ded92b5f409ef4))
- Ensure nil subject is not allowed
  ([#449](https://github.com/ory/keto/issues/449))
  ([7a0fcfc](https://github.com/ory/keto/commit/7a0fcfc4fe83776fa09cf78ee11f407610554d04)):

  The nodejs gRPC client was a great fuzzer and pointed me to some nil pointer
  dereference panics. This adds some input validation to prevent panics.

- Ensure persister errors are handled by sqlcon
  ([#473](https://github.com/ory/keto/issues/473))
  ([4343c4a](https://github.com/ory/keto/commit/4343c4acd8f917fb7ae131e67bca6855d4d61694))
- Handle pagination and errors in the check/expand engines
  ([#398](https://github.com/ory/keto/issues/398))
  ([5eb1a7d](https://github.com/ory/keto/commit/5eb1a7d49af6b43707c122de8727cbd72285cb5c))
- Ignore dist
  ([ba816ea](https://github.com/ory/keto/commit/ba816ea2ca39962f02c08e0c7b75cfe3cf1d963d))
- Ignore x/net false positives
  ([d8b36cb](https://github.com/ory/keto/commit/d8b36cb1812abf7265ac15c29780222be025186b))
- Improve CLI remote sourcing ([#474](https://github.com/ory/keto/issues/474))
  ([a85f4d7](https://github.com/ory/keto/commit/a85f4d7470ac3744476e82e5889b97d5a0680473))
- Improve handlers and add tests
  ([#470](https://github.com/ory/keto/issues/470))
  ([ca5ccb9](https://github.com/ory/keto/commit/ca5ccb9c237fdcc4db031ec97a75616a859cbf8f))
- Insert relation tuples without fmt.Sprintf
  ([#443](https://github.com/ory/keto/issues/443))
  ([fe507bb](https://github.com/ory/keto/commit/fe507bb4ea719780e732d098291aa190d6b1c441))
- Minor bugfixes ([#371](https://github.com/ory/keto/issues/371))
  ([185ee1e](https://github.com/ory/keto/commit/185ee1e51bc4bcdee028f71fcaf207b7e342313b))
- Move dockerfile to where it belongs
  ([f087843](https://github.com/ory/keto/commit/f087843ac8f24e741bf39fe65ee5b0a9adf9a5bb))
- Namespace migrator ([#417](https://github.com/ory/keto/issues/417))
  ([ea79300](https://github.com/ory/keto/commit/ea7930064f490b063a712b4e18521f8996931a13)),
  closes [#404](https://github.com/ory/keto/issues/404)
- Remove SQL logging ([#455](https://github.com/ory/keto/issues/455))
  ([d8e2a86](https://github.com/ory/keto/commit/d8e2a869db2a9cfb44423b434330536036b2f421))
- Rename /relationtuple endpoint to /relation-tuples
  ([#519](https://github.com/ory/keto/issues/519))
  ([8eb55f6](https://github.com/ory/keto/commit/8eb55f6269399f2bc5f000b8a768bcdf356c756f))
- Resolve gitignore build
  ([6f04bbb](https://github.com/ory/keto/commit/6f04bbb6057779b4d73d3f94677cea365843f7ac))
- Resolve goreleaser issues
  ([d32767f](https://github.com/ory/keto/commit/d32767f32856cf5bd48514c5d61746417fbed6f5))
- Resolve windows build issues
  ([8bcdfbf](https://github.com/ory/keto/commit/8bcdfbfbdb0b10c03ff93838e8fe6e778236e96d))
- Rewrite check engine to search starting at the object
  ([#310](https://github.com/ory/keto/issues/310))
  ([7d99694](https://github.com/ory/keto/commit/7d9969414ebc8cf6ef5d211ad34f8ae01bd3b4ee)),
  closes [#302](https://github.com/ory/keto/issues/302)
- Secure query building ([#442](https://github.com/ory/keto/issues/442))
  ([c7d2770](https://github.com/ory/keto/commit/c7d2770ed570238fd1262bcc4e5b4afa6c12d80e))
- Strict version enforcement in docker
  ([e45b28f](https://github.com/ory/keto/commit/e45b28fec626db35f1bd4580e5b11c9c94a02669))
- Update dd-trace to fix build issues
  ([2ad489f](https://github.com/ory/keto/commit/2ad489f0d9cae3191718d36823fe25df58ab95e6))
- Update docker to go 1.16 and alpine
  ([c63096c](https://github.com/ory/keto/commit/c63096cb53d2171f22f4a0d4a9ac3c9bfac89d01))
- Use errors.WithStack everywhere
  ([#462](https://github.com/ory/keto/issues/462))
  ([5f25bce](https://github.com/ory/keto/commit/5f25bceea35179c67d24dd95f698dc57b789d87a)),
  closes [#437](https://github.com/ory/keto/issues/437):

  Fixed all occurrences found using the search pattern `return .*, err\n`.

- Use package name in pkger
  ([6435939](https://github.com/ory/keto/commit/6435939ad7e5899505cd0e6261f5dfc819c9ca42))
- **schema:** Add trace level to logger
  ([a5a1402](https://github.com/ory/keto/commit/a5a1402c61e1a37b1a9a349ad5736eaca66bd6a4))
- Use make() to initialize slices
  ([#250](https://github.com/ory/keto/issues/250))
  ([84f028d](https://github.com/ory/keto/commit/84f028dc35665174542e103c0aefc635bb6d3e52)),
  closes [#217](https://github.com/ory/keto/issues/217)

### Build System

- Pin dependency versions of buf and protoc plugins
  ([#338](https://github.com/ory/keto/issues/338))
  ([5a2fd1c](https://github.com/ory/keto/commit/5a2fd1cc8dff02aa7017771adc0d9101f6c86775))

### Code Generation

- Pin v0.6.0-alpha.1 release commit
  ([875af25](https://github.com/ory/keto/commit/875af25f89b813455148e58884dcdf1cd3600b86))

### Code Refactoring

- Data structures ([#279](https://github.com/ory/keto/issues/279))
  ([1316077](https://github.com/ory/keto/commit/131607762d0006e4cf4f93e8731ef7648348b2ec))

### Documentation

- Add check- and expand-API guides
  ([#493](https://github.com/ory/keto/issues/493))
  ([09a25b4](https://github.com/ory/keto/commit/09a25b4063abcfdcd4c0de315a2ef088d6d4e72e))
- Add current features overview ([#505](https://github.com/ory/keto/issues/505))
  ([605afa0](https://github.com/ory/keto/commit/605afa029794ad115bba02e004e1596cea038e8e))
- Add missing pages ([#518](https://github.com/ory/keto/issues/518))
  ([43cbaa9](https://github.com/ory/keto/commit/43cbaa9140cfa0ea3c72f699f6bb34f5ed31d8dd))
- Add namespace and relation naming conventions
  ([#510](https://github.com/ory/keto/issues/510))
  ([dd31865](https://github.com/ory/keto/commit/dd318653178cd45da47f3e7cef507b42708363ef))
- Add performance page ([#413](https://github.com/ory/keto/issues/413))
  ([6fe0639](https://github.com/ory/keto/commit/6fe0639d36087b5ecd555eb6fe5ce949f3f6f0d7)):

  This also refactored the server startup. Functionality did not change.

- Add production guide
  ([a9163c7](https://github.com/ory/keto/commit/a9163c7690c55c8191650c4dfb464b75ea02446b))
- Add zanzibar overview to README.md
  ([#265](https://github.com/ory/keto/issues/265))
  ([15a95b2](https://github.com/ory/keto/commit/15a95b28e745592353e4656d42a9d0bd20ce468f))
- API overview ([#501](https://github.com/ory/keto/issues/501))
  ([05fe03b](https://github.com/ory/keto/commit/05fe03b5bf7a3f790aa6c9c1d3fcdb31304ef6af))
- Concepts ([#429](https://github.com/ory/keto/issues/429))
  ([2f2c885](https://github.com/ory/keto/commit/2f2c88527b3f6d1d46a5c287d8aca0874d18a28d))
- Delete old redirect homepage
  ([c0a3784](https://github.com/ory/keto/commit/c0a378448f8c7723bae68f7b52a019b697b25863))
- Document gRPC SKDs
  ([7583fe8](https://github.com/ory/keto/commit/7583fe8933f6676b4e37477098b1d43d12819b8b))
- Fix grammatical error ([#222](https://github.com/ory/keto/issues/222))
  ([256a0d2](https://github.com/ory/keto/commit/256a0d2e53fe1eb859e41fc539870ae1d5a493d2))
- Fix regression issues
  ([9697bb4](https://github.com/ory/keto/commit/9697bb43dd23c0d1fae74ea42e848883c45dae77))
- Generate gRPC reference page ([#488](https://github.com/ory/keto/issues/488))
  ([93ebe6d](https://github.com/ory/keto/commit/93ebe6db7e887d708503a54c5ec943254e37ca43))
- Improve CLI documentation ([#503](https://github.com/ory/keto/issues/503))
  ([be9327f](https://github.com/ory/keto/commit/be9327f7b28152a78f731043acf83b7092e42e29))
- Minor fixes ([#532](https://github.com/ory/keto/issues/532))
  ([638342e](https://github.com/ory/keto/commit/638342eb9519d9bf609926fb87558071e2815fb3))
- Move development section
  ([9ff393f](https://github.com/ory/keto/commit/9ff393f6cba1fb0a33918377ce505455c34d9dfc))
- Move to json sidebar
  ([257bf96](https://github.com/ory/keto/commit/257bf96044df37c3d7af8a289fb67098d48da1a3))
- Remove duplicate "is"
  ([ca3277d](https://github.com/ory/keto/commit/ca3277d82c1508797bc8c663963407d2e4d9112f))
- Remove duplicate template
  ([1d3b38e](https://github.com/ory/keto/commit/1d3b38e4045b0b874bb1186ea628f5a37353a2e6))
- Remove old documentation ([#426](https://github.com/ory/keto/issues/426))
  ([eb76913](https://github.com/ory/keto/commit/eb7691306018678e024211b51627a1c27e780a6b))
- Replace TODO links ([#512](https://github.com/ory/keto/issues/512))
  ([ad8e20b](https://github.com/ory/keto/commit/ad8e20b3bef2bc46b3a32c2c9ccb6e16e4bad22c))
- Resolve broken links
  ([0d0a50b](https://github.com/ory/keto/commit/0d0a50b3f4112893f32c81adc8edd137b5a62541))
- Simple access check guide ([#451](https://github.com/ory/keto/issues/451))
  ([e0485af](https://github.com/ory/keto/commit/e0485afc46a445868580aa541e962e80cbea0670)):

  This also enables gRPC go, gRPC nodejs, cURL, and Keto CLI code samples to be
  tested.

- Update comment in write response
  ([#329](https://github.com/ory/keto/issues/329))
  ([4ca0baf](https://github.com/ory/keto/commit/4ca0baf62e34402e749e870fe8c0cc893684192c))
- Update install instructions
  ([d2e4123](https://github.com/ory/keto/commit/d2e4123f3e2e58da8be181a0a542e3dcc1313e16))
- Update introduction
  ([5f71d73](https://github.com/ory/keto/commit/5f71d73e2ee95d02abc4cd42a76c98a35942df0c))
- Update README ([#515](https://github.com/ory/keto/issues/515))
  ([18d3cd6](https://github.com/ory/keto/commit/18d3cd61b0a79400170dc0f89860b4614cc4a543)):

  Also format all markdown files in the root.

- Update repository templates
  ([db505f9](https://github.com/ory/keto/commit/db505f9e10755bc20c4623c4f5f99f33283dffda))
- Update repository templates
  ([6c056bb](https://github.com/ory/keto/commit/6c056bb2043af6e82f06fdfa509ab3fa0d5e5d06))
- Update SDK links ([#514](https://github.com/ory/keto/issues/514))
  ([f920fbf](https://github.com/ory/keto/commit/f920fbfc8dcc6711ad9e046578a4506179952be7))
- Update swagger documentation for REST endpoints
  ([c363de6](https://github.com/ory/keto/commit/c363de61edf912fef85acc6bcdac6e1c15c48f4f))
- Use mdx for api reference
  ([340f3a3](https://github.com/ory/keto/commit/340f3a3dd20c82c743e7b3ad6aaf06a4c118b5a1))
- Various improvements and updates
  ([#486](https://github.com/ory/keto/issues/486))
  ([a812ace](https://github.com/ory/keto/commit/a812ace2303214e0e7acb2e283efa1cff0d5d279))

### Features

- Add .dockerignore
  ([8b0ff06](https://github.com/ory/keto/commit/8b0ff066b2508ef2f3629f9a3e2fce601b8dcce1))
- Add and automate version schema
  ([b01eef8](https://github.com/ory/keto/commit/b01eef8d4d5834b5888cb369ecf01ee01b40c24c))
- Add check engine ([#277](https://github.com/ory/keto/issues/277))
  ([396c1ae](https://github.com/ory/keto/commit/396c1ae33b777031f8d59549d9de4a88e3f6b10a))
- Add gRPC health status ([#427](https://github.com/ory/keto/issues/427))
  ([51c4223](https://github.com/ory/keto/commit/51c4223d6cb89a9bfbc115ef20db8350aeb2e8af))
- Add is_last_page to list response
  ([#425](https://github.com/ory/keto/issues/425))
  ([b73d91f](https://github.com/ory/keto/commit/b73d91f061ab155c53d802263c0263aa39e64bdf))
- Add POST REST handler for policy check
  ([7d89860](https://github.com/ory/keto/commit/7d89860bc4a790a69f5bea5b0dbe4a2938c6729f))
- Add relation write API ([#275](https://github.com/ory/keto/issues/275))
  ([f2ddb9d](https://github.com/ory/keto/commit/f2ddb9d884ed71037b5371c00bb10b63d25d47c0))
- Add REST and gRPC logger middlewares
  ([#436](https://github.com/ory/keto/issues/436))
  ([615eb0b](https://github.com/ory/keto/commit/615eb0bec3bdc0fd26abc7af0b8990269b0cbedd))
- Add SQA telemetry ([#535](https://github.com/ory/keto/issues/535))
  ([9f6472b](https://github.com/ory/keto/commit/9f6472b0c996505d41058e9b55afa8fd6b9bb2d5))
- Add sql persister ([#350](https://github.com/ory/keto/issues/350))
  ([d595d52](https://github.com/ory/keto/commit/d595d52dabb8f4953b5c23d3a8154cac13d00306))
- Add tracing ([#536](https://github.com/ory/keto/issues/536))
  ([b57a144](https://github.com/ory/keto/commit/b57a144e0a7ec39d5831dbb79840c2b25c044e6a))
- Allow to apply namespace migrations together with regular migrations
  ([#441](https://github.com/ory/keto/issues/441))
  ([57e2bbc](https://github.com/ory/keto/commit/57e2bbc5eaebe43834f2432eb1ee2820d9cb2988))
- Delete relation tuples ([#457](https://github.com/ory/keto/issues/457))
  ([3ec8afa](https://github.com/ory/keto/commit/3ec8afa68c5b5ddc26609b9afd17cc0d06cd82bf)),
  closes [#452](https://github.com/ory/keto/issues/452)
- Dockerfile and docker compose example
  ([#390](https://github.com/ory/keto/issues/390))
  ([10cd0b3](https://github.com/ory/keto/commit/10cd0b39c12ef96710bda6ff013f7c5eeae97118))
- Expand API ([#285](https://github.com/ory/keto/issues/285))
  ([a3ca0b8](https://github.com/ory/keto/commit/a3ca0b8a109b63f443e359cd8ff18a7b3e489f84))
- Expand GPRC service and CLI ([#383](https://github.com/ory/keto/issues/383))
  ([acf2154](https://github.com/ory/keto/commit/acf21546d3e135deb77c853b751a3da3a7b16f00))
- First API draft and generation
  ([#315](https://github.com/ory/keto/issues/315))
  ([bda5d8b](https://github.com/ory/keto/commit/bda5d8b7e90d749600f5b5e169df8a6ec3705b22))
- GRPC status codes and improved error messages
  ([#467](https://github.com/ory/keto/issues/467))
  ([4a4f8c6](https://github.com/ory/keto/commit/4a4f8c6b323664329414b61e7d80d7838face730))
- GRPC version API ([#475](https://github.com/ory/keto/issues/475))
  ([89cc46f](https://github.com/ory/keto/commit/89cc46fe4a13b062693d3db4f803834ba37f4e48))
- Implement goreleaser pipeline
  ([888ac43](https://github.com/ory/keto/commit/888ac43e6f706f619b2f1b58271dd027094c9ae9)),
  closes [#410](https://github.com/ory/keto/issues/410)
- Incorporate new GRPC API structure
  ([#331](https://github.com/ory/keto/issues/331))
  ([e0916ad](https://github.com/ory/keto/commit/e0916ad00c81b24177cfe45faf77b93d2c33dc1f))
- Koanf and namespace configuration
  ([#367](https://github.com/ory/keto/issues/367))
  ([3ad32bc](https://github.com/ory/keto/commit/3ad32bc13a4d96135be8031eb6fe4c15868272ca))
- Namespace configuration ([#324](https://github.com/ory/keto/issues/324))
  ([b94f50d](https://github.com/ory/keto/commit/b94f50d1800c47a43561df5009cb38b44ccd0088))
- Namespace migrate status CLI ([#508](https://github.com/ory/keto/issues/508))
  ([e3f7ad9](https://github.com/ory/keto/commit/e3f7ad91585b616e97f85ce0f55c76406b6c4d0a)):

  This also refactors the current `migrate` and `namespace migrate` commands.

- Nodejs gRPC definitions ([#447](https://github.com/ory/keto/issues/447))
  ([3b5c313](https://github.com/ory/keto/commit/3b5c31326645adb2d5b14ced901771a7ba00fd1c)):

  Includes Typescript definitions.

- Read API ([#269](https://github.com/ory/keto/issues/269))
  ([de5119a](https://github.com/ory/keto/commit/de5119a6e3c7563cfc2e1ada12d47b27ebd7faaa)):

  This is a first draft of the read API. It is reachable by REST and gRPC calls.
  The main purpose of this PR is to establish the basic repository structure and
  define the API.

- Relationtuple parse command ([#490](https://github.com/ory/keto/issues/490))
  ([91a3cf4](https://github.com/ory/keto/commit/91a3cf47fbdb8203b799cf7c69bcf3dbbfb98b3a)):

  This command parses the relation tuple format used in the docs. It greatly
  improves the experience when copying something from the documentation. It can
  especially be used to pipe relation tuples into other commands, e.g.:

  ```shell
  echo "messages:02y_15_4w350m3#decypher@john" | \
    keto relation-tuple parse - --format json | \
    keto relation-tuple create -
  ```

- REST patch relation tuples ([#491](https://github.com/ory/keto/issues/491))
  ([d38618a](https://github.com/ory/keto/commit/d38618a9e647902ce019396ff1c33973020bf797)):

  The new PATCH handler allows transactional changes similar to the already
  existing gRPC service.

- Separate and multiplex ports based on read/write privilege
  ([#397](https://github.com/ory/keto/issues/397))
  ([6918ac3](https://github.com/ory/keto/commit/6918ac3bfa355cbd551e44376c214f412e3414e4))
- Swagger SDK ([#476](https://github.com/ory/keto/issues/476))
  ([011888c](https://github.com/ory/keto/commit/011888c2b7e2d0f7b8923c994c70e62d374a2830))

### Tests

- Add command tests ([#487](https://github.com/ory/keto/issues/487))
  ([61c28e4](https://github.com/ory/keto/commit/61c28e48a5c3f623e5cc133e69ba368c5103f414))
- Add dedicated persistence tests
  ([#416](https://github.com/ory/keto/issues/416))
  ([4e98906](https://github.com/ory/keto/commit/4e9890605edf3ea26134917a95bfa6fbb176565e))
- Add handler tests ([#478](https://github.com/ory/keto/issues/478))
  ([9315a77](https://github.com/ory/keto/commit/9315a77820d50e400b78f2f019a871be022a9887))
- Add initial e2e test ([#380](https://github.com/ory/keto/issues/380))
  ([dc5d3c9](https://github.com/ory/keto/commit/dc5d3c9d02604fddbfa56ac5ebbc1fef56a881d9))
- Add relationtuple definition tests
  ([#415](https://github.com/ory/keto/issues/415))
  ([2e3dcb2](https://github.com/ory/keto/commit/2e3dcb200a7769dc8710d311ca08a7515012fbdd))
- Enable GRPC client in e2e test
  ([#382](https://github.com/ory/keto/issues/382))
  ([4e5c6ae](https://github.com/ory/keto/commit/4e5c6aed56e5a449003956ec114ec131be068aaf))
- Improve docs sample tests ([#461](https://github.com/ory/keto/issues/461))
  ([6e0e5e6](https://github.com/ory/keto/commit/6e0e5e6184916e894fd4694cfa3a158f11fae11f))

# [0.5.6-alpha.1](https://github.com/ory/keto/compare/v0.5.5-alpha.1...v0.5.6-alpha.1) (2020-05-28)

This release bumps vulnerable transient dependencies (those are not actually
used in ORY Keto) and updates several documentation pages and improves
structured logging output. Additionally, ORY Keto now uses the updated release
pipeline!

### Bug Fixes

- Update install script
  ([21e1bf0](https://github.com/ory/keto/commit/21e1bf05177576a9d743bd11744ef6a42be50b8d))

### Chores

- Pin v0.5.6-alpha.1 release commit
  ([ed0da08](https://github.com/ory/keto/commit/ed0da08a03a910660358fc56c568692325749b6d))

# [0.5.5-alpha.1](https://github.com/ory/keto/compare/v0.5.4-alpha.1...v0.5.5-alpha.1) (2020-05-28)

This release bumps vulnerable transient dependencies (those are not actually
used in ORY Keto) and updates several documentation pages and improves
structured logging output. Additionally, ORY Keto now uses the updated release
pipeline!

### Bug Fixes

- Move deps to go_mod_indirect_pins
  ([dd3e971](https://github.com/ory/keto/commit/dd3e971ac418baf10c1b33005acc7e6f66fb0d85))
- Resolve test issues
  ([9bd9956](https://github.com/ory/keto/commit/9bd9956e33731f1619c32e1e6b7c78f42e7c47c3))
- Update install.sh script
  ([f64d320](https://github.com/ory/keto/commit/f64d320b6424fe3256eb7fad1c94dcc1ef0bf487))
- Use semver-regex replacer func
  ([2cc3bbb](https://github.com/ory/keto/commit/2cc3bbb2d75ba5fa7a3653d7adcaa712ff38c603))

### Chores

- Pin v0.5.5-alpha.1 release commit
  ([4666a0f](https://github.com/ory/keto/commit/4666a0f258f253d19a14eca34f4b7049f2d0afa2))

### Documentation

- Add missing colon in docker run command
  ([#193](https://github.com/ory/keto/issues/193))
  ([383063d](https://github.com/ory/keto/commit/383063d260d995665da4c02c9a7bac7e06a2c8d3))
- Update github templates ([#182](https://github.com/ory/keto/issues/182))
  ([72ea09b](https://github.com/ory/keto/commit/72ea09bbbf9925d7705842703b32826376f636e4))
- Update github templates ([#184](https://github.com/ory/keto/issues/184))
  ([ed546b7](https://github.com/ory/keto/commit/ed546b7a2b9ee690284a48c641edd1570464d71f))
- Update github templates ([#188](https://github.com/ory/keto/issues/188))
  ([ebd75b2](https://github.com/ory/keto/commit/ebd75b2f6545ff4372773f6370300c7b2ca71c51))
- Update github templates ([#189](https://github.com/ory/keto/issues/189))
  ([fd4c0b1](https://github.com/ory/keto/commit/fd4c0b17bcb1c281baac1772ab94e305ec8c5c86))
- Update github templates ([#195](https://github.com/ory/keto/issues/195))
  ([ba0943c](https://github.com/ory/keto/commit/ba0943c45d36ef10bdf1169f0aeef439a3a67d28))
- Update linux install guide ([#191](https://github.com/ory/keto/issues/191))
  ([7d8b24b](https://github.com/ory/keto/commit/7d8b24bddb9c92feb78c7b65f39434d538773b58))
- Update repository templates
  ([ea65b5c](https://github.com/ory/keto/commit/ea65b5c5ada0a7453326fa755aa914306f1b1851))
- Use central banner repo for README
  ([0d95d97](https://github.com/ory/keto/commit/0d95d97504df4d0ab57d18dc6d0a824a3f8f5896))
- Use correct banner
  ([c6dfe28](https://github.com/ory/keto/commit/c6dfe280fd962169c424834cea040a408c1bc83f))
- Use correct version
  ([5f7030c](https://github.com/ory/keto/commit/5f7030c9069fe392200be72f8ce1a93890fbbba8)),
  closes [#200](https://github.com/ory/keto/issues/200)
- Use correct versions in install docs
  ([52e6c34](https://github.com/ory/keto/commit/52e6c34780ed41c169504d71c39459898b5d14f9))

# [0.5.4-alpha.1](https://github.com/ory/keto/compare/v0.5.3-alpha.3...v0.5.4-alpha.1) (2020-04-07)

fix: resolve panic when executing migrations (#178)

Closes #177

### Bug Fixes

- Resolve panic when executing migrations
  ([#178](https://github.com/ory/keto/issues/178))
  ([7e83fee](https://github.com/ory/keto/commit/7e83feefaad041c60f09232ac44ed8b7240c6558)),
  closes [#177](https://github.com/ory/keto/issues/177)

# [0.5.3-alpha.3](https://github.com/ory/keto/compare/v0.5.3-alpha.2...v0.5.3-alpha.3) (2020-04-06)

autogen(docs): regenerate and update changelog

### Code Generation

- **docs:** Regenerate and update changelog
  ([769cef9](https://github.com/ory/keto/commit/769cef90f27ba9c203d3faf47272287ab17dc7eb))

### Code Refactoring

- Move docs to this repository ([#172](https://github.com/ory/keto/issues/172))
  ([312480d](https://github.com/ory/keto/commit/312480de3cefc5b72916ba95d8287443cf3ccb3d))

### Documentation

- Regenerate and update changelog
  ([dda79b1](https://github.com/ory/keto/commit/dda79b106a18bc33d70ae60e352118b0d288d26b))
- Regenerate and update changelog
  ([9048dd8](https://github.com/ory/keto/commit/9048dd8d8a0f0654072b3d4b77261fe947a34ece))
- Regenerate and update changelog
  ([806f68c](https://github.com/ory/keto/commit/806f68c603781742e0177ec0b2deecaf64c5b721))
- Regenerate and update changelog
  ([8905ee7](https://github.com/ory/keto/commit/8905ee74d4ec394af92240e180cc5d7f6493cb2f))
- Regenerate and update changelog
  ([203c1cc](https://github.com/ory/keto/commit/203c1cc659a72f81a370d7b9b7fbda60e7c96c9e))
- Regenerate and update changelog
  ([8875a95](https://github.com/ory/keto/commit/8875a95b35df57668acb27820a3aff1cdfbe8b30))
- Regenerate and update changelog
  ([28ddd3e](https://github.com/ory/keto/commit/28ddd3e1483afe8571b3d2bf9afcc31386d85f7f))
- Regenerate and update changelog
  ([927c4ed](https://github.com/ory/keto/commit/927c4edc4a770133bcb34bc044dd5c5e0eb3ffb7))
- Updates issue and pull request templates
  ([#168](https://github.com/ory/keto/issues/168))
  ([29a38a8](https://github.com/ory/keto/commit/29a38a85c61ec2c8d0ad2ce6d5a0f9e9d74b52f7))
- Updates issue and pull request templates
  ([#169](https://github.com/ory/keto/issues/169))
  ([99b7d5d](https://github.com/ory/keto/commit/99b7d5de24fed1aed746c4447a390d084632f89a))
- Updates issue and pull request templates
  ([#171](https://github.com/ory/keto/issues/171))
  ([7a9876b](https://github.com/ory/keto/commit/7a9876b8ed4282f50f886a025033641bd027a0e2))

# [0.5.3-alpha.1](https://github.com/ory/keto/compare/v0.5.2...v0.5.3-alpha.1) (2020-04-03)

chore: move to ory analytics fork (#167)

### Chores

- Move to ory analytics fork ([#167](https://github.com/ory/keto/issues/167))
  ([f824011](https://github.com/ory/keto/commit/f824011b4d19058504b3a43ed53a420619444a51))

# [0.5.2](https://github.com/ory/keto/compare/v0.5.1-alpha.1...v0.5.2) (2020-04-02)

docs: Regenerate and update changelog

### Documentation

- Regenerate and update changelog
  ([1e52100](https://github.com/ory/keto/commit/1e521001a43a0a13e2224e1a44956442ac6ffbc7))
- Regenerate and update changelog
  ([e4d32a6](https://github.com/ory/keto/commit/e4d32a62c1ae96115ea50bb471f5ff2ce2f2c4b9))

# [0.5.0](https://github.com/ory/keto/compare/v0.4.5-alpha.1...v0.5.0) (2020-04-02)

docs: use real json bool type in swagger (#162)

Closes #160

### Bug Fixes

- Move to ory sqa service ([#159](https://github.com/ory/keto/issues/159))
  ([c3bf1b1](https://github.com/ory/keto/commit/c3bf1b1964a14be4cc296aae98d0739e65917e18))
- Use correct response mode for removeOryAccessControlPolicyRoleMe…
  ([#161](https://github.com/ory/keto/issues/161))
  ([17543cf](https://github.com/ory/keto/commit/17543cfef63a1d040a2234bd63b210fb9c4f6015))

### Documentation

- Regenerate and update changelog
  ([6a77f75](https://github.com/ory/keto/commit/6a77f75d66e89420f2daf2fae011d31bcfa34008))
- Regenerate and update changelog
  ([c8c9d29](https://github.com/ory/keto/commit/c8c9d29e77ef53e1196cc6fe600c53d93376229b))
- Regenerate and update changelog
  ([fe8327d](https://github.com/ory/keto/commit/fe8327d951394084df7785166c9a9578c1ab0643))
- Regenerate and update changelog
  ([b5b1d66](https://github.com/ory/keto/commit/b5b1d66a4b933df8789337cce3f6d6bf391b617b))
- Update forum and chat links
  ([e96d7ba](https://github.com/ory/keto/commit/e96d7ba3dcc693c22eb983b3f58a05c9c6adbda7))
- Updates issue and pull request templates
  ([#158](https://github.com/ory/keto/issues/158))
  ([ab14cfa](https://github.com/ory/keto/commit/ab14cfa51ce195b26a83c050452530a5008589d7))
- Use real json bool type in swagger
  ([#162](https://github.com/ory/keto/issues/162))
  ([5349e7f](https://github.com/ory/keto/commit/5349e7f910ad22558a01b76be62db2136b5eb301)),
  closes [#160](https://github.com/ory/keto/issues/160)

# [0.4.5-alpha.1](https://github.com/ory/keto/compare/v0.4.4-alpha.1...v0.4.5-alpha.1) (2020-02-29)

docs: Regenerate and update changelog

### Bug Fixes

- **driver:** Extract scheme from DSN using sqlcon.GetDriverName
  ([#156](https://github.com/ory/keto/issues/156))
  ([187e289](https://github.com/ory/keto/commit/187e289f1a235b5cacf2a0b7ca5e98c384fa7a14)),
  closes [#145](https://github.com/ory/keto/issues/145)

### Documentation

- Regenerate and update changelog
  ([41513da](https://github.com/ory/keto/commit/41513da35ea038f3c4cc2d98b9796cee5b5a8b92))

# [0.4.4-alpha.1](https://github.com/ory/keto/compare/v0.4.3-alpha.2...v0.4.4-alpha.1) (2020-02-14)

docs: Regenerate and update changelog

### Bug Fixes

- **goreleaser:** Update brew section
  ([0918ff3](https://github.com/ory/keto/commit/0918ff3032eeecd26c67d6249c7e28e71ee110af))

### Documentation

- Prepare ecosystem automation
  ([2e39be7](https://github.com/ory/keto/commit/2e39be79ebad1cec021ae3ee4b0a75ffea4b7424))
- Regenerate and update changelog
  ([009c4c4](https://github.com/ory/keto/commit/009c4c4e4fd4c5607cc30cc9622fd0f82e3891f3))
- Regenerate and update changelog
  ([49f3c4b](https://github.com/ory/keto/commit/49f3c4ba34df5879d8f48cc96bf0df9dad820362))
- Updates issue and pull request templates
  ([#153](https://github.com/ory/keto/issues/153))
  ([7fb7521](https://github.com/ory/keto/commit/7fb752114e1e2a91ab96fdb546835de8aee4926b))

### Features

- **ci:** Add nancy vuln scanner
  ([#152](https://github.com/ory/keto/issues/152))
  ([c19c2b9](https://github.com/ory/keto/commit/c19c2b9efe8299b8878cc8099fe314d8dcda3a08))

### Unclassified

- Update CHANGELOG [ci skip]
  ([63fe513](https://github.com/ory/keto/commit/63fe513d22ec3747a95cdb8f797ba1eba5ca344f))
- Update CHANGELOG [ci skip]
  ([7b7c3ac](https://github.com/ory/keto/commit/7b7c3ac6c06c072fea1b64624ea79a3fd406b09c))
- Update CHANGELOG [ci skip]
  ([8886392](https://github.com/ory/keto/commit/8886392b39fb46ad338c8284866d4dae64ad1826))
- Update CHANGELOG [ci skip]
  ([5bbc284](https://github.com/ory/keto/commit/5bbc2844c49b0a68ba3bd8b003d91f87e2aed9e2))

# [0.4.3-alpha.2](https://github.com/ory/keto/compare/v0.4.3-alpha.1...v0.4.3-alpha.2) (2020-01-31)

Update README.md

### Unclassified

- Update README.md
  ([0ab9c6f](https://github.com/ory/keto/commit/0ab9c6f372a1538a958a68b34315c9167b5a9093))
- Update CHANGELOG [ci skip]
  ([f0a1428](https://github.com/ory/keto/commit/f0a1428f4b99ceb35ff4f1e839bc5237e19db628))

# [0.4.3-alpha.1](https://github.com/ory/keto/compare/v0.4.2-alpha.1...v0.4.3-alpha.1) (2020-01-23)

Disable access logging for health endpoints (#151)

Closes #150

### Unclassified

- Disable access logging for health endpoints (#151)
  ([6ca0c09](https://github.com/ory/keto/commit/6ca0c09b5618122762475cffdc9c32adf28456a1)),
  closes [#151](https://github.com/ory/keto/issues/151)
  [#150](https://github.com/ory/keto/issues/150)

# [0.4.2-alpha.1](https://github.com/ory/keto/compare/v0.4.1-beta.1...v0.4.2-alpha.1) (2020-01-14)

Update CHANGELOG [ci skip]

### Unclassified

- Update CHANGELOG [ci skip]
  ([afaabde](https://github.com/ory/keto/commit/afaabde63affcf568e3090e55b4b957edff2890c))

# [0.4.1-beta.1](https://github.com/ory/keto/compare/v0.4.0-sandbox...v0.4.1-beta.1) (2020-01-13)

Update CHANGELOG [ci skip]

### Unclassified

- Update CHANGELOG [ci skip]
  ([e3ca5a7](https://github.com/ory/keto/commit/e3ca5a7d8b9827ffc7b31a8b5e459db3e912a590))
- Update SDK
  ([5dd6237](https://github.com/ory/keto/commit/5dd623755d4832f33c3dcefb778a9a70eace7b52))

# [0.4.0-alpha.1](https://github.com/ory/keto/compare/v0.3.9-sandbox...v0.4.0-alpha.1) (2020-01-13)

Move to new SDK generators (#146)

### Unclassified

- Move to new SDK generators (#146)
  ([4f51a09](https://github.com/ory/keto/commit/4f51a0948723efc092f1887b111d1e6dd590a075)),
  closes [#146](https://github.com/ory/keto/issues/146)
- Fix typos in the README (#144)
  ([85d838c](https://github.com/ory/keto/commit/85d838c0872c73eb70b5bfff1ccb175b07f6b1e4)),
  closes [#144](https://github.com/ory/keto/issues/144)

# [0.3.9-sandbox](https://github.com/ory/keto/compare/v0.3.8-sandbox...v0.3.9-sandbox) (2019-12-16)

Update go modules

### Unclassified

- Update go modules
  ([1151e07](https://github.com/ory/keto/commit/1151e0755c974b0aea86be5aaeae365ea9aef094))

# [0.3.7-sandbox](https://github.com/ory/keto/compare/v0.3.6-sandbox...v0.3.7-sandbox) (2019-12-11)

Update documentation banner image (#143)

### Unclassified

- Update documentation banner image (#143)
  ([e444755](https://github.com/ory/keto/commit/e4447552031a4f26ec21a336071b0bb19843df61)),
  closes [#143](https://github.com/ory/keto/issues/143)
- Revert incorrect license changes
  ([094c4f3](https://github.com/ory/keto/commit/094c4f30184d77a05044087c13e71ce4adb4d735))
- Fix invalid pseudo version ([#138](https://github.com/ory/keto/issues/138))
  ([79b4457](https://github.com/ory/keto/commit/79b4457f0162197ba267edbb8c0031c47e03bade))

# [0.3.6-sandbox](https://github.com/ory/keto/compare/v0.3.5-sandbox...v0.3.6-sandbox) (2019-10-16)

Resolve issues with mysql tests (#137)

### Unclassified

- Resolve issues with mysql tests (#137)
  ([ef5aec8](https://github.com/ory/keto/commit/ef5aec8e493199c46b78e8f1257aa41df9545f28)),
  closes [#137](https://github.com/ory/keto/issues/137)

# [0.3.5-sandbox](https://github.com/ory/keto/compare/v0.3.4-sandbox...v0.3.5-sandbox) (2019-08-21)

Implement roles and policies filter (#124)

### Documentation

- Incorporates changes from version v0.3.3-sandbox
  ([57686d2](https://github.com/ory/keto/commit/57686d2e30b229cae33e717eb8b3db9da3bdaf0a))
- README grammar fixes ([#114](https://github.com/ory/keto/issues/114))
  ([e592736](https://github.com/ory/keto/commit/e5927360300d8c4fbea841c1c2fb92b48b77885e))
- Updates issue and pull request templates
  ([#110](https://github.com/ory/keto/issues/110))
  ([80c8516](https://github.com/ory/keto/commit/80c8516efbcf33902d8a45f1dc7dbafff2aab8b1))
- Updates issue and pull request templates
  ([#111](https://github.com/ory/keto/issues/111))
  ([22305d0](https://github.com/ory/keto/commit/22305d0a9b5114de8125c16030bbcd1de695ae9b))
- Updates issue and pull request templates
  ([#112](https://github.com/ory/keto/issues/112))
  ([dccada9](https://github.com/ory/keto/commit/dccada9a2189bbd899c5c4a18665a92113fe6cd7))
- Updates issue and pull request templates
  ([#125](https://github.com/ory/keto/issues/125))
  ([15f373a](https://github.com/ory/keto/commit/15f373a16b8cfbd6cdad2bda5f161e171c566137))
- Updates issue and pull request templates
  ([#128](https://github.com/ory/keto/issues/128))
  ([eaf8e33](https://github.com/ory/keto/commit/eaf8e33f3904484635924bdac190c8dc7b60f939))
- Updates issue and pull request templates
  ([#130](https://github.com/ory/keto/issues/130))
  ([a440d14](https://github.com/ory/keto/commit/a440d142275a7a91a0a6bb487fe47d22247f4988))
- Updates issue and pull request templates
  ([#131](https://github.com/ory/keto/issues/131))
  ([dbf2cb2](https://github.com/ory/keto/commit/dbf2cb23c5b6f0f1ee0be5c0b5a58fb0c3dbefd1))
- Updates issue and pull request templates
  ([#132](https://github.com/ory/keto/issues/132))
  ([e121048](https://github.com/ory/keto/commit/e121048d10627ed32a07e26455efd69248f1bd95))
- Updates issue and pull request templates
  ([#133](https://github.com/ory/keto/issues/133))
  ([1b7490a](https://github.com/ory/keto/commit/1b7490abc1d5d0501b66595eb2d92834b6fb0345))

### Unclassified

- Implement roles and policies filter (#124)
  ([db94481](https://github.com/ory/keto/commit/db9448103621a6a8cd086a4cef6c6a22398e621f)),
  closes [#124](https://github.com/ory/keto/issues/124)
- Add adopters placeholder ([#129](https://github.com/ory/keto/issues/129))
  ([b814838](https://github.com/ory/keto/commit/b8148388b8bea97d1f1b4b54de2f0b8ef6b8b6c7))
- Improve documentation (#126)
  ([aabb04d](https://github.com/ory/keto/commit/aabb04d5f283d3c73eb3f3531b4e470ae716db5e)),
  closes [#126](https://github.com/ory/keto/issues/126)
- Create FUNDING.yml
  ([571b447](https://github.com/ory/keto/commit/571b447ed3a02f43623ef5c5adc09682b5f379bd))
- Use non-root user in image ([#116](https://github.com/ory/keto/issues/116))
  ([a493e55](https://github.com/ory/keto/commit/a493e550a8bb86d99164f4ea76dbcecf76c9c2c1))
- Remove binary license (#117)
  ([6e85f7c](https://github.com/ory/keto/commit/6e85f7c6f430e88fb4117a131f57bd69466a8ca1)),
  closes [#117](https://github.com/ory/keto/issues/117)

# [0.3.3-sandbox](https://github.com/ory/keto/compare/v0.3.1-sandbox...v0.3.3-sandbox) (2019-05-18)

ci: Resolve goreleaser issues (#108)

### Continuous Integration

- Resolve goreleaser issues ([#108](https://github.com/ory/keto/issues/108))
  ([5753f27](https://github.com/ory/keto/commit/5753f27a9e89ccdda7c02969217c253aa72cb94b))

### Documentation

- Incorporates changes from version v0.3.1-sandbox
  ([b8a0029](https://github.com/ory/keto/commit/b8a002937483a0f71fe5aba26bb18beb41886249))
- Updates issue and pull request templates
  ([#106](https://github.com/ory/keto/issues/106))
  ([54a5a27](https://github.com/ory/keto/commit/54a5a27f24a90ab3c5f9915f36582b85eecd0d62))

# [0.3.1-sandbox](https://github.com/ory/keto/compare/v0.3.0-sandbox...v0.3.1-sandbox) (2019-04-29)

ci: Use image that includes bash/sh for release docs (#103)

Signed-off-by: aeneasr <aeneas@ory.sh>

### Continuous Integration

- Use image that includes bash/sh for release docs
  ([#103](https://github.com/ory/keto/issues/103))
  ([e9d3027](https://github.com/ory/keto/commit/e9d3027fc62b20f28cd7a023222390e24d565eb1))

### Documentation

- Incorporates changes from version v0.3.0-sandbox
  ([605d2f4](https://github.com/ory/keto/commit/605d2f43621b806b750edc81d439edc92cfb7c38))

### Unclassified

- Allow configuration files and update UPGRADE guide. (#102)
  ([3934dc6](https://github.com/ory/keto/commit/3934dc6e690822358067b43920048d45a4b7799b)),
  closes [#102](https://github.com/ory/keto/issues/102)

# [0.3.0-sandbox](https://github.com/ory/keto/compare/v0.2.3-sandbox+oryOS.10...v0.3.0-sandbox) (2019-04-29)

docker: Remove full tag from build pipeline (#101)

Signed-off-by: aeneasr <aeneas@ory.sh>

### Documentation

- Update patrons
  ([c8dc7cd](https://github.com/ory/keto/commit/c8dc7cdc68676970328b55648b8d6e469c77fbfd))

### Unclassified

- Improve naming for ory policies
  ([#100](https://github.com/ory/keto/issues/100))
  ([b39703d](https://github.com/ory/keto/commit/b39703d362d333213fcb7d3782e363d09b6dabbd))
- Remove full tag from build pipeline
  ([#101](https://github.com/ory/keto/issues/101))
  ([602a273](https://github.com/ory/keto/commit/602a273dc5a0c29e80a22f04adb937ab385c4512))
- Remove duplicate code in Makefile (#99)
  ([04f5223](https://github.com/ory/keto/commit/04f52231509dd0f3a57d745918fc43fff7c595ff)),
  closes [#99](https://github.com/ory/keto/issues/99)
- Add tracing support and general improvements (#98)
  ([63b3946](https://github.com/ory/keto/commit/63b3946e0ae1fa23c6a359e9a64b296addff868c)),
  closes [#98](https://github.com/ory/keto/issues/98):

  This patch improves the internal configuration and service management. It adds
  support for distributed tracing and resolves several issues in the release
  pipeline and CLI.

  Additionally, composable docker-compose configuration files have been added.

  Several bugs have been fixed in the release management pipeline.

- Add content-type in the response of allowed
  ([#90](https://github.com/ory/keto/issues/90))
  ([39a1486](https://github.com/ory/keto/commit/39a1486dc53456189d30380460a9aeba198fa9e9))
- Fix disable-telemetry check ([#85](https://github.com/ory/keto/issues/85))
  ([38b5383](https://github.com/ory/keto/commit/38b538379973fa34bd2bf24dcb2e6dbedf324e1e))
- Fix remove member from role ([#87](https://github.com/ory/keto/issues/87))
  ([698e161](https://github.com/ory/keto/commit/698e161989331ca5a3a0769301d9694ef805a876)),
  closes [#74](https://github.com/ory/keto/issues/74)
- Fix the type of conditions in the policy
  ([#86](https://github.com/ory/keto/issues/86))
  ([fc1ced6](https://github.com/ory/keto/commit/fc1ced63bd39c9fbf437e419dfc384343e36e0ee))
- Move Go SDK generation to go-swagger
  ([#94](https://github.com/ory/keto/issues/94))
  ([9f48a95](https://github.com/ory/keto/commit/9f48a95187a7b6160108cd7d0301590de2e58f07)),
  closes [#92](https://github.com/ory/keto/issues/92)
- Send 403 when authorization result is negative
  ([#93](https://github.com/ory/keto/issues/93))
  ([de806d8](https://github.com/ory/keto/commit/de806d892819db63c1abc259ab06ee08d87895dc)),
  closes [#75](https://github.com/ory/keto/issues/75)
- Update dependencies ([#91](https://github.com/ory/keto/issues/91))
  ([4d44174](https://github.com/ory/keto/commit/4d4417474ebf8cc69d01e5ac82633b966cdefbc7))
- storage/memory: Fix upsert with pre-existing key will causes duplicate records
  (#88)
  ([1cb8a36](https://github.com/ory/keto/commit/1cb8a36a08883b785d9bb0a4be1ddc00f1f9d358)),
  closes [#88](https://github.com/ory/keto/issues/88)
  [#80](https://github.com/ory/keto/issues/80)

# [0.2.3-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.2.2-sandbox+oryOS.10...v0.2.3-sandbox+oryOS.10) (2019-02-05)

dist: Fix packr build pipeline (#84)

Closes #73 Closes #81

Signed-off-by: aeneasr <aeneas@ory.sh>

### Documentation

- Add documentation for glob matching
  ([5c8babb](https://github.com/ory/keto/commit/5c8babbfbae01a78f30cfbff92d8e9c3a6b09027))
- Incorporates changes from version v0.2.2-sandbox+oryOS.10
  ([ed7af3f](https://github.com/ory/keto/commit/ed7af3fa4e5d1d0d03b5366f4cf865a5b82ec293))
- Properly generate api.swagger.json
  ([18e3f84](https://github.com/ory/keto/commit/18e3f84cdeee317f942d61753399675c98886e5d))

### Unclassified

- Add placeholder go file for rego inclusion
  ([6a6f64d](https://github.com/ory/keto/commit/6a6f64d8c59b496f6cf360f55eba1e16bf5380f1))
- Add support for glob matching
  ([bb76c6b](https://github.com/ory/keto/commit/bb76c6bebe522fc25448c4f4e4d1ef7c530a725f))
- Ex- and import rego subdirectories for `go get`
  [#77](https://github.com/ory/keto/issues/77)
  ([59cc053](https://github.com/ory/keto/commit/59cc05328f068fc3046b2dbc022a562fd5d67960)),
  closes [#73](https://github.com/ory/keto/issues/73)
- Fix packr build pipeline ([#84](https://github.com/ory/keto/issues/84))
  ([65a87d5](https://github.com/ory/keto/commit/65a87d564d237bc979bb5962beff7d3703d9689f)),
  closes [#73](https://github.com/ory/keto/issues/73)
  [#81](https://github.com/ory/keto/issues/81)
- Import glob in rego/doc.go
  ([7798442](https://github.com/ory/keto/commit/7798442553cfe7989a23d2c389c8c63a24013543))
- Properly handle dbal error
  ([6811607](https://github.com/ory/keto/commit/6811607ea79c8f3155a17bc1aea566e9e4680616))
- Properly handle TLS certificates if set
  ([36399f0](https://github.com/ory/keto/commit/36399f09261d4f3cb5e053679eee3cb15da2df19)),
  closes [#73](https://github.com/ory/keto/issues/73)

# [0.2.2-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.2.1-sandbox+oryOS.10...v0.2.2-sandbox+oryOS.10) (2018-12-13)

ci: Fix docker push arguments in publish task

Signed-off-by: aeneasr <aeneas@ory.sh>

### Continuous Integration

- Fix docker push arguments in publish task
  ([f03c77c](https://github.com/ory/keto/commit/f03c77c6b7461ab81cb03265cbec909ac45c2259))

# [0.2.1-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.2.0-sandbox+oryOS.10...v0.2.1-sandbox+oryOS.10) (2018-12-13)

ci: Fix docker release task

Signed-off-by: aeneasr <aeneas@ory.sh>

### Continuous Integration

- Fix docker release task
  ([7a0414f](https://github.com/ory/keto/commit/7a0414f614b6cc8b1d78cfbb773a2f0192d00d23))

# [0.2.0-sandbox+oryOS.10](https://github.com/ory/keto/compare/v0.0.1...v0.2.0-sandbox+oryOS.10) (2018-12-13)

all: gofmt

Signed-off-by: aeneasr <aeneas@ory.sh>

### Documentation

- Adds banner
  ([0ec1d8f](https://github.com/ory/keto/commit/0ec1d8f5e843465d17ac4c8f91f18e5badf16900))
- Adds GitHub templates & code of conduct
  ([#31](https://github.com/ory/keto/issues/31))
  ([a11e898](https://github.com/ory/keto/commit/a11e8980f2af528f1357659966123d0cbf7d43db))
- Adds link to examples repository
  ([#32](https://github.com/ory/keto/issues/32))
  ([7061a2a](https://github.com/ory/keto/commit/7061a2aa31652a9e0c2d449facb1201bfa11fd3f))
- Adds security console image
  ([fd27fc9](https://github.com/ory/keto/commit/fd27fc9cce50beb3d0189e0a93300879fd7149db))
- Changes hydra to keto in readme
  ([9dab531](https://github.com/ory/keto/commit/9dab531744cf5b0ae98862945d44b07535595781))
- Deprecate old versions in logs
  ([955d647](https://github.com/ory/keto/commit/955d647307a48ee7cf2d3f9fb4263072adf42299))
- Incorporates changes from version
  ([85c4d81](https://github.com/ory/keto/commit/85c4d81a192e92f874c106b91cfa6fb404d9a34a))
- Incorporates changes from version v0.0.0-testrelease.1
  ([6062dd4](https://github.com/ory/keto/commit/6062dd4a894607f5f1ead119af20cc8bdbe15bef))
- Incorporates changes from version v0.0.1-1-g85c4d81
  ([f4606fc](https://github.com/ory/keto/commit/f4606fce0326bece2a89dadc029bc5ce9778df18))
- Incorporates changes from version v0.0.1-11-g114914f
  ([92a4dca](https://github.com/ory/keto/commit/92a4dca7a41dcf3a88c4800bf6d2217f33cfcdd1))
- Incorporates changes from version v0.0.1-16-g7d8a8ad
  ([2b76a83](https://github.com/ory/keto/commit/2b76a83755153b3f8a2b8d28c5b0029d96d567b6))
- Incorporates changes from version v0.0.1-18-g099e7e0
  ([70b12ad](https://github.com/ory/keto/commit/70b12adf5bcc0e890d6707e11e891e6cedfb3d87))
- Incorporates changes from version v0.0.1-20-g97ccbe6
  ([b21d56e](https://github.com/ory/keto/commit/b21d56e599c7eb4c1769bc18878f7d5818b73023))
- Incorporates changes from version v0.0.1-30-gaf2c3b5
  ([a1d0dcc](https://github.com/ory/keto/commit/a1d0dcc78a9506260f86df00e4dff8ab02909ce1))
- Incorporates changes from version v0.0.1-32-gedb5a60
  ([a5c369a](https://github.com/ory/keto/commit/a5c369a90da67c96bbde60e673c67f50b841fadd))
- Incorporates changes from version v0.0.1-6-g570783e
  ([0fcbbcb](https://github.com/ory/keto/commit/0fcbbcb02f1d748f9c733c86368b223b2ee4c6e2))
- Incorporates changes from version v0.0.1-7-g0fcbbcb
  ([c0141a8](https://github.com/ory/keto/commit/c0141a8ec22ea1260bf2d45d72dfe06737ec0115))
- Incorporates changes from version v0.1.0-sandbox
  ([9ee0664](https://github.com/ory/keto/commit/9ee06646d2cfb2d69abdcc411e31d14957437a1e))
- Incorporates changes from version v1.0.0-beta.1-1-g162d7b8
  ([647c5a9](https://github.com/ory/keto/commit/647c5a9e1bc8d9d635bf6f2511c3faa9a9daefef))
- Incorporates changes from version v1.0.0-beta.2-11-g2b280bb
  ([936889d](https://github.com/ory/keto/commit/936889d760f04a03d498f65331d653cbad3702d0))
- Incorporates changes from version v1.0.0-beta.2-13-g382e1d3
  ([883df44](https://github.com/ory/keto/commit/883df44a922f3daee86597af467072555cadc7e7))
- Incorporates changes from version v1.0.0-beta.2-15-g74450da
  ([48dd9f1](https://github.com/ory/keto/commit/48dd9f1ffbeaa99ac8dc27085c5a50f9244bf9c3))
- Incorporates changes from version v1.0.0-beta.2-3-gf623c52
  ([b6b90e5](https://github.com/ory/keto/commit/b6b90e5b2180921f78064a60666704b4e72679b6))
- Incorporates changes from version v1.0.0-beta.2-5-g3852be5
  ([3f09090](https://github.com/ory/keto/commit/3f09090a2f82f3f29154c19217cea0a10d65ea3a))
- Incorporates changes from version v1.0.0-beta.2-9-gc785187
  ([4c30a3c](https://github.com/ory/keto/commit/4c30a3c0ad83ba80e1857b41211e7ddade06c4cf))
- Incorporates changes from version v1.0.0-beta.3-1-g06adbf1
  ([0ba3c06](https://github.com/ory/keto/commit/0ba3c0674832b641ef5e0c3f0d60d81ed3a647b2))
- Incorporates changes from version v1.0.0-beta.3-10-g9994967
  ([d2345ca](https://github.com/ory/keto/commit/d2345ca3beb354d6ee7c7926c1a5ddb425d6b405))
- Incorporates changes from version v1.0.0-beta.3-12-gc28b521
  ([b4d792f](https://github.com/ory/keto/commit/b4d792f74055853f05ca46c67625ffd432fc74fd))
- Incorporates changes from version v1.0.0-beta.3-3-g9e16605
  ([c43bf2b](https://github.com/ory/keto/commit/c43bf2b5232bed9106dd47d7eb53d2f93bfe260d))
- Incorporates changes from version v1.0.0-beta.3-5-ga11e898
  ([b9d9b8e](https://github.com/ory/keto/commit/b9d9b8ee33ab957f43f99c427a88ade847e79ed0))
- Incorporates changes from version v1.0.0-beta.3-8-g7061a2a
  ([d76ff9d](https://github.com/ory/keto/commit/d76ff9dc9a4c8a8f1286eeb139d8f5af9617f421))
- Incorporates changes from version v1.0.0-beta.5
  ([0dc314c](https://github.com/ory/keto/commit/0dc314c7888020b40e12eb59fd77135044fd063b))
- Incorporates changes from version v1.0.0-beta.6-1-g5e97104
  ([f14c8ed](https://github.com/ory/keto/commit/f14c8edd7204a811e333ea84429cf837b4e7d27b))
- Incorporates changes from version v1.0.0-beta.8
  ([5045b59](https://github.com/ory/keto/commit/5045b59e2a83d6ab047b1b95c581d7c34e96a2e0))
- Incorporates changes from version v1.0.0-beta.9
  ([be2f035](https://github.com/ory/keto/commit/be2f03524721ef47ecb1c9aec57c2696174e0657))
- Properly sets up changelog TOC
  ([e0acd67](https://github.com/ory/keto/commit/e0acd670ab19c0d6fd36733fea164e2b0414597d))
- Puts toc in the right place
  ([114914f](https://github.com/ory/keto/commit/114914fa354f784b310bc9dfd232a011e0d98d99))
- Revert changes from test release
  ([ab3a64d](https://github.com/ory/keto/commit/ab3a64d3d41292364c5947db98c4d27a8223853e))
- Update documentation links ([#67](https://github.com/ory/keto/issues/67))
  ([d22d413](https://github.com/ory/keto/commit/d22d413c7a001ccaa96b4c013665153f41831614))
- Update link to security console
  ([846ce4b](https://github.com/ory/keto/commit/846ce4baa9da5954bd30996f489885a026c48185))
- Update migration guide
  ([3c44b58](https://github.com/ory/keto/commit/3c44b58613e46ed39d42463537773fe9d95a54da))
- Update to latest changes
  ([1625123](https://github.com/ory/keto/commit/1625123ed342f019d5e7ab440eb37da310570842))
- Updates copyright notice
  ([9dd5578](https://github.com/ory/keto/commit/9dd557825dfd3b9d589c9db2ccb201638debbaae))
- Updates installation guide
  ([f859645](https://github.com/ory/keto/commit/f859645f230f405cfabed0c1b9a2b67b1a3841d3))
- Updates issue and pull request templates
  ([#52](https://github.com/ory/keto/issues/52))
  ([941cae6](https://github.com/ory/keto/commit/941cae6fee058f68eabbbf4dd9cafad4760e108f))
- Updates issue and pull request templates
  ([#53](https://github.com/ory/keto/issues/53))
  ([7b222d2](https://github.com/ory/keto/commit/7b222d285e74c0db482136b23f37072216b3acb0))
- Updates issue and pull request templates
  ([#54](https://github.com/ory/keto/issues/54))
  ([f098639](https://github.com/ory/keto/commit/f098639b5e748151810848fdd3173e0246bc03dc))
- Updates link to guide and header
  ([437c255](https://github.com/ory/keto/commit/437c255ecfff4127fb586cc069e07f86988ad1ba))
- Updates link to open collective
  ([382e1d3](https://github.com/ory/keto/commit/382e1d34c7da0ba0447b78506a749bd7f0085f48))
- Updates links to docs
  ([d84be3b](https://github.com/ory/keto/commit/d84be3b6a8e5eb284ec3fb137ee774ba5ee0d529))
- Updates newsletter link in README
  ([2dc36b2](https://github.com/ory/keto/commit/2dc36b21c8af8e3e39f093198715ea24b65d65af))

### Unclassified

- Add Go SDK factory
  ([99db7e6](https://github.com/ory/keto/commit/99db7e6d4edac88794266a01ddfab9cd0632e95a))
- Add go SDK interface
  ([3dd5f7d](https://github.com/ory/keto/commit/3dd5f7d61bb460c34744b84a34755bfb8219b304))
- Add health handlers
  ([bddb949](https://github.com/ory/keto/commit/bddb949459d05002b0f8882d981e4f63fdddf25f))
- Add policy list handler
  ([a290619](https://github.com/ory/keto/commit/a290619d01d15eb8e3b4e33ede1058d316ee807a))
- Add role iterator in list handler
  ([a3eb696](https://github.com/ory/keto/commit/a3eb6961783f7b562f0a0d0a7e2819bffebce5b8))
- Add SDK generation to circle ci
  ([9b37165](https://github.com/ory/keto/commit/9b37165873bcb0cc5dc60d2514d9824a073466a1))
- Adds ability to update a role using PUT
  ([#14](https://github.com/ory/keto/issues/14))
  ([97ccbe6](https://github.com/ory/keto/commit/97ccbe6d808823c56901ad237878aa6d53cddeeb)):

  - transfer UpdateRoleMembers from https://github.com/ory/hydra/pull/768 to
    keto

  - fix tests by using right http method & correcting sql request

  - Change behavior to overwrite the whole role instead of just the members.

  * small sql migration fix

- Adds log message when telemetry is active
  ([f623c52](https://github.com/ory/keto/commit/f623c52655ff85b7f7209eb73e94eb66a297c5b7))
- Clean up vendor dependencies
  ([9a33c23](https://github.com/ory/keto/commit/9a33c23f4d37ab88b4d643fd79204334d73404c6))
- Do not split empty scope ([#45](https://github.com/ory/keto/issues/45))
  ([b29cf8c](https://github.com/ory/keto/commit/b29cf8cc92607e13457dba8331f5c9286054c8c1))
- Fix typo in help command in env var name
  ([#39](https://github.com/ory/keto/issues/39))
  ([8a5016c](https://github.com/ory/keto/commit/8a5016cd75be78bb42a9a38bfd453ad5722db9db)),
  closes [#25](https://github.com/ory/keto/issues/25)
- Fixes environment variable typos
  ([566d588](https://github.com/ory/keto/commit/566d588e4fca12399966718b725fe4461a28e51e))
- Fixes typo in help command
  ([74450da](https://github.com/ory/keto/commit/74450da18a27513820328c28f72203653c664367)),
  closes [#25](https://github.com/ory/keto/issues/25)
- Format code
  ([637c78c](https://github.com/ory/keto/commit/637c78cba697682b544473a3af9b6ae7715561aa))
- Gofmt
  ([a8d7f9f](https://github.com/ory/keto/commit/a8d7f9f546ae2f3b8c3fa643d8e19b68ca26cc67))
- Improve compose documentation
  ([6870443](https://github.com/ory/keto/commit/68704435f3c299b853f4ff5cacae285b09ada3b5))
- Improves usage of metrics middleware
  ([726c4be](https://github.com/ory/keto/commit/726c4bedfc3f02fdac380930e32f37c251e51aa4))
- Improves usage of metrics middleware
  ([301f386](https://github.com/ory/keto/commit/301f38605af66abae4d28ed0cac90d0b82b655c4))
- Introduce docker-compose file for testing
  ([ba857e3](https://github.com/ory/keto/commit/ba857e3859966e857c5a741825411575e17446de))
- Introduces health and version endpoints
  ([6a9da74](https://github.com/ory/keto/commit/6a9da74f693ee6c15a775ab8d652582aea093601))
- List roles from keto_role table ([#28](https://github.com/ory/keto/issues/28))
  ([9e16605](https://github.com/ory/keto/commit/9e166054b8d474fbce6983d5d00eeeb062fc79b1))
- Properly names flags
  ([af2c3b5](https://github.com/ory/keto/commit/af2c3b5bc96e95fb31b1db5c7fe6dfd6b6fc5b20))
- Require explicit CORS enabling ([#42](https://github.com/ory/keto/issues/42))
  ([9a45107](https://github.com/ory/keto/commit/9a45107af304b2a8e663a532e4f6e4536f15888c))
- Update dependencies
  ([663d8b1](https://github.com/ory/keto/commit/663d8b13e99694a57752cd60a68342b81b041c66))
- Switch to rego as policy decision engine (#48)
  ([ee9bcf2](https://github.com/ory/keto/commit/ee9bcf2719178e5a8dccca083a90313947a8a63b)),
  closes [#48](https://github.com/ory/keto/issues/48)
- Update hydra to v1.0.0-beta.6 ([#35](https://github.com/ory/keto/issues/35))
  ([5e97104](https://github.com/ory/keto/commit/5e971042afff06e2a6ee3b54d2fea31687203623))
- Update npm package registry
  ([a53d3d2](https://github.com/ory/keto/commit/a53d3d23e11fde5dcfbb27a2add1049f4d8e10e6))
- Enable TLS option to serve API (#46)
  ([2f62063](https://github.com/ory/keto/commit/2f620632d0375bf9c7e58dbfb49627c02c66abf3)),
  closes [#46](https://github.com/ory/keto/issues/46)
- Make introspection authorization optional
  ([e5460ad](https://github.com/ory/keto/commit/e5460ad884cd018cd6177324b949cd66bfd53bc7))
- Properly output telemetry information
  ([#33](https://github.com/ory/keto/issues/33))
  ([9994967](https://github.com/ory/keto/commit/9994967b0ca54a62b8b0088fe02be9e890d9574b))
- Remove ORY Hydra dependency ([#44](https://github.com/ory/keto/issues/44))
  ([d487344](https://github.com/ory/keto/commit/d487344fe7e07cb6370371c6b0b6cf3cca767ed1))
- Resolves an issue with the hydra migrate command
  ([2b280bb](https://github.com/ory/keto/commit/2b280bb57c9073a9c8384cde0b14a6991cfacdb6)),
  closes [#23](https://github.com/ory/keto/issues/23)
- Upgrade superagent version ([#41](https://github.com/ory/keto/issues/41))
  ([9c80dbc](https://github.com/ory/keto/commit/9c80dbcc1cc63243839b58ca56ac9be104797887))
- gofmt
  ([777b1be](https://github.com/ory/keto/commit/777b1be1378d314e7cfde0c34450afcce7e590a5))
- Updates README.md (#34)
  ([c28b521](https://github.com/ory/keto/commit/c28b5219fd64314a75ee3c848a80a0c5974ebb7d)),
  closes [#34](https://github.com/ory/keto/issues/34)
- Properly parses cors options
  ([edb5a60](https://github.com/ory/keto/commit/edb5a600f2ce16c0847ee5ef399fa5a41b1e736a))
- Removes additional output if no args are passed
  ([703e124](https://github.com/ory/keto/commit/703e1246ce0fd89066b497c45f0c6cadeb06c331))
- Resolves broken role test
  ([b6c7f9c](https://github.com/ory/keto/commit/b6c7f9c33c4c1f43164d6da0ec7f2553f1f4c598))
- Resolves minor typos and updates install guide
  ([3852be5](https://github.com/ory/keto/commit/3852be56cb81df966a85d4c828de0397d9e74768))
- Updates to latest sqlcon
  ([2c9f643](https://github.com/ory/keto/commit/2c9f643042ff4edffae8bd41834d2a57c923871c))
- Use roles in warden decision
  ([c785187](https://github.com/ory/keto/commit/c785187e31fc7a4b8b762a5e27fac66dcaa97513)),
  closes [#21](https://github.com/ory/keto/issues/21)
  [#19](https://github.com/ory/keto/issues/19)
- authn/client: Payload is now prefixed with client
  ([8584d94](https://github.com/ory/keto/commit/8584d94cfb18deb37ae32ae601f4cd15c14067e7))

# [0.0.1](https://github.com/ory/keto/compare/4f00bc96ece3180a888718ec3c41c69106c86f56...v0.0.1) (2018-05-20)

authn: Checks token_type is "access_token", if set

Closes #1

### Documentation

- Incorporates changes from version
  ([b5445a0](https://github.com/ory/keto/commit/b5445a0fc5b6f813cd1731b20c8c5c79d7c4cdf8))
- Incorporates changes from version
  ([295ff99](https://github.com/ory/keto/commit/295ff998af55777823b04f423e365fd58e61753b))
- Incorporates changes from version
  ([bd44d41](https://github.com/ory/keto/commit/bd44d41b2781e33353082397c47390a27f749e16))
- Updates readme and upgrades
  ([0f95dbb](https://github.com/ory/keto/commit/0f95dbb967fd17b607caa999ae30453f5f599739))
- Uses keto repo for changelog
  ([14c0b2a](https://github.com/ory/keto/commit/14c0b2a2bd31566f2b9048831f894aba05c5b15d))

### Unclassified

- Adds migrate commands to the proper parent command
  ([231c70d](https://github.com/ory/keto/commit/231c70d816b0736a51eddc1fa0445bac672b1b2f))
- Checks token_type is "access_token", if set
  ([d2b8f5d](https://github.com/ory/keto/commit/d2b8f5d313cce597566bd18e4f3bea4a423a62ee)),
  closes [#1](https://github.com/ory/keto/issues/1)
- Removes old test
  ([07b733b](https://github.com/ory/keto/commit/07b733bfae4b733e3e2124545b92c537dabbdcf0))
- Renames subject to sub in response payloads
  ([ca4d540](https://github.com/ory/keto/commit/ca4d5408000be2b896d38eaaf5e67a3fc0a566da))
- Tells linguist to ignore SDK files
  ([f201eb9](https://github.com/ory/keto/commit/f201eb95f3309a60ac50f42cfba0bae2e38e8d13))
- Retries SQL connection on migrate commands
  ([3d33d73](https://github.com/ory/keto/commit/3d33d73c009077c5bf30ae4b03802904bfb5d5b2)):

  This patch also introduces a fatal error if migrations fail

- cmd/server: Resolves DBAL not handling postgres properly
  ([dedc32a](https://github.com/ory/keto/commit/dedc32ab218923243b1955ce5bcbbdc5cc416953))
- cmd/server: Improves error message in migrate command
  ([4b17ce8](https://github.com/ory/keto/commit/4b17ce8848113cae807840182d1a318190c2a9b3))
- Resolves travis and docker issues
  ([6f4779c](https://github.com/ory/keto/commit/6f4779cc51bf4f2ee5b97541fb77d8f882497710))
- Adds OAuth2 Client Credentials authenticator and warden endpoint
  ([c55139b](https://github.com/ory/keto/commit/c55139b51e636834759706499a2aec1451f4fbd9))
- Adds SDK helpers
  ([a1c2608](https://github.com/ory/keto/commit/a1c260801d9366fccf4bfb4fc64b2c67fc594565))
- Resolves SDK and test issues (#4)
  ([2d4cd98](https://github.com/ory/keto/commit/2d4cd9805af3081bbcbea3f806ca066d35385a4b)),
  closes [#4](https://github.com/ory/keto/issues/4)
- Initial project commit
  ([a592e51](https://github.com/ory/keto/commit/a592e5126f130f8b673fff6c894fdbd9fb56f81c))
- Initial commit
  ([4f00bc9](https://github.com/ory/keto/commit/4f00bc96ece3180a888718ec3c41c69106c86f56))
