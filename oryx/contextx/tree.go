// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package contextx

import "context"

type ContextKey int

const (
	ValidContextKey ContextKey = iota + 1
)

var RootContext = context.WithValue(context.Background(), ValidContextKey, true)

func IsRootContext(ctx context.Context) bool {
	is, ok := ctx.Value(ValidContextKey).(bool)
	return is && ok
}
