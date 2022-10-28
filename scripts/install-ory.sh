#!/usr/bin/env bash
set -euo pipefail

source ./scripts/install-helpers.sh

BINDIR=${BINDIR:-./.bin}
ORY_VERSION="0.1.44"

OS=$(uname_os)
BINNAME="ory"
if [ "${OS}" = "windows" ]; then
    BINNAME="ory.exe"
fi

if check_binary_lock "$BINDIR" "$BINNAME"; then
    echo "ory ${ORY_VERSION} already installed"
    exit 0
fi

curl -sSfL https://raw.githubusercontent.com/ory/meta/master/install.sh | bash -s -- -b "$BINDIR" ory "v$ORY_VERSION"

regenerate_lockfile "$BINDIR"
