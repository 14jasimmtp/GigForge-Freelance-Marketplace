// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: pb/admin/admin.proto

package admin

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

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	AdminLogin(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
	BlockUser(ctx context.Context, in *BlockReq, opts ...grpc.CallOption) (*BlockRes, error)
	UnBlockUser(ctx context.Context, in *BlockReq, opts ...grpc.CallOption) (*BlockRes, error)
	AddSkill(ctx context.Context, in *AddSkillReq, opts ...grpc.CallOption) (*AddSkillRes, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) AdminLogin(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	out := new(LoginRes)
	err := c.cc.Invoke(ctx, "/admin.AdminService/AdminLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) BlockUser(ctx context.Context, in *BlockReq, opts ...grpc.CallOption) (*BlockRes, error) {
	out := new(BlockRes)
	err := c.cc.Invoke(ctx, "/admin.AdminService/BlockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UnBlockUser(ctx context.Context, in *BlockReq, opts ...grpc.CallOption) (*BlockRes, error) {
	out := new(BlockRes)
	err := c.cc.Invoke(ctx, "/admin.AdminService/UnBlockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AddSkill(ctx context.Context, in *AddSkillReq, opts ...grpc.CallOption) (*AddSkillRes, error) {
	out := new(AddSkillRes)
	err := c.cc.Invoke(ctx, "/admin.AdminService/AddSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	AdminLogin(context.Context, *LoginReq) (*LoginRes, error)
	BlockUser(context.Context, *BlockReq) (*BlockRes, error)
	UnBlockUser(context.Context, *BlockReq) (*BlockRes, error)
	AddSkill(context.Context, *AddSkillReq) (*AddSkillRes, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) AdminLogin(context.Context, *LoginReq) (*LoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogin not implemented")
}
func (UnimplementedAdminServiceServer) BlockUser(context.Context, *BlockReq) (*BlockRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockUser not implemented")
}
func (UnimplementedAdminServiceServer) UnBlockUser(context.Context, *BlockReq) (*BlockRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnBlockUser not implemented")
}
func (UnimplementedAdminServiceServer) AddSkill(context.Context, *AddSkillReq) (*AddSkillRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSkill not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_AdminLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/AdminLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminLogin(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_BlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).BlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/BlockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).BlockUser(ctx, req.(*BlockReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UnBlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UnBlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/UnBlockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UnBlockUser(ctx, req.(*BlockReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AddSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSkillReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AddSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/AddSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AddSkill(ctx, req.(*AddSkillReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminLogin",
			Handler:    _AdminService_AdminLogin_Handler,
		},
		{
			MethodName: "BlockUser",
			Handler:    _AdminService_BlockUser_Handler,
		},
		{
			MethodName: "UnBlockUser",
			Handler:    _AdminService_UnBlockUser_Handler,
		},
		{
			MethodName: "AddSkill",
			Handler:    _AdminService_AddSkill_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/admin/admin.proto",
}
