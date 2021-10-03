// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpcupstream

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

// UpstreamRegistryClient is the client API for UpstreamRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpstreamRegistryClient interface {
	FindUpstream(ctx context.Context, in *FindUpstreamRequest, opts ...grpc.CallOption) (*FindUpstreamReply, error)
	MapAuth(ctx context.Context, in *MapAuthRequest, opts ...grpc.CallOption) (*MapAuthReply, error)
}

type upstreamRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewUpstreamRegistryClient(cc grpc.ClientConnInterface) UpstreamRegistryClient {
	return &upstreamRegistryClient{cc}
}

func (c *upstreamRegistryClient) FindUpstream(ctx context.Context, in *FindUpstreamRequest, opts ...grpc.CallOption) (*FindUpstreamReply, error) {
	out := new(FindUpstreamReply)
	err := c.cc.Invoke(ctx, "/grpcupstream.UpstreamRegistry/FindUpstream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamRegistryClient) MapAuth(ctx context.Context, in *MapAuthRequest, opts ...grpc.CallOption) (*MapAuthReply, error) {
	out := new(MapAuthReply)
	err := c.cc.Invoke(ctx, "/grpcupstream.UpstreamRegistry/MapAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpstreamRegistryServer is the server API for UpstreamRegistry service.
// All implementations must embed UnimplementedUpstreamRegistryServer
// for forward compatibility
type UpstreamRegistryServer interface {
	FindUpstream(context.Context, *FindUpstreamRequest) (*FindUpstreamReply, error)
	MapAuth(context.Context, *MapAuthRequest) (*MapAuthReply, error)
	mustEmbedUnimplementedUpstreamRegistryServer()
}

// UnimplementedUpstreamRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedUpstreamRegistryServer struct {
}

func (UnimplementedUpstreamRegistryServer) FindUpstream(context.Context, *FindUpstreamRequest) (*FindUpstreamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUpstream not implemented")
}
func (UnimplementedUpstreamRegistryServer) MapAuth(context.Context, *MapAuthRequest) (*MapAuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapAuth not implemented")
}
func (UnimplementedUpstreamRegistryServer) mustEmbedUnimplementedUpstreamRegistryServer() {}

// UnsafeUpstreamRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UpstreamRegistryServer will
// result in compilation errors.
type UnsafeUpstreamRegistryServer interface {
	mustEmbedUnimplementedUpstreamRegistryServer()
}

func RegisterUpstreamRegistryServer(s grpc.ServiceRegistrar, srv UpstreamRegistryServer) {
	s.RegisterService(&UpstreamRegistry_ServiceDesc, srv)
}

func _UpstreamRegistry_FindUpstream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUpstreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamRegistryServer).FindUpstream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcupstream.UpstreamRegistry/FindUpstream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamRegistryServer).FindUpstream(ctx, req.(*FindUpstreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UpstreamRegistry_MapAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MapAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamRegistryServer).MapAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcupstream.UpstreamRegistry/MapAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamRegistryServer).MapAuth(ctx, req.(*MapAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UpstreamRegistry_ServiceDesc is the grpc.ServiceDesc for UpstreamRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UpstreamRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcupstream.UpstreamRegistry",
	HandlerType: (*UpstreamRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindUpstream",
			Handler:    _UpstreamRegistry_FindUpstream_Handler,
		},
		{
			MethodName: "MapAuth",
			Handler:    _UpstreamRegistry_MapAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "upstream.proto",
}