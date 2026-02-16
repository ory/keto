#!/bin/bash
set -euo pipefail

export KETO_READ_REMOTE="127.0.0.1:4466"

keto relation-tuple get --namespace Chat --relation member --subject-set User:PM --format json --insecure-disable-transport-security | \
  jq ".relation_tuples[] | .object" -r | sort
