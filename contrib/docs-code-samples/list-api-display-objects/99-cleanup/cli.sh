#!/bin/bash
set -euo pipefail

export KETO_READ_REMOTE="127.0.0.1:4466"
export KETO_WRITE_REMOTE="127.0.0.1:4467"

keto relation-tuple get --namespace chats --format json --insecure-disable-transport-security | \
  jq ".relation_tuples" | \
    keto relation-tuple delete - -q > /dev/null --insecure-disable-transport-security
