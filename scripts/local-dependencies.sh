#!/bin/bash

REPO_ROOT=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )/.." &> /dev/null && pwd )
export PATH="$REPO_ROOT/.bin/brew/bin:$REPO_ROOT/.bin/gobin:$PATH"
