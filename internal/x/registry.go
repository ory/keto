// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package x

import (
	"context"

	"connectrpc.com/connect"
	"github.com/gofrs/uuid"
)

type NetworkIDProvider interface {
	NetworkID(context.Context) uuid.UUID
}

type TransactorProvider interface {
	Transactor() interface {
		Transaction(ctx context.Context, f func(ctx context.Context) error) error
	}
}

type HandlerOptionsProvider interface {
	HandlerOptions() []connect.HandlerOption
}
