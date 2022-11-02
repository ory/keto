// Copyright © 2022 Ory Corp

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
