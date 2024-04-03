package service

import (
	"context"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pb/auth"
)

type ServiceIfc interface {
	Login(ctx context.Context, req *auth.UserLoginReq) (*auth.UserLoginRes, error)
	Signup(ctx context.Context,user *auth.UserSignupReq) (*auth.UserSignupRes,error)
	Verify(ctx context.Context,req *auth.VerifyReq)(*auth.VerifyRes,error)
	// ForgotPassword(ctx context.Context,req *auth.FPreq) (*auth.FPres,error)
	// ResetPassword(ctx context.Context,req *auth.RPreq)(*auth.RPres,error)
}
