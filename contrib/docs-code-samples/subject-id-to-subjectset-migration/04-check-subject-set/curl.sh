#!/bin/bash
set -euo pipefail

curl -G --silent \
     --retry 7 --retry-connrefused \
     --data-urlencode "namespace=File" \
     --data-urlencode "object=data.txt" \
     --data-urlencode "relation=viewer" \
     --data-urlencode "subject_set.namespace=User" \
     --data-urlencode "subject_set.object=alice" \
     --data-urlencode "subject_set.relation=" \
     http://127.0.0.1:4466/relation-tuples/check \
  | jq -r 'if .allowed == true then "Allowed" else "Denied" end'
