#!/bin/bash
set -euo pipefail

diff <("$(dirname "$0")"/cli.sh) <(echo "true")
echo "$0 PASSED"
