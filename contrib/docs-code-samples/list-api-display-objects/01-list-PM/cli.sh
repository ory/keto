#!/bin/bash
set -euo pipefail

export KETO_READ_REMOTE="127.0.0.1:4466"

keto relation-tuple get --namespace chats --relation member --subject-id PM --format json --insecure-disable-transport-security | \
  jq ".relation_tuples[] | .object" -r | sort
