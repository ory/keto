#!/usr/bin/env bash
set -euo pipefail

source ./scripts/install-helpers.sh

BINDIR=${BINDIR:-./.bin}
TRIVY_VERSION="0.32.1"

OS=$(uname_os)
BINNAME="trivy"
if [ "${OS}" = "windows" ]; then
    BINNAME="trivy.exe"
fi

if [[ "$("$BINDIR/$BINNAME" version)" == *"$TRIVY_VERSION"* ]]; then
    echo "trivy ${TRIVY_VERSION} already installed"
    exit 0
fi

curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | bash -s -- -b "$BINDIR" "v$TRIVY_VERSION"
