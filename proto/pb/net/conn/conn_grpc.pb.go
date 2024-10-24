// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: net/conn/conn.proto

package conn

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Conn_Read_FullMethodName             = "/net.conn.Conn/Read"
	Conn_Write_FullMethodName            = "/net.conn.Conn/Write"
	Conn_Close_FullMethodName            = "/net.conn.Conn/Close"
	Conn_SetDeadline_FullMethodName      = "/net.conn.Conn/SetDeadline"
	Conn_SetReadDeadline_FullMethodName  = "/net.conn.Conn/SetReadDeadline"
	Conn_SetWriteDeadline_FullMethodName = "/net.conn.Conn/SetWriteDeadline"
)

// ConnClient is the client API for Conn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnClient interface {
	// Read reads data from the connection.
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	// Write writes data to the connection.
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	// Close closes the connection.
	Close(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// SetDeadline sets the read and write deadlines associated
	// with the connection.
	SetDeadline(ctx context.Context, in *SetDeadlineRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// SetReadDeadline sets the deadline for future Read calls
	// and any currently-blocked Read call.
	SetReadDeadline(ctx context.Context, in *SetDeadlineRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// SetWriteDeadline sets the deadline for future Write calls
	// and any currently-blocked Write call.
	SetWriteDeadline(ctx context.Context, in *SetDeadlineRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type connClient struct {
	cc grpc.ClientConnInterface
}

func NewConnClient(cc grpc.ClientConnInterface) ConnClient {
	return &connClient{cc}
}

func (c *connClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, Conn_Read_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, Conn_Write_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connClient) Close(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Conn_Close_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connClient) SetDeadline(ctx context.Context, in *SetDeadlineRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Conn_SetDeadline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connClient) SetReadDeadline(ctx context.Context, in *SetDeadlineRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Conn_SetReadDeadline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connClient) SetWriteDeadline(ctx context.Context, in *SetDeadlineRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Conn_SetWriteDeadline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnServer is the server API for Conn service.
// All implementations must embed UnimplementedConnServer
// for forward compatibility
type ConnServer interface {
	// Read reads data from the connection.
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	// Write writes data to the connection.
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
	// Close closes the connection.
	Close(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// SetDeadline sets the read and write deadlines associated
	// with the connection.
	SetDeadline(context.Context, *SetDeadlineRequest) (*emptypb.Empty, error)
	// SetReadDeadline sets the deadline for future Read calls
	// and any currently-blocked Read call.
	SetReadDeadline(context.Context, *SetDeadlineRequest) (*emptypb.Empty, error)
	// SetWriteDeadline sets the deadline for future Write calls
	// and any currently-blocked Write call.
	SetWriteDeadline(context.Context, *SetDeadlineRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedConnServer()
}

// UnimplementedConnServer must be embedded to have forward compatible implementations.
type UnimplementedConnServer struct {
}

func (UnimplementedConnServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedConnServer) Write(context.Context, *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedConnServer) Close(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}
func (UnimplementedConnServer) SetDeadline(context.Context, *SetDeadlineRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDeadline not implemented")
}
func (UnimplementedConnServer) SetReadDeadline(context.Context, *SetDeadlineRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetReadDeadline not implemented")
}
func (UnimplementedConnServer) SetWriteDeadline(context.Context, *SetDeadlineRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetWriteDeadline not implemented")
}
func (UnimplementedConnServer) mustEmbedUnimplementedConnServer() {}

// UnsafeConnServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnServer will
// result in compilation errors.
type UnsafeConnServer interface {
	mustEmbedUnimplementedConnServer()
}

func RegisterConnServer(s grpc.ServiceRegistrar, srv ConnServer) {
	s.RegisterService(&Conn_ServiceDesc, srv)
}

func _Conn_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conn_Read_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conn_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conn_Write_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conn_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conn_Close_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServer).Close(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conn_SetDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeadlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServer).SetDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conn_SetDeadline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServer).SetDeadline(ctx, req.(*SetDeadlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conn_SetReadDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeadlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServer).SetReadDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conn_SetReadDeadline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServer).SetReadDeadline(ctx, req.(*SetDeadlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conn_SetWriteDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeadlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServer).SetWriteDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conn_SetWriteDeadline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServer).SetWriteDeadline(ctx, req.(*SetDeadlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Conn_ServiceDesc is the grpc.ServiceDesc for Conn service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Conn_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "net.conn.Conn",
	HandlerType: (*ConnServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Conn_Read_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _Conn_Write_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Conn_Close_Handler,
		},
		{
			MethodName: "SetDeadline",
			Handler:    _Conn_SetDeadline_Handler,
		},
		{
			MethodName: "SetReadDeadline",
			Handler:    _Conn_SetReadDeadline_Handler,
		},
		{
			MethodName: "SetWriteDeadline",
			Handler:    _Conn_SetWriteDeadline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "net/conn/conn.proto",
}