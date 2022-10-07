// Copyright © 2022 Ory Corp

package schema

func Fuzz(data []byte) int {
	Parse(string(data))
	return 0
}
