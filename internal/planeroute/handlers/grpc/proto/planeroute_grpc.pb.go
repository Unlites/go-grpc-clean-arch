// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.0
// source: internal/planeroute/handlers/grpc/proto/planeroute.proto

package proto

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

// PlaneRouteClient is the client API for PlaneRoute service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlaneRouteClient interface {
	AddPlaneRoute(ctx context.Context, in *PlaneRouteRequest, opts ...grpc.CallOption) (*ResultResponse, error)
	StreamCurrentRoutes(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (PlaneRoute_StreamCurrentRoutesClient, error)
}

type planeRouteClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaneRouteClient(cc grpc.ClientConnInterface) PlaneRouteClient {
	return &planeRouteClient{cc}
}

func (c *planeRouteClient) AddPlaneRoute(ctx context.Context, in *PlaneRouteRequest, opts ...grpc.CallOption) (*ResultResponse, error) {
	out := new(ResultResponse)
	err := c.cc.Invoke(ctx, "/planeroute.PlaneRoute/AddPlaneRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planeRouteClient) StreamCurrentRoutes(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (PlaneRoute_StreamCurrentRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &PlaneRoute_ServiceDesc.Streams[0], "/planeroute.PlaneRoute/StreamCurrentRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &planeRouteStreamCurrentRoutesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PlaneRoute_StreamCurrentRoutesClient interface {
	Recv() (*PlaneRouteResponse, error)
	grpc.ClientStream
}

type planeRouteStreamCurrentRoutesClient struct {
	grpc.ClientStream
}

func (x *planeRouteStreamCurrentRoutesClient) Recv() (*PlaneRouteResponse, error) {
	m := new(PlaneRouteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PlaneRouteServer is the server API for PlaneRoute service.
// All implementations must embed UnimplementedPlaneRouteServer
// for forward compatibility
type PlaneRouteServer interface {
	AddPlaneRoute(context.Context, *PlaneRouteRequest) (*ResultResponse, error)
	StreamCurrentRoutes(*EmptyRequest, PlaneRoute_StreamCurrentRoutesServer) error
	mustEmbedUnimplementedPlaneRouteServer()
}

// UnimplementedPlaneRouteServer must be embedded to have forward compatible implementations.
type UnimplementedPlaneRouteServer struct {
}

func (UnimplementedPlaneRouteServer) AddPlaneRoute(context.Context, *PlaneRouteRequest) (*ResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPlaneRoute not implemented")
}
func (UnimplementedPlaneRouteServer) StreamCurrentRoutes(*EmptyRequest, PlaneRoute_StreamCurrentRoutesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamCurrentRoutes not implemented")
}
func (UnimplementedPlaneRouteServer) mustEmbedUnimplementedPlaneRouteServer() {}

// UnsafePlaneRouteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlaneRouteServer will
// result in compilation errors.
type UnsafePlaneRouteServer interface {
	mustEmbedUnimplementedPlaneRouteServer()
}

func RegisterPlaneRouteServer(s grpc.ServiceRegistrar, srv PlaneRouteServer) {
	s.RegisterService(&PlaneRoute_ServiceDesc, srv)
}

func _PlaneRoute_AddPlaneRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaneRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaneRouteServer).AddPlaneRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/planeroute.PlaneRoute/AddPlaneRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaneRouteServer).AddPlaneRoute(ctx, req.(*PlaneRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaneRoute_StreamCurrentRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PlaneRouteServer).StreamCurrentRoutes(m, &planeRouteStreamCurrentRoutesServer{stream})
}

type PlaneRoute_StreamCurrentRoutesServer interface {
	Send(*PlaneRouteResponse) error
	grpc.ServerStream
}

type planeRouteStreamCurrentRoutesServer struct {
	grpc.ServerStream
}

func (x *planeRouteStreamCurrentRoutesServer) Send(m *PlaneRouteResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PlaneRoute_ServiceDesc is the grpc.ServiceDesc for PlaneRoute service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlaneRoute_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "planeroute.PlaneRoute",
	HandlerType: (*PlaneRouteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPlaneRoute",
			Handler:    _PlaneRoute_AddPlaneRoute_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamCurrentRoutes",
			Handler:       _PlaneRoute_StreamCurrentRoutes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/planeroute/handlers/grpc/proto/planeroute.proto",
}