// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package schema

import (
	"context"
	"io"
	"net/http"

	"connectrpc.com/connect"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/ory/herodot"
	"github.com/ory/x/httprouterx"
	"github.com/ory/x/httpx"
	"github.com/ory/x/logrusx"

	opl "github.com/ory/keto/gen/go/ory/keto/opl/v1alpha1"
	"github.com/ory/keto/gen/go/ory/keto/opl/v1alpha1/oplconnect"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoapi"
)

type (
	handlerDependencies interface {
		logrusx.Provider
		httpx.WriterProvider
		x.HandlerOptionsProvider
	}
	Handler struct {
		oplconnect.UnimplementedSyntaxServiceHandler
		d handlerDependencies
	}
)

const RouteBase = "/opl/syntax/check"

func NewHandler(d handlerDependencies) *Handler {
	return &Handler{d: d}
}

func (h *Handler) RegisterSyntaxRoutes(r httprouterx.Router) {
	r.POST(RouteBase, h.postCheckOplSyntax)
	r.Handle(oplconnect.NewSyntaxServiceHandler(h, h.d.HandlerOptions()...))
}

func (h *Handler) ProtoFiles() []protoreflect.FileDescriptor {
	return []protoreflect.FileDescriptor{opl.File_ory_keto_opl_v1alpha1_syntax_service_proto}
}

func (h *Handler) Check(_ context.Context, request *connect.Request[opl.CheckRequest]) (*connect.Response[opl.CheckResponse], error) {
	_, parseErrors := Parse(string(request.Msg.GetContent()))
	apiErrors := make([]*opl.ParseError, len(parseErrors))
	for i, e := range parseErrors {
		apiErrors[i] = e.ToProto()
	}
	return connect.NewResponse(&opl.CheckResponse{ParseErrors: apiErrors}), nil
}

// Check OPL Syntax Request Parameters
//
// swagger:parameters checkOplSyntax
//
//lint:ignore U1000 Used to generate Swagger and OpenAPI definitions
type checkOplSyntax struct {
	// in: body
	Body checkOplSyntaxBody
}

// Ory Permission Language Document
//
// swagger:model checkOplSyntaxBody
//
//lint:ignore U1000 Used to generate Swagger and OpenAPI definitions
type checkOplSyntaxBody string

// swagger:route POST /opl/syntax/check relationship checkOplSyntax
//
// # Check the syntax of an OPL file
//
// The OPL file is expected in the body of the request.
//
//	Consumes:
//	- text/plain
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Responses:
//	  200: checkOplSyntaxResult
//	  400: errorGeneric
//	  default: errorGeneric
//
//	Extensions:
//	  x-ory-ratelimit-bucket: keto-admin-medium
func (h *Handler) postCheckOplSyntax(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrBadRequest().WithError(err.Error())))
		return
	}
	_, parseErrors := Parse(string(data))
	apiErrors := make([]*ketoapi.ParseError, len(parseErrors))
	for i, e := range parseErrors {
		apiErrors[i] = e.ToAPI()
	}
	h.d.Writer().Write(w, r, &ketoapi.CheckOPLSyntaxResponse{Errors: apiErrors})
}
