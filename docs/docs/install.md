---
id: install
title: Installation
---

Installing Ory Keto on any system is straightforward. We provide pre-built
binaries, Docker images, and support a number of package managers.

## Docker

We recommend using Docker to run Ory Keto:

```shell
$ docker pull oryd/keto:v0.7.0-alpha.0.pre.5
$ docker run --rm -it oryd/keto:v0.7.0-alpha.0.pre.5 help
```

## macOS

You can install Ory Keto using [homebrew](https://brew.sh/) on macOS:

```shell
$ brew tap ory/keto
$ brew install ory/keto/keto
$ keto help
```

## Linux

On linux, you can use `bash <(curl ...)` to fetch the latest stable binary
using:

```shell
$ bash <(curl https://raw.githubusercontent.com/ory/keto/master/install.sh) -b . v0.7.0-alpha.0.pre.5
$ ./keto help
```

You may want to move Ory Keto into your `$PATH`:

```shell
$ sudo mv ./keto /usr/local/bin/
$ keto help
```

## Windows

You can install Ory Keto using [scoop](https://scoop.sh) on Windows:

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

If you wish to compile Ory Keto yourself, you need to install and set up
[Go 1.16+](https://golang.org/) and add `$GOPATH/bin` to your `$PATH`.

The following commands check out the latest release tag of Ory Keto, compile it,
and set up flags so that `keto version` works as expected. Please note that this
will only work with a Linux shell like bash or sh.

```shell
$ git clone https://github.com/ory/keto.git
$ cd keto
$ git checkout v0.7.0-alpha.0.pre.5
$ make install
$ keto help
```
