#!/bin/bash
set -euo pipefail

diff <("$(dirname "$0")"/curl.sh 2> /dev/null) <(echo "201 Created!")
echo "$0 PASSED"
