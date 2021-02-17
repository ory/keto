#!/bin/bash
set -euo pipefail

curl -G \
     --data-urlencode "subject=john" \
     --data-urlencode "relation=decypher" \
     --data-urlencode "namespace=message" \
     --data-urlencode "object=02y_15_4w350m3" \
     -w "%{http_code}" \
     http://127.0.0.1:4466/check

# Expected Output: 200 allowed
