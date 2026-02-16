#!/bin/bash
set -euo pipefail

export KETO_READ_REMOTE="127.0.0.1:4466"

keto relation-tuple get --namespace Chat --object coffee-break --relation member --format json --insecure-disable-transport-security | \
  jq "[.relation_tuples[] | .subject_set.object] | sort | .[]" -r
