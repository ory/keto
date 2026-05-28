#!/bin/bash
set -euo pipefail

export KETO_READ_REMOTE="127.0.0.1:4466"

keto check User:alice viewer File:data.txt --insecure-disable-transport-security
