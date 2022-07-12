#!/usr/bin/env bash
set -euo pipefail

for i in {1..500} ; do
  d="$(mktemp -d)"
  for j in {1..20} ; do
    echo "[" > "$d/$j.json"
    for k in {1..100}; do
      echo '{"namespace": "default", "object": "'$i'", "subject_id": "'$j'", "relation": "'$k'"},' >> "$d/$j.json"
    done
    echo '{"namespace": "default", "object": "'$i'", "subject_id": "'$j'", "relation": "last"}]' >> "$d/$j.json"
  done
  ./keto relation-tuple create "$d"
done
