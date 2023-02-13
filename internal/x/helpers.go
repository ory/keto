// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package x

import (
	"github.com/gofrs/uuid"
)

func UUIDs(n int) []uuid.UUID {
	res := make([]uuid.UUID, n)
	for i := range res {
		res[i] = uuid.Must(uuid.NewV4())
	}
	return res
}
