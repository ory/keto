#!/bin/bash
set -euo pipefail

curl -G --silent \
     --data-urlencode "namespace=files" \
     --data-urlencode "relation=access" \
     --data-urlencode "object=/photos/beach.jpg" \
     --data-urlencode "max-depth=3" \
     http://127.0.0.1:4466/relation-tuples/expand | \
  jq
