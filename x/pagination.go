package x

type (
	paginationOptions struct {
		Page, PerPage int
	}
	PaginationOptionSetter func(*paginationOptions) *paginationOptions
)

func WithPage(page int) PaginationOptionSetter {
	return func(opts *paginationOptions) *paginationOptions {
		opts.Page = page
		return opts
	}
}

func WithPerPage(perPage int) PaginationOptionSetter {
	return func(opts *paginationOptions) *paginationOptions {
		opts.PerPage = perPage
		return opts
	}
}

func GetPaginationOptions(modifiers ...PaginationOptionSetter) *paginationOptions {
	opts := &paginationOptions{
		Page:    0,
		PerPage: 100,
	}
	for _, f := range modifiers {
		opts = f(opts)
	}
	return opts
}
