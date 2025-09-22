package api

import (
	"context"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"connectrpc.com/vanguard"
	"connectrpc.com/vanguard/vanguardgrpc"
	"github.com/urfave/negroni"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func init() {
	encoding.RegisterCodec(vanguardgrpc.NewCodec(&vanguard.JSONCodec{
		MarshalOptions:   protojson.MarshalOptions{EmitUnpopulated: true},
		UnmarshalOptions: protojson.UnmarshalOptions{DiscardUnknown: true},
	}))
}

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

var (
	invalidFieldPathRe = regexp.MustCompile(`in field path ".*":`)
	unexpectedBodyRe   = regexp.MustCompile(`request should have no body; instead got \d bytes`)
)

var setErrorResponse = negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if ct := r.Header.Get("content-type"); strings.HasPrefix(ct, "application/grpc") || strings.HasPrefix(ct, "application/connect") {
		next(w, r)
		return
	}

	resp := &response{w: w}
	next(resp, r)
	resp.flush()
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
