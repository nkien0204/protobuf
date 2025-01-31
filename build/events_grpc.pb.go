// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: events.proto

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

// RouteGuideClient is the client API for RouteGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteGuideClient interface {
	// Bidirectional streaming RPC
	PingPong(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_PingPongClient, error)
}

type routeGuideClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteGuideClient(cc grpc.ClientConnInterface) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) PingPong(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_PingPongClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteGuide_ServiceDesc.Streams[0], "/events.RouteGuide/PingPong", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuidePingPongClient{stream}
	return x, nil
}

type RouteGuide_PingPongClient interface {
	Send(*RpcHeartBeatEvent) error
	Recv() (*RpcHeartBeatEvent, error)
	grpc.ClientStream
}

type routeGuidePingPongClient struct {
	grpc.ClientStream
}

func (x *routeGuidePingPongClient) Send(m *RpcHeartBeatEvent) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuidePingPongClient) Recv() (*RpcHeartBeatEvent, error) {
	m := new(RpcHeartBeatEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuideServer is the server API for RouteGuide service.
// All implementations must embed UnimplementedRouteGuideServer
// for forward compatibility
type RouteGuideServer interface {
	// Bidirectional streaming RPC
	PingPong(RouteGuide_PingPongServer) error
	mustEmbedUnimplementedRouteGuideServer()
}

// UnimplementedRouteGuideServer must be embedded to have forward compatible implementations.
type UnimplementedRouteGuideServer struct {
}

func (UnimplementedRouteGuideServer) PingPong(RouteGuide_PingPongServer) error {
	return status.Errorf(codes.Unimplemented, "method PingPong not implemented")
}
func (UnimplementedRouteGuideServer) mustEmbedUnimplementedRouteGuideServer() {}

// UnsafeRouteGuideServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteGuideServer will
// result in compilation errors.
type UnsafeRouteGuideServer interface {
	mustEmbedUnimplementedRouteGuideServer()
}

func RegisterRouteGuideServer(s grpc.ServiceRegistrar, srv RouteGuideServer) {
	s.RegisterService(&RouteGuide_ServiceDesc, srv)
}

func _RouteGuide_PingPong_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).PingPong(&routeGuidePingPongServer{stream})
}

type RouteGuide_PingPongServer interface {
	Send(*RpcHeartBeatEvent) error
	Recv() (*RpcHeartBeatEvent, error)
	grpc.ServerStream
}

type routeGuidePingPongServer struct {
	grpc.ServerStream
}

func (x *routeGuidePingPongServer) Send(m *RpcHeartBeatEvent) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuidePingPongServer) Recv() (*RpcHeartBeatEvent, error) {
	m := new(RpcHeartBeatEvent)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuide_ServiceDesc is the grpc.ServiceDesc for RouteGuide service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteGuide_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "events.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PingPong",
			Handler:       _RouteGuide_PingPong_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "events.proto",
}
