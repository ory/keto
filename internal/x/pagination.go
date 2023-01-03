// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package x

type (
	PaginationOptions struct {
		Token string `json:"page_token"`
		Size  int    `json:"page_size"`
	}
	PaginationOptionSetter func(*PaginationOptions) *PaginationOptions
)

func WithToken(t string) PaginationOptionSetter {
	return func(opts *PaginationOptions) *PaginationOptions {
		opts.Token = t
		return opts
	}
}

func WithSize(size int) PaginationOptionSetter {
	return func(opts *PaginationOptions) *PaginationOptions {
		opts.Size = size
		return opts
	}
}

func GetPaginationOptions(modifiers ...PaginationOptionSetter) *PaginationOptions {
	opts := &PaginationOptions{}
	for _, f := range modifiers {
		opts = f(opts)
	}
	return opts
}
