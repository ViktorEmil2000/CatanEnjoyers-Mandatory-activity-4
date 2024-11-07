// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: Boot/Boot.proto

package CatanEnjoyers_Mandatory_activity_4

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
	Boot_BootstrapConnect_FullMethodName = "/Boot.Boot/BootstrapConnect"
)

// BootClient is the client API for Boot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BootClient interface {
	BootstrapConnect(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ClientToBoot, BootToClient], error)
}

type bootClient struct {
	cc grpc.ClientConnInterface
}

func NewBootClient(cc grpc.ClientConnInterface) BootClient {
	return &bootClient{cc}
}

func (c *bootClient) BootstrapConnect(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ClientToBoot, BootToClient], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Boot_ServiceDesc.Streams[0], Boot_BootstrapConnect_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ClientToBoot, BootToClient]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Boot_BootstrapConnectClient = grpc.BidiStreamingClient[ClientToBoot, BootToClient]

// BootServer is the server API for Boot service.
// All implementations must embed UnimplementedBootServer
// for forward compatibility.
type BootServer interface {
	BootstrapConnect(grpc.BidiStreamingServer[ClientToBoot, BootToClient]) error
	mustEmbedUnimplementedBootServer()
}

// UnimplementedBootServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBootServer struct{}

func (UnimplementedBootServer) BootstrapConnect(grpc.BidiStreamingServer[ClientToBoot, BootToClient]) error {
	return status.Errorf(codes.Unimplemented, "method BootstrapConnect not implemented")
}
func (UnimplementedBootServer) mustEmbedUnimplementedBootServer() {}
func (UnimplementedBootServer) testEmbeddedByValue()              {}

// UnsafeBootServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BootServer will
// result in compilation errors.
type UnsafeBootServer interface {
	mustEmbedUnimplementedBootServer()
}

func RegisterBootServer(s grpc.ServiceRegistrar, srv BootServer) {
	// If the following call pancis, it indicates UnimplementedBootServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Boot_ServiceDesc, srv)
}

func _Boot_BootstrapConnect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BootServer).BootstrapConnect(&grpc.GenericServerStream[ClientToBoot, BootToClient]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Boot_BootstrapConnectServer = grpc.BidiStreamingServer[ClientToBoot, BootToClient]

// Boot_ServiceDesc is the grpc.ServiceDesc for Boot service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Boot_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Boot.Boot",
	HandlerType: (*BootServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BootstrapConnect",
			Handler:       _Boot_BootstrapConnect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "Boot/Boot.proto",
}
