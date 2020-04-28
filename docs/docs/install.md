---
id: install
title: Installation
---

Installing ORY Keto on any system is straightforward. We provide pre-built
binaries, Docker images, and support a number of package managers.

## Docker

We recommend using Docker to run ORY Keto:

```shell
$ docker pull oryd/keto:v0.4.4-alpha.1
$ docker run --rm -it oryd/keto:v0.4.4-alpha.1 help
```

## macOS

You can install ORY Keto using [homebrew](https://brew.sh/) on macOS:

```shell
$ brew tap ory/keto
$ brew install ory/keto/keto
$ keto help
```

## Linux

On linux, you can use `bash <(curl ...)` to fetch the latest stable binary
using:

```shell
$ bash <(curl https://raw.githubusercontent.com/ory/keto/master/install.sh) -b . v0.4.4-alpha.1
$ ./keto help
```

You may want to move ORY Keto into your `$PATH`:

```shell
$ sudo mv ./keto /usr/local/bin/
$ keto help
```

## Windows

You can install ORY Keto using [scoop](https://scoop.sh) on Windows:

```shell
> scoop bucket add ory-keto https://github.com/ory/scoop-keto.git
> scoop install keto
> keto help
```

## Download Binaries

The client and server **binaries are downloadable via
[GitHub releases](https://github.com/ory/keto/releases)**. There is currently no
installer available. You have to add the Keto binary to the PATH environment
variable yourself or put the binary in a location that is already in your
`$PATH`, for example `/usr/local/bin`.

Once installed, you should be able to run:

```shell
$ keto help
```

## Building from Source

If you wish to compile ORY Keto yourself, you need to install and set up
[Go 1.12+](https://golang.org/) and add `$GOPATH/bin` to your `$PATH`.

The following commands check out the latest release tag of ORY Keto, compile it,
and set up flags so that `keto version` works as expected. Please note that this
will only work with a Linux shell like bash or sh.

```shell
$ go get -d -u github.com/ory/keto
$ cd $(go env GOPATH)/src/github.com/ory/keto
$ GO111MODULE=on make install-stable
$ $(go env GOPATH)/bin/keto help
```
