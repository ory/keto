// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package x

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gofrs/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

func UUIDs(n int) []uuid.UUID {
	res := make([]uuid.UUID, n)
	for i := range res {
		res[i] = uuid.Must(uuid.NewV4())
	}
	return res
}

var GRPCGatewayMuxOptions = []runtime.ServeMuxOption{
	runtime.WithForwardResponseOption(HttpResponseModifier),
	runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
		md := make(metadata.MD)

		contentLength, _ := strconv.Atoi(req.Header.Get("Content-Length"))
		md.Set("hasbody", strconv.FormatBool(contentLength > 0))
		md.Set("path", req.URL.Path)
		return md
	}),
	runtime.WithErrorHandler(func(_ context.Context, _ *runtime.ServeMux, _ runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
		var customStatus *runtime.HTTPStatusError
		if errors.As(err, &customStatus) {
			err = customStatus.Err
		}

		s := status.Convert(err)

		st := runtime.HTTPStatusFromCode(s.Code())
		if customStatus != nil {
			st = customStatus.HTTPStatus
		}
		w.WriteHeader(st)

		errResponse := rts.ErrorResponse{
			Error: &rts.ErrorResponse_Error{
				Code:    int64(st),
				Status:  http.StatusText(st),
				Message: s.Message(),
			},
		}
		for _, detail := range s.Details() {
			switch t := detail.(type) {
			case *errdetails.ErrorInfo:
				errResponse.Error.Reason = t.Reason
			case *errdetails.DebugInfo:
				errResponse.Error.Debug = t.Detail
			case *errdetails.RequestInfo:
				errResponse.Error.Request = t.RequestId
			case *errdetails.BadRequest:
				errResponse.Error.Details = make(map[string]string, len(t.FieldViolations))
				for _, v := range t.FieldViolations {
					errResponse.Error.Details[v.Field] = v.Description
				}
			}
		}

		buf, _ := json.Marshal(&errResponse)
		_, _ = w.Write(buf)
	}),
}
