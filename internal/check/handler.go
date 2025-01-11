// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check

import (
	"context"
	"net/http"

	"github.com/ory/herodot"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ory/keto/internal/check/checkgroup"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/internal/x/api"
	"github.com/ory/keto/ketoapi"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type (
	handlerDependencies interface {
		EngineProvider
		relationtuple.ManagerProvider
		relationtuple.MapperProvider
		x.LoggerProvider
		x.WriterProvider
		config.Provider
	}
	Handler struct {
		d handlerDependencies
	}
)

var (
	_ rts.CheckServiceServer = (*Handler)(nil)
)

func NewHandler(d handlerDependencies) *Handler {
	return &Handler{d: d}
}

const (
	RouteBase        = "/relation-tuples/check"
	OpenAPIRouteBase = RouteBase + "/openapi"
	BatchRoute       = "/relation-tuples/batch/check"
)

func (h *Handler) RegisterReadGRPC(s *grpc.Server) {
	rts.RegisterCheckServiceServer(s, h)
}

func (h *Handler) Check(ctx context.Context, req *rts.CheckRequest) (res *rts.CheckResponse, err error) {
	tuple := (&ketoapi.RelationTuple{}).FromCheckRequest(req)

	// Check if we should set the HTTP status code to 403 instead of 200 if the check fails.

	if api.RequestPath(ctx) == RouteBase {
		defer func() {
			if res != nil && !res.Allowed {
				api.SetStatusCode(ctx, http.StatusForbidden)
			}
		}()
	}

	if tuple.SubjectID == nil && tuple.SubjectSet == nil {
		return nil, ketoapi.ErrNilSubject
	}

	internalTuple, err := h.d.ReadOnlyMapper().FromTuple(ctx, tuple)
	if errors.Is(err, herodot.ErrNotFound) {
		res = &rts.CheckResponse{Allowed: false}
		return res, nil
	}
	if err != nil {
		return nil, err
	}
	allowed, err := h.d.PermissionEngine().CheckIsMember(ctx, internalTuple[0], int(req.MaxDepth))
	if err != nil {
		return nil, err
	}

	res = &rts.CheckResponse{Allowed: allowed}
	return res, nil
}

// BatchCheck is the gRPC entry point for checking batches of tuples
func (h *Handler) BatchCheck(ctx context.Context, req *rts.BatchCheckRequest) (*rts.BatchCheckResponse, error) {
	// Handle the case where the tuples are passed in the request body.
	if req.RestBody != nil {
		if req.Tuples != nil {
			return nil, status.Error(codes.InvalidArgument, "either tuples or rest_body should be set, not both")
		}
		req.Tuples = req.RestBody.Tuples
	}

	if len(req.Tuples) > h.d.Config(ctx).BatchCheckMaxBatchSize() {
		return nil, status.Errorf(codes.InvalidArgument,
			"batch exceeds max size of %v", h.d.Config(ctx).BatchCheckMaxBatchSize())
	}

	ketoTuples := make([]*ketoapi.RelationTuple, len(req.Tuples))
	for i, tuple := range req.Tuples {
		ketoTuples[i] = (&ketoapi.RelationTuple{}).FromProto(tuple)
	}

	results, err := h.d.PermissionEngine().BatchCheck(ctx, ketoTuples, int(req.MaxDepth))
	if err != nil {
		return nil, err
	}

	responses := make([]*rts.CheckResponseWithError, len(results))
	for i, result := range results {
		errMsg := ""
		if result.Err != nil {
			errMsg = result.Err.Error()
		}
		responses[i] = &rts.CheckResponseWithError{
			Allowed: result.Membership == checkgroup.IsMember,
			Error:   errMsg,
		}
	}

	return &rts.BatchCheckResponse{
		Results: responses,
	}, nil
}
