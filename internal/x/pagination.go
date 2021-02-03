package x

type (
	paginationOptions struct {
		Token string
		Size  int
	}
	PaginationOptionSetter func(*paginationOptions) *paginationOptions
)

const (
	PageTokenEnd = "no other page"
)

func WithToken(t string) PaginationOptionSetter {
	return func(opts *paginationOptions) *paginationOptions {
		opts.Token = t
		return opts
	}
}

func WithSize(size int) PaginationOptionSetter {
	return func(opts *paginationOptions) *paginationOptions {
		opts.Size = size
		return opts
	}
}

func GetPaginationOptions(modifiers ...PaginationOptionSetter) *paginationOptions {
	opts := &paginationOptions{}
	for _, f := range modifiers {
		opts = f(opts)
	}
	return opts
}
