#!/bin/bash
set -euo pipefail

export KETO_READ_REMOTE="127.0.0.1:4466"
export KETO_WRITE_REMOTE="127.0.0.1:4467"

keto relation-tuple get --namespace files --format json | \
  jq ".relation_tuples" | \
    keto relation-tuple delete - -q > /dev/null

keto relation-tuple get --namespace directories --format json | \
  jq ".relation_tuples" | \
    keto relation-tuple delete - -q > /dev/null
