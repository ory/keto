#!/bin/bash
set -euo pipefail

export KETO_READ_REMOTE="127.0.0.1:4466"

keto relation-tuple get --namespace chats --object coffee-break --relation member --format json | \
  jq ".relation_tuples[] | .subject_id" -r
