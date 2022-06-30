#!/bin/bash
set -euo pipefail

PATH=".bin/brew/bin:$PATH"

help() {
  echo "Usage: brew-update-version.sh <brew-formula> <version>"
  echo "Example: brew-update-version.sh ory/tap/cli 1.0.0"
}

if [ $# -ne 2 ]; then
  help
  exit 1
fi

# copy the formula with the specific version to the local tap
brew extract --version="$2" "$1" keto/tools
# sync the formulas to the directory included in git
rsync -a --delete .bin/brew/Library/Taps/keto/homebrew-tools/Formula/ .bin/formula-pins/
