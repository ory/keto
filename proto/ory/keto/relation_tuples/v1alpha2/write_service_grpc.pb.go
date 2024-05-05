// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: ory/keto/relation_tuples/v1alpha2/write_service.proto

package rts

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	WriteService_TransactRelationTuples_FullMethodName = "/ory.keto.relation_tuples.v1alpha2.WriteService/TransactRelationTuples"
	WriteService_DeleteRelationTuples_FullMethodName   = "/ory.keto.relation_tuples.v1alpha2.WriteService/DeleteRelationTuples"
)

// WriteServiceClient is the client API for WriteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WriteServiceClient interface {
	// Writes one or more relationships in a single transaction.
	TransactRelationTuples(ctx context.Context, in *TransactRelationTuplesRequest, opts ...grpc.CallOption) (*TransactRelationTuplesResponse, error)
	// Deletes relationships based on relation query
	DeleteRelationTuples(ctx context.Context, in *DeleteRelationTuplesRequest, opts ...grpc.CallOption) (*DeleteRelationTuplesResponse, error)
}

type writeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWriteServiceClient(cc grpc.ClientConnInterface) WriteServiceClient {
	return &writeServiceClient{cc}
}

func (c *writeServiceClient) TransactRelationTuples(ctx context.Context, in *TransactRelationTuplesRequest, opts ...grpc.CallOption) (*TransactRelationTuplesResponse, error) {
	out := new(TransactRelationTuplesResponse)
	err := c.cc.Invoke(ctx, WriteService_TransactRelationTuples_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writeServiceClient) DeleteRelationTuples(ctx context.Context, in *DeleteRelationTuplesRequest, opts ...grpc.CallOption) (*DeleteRelationTuplesResponse, error) {
	out := new(DeleteRelationTuplesResponse)
	err := c.cc.Invoke(ctx, WriteService_DeleteRelationTuples_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WriteServiceServer is the server API for WriteService service.
// All implementations should embed UnimplementedWriteServiceServer
// for forward compatibility
type WriteServiceServer interface {
	// Writes one or more relationships in a single transaction.
	TransactRelationTuples(context.Context, *TransactRelationTuplesRequest) (*TransactRelationTuplesResponse, error)
	// Deletes relationships based on relation query
	DeleteRelationTuples(context.Context, *DeleteRelationTuplesRequest) (*DeleteRelationTuplesResponse, error)
}

// UnimplementedWriteServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWriteServiceServer struct {
}

func (UnimplementedWriteServiceServer) TransactRelationTuples(context.Context, *TransactRelationTuplesRequest) (*TransactRelationTuplesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransactRelationTuples not implemented")
}
func (UnimplementedWriteServiceServer) DeleteRelationTuples(context.Context, *DeleteRelationTuplesRequest) (*DeleteRelationTuplesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRelationTuples not implemented")
}

// UnsafeWriteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WriteServiceServer will
// result in compilation errors.
type UnsafeWriteServiceServer interface {
	mustEmbedUnimplementedWriteServiceServer()
}

func RegisterWriteServiceServer(s grpc.ServiceRegistrar, srv WriteServiceServer) {
	s.RegisterService(&WriteService_ServiceDesc, srv)
}

func _WriteService_TransactRelationTuples_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactRelationTuplesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriteServiceServer).TransactRelationTuples(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WriteService_TransactRelationTuples_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriteServiceServer).TransactRelationTuples(ctx, req.(*TransactRelationTuplesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WriteService_DeleteRelationTuples_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRelationTuplesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriteServiceServer).DeleteRelationTuples(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WriteService_DeleteRelationTuples_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriteServiceServer).DeleteRelationTuples(ctx, req.(*DeleteRelationTuplesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WriteService_ServiceDesc is the grpc.ServiceDesc for WriteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WriteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ory.keto.relation_tuples.v1alpha2.WriteService",
	HandlerType: (*WriteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransactRelationTuples",
			Handler:    _WriteService_TransactRelationTuples_Handler,
		},
		{
			MethodName: "DeleteRelationTuples",
			Handler:    _WriteService_DeleteRelationTuples_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ory/keto/relation_tuples/v1alpha2/write_service.proto",
}
