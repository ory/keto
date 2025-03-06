// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: ory/keto/relation_tuples/v1alpha2/version.proto

package rts

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	VersionService_GetVersion_FullMethodName = "/ory.keto.relation_tuples.v1alpha2.VersionService/GetVersion"
)

// VersionServiceClient is the client API for VersionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The service returning the specific Ory Keto instance version.
//
// This service is part of the [read-APIs](../concepts/25_api-overview.mdx#read-apis) and [write-APIs](../concepts/25_api-overview.mdx#write-apis).
type VersionServiceClient interface {
	// Returns the version of the Ory Keto instance.
	//
	// This endpoint returns the service version typically notated using semantic versioning.
	//
	// If the service supports TLS Edge Termination, this endpoint does not require the
	// X-Forwarded-Proto header to be set.
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error)
}

type versionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVersionServiceClient(cc grpc.ClientConnInterface) VersionServiceClient {
	return &versionServiceClient{cc}
}

func (c *versionServiceClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVersionResponse)
	err := c.cc.Invoke(ctx, VersionService_GetVersion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VersionServiceServer is the server API for VersionService service.
// All implementations should embed UnimplementedVersionServiceServer
// for forward compatibility.
//
// The service returning the specific Ory Keto instance version.
//
// This service is part of the [read-APIs](../concepts/25_api-overview.mdx#read-apis) and [write-APIs](../concepts/25_api-overview.mdx#write-apis).
type VersionServiceServer interface {
	// Returns the version of the Ory Keto instance.
	//
	// This endpoint returns the service version typically notated using semantic versioning.
	//
	// If the service supports TLS Edge Termination, this endpoint does not require the
	// X-Forwarded-Proto header to be set.
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)
}

// UnimplementedVersionServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVersionServiceServer struct{}

func (UnimplementedVersionServiceServer) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedVersionServiceServer) testEmbeddedByValue() {}

// UnsafeVersionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VersionServiceServer will
// result in compilation errors.
type UnsafeVersionServiceServer interface {
	mustEmbedUnimplementedVersionServiceServer()
}

func RegisterVersionServiceServer(s grpc.ServiceRegistrar, srv VersionServiceServer) {
	// If the following call pancis, it indicates UnimplementedVersionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VersionService_ServiceDesc, srv)
}

func _VersionService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VersionService_GetVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServiceServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VersionService_ServiceDesc is the grpc.ServiceDesc for VersionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VersionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ory.keto.relation_tuples.v1alpha2.VersionService",
	HandlerType: (*VersionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _VersionService_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ory/keto/relation_tuples/v1alpha2/version.proto",
}
