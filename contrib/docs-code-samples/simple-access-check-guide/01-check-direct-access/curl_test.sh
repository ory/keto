#!/bin/bash
set -euo pipefail

diff <("$(dirname "$0")"/curl.sh 2> /dev/null) <(echo -e "200\nallowed")
echo "$0 PASSED"
