// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.0
// source: proto/geometry.proto

package gtpc

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

// GeometryServiceClient is the client API for GeometryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeometryServiceClient interface {
	Area(ctx context.Context, in *RectRequest, opts ...grpc.CallOption) (*AreaResponse, error)
	Perimeter(ctx context.Context, in *RectRequest, opts ...grpc.CallOption) (*PermiterResponse, error)
}

type geometryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGeometryServiceClient(cc grpc.ClientConnInterface) GeometryServiceClient {
	return &geometryServiceClient{cc}
}

func (c *geometryServiceClient) Area(ctx context.Context, in *RectRequest, opts ...grpc.CallOption) (*AreaResponse, error) {
	out := new(AreaResponse)
	err := c.cc.Invoke(ctx, "/geometry.GeometryService/Area", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geometryServiceClient) Perimeter(ctx context.Context, in *RectRequest, opts ...grpc.CallOption) (*PermiterResponse, error) {
	out := new(PermiterResponse)
	err := c.cc.Invoke(ctx, "/geometry.GeometryService/Perimeter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeometryServiceServer is the server API for GeometryService service.
// All implementations must embed UnimplementedGeometryServiceServer
// for forward compatibility
type GeometryServiceServer interface {
	Area(context.Context, *RectRequest) (*AreaResponse, error)
	Perimeter(context.Context, *RectRequest) (*PermiterResponse, error)
	mustEmbedUnimplementedGeometryServiceServer()
}

// UnimplementedGeometryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGeometryServiceServer struct {
}

func (UnimplementedGeometryServiceServer) Area(context.Context, *RectRequest) (*AreaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Area not implemented")
}
func (UnimplementedGeometryServiceServer) Perimeter(context.Context, *RectRequest) (*PermiterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Perimeter not implemented")
}
func (UnimplementedGeometryServiceServer) mustEmbedUnimplementedGeometryServiceServer() {}

// UnsafeGeometryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeometryServiceServer will
// result in compilation errors.
type UnsafeGeometryServiceServer interface {
	mustEmbedUnimplementedGeometryServiceServer()
}

func RegisterGeometryServiceServer(s grpc.ServiceRegistrar, srv GeometryServiceServer) {
	s.RegisterService(&GeometryService_ServiceDesc, srv)
}

func _GeometryService_Area_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeometryServiceServer).Area(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geometry.GeometryService/Area",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeometryServiceServer).Area(ctx, req.(*RectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeometryService_Perimeter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeometryServiceServer).Perimeter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geometry.GeometryService/Perimeter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeometryServiceServer).Perimeter(ctx, req.(*RectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GeometryService_ServiceDesc is the grpc.ServiceDesc for GeometryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GeometryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "geometry.GeometryService",
	HandlerType: (*GeometryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Area",
			Handler:    _GeometryService_Area_Handler,
		},
		{
			MethodName: "Perimeter",
			Handler:    _GeometryService_Perimeter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/geometry.proto",
}
