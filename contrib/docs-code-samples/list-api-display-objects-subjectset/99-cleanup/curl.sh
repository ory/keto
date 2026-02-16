#!/bin/bash
set -euo pipefail

export KETO_READ_REMOTE="127.0.0.1:4466"
export KETO_WRITE_REMOTE="127.0.0.1:4467"

curl -G --silent \
     --retry 7 --retry-connrefused \
     --data-urlencode "namespace=Chat" \
     http://127.0.0.1:4466/relation-tuples | \
  jq "[ .relation_tuples[] | { relation_tuple: . , action: \"delete\" } ]" -c | \
    curl -X PATCH --silent --fail \
      -H 'Content-Type: application/json' \
      --retry 7 --retry-connrefused \
      --data @- \
      http://127.0.0.1:4467/admin/relation-tuples
