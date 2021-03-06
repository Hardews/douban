// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: user.proto

package proto

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

// DoubanClient is the client API for Douban service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DoubanClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*SuccessfulResp, error)
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*SuccessfulResp, error)
	ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*SuccessfulResp, error)
}

type doubanClient struct {
	cc grpc.ClientConnInterface
}

func NewDoubanClient(cc grpc.ClientConnInterface) DoubanClient {
	return &doubanClient{cc}
}

func (c *doubanClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*SuccessfulResp, error) {
	out := new(SuccessfulResp)
	err := c.cc.Invoke(ctx, "/server.Douban/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doubanClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*SuccessfulResp, error) {
	out := new(SuccessfulResp)
	err := c.cc.Invoke(ctx, "/server.Douban/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doubanClient) ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*SuccessfulResp, error) {
	out := new(SuccessfulResp)
	err := c.cc.Invoke(ctx, "/server.Douban/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DoubanServer is the server API for Douban service.
// All implementations must embed UnimplementedDoubanServer
// for forward compatibility
type DoubanServer interface {
	Login(context.Context, *LoginReq) (*SuccessfulResp, error)
	Register(context.Context, *RegisterReq) (*SuccessfulResp, error)
	ChangePassword(context.Context, *ChangePasswordReq) (*SuccessfulResp, error)
	mustEmbedUnimplementedDoubanServer()
}

// UnimplementedDoubanServer must be embedded to have forward compatible implementations.
type UnimplementedDoubanServer struct {
}

func (UnimplementedDoubanServer) Login(context.Context, *LoginReq) (*SuccessfulResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedDoubanServer) Register(context.Context, *RegisterReq) (*SuccessfulResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedDoubanServer) ChangePassword(context.Context, *ChangePasswordReq) (*SuccessfulResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedDoubanServer) mustEmbedUnimplementedDoubanServer() {}

// UnsafeDoubanServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DoubanServer will
// result in compilation errors.
type UnsafeDoubanServer interface {
	mustEmbedUnimplementedDoubanServer()
}

func RegisterDoubanServer(s grpc.ServiceRegistrar, srv DoubanServer) {
	s.RegisterService(&Douban_ServiceDesc, srv)
}

func _Douban_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoubanServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Douban/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoubanServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Douban_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoubanServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Douban/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoubanServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Douban_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoubanServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Douban/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoubanServer).ChangePassword(ctx, req.(*ChangePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Douban_ServiceDesc is the grpc.ServiceDesc for Douban service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Douban_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.Douban",
	HandlerType: (*DoubanServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Douban_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Douban_Register_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _Douban_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
