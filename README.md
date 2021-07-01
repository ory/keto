<h1 align="center"><img src="https://raw.githubusercontent.com/ory/meta/master/static/banners/keto.svg" alt="ORY Keto - Open Source & Cloud Native Access Control Server"></h1>

<h4 align="center">    
    <a href="https://www.ory.sh/chat">Chat</a> |
    <a href="https://github.com/ory/keto/discussions">Discusssions</a> |
    <a href="http://eepurl.com/di390P">Newsletter</a><br/><br/>
    <a href="https://www.ory.sh/docs/keto/">Guide</a> |
    <a href="https://www.ory.sh/docs/keto/sdk/api">API Docs</a> |
    <a href="https://godoc.org/github.com/ory/keto">Code Docs</a><br/><br/>
    <a href="https://opencollective.com/ory">Support this project!</a>
</h4>

---

<p align="left">
    <a href="https://circleci.com/gh/ory/keto/tree/master"><img src="https://circleci.com/gh/ory/keto.svg?style=shield" alt="Build Status"></a>
    <a href="https://coveralls.io/github/ory/keto?branch=master"> <img src="https://coveralls.io/repos/ory/keto/badge.svg?branch=master&service=github" alt="Coverage Status"></a>
    <a href="https://goreportcard.com/report/github.com/ory/keto"><img src="https://goreportcard.com/badge/github.com/ory/keto" alt="Go Report Card"></a>
    <a href="https://pkg.go.dev/github.com/ory/keto"><img src="https://pkg.go.dev/badge/www.github.com/ory/keto" alt="PkgGoDev"></a>
    <a href="#backers" alt="sponsors on Open Collective"><img src="https://opencollective.com/ory/backers/badge.svg" /></a> <a href="#sponsors" alt="Sponsors on Open Collective"><img src="https://opencollective.com/ory/sponsors/badge.svg" /></a>
    <a href="https://github.com/ory/keto/blob/master/CODE_OF_CONDUCT.md" alt="Ory Code of Conduct"><img src="https://img.shields.io/badge/ory-code%20of%20conduct-green" /></a>
</p>
Ory Keto is the first and only open source implementation of "Zanzibar: Google's
Consistent, Global Authorization System":

> Determining whether online users are authorized to access digital objects is
> central to preserving privacy. This paper presents the design, implementation,
> and deployment of Zanzibar, a global system for storing and evaluating access
> control lists. Zanzibar provides a uniform data model and configuration
> language for expressing a wide range of access control policies from hundreds
> of client services at Google, including Calendar, Cloud, Drive, Maps, Photos,
> and YouTube. Its authorization decisions respect causal ordering of user
> actions and thus provide external consistency amid changes to access control
> lists and object contents. Zanzibar scales to trillions of access control
> lists and millions of authorization requests per second to support services
> used by billions of people. It has maintained 95th-percentile latency of less
> than 10 milliseconds and availability of greater than 99.999% over 3 years of
> production use.
>
> [Source](https://research.google/pubs/pub48190/)

If you need to know if a user (or robot, car, service) is allowed to do
something - Ory Keto is the right fit for you.

Currently, Ory Keto implements the basic API contracts for managing and checking
relations ("permissions") with HTTP and gRPC APIs. Future versions will include
features such as userset rewrites (e.g. RBAC-style role-permission models),
Zookies, and more. An overview of what is implemented and upcoming can be found
at
[Implemented and Planned Features](https://www.ory.sh/keto/docs/next/implemented-planned-features).

---

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Who's using it?](#whos-using-it)
  - [Installation](#installation)
- [Ecosystem](#ecosystem)
  - [ORY Kratos: Identity and User Infrastructure and Management](#ory-kratos-identity-and-user-infrastructure-and-management)
  - [ORY Hydra: OAuth2 & OpenID Connect Server](#ory-hydra-oauth2--openid-connect-server)
  - [ORY Oathkeeper: Identity & Access Proxy](#ory-oathkeeper-identity--access-proxy)
  - [ORY Keto: Access Control Policies as a Server](#ory-keto-access-control-policies-as-a-server)
- [Security](#security)
  - [Disclosing vulnerabilities](#disclosing-vulnerabilities)
- [Telemetry](#telemetry)
  - [Guide](#guide)
  - [HTTP API documentation](#http-api-documentation)
  - [Upgrading and Changelog](#upgrading-and-changelog)
  - [Command line documentation](#command-line-documentation)
  - [Develop](#develop)
    - [Dependencies](#dependencies)
    - [Install from source](#install-from-source)
    - [Formatting Code](#formatting-code)
    - [Running Tests](#running-tests)
      - [Short Tests](#short-tests)
      - [Regular Tests](#regular-tests)
      - [End-to-End Tests](#end-to-end-tests)
    - [Build Docker](#build-docker)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Who's using it?

<!--BEGIN ADOPTERS-->

The Ory community stands on the shoulders of individuals, companies, and
maintainers. We thank everyone involved - from submitting bug reports and
feature requests, to contributing patches, to sponsoring our work. Our community
is 1000+ strong and growing rapidly. The Ory stack protects 16.000.000.000+ API
requests every month with over 250.000+ active service nodes. We would have
never been able to achieve this without each and everyone of you!

The following list represents companies that have accompanied us along the way
and that have made outstanding contributions to our ecosystem. _If you think
that your company deserves a spot here, reach out to
<a href="mailto:office-muc@ory.sh">office-muc@ory.sh</a> now_!

**Please consider giving back by becoming a sponsor of our open source work on
<a href="https://www.patreon.com/_ory">Patreon</a> or
<a href="https://opencollective.com/ory">Open Collective</a>.**

<table>
    <thead>
        <tr>
            <th>Type</th>
            <th>Name</th>
            <th>Logo</th>
            <th>Website</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Sponsor</td>
            <td>Raspberry PI Foundation</td>
            <td align="center"><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/raspi.svg" alt="Raspberry PI Foundation"></td>
            <td><a href="https://www.raspberrypi.org/">raspberrypi.org</a></td>
        </tr>
        <tr>
            <td>Contributor</td>
            <td>Kyma Project</td>
            <td align="center"><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/kyma.svg" alt="Kyma Project"></td>
            <td><a href="https://kyma-project.io">kyma-project.io</a></td>
        </tr>
        <tr>
            <td>Sponsor</td>
            <td>ThoughtWorks</td>
            <td align="center"><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/tw.svg" alt="ThoughtWorks"></td>
            <td><a href="https://www.thoughtworks.com/">thoughtworks.com</a></td>
        </tr>
        <tr>
            <td>Sponsor</td>
            <td>Tulip</td>
            <td align="center"><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/tulip.svg" alt="Tulip Retail"></td>
            <td><a href="https://tulip.com/">tulip.com</a></td>
        </tr>
        <tr>
            <td>Sponsor</td>
            <td>Cashdeck / All My Funds</td>
            <td align="center"><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/allmyfunds.svg" alt="All My Funds"></td>
            <td><a href="https://cashdeck.com.au/">cashdeck.com.au</a></td>
        </tr>
        <tr>
            <td>Sponsor</td>
            <td>3Rein</td>
            <td align="center"><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/3R-horiz.svg" alt="3Rein"></td>
            <td><a href="https://3rein.com/">3rein.com</a></td>
        </tr>
        <tr>
            <td>Contributor</td>
            <td>Hootsuite</td>
            <td align="center"><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/hootsuite.svg" alt="Hootsuite"></td>
            <td><a href="https://hootsuite.com/">hootsuite.com</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Segment</td>
            <td align="center"><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/segment.svg" alt="Segment"></td>
            <td><a href="https://segment.com/">segment.com</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Arduino</td>
            <td align="center"><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/arduino.svg" alt="Arduino"></td>
            <td><a href="https://www.arduino.cc/">arduino.cc</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>DataDetect</td>
            <td align="center"><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/datadetect.svg" alt="Datadetect"></td>
            <td><a href="https://unifiedglobalarchiving.com/data-detect/">unifiedglobalarchiving.com/data-detect/</a></td>
        </tr>        
        <tr>
            <td>Adopter *</td>
            <td>Sainsbury's</td>
            <td align="center"><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/sainsburys.svg" alt="Sainsbury's"></td>
            <td><a href="https://www.sainsburys.co.uk/">sainsburys.co.uk</a></td>
        </tr>
                <tr>
            <td>Adopter *</td>
            <td>Contraste</td>
            <td align="center"><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/contraste.svg" alt="Contraste"></td>
            <td><a href="https://www.contraste.com/en">contraste.com</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Reyah</td>
            <td align="center"><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/reyah.svg" alt="Reyah"></td>
            <td><a href="https://reyah.eu/">reyah.eu</a></td>
        </tr>        
        <tr>
            <td>Adopter *</td>
            <td>Zero</td>
            <td align="center"><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/commitzero.svg" alt="Project Zero by Commit"></td>
            <td><a href="https://getzero.dev/">getzero.dev</a></td>
        </tr>
        <tr>
            <td>Sponsor</td>
            <td>OrderMyGear</td>
            <td align="center"><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/ordermygear.svg" alt="OrderMyGear"></td>
            <td><a href="https://www.ordermygear.com/">ordermygear.com</a></td>
        </tr>
        <tr>
            <td>Sponsor</td>
            <td>Spiri.bo</td>
            <td align="center"><img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/spiribo.svg" alt="Spiri.bo"></td>
            <td><a href="https://spiri.bo/">spiri.bo</a></td>
        </tr>
    </tbody>
</table>

We also want to thank all individual contributors

<a href="https://opencollective.com/ory" target="_blank"><img src="https://opencollective.com/ory/contributors.svg?width=890&button=false" /></a>

as well as all of our backers

<a href="https://opencollective.com/ory#backers" target="_blank"><img src="https://opencollective.com/ory/backers.svg?width=890"></a>

and past & current supporters (in alphabetical order) on
[Patreon](https://www.patreon.com/_ory): Alexander Alimovs, Billy, Chancy
Kennedy, Drozzy, Edwin Trejos, Howard Edidin, Ken Adler Oz Haven, Stefan Hans,
TheCrealm.

<em>\* Uses one of Ory's major projects in production.</em>

<!--END ADOPTERS-->

### Installation

Head over to the documentation to learn about ways of
[installing ORY Keto](https://www.ory.sh/docs/next/keto/install).

## Ecosystem

<!--BEGIN ECOSYSTEM-->

We build Ory on several guiding principles when it comes to our architecture
design:

- Minimal dependencies
- Runs everywhere
- Scales without effort
- Minimize room for human and network errors

Ory's architecture designed to run best on a Container Orchestration Systems
such as Kubernetes, CloudFoundry, OpenShift, and similar projects. Binaries are
small (5-15MB) and available for all popular processor types (ARM, AMD64, i386)
and operating systems (FreeBSD, Linux, macOS, Windows) without system
dependencies (Java, Node, Ruby, libxml, ...).

### Ory Kratos: Identity and User Infrastructure and Management

[Ory Kratos](https://github.com/ory/kratos) is an API-first Identity and User
Management system that is built according to
[cloud architecture best practices](https://www.ory.sh/docs/next/ecosystem/software-architecture-philosophy).
It implements core use cases that almost every software application needs to
deal with: Self-service Login and Registration, Multi-Factor Authentication
(MFA/2FA), Account Recovery and Verification, Profile and Account Management.

### Ory Hydra: OAuth2 & OpenID Connect Server

[Ory Hydra](https://github.com/ory/hydra) is an OpenID Certified™ OAuth2 and
OpenID Connect Provider which easily connects to any existing identity system by
writing a tiny "bridge" application. Gives absolute control over user interface
and user experience flows.

### Ory Oathkeeper: Identity & Access Proxy

[Ory Oathkeeper](https://github.com/ory/oathkeeper) is a BeyondCorp/Zero Trust
Identity & Access Proxy (IAP) with configurable authentication, authorization,
and request mutation rules for your web services: Authenticate JWT, Access
Tokens, API Keys, mTLS; Check if the contained subject is allowed to perform the
request; Encode resulting content into custom headers (`X-User-ID`), JSON Web
Tokens and more!

### Ory Keto: Access Control Policies as a Server

[Ory Keto](https://github.com/ory/keto) is a policy decision point. It uses a
set of access control policies, similar to AWS IAM Policies, in order to
determine whether a subject (user, application, service, car, ...) is authorized
to perform a certain action on a resource.

<!--END ECOSYSTEM-->

## Security

### Disclosing vulnerabilities

If you think you found a security vulnerability, please refrain from posting it
publicly on the forums, the chat, or GitHub and send us an email to
[hi@ory.am](mailto:hi@ory.am) instead.

## Telemetry

Our services collect summarized, anonymized data which can optionally be turned
off. Click [here](https://www.ory.sh/docs/ecosystem/sqa) to learn more.

### Guide

The Guide is available [here](https://www.ory.sh/docs/next/keto/).

### HTTP API documentation

The HTTP API is documented [here](https://www.ory.sh/docs/next/keto/sdk/api).

### Upgrading and Changelog

New releases might introduce breaking changes. To help you identify and
incorporate those changes, we document these changes in
[UPGRADE.md](./UPGRADE.md) and [CHANGELOG.md](./CHANGELOG.md).

### Command line documentation

Run `keto -h` or `keto help`.

### Develop

We encourage all contributions and encourage you to read our
[contribution guidelines](./CONTRIBUTING.md)

#### Dependencies

You need Go 1.16+ and (for the test suites):

- Docker and Docker Compose
- GNU Make 4.3
- NodeJS / npm@v7

It is possible to develop ORY Keto on Windows, but please be aware that all
guides assume a Unix shell like bash or zsh.

#### Install from source

<pre type="make/command">
make install
</pre>

#### Formatting Code

You can format all code using <code type="make/command">make format</code>. Our
CI checks if your code is properly formatted.

#### Running Tests

There are two types of tests you can run:

- Short tests (do not require a SQL database like PostgreSQL)
- Regular tests (do require PostgreSQL, MySQL, CockroachDB)

##### Short Tests

Short tests run fairly quickly. You can either test all of the code at once

```shell script
go test -short -tags sqlite ./...
```

or test just a specific module:

```shell script
go test -tags sqlite -short ./internal/check/...
```

##### Regular Tests

Regular tests require a database set up. Our test suite is able to work with
docker directly (using [ory/dockertest](https://github.com/ory/dockertest)) but
we encourage to use the script instead. Using dockertest can bloat the number of
Docker Images on your system and starting them on each run is quite slow.
Instead we recommend doing:

```shell
source ./scripts/test-resetdb.sh
go test -tags sqlite ./...
```

##### End-to-End Tests

The e2e tests are part of the normal `go test`. To only run the e2e test, use

```shell
source ./scripts/test-resetdb.sh
go test -tags sqlite ./internal/e2e/...
```

or add the `-short` tag to only test against sqlite in-memory.

#### Build Docker

You can build a development Docker Image using:

<pre type="make/command">
make docker
</pre>
