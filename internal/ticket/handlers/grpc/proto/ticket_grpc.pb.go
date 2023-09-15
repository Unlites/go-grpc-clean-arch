// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.0
// source: internal/ticket/handlers/grpc/proto/ticket.proto

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

// TicketClient is the client API for Ticket service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketClient interface {
	AddTicket(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*ResultResponse, error)
}

type ticketClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketClient(cc grpc.ClientConnInterface) TicketClient {
	return &ticketClient{cc}
}

func (c *ticketClient) AddTicket(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*ResultResponse, error) {
	out := new(ResultResponse)
	err := c.cc.Invoke(ctx, "/ticket.Ticket/AddTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketServer is the server API for Ticket service.
// All implementations must embed UnimplementedTicketServer
// for forward compatibility
type TicketServer interface {
	AddTicket(context.Context, *TicketRequest) (*ResultResponse, error)
	mustEmbedUnimplementedTicketServer()
}

// UnimplementedTicketServer must be embedded to have forward compatible implementations.
type UnimplementedTicketServer struct {
}

func (UnimplementedTicketServer) AddTicket(context.Context, *TicketRequest) (*ResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTicket not implemented")
}
func (UnimplementedTicketServer) mustEmbedUnimplementedTicketServer() {}

// UnsafeTicketServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketServer will
// result in compilation errors.
type UnsafeTicketServer interface {
	mustEmbedUnimplementedTicketServer()
}

func RegisterTicketServer(s grpc.ServiceRegistrar, srv TicketServer) {
	s.RegisterService(&Ticket_ServiceDesc, srv)
}

func _Ticket_AddTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServer).AddTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket.Ticket/AddTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServer).AddTicket(ctx, req.(*TicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Ticket_ServiceDesc is the grpc.ServiceDesc for Ticket service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ticket_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ticket.Ticket",
	HandlerType: (*TicketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTicket",
			Handler:    _Ticket_AddTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/ticket/handlers/grpc/proto/ticket.proto",
}
