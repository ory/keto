---
id: index
title: Overview
---

All SDKs use automated code generation provided by
[`openapi-generator`](https://github.com/OpenAPITools/openapi-generator) and
[protoc](https://github.com/protocolbuffers/protobuf). Unfortunately,
`openapi-generator` has serious breaking changes in the generated code when
upgrading versions. Therefore, we do not make backwards compatibility promises
with regards to the generated SDKs. We hope to improve this process in the
future.

Before you check out the SDKs, head over to the
[REST API](../reference/rest-api.mdx) and [gRPC API](../reference/proto-api.mdx)
documentation which includes code samples for common programming languages for
each REST endpoint.

We publish our SDKs for popular languages in their respective package
repositories:

- [Dart](https://pub.dev/packages/ory_keto_client)
- [.NET](https://www.nuget.org/packages/Ory.Keto.Client/)
- [Go REST](https://github.com/ory/keto-client-go)
- [Go gRPC](https://github.com/ory/keto/blob/master/proto/ory/keto/acl)
- [Java](https://search.maven.org/artifact/sh.ory.keto/keto-client)
- [PHP](https://packagist.org/packages/ory/keto-client)
- [Python](https://pypi.org/project/ory-keto-client/)
- [Ruby](https://rubygems.org/gems/ory-keto-client)
- [Rust](https://crates.io/crates/ory-keto-client)
- [NodeJS REST](https://www.npmjs.com/package/@ory/keto-client) (with
  TypeScript)
- [NodeJS gRPC](https://www.npmjs.com/package/@ory/keto-grpc-client) (with
  TypeScript)

Take a look at the source:
[Generated SDKs for Ory Keto](https://github.com/ory/sdk/tree/master/clients/keto/)

Missing your programming language?
[Create an issue](https://github.com/ory/keto/issues) and help us build, test
and publish the SDK for your programming language!
