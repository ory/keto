#!/bin/bash

set -euo pipefail

if [ ! -d .bin/brew ] || [ ! -z "${BREW_FORCE_REINSTALL:-}" ]; then
  rm -rfd .bin/brew
  git clone --depth=1 https://github.com/Homebrew/brew .bin/brew
else
  echo "Using existing brew installation. To force a reinstall, set BREW_FORCE_REINSTALL=1"
fi
