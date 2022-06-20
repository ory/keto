package x

import "github.com/gofrs/uuid"

func Ptr[T any](v T) *T {
	return &v
}

func UUIDs(n int) []uuid.UUID {
	res := make([]uuid.UUID, n)
	for i := range res {
		res[i] = uuid.Must(uuid.NewV4())
	}
	return res
}
