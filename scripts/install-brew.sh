#!/bin/bash

set -euo pipefail

if [ ! -d .bin/brew ] || [ ! -z "${FORCE_BREW_REINSTALL:-}" ]; then
  rm -rfd .bin/brew
  git clone --depth=1 https://github.com/Homebrew/brew .bin/brew

  # create local tap for pinned versions
  .bin/brew/bin/brew tap-new --no-git keto/tools
  # populate local tap with pinned versions
  rsync -a --delete .bin/formula-pins/ .bin/brew/Library/Taps/keto/homebrew-tools/Formula/
else
  echo "Using existing brew installation. To force a reinstall, set the env var FORCE_BREW_REINSTALL"
fi
