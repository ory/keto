#!/bin/bash
set -euo pipefail

diff <("$(dirname "$0")"/cli.sh) <(echo "messages:02y_15_4w350m3#decypher@john")
echo "CLI PASSED"
