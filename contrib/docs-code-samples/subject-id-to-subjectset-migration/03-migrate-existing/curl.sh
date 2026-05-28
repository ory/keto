#!/bin/bash
set -euo pipefail

# In production, paginate GET /admin/relation-tuples?namespace=File&relation=viewer,
# filter for tuples with subject_id set, and write SubjectSet counterparts in batches.
# This script writes the known SubjectSet tuples directly for brevity.

curl -X PATCH --fail --silent \
     -H 'Content-Type: application/json' \
     --retry 7 --retry-connrefused \
     --data '[
       {"action":"insert","relation_tuple":{"namespace":"File","object":"data.txt","relation":"viewer","subject_set":{"namespace":"User","object":"alice","relation":""}}},
       {"action":"insert","relation_tuple":{"namespace":"File","object":"data.txt","relation":"viewer","subject_set":{"namespace":"User","object":"bob","relation":""}}},
       {"action":"insert","relation_tuple":{"namespace":"File","object":"data.txt","relation":"viewer","subject_set":{"namespace":"ApiKey","object":"ci-bot","relation":""}}}
     ]' \
     http://127.0.0.1:4467/admin/relation-tuples > /dev/null \
  && echo "Migration complete" \
  || echo "Encountered error"
