#!/bin/bash
set -euo pipefail

export KETO_WRITE_REMOTE="127.0.0.1:4467"

echo 'Chat:memes#member@User:PM
Chat:memes#member@User:Vincent
Chat:memes#member@User:Julia

Chat:cars#member@User:PM
Chat:cars#member@User:Julia

Chat:coffee-break#member@User:PM
Chat:coffee-break#member@User:Vincent
Chat:coffee-break#member@User:Julia
Chat:coffee-break#member@User:Patrik' | \
  keto relation-tuple parse - --format json | \
    keto relation-tuple create - >/dev/null --insecure-disable-transport-security \
    && echo "Successfully created tuples" \
    || echo "Encountered error"
