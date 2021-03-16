#!/bin/bash
set -euo pipefail

export KETO_READ_REMOTE="127.0.0.1:4466"

keto relation-tuple get chats --relation member --subject PM --format json | \
  jq ".relation_tuples[] | .object" -r
