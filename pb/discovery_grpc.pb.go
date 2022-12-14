// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: discovery.proto

package pb

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

// ServiceDiscoveryClient is the client API for ServiceDiscovery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceDiscoveryClient interface {
	GetAllServices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ServiceDiscovery_GetAllServicesClient, error)
	Discovery(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (ServiceDiscovery_DiscoveryClient, error)
}

type serviceDiscoveryClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceDiscoveryClient(cc grpc.ClientConnInterface) ServiceDiscoveryClient {
	return &serviceDiscoveryClient{cc}
}

func (c *serviceDiscoveryClient) GetAllServices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ServiceDiscovery_GetAllServicesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServiceDiscovery_ServiceDesc.Streams[0], "/ServiceDiscovery/getAllServices", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceDiscoveryGetAllServicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServiceDiscovery_GetAllServicesClient interface {
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type serviceDiscoveryGetAllServicesClient struct {
	grpc.ClientStream
}

func (x *serviceDiscoveryGetAllServicesClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceDiscoveryClient) Discovery(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (ServiceDiscovery_DiscoveryClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServiceDiscovery_ServiceDesc.Streams[1], "/ServiceDiscovery/discovery", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceDiscoveryDiscoveryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServiceDiscovery_DiscoveryClient interface {
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type serviceDiscoveryDiscoveryClient struct {
	grpc.ClientStream
}

func (x *serviceDiscoveryDiscoveryClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceDiscoveryServer is the server API for ServiceDiscovery service.
// All implementations must embed UnimplementedServiceDiscoveryServer
// for forward compatibility
type ServiceDiscoveryServer interface {
	GetAllServices(*Empty, ServiceDiscovery_GetAllServicesServer) error
	Discovery(*DiscoveryRequest, ServiceDiscovery_DiscoveryServer) error
	mustEmbedUnimplementedServiceDiscoveryServer()
}

// UnimplementedServiceDiscoveryServer must be embedded to have forward compatible implementations.
type UnimplementedServiceDiscoveryServer struct {
}

func (UnimplementedServiceDiscoveryServer) GetAllServices(*Empty, ServiceDiscovery_GetAllServicesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllServices not implemented")
}
func (UnimplementedServiceDiscoveryServer) Discovery(*DiscoveryRequest, ServiceDiscovery_DiscoveryServer) error {
	return status.Errorf(codes.Unimplemented, "method Discovery not implemented")
}
func (UnimplementedServiceDiscoveryServer) mustEmbedUnimplementedServiceDiscoveryServer() {}

// UnsafeServiceDiscoveryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceDiscoveryServer will
// result in compilation errors.
type UnsafeServiceDiscoveryServer interface {
	mustEmbedUnimplementedServiceDiscoveryServer()
}

func RegisterServiceDiscoveryServer(s grpc.ServiceRegistrar, srv ServiceDiscoveryServer) {
	s.RegisterService(&ServiceDiscovery_ServiceDesc, srv)
}

func _ServiceDiscovery_GetAllServices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceDiscoveryServer).GetAllServices(m, &serviceDiscoveryGetAllServicesServer{stream})
}

type ServiceDiscovery_GetAllServicesServer interface {
	Send(*DiscoveryResponse) error
	grpc.ServerStream
}

type serviceDiscoveryGetAllServicesServer struct {
	grpc.ServerStream
}

func (x *serviceDiscoveryGetAllServicesServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ServiceDiscovery_Discovery_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DiscoveryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceDiscoveryServer).Discovery(m, &serviceDiscoveryDiscoveryServer{stream})
}

type ServiceDiscovery_DiscoveryServer interface {
	Send(*DiscoveryResponse) error
	grpc.ServerStream
}

type serviceDiscoveryDiscoveryServer struct {
	grpc.ServerStream
}

func (x *serviceDiscoveryDiscoveryServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ServiceDiscovery_ServiceDesc is the grpc.ServiceDesc for ServiceDiscovery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceDiscovery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ServiceDiscovery",
	HandlerType: (*ServiceDiscoveryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getAllServices",
			Handler:       _ServiceDiscovery_GetAllServices_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "discovery",
			Handler:       _ServiceDiscovery_Discovery_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "discovery.proto",
}
