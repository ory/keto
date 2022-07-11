#!/usr/bin/env sh
set -euo pipefail

echo "$PATH"

keto status --block
echo "Keto is ready"

keto relation-tuple get --relation last --subject-id 20 --read-remote migration-v09-keto-v0.8-1:4466

d="$(mktemp -d)"

for i in $(seq 1 500) ; do
  rm -rf "$d/*"
  for j in $(seq 1 20) ; do
    echo "[" > "$d/$j.json"
    for k in $(seq 1 100); do
      echo '{"namespace": "default", "object": "'$i'", "subject_id": "'$j'", "relation": "'$k'"},' >> "$d/$j.json"
    done
    echo '{"namespace": "default", "object": "'$i'", "subject_id": "'$j'", "relation": "last"}]' >> "$d/$j.json"
  done
  echo "Creating page $i"
  keto relation-tuple create "$d"
done
