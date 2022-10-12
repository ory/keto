#!/usr/bin/env bash
set -euo pipefail

source ./scripts/install-helpers.sh

BINDIR=${BINDIR:-./.bin}
GRYPE_VERSION="0.50.2"

OS=$(uname_os)
BINNAME="grype"
if [ "${OS}" = "windows" ]; then
    BINNAME="grype.exe"
fi

if [[ "$("$BINDIR/$BINNAME" version)" == *"$GRYPE_VERSION"* ]]; then
    echo "grype ${GRYPE_VERSION} already installed"
    exit 0
fi

curl -sSfL https://raw.githubusercontent.com/anchore/grype/main/install.sh | bash -s -- -b "$BINDIR" "v$GRYPE_VERSION"
