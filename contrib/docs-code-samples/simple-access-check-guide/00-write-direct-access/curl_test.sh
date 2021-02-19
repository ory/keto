#!/bin/bash
set -euo pipefail

diff <("$(dirname "$0")"/curl.sh 2> /dev/null) <(echo "Created 201")
echo "$0 PASSED"
