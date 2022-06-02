// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package __

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

// ServerStreamingClient is the client API for ServerStreaming service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerStreamingClient interface {
	// A Client streaming RPC.
	GetServerResponse(ctx context.Context, in *Number, opts ...grpc.CallOption) (ServerStreaming_GetServerResponseClient, error)
}

type serverStreamingClient struct {
	cc grpc.ClientConnInterface
}

func NewServerStreamingClient(cc grpc.ClientConnInterface) ServerStreamingClient {
	return &serverStreamingClient{cc}
}

func (c *serverStreamingClient) GetServerResponse(ctx context.Context, in *Number, opts ...grpc.CallOption) (ServerStreaming_GetServerResponseClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServerStreaming_ServiceDesc.Streams[0], "/serverstreaming.ServerStreaming/GetServerResponse", opts...)
	if err != nil {
		return nil, err
	}
	x := &serverStreamingGetServerResponseClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServerStreaming_GetServerResponseClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type serverStreamingGetServerResponseClient struct {
	grpc.ClientStream
}

func (x *serverStreamingGetServerResponseClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerStreamingServer is the server API for ServerStreaming service.
// All implementations must embed UnimplementedServerStreamingServer
// for forward compatibility
type ServerStreamingServer interface {
	// A Client streaming RPC.
	GetServerResponse(*Number, ServerStreaming_GetServerResponseServer) error
	mustEmbedUnimplementedServerStreamingServer()
}

// UnimplementedServerStreamingServer must be embedded to have forward compatible implementations.
type UnimplementedServerStreamingServer struct {
}

func (UnimplementedServerStreamingServer) GetServerResponse(*Number, ServerStreaming_GetServerResponseServer) error {
	return status.Errorf(codes.Unimplemented, "method GetServerResponse not implemented")
}
func (UnimplementedServerStreamingServer) mustEmbedUnimplementedServerStreamingServer() {}

// UnsafeServerStreamingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerStreamingServer will
// result in compilation errors.
type UnsafeServerStreamingServer interface {
	mustEmbedUnimplementedServerStreamingServer()
}

func RegisterServerStreamingServer(s grpc.ServiceRegistrar, srv ServerStreamingServer) {
	s.RegisterService(&ServerStreaming_ServiceDesc, srv)
}

func _ServerStreaming_GetServerResponse_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Number)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServerStreamingServer).GetServerResponse(m, &serverStreamingGetServerResponseServer{stream})
}

type ServerStreaming_GetServerResponseServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type serverStreamingGetServerResponseServer struct {
	grpc.ServerStream
}

func (x *serverStreamingGetServerResponseServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

// ServerStreaming_ServiceDesc is the grpc.ServiceDesc for ServerStreaming service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerStreaming_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "serverstreaming.ServerStreaming",
	HandlerType: (*ServerStreamingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetServerResponse",
			Handler:       _ServerStreaming_GetServerResponse_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "serverstreaming.proto",
}
