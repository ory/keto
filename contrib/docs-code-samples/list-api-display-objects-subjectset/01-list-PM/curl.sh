#!/bin/bash
set -euo pipefail

curl -G --silent \
     --retry 7 --retry-connrefused \
     --data-urlencode "namespace=Chat" \
     --data-urlencode "relation=member" \
     --data-urlencode "subject_set.namespace=User" \
     --data-urlencode "subject_set.object=PM" \
     --data-urlencode "subject_set.relation=" \
     http://127.0.0.1:4466/relation-tuples | \
  jq ".relation_tuples[] | .object" -r | sort
