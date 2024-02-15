<h1 align="center"><img src="https://raw.githubusercontent.com/ory/meta/master/static/banners/keto.svg" alt="Ory Keto - Open Source & Cloud Native Access Control Server"></h1>

<h4 align="center">    
    <a href="https://www.ory.sh/chat">Chat</a> |
    <a href="https://github.com/ory/keto/discussions">Discusssions</a> |
    <a href="http://eepurl.com/di390P">Newsletter</a><br/><br/>
    <a href="https://www.ory.sh/docs/keto/">Guide</a> |
    <a href="https://www.ory.sh/docs/keto/sdk/api">API Docs</a> |
    <a href="https://godoc.org/github.com/ory/keto">Code Docs</a><br/><br/>
    <a href="https://opencollective.com/ory">Support this project!</a><br/><br/>
    <a href="https://www.ory.sh/jobs/">Work in Open Source, Ory is hiring!</a>
</h4>

---

<p align="left">
    <a href="https://github.com/ory/keto/actions/workflows/ci.yaml"><img src="https://github.com/ory/keto/actions/workflows/ci.yaml/badge.svg?branch=master&event=push" alt="CI Tasks for Ory keto"></a>
    <a href="https://coveralls.io/github/ory/keto?branch=master"> <img src="https://coveralls.io/repos/ory/keto/badge.svg?branch=master&service=github" alt="Coverage Status"></a>
    <a href="https://goreportcard.com/report/github.com/ory/keto"><img src="https://goreportcard.com/badge/github.com/ory/keto" alt="Go Report Card"></a>
    <a href="https://pkg.go.dev/github.com/ory/keto"><img src="https://pkg.go.dev/badge/www.github.com/ory/keto" alt="PkgGoDev"></a>
    <a href="#backers" alt="sponsors on Open Collective"><img src="https://opencollective.com/ory/backers/badge.svg" /></a> <a href="#sponsors" alt="Sponsors on Open Collective"><img src="https://opencollective.com/ory/sponsors/badge.svg" /></a>
    <a href="https://github.com/ory/keto/blob/master/CODE_OF_CONDUCT.md" alt="Ory Code of Conduct"><img src="https://img.shields.io/badge/ory-code%20of%20conduct-green" /></a>
</p>

Ory Keto is the first and most popular open source implementation of "Zanzibar:
Google's Consistent, Global Authorization System"!

## Get Started

You can use
[Docker to run Ory Keto locally](https://www.ory.sh/docs/keto/install) or use
the Ory CLI to try out Ory Keto:

```sh
# This example works best in Bash
bash <(curl https://raw.githubusercontent.com/ory/meta/master/install.sh) -b . ory
sudo mv ./ory /usr/local/bin/

# Or with Homebrew installed
brew install ory/tap/cli
```

create a new project (you may also use
[Docker](https://www.ory.sh/docs/keto/install))

```sh
ory create project --name "Ory Keto Example"
export project_id="{set to the id from output}"
```

and follow the quick & easy steps below.

## Create a namespace with the Ory Permission Language

```sh
# Write a simple configuration with one namespace
echo "class Document implements Namespace {}" > config.ts

# Apply that configuration
ory patch opl --project $project_id -f file://./config.ts

# Create a relationship that grants tom access to a document
echo "Document:secret#read@tom" \
  | ory parse relation-tuples --project=$project_id --format=json - \
  | ory create relation-tuples --project=$project_id -

# List all relationships
ory list relation-tuples --project=$project_id
# Output:
#   NAMESPACE	OBJECT	RELATION NAME	SUBJECT
#   Document	secret	read		tom
#
#   NEXT PAGE TOKEN
#   IS LAST PAGE	true
```

Now, check out your project on the [Ory Network](https://console.ory.sh/) or
continue with a [more in-depth guide](https://www.ory.sh/docs/keto/quickstart).

### Ory Keto on the Ory Network

The [Ory Network](https://www.ory.sh/cloud) is the fastest, most secure and
worry-free way to use Ory's Services. **Ory Permissions** is powered by the Ory
Keto open source permission server, and it's fully API-compatible.

The Ory Network provides the infrastructure for modern end-to-end security:

- Identity & credential management scaling to billions of users and devices
- Registration, Login and Account management flows for passkey, biometric,
  social, SSO and multi-factor authentication
- Pre-built login, registration and account management pages and components
- OAuth2 and OpenID provider for single sign on, API access and
  machine-to-machine authorization
- **Low-latency permission checks based on Google's Zanzibar model and with
  built-in support for the Ory Permission Language**

It's fully managed, highly available, developer & compliance-friendly!

- GDPR-friendly secure storage with data locality
- Cloud-native APIs, compatible with Ory's Open Source servers
- Comprehensive admin tools with the web-based Ory Console and the Ory Command
  Line Interface (CLI)
- Extensive documentation, straightforward examples and easy-to-follow guides
- Fair, usage-based [pricing](https://www.ory.sh/pricing)

Sign up for a
[**free developer account**](https://console.ory.sh/registration?utm_source=github&utm_medium=banner&utm_campaign=keto-readme)
today!

## Ory Network Hybrid Support Plan

Ory offers a support plan for Ory Network Hybrid, including Ory on private cloud
deployments. If you have a self-hosted solution and would like help, consider a
support plan! The team at Ory has years of experience in cloud computing. Ory's
offering is the only official program for qualified support from the
maintainers. For more information see the
**[website](https://www.ory.sh/support/)** or
**[book a meeting](https://www.ory.sh/contact/)**!

## Ory Permissions, Keto and the Google's Zanzibar model

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
something - Ory Permissions and Ory Keto are the right fit for you.

Currently, Ory Permissions [on the Ory Network] and the open-source Ory Keto
permission server implement the API contracts for managing and checking
relations ("permissions") with HTTP and gRPC APIs, as well as global rules
defined through the Ory Permission Language ("userset rewrites"). Future
versions will include features such as Zookies, reverse permission lookups, and
more.

---

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Ory Permissions, Keto and the Google's Zanzibar model](#ory-permissions-keto-and-the-googles-zanzibar-model)
- [Who's Using It?](#whos-using-it)
  - [Installation](#installation)
- [Ecosystem](#ecosystem)
  - [Ory Kratos: Identity and User Infrastructure and Management](#ory-kratos-identity-and-user-infrastructure-and-management)
  - [Ory Hydra: OAuth2 & OpenID Connect Server](#ory-hydra-oauth2--openid-connect-server)
  - [Ory Oathkeeper: Identity & Access Proxy](#ory-oathkeeper-identity--access-proxy)
  - [Ory Keto: Access Control Policies as a Server](#ory-keto-access-control-policies-as-a-server)
- [Security](#security)
  - [Disclosing Vulnerabilities](#disclosing-vulnerabilities)
- [Telemetry](#telemetry)
  - [Guide](#guide)
  - [HTTP API Documentation](#http-api-documentation)
  - [Upgrading and Changelog](#upgrading-and-changelog)
  - [Command Line Documentation](#command-line-documentation)
  - [Develop](#develop)
    - [Dependencies](#dependencies)
    - [Install From Source](#install-from-source)
    - [Formatting Code](#formatting-code)
    - [Running Tests](#running-tests)
      - [Short Tests](#short-tests)
      - [Regular Tests](#regular-tests)
      - [End-to-End Tests](#end-to-end-tests)
    - [Build Docker](#build-docker)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Who's Using It?

<!--BEGIN ADOPTERS-->

The Ory community stands on the shoulders of individuals, companies, and
maintainers. The Ory team thanks everyone involved - from submitting bug reports
and feature requests, to contributing patches and documentation. The Ory
community counts more than 33.000 members and is growing rapidly. The Ory stack
protects 60.000.000.000+ API requests every month with over 400.000+ active
service nodes. None of this would have been possible without each and everyone
of you!

The following list represents companies that have accompanied us along the way
and that have made outstanding contributions to our ecosystem. _If you think
that your company deserves a spot here, reach out to
<a href="mailto:office@ory.sh">office@ory.sh</a> now_!

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
            <td>Adopter *</td>
            <td>Raspberry PI Foundation</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/raspi.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/raspi.svg" alt="Raspberry PI Foundation">
                </picture>
            </td>
            <td><a href="https://www.raspberrypi.org/">raspberrypi.org</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Kyma Project</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/kyma.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/kyma.svg" alt="Kyma Project">
                </picture>
            </td>
            <td><a href="https://kyma-project.io">kyma-project.io</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Tulip</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/tulip.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/tulip.svg" alt="Tulip Retail">
                </picture>
            </td>
            <td><a href="https://tulip.com/">tulip.com</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Cashdeck / All My Funds</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/allmyfunds.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/allmyfunds.svg" alt="All My Funds">
                </picture>
            </td>
            <td><a href="https://cashdeck.com.au/">cashdeck.com.au</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Hootsuite</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/hootsuite.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/hootsuite.svg" alt="Hootsuite">
                </picture>
            </td>
            <td><a href="https://hootsuite.com/">hootsuite.com</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Segment</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/segment.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/segment.svg" alt="Segment">
                </picture>
            </td>
            <td><a href="https://segment.com/">segment.com</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Arduino</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/arduino.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/arduino.svg" alt="Arduino">
                </picture>
            </td>
            <td><a href="https://www.arduino.cc/">arduino.cc</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>DataDetect</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/datadetect.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/datadetect.svg" alt="Datadetect">
                </picture>
            </td>
            <td><a href="https://unifiedglobalarchiving.com/data-detect/">unifiedglobalarchiving.com/data-detect/</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Sainsbury's</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/sainsburys.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/sainsburys.svg" alt="Sainsbury's">
                </picture>
            </td>
            <td><a href="https://www.sainsburys.co.uk/">sainsburys.co.uk</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Contraste</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/contraste.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/contraste.svg" alt="Contraste">
                </picture>
            </td>
            <td><a href="https://www.contraste.com/en">contraste.com</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Reyah</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/reyah.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/reyah.svg" alt="Reyah">
                </picture>
            </td>
            <td><a href="https://reyah.eu/">reyah.eu</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Zero</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/commitzero.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/commitzero.svg" alt="Project Zero by Commit">
                </picture>
            </td>
            <td><a href="https://getzero.dev/">getzero.dev</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Padis</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/padis.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/padis.svg" alt="Padis">
                </picture>
            </td>
            <td><a href="https://padis.io/">padis.io</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Cloudbear</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/cloudbear.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/cloudbear.svg" alt="Cloudbear">
                </picture>
            </td>
            <td><a href="https://cloudbear.eu/">cloudbear.eu</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Security Onion Solutions</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/securityonion.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/securityonion.svg" alt="Security Onion Solutions">
                </picture>
            </td>
            <td><a href="https://securityonionsolutions.com/">securityonionsolutions.com</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Factly</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/factly.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/factly.svg" alt="Factly">
                </picture>
            </td>
            <td><a href="https://factlylabs.com/">factlylabs.com</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Nortal</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/nortal.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/nortal.svg" alt="Nortal">
                </picture>
            </td>
            <td><a href="https://nortal.com/">nortal.com</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>OrderMyGear</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/ordermygear.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/ordermygear.svg" alt="OrderMyGear">
                </picture>
            </td>
            <td><a href="https://www.ordermygear.com/">ordermygear.com</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Spiri.bo</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/spiribo.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/spiribo.svg" alt="Spiri.bo">
                </picture>
            </td>
            <td><a href="https://spiri.bo/">spiri.bo</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Strivacity</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/strivacity.svg" />
                    <img height="16px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/strivacity.svg" alt="Spiri.bo">
                </picture>
            </td>
            <td><a href="https://strivacity.com/">strivacity.com</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Hanko</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/hanko.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/hanko.svg" alt="Hanko">
                </picture>
            </td>
            <td><a href="https://hanko.io/">hanko.io</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Rabbit</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/rabbit.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/rabbit.svg" alt="Rabbit">
                </picture>
            </td>
            <td><a href="https://rabbit.co.th/">rabbit.co.th</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>inMusic</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/inmusic.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/inmusic.svg" alt="InMusic">
                </picture>
            </td>
            <td><a href="https://inmusicbrands.com/">inmusicbrands.com</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Buhta</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/buhta.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/buhta.svg" alt="Buhta">
                </picture>
            </td>
            <td><a href="https://buhta.com/">buhta.com</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Connctd</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/connctd.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/connctd.svg" alt="Connctd">
                </picture>
            </td>
            <td><a href="https://connctd.com/">connctd.com</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>Paralus</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/paralus.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/paralus.svg" alt="Paralus">
                </picture>
            </td>
            <td><a href="https://www.paralus.io/">paralus.io</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>TIER IV</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/tieriv.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/tieriv.svg" alt="TIER IV">
                </picture>
            </td>
            <td><a href="https://tier4.jp/en/">tier4.jp</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>R2Devops</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/r2devops.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/r2devops.svg" alt="R2Devops">
                </picture>
            </td>
            <td><a href="https://r2devops.io/">r2devops.io</a></td>
        </tr>
        <tr>
            <td>Adopter *</td>
            <td>LunaSec</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/lunasec.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/lunasec.svg" alt="LunaSec">
                </picture>
            </td>
            <td><a href="https://www.lunasec.io/">lunasec.io</a></td>
        </tr>
            <tr>
            <td>Adopter *</td>
            <td>Serlo</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/serlo.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/serlo.svg" alt="Serlo">
                </picture>
            </td>
            <td><a href="https://serlo.org/">serlo.org</a></td>
        </tr>
        </tr>
            <tr>
            <td>Adopter *</td>
            <td>dyrector.io</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/dyrector_io.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/dyrector_io.svg" alt="dyrector.io">
                </picture>
            </td>
            <td><a href="https://dyrector.io/">dyrector.io</a></td>
        </tr>
        </tr>
            <tr>
            <td>Adopter *</td>
            <td>Stackspin</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/stackspin.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/stackspin.svg" alt="stackspin.net">
                </picture>
            </td>
            <td><a href="https://www.stackspin.net/">stackspin.net</a></td>
        </tr>
        </tr>
            <tr>
            <td>Adopter *</td>
            <td>Amplitude</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/amplitude.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/amplitude.svg" alt="amplitude.com">
                </picture>
            </td>
            <td><a href="https://amplitude.com/">amplitude.com</a></td>
        </tr>
         <tr>
            <td>Adopter *</td>
            <td>Pinniped</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/pinniped.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/pinniped.svg" alt="pinniped.dev">
                </picture>
            </td>
            <td><a href="https://pinniped.dev/">pinniped.dev</a></td>
        </tr>         
        <tr>
            <td>Adopter *</td>
            <td>Pvotal</td>
            <td align="center">
                <picture>
                    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/ory/meta/master/static/adopters/light/pvotal.svg" />
                    <img height="32px" src="https://raw.githubusercontent.com/ory/meta/master/static/adopters/dark/pvotal.svg" alt="pvotal.tech">
                </picture>
            </td>
            <td><a href="https://pvotal.tech/">pvotal.tech</a></td>
        </tr>
    </tbody>
</table>

Many thanks to all individual contributors

<a href="https://opencollective.com/ory" target="_blank"><img src="https://opencollective.com/ory/contributors.svg?width=890&limit=714&button=false" /></a>

<em>\* Uses one of Ory's major projects in production.</em>

<!--END ADOPTERS-->

## Installation

Head over to the documentation to learn about ways of
[installing Ory Keto](https://www.ory.sh/docs/keto/install).

## Ecosystem

<!--BEGIN ECOSYSTEM-->

We build Ory on several guiding principles when it comes to our architecture
design:

- Minimal dependencies
- Runs everywhere
- Scales without effort
- Minimize room for human and network errors

Ory's architecture is designed to run best on a Container Orchestration system
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
(MFA/2FA), Account Recovery and Verification, Profile, and Account Management.

### Ory Hydra: OAuth2 & OpenID Connect Server

[Ory Hydra](https://github.com/ory/hydra) is an OpenID Certified™ OAuth2 and
OpenID Connect Provider which easily connects to any existing identity system by
writing a tiny "bridge" application. It gives absolute control over the user
interface and user experience flows.

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

### Disclosing Vulnerabilities

If you think you found a security vulnerability, please refrain from posting it
publicly on the forums, the chat, or GitHub. You can find all info for
responsible disclosure in our
[security.txt](https://www.ory.sh/.well-known/security.txt).

## Telemetry

Our services collect summarized, anonymized data which can optionally be turned
off. Click [here](https://www.ory.sh/docs/ecosystem/sqa) to learn more.

### Guide

The Guide is available [here](https://www.ory.sh/docs/keto/).

### HTTP API Documentation

The HTTP API is documented [here](https://www.ory.sh/docs/keto/sdk/api).

### Upgrading and Changelog

New releases might introduce breaking changes. To help you identify and
incorporate those changes, we document these changes in
[UPGRADE.md](./UPGRADE.md) and [CHANGELOG.md](./CHANGELOG.md).

### Command Line Documentation

Run `keto -h` or `keto help`.

### Develop

We encourage all contributions and recommend you read our
[contribution guidelines](./CONTRIBUTING.md).

#### Dependencies

You need Go 1.19+ and (for the test suites):

- Docker and Docker Compose
- GNU Make 4.3
- NodeJS / npm >= v7

It is possible to develop Ory Keto on Windows, but please be aware that all
guides assume a Unix shell like bash or zsh.

#### Install From Source

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

Short tests run fairly quickly. You can either test all of the code at once:

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

The e2e tests are part of the normal `go test`. To only run the e2e test, use:

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
