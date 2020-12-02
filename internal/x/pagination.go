package x

type (
	paginationOptions struct {
		Token string
		Size  uint
	}
	PaginationOptionSetter func(*paginationOptions) *paginationOptions
)

func WithToken(t string) PaginationOptionSetter {
	return func(opts *paginationOptions) *paginationOptions {
		opts.Token = t
		return opts
	}
}

func WithSize(size uint) PaginationOptionSetter {
	return func(opts *paginationOptions) *paginationOptions {
		opts.Size = size
		return opts
	}
}

func GetPaginationOptions(modifiers ...PaginationOptionSetter) *paginationOptions {
	opts := &paginationOptions{
		Token: "",
		Size:  100,
	}
	for _, f := range modifiers {
		opts = f(opts)
	}
	return opts
}
