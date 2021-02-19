#!/bin/bash
set -euo pipefail

diff <("$(dirname "$0")"/curl.sh 2> /dev/null) <(echo '"allowed" 200')
echo "$0 PASSED"
