#!/bin/bash
set -euo pipefail

source ./scripts/install-helpers.sh

SWAGGER_VERSION=0.30.3

usage() {
	cat <<EOF
Install swagger $SWAGGER_VERSION from the official release page.

Env Vars:
  BINDIR  - The directory to install the binary to. Defaults to ./.bin

EOF
	exit 2
}

execute() {
	echo "downloading file ${BINDIR}"
	test ! -d "${BINDIR}" && install -d "${BINDIR}"
	http_download "${BINDIR}/${BINNAME}" "${BINARY_URL}" ""
	chmod +x "${BINDIR}/${BINNAME}"
  echo "installed swagger to ${BINDIR}"
}

OWNER=go-swagger
REPO=go-swagger

OS=$(uname_os)
ARCH=$(uname_arch)

GITHUB_DOWNLOAD=https://github.com/${OWNER}/${REPO}/releases/download
BINDIR=${BINDIR:-./.bin}
NAME="swagger_${OS}_${ARCH}"
BINARY_URL=${GITHUB_DOWNLOAD}/v${SWAGGER_VERSION}/${NAME}

BINNAME="swagger"
if [ "${OS}" = "windows" ]; then
    BINNAME="swagger.exe"
fi

if [[ "$("$BINDIR/$BINNAME" version)" == *"$SWAGGER_VERSION"* ]]; then
    echo "swagger ${SWAGGER_VERSION} already installed"
    exit 0
fi

execute
