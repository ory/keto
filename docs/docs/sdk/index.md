---
id: index
title: Overview
---

All SDKs use automated code generation provided by
[`openapi-generator`](https://github.com/OpenAPITools/openapi-generator).
Unfortunately, `openapi-generator` has serious breaking changes in the generated
code when upgrading versions. Therefore, we do not make backwards compatibility
promises with regards to the generated SDKs. We hope to improve this process in
the future.

Before you check out the SDKs, head over to the [REST API](../reference/api.mdx)
documentation which includes code samples for common programming languages for
each REST endpoint.

We publish our SDKs for popular languages in their respective package
repositories:

- [Dart](https://github.com/ory/sdk/tree/master/clients/keto/dart)
- [.NET](https://github.com/ory/sdk/tree/master/clients/keto/dotnet)
- [Go](https://github.com/ory/keto-client-go)
- [Java](https://search.maven.org/artifact/sh.ory.keto/keto-client)
- [PHP](https://packagist.org/packages/ory/keto-client)
- [Python](https://pypi.org/project/ory-keto-client/)
- [Ruby](https://rubygems.org/gems/ory-keto-client)
- [Rust](https://github.com/ory/sdk/tree/master/clients/keto/rust)
- [NodeJS](https://www.npmjs.com/package/@oryd/keto-client) (with TypeScript)

Missing your programming language?
[Create an issue](https://github.com/ory/keto/issues) and help us build, test
and publish the SDK for your programming language!
