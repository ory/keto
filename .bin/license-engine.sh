#!/bin/bash

# This script detects non-compliant licenses in the output of language-specific license checkers.

# These licenses are allowed.
# These are the exact and complete license strings for 100% legal certainty, no regexes.
ALLOWED_LICENSES=(
	'0BSD'
	'AFLv2.1'
	'(AFL-2.1 OR BSD-3-Clause)'
	'Apache-2.0'
	'Apache*'
	'BSD*'
	'BSD-2-Clause'
	'(BSD-2-Clause OR MIT OR Apache-2.0)'
	'BSD-3-Clause'
	'(BSD-3-Clause OR GPL-2.0)'
	'CC0-1.0'
	'CC-BY-3.0'
	'CC-BY-4.0'
	'(CC-BY-4.0 AND MIT)'
	'ISC'
	'LGPL-2.1' # LGPL allows commercial use, requires only that modifications to LGPL-protected libraries are published under a GPL-compatible license
	'MIT'
	'MIT*'
	'(MIT AND BSD-3-Clause)'
	'(MIT AND Zlib)'
	'(MIT OR Apache-2.0)'
	'(MIT OR CC0-1.0)'
	'MPL-2.0'
	'(MPL-2.0 OR Apache-2.0)'
	'Python-2.0' # the Python-2.0 is a permissive license, see https://en.wikipedia.org/wiki/Python_License
	'Unlicense'
	'(WTFPL OR MIT)'
)

# These modules don't work with the current license checkers
# and have been manually verified to have a compatible license (regex format).
APPROVED_MODULES=(
	'https://github.com/ory-corp/cloud/'            # Ory IP
	'github.com/ory/kratos-client-go'               # Apache-2.0
	'github.com/ory/hydra-client-go'                # Apache-2.0
	'github.com/gobuffalo/github_flavored_markdown' # MIT
	'github.com/ory/keto/.*'
	'buffers@0.1.1'                                     # MIT: original source at http://github.com/substack/node-bufferlist is deleted but a fork at https://github.com/pkrumins/node-bufferlist/blob/master/LICENSE contains the original license by the original author (James Halliday)
	'https://github.com/iconify/iconify/packages/react' # MIT: license is in root of monorepo at https://github.com/iconify/iconify/blob/main/license.txt
	'github.com/gobuffalo/.*'                           # MIT: license is in root of monorepo at https://github.com/gobuffalo/github_flavored_markdown/blob/main/LICENSE
	'github.com/ory-corp/cloud/.*'                      # Ory IP
)

# These lines in the output should be ignored (plain text, no regex).
IGNORE_LINES=(
	'"module name","licenses"' # header of license output for Node.js
)

echo_green() {
	printf "\e[1;92m%s\e[0m\n" "$@"
}

echo_red() {
	printf "\e[0;91m%s\e[0m\n" "$@"
}

# capture STDIN
input=$(cat -)

# remove ignored lines
for ignored in "${IGNORE_LINES[@]}"; do
	input=$(echo "$input" | grep -vF "$ignored")
done

# remove pre-approved modules
for approved in "${APPROVED_MODULES[@]}"; do
	input=$(echo "$input" | grep -v "\"${approved}\"")
	input=$(echo "$input" | grep -v "\"Custom: ${approved}\"")
done

# remove allowed licenses
for allowed in "${ALLOWED_LICENSES[@]}"; do
	input=$(echo "$input" | grep -vF "\"${allowed}\"")
done

# anything left in the input at this point is a module with an invalid license

# print outcome
if [ -z "$input" ]; then
	echo_green "Licenses are okay."
else
	echo_red "Unknown licenses found!"
	echo
	echo "$input"
	exit 1
fi
