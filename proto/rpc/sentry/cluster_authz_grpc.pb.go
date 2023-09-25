// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/rpc/sentry/cluster_authz.proto

package sentry

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
	ClusterAuthorizationService_GetUserAuthorization_FullMethodName = "/paralus.dev.sentry.rpc.ClusterAuthorizationService/GetUserAuthorization"
)

// ClusterAuthorizationServiceClient is the client API for ClusterAuthorizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterAuthorizationServiceClient interface {
	GetUserAuthorization(ctx context.Context, in *GetUserAuthorizationRequest, opts ...grpc.CallOption) (*GetUserAuthorizationResponse, error)
}

type clusterAuthorizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterAuthorizationServiceClient(cc grpc.ClientConnInterface) ClusterAuthorizationServiceClient {
	return &clusterAuthorizationServiceClient{cc}
}

func (c *clusterAuthorizationServiceClient) GetUserAuthorization(ctx context.Context, in *GetUserAuthorizationRequest, opts ...grpc.CallOption) (*GetUserAuthorizationResponse, error) {
	out := new(GetUserAuthorizationResponse)
	err := c.cc.Invoke(ctx, ClusterAuthorizationService_GetUserAuthorization_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterAuthorizationServiceServer is the server API for ClusterAuthorizationService service.
// All implementations should embed UnimplementedClusterAuthorizationServiceServer
// for forward compatibility
type ClusterAuthorizationServiceServer interface {
	GetUserAuthorization(context.Context, *GetUserAuthorizationRequest) (*GetUserAuthorizationResponse, error)
}

// UnimplementedClusterAuthorizationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClusterAuthorizationServiceServer struct {
}

func (UnimplementedClusterAuthorizationServiceServer) GetUserAuthorization(context.Context, *GetUserAuthorizationRequest) (*GetUserAuthorizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAuthorization not implemented")
}

// UnsafeClusterAuthorizationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterAuthorizationServiceServer will
// result in compilation errors.
type UnsafeClusterAuthorizationServiceServer interface {
	mustEmbedUnimplementedClusterAuthorizationServiceServer()
}

func RegisterClusterAuthorizationServiceServer(s grpc.ServiceRegistrar, srv ClusterAuthorizationServiceServer) {
	s.RegisterService(&ClusterAuthorizationService_ServiceDesc, srv)
}

func _ClusterAuthorizationService_GetUserAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAuthorizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterAuthorizationServiceServer).GetUserAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterAuthorizationService_GetUserAuthorization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterAuthorizationServiceServer).GetUserAuthorization(ctx, req.(*GetUserAuthorizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterAuthorizationService_ServiceDesc is the grpc.ServiceDesc for ClusterAuthorizationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterAuthorizationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "paralus.dev.sentry.rpc.ClusterAuthorizationService",
	HandlerType: (*ClusterAuthorizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserAuthorization",
			Handler:    _ClusterAuthorizationService_GetUserAuthorization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rpc/sentry/cluster_authz.proto",
}
