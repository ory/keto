// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package x

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"github.com/ory/herodot"
)

func UUIDs(n int) []uuid.UUID {
	res := make([]uuid.UUID, n)
	for i := range res {
		res[i] = uuid.Must(uuid.NewV4())
	}
	return res
}

var ValidationInterceptor grpc.UnaryServerInterceptor = func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if req, ok := req.(proto.Message); ok {
		if err := protovalidate.Validate(req); err != nil {
			return nil, herodot.ErrBadRequest.WithWrap(err).WithReason(err.Error())
		}
	}
	return handler(ctx, req)
}
