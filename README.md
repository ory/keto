# ORY Keto

<h4 align="center">
    <a href="https://discord.gg/PAMQWkr">Chat</a> |
    <a href="https://community.ory.am/">Forums</a> |
    <a href="http://eepurl.com/bKT3N9">Newsletter</a><br/><br/>
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
- [Ecosystem](#ecosystem)
  - [ORY Security Console: Administrative User Interface](#ory-security-console-administrative-user-interface)
  - [ORY Hydra: OAuth2 & OpenID Connect Server](#ory-hydra-oauth2--openid-connect-server)
  - [ORY Oathkeeper: Identity & Access Proxy](#ory-oathkeeper-identity--access-proxy)
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

## Installation

The easiest way to install ORY Keto is using the [Docker Hub Image](https://hub.docker.com/r/oryd/keto/):

```
docker run oryd/keto:<version> help
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
[here](https://www.ory.sh/docs/1-hydra/9-telemetry) to learn more.

### Guide

The Guide is available [here](https://www.ory.sh/docs/3-keto/).

### HTTP API documentation

The HTTP API is documented [here](https://www.ory.sh/docs/api/keto).

### Upgrading and Changelog

New releases might introduce breaking changes. To help you identify and incorporate those changes, we document these
changes in [UPGRADE.md](./UPGRADE.md) and [CHANGELOG.md](./CHANGELOG.md).

### Command line documentation

Run `hydra -h` or `hydra help`.
