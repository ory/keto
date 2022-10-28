#!/bin/env bash
set -euo pipefail

source ./scripts/install-helpers.sh

PROTOBUF_JS_VERSION=3.21.2
CHECKSUM=''

usage() {
	cat <<EOF
Install protoc $PROTOBUF_JS_VERSION from the official release page.

Env Vars:
  BINDIR  - The directory to install the binary to. Defaults to ./.bin

EOF
	exit 2
}

execute() {
	tmpdir=$(mktemp -d)
	echo "downloading files into ${tmpdir}"
	http_download "${tmpdir}/${TARBALL}" "${TARBALL_URL}" ""
	srcdir="${tmpdir}/bin"
	(cd "${tmpdir}" && untar "${TARBALL}")
	test ! -d "${BINDIR}" && install -d "${BINDIR}"

  install "${srcdir}/${BINNAME}" "${BINDIR}/"
  echo "installed protoc to ${BINDIR}"
	rm -rf "${tmpdir}"
}

OWNER=protocolbuffers
REPO=protobuf-javascript

OS=$(uname_os)
ARCH=$(uname_arch)

case "${OS}" in
darwin) OS=osx ;;
esac

case "${ARCH}" in
arm64) ARCH=aarch_64 ;;
amd64) ARCH=x86_64 ;;
esac

GITHUB_DOWNLOAD=https://github.com/${OWNER}/${REPO}/releases/download
BINDIR=${BINDIR:-./.bin}
NAME="protobuf-javascript-${PROTOBUF_JS_VERSION}-${OS}-${ARCH}"
TARBALL="${NAME}.zip"
TARBALL_URL=${GITHUB_DOWNLOAD}/v${PROTOBUF_JS_VERSION}/${TARBALL}

BINNAME="protoc-gen-js"
if [ "${OS}" = "windows" ]; then
    BINNAME="protoc-gen-js.exe"
fi

if check_binary_version "$BINDIR" "$BINNAME" "$PROTOBUF_JS_VERSION"; then
    echo "protoc-gen-js ${PROTOBUF_JS_VERSION} already installed"
    exit 0
fi

execute

update_binary_version "$BINDIR" "$BINNAME" "$PROTOBUF_JS_VERSION"
