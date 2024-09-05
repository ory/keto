package api

import (
	"context"
	"encoding/json"
	"log"
	"maps"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"

	"connectrpc.com/vanguard/vanguardgrpc"
	"github.com/urfave/negroni"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"

	v1alpha2 "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type (
	Server struct {
		GRPCServer *grpc.Server
	}

	ServerOption func(o *serverOptions)

	serverOptions struct {
		grpcOptions []grpc.ServerOption
	}
)

func WithGRPCOption(grpcOption grpc.ServerOption) ServerOption {
	return func(o *serverOptions) {
		o.grpcOptions = append(o.grpcOptions, grpcOption)
	}
}

func NewServer(opt ...ServerOption) *Server {
	options := new(serverOptions)
	for _, o := range opt {
		o(options)
	}

	grpcServer := grpc.NewServer(options.grpcOptions...)
	reflection.Register(grpcServer)

	return &Server{GRPCServer: grpcServer}
}

func (s *Server) Handler() (http.Handler, error) {
	// Create a vanguard handler for all services registered in grpcServer
	handler, err := vanguardgrpc.NewTranscoder(s.GRPCServer)
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	mux.Handle("/", handler)
	mw := negroni.New()
	mw.Use(setErrorResponse)
	mw.Use(setRequestPath)
	mw.UseHandler(mux)

	return mw, nil
}

func (s *Server) RegisterService(desc *grpc.ServiceDesc, impl any) {
	s.GRPCServer.RegisterService(desc, impl)
}

var setErrorResponse = negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if ct := r.Header.Get("content-type"); strings.HasPrefix(ct, "application/grpc") || strings.HasPrefix(ct, "application/connect") {
		next(w, r)
		return
	}

	rr := httptest.NewRecorder()
	next(rr, r)

	maps.Copy(w.Header(), rr.Header())

	var spb statuspb.Status
	body := rr.Body.Bytes()
	protojson.Unmarshal(body, &spb)

	if spb.Code == int32(codes.Unknown) && strings.HasPrefix(spb.Message, "proto: ") {
		// Special case: error deserializing request into a protobuf is a bad request.
		w.WriteHeader(http.StatusBadRequest)
	} else if w.Header().Get("x-http-code") != "" {
		statusCode, err := strconv.Atoi(w.Header().Get("x-http-code"))
		if err != nil {
			log.Println("error:", err)
		}
		w.Header().Del("x-http-code")
		w.WriteHeader(statusCode)
	} else {
		w.WriteHeader(rr.Code)
	}

	if spb.Code == 0 && spb.Message == "" {
		_, _ = w.Write(body)
		return
	}

	s := status.FromProto(&spb)
	errResponse := v1alpha2.ErrorResponse{
		Error: &v1alpha2.ErrorResponse_Error{
			Code:    int64(rr.Code),
			Status:  http.StatusText(rr.Code),
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

	json.NewEncoder(w).Encode(&errResponse)
})

var setRequestPath = negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	r.Header.Set("x-path", r.URL.Path)
	next(w, r)
})

func SetStatusCode(ctx context.Context, statusCode int) {
	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", strconv.Itoa(statusCode)))
}

func SetLocationHeader(ctx context.Context, location string) {
	_ = grpc.SetHeader(ctx, metadata.Pairs("Location", location))
}

func RequestPath(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		path := md["x-path"]
		if len(path) > 0 {
			return path[0]
		}
	}
	return ""
}
