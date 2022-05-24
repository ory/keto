#!/usr/bin/env bash
set -euo pipefail

for i in {6..500} ; do
  for j in {1..500} ; do
      for k in {1..4}; do
        ./keto relation-tuple create <(echo '{"namespace": "default", "object": "'$i'", "subject_id": "'$j'", "relation": "'$k'"}')
      done
  done
done
