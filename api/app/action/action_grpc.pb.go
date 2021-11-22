// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package action

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	CreateAction(ctx context.Context, in *CreateActionRequest, opts ...grpc.CallOption) (*Action, error)
	QueryAction(ctx context.Context, in *QueryActionRequest, opts ...grpc.CallOption) (*ActionSet, error)
	DescribeAction(ctx context.Context, in *DescribeActionRequest, opts ...grpc.CallOption) (*Action, error)
	UpdateAction(ctx context.Context, in *UpdateActionRequest, opts ...grpc.CallOption) (*Action, error)
	DeleteAction(ctx context.Context, in *DeleteActionRequest, opts ...grpc.CallOption) (*Action, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) CreateAction(ctx context.Context, in *CreateActionRequest, opts ...grpc.CallOption) (*Action, error) {
	out := new(Action)
	err := c.cc.Invoke(ctx, "/infraboard.workflow.action.Service/CreateAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryAction(ctx context.Context, in *QueryActionRequest, opts ...grpc.CallOption) (*ActionSet, error) {
	out := new(ActionSet)
	err := c.cc.Invoke(ctx, "/infraboard.workflow.action.Service/QueryAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DescribeAction(ctx context.Context, in *DescribeActionRequest, opts ...grpc.CallOption) (*Action, error) {
	out := new(Action)
	err := c.cc.Invoke(ctx, "/infraboard.workflow.action.Service/DescribeAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateAction(ctx context.Context, in *UpdateActionRequest, opts ...grpc.CallOption) (*Action, error) {
	out := new(Action)
	err := c.cc.Invoke(ctx, "/infraboard.workflow.action.Service/UpdateAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteAction(ctx context.Context, in *DeleteActionRequest, opts ...grpc.CallOption) (*Action, error) {
	out := new(Action)
	err := c.cc.Invoke(ctx, "/infraboard.workflow.action.Service/DeleteAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	CreateAction(context.Context, *CreateActionRequest) (*Action, error)
	QueryAction(context.Context, *QueryActionRequest) (*ActionSet, error)
	DescribeAction(context.Context, *DescribeActionRequest) (*Action, error)
	UpdateAction(context.Context, *UpdateActionRequest) (*Action, error)
	DeleteAction(context.Context, *DeleteActionRequest) (*Action, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) CreateAction(context.Context, *CreateActionRequest) (*Action, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAction not implemented")
}
func (UnimplementedServiceServer) QueryAction(context.Context, *QueryActionRequest) (*ActionSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAction not implemented")
}
func (UnimplementedServiceServer) DescribeAction(context.Context, *DescribeActionRequest) (*Action, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeAction not implemented")
}
func (UnimplementedServiceServer) UpdateAction(context.Context, *UpdateActionRequest) (*Action, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAction not implemented")
}
func (UnimplementedServiceServer) DeleteAction(context.Context, *DeleteActionRequest) (*Action, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAction not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_CreateAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.workflow.action.Service/CreateAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateAction(ctx, req.(*CreateActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.workflow.action.Service/QueryAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryAction(ctx, req.(*QueryActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DescribeAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DescribeAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.workflow.action.Service/DescribeAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DescribeAction(ctx, req.(*DescribeActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.workflow.action.Service/UpdateAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateAction(ctx, req.(*UpdateActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.workflow.action.Service/DeleteAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteAction(ctx, req.(*DeleteActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.workflow.action.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAction",
			Handler:    _Service_CreateAction_Handler,
		},
		{
			MethodName: "QueryAction",
			Handler:    _Service_QueryAction_Handler,
		},
		{
			MethodName: "DescribeAction",
			Handler:    _Service_DescribeAction_Handler,
		},
		{
			MethodName: "UpdateAction",
			Handler:    _Service_UpdateAction_Handler,
		},
		{
			MethodName: "DeleteAction",
			Handler:    _Service_DeleteAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/app/action/pb/action.proto",
}