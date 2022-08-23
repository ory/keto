#!/bin/bash
set -euo pipefail

export KETO_READ_REMOTE="127.0.0.1:4466"

keto expand access files /photos/beach.jpg --format json-pretty --max-depth 3 --insecure-disable-transport-security
