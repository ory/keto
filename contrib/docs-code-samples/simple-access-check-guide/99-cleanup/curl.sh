#!/bin/bash
set -euo pipefail

curl -X DELETE -G --silent \
     --data-urlencode "subject_id=john" \
     --data-urlencode "relation=decypher" \
     --data-urlencode "namespace=messages" \
     --data-urlencode "object=02y_15_4w350m3" \
     -H 'Content-Type: application/json' \
     http://127.0.0.1:4467/admin/relation-tuples
