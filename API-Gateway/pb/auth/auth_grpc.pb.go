// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: pb/auth/auth.proto

package auth

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Login(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginRes, error)
	Signup(ctx context.Context, in *UserSignupReq, opts ...grpc.CallOption) (*UserSignupRes, error)
	Verify(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyRes, error)
	ForgotPassword(ctx context.Context, in *FPreq, opts ...grpc.CallOption) (*FPres, error)
	ResetPassword(ctx context.Context, in *RPreq, opts ...grpc.CallOption) (*RPres, error)
	AddProfileDescription(ctx context.Context, in *APDReq, opts ...grpc.CallOption) (*APDRes, error)
	UpdateProfileDescription(ctx context.Context, in *UPDReq, opts ...grpc.CallOption) (*UPDRes, error)
	AddEducation(ctx context.Context, in *AddEducationReq, opts ...grpc.CallOption) (*AddEducationRes, error)
	UpdateEducation(ctx context.Context, in *UpdateEducationReq, opts ...grpc.CallOption) (*UpdateEducationRes, error)
	DeleteEducation(ctx context.Context, in *DeleteEducationReq, opts ...grpc.CallOption) (*DeleteEducationRes, error)
	GetProfile(ctx context.Context, in *GetProfileReq, opts ...grpc.CallOption) (*GetProfileRes, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginRes, error) {
	out := new(UserLoginRes)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Signup(ctx context.Context, in *UserSignupReq, opts ...grpc.CallOption) (*UserSignupRes, error) {
	out := new(UserSignupRes)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Verify(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyRes, error) {
	out := new(VerifyRes)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ForgotPassword(ctx context.Context, in *FPreq, opts ...grpc.CallOption) (*FPres, error) {
	out := new(FPres)
	err := c.cc.Invoke(ctx, "/auth.AuthService/ForgotPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ResetPassword(ctx context.Context, in *RPreq, opts ...grpc.CallOption) (*RPres, error) {
	out := new(RPres)
	err := c.cc.Invoke(ctx, "/auth.AuthService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AddProfileDescription(ctx context.Context, in *APDReq, opts ...grpc.CallOption) (*APDRes, error) {
	out := new(APDRes)
	err := c.cc.Invoke(ctx, "/auth.AuthService/AddProfileDescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateProfileDescription(ctx context.Context, in *UPDReq, opts ...grpc.CallOption) (*UPDRes, error) {
	out := new(UPDRes)
	err := c.cc.Invoke(ctx, "/auth.AuthService/UpdateProfileDescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AddEducation(ctx context.Context, in *AddEducationReq, opts ...grpc.CallOption) (*AddEducationRes, error) {
	out := new(AddEducationRes)
	err := c.cc.Invoke(ctx, "/auth.AuthService/AddEducation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateEducation(ctx context.Context, in *UpdateEducationReq, opts ...grpc.CallOption) (*UpdateEducationRes, error) {
	out := new(UpdateEducationRes)
	err := c.cc.Invoke(ctx, "/auth.AuthService/UpdateEducation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteEducation(ctx context.Context, in *DeleteEducationReq, opts ...grpc.CallOption) (*DeleteEducationRes, error) {
	out := new(DeleteEducationRes)
	err := c.cc.Invoke(ctx, "/auth.AuthService/DeleteEducation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetProfile(ctx context.Context, in *GetProfileReq, opts ...grpc.CallOption) (*GetProfileRes, error) {
	out := new(GetProfileRes)
	err := c.cc.Invoke(ctx, "/auth.AuthService/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	Login(context.Context, *UserLoginReq) (*UserLoginRes, error)
	Signup(context.Context, *UserSignupReq) (*UserSignupRes, error)
	Verify(context.Context, *VerifyReq) (*VerifyRes, error)
	ForgotPassword(context.Context, *FPreq) (*FPres, error)
	ResetPassword(context.Context, *RPreq) (*RPres, error)
	AddProfileDescription(context.Context, *APDReq) (*APDRes, error)
	UpdateProfileDescription(context.Context, *UPDReq) (*UPDRes, error)
	AddEducation(context.Context, *AddEducationReq) (*AddEducationRes, error)
	UpdateEducation(context.Context, *UpdateEducationReq) (*UpdateEducationRes, error)
	DeleteEducation(context.Context, *DeleteEducationReq) (*DeleteEducationRes, error)
	GetProfile(context.Context, *GetProfileReq) (*GetProfileRes, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Login(context.Context, *UserLoginReq) (*UserLoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) Signup(context.Context, *UserSignupReq) (*UserSignupRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedAuthServiceServer) Verify(context.Context, *VerifyReq) (*VerifyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedAuthServiceServer) ForgotPassword(context.Context, *FPreq) (*FPres, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgotPassword not implemented")
}
func (UnimplementedAuthServiceServer) ResetPassword(context.Context, *RPreq) (*RPres, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedAuthServiceServer) AddProfileDescription(context.Context, *APDReq) (*APDRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProfileDescription not implemented")
}
func (UnimplementedAuthServiceServer) UpdateProfileDescription(context.Context, *UPDReq) (*UPDRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfileDescription not implemented")
}
func (UnimplementedAuthServiceServer) AddEducation(context.Context, *AddEducationReq) (*AddEducationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEducation not implemented")
}
func (UnimplementedAuthServiceServer) UpdateEducation(context.Context, *UpdateEducationReq) (*UpdateEducationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEducation not implemented")
}
func (UnimplementedAuthServiceServer) DeleteEducation(context.Context, *DeleteEducationReq) (*DeleteEducationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEducation not implemented")
}
func (UnimplementedAuthServiceServer) GetProfile(context.Context, *GetProfileReq) (*GetProfileRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*UserLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Signup(ctx, req.(*UserSignupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Verify(ctx, req.(*VerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FPreq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/ForgotPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ForgotPassword(ctx, req.(*FPreq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPreq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ResetPassword(ctx, req.(*RPreq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AddProfileDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AddProfileDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/AddProfileDescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AddProfileDescription(ctx, req.(*APDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateProfileDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UPDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateProfileDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/UpdateProfileDescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateProfileDescription(ctx, req.(*UPDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AddEducation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEducationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AddEducation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/AddEducation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AddEducation(ctx, req.(*AddEducationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateEducation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEducationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateEducation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/UpdateEducation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateEducation(ctx, req.(*UpdateEducationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeleteEducation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEducationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeleteEducation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/DeleteEducation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeleteEducation(ctx, req.(*DeleteEducationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetProfile(ctx, req.(*GetProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Signup",
			Handler:    _AuthService_Signup_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _AuthService_Verify_Handler,
		},
		{
			MethodName: "ForgotPassword",
			Handler:    _AuthService_ForgotPassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _AuthService_ResetPassword_Handler,
		},
		{
			MethodName: "AddProfileDescription",
			Handler:    _AuthService_AddProfileDescription_Handler,
		},
		{
			MethodName: "UpdateProfileDescription",
			Handler:    _AuthService_UpdateProfileDescription_Handler,
		},
		{
			MethodName: "AddEducation",
			Handler:    _AuthService_AddEducation_Handler,
		},
		{
			MethodName: "UpdateEducation",
			Handler:    _AuthService_UpdateEducation_Handler,
		},
		{
			MethodName: "DeleteEducation",
			Handler:    _AuthService_DeleteEducation_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _AuthService_GetProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/auth/auth.proto",
}
