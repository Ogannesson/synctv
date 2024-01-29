// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: proto/provider/plugin.proto

package providerpb

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

const (
	Oauth2Plugin_Init_FullMethodName         = "/proto.Oauth2Plugin/Init"
	Oauth2Plugin_Provider_FullMethodName     = "/proto.Oauth2Plugin/Provider"
	Oauth2Plugin_NewAuthURL_FullMethodName   = "/proto.Oauth2Plugin/NewAuthURL"
	Oauth2Plugin_GetToken_FullMethodName     = "/proto.Oauth2Plugin/GetToken"
	Oauth2Plugin_RefreshToken_FullMethodName = "/proto.Oauth2Plugin/RefreshToken"
	Oauth2Plugin_GetUserInfo_FullMethodName  = "/proto.Oauth2Plugin/GetUserInfo"
)

// Oauth2PluginClient is the client API for Oauth2Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Oauth2PluginClient interface {
	Init(ctx context.Context, in *InitReq, opts ...grpc.CallOption) (*Enpty, error)
	Provider(ctx context.Context, in *Enpty, opts ...grpc.CallOption) (*ProviderResp, error)
	NewAuthURL(ctx context.Context, in *NewAuthURLReq, opts ...grpc.CallOption) (*NewAuthURLResp, error)
	GetToken(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*Token, error)
	RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*RefreshTokenResp, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
}

type oauth2PluginClient struct {
	cc grpc.ClientConnInterface
}

func NewOauth2PluginClient(cc grpc.ClientConnInterface) Oauth2PluginClient {
	return &oauth2PluginClient{cc}
}

func (c *oauth2PluginClient) Init(ctx context.Context, in *InitReq, opts ...grpc.CallOption) (*Enpty, error) {
	out := new(Enpty)
	err := c.cc.Invoke(ctx, Oauth2Plugin_Init_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2PluginClient) Provider(ctx context.Context, in *Enpty, opts ...grpc.CallOption) (*ProviderResp, error) {
	out := new(ProviderResp)
	err := c.cc.Invoke(ctx, Oauth2Plugin_Provider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2PluginClient) NewAuthURL(ctx context.Context, in *NewAuthURLReq, opts ...grpc.CallOption) (*NewAuthURLResp, error) {
	out := new(NewAuthURLResp)
	err := c.cc.Invoke(ctx, Oauth2Plugin_NewAuthURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2PluginClient) GetToken(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, Oauth2Plugin_GetToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2PluginClient) RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*RefreshTokenResp, error) {
	out := new(RefreshTokenResp)
	err := c.cc.Invoke(ctx, Oauth2Plugin_RefreshToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2PluginClient) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	out := new(GetUserInfoResp)
	err := c.cc.Invoke(ctx, Oauth2Plugin_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Oauth2PluginServer is the server API for Oauth2Plugin service.
// All implementations must embed UnimplementedOauth2PluginServer
// for forward compatibility
type Oauth2PluginServer interface {
	Init(context.Context, *InitReq) (*Enpty, error)
	Provider(context.Context, *Enpty) (*ProviderResp, error)
	NewAuthURL(context.Context, *NewAuthURLReq) (*NewAuthURLResp, error)
	GetToken(context.Context, *GetTokenReq) (*Token, error)
	RefreshToken(context.Context, *RefreshTokenReq) (*RefreshTokenResp, error)
	GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error)
	mustEmbedUnimplementedOauth2PluginServer()
}

// UnimplementedOauth2PluginServer must be embedded to have forward compatible implementations.
type UnimplementedOauth2PluginServer struct {
}

func (UnimplementedOauth2PluginServer) Init(context.Context, *InitReq) (*Enpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedOauth2PluginServer) Provider(context.Context, *Enpty) (*ProviderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Provider not implemented")
}
func (UnimplementedOauth2PluginServer) NewAuthURL(context.Context, *NewAuthURLReq) (*NewAuthURLResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAuthURL not implemented")
}
func (UnimplementedOauth2PluginServer) GetToken(context.Context, *GetTokenReq) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedOauth2PluginServer) RefreshToken(context.Context, *RefreshTokenReq) (*RefreshTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedOauth2PluginServer) GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedOauth2PluginServer) mustEmbedUnimplementedOauth2PluginServer() {}

// UnsafeOauth2PluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Oauth2PluginServer will
// result in compilation errors.
type UnsafeOauth2PluginServer interface {
	mustEmbedUnimplementedOauth2PluginServer()
}

func RegisterOauth2PluginServer(s grpc.ServiceRegistrar, srv Oauth2PluginServer) {
	s.RegisterService(&Oauth2Plugin_ServiceDesc, srv)
}

func _Oauth2Plugin_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2PluginServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Oauth2Plugin_Init_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2PluginServer).Init(ctx, req.(*InitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2Plugin_Provider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Enpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2PluginServer).Provider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Oauth2Plugin_Provider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2PluginServer).Provider(ctx, req.(*Enpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2Plugin_NewAuthURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAuthURLReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2PluginServer).NewAuthURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Oauth2Plugin_NewAuthURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2PluginServer).NewAuthURL(ctx, req.(*NewAuthURLReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2Plugin_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2PluginServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Oauth2Plugin_GetToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2PluginServer).GetToken(ctx, req.(*GetTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2Plugin_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2PluginServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Oauth2Plugin_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2PluginServer).RefreshToken(ctx, req.(*RefreshTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2Plugin_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2PluginServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Oauth2Plugin_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2PluginServer).GetUserInfo(ctx, req.(*GetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Oauth2Plugin_ServiceDesc is the grpc.ServiceDesc for Oauth2Plugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Oauth2Plugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Oauth2Plugin",
	HandlerType: (*Oauth2PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _Oauth2Plugin_Init_Handler,
		},
		{
			MethodName: "Provider",
			Handler:    _Oauth2Plugin_Provider_Handler,
		},
		{
			MethodName: "NewAuthURL",
			Handler:    _Oauth2Plugin_NewAuthURL_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _Oauth2Plugin_GetToken_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _Oauth2Plugin_RefreshToken_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _Oauth2Plugin_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/provider/plugin.proto",
}
