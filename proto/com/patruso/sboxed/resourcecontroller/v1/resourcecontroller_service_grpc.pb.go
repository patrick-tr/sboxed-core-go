// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/com/patruso/sboxed/resourcecontroller/v1/resourcecontroller_service.proto

package resourcecontroller

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
	ResourceControllerService_ListSandboxes_FullMethodName   = "/proto.com.patruso.sboxed.resourcecontroller.v1.ResourceControllerService/ListSandboxes"
	ResourceControllerService_ListRoutes_FullMethodName      = "/proto.com.patruso.sboxed.resourcecontroller.v1.ResourceControllerService/ListRoutes"
	ResourceControllerService_ListRouteGroups_FullMethodName = "/proto.com.patruso.sboxed.resourcecontroller.v1.ResourceControllerService/ListRouteGroups"
)

// ResourceControllerServiceClient is the client API for ResourceControllerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceControllerServiceClient interface {
	ListSandboxes(ctx context.Context, in *ListSandboxesRequest, opts ...grpc.CallOption) (*ListSandboxesResponse, error)
	ListRoutes(ctx context.Context, in *ListRoutesRequest, opts ...grpc.CallOption) (*ListRoutesResponse, error)
	ListRouteGroups(ctx context.Context, in *ListRouteGroupsRequest, opts ...grpc.CallOption) (*ListRouteGroupsResponse, error)
}

type resourceControllerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceControllerServiceClient(cc grpc.ClientConnInterface) ResourceControllerServiceClient {
	return &resourceControllerServiceClient{cc}
}

func (c *resourceControllerServiceClient) ListSandboxes(ctx context.Context, in *ListSandboxesRequest, opts ...grpc.CallOption) (*ListSandboxesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSandboxesResponse)
	err := c.cc.Invoke(ctx, ResourceControllerService_ListSandboxes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceControllerServiceClient) ListRoutes(ctx context.Context, in *ListRoutesRequest, opts ...grpc.CallOption) (*ListRoutesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRoutesResponse)
	err := c.cc.Invoke(ctx, ResourceControllerService_ListRoutes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceControllerServiceClient) ListRouteGroups(ctx context.Context, in *ListRouteGroupsRequest, opts ...grpc.CallOption) (*ListRouteGroupsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRouteGroupsResponse)
	err := c.cc.Invoke(ctx, ResourceControllerService_ListRouteGroups_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceControllerServiceServer is the server API for ResourceControllerService service.
// All implementations must embed UnimplementedResourceControllerServiceServer
// for forward compatibility.
type ResourceControllerServiceServer interface {
	ListSandboxes(context.Context, *ListSandboxesRequest) (*ListSandboxesResponse, error)
	ListRoutes(context.Context, *ListRoutesRequest) (*ListRoutesResponse, error)
	ListRouteGroups(context.Context, *ListRouteGroupsRequest) (*ListRouteGroupsResponse, error)
	mustEmbedUnimplementedResourceControllerServiceServer()
}

// UnimplementedResourceControllerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedResourceControllerServiceServer struct{}

func (UnimplementedResourceControllerServiceServer) ListSandboxes(context.Context, *ListSandboxesRequest) (*ListSandboxesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSandboxes not implemented")
}
func (UnimplementedResourceControllerServiceServer) ListRoutes(context.Context, *ListRoutesRequest) (*ListRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoutes not implemented")
}
func (UnimplementedResourceControllerServiceServer) ListRouteGroups(context.Context, *ListRouteGroupsRequest) (*ListRouteGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRouteGroups not implemented")
}
func (UnimplementedResourceControllerServiceServer) mustEmbedUnimplementedResourceControllerServiceServer() {
}
func (UnimplementedResourceControllerServiceServer) testEmbeddedByValue() {}

// UnsafeResourceControllerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceControllerServiceServer will
// result in compilation errors.
type UnsafeResourceControllerServiceServer interface {
	mustEmbedUnimplementedResourceControllerServiceServer()
}

func RegisterResourceControllerServiceServer(s grpc.ServiceRegistrar, srv ResourceControllerServiceServer) {
	// If the following call pancis, it indicates UnimplementedResourceControllerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ResourceControllerService_ServiceDesc, srv)
}

func _ResourceControllerService_ListSandboxes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSandboxesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceControllerServiceServer).ListSandboxes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceControllerService_ListSandboxes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceControllerServiceServer).ListSandboxes(ctx, req.(*ListSandboxesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceControllerService_ListRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceControllerServiceServer).ListRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceControllerService_ListRoutes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceControllerServiceServer).ListRoutes(ctx, req.(*ListRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceControllerService_ListRouteGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRouteGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceControllerServiceServer).ListRouteGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceControllerService_ListRouteGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceControllerServiceServer).ListRouteGroups(ctx, req.(*ListRouteGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceControllerService_ServiceDesc is the grpc.ServiceDesc for ResourceControllerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceControllerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.com.patruso.sboxed.resourcecontroller.v1.ResourceControllerService",
	HandlerType: (*ResourceControllerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSandboxes",
			Handler:    _ResourceControllerService_ListSandboxes_Handler,
		},
		{
			MethodName: "ListRoutes",
			Handler:    _ResourceControllerService_ListRoutes_Handler,
		},
		{
			MethodName: "ListRouteGroups",
			Handler:    _ResourceControllerService_ListRouteGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/com/patruso/sboxed/resourcecontroller/v1/resourcecontroller_service.proto",
}
