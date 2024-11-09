// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: p2p/P2P.proto

package p2p

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
	P2P_ClientConnect_FullMethodName   = "/p2p.P2P/ClientConnect"
	P2P_RequestResponse_FullMethodName = "/p2p.P2P/RequestResponse"
)

// P2PClient is the client API for P2P service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type P2PClient interface {
	ClientConnect(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[FromClient, FromServer], error)
	RequestResponse(ctx context.Context, in *ResquestFromClient, opts ...grpc.CallOption) (*ResponseFromServer, error)
}

type p2PClient struct {
	cc grpc.ClientConnInterface
}

func NewP2PClient(cc grpc.ClientConnInterface) P2PClient {
	return &p2PClient{cc}
}

func (c *p2PClient) ClientConnect(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[FromClient, FromServer], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &P2P_ServiceDesc.Streams[0], P2P_ClientConnect_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[FromClient, FromServer]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type P2P_ClientConnectClient = grpc.BidiStreamingClient[FromClient, FromServer]

func (c *p2PClient) RequestResponse(ctx context.Context, in *ResquestFromClient, opts ...grpc.CallOption) (*ResponseFromServer, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseFromServer)
	err := c.cc.Invoke(ctx, P2P_RequestResponse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// P2PServer is the server API for P2P service.
// All implementations must embed UnimplementedP2PServer
// for forward compatibility.
type P2PServer interface {
	ClientConnect(grpc.BidiStreamingServer[FromClient, FromServer]) error
	RequestResponse(context.Context, *ResquestFromClient) (*ResponseFromServer, error)
	mustEmbedUnimplementedP2PServer()
}

// UnimplementedP2PServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedP2PServer struct{}

func (UnimplementedP2PServer) ClientConnect(grpc.BidiStreamingServer[FromClient, FromServer]) error {
	return status.Errorf(codes.Unimplemented, "method ClientConnect not implemented")
}
func (UnimplementedP2PServer) RequestResponse(context.Context, *ResquestFromClient) (*ResponseFromServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestResponse not implemented")
}
func (UnimplementedP2PServer) mustEmbedUnimplementedP2PServer() {}
func (UnimplementedP2PServer) testEmbeddedByValue()             {}

// UnsafeP2PServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to P2PServer will
// result in compilation errors.
type UnsafeP2PServer interface {
	mustEmbedUnimplementedP2PServer()
}

func RegisterP2PServer(s grpc.ServiceRegistrar, srv P2PServer) {
	// If the following call pancis, it indicates UnimplementedP2PServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&P2P_ServiceDesc, srv)
}

func _P2P_ClientConnect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(P2PServer).ClientConnect(&grpc.GenericServerStream[FromClient, FromServer]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type P2P_ClientConnectServer = grpc.BidiStreamingServer[FromClient, FromServer]

func _P2P_RequestResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResquestFromClient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PServer).RequestResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: P2P_RequestResponse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PServer).RequestResponse(ctx, req.(*ResquestFromClient))
	}
	return interceptor(ctx, in, info, handler)
}

// P2P_ServiceDesc is the grpc.ServiceDesc for P2P service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var P2P_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "p2p.P2P",
	HandlerType: (*P2PServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestResponse",
			Handler:    _P2P_RequestResponse_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientConnect",
			Handler:       _P2P_ClientConnect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "p2p/P2P.proto",
}
