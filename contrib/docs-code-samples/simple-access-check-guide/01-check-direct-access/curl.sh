#!/bin/bash
set -euo pipefail

curl -G --silent \
     --data-urlencode "subject_id=john" \
     --data-urlencode "relation=decypher" \
     --data-urlencode "namespace=messages" \
     --data-urlencode "object=02y_15_4w350m3" \
     http://127.0.0.1:4466/relation-tuples/check \
  | jq -r 'if .allowed == true then "Allowed" else "Denied" end'
