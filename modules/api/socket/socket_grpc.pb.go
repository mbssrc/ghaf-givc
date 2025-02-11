// Copyright 2024 TII (SSRC) and the Ghaf contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: socket/socket.proto

package socketproxy

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
	SocketStream_TransferData_FullMethodName = "/socketproxy.SocketStream/TransferData"
)

// SocketStreamClient is the client API for SocketStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SocketStreamClient interface {
	TransferData(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[BytePacket, BytePacket], error)
}

type socketStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewSocketStreamClient(cc grpc.ClientConnInterface) SocketStreamClient {
	return &socketStreamClient{cc}
}

func (c *socketStreamClient) TransferData(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[BytePacket, BytePacket], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &SocketStream_ServiceDesc.Streams[0], SocketStream_TransferData_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[BytePacket, BytePacket]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SocketStream_TransferDataClient = grpc.BidiStreamingClient[BytePacket, BytePacket]

// SocketStreamServer is the server API for SocketStream service.
// All implementations must embed UnimplementedSocketStreamServer
// for forward compatibility.
type SocketStreamServer interface {
	TransferData(grpc.BidiStreamingServer[BytePacket, BytePacket]) error
	mustEmbedUnimplementedSocketStreamServer()
}

// UnimplementedSocketStreamServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSocketStreamServer struct{}

func (UnimplementedSocketStreamServer) TransferData(grpc.BidiStreamingServer[BytePacket, BytePacket]) error {
	return status.Errorf(codes.Unimplemented, "method TransferData not implemented")
}
func (UnimplementedSocketStreamServer) mustEmbedUnimplementedSocketStreamServer() {}
func (UnimplementedSocketStreamServer) testEmbeddedByValue()                      {}

// UnsafeSocketStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SocketStreamServer will
// result in compilation errors.
type UnsafeSocketStreamServer interface {
	mustEmbedUnimplementedSocketStreamServer()
}

func RegisterSocketStreamServer(s grpc.ServiceRegistrar, srv SocketStreamServer) {
	// If the following call pancis, it indicates UnimplementedSocketStreamServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SocketStream_ServiceDesc, srv)
}

func _SocketStream_TransferData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SocketStreamServer).TransferData(&grpc.GenericServerStream[BytePacket, BytePacket]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SocketStream_TransferDataServer = grpc.BidiStreamingServer[BytePacket, BytePacket]

// SocketStream_ServiceDesc is the grpc.ServiceDesc for SocketStream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SocketStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "socketproxy.SocketStream",
	HandlerType: (*SocketStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TransferData",
			Handler:       _SocketStream_TransferData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "socket/socket.proto",
}
