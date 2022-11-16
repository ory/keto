#!/usr/bin/env bash
set -euo pipefail

source ./scripts/install-helpers.sh

BINDIR=${BINDIR:-./.bin}
LICENCES_VERSION="0.1.48"

OS=$(uname_os)
BINNAME="licences"
if [ "${OS}" = "windows" ]; then
    BINNAME="licences.exe"
fi

if [[ "$("$BINDIR/$BINNAME" version)" == *"$LICENCES_VERSION"* ]]; then
    echo "licences ${LICENCES_VERSION} already installed"
    exit 0
fi

curl -sSfL https://raw.githubusercontent.com/ory/ci/master/licenses/install | bash -s -- -b "$BINDIR" licences "v$LICENCES_VERSION"
