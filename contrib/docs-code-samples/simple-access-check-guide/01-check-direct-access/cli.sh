#!/bin/bash
set -euo pipefail

export KETO_READ_REMOTE="127.0.0.1:4466"

keto check john decypher messages 02y_15_4w350m3 --insecure-disable-transport-security
