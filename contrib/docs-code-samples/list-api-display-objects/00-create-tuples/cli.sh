#!/bin/bash
set -euo pipefail

export KETO_WRITE_REMOTE="127.0.0.1:4467"

echo 'chats:memes#member@PM
chats:memes#member@Vincent
chats:memes#member@Julia

chats:cars#member@PM
chats:cars#member@Julia

chats:coffee-break#member@PM
chats:coffee-break#member@Vincent
chats:coffee-break#member@Julia
chats:coffee-break#member@Patrik' | \
  keto relation-tuple parse - --format json | \
    keto relation-tuple create - >/dev/null --insecure-disable-transport-security \
    && echo "Successfully created tuples" \
    || echo "Encountered error"
