#!/bin/bash
set -euo pipefail

curl -G \
     --data-urlencode "subject=john" \
     --data-urlencode "relation=decypher" \
     --data-urlencode "namespace=messages" \
     --data-urlencode "object=02y_15_4w350m3" \
     -w " %{response_code}\n" \
     http://127.0.0.1:4466/check

# Expected Output:
#   "allowed" 200
