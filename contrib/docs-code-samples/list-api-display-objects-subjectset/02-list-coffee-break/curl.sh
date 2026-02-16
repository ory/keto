#!/bin/bash
set -euo pipefail

curl -G --silent \
     --retry 7 --retry-connrefused \
     --data-urlencode "namespace=Chat" \
     --data-urlencode "object=coffee-break" \
     --data-urlencode "relation=member" \
     http://127.0.0.1:4466/relation-tuples | \
  jq "[.relation_tuples[] | .subject_set.object] | sort | .[]" -r
