<h1 align="center"><img src="./docs/images/banner_keto.png" alt="ORY Keto - Open Source & Cloud Native Access Control Server"></h1>

<h4 align="center">    
    <a href="https://discord.gg/PAMQWkr">Chat</a> |
    <a href="https://community.ory.am/">Forums</a> |
    <a href="http://eepurl.com/di390P">Newsletter</a><br/><br/>
    <a href="https://www.ory.sh/docs/guides/master/keto/">Guide</a> |
    <a href="https://www.ory.sh/docs/api/keto?version=master">API Docs</a> |
    <a href="https://godoc.org/github.com/ory/keto">Code Docs</a><br/><br/>
    <a href="https://opencollective.com/ory-hydra">Support this project!</a>
</h4>

This service is a policy decision point. It uses a set of access control policies, similar to
[AWS IAM Policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html), in order to determine whether
a subject (user, application, service, car, ...) is authorized to perform a certain action on a resource.

<p align="left">
    <a href="https://circleci.com/gh/ory/keto/tree/master"><img src="https://circleci.com/gh/ory/keto/tree/master.svg?style=shield" alt="Build Status"></a>
    <a href="https://coveralls.io/github/ory/keto?branch=master"><img src="https://coveralls.io/repos/ory/keto/badge.svg?branch=master&service=github" alt="Coverage Status"></a>
    <a href="https://goreportcard.com/report/github.com/ory/keto"><img src="https://goreportcard.com/badge/github.com/ory/keto" alt="Go Report Card"></a>
</p>


---

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Introduction](#introduction)
  - [Installation](#installation)
    - [Download binaries](#download-binaries)
    - [Using Docker](#using-docker)
    - [Building from source](#building-from-source)
- [Ecosystem](#ecosystem)
  - [ORY Security Console: Administrative User Interface](#ory-security-console-administrative-user-interface)
  - [ORY Hydra: OAuth2 & OpenID Connect Server](#ory-hydra-oauth2-&-openid-connect-server)
  - [ORY Oathkeeper: Identity & Access Proxy](#ory-oathkeeper-identity-&-access-proxy)
- [Security](#security)
  - [Disclosing vulnerabilities](#disclosing-vulnerabilities)
- [Telemetry](#telemetry)
  - [Guide](#guide)
  - [HTTP API documentation](#http-api-documentation)
  - [Upgrading and Changelog](#upgrading-and-changelog)
  - [Command line documentation](#command-line-documentation)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Introduction

With ORY Keto, you can model Access Control Lists, Role Based Access Control, and fine-grained permission sets.
This server implementation uses [ORY Ladon](https://github.com/ory/ladon) as the decision engine.

ORY Keto is possible to resolve credentials using various authentication mechanisms:

* OAuth 2.0 Access Tokens using the OAuth 2.0 Introspection standard.
* Plaintext when you already know the user ID.
* JSON Web Tokens (coming soon).
* SAML (coming soon).

### Installation

There are various ways of installing ORY keto on your system.

#### Download binaries

The client and server **binaries are downloadable at [releases](https://github.com/ory/keto/releases)**.
There is currently no installer available. You have to add the ORY keto binary to the PATH environment variable yourself or put
the binary in a location that is already in your path (`/usr/bin`, ...).
If you do not understand what that all of this means, ask in our [chat channel](https://www.ory.sh/chat). We are happy to help.

#### Using Docker

**Starting the host** is easiest with docker. The host process handles HTTP requests and is backed by a database.
Read how to install docker on [Linux](https://docs.docker.com/linux/), [OSX](https://docs.docker.com/mac/) or
[Windows](https://docs.docker.com/windows/). ORY keto is available on [Docker Hub](https://hub.docker.com/r/oryd/keto/).

You can use ORY keto without a database, but be aware that restarting, scaling
or stopping the container will **lose all data**:

```
$ docker run -e "DATABASE_URL=memory" -d --name my-keto -p 4466:4466 oryd/keto
ec91228cb105db315553499c81918258f52cee9636ea2a4821bdb8226872f54b
```

#### Building from source

If you wish to compile ORY keto yourself, you need to install and set up [Go 1.10+](https://golang.org/) and add `$GOPATH/bin`
to your `$PATH` as well as [golang/dep](http://github.com/golang/dep).

The following commands will check out the latest release tag of ORY keto and compile it and set up flags so that `keto version`
works as expected. Please note that this will only work with a linux shell like bash or sh.

```
go get -d -u github.com/ory/keto
cd $(go env GOPATH)/src/github.com/ory/keto
keto_LATEST=$(git describe --abbrev=0 --tags)
git checkout $keto_LATEST
dep ensure -vendor-only
go install \
    -ldflags "-X github.com/ory/keto/cmd.Version=$keto_LATEST -X github.com/ory/keto/cmd.BuildTime=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` -X github.com/ory/keto/cmd.GitHash=`git rev-parse HEAD`" \
    github.com/ory/keto
git checkout master
keto help
```

## Ecosystem

<a href="https://console.ory.am/auth/login">
    <img align="right" width="30%" src="docs/images/sec-console.png" alt="ORY Security Console">
</a>

### ORY Security Console: Administrative User Interface

The [ORY Security Console](https://console.ory.am/auth/login) is a visual admin interface for managing ORY Hydra,
ORY Oathkeeper, and ORY Keto.

### ORY Hydra: OAuth2 & OpenID Connect Server

[ORY Hydra](https://github.com/ory/hydra) ORY Hydra is a hardened OAuth2 and OpenID Connect server optimized
for low-latency, high throughput, and low resource consumption. ORY Hydra is not an identity provider
(user sign up, user log in, password reset flow), but connects to your existing identity provider through a consent app.

### ORY Oathkeeper: Identity & Access Proxy

[ORY Oathkeeper](https://github.com/ory/oathkeeper) is a BeyondCorp/Zero Trust Identity & Access Proxy (IAP) built
on top of OAuth2 and ORY Hydra.

## Security

### Disclosing vulnerabilities

If you think you found a security vulnerability, please refrain from posting it publicly on the forums, the chat, or GitHub
and send us an email to [hi@ory.am](mailto:hi@ory.am) instead.

## Telemetry

Our services collect summarized, anonymized data which can optionally be turned off. Click
[here](https://www.ory.sh/docs/guides/master/telemetry/) to learn more.

### Guide

The Guide is available [here](https://www.ory.sh/docs/guides/master/keto/).

### HTTP API documentation

The HTTP API is documented [here](https://www.ory.sh/docs/api/keto?version=master).

### Upgrading and Changelog

New releases might introduce breaking changes. To help you identify and incorporate those changes, we document these
changes in [UPGRADE.md](./UPGRADE.md) and [CHANGELOG.md](./CHANGELOG.md).

### Command line documentation

Run `keto -h` or `keto help`.
