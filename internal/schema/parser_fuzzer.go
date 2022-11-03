// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package schema

func Fuzz(data []byte) int {
	Parse(string(data))
	return 0
}
