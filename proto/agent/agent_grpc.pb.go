// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package agent

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PluginClient is the client API for Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginClient interface {
	GetEnvironmentSchema(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetEnvironmentSchema_Response, error)
	GetQuerySchema(ctx context.Context, in *GetQuerySchema_Request, opts ...grpc.CallOption) (*GetQuerySchema_Response, error)
	ValidEnvironmentConfig(ctx context.Context, in *ValidEnvironmentConfig_Request, opts ...grpc.CallOption) (*ValidEnvironmentConfig_Response, error)
	ValidQueryConfig(ctx context.Context, in *ValidQueryConfig_Request, opts ...grpc.CallOption) (*ValidQueryConfig_Response, error)
	Recover(ctx context.Context, opts ...grpc.CallOption) (Plugin_RecoverClient, error)
	Query(ctx context.Context, in *Query_Request, opts ...grpc.CallOption) (*Query_Response, error)
}

type pluginClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginClient(cc grpc.ClientConnInterface) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) GetEnvironmentSchema(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetEnvironmentSchema_Response, error) {
	out := new(GetEnvironmentSchema_Response)
	err := c.cc.Invoke(ctx, "/agent.Plugin/GetEnvironmentSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) GetQuerySchema(ctx context.Context, in *GetQuerySchema_Request, opts ...grpc.CallOption) (*GetQuerySchema_Response, error) {
	out := new(GetQuerySchema_Response)
	err := c.cc.Invoke(ctx, "/agent.Plugin/GetQuerySchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) ValidEnvironmentConfig(ctx context.Context, in *ValidEnvironmentConfig_Request, opts ...grpc.CallOption) (*ValidEnvironmentConfig_Response, error) {
	out := new(ValidEnvironmentConfig_Response)
	err := c.cc.Invoke(ctx, "/agent.Plugin/ValidEnvironmentConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) ValidQueryConfig(ctx context.Context, in *ValidQueryConfig_Request, opts ...grpc.CallOption) (*ValidQueryConfig_Response, error) {
	out := new(ValidQueryConfig_Response)
	err := c.cc.Invoke(ctx, "/agent.Plugin/ValidQueryConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Recover(ctx context.Context, opts ...grpc.CallOption) (Plugin_RecoverClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Plugin_serviceDesc.Streams[0], "/agent.Plugin/Recover", opts...)
	if err != nil {
		return nil, err
	}
	x := &pluginRecoverClient{stream}
	return x, nil
}

type Plugin_RecoverClient interface {
	Send(*Recover_Request) error
	CloseAndRecv() (*Recover_Response, error)
	grpc.ClientStream
}

type pluginRecoverClient struct {
	grpc.ClientStream
}

func (x *pluginRecoverClient) Send(m *Recover_Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pluginRecoverClient) CloseAndRecv() (*Recover_Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Recover_Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pluginClient) Query(ctx context.Context, in *Query_Request, opts ...grpc.CallOption) (*Query_Response, error) {
	out := new(Query_Response)
	err := c.cc.Invoke(ctx, "/agent.Plugin/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServer is the server API for Plugin service.
// All implementations must embed UnimplementedPluginServer
// for forward compatibility
type PluginServer interface {
	GetEnvironmentSchema(context.Context, *empty.Empty) (*GetEnvironmentSchema_Response, error)
	GetQuerySchema(context.Context, *GetQuerySchema_Request) (*GetQuerySchema_Response, error)
	ValidEnvironmentConfig(context.Context, *ValidEnvironmentConfig_Request) (*ValidEnvironmentConfig_Response, error)
	ValidQueryConfig(context.Context, *ValidQueryConfig_Request) (*ValidQueryConfig_Response, error)
	Recover(Plugin_RecoverServer) error
	Query(context.Context, *Query_Request) (*Query_Response, error)
	mustEmbedUnimplementedPluginServer()
}

// UnimplementedPluginServer must be embedded to have forward compatible implementations.
type UnimplementedPluginServer struct {
}

func (*UnimplementedPluginServer) GetEnvironmentSchema(context.Context, *empty.Empty) (*GetEnvironmentSchema_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEnvironmentSchema not implemented")
}
func (*UnimplementedPluginServer) GetQuerySchema(context.Context, *GetQuerySchema_Request) (*GetQuerySchema_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuerySchema not implemented")
}
func (*UnimplementedPluginServer) ValidEnvironmentConfig(context.Context, *ValidEnvironmentConfig_Request) (*ValidEnvironmentConfig_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidEnvironmentConfig not implemented")
}
func (*UnimplementedPluginServer) ValidQueryConfig(context.Context, *ValidQueryConfig_Request) (*ValidQueryConfig_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidQueryConfig not implemented")
}
func (*UnimplementedPluginServer) Recover(Plugin_RecoverServer) error {
	return status.Errorf(codes.Unimplemented, "method Recover not implemented")
}
func (*UnimplementedPluginServer) Query(context.Context, *Query_Request) (*Query_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (*UnimplementedPluginServer) mustEmbedUnimplementedPluginServer() {}

func RegisterPluginServer(s *grpc.Server, srv PluginServer) {
	s.RegisterService(&_Plugin_serviceDesc, srv)
}

func _Plugin_GetEnvironmentSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).GetEnvironmentSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Plugin/GetEnvironmentSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).GetEnvironmentSchema(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_GetQuerySchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuerySchema_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).GetQuerySchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Plugin/GetQuerySchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).GetQuerySchema(ctx, req.(*GetQuerySchema_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_ValidEnvironmentConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidEnvironmentConfig_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).ValidEnvironmentConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Plugin/ValidEnvironmentConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).ValidEnvironmentConfig(ctx, req.(*ValidEnvironmentConfig_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_ValidQueryConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidQueryConfig_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).ValidQueryConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Plugin/ValidQueryConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).ValidQueryConfig(ctx, req.(*ValidQueryConfig_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Recover_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PluginServer).Recover(&pluginRecoverServer{stream})
}

type Plugin_RecoverServer interface {
	SendAndClose(*Recover_Response) error
	Recv() (*Recover_Request, error)
	grpc.ServerStream
}

type pluginRecoverServer struct {
	grpc.ServerStream
}

func (x *pluginRecoverServer) SendAndClose(m *Recover_Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pluginRecoverServer) Recv() (*Recover_Request, error) {
	m := new(Recover_Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Plugin_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Plugin/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Query(ctx, req.(*Query_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Plugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agent.Plugin",
	HandlerType: (*PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEnvironmentSchema",
			Handler:    _Plugin_GetEnvironmentSchema_Handler,
		},
		{
			MethodName: "GetQuerySchema",
			Handler:    _Plugin_GetQuerySchema_Handler,
		},
		{
			MethodName: "ValidEnvironmentConfig",
			Handler:    _Plugin_ValidEnvironmentConfig_Handler,
		},
		{
			MethodName: "ValidQueryConfig",
			Handler:    _Plugin_ValidQueryConfig_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Plugin_Query_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Recover",
			Handler:       _Plugin_Recover_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/agent/agent.proto",
}
