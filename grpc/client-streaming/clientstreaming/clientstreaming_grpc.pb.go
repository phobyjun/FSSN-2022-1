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

// ClientStreamingClient is the client API for ClientStreaming service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientStreamingClient interface {
	GetServerResponse(ctx context.Context, opts ...grpc.CallOption) (ClientStreaming_GetServerResponseClient, error)
}

type clientStreamingClient struct {
	cc grpc.ClientConnInterface
}

func NewClientStreamingClient(cc grpc.ClientConnInterface) ClientStreamingClient {
	return &clientStreamingClient{cc}
}

func (c *clientStreamingClient) GetServerResponse(ctx context.Context, opts ...grpc.CallOption) (ClientStreaming_GetServerResponseClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientStreaming_ServiceDesc.Streams[0], "/bidirectional.ClientStreaming/GetServerResponse", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientStreamingGetServerResponseClient{stream}
	return x, nil
}

type ClientStreaming_GetServerResponseClient interface {
	Send(*Message) error
	CloseAndRecv() (*Number, error)
	grpc.ClientStream
}

type clientStreamingGetServerResponseClient struct {
	grpc.ClientStream
}

func (x *clientStreamingGetServerResponseClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientStreamingGetServerResponseClient) CloseAndRecv() (*Number, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Number)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientStreamingServer is the server API for ClientStreaming service.
// All implementations must embed UnimplementedClientStreamingServer
// for forward compatibility
type ClientStreamingServer interface {
	GetServerResponse(ClientStreaming_GetServerResponseServer) error
	mustEmbedUnimplementedClientStreamingServer()
}

// UnimplementedClientStreamingServer must be embedded to have forward compatible implementations.
type UnimplementedClientStreamingServer struct {
}

func (UnimplementedClientStreamingServer) GetServerResponse(ClientStreaming_GetServerResponseServer) error {
	return status.Errorf(codes.Unimplemented, "method GetServerResponse not implemented")
}
func (UnimplementedClientStreamingServer) mustEmbedUnimplementedClientStreamingServer() {}

// UnsafeClientStreamingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientStreamingServer will
// result in compilation errors.
type UnsafeClientStreamingServer interface {
	mustEmbedUnimplementedClientStreamingServer()
}

func RegisterClientStreamingServer(s grpc.ServiceRegistrar, srv ClientStreamingServer) {
	s.RegisterService(&ClientStreaming_ServiceDesc, srv)
}

func _ClientStreaming_GetServerResponse_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientStreamingServer).GetServerResponse(&clientStreamingGetServerResponseServer{stream})
}

type ClientStreaming_GetServerResponseServer interface {
	SendAndClose(*Number) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type clientStreamingGetServerResponseServer struct {
	grpc.ServerStream
}

func (x *clientStreamingGetServerResponseServer) SendAndClose(m *Number) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientStreamingGetServerResponseServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientStreaming_ServiceDesc is the grpc.ServiceDesc for ClientStreaming service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientStreaming_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bidirectional.ClientStreaming",
	HandlerType: (*ClientStreamingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetServerResponse",
			Handler:       _ClientStreaming_GetServerResponse_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "clientstreaming.proto",
}
